package event

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/donation/backend-end/abi"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/db"
	"github.com/haitao-sun03/donation/backend-end/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据编码函数
func encodeDonationData(donor common.Address, amount *big.Int) []byte {
	// abi := getContractABI("CampaignRelContract")
	// 定义 ABI 参数类型
	addressType, _ := ethabi.NewType("address", "", nil)
	uint256Type, _ := ethabi.NewType("uint256", "", nil)
	arguments := ethabi.Arguments{
		{Type: addressType},
		{Type: uint256Type},
	}

	data, err := arguments.Pack(donor, amount)
	if err != nil {
		panic(fmt.Sprintf("ABI Pack error: %v", err))
	}
	return data
}

// TestSetup contains common resources for tests
type TestSetup struct {
	DB      *sqlmock.Sqlmock
	MockDB  *db.GormDB
	Handler *DonationHandler
}

// SetupTest initializes common test resources
func SetupTest(t *testing.T) (*TestSetup, func()) {
	config.Init(true)
	_db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("some_version"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: _db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)
	assert.NotNil(t, gormDB)

	mockDB := &db.GormDB{InnerDB: gormDB}
	handler := NewDonationHandler(mockDB)

	cleanup := func() {
		_db.Close()
	}

	return &TestSetup{
		DB:      &mock,
		MockDB:  mockDB,
		Handler: handler,
	}, cleanup
}

