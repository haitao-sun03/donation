// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DonationManageMetaData contains all meta data concerning the DonationManage contract.
var DonationManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ActiveCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CancellCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CompletedCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"donater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refunder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"starter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDonationsManageContract.CampaignNature\",\"name\":\"nature\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDonationsManageContract.Beneficiary\",\"name\":\"beneficiary\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"purpose\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"expectedImpact\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumDonationsManageContract.CampaignStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StartCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"activeCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDonated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"enumDonationsManageContract.CampaignNature\",\"name\":\"nature\",\"type\":\"uint8\"},{\"internalType\":\"enumDonationsManageContract.Beneficiary\",\"name\":\"beneficiary\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"purpose\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"expectedImpact\",\"type\":\"string\"},{\"internalType\":\"enumDonationsManageContract.CampaignStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"starter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignsToCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"cancellCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"completedCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"enumDonationsManageContract.CampaignNature\",\"name\":\"_nature\",\"type\":\"uint8\"},{\"internalType\":\"enumDonationsManageContract.Beneficiary\",\"name\":\"_beneficiary\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_purpose\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_expectedImpact\",\"type\":\"string\"}],\"name\":\"createCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"getCampaignDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDonated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"enumDonationsManageContract.CampaignStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"starter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"refundToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DonationManageABI is the input ABI used to generate the binding from.
// Deprecated: Use DonationManageMetaData.ABI instead.
var DonationManageABI = DonationManageMetaData.ABI

// DonationManage is an auto generated Go binding around an Ethereum contract.
type DonationManage struct {
	DonationManageCaller     // Read-only binding to the contract
	DonationManageTransactor // Write-only binding to the contract
	DonationManageFilterer   // Log filterer for contract events
}

// DonationManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DonationManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DonationManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DonationManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DonationManageSession struct {
	Contract     *DonationManage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DonationManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DonationManageCallerSession struct {
	Contract *DonationManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DonationManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DonationManageTransactorSession struct {
	Contract     *DonationManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DonationManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DonationManageRaw struct {
	Contract *DonationManage // Generic contract binding to access the raw methods on
}

// DonationManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DonationManageCallerRaw struct {
	Contract *DonationManageCaller // Generic read-only contract binding to access the raw methods on
}

// DonationManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DonationManageTransactorRaw struct {
	Contract *DonationManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDonationManage creates a new instance of DonationManage, bound to a specific deployed contract.
func NewDonationManage(address common.Address, backend bind.ContractBackend) (*DonationManage, error) {
	contract, err := bindDonationManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DonationManage{DonationManageCaller: DonationManageCaller{contract: contract}, DonationManageTransactor: DonationManageTransactor{contract: contract}, DonationManageFilterer: DonationManageFilterer{contract: contract}}, nil
}

// NewDonationManageCaller creates a new read-only instance of DonationManage, bound to a specific deployed contract.
func NewDonationManageCaller(address common.Address, caller bind.ContractCaller) (*DonationManageCaller, error) {
	contract, err := bindDonationManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DonationManageCaller{contract: contract}, nil
}

// NewDonationManageTransactor creates a new write-only instance of DonationManage, bound to a specific deployed contract.
func NewDonationManageTransactor(address common.Address, transactor bind.ContractTransactor) (*DonationManageTransactor, error) {
	contract, err := bindDonationManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DonationManageTransactor{contract: contract}, nil
}

// NewDonationManageFilterer creates a new log filterer instance of DonationManage, bound to a specific deployed contract.
func NewDonationManageFilterer(address common.Address, filterer bind.ContractFilterer) (*DonationManageFilterer, error) {
	contract, err := bindDonationManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DonationManageFilterer{contract: contract}, nil
}

// bindDonationManage binds a generic wrapper to an already deployed contract.
func bindDonationManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DonationManageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DonationManage *DonationManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonationManage.Contract.DonationManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DonationManage *DonationManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonationManage.Contract.DonationManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DonationManage *DonationManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonationManage.Contract.DonationManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DonationManage *DonationManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonationManage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DonationManage *DonationManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonationManage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DonationManage *DonationManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonationManage.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_DonationManage *DonationManageCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_DonationManage *DonationManageSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _DonationManage.Contract.UPGRADEINTERFACEVERSION(&_DonationManage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_DonationManage *DonationManageCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _DonationManage.Contract.UPGRADEINTERFACEVERSION(&_DonationManage.CallOpts)
}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_DonationManage *DonationManageCaller) CampaignCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "campaignCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_DonationManage *DonationManageSession) CampaignCount() (*big.Int, error) {
	return _DonationManage.Contract.CampaignCount(&_DonationManage.CallOpts)
}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_DonationManage *DonationManageCallerSession) CampaignCount() (*big.Int, error) {
	return _DonationManage.Contract.CampaignCount(&_DonationManage.CallOpts)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status, address starter)