func TestDonationHandler_FirstDonation(t *testing.T) {
	setup, cleanup := SetupTest(t)
	defer cleanup()

	campaignId := "123"
	donor := "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	var amount int64 = 100
	// 模拟事件日志
	vLog := types.Log{
		Topics: []common.Hash{
			crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")),
			common.BigToHash(big.NewInt(123)),
		},
		Data: encodeDonationData(common.HexToAddress(donor), big.NewInt(amount)),
	}
	mock := *setup.DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM `campaigns` WHERE campaign_id = \\? AND `campaigns`.`deleted_at` IS NULL ORDER BY `campaigns`.`id` LIMIT \\? FOR UPDATE").
		WithArgs(campaignId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"campaign_id", "total_donated"}).
			AddRow(campaignId, strconv.FormatInt(amount, 10)))
	mock.ExpectExec("UPDATE `campaigns` SET `total_donated`=\\?,`updated_at`=\\? WHERE campaign_id = \\? AND `campaigns`.`deleted_at` IS NULL").
		WithArgs("200", sqlmock.AnyArg(), campaignId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT \\* FROM `donation_records` WHERE campaign_id = \\? AND donor = \\? AND `donation_records`.`deleted_at` IS NULL ORDER BY `donation_records`.`id` LIMIT \\?").
		WithArgs(campaignId, donor, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	campaignIdInt, _ := strconv.Atoi(campaignId)

	// 已经捐赠过，不需要再次设置角色
	mock.ExpectQuery("SELECT \\* FROM `user_activity_roles` WHERE address = \\? AND campaign_id = \\? AND role = \\? ORDER BY `user_activity_roles`.`address` LIMIT \\?").
		WithArgs(donor, campaignIdInt, "donor", 1).
		WillReturnRows(sqlmock.NewRows([]string{"address", "campaign_id", "role"}).
			AddRow(
				donor,           // address
				campaignIdInt,   // campaign_id
				model.DonorRole, // role
			))

	mock.ExpectExec("INSERT INTO `donation_records` \\(`created_at`,`updated_at`,`deleted_at`,`campaign_id`,`donor`,`amount`,`is_refund`,`mint_level`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), campaignId, donor, strconv.FormatInt(amount, 10), 0, "").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	handler := setup.Handler
	data, err := handler.Parse(vLog)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	// 确保 data 是 *abi.DonationManageDonate 类型
	donationData, ok := data.(*abi.DonationManageDonate)
	assert.True(t, ok, "data should be of type *abi.DonationManageDonate")
	assert.Equal(t, donationData.Amount, big.NewInt(amount))
	assert.Equal(t, donationData.Donater, common.HexToAddress(donor))
	str, _ := new(big.Int).SetString(campaignId, 10)
	assert.Equal(t, donationData.Id, str)
	err = handler.Handle(donationData)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDonationHandler_NotFirstDonation(t *testing.T) {
	setup, cleanup := SetupTest(t)
	defer cleanup()

	campaignId := "123"
	donor := "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	var campaignTotalAmount int64 = 10000
	var prepareDonateAmount int64 = 100
	haveBeenDonoteAmountOfThisDonator := "5000"
	// 模拟事件日志
	vLog := types.Log{
		Topics: []common.Hash{
			crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")),
			common.BigToHash(big.NewInt(123)),
		},
		Data: encodeDonationData(common.HexToAddress(donor), big.NewInt(prepareDonateAmount)),
	}
	mock := *setup.DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM `campaigns` WHERE campaign_id = \\? AND `campaigns`.`deleted_at` IS NULL ORDER BY `campaigns`.`id` LIMIT \\? FOR UPDATE").
		WithArgs(campaignId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"campaign_id", "total_donated"}).
			AddRow(campaignId, strconv.FormatInt(campaignTotalAmount, 10)))

	campaignTotalAmountBig := big.NewInt(campaignTotalAmount)
	campaignTotal := new(big.Int).Add(campaignTotalAmountBig, big.NewInt(prepareDonateAmount))

	mock.ExpectExec("UPDATE `campaigns` SET `total_donated`=\\?,`updated_at`=\\? WHERE campaign_id = \\? AND `campaigns`.`deleted_at` IS NULL").
		WithArgs(campaignTotal.String(), sqlmock.AnyArg(), campaignId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT \\* FROM `donation_records` WHERE campaign_id = \\? AND donor = \\? AND `donation_records`.`deleted_at` IS NULL ORDER BY `donation_records`.`id` LIMIT \\?").
		WithArgs(campaignId, donor, 1).
		WillReturnRows(sqlmock.NewRows([]string{"campaign_id", "donor", "amount"}).
			AddRow(campaignId, donor, haveBeenDonoteAmountOfThisDonator))

	campaignIdInt, _ := strconv.Atoi(campaignId)

	// 已经捐赠过，不需要再次设置角色
	mock.ExpectQuery("SELECT \\* FROM `user_activity_roles` WHERE address = \\? AND campaign_id = \\? AND role = \\? ORDER BY `user_activity_roles`.`address` LIMIT \\?").
		WithArgs(donor, campaignIdInt, "donor", 1).
		WillReturnRows(sqlmock.NewRows([]string{"address", "campaign_id", "role"}).
			AddRow(
				donor,           // address
				campaignIdInt,   // campaign_id
				model.DonorRole, // role
			))

	oldAmount, _ := new(big.Int).SetString(haveBeenDonoteAmountOfThisDonator, 10)
	totalAomunt := new(big.Int).Add(oldAmount, big.NewInt(prepareDonateAmount))
	mock.ExpectExec("UPDATE `donation_records` SET `amount`=\\?,`updated_at`=\\? WHERE campaign_id = \\? AND donor = \\? AND `donation_records`.`deleted_at` IS NULL").
		WithArgs(totalAomunt.String(), sqlmock.AnyArg(), campaignId, donor).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	handler := setup.Handler
	data, err := handler.Parse(vLog)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	// 确保 data 是 *abi.DonationManageDonate 类型
	donationData, ok := data.(*abi.DonationManageDonate)
	assert.True(t, ok, "data should be of type *abi.DonationManageDonate")
	assert.Equal(t, donationData.Amount, big.NewInt(prepareDonateAmount))
	assert.Equal(t, donationData.Donater, common.HexToAddress(donor))
	str, _ := new(big.Int).SetString(campaignId, 10)
	assert.Equal(t, donationData.Id, str)
	err = handler.Handle(donationData)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