func (_DonationManage *DonationManageCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Title          string
	Goal           *big.Int
	TotalDonated   *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	Nature         uint8
	Beneficiary    uint8
	Purpose        string
	ExpectedImpact string
	Status         uint8
	Starter        common.Address
}, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "campaigns", arg0)

	outstruct := new(struct {
		Title          string
		Goal           *big.Int
		TotalDonated   *big.Int
		StartTime      *big.Int
		EndTime        *big.Int
		Nature         uint8
		Beneficiary    uint8
		Purpose        string
		ExpectedImpact string
		Status         uint8
		Starter        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Goal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalDonated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Nature = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Beneficiary = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Purpose = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.ExpectedImpact = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.Starter = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status, address starter)
func (_DonationManage *DonationManageSession) Campaigns(arg0 *big.Int) (struct {
	Title          string
	Goal           *big.Int
	TotalDonated   *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	Nature         uint8
	Beneficiary    uint8
	Purpose        string
	ExpectedImpact string
	Status         uint8
	Starter        common.Address
}, error) {
	return _DonationManage.Contract.Campaigns(&_DonationManage.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status, address starter)
func (_DonationManage *DonationManageCallerSession) Campaigns(arg0 *big.Int) (struct {
	Title          string
	Goal           *big.Int
	TotalDonated   *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	Nature         uint8
	Beneficiary    uint8
	Purpose        string
	ExpectedImpact string
	Status         uint8
	Starter        common.Address
}, error) {
	return _DonationManage.Contract.Campaigns(&_DonationManage.CallOpts, arg0)
}

// CampaignsToCheck is a free data retrieval call binding the contract method 0xbe6c81a5.
//
// Solidity: function campaignsToCheck(uint256 ) view returns(uint256)
func (_DonationManage *DonationManageCaller) CampaignsToCheck(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "campaignsToCheck", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignsToCheck is a free data retrieval call binding the contract method 0xbe6c81a5.
//
// Solidity: function campaignsToCheck(uint256 ) view returns(uint256)
func (_DonationManage *DonationManageSession) CampaignsToCheck(arg0 *big.Int) (*big.Int, error) {
	return _DonationManage.Contract.CampaignsToCheck(&_DonationManage.CallOpts, arg0)
}

// CampaignsToCheck is a free data retrieval call binding the contract method 0xbe6c81a5.
//
// Solidity: function campaignsToCheck(uint256 ) view returns(uint256)
func (_DonationManage *DonationManageCallerSession) CampaignsToCheck(arg0 *big.Int) (*big.Int, error) {
	return _DonationManage.Contract.CampaignsToCheck(&_DonationManage.CallOpts, arg0)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_DonationManage *DonationManageCaller) CheckUpkeep(opts *bind.CallOpts, arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "checkUpkeep", arg0)

	outstruct := new(struct {
		UpkeepNeeded bool
		PerformData  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpkeepNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PerformData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_DonationManage *DonationManageSession) CheckUpkeep(arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _DonationManage.Contract.CheckUpkeep(&_DonationManage.CallOpts, arg0)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_DonationManage *DonationManageCallerSession) CheckUpkeep(arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _DonationManage.Contract.CheckUpkeep(&_DonationManage.CallOpts, arg0)
}

// GetCampaignDetails is a free data retrieval call binding the contract method 0xd993de62.
//
// Solidity: function getCampaignDetails(uint256 _campaignId) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 status, address starter)
func (_DonationManage *DonationManageCaller) GetCampaignDetails(opts *bind.CallOpts, _campaignId *big.Int) (struct {
	Title        string
	Goal         *big.Int
	TotalDonated *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Status       uint8
	Starter      common.Address
}, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "getCampaignDetails", _campaignId)

	outstruct := new(struct {
		Title        string
		Goal         *big.Int
		TotalDonated *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Status       uint8
		Starter      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Goal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalDonated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Starter = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetCampaignDetails is a free data retrieval call binding the contract method 0xd993de62.
//
// Solidity: function getCampaignDetails(uint256 _campaignId) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 status, address starter)
func (_DonationManage *DonationManageSession) GetCampaignDetails(_campaignId *big.Int) (struct {
	Title        string
	Goal         *big.Int
	TotalDonated *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Status       uint8
	Starter      common.Address
}, error) {
	return _DonationManage.Contract.GetCampaignDetails(&_DonationManage.CallOpts, _campaignId)
}

// GetCampaignDetails is a free data retrieval call binding the contract method 0xd993de62.
//
// Solidity: function getCampaignDetails(uint256 _campaignId) view returns(string title, uint256 goal, uint256 totalDonated, uint256 startTime, uint256 endTime, uint8 status, address starter)
func (_DonationManage *DonationManageCallerSession) GetCampaignDetails(_campaignId *big.Int) (struct {
	Title        string
	Goal         *big.Int
	TotalDonated *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Status       uint8
	Starter      common.Address
}, error) {
	return _DonationManage.Contract.GetCampaignDetails(&_DonationManage.CallOpts, _campaignId)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_DonationManage *DonationManageCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_DonationManage *DonationManageSession) GetVersion() (string, error) {
	return _DonationManage.Contract.GetVersion(&_DonationManage.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_DonationManage *DonationManageCallerSession) GetVersion() (string, error) {
	return _DonationManage.Contract.GetVersion(&_DonationManage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DonationManage *DonationManageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DonationManage *DonationManageSession) Owner() (common.Address, error) {
	return _DonationManage.Contract.Owner(&_DonationManage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DonationManage *DonationManageCallerSession) Owner() (common.Address, error) {
	return _DonationManage.Contract.Owner(&_DonationManage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DonationManage *DonationManageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DonationManage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DonationManage *DonationManageSession) ProxiableUUID() ([32]byte, error) {
	return _DonationManage.Contract.ProxiableUUID(&_DonationManage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_DonationManage *DonationManageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _DonationManage.Contract.ProxiableUUID(&_DonationManage.CallOpts)
}

// ActiveCampaign is a paid mutator transaction binding the contract method 0x23be7dac.
//
// Solidity: function activeCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactor) ActiveCampaign(opts *bind.TransactOpts, _campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "activeCampaign", _campaignId)
}

// ActiveCampaign is a paid mutator transaction binding the contract method 0x23be7dac.
//
// Solidity: function activeCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageSession) ActiveCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.ActiveCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// ActiveCampaign is a paid mutator transaction binding the contract method 0x23be7dac.
//
// Solidity: function activeCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactorSession) ActiveCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.ActiveCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// CancellCampaign is a paid mutator transaction binding the contract method 0xb4e5be29.
//
// Solidity: function cancellCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactor) CancellCampaign(opts *bind.TransactOpts, _campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "cancellCampaign", _campaignId)
}

// CancellCampaign is a paid mutator transaction binding the contract method 0xb4e5be29.
//
// Solidity: function cancellCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageSession) CancellCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.CancellCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// CancellCampaign is a paid mutator transaction binding the contract method 0xb4e5be29.
//
// Solidity: function cancellCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactorSession) CancellCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.CancellCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// CompletedCampaign is a paid mutator transaction binding the contract method 0x5177af14.
//
// Solidity: function completedCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactor) CompletedCampaign(opts *bind.TransactOpts, _campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "completedCampaign", _campaignId)
}

// CompletedCampaign is a paid mutator transaction binding the contract method 0x5177af14.
//
// Solidity: function completedCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageSession) CompletedCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.CompletedCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// CompletedCampaign is a paid mutator transaction binding the contract method 0x5177af14.
//
// Solidity: function completedCampaign(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactorSession) CompletedCampaign(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.CompletedCampaign(&_DonationManage.TransactOpts, _campaignId)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9d1a3691.
//
// Solidity: function createCampaign(string _title, uint256 _goal, uint256 _startTime, uint256 _endTime, uint8 _nature, uint8 _beneficiary, string _purpose, string _expectedImpact) returns()
func (_DonationManage *DonationManageTransactor) CreateCampaign(opts *bind.TransactOpts, _title string, _goal *big.Int, _startTime *big.Int, _endTime *big.Int, _nature uint8, _beneficiary uint8, _purpose string, _expectedImpact string) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "createCampaign", _title, _goal, _startTime, _endTime, _nature, _beneficiary, _purpose, _expectedImpact)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9d1a3691.
//
// Solidity: function createCampaign(string _title, uint256 _goal, uint256 _startTime, uint256 _endTime, uint8 _nature, uint8 _beneficiary, string _purpose, string _expectedImpact) returns()
func (_DonationManage *DonationManageSession) CreateCampaign(_title string, _goal *big.Int, _startTime *big.Int, _endTime *big.Int, _nature uint8, _beneficiary uint8, _purpose string, _expectedImpact string) (*types.Transaction, error) {
	return _DonationManage.Contract.CreateCampaign(&_DonationManage.TransactOpts, _title, _goal, _startTime, _endTime, _nature, _beneficiary, _purpose, _expectedImpact)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x9d1a3691.
//
// Solidity: function createCampaign(string _title, uint256 _goal, uint256 _startTime, uint256 _endTime, uint8 _nature, uint8 _beneficiary, string _purpose, string _expectedImpact) returns()
func (_DonationManage *DonationManageTransactorSession) CreateCampaign(_title string, _goal *big.Int, _startTime *big.Int, _endTime *big.Int, _nature uint8, _beneficiary uint8, _purpose string, _expectedImpact string) (*types.Transaction, error) {
	return _DonationManage.Contract.CreateCampaign(&_DonationManage.TransactOpts, _title, _goal, _startTime, _endTime, _nature, _beneficiary, _purpose, _expectedImpact)
}

// Donate is a paid mutator transaction binding the contract method 0x0cdd53f6.
//
// Solidity: function donate(uint256 _campaignId, uint256 _amount) returns()
func (_DonationManage *DonationManageTransactor) Donate(opts *bind.TransactOpts, _campaignId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "donate", _campaignId, _amount)
}

// Donate is a paid mutator transaction binding the contract method 0x0cdd53f6.
//
// Solidity: function donate(uint256 _campaignId, uint256 _amount) returns()
func (_DonationManage *DonationManageSession) Donate(_campaignId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.Donate(&_DonationManage.TransactOpts, _campaignId, _amount)
}

// Donate is a paid mutator transaction binding the contract method 0x0cdd53f6.
//
// Solidity: function donate(uint256 _campaignId, uint256 _amount) returns()
func (_DonationManage *DonationManageTransactorSession) Donate(_campaignId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.Donate(&_DonationManage.TransactOpts, _campaignId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _nft) returns()
func (_DonationManage *DonationManageTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _nft common.Address) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "initialize", _token, _nft)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _nft) returns()
func (_DonationManage *DonationManageSession) Initialize(_token common.Address, _nft common.Address) (*types.Transaction, error) {
	return _DonationManage.Contract.Initialize(&_DonationManage.TransactOpts, _token, _nft)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _nft) returns()
func (_DonationManage *DonationManageTransactorSession) Initialize(_token common.Address, _nft common.Address) (*types.Transaction, error) {
	return _DonationManage.Contract.Initialize(&_DonationManage.TransactOpts, _token, _nft)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_DonationManage *DonationManageTransactor) PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "performUpkeep", performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_DonationManage *DonationManageSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _DonationManage.Contract.PerformUpkeep(&_DonationManage.TransactOpts, performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_DonationManage *DonationManageTransactorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _DonationManage.Contract.PerformUpkeep(&_DonationManage.TransactOpts, performData)
}

// RefundToken is a paid mutator transaction binding the contract method 0x17191704.
//
// Solidity: function refundToken(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactor) RefundToken(opts *bind.TransactOpts, _campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "refundToken", _campaignId)
}

// RefundToken is a paid mutator transaction binding the contract method 0x17191704.
//
// Solidity: function refundToken(uint256 _campaignId) returns()
func (_DonationManage *DonationManageSession) RefundToken(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.RefundToken(&_DonationManage.TransactOpts, _campaignId)
}

// RefundToken is a paid mutator transaction binding the contract method 0x17191704.
//
// Solidity: function refundToken(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactorSession) RefundToken(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.RefundToken(&_DonationManage.TransactOpts, _campaignId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DonationManage *DonationManageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DonationManage *DonationManageSession) RenounceOwnership() (*types.Transaction, error) {
	return _DonationManage.Contract.RenounceOwnership(&_DonationManage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DonationManage *DonationManageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DonationManage.Contract.RenounceOwnership(&_DonationManage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DonationManage *DonationManageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DonationManage *DonationManageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DonationManage.Contract.TransferOwnership(&_DonationManage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DonationManage *DonationManageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DonationManage.Contract.TransferOwnership(&_DonationManage.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DonationManage *DonationManageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DonationManage *DonationManageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DonationManage.Contract.UpgradeToAndCall(&_DonationManage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DonationManage *DonationManageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DonationManage.Contract.UpgradeToAndCall(&_DonationManage.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactor) Withdraw(opts *bind.TransactOpts, _campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.contract.Transact(opts, "withdraw", _campaignId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _campaignId) returns()
func (_DonationManage *DonationManageSession) Withdraw(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.Withdraw(&_DonationManage.TransactOpts, _campaignId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _campaignId) returns()
func (_DonationManage *DonationManageTransactorSession) Withdraw(_campaignId *big.Int) (*types.Transaction, error) {
	return _DonationManage.Contract.Withdraw(&_DonationManage.TransactOpts, _campaignId)
}

// DonationManageActiveCampaignIterator is returned from FilterActiveCampaign and is used to iterate over the raw logs and unpacked data for ActiveCampaign events raised by the DonationManage contract.
type DonationManageActiveCampaignIterator struct {
	Event *DonationManageActiveCampaign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageActiveCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageActiveCampaign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageActiveCampaign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageActiveCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageActiveCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageActiveCampaign represents a ActiveCampaign event raised by the DonationManage contract.
type DonationManageActiveCampaign struct {
	Id     *big.Int
	Caller common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActiveCampaign is a free log retrieval operation binding the contract event 0x45e7ca8c14330ef52434b43a323a650db341a2b85a5160f66d892e3a9623c227.
//
// Solidity: event ActiveCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) FilterActiveCampaign(opts *bind.FilterOpts, id []*big.Int) (*DonationManageActiveCampaignIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "ActiveCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageActiveCampaignIterator{contract: _DonationManage.contract, event: "ActiveCampaign", logs: logs, sub: sub}, nil
}

// WatchActiveCampaign is a free log subscription operation binding the contract event 0x45e7ca8c14330ef52434b43a323a650db341a2b85a5160f66d892e3a9623c227.
//
// Solidity: event ActiveCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) WatchActiveCampaign(opts *bind.WatchOpts, sink chan<- *DonationManageActiveCampaign, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "ActiveCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageActiveCampaign)
				if err := _DonationManage.contract.UnpackLog(event, "ActiveCampaign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActiveCampaign is a log parse operation binding the contract event 0x45e7ca8c14330ef52434b43a323a650db341a2b85a5160f66d892e3a9623c227.
//
// Solidity: event ActiveCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) ParseActiveCampaign(log types.Log) (*DonationManageActiveCampaign, error) {
	event := new(DonationManageActiveCampaign)
	if err := _DonationManage.contract.UnpackLog(event, "ActiveCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageCancellCampaignIterator is returned from FilterCancellCampaign and is used to iterate over the raw logs and unpacked data for CancellCampaign events raised by the DonationManage contract.
type DonationManageCancellCampaignIterator struct {
	Event *DonationManageCancellCampaign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageCancellCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageCancellCampaign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageCancellCampaign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageCancellCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageCancellCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageCancellCampaign represents a CancellCampaign event raised by the DonationManage contract.
type DonationManageCancellCampaign struct {
	Id     *big.Int
	Caller common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancellCampaign is a free log retrieval operation binding the contract event 0x1c37c67694e49ddaa169e8b7f9dafd655ba8cfc26df25bf66ea18d649d5e7096.
//
// Solidity: event CancellCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) FilterCancellCampaign(opts *bind.FilterOpts, id []*big.Int) (*DonationManageCancellCampaignIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "CancellCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageCancellCampaignIterator{contract: _DonationManage.contract, event: "CancellCampaign", logs: logs, sub: sub}, nil
}

// WatchCancellCampaign is a free log subscription operation binding the contract event 0x1c37c67694e49ddaa169e8b7f9dafd655ba8cfc26df25bf66ea18d649d5e7096.
//
// Solidity: event CancellCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) WatchCancellCampaign(opts *bind.WatchOpts, sink chan<- *DonationManageCancellCampaign, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "CancellCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageCancellCampaign)
				if err := _DonationManage.contract.UnpackLog(event, "CancellCampaign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancellCampaign is a log parse operation binding the contract event 0x1c37c67694e49ddaa169e8b7f9dafd655ba8cfc26df25bf66ea18d649d5e7096.
//
// Solidity: event CancellCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) ParseCancellCampaign(log types.Log) (*DonationManageCancellCampaign, error) {
	event := new(DonationManageCancellCampaign)
	if err := _DonationManage.contract.UnpackLog(event, "CancellCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageCompletedCampaignIterator is returned from FilterCompletedCampaign and is used to iterate over the raw logs and unpacked data for CompletedCampaign events raised by the DonationManage contract.
type DonationManageCompletedCampaignIterator struct {
	Event *DonationManageCompletedCampaign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageCompletedCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageCompletedCampaign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageCompletedCampaign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageCompletedCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageCompletedCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageCompletedCampaign represents a CompletedCampaign event raised by the DonationManage contract.
type DonationManageCompletedCampaign struct {
	Id     *big.Int
	Caller common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCompletedCampaign is a free log retrieval operation binding the contract event 0xcab935963eaa03f3973204c91b5bc658726a1015d91fc361d82f67c407172300.
//
// Solidity: event CompletedCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) FilterCompletedCampaign(opts *bind.FilterOpts, id []*big.Int) (*DonationManageCompletedCampaignIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "CompletedCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageCompletedCampaignIterator{contract: _DonationManage.contract, event: "CompletedCampaign", logs: logs, sub: sub}, nil
}

// WatchCompletedCampaign is a free log subscription operation binding the contract event 0xcab935963eaa03f3973204c91b5bc658726a1015d91fc361d82f67c407172300.
//
// Solidity: event CompletedCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) WatchCompletedCampaign(opts *bind.WatchOpts, sink chan<- *DonationManageCompletedCampaign, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "CompletedCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageCompletedCampaign)
				if err := _DonationManage.contract.UnpackLog(event, "CompletedCampaign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompletedCampaign is a log parse operation binding the contract event 0xcab935963eaa03f3973204c91b5bc658726a1015d91fc361d82f67c407172300.
//
// Solidity: event CompletedCampaign(uint256 indexed id, address caller, uint256 time)
func (_DonationManage *DonationManageFilterer) ParseCompletedCampaign(log types.Log) (*DonationManageCompletedCampaign, error) {
	event := new(DonationManageCompletedCampaign)
	if err := _DonationManage.contract.UnpackLog(event, "CompletedCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the DonationManage contract.
type DonationManageDonateIterator struct {
	Event *DonationManageDonate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageDonate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageDonate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageDonate represents a Donate event raised by the DonationManage contract.
type DonationManageDonate struct {
	Id      *big.Int
	Donater common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x48def02da8c59b50861e25bab759a8a0c0de6fab6448ef85f32ce59397947d80.
//
// Solidity: event Donate(uint256 indexed id, address donater, uint256 amount)
func (_DonationManage *DonationManageFilterer) FilterDonate(opts *bind.FilterOpts, id []*big.Int) (*DonationManageDonateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "Donate", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageDonateIterator{contract: _DonationManage.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x48def02da8c59b50861e25bab759a8a0c0de6fab6448ef85f32ce59397947d80.
//
// Solidity: event Donate(uint256 indexed id, address donater, uint256 amount)
func (_DonationManage *DonationManageFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *DonationManageDonate, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "Donate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageDonate)
				if err := _DonationManage.contract.UnpackLog(event, "Donate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDonate is a log parse operation binding the contract event 0x48def02da8c59b50861e25bab759a8a0c0de6fab6448ef85f32ce59397947d80.
//
// Solidity: event Donate(uint256 indexed id, address donater, uint256 amount)
func (_DonationManage *DonationManageFilterer) ParseDonate(log types.Log) (*DonationManageDonate, error) {
	event := new(DonationManageDonate)
	if err := _DonationManage.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DonationManage contract.
type DonationManageInitializedIterator struct {
	Event *DonationManageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageInitialized represents a Initialized event raised by the DonationManage contract.
type DonationManageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DonationManage *DonationManageFilterer) FilterInitialized(opts *bind.FilterOpts) (*DonationManageInitializedIterator, error) {

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DonationManageInitializedIterator{contract: _DonationManage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DonationManage *DonationManageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DonationManageInitialized) (event.Subscription, error) {

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageInitialized)
				if err := _DonationManage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DonationManage *DonationManageFilterer) ParseInitialized(log types.Log) (*DonationManageInitialized, error) {
	event := new(DonationManageInitialized)
	if err := _DonationManage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DonationManage contract.
type DonationManageOwnershipTransferredIterator struct {
	Event *DonationManageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageOwnershipTransferred represents a OwnershipTransferred event raised by the DonationManage contract.
type DonationManageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DonationManage *DonationManageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DonationManageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageOwnershipTransferredIterator{contract: _DonationManage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DonationManage *DonationManageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DonationManageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageOwnershipTransferred)
				if err := _DonationManage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DonationManage *DonationManageFilterer) ParseOwnershipTransferred(log types.Log) (*DonationManageOwnershipTransferred, error) {
	event := new(DonationManageOwnershipTransferred)
	if err := _DonationManage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the DonationManage contract.
type DonationManageRefundIterator struct {
	Event *DonationManageRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageRefund represents a Refund event raised by the DonationManage contract.
type DonationManageRefund struct {
	Id           *big.Int
	Refunder     common.Address
	RefundAmount *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed id, address refunder, uint256 refundAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) FilterRefund(opts *bind.FilterOpts, id []*big.Int) (*DonationManageRefundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "Refund", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageRefundIterator{contract: _DonationManage.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed id, address refunder, uint256 refundAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *DonationManageRefund, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "Refund", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageRefund)
				if err := _DonationManage.contract.UnpackLog(event, "Refund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefund is a log parse operation binding the contract event 0xaba4210cad11f1112f31ba136a61cdc0f6ab4c2b478da8c93eec1d9769e97a89.
//
// Solidity: event Refund(uint256 indexed id, address refunder, uint256 refundAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) ParseRefund(log types.Log) (*DonationManageRefund, error) {
	event := new(DonationManageRefund)
	if err := _DonationManage.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageStartCampaignIterator is returned from FilterStartCampaign and is used to iterate over the raw logs and unpacked data for StartCampaign events raised by the DonationManage contract.
type DonationManageStartCampaignIterator struct {
	Event *DonationManageStartCampaign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageStartCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageStartCampaign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageStartCampaign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageStartCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageStartCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageStartCampaign represents a StartCampaign event raised by the DonationManage contract.
type DonationManageStartCampaign struct {
	Id             *big.Int
	Starter        common.Address
	Title          string
	Goal           *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	Nature         uint8
	Beneficiary    uint8
	Purpose        string
	ExpectedImpact string
	Status         uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStartCampaign is a free log retrieval operation binding the contract event 0xbe3dd7dc6384f58902a35b3dbc09ad6f5cc91b4499b64a64574f25c4c13cdef6.
//
// Solidity: event StartCampaign(uint256 indexed id, address starter, string title, uint256 goal, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status)
func (_DonationManage *DonationManageFilterer) FilterStartCampaign(opts *bind.FilterOpts, id []*big.Int) (*DonationManageStartCampaignIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "StartCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageStartCampaignIterator{contract: _DonationManage.contract, event: "StartCampaign", logs: logs, sub: sub}, nil
}

// WatchStartCampaign is a free log subscription operation binding the contract event 0xbe3dd7dc6384f58902a35b3dbc09ad6f5cc91b4499b64a64574f25c4c13cdef6.
//
// Solidity: event StartCampaign(uint256 indexed id, address starter, string title, uint256 goal, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status)
func (_DonationManage *DonationManageFilterer) WatchStartCampaign(opts *bind.WatchOpts, sink chan<- *DonationManageStartCampaign, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "StartCampaign", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageStartCampaign)
				if err := _DonationManage.contract.UnpackLog(event, "StartCampaign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartCampaign is a log parse operation binding the contract event 0xbe3dd7dc6384f58902a35b3dbc09ad6f5cc91b4499b64a64574f25c4c13cdef6.
//
// Solidity: event StartCampaign(uint256 indexed id, address starter, string title, uint256 goal, uint256 startTime, uint256 endTime, uint8 nature, uint8 beneficiary, string purpose, string expectedImpact, uint8 status)
func (_DonationManage *DonationManageFilterer) ParseStartCampaign(log types.Log) (*DonationManageStartCampaign, error) {
	event := new(DonationManageStartCampaign)
	if err := _DonationManage.contract.UnpackLog(event, "StartCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the DonationManage contract.
type DonationManageUpgradedIterator struct {
	Event *DonationManageUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageUpgraded represents a Upgraded event raised by the DonationManage contract.
type DonationManageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DonationManage *DonationManageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DonationManageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageUpgradedIterator{contract: _DonationManage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DonationManage *DonationManageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DonationManageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageUpgraded)
				if err := _DonationManage.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DonationManage *DonationManageFilterer) ParseUpgraded(log types.Log) (*DonationManageUpgraded, error) {
	event := new(DonationManageUpgraded)
	if err := _DonationManage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DonationManageWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the DonationManage contract.
type DonationManageWithdrawIterator struct {
	Event *DonationManageWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DonationManageWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonationManageWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DonationManageWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DonationManageWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DonationManageWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DonationManageWithdraw represents a Withdraw event raised by the DonationManage contract.
type DonationManageWithdraw struct {
	Id             *big.Int
	Withdrawer     common.Address
	WithdrawAmount *big.Int
	Time           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 indexed id, address withdrawer, uint256 withdrawAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) FilterWithdraw(opts *bind.FilterOpts, id []*big.Int) (*DonationManageWithdrawIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.FilterLogs(opts, "Withdraw", idRule)
	if err != nil {
		return nil, err
	}
	return &DonationManageWithdrawIterator{contract: _DonationManage.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 indexed id, address withdrawer, uint256 withdrawAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DonationManageWithdraw, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DonationManage.contract.WatchLogs(opts, "Withdraw", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DonationManageWithdraw)
				if err := _DonationManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 indexed id, address withdrawer, uint256 withdrawAmount, uint256 time)
func (_DonationManage *DonationManageFilterer) ParseWithdraw(log types.Log) (*DonationManageWithdraw, error) {
	event := new(DonationManageWithdraw)
	if err := _DonationManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
