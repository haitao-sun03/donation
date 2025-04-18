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

// IL1MessageServiceClaimMessageWithProofParams is an auto generated low-level Go binding around an user-defined struct.
type IL1MessageServiceClaimMessageWithProofParams struct {
	Proof         [][32]byte
	MessageNumber *big.Int
	LeafIndex     uint32
	From          common.Address
	To            common.Address
	Fee           *big.Int
	Value         *big.Int
	FeeRecipient  common.Address
	MerkleRoot    [32]byte
	Data          []byte
}

// ILineaRollupBlobSubmission is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupBlobSubmission struct {
	DataEvaluationClaim *big.Int
	KzgCommitment       []byte
	KzgProof            []byte
	FinalStateRootHash  [32]byte
	SnarkHash           [32]byte
}

// ILineaRollupCompressedCalldataSubmission is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupCompressedCalldataSubmission struct {
	FinalStateRootHash [32]byte
	SnarkHash          [32]byte
	CompressedData     []byte
}

// ILineaRollupFinalizationDataV3 is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupFinalizationDataV3 struct {
	ParentStateRootHash                     [32]byte
	EndBlockNumber                          *big.Int
	ShnarfData                              ILineaRollupShnarfData
	LastFinalizedTimestamp                  *big.Int
	FinalTimestamp                          *big.Int
	LastFinalizedL1RollingHash              [32]byte
	L1RollingHash                           [32]byte
	LastFinalizedL1RollingHashMessageNumber *big.Int
	L1RollingHashMessageNumber              *big.Int
	L2MerkleTreesDepth                      *big.Int
	L2MerkleRoots                           [][32]byte
	L2MessagingBlocksOffsets                []byte
}

// ILineaRollupInitializationData is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupInitializationData struct {
	InitialStateRootHash     [32]byte
	InitialL2BlockNumber     *big.Int
	GenesisTimestamp         *big.Int
	DefaultVerifier          common.Address
	RateLimitPeriodInSeconds *big.Int
	RateLimitAmountInWei     *big.Int
	RoleAddresses            []IPermissionsManagerRoleAddress
	PauseTypeRoles           []IPauseManagerPauseTypeRole
	UnpauseTypeRoles         []IPauseManagerPauseTypeRole
	FallbackOperator         common.Address
	DefaultAdmin             common.Address
}

// ILineaRollupShnarfData is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupShnarfData struct {
	ParentShnarf        [32]byte
	SnarkHash           [32]byte
	FinalStateRootHash  [32]byte
	DataEvaluationPoint [32]byte
	DataEvaluationClaim [32]byte
}

// IPauseManagerPauseTypeRole is an auto generated low-level Go binding around an user-defined struct.
type IPauseManagerPauseTypeRole struct {
	PauseType uint8
	Role      [32]byte
}

// IPermissionsManagerRoleAddress is an auto generated low-level Go binding around an user-defined struct.
type IPermissionsManagerRoleAddress struct {
	AddressWithRole common.Address
	Role            [32]byte
}

// LineaSepoliaVerifyMetaData contains all meta data concerning the LineaSepoliaVerify contract.
var LineaSepoliaVerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ArrayLengthsDoNotMatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"emptyBlobIndex\",\"type\":\"uint256\"}],\"name\":\"BlobSubmissionDataEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlobSubmissionDataIsMissing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesLengthNotMultipleOf32\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bytesLength\",\"type\":\"uint256\"}],\"name\":\"BytesLengthNotMultipleOfTwo\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currentDataHash\",\"type\":\"bytes32\"}],\"name\":\"DataAlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"EmptyBlobDataAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySubmissionData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeePaymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"}],\"name\":\"FinalBlobNotSubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint256\"}],\"name\":\"FinalBlockNumberLessThanOrEqualToLastFinalizedBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalBlockStateEqualsZeroHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"FinalShnarfWrong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"FinalizationInTheFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"FinalizationStateIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstByteIsNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errorReason\",\"type\":\"string\"}],\"name\":\"InvalidProofOrProofVerificationRanOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"}],\"name\":\"IsNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"}],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"name\":\"L1RollingHashDoesNotExistOnL1\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"L2MerkleRootAlreadyAnchored\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2MerkleRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastFinalizationTimeNotLapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxAllowedIndex\",\"type\":\"uint32\"}],\"name\":\"LeafIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"MessageAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageDoesNotExistOrHasAlreadyBeenClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"MessageSendingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"name\":\"MissingMessageNumberForRollingHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"}],\"name\":\"MissingRollingHashForMessageNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNonFallbackOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"}],\"name\":\"ParentBlobNotSubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PointEvaluationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fieldElements\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blsCurveModulus\",\"type\":\"uint256\"}],\"name\":\"PointEvaluationResponseInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"PrecompileReturnDataLengthWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"ProofLengthDifferentThanMerkleDepth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartingRootHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueSentTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroHashNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resettingAddress\",\"type\":\"address\"}],\"name\":\"AmountUsedInPeriodReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentStateRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"}],\"name\":\"DataFinalizedV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentShnarf\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"}],\"name\":\"DataSubmittedV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOperator\",\"type\":\"address\"}],\"name\":\"FallbackOperatorAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOperator\",\"type\":\"address\"}],\"name\":\"FallbackOperatorRoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2MerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"treeDepth\",\"type\":\"uint256\"}],\"name\":\"L2MerkleRootAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"}],\"name\":\"L2MessagingBlockAnchored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"amountChangeBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"amountUsedLoweredToLimit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"usedAmountResetToZero\",\"type\":\"bool\"}],\"name\":\"LimitAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"previousVersion\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"newVersion\",\"type\":\"bytes8\"}],\"name\":\"LineaRollupVersionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"PauseTypeRoleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodInSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPeriodEnd\",\"type\":\"uint256\"}],\"name\":\"RateLimitInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RollingHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"unPauseType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"UnPauseTypeRoleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierSetBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRACT_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_SHNARF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_SENT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ALL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_BLOB_SUBMISSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_FINALIZATION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_L1_L2_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_L2_L1_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_ALL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_BLOB_SUBMISSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_FINALIZATION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_L1_L2_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_L2_L1_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USED_RATE_LIMIT_RESETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_UNSETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"}],\"name\":\"blobShnarfExists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"exists\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIL1MessageService.ClaimMessageWithProofParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"claimMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFinalizedShnarf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFinalizedState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2StoredL1MessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2StoredL1RollingHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodAmountInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataEndingBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endingBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataFinalStateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataParents\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataShnarfHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarfHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataStartingBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startingBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_aggregatedProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationPoint\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationClaim\",\"type\":\"bytes32\"}],\"internalType\":\"structILineaRollup.ShnarfData\",\"name\":\"shnarfData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lastFinalizedL1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedL1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2MerkleTreesDepth\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2MerkleRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"l2MessagingBlocksOffsets\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.FinalizationDataV3\",\"name\":\"_finalizationData\",\"type\":\"tuple\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"inboxL2L1MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageStatus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"initialStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"initialL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"defaultVerifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rateLimitPeriodInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateLimitAmountInWei\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addressWithRole\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPermissionsManager.RoleAddress[]\",\"name\":\"roleAddresses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPauseManager.PauseTypeRole[]\",\"name\":\"pauseTypeRoles\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPauseManager.PauseTypeRole[]\",\"name\":\"unpauseTypeRoles\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"fallbackOperator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"internalType\":\"structILineaRollup.InitializationData\",\"name\":\"_initializationData\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messageNumber\",\"type\":\"uint256\"}],\"name\":\"isMessageClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pauseTypeIsPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"l2MerkleRootsDepths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"treeDepth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"outboxL1L2MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageStatus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"pauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"pauseTypeStatuses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pauseStatus\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addressWithRole\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPermissionsManager.RoleAddress[]\",\"name\":\"_roleAddresses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPauseManager.PauseTypeRole[]\",\"name\":\"_pauseTypeRoles\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"pauseType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"internalType\":\"structIPauseManager.PauseTypeRole[]\",\"name\":\"_unpauseTypeRoles\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_fallbackOperator\",\"type\":\"address\"}],\"name\":\"reinitializeLineaRollupV6\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetAmountUsedInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resetRateLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"}],\"name\":\"rollingHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messageNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_rollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lastFinalizedTimestamp\",\"type\":\"uint256\"}],\"name\":\"setFallbackOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"stateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dataEvaluationClaim\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"kzgCommitment\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"kzgProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"}],\"internalType\":\"structILineaRollup.BlobSubmission[]\",\"name\":\"_blobSubmissions\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_finalBlobShnarf\",\"type\":\"bytes32\"}],\"name\":\"submitBlobs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedData\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.CompressedCalldataSubmission\",\"name\":\"_submission\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_expectedShnarf\",\"type\":\"bytes32\"}],\"name\":\"submitDataAsCalldata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMigrationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPauseManager.PauseType\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"unPauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"unsetVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LineaSepoliaVerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use LineaSepoliaVerifyMetaData.ABI instead.
var LineaSepoliaVerifyABI = LineaSepoliaVerifyMetaData.ABI

// LineaSepoliaVerify is an auto generated Go binding around an Ethereum contract.
type LineaSepoliaVerify struct {
	LineaSepoliaVerifyCaller     // Read-only binding to the contract
	LineaSepoliaVerifyTransactor // Write-only binding to the contract
	LineaSepoliaVerifyFilterer   // Log filterer for contract events
}

// LineaSepoliaVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type LineaSepoliaVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaSepoliaVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LineaSepoliaVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaSepoliaVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LineaSepoliaVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LineaSepoliaVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LineaSepoliaVerifySession struct {
	Contract     *LineaSepoliaVerify // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LineaSepoliaVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LineaSepoliaVerifyCallerSession struct {
	Contract *LineaSepoliaVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LineaSepoliaVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LineaSepoliaVerifyTransactorSession struct {
	Contract     *LineaSepoliaVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LineaSepoliaVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type LineaSepoliaVerifyRaw struct {
	Contract *LineaSepoliaVerify // Generic contract binding to access the raw methods on
}

// LineaSepoliaVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LineaSepoliaVerifyCallerRaw struct {
	Contract *LineaSepoliaVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// LineaSepoliaVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LineaSepoliaVerifyTransactorRaw struct {
	Contract *LineaSepoliaVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLineaSepoliaVerify creates a new instance of LineaSepoliaVerify, bound to a specific deployed contract.
func NewLineaSepoliaVerify(address common.Address, backend bind.ContractBackend) (*LineaSepoliaVerify, error) {
	contract, err := bindLineaSepoliaVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerify{LineaSepoliaVerifyCaller: LineaSepoliaVerifyCaller{contract: contract}, LineaSepoliaVerifyTransactor: LineaSepoliaVerifyTransactor{contract: contract}, LineaSepoliaVerifyFilterer: LineaSepoliaVerifyFilterer{contract: contract}}, nil
}

// NewLineaSepoliaVerifyCaller creates a new read-only instance of LineaSepoliaVerify, bound to a specific deployed contract.
func NewLineaSepoliaVerifyCaller(address common.Address, caller bind.ContractCaller) (*LineaSepoliaVerifyCaller, error) {
	contract, err := bindLineaSepoliaVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyCaller{contract: contract}, nil
}

// NewLineaSepoliaVerifyTransactor creates a new write-only instance of LineaSepoliaVerify, bound to a specific deployed contract.
func NewLineaSepoliaVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*LineaSepoliaVerifyTransactor, error) {
	contract, err := bindLineaSepoliaVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyTransactor{contract: contract}, nil
}

// NewLineaSepoliaVerifyFilterer creates a new log filterer instance of LineaSepoliaVerify, bound to a specific deployed contract.
func NewLineaSepoliaVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*LineaSepoliaVerifyFilterer, error) {
	contract, err := bindLineaSepoliaVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyFilterer{contract: contract}, nil
}

// bindLineaSepoliaVerify binds a generic wrapper to an already deployed contract.
func bindLineaSepoliaVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LineaSepoliaVerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LineaSepoliaVerify *LineaSepoliaVerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LineaSepoliaVerify.Contract.LineaSepoliaVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LineaSepoliaVerify *LineaSepoliaVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.LineaSepoliaVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LineaSepoliaVerify *LineaSepoliaVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.LineaSepoliaVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LineaSepoliaVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTVERSION is a free data retrieval call binding the contract method 0x38b90333.
//
// Solidity: function CONTRACT_VERSION() view returns(string)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CONTRACTVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "CONTRACT_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CONTRACTVERSION is a free data retrieval call binding the contract method 0x38b90333.
//
// Solidity: function CONTRACT_VERSION() view returns(string)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CONTRACTVERSION() (string, error) {
	return _LineaSepoliaVerify.Contract.CONTRACTVERSION(&_LineaSepoliaVerify.CallOpts)
}

// CONTRACTVERSION is a free data retrieval call binding the contract method 0x38b90333.
//
// Solidity: function CONTRACT_VERSION() view returns(string)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CONTRACTVERSION() (string, error) {
	return _LineaSepoliaVerify.Contract.CONTRACTVERSION(&_LineaSepoliaVerify.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DEFAULTADMINROLE(&_LineaSepoliaVerify.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DEFAULTADMINROLE(&_LineaSepoliaVerify.CallOpts)
}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) GENESISSHNARF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "GENESIS_SHNARF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) GENESISSHNARF() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.GENESISSHNARF(&_LineaSepoliaVerify.CallOpts)
}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) GENESISSHNARF() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.GENESISSHNARF(&_LineaSepoliaVerify.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) INBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "INBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _LineaSepoliaVerify.Contract.INBOXSTATUSRECEIVED(&_LineaSepoliaVerify.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _LineaSepoliaVerify.Contract.INBOXSTATUSRECEIVED(&_LineaSepoliaVerify.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) INBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "INBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _LineaSepoliaVerify.Contract.INBOXSTATUSUNKNOWN(&_LineaSepoliaVerify.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _LineaSepoliaVerify.Contract.INBOXSTATUSUNKNOWN(&_LineaSepoliaVerify.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) OPERATORROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.OPERATORROLE(&_LineaSepoliaVerify.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) OPERATORROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.OPERATORROLE(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) OUTBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "OUTBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSRECEIVED(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSRECEIVED(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) OUTBOXSTATUSSENT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "OUTBOX_STATUS_SENT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) OUTBOXSTATUSSENT() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSSENT(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) OUTBOXSTATUSSENT() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSSENT(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) OUTBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "OUTBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSUNKNOWN(&_LineaSepoliaVerify.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _LineaSepoliaVerify.Contract.OUTBOXSTATUSUNKNOWN(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PAUSEALLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "PAUSE_ALL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PAUSEALLROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEALLROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PAUSEALLROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEALLROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0x12d3fa9a.
//
// Solidity: function PAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PAUSEBLOBSUBMISSIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "PAUSE_BLOB_SUBMISSION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0x12d3fa9a.
//
// Solidity: function PAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PAUSEBLOBSUBMISSIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEBLOBSUBMISSIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0x12d3fa9a.
//
// Solidity: function PAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PAUSEBLOBSUBMISSIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEBLOBSUBMISSIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x63213155.
//
// Solidity: function PAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PAUSEFINALIZATIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "PAUSE_FINALIZATION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x63213155.
//
// Solidity: function PAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PAUSEFINALIZATIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEFINALIZATIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x63213155.
//
// Solidity: function PAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PAUSEFINALIZATIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEFINALIZATIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xb9174ba3.
//
// Solidity: function PAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PAUSEL1L2ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "PAUSE_L1_L2_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xb9174ba3.
//
// Solidity: function PAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PAUSEL1L2ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEL1L2ROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xb9174ba3.
//
// Solidity: function PAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PAUSEL1L2ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEL1L2ROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x8de49487.
//
// Solidity: function PAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PAUSEL2L1ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "PAUSE_L2_L1_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x8de49487.
//
// Solidity: function PAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PAUSEL2L1ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEL2L1ROLE(&_LineaSepoliaVerify.CallOpts)
}

// PAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x8de49487.
//
// Solidity: function PAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PAUSEL2L1ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.PAUSEL2L1ROLE(&_LineaSepoliaVerify.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) RATELIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "RATE_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.RATELIMITSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.RATELIMITSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEALLROLE is a free data retrieval call binding the contract method 0x6a906b80.
//
// Solidity: function UNPAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) UNPAUSEALLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "UNPAUSE_ALL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEALLROLE is a free data retrieval call binding the contract method 0x6a906b80.
//
// Solidity: function UNPAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UNPAUSEALLROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEALLROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEALLROLE is a free data retrieval call binding the contract method 0x6a906b80.
//
// Solidity: function UNPAUSE_ALL_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) UNPAUSEALLROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEALLROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0xb59faa60.
//
// Solidity: function UNPAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) UNPAUSEBLOBSUBMISSIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "UNPAUSE_BLOB_SUBMISSION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0xb59faa60.
//
// Solidity: function UNPAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UNPAUSEBLOBSUBMISSIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEBLOBSUBMISSIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEBLOBSUBMISSIONROLE is a free data retrieval call binding the contract method 0xb59faa60.
//
// Solidity: function UNPAUSE_BLOB_SUBMISSION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) UNPAUSEBLOBSUBMISSIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEBLOBSUBMISSIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x03134d1d.
//
// Solidity: function UNPAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) UNPAUSEFINALIZATIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "UNPAUSE_FINALIZATION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x03134d1d.
//
// Solidity: function UNPAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UNPAUSEFINALIZATIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEFINALIZATIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEFINALIZATIONROLE is a free data retrieval call binding the contract method 0x03134d1d.
//
// Solidity: function UNPAUSE_FINALIZATION_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) UNPAUSEFINALIZATIONROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEFINALIZATIONROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xcc6f7251.
//
// Solidity: function UNPAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) UNPAUSEL1L2ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "UNPAUSE_L1_L2_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xcc6f7251.
//
// Solidity: function UNPAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UNPAUSEL1L2ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEL1L2ROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEL1L2ROLE is a free data retrieval call binding the contract method 0xcc6f7251.
//
// Solidity: function UNPAUSE_L1_L2_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) UNPAUSEL1L2ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEL1L2ROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x3b12eccb.
//
// Solidity: function UNPAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) UNPAUSEL2L1ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "UNPAUSE_L2_L1_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x3b12eccb.
//
// Solidity: function UNPAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UNPAUSEL2L1ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEL2L1ROLE(&_LineaSepoliaVerify.CallOpts)
}

// UNPAUSEL2L1ROLE is a free data retrieval call binding the contract method 0x3b12eccb.
//
// Solidity: function UNPAUSE_L2_L1_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) UNPAUSEL2L1ROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.UNPAUSEL2L1ROLE(&_LineaSepoliaVerify.CallOpts)
}

// USEDRATELIMITRESETTERROLE is a free data retrieval call binding the contract method 0x5230eef2.
//
// Solidity: function USED_RATE_LIMIT_RESETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) USEDRATELIMITRESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "USED_RATE_LIMIT_RESETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USEDRATELIMITRESETTERROLE is a free data retrieval call binding the contract method 0x5230eef2.
//
// Solidity: function USED_RATE_LIMIT_RESETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) USEDRATELIMITRESETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.USEDRATELIMITRESETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// USEDRATELIMITRESETTERROLE is a free data retrieval call binding the contract method 0x5230eef2.
//
// Solidity: function USED_RATE_LIMIT_RESETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) USEDRATELIMITRESETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.USEDRATELIMITRESETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) VERIFIERSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "VERIFIER_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) VERIFIERSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.VERIFIERSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) VERIFIERSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.VERIFIERSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// VERIFIERUNSETTERROLE is a free data retrieval call binding the contract method 0xd722bbfc.
//
// Solidity: function VERIFIER_UNSETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) VERIFIERUNSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "VERIFIER_UNSETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERUNSETTERROLE is a free data retrieval call binding the contract method 0xd722bbfc.
//
// Solidity: function VERIFIER_UNSETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) VERIFIERUNSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.VERIFIERUNSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// VERIFIERUNSETTERROLE is a free data retrieval call binding the contract method 0xd722bbfc.
//
// Solidity: function VERIFIER_UNSETTER_ROLE() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) VERIFIERUNSETTERROLE() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.VERIFIERUNSETTERROLE(&_LineaSepoliaVerify.CallOpts)
}

// BlobShnarfExists is a free data retrieval call binding the contract method 0x2130d812.
//
// Solidity: function blobShnarfExists(bytes32 shnarf) view returns(uint256 exists)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) BlobShnarfExists(opts *bind.CallOpts, shnarf [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "blobShnarfExists", shnarf)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlobShnarfExists is a free data retrieval call binding the contract method 0x2130d812.
//
// Solidity: function blobShnarfExists(bytes32 shnarf) view returns(uint256 exists)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) BlobShnarfExists(shnarf [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.BlobShnarfExists(&_LineaSepoliaVerify.CallOpts, shnarf)
}

// BlobShnarfExists is a free data retrieval call binding the contract method 0x2130d812.
//
// Solidity: function blobShnarfExists(bytes32 shnarf) view returns(uint256 exists)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) BlobShnarfExists(shnarf [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.BlobShnarfExists(&_LineaSepoliaVerify.CallOpts, shnarf)
}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentFinalizedShnarf(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentFinalizedShnarf")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentFinalizedShnarf() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentFinalizedShnarf(&_LineaSepoliaVerify.CallOpts)
}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentFinalizedShnarf() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentFinalizedShnarf(&_LineaSepoliaVerify.CallOpts)
}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentFinalizedState(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentFinalizedState")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentFinalizedState() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentFinalizedState(&_LineaSepoliaVerify.CallOpts)
}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentFinalizedState() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentFinalizedState(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentL2BlockNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2BlockNumber(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentL2BlockNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2BlockNumber(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentL2StoredL1MessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentL2StoredL1MessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentL2StoredL1MessageNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2StoredL1MessageNumber(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentL2StoredL1MessageNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2StoredL1MessageNumber(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentL2StoredL1RollingHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentL2StoredL1RollingHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentL2StoredL1RollingHash() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2StoredL1RollingHash(&_LineaSepoliaVerify.CallOpts)
}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentL2StoredL1RollingHash() ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.CurrentL2StoredL1RollingHash(&_LineaSepoliaVerify.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentPeriodAmountInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentPeriodAmountInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentPeriodAmountInWei(&_LineaSepoliaVerify.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentPeriodAmountInWei(&_LineaSepoliaVerify.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentPeriodEnd() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentPeriodEnd(&_LineaSepoliaVerify.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentPeriodEnd() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentPeriodEnd(&_LineaSepoliaVerify.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) CurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "currentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) CurrentTimestamp() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentTimestamp(&_LineaSepoliaVerify.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) CurrentTimestamp() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.CurrentTimestamp(&_LineaSepoliaVerify.CallOpts)
}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DataEndingBlock(opts *bind.CallOpts, dataHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "dataEndingBlock", dataHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DataEndingBlock(dataHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.DataEndingBlock(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DataEndingBlock(dataHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.DataEndingBlock(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DataFinalStateRootHashes(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "dataFinalStateRootHashes", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DataFinalStateRootHashes(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataFinalStateRootHashes(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DataFinalStateRootHashes(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataFinalStateRootHashes(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DataParents(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "dataParents", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DataParents(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataParents(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DataParents(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataParents(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DataShnarfHashes(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "dataShnarfHashes", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DataShnarfHashes(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataShnarfHashes(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DataShnarfHashes(dataHash [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.DataShnarfHashes(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) DataStartingBlock(opts *bind.CallOpts, dataHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "dataStartingBlock", dataHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) DataStartingBlock(dataHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.DataStartingBlock(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) DataStartingBlock(dataHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.DataStartingBlock(&_LineaSepoliaVerify.CallOpts, dataHash)
}

// FallbackOperator is a free data retrieval call binding the contract method 0xcf5b2764.
//
// Solidity: function fallbackOperator() view returns(address)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) FallbackOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "fallbackOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FallbackOperator is a free data retrieval call binding the contract method 0xcf5b2764.
//
// Solidity: function fallbackOperator() view returns(address)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) FallbackOperator() (common.Address, error) {
	return _LineaSepoliaVerify.Contract.FallbackOperator(&_LineaSepoliaVerify.CallOpts)
}

// FallbackOperator is a free data retrieval call binding the contract method 0xcf5b2764.
//
// Solidity: function fallbackOperator() view returns(address)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) FallbackOperator() (common.Address, error) {
	return _LineaSepoliaVerify.Contract.FallbackOperator(&_LineaSepoliaVerify.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.GetRoleAdmin(&_LineaSepoliaVerify.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.GetRoleAdmin(&_LineaSepoliaVerify.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LineaSepoliaVerify.Contract.HasRole(&_LineaSepoliaVerify.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LineaSepoliaVerify.Contract.HasRole(&_LineaSepoliaVerify.CallOpts, role, account)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) InboxL2L1MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "inboxL2L1MessageStatus", messageHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) InboxL2L1MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.InboxL2L1MessageStatus(&_LineaSepoliaVerify.CallOpts, messageHash)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) InboxL2L1MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.InboxL2L1MessageStatus(&_LineaSepoliaVerify.CallOpts, messageHash)
}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool isClaimed)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) IsMessageClaimed(opts *bind.CallOpts, _messageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "isMessageClaimed", _messageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool isClaimed)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) IsMessageClaimed(_messageNumber *big.Int) (bool, error) {
	return _LineaSepoliaVerify.Contract.IsMessageClaimed(&_LineaSepoliaVerify.CallOpts, _messageNumber)
}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool isClaimed)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) IsMessageClaimed(_messageNumber *big.Int) (bool, error) {
	return _LineaSepoliaVerify.Contract.IsMessageClaimed(&_LineaSepoliaVerify.CallOpts, _messageNumber)
}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool pauseTypeIsPaused)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) IsPaused(opts *bind.CallOpts, _pauseType uint8) (bool, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "isPaused", _pauseType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool pauseTypeIsPaused)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) IsPaused(_pauseType uint8) (bool, error) {
	return _LineaSepoliaVerify.Contract.IsPaused(&_LineaSepoliaVerify.CallOpts, _pauseType)
}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool pauseTypeIsPaused)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) IsPaused(_pauseType uint8) (bool, error) {
	return _LineaSepoliaVerify.Contract.IsPaused(&_LineaSepoliaVerify.CallOpts, _pauseType)
}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) L2MerkleRootsDepths(opts *bind.CallOpts, merkleRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "l2MerkleRootsDepths", merkleRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) L2MerkleRootsDepths(merkleRoot [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.L2MerkleRootsDepths(&_LineaSepoliaVerify.CallOpts, merkleRoot)
}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) L2MerkleRootsDepths(merkleRoot [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.L2MerkleRootsDepths(&_LineaSepoliaVerify.CallOpts, merkleRoot)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) LimitInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "limitInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) LimitInWei() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.LimitInWei(&_LineaSepoliaVerify.CallOpts)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) LimitInWei() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.LimitInWei(&_LineaSepoliaVerify.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) NextMessageNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.NextMessageNumber(&_LineaSepoliaVerify.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) NextMessageNumber() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.NextMessageNumber(&_LineaSepoliaVerify.CallOpts)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) OutboxL1L2MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "outboxL1L2MessageStatus", messageHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) OutboxL1L2MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.OutboxL1L2MessageStatus(&_LineaSepoliaVerify.CallOpts, messageHash)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) OutboxL1L2MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.OutboxL1L2MessageStatus(&_LineaSepoliaVerify.CallOpts, messageHash)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PauseTypeStatuses(opts *bind.CallOpts, pauseType [32]byte) (bool, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "pauseTypeStatuses", pauseType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PauseTypeStatuses(pauseType [32]byte) (bool, error) {
	return _LineaSepoliaVerify.Contract.PauseTypeStatuses(&_LineaSepoliaVerify.CallOpts, pauseType)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PauseTypeStatuses(pauseType [32]byte) (bool, error) {
	return _LineaSepoliaVerify.Contract.PauseTypeStatuses(&_LineaSepoliaVerify.CallOpts, pauseType)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) PeriodInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "periodInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PeriodInSeconds() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.PeriodInSeconds(&_LineaSepoliaVerify.CallOpts)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) PeriodInSeconds() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.PeriodInSeconds(&_LineaSepoliaVerify.CallOpts)
}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) RollingHashes(opts *bind.CallOpts, messageNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "rollingHashes", messageNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) RollingHashes(messageNumber *big.Int) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.RollingHashes(&_LineaSepoliaVerify.CallOpts, messageNumber)
}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) RollingHashes(messageNumber *big.Int) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.RollingHashes(&_LineaSepoliaVerify.CallOpts, messageNumber)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address originalSender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address originalSender)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) Sender() (common.Address, error) {
	return _LineaSepoliaVerify.Contract.Sender(&_LineaSepoliaVerify.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address originalSender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) Sender() (common.Address, error) {
	return _LineaSepoliaVerify.Contract.Sender(&_LineaSepoliaVerify.CallOpts)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) StateRootHashes(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "stateRootHashes", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) StateRootHashes(blockNumber *big.Int) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.StateRootHashes(&_LineaSepoliaVerify.CallOpts, blockNumber)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) StateRootHashes(blockNumber *big.Int) ([32]byte, error) {
	return _LineaSepoliaVerify.Contract.StateRootHashes(&_LineaSepoliaVerify.CallOpts, blockNumber)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LineaSepoliaVerify.Contract.SupportsInterface(&_LineaSepoliaVerify.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LineaSepoliaVerify.Contract.SupportsInterface(&_LineaSepoliaVerify.CallOpts, interfaceId)
}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) SystemMigrationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "systemMigrationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SystemMigrationBlock() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.SystemMigrationBlock(&_LineaSepoliaVerify.CallOpts)
}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) SystemMigrationBlock() (*big.Int, error) {
	return _LineaSepoliaVerify.Contract.SystemMigrationBlock(&_LineaSepoliaVerify.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCaller) Verifiers(opts *bind.CallOpts, proofType *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LineaSepoliaVerify.contract.Call(opts, &out, "verifiers", proofType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) Verifiers(proofType *big.Int) (common.Address, error) {
	return _LineaSepoliaVerify.Contract.Verifiers(&_LineaSepoliaVerify.CallOpts, proofType)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyCallerSession) Verifiers(proofType *big.Int) (common.Address, error) {
	return _LineaSepoliaVerify.Contract.Verifiers(&_LineaSepoliaVerify.CallOpts, proofType)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) ClaimMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "claimMessage", _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ClaimMessage(&_LineaSepoliaVerify.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ClaimMessage(&_LineaSepoliaVerify.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) ClaimMessageWithProof(opts *bind.TransactOpts, _params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "claimMessageWithProof", _params)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) ClaimMessageWithProof(_params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ClaimMessageWithProof(&_LineaSepoliaVerify.TransactOpts, _params)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) ClaimMessageWithProof(_params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ClaimMessageWithProof(&_LineaSepoliaVerify.TransactOpts, _params)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5603c65f.
//
// Solidity: function finalizeBlocks(bytes _aggregatedProof, uint256 _proofType, (bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) FinalizeBlocks(opts *bind.TransactOpts, _aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV3) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "finalizeBlocks", _aggregatedProof, _proofType, _finalizationData)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5603c65f.
//
// Solidity: function finalizeBlocks(bytes _aggregatedProof, uint256 _proofType, (bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) FinalizeBlocks(_aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV3) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.FinalizeBlocks(&_LineaSepoliaVerify.TransactOpts, _aggregatedProof, _proofType, _finalizationData)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5603c65f.
//
// Solidity: function finalizeBlocks(bytes _aggregatedProof, uint256 _proofType, (bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) FinalizeBlocks(_aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV3) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.FinalizeBlocks(&_LineaSepoliaVerify.TransactOpts, _aggregatedProof, _proofType, _finalizationData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.GrantRole(&_LineaSepoliaVerify.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.GrantRole(&_LineaSepoliaVerify.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa98e773d.
//
// Solidity: function initialize((bytes32,uint256,uint256,address,uint256,uint256,(address,bytes32)[],(uint8,bytes32)[],(uint8,bytes32)[],address,address) _initializationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) Initialize(opts *bind.TransactOpts, _initializationData ILineaRollupInitializationData) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "initialize", _initializationData)
}

// Initialize is a paid mutator transaction binding the contract method 0xa98e773d.
//
// Solidity: function initialize((bytes32,uint256,uint256,address,uint256,uint256,(address,bytes32)[],(uint8,bytes32)[],(uint8,bytes32)[],address,address) _initializationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) Initialize(_initializationData ILineaRollupInitializationData) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.Initialize(&_LineaSepoliaVerify.TransactOpts, _initializationData)
}

// Initialize is a paid mutator transaction binding the contract method 0xa98e773d.
//
// Solidity: function initialize((bytes32,uint256,uint256,address,uint256,uint256,(address,bytes32)[],(uint8,bytes32)[],(uint8,bytes32)[],address,address) _initializationData) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) Initialize(_initializationData ILineaRollupInitializationData) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.Initialize(&_LineaSepoliaVerify.TransactOpts, _initializationData)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) PauseByType(opts *bind.TransactOpts, _pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "pauseByType", _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) PauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.PauseByType(&_LineaSepoliaVerify.TransactOpts, _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) PauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.PauseByType(&_LineaSepoliaVerify.TransactOpts, _pauseType)
}

// ReinitializeLineaRollupV6 is a paid mutator transaction binding the contract method 0xc0c4e584.
//
// Solidity: function reinitializeLineaRollupV6((address,bytes32)[] _roleAddresses, (uint8,bytes32)[] _pauseTypeRoles, (uint8,bytes32)[] _unpauseTypeRoles, address _fallbackOperator) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) ReinitializeLineaRollupV6(opts *bind.TransactOpts, _roleAddresses []IPermissionsManagerRoleAddress, _pauseTypeRoles []IPauseManagerPauseTypeRole, _unpauseTypeRoles []IPauseManagerPauseTypeRole, _fallbackOperator common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "reinitializeLineaRollupV6", _roleAddresses, _pauseTypeRoles, _unpauseTypeRoles, _fallbackOperator)
}

// ReinitializeLineaRollupV6 is a paid mutator transaction binding the contract method 0xc0c4e584.
//
// Solidity: function reinitializeLineaRollupV6((address,bytes32)[] _roleAddresses, (uint8,bytes32)[] _pauseTypeRoles, (uint8,bytes32)[] _unpauseTypeRoles, address _fallbackOperator) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) ReinitializeLineaRollupV6(_roleAddresses []IPermissionsManagerRoleAddress, _pauseTypeRoles []IPauseManagerPauseTypeRole, _unpauseTypeRoles []IPauseManagerPauseTypeRole, _fallbackOperator common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ReinitializeLineaRollupV6(&_LineaSepoliaVerify.TransactOpts, _roleAddresses, _pauseTypeRoles, _unpauseTypeRoles, _fallbackOperator)
}

// ReinitializeLineaRollupV6 is a paid mutator transaction binding the contract method 0xc0c4e584.
//
// Solidity: function reinitializeLineaRollupV6((address,bytes32)[] _roleAddresses, (uint8,bytes32)[] _pauseTypeRoles, (uint8,bytes32)[] _unpauseTypeRoles, address _fallbackOperator) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) ReinitializeLineaRollupV6(_roleAddresses []IPermissionsManagerRoleAddress, _pauseTypeRoles []IPauseManagerPauseTypeRole, _unpauseTypeRoles []IPauseManagerPauseTypeRole, _fallbackOperator common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ReinitializeLineaRollupV6(&_LineaSepoliaVerify.TransactOpts, _roleAddresses, _pauseTypeRoles, _unpauseTypeRoles, _fallbackOperator)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) RenounceRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "renounceRole", _role, _account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) RenounceRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.RenounceRole(&_LineaSepoliaVerify.TransactOpts, _role, _account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) RenounceRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.RenounceRole(&_LineaSepoliaVerify.TransactOpts, _role, _account)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) ResetAmountUsedInPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "resetAmountUsedInPeriod")
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ResetAmountUsedInPeriod(&_LineaSepoliaVerify.TransactOpts)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ResetAmountUsedInPeriod(&_LineaSepoliaVerify.TransactOpts)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) ResetRateLimitAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "resetRateLimitAmount", _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ResetRateLimitAmount(&_LineaSepoliaVerify.TransactOpts, _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.ResetRateLimitAmount(&_LineaSepoliaVerify.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.RevokeRole(&_LineaSepoliaVerify.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.RevokeRole(&_LineaSepoliaVerify.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "sendMessage", _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SendMessage(&_LineaSepoliaVerify.TransactOpts, _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SendMessage(&_LineaSepoliaVerify.TransactOpts, _to, _fee, _calldata)
}

// SetFallbackOperator is a paid mutator transaction binding the contract method 0xbcc3003d.
//
// Solidity: function setFallbackOperator(uint256 _messageNumber, bytes32 _rollingHash, uint256 _lastFinalizedTimestamp) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) SetFallbackOperator(opts *bind.TransactOpts, _messageNumber *big.Int, _rollingHash [32]byte, _lastFinalizedTimestamp *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "setFallbackOperator", _messageNumber, _rollingHash, _lastFinalizedTimestamp)
}

// SetFallbackOperator is a paid mutator transaction binding the contract method 0xbcc3003d.
//
// Solidity: function setFallbackOperator(uint256 _messageNumber, bytes32 _rollingHash, uint256 _lastFinalizedTimestamp) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SetFallbackOperator(_messageNumber *big.Int, _rollingHash [32]byte, _lastFinalizedTimestamp *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SetFallbackOperator(&_LineaSepoliaVerify.TransactOpts, _messageNumber, _rollingHash, _lastFinalizedTimestamp)
}

// SetFallbackOperator is a paid mutator transaction binding the contract method 0xbcc3003d.
//
// Solidity: function setFallbackOperator(uint256 _messageNumber, bytes32 _rollingHash, uint256 _lastFinalizedTimestamp) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) SetFallbackOperator(_messageNumber *big.Int, _rollingHash [32]byte, _lastFinalizedTimestamp *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SetFallbackOperator(&_LineaSepoliaVerify.TransactOpts, _messageNumber, _rollingHash, _lastFinalizedTimestamp)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) SetVerifierAddress(opts *bind.TransactOpts, _newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "setVerifierAddress", _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SetVerifierAddress(&_LineaSepoliaVerify.TransactOpts, _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SetVerifierAddress(&_LineaSepoliaVerify.TransactOpts, _newVerifierAddress, _proofType)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x99467a35.
//
// Solidity: function submitBlobs((uint256,bytes,bytes,bytes32,bytes32)[] _blobSubmissions, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) SubmitBlobs(opts *bind.TransactOpts, _blobSubmissions []ILineaRollupBlobSubmission, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "submitBlobs", _blobSubmissions, _parentShnarf, _finalBlobShnarf)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x99467a35.
//
// Solidity: function submitBlobs((uint256,bytes,bytes,bytes32,bytes32)[] _blobSubmissions, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SubmitBlobs(_blobSubmissions []ILineaRollupBlobSubmission, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SubmitBlobs(&_LineaSepoliaVerify.TransactOpts, _blobSubmissions, _parentShnarf, _finalBlobShnarf)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x99467a35.
//
// Solidity: function submitBlobs((uint256,bytes,bytes,bytes32,bytes32)[] _blobSubmissions, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) SubmitBlobs(_blobSubmissions []ILineaRollupBlobSubmission, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SubmitBlobs(&_LineaSepoliaVerify.TransactOpts, _blobSubmissions, _parentShnarf, _finalBlobShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0x6854f6bc.
//
// Solidity: function submitDataAsCalldata((bytes32,bytes32,bytes) _submission, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) SubmitDataAsCalldata(opts *bind.TransactOpts, _submission ILineaRollupCompressedCalldataSubmission, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "submitDataAsCalldata", _submission, _parentShnarf, _expectedShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0x6854f6bc.
//
// Solidity: function submitDataAsCalldata((bytes32,bytes32,bytes) _submission, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) SubmitDataAsCalldata(_submission ILineaRollupCompressedCalldataSubmission, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SubmitDataAsCalldata(&_LineaSepoliaVerify.TransactOpts, _submission, _parentShnarf, _expectedShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0x6854f6bc.
//
// Solidity: function submitDataAsCalldata((bytes32,bytes32,bytes) _submission, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) SubmitDataAsCalldata(_submission ILineaRollupCompressedCalldataSubmission, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.SubmitDataAsCalldata(&_LineaSepoliaVerify.TransactOpts, _submission, _parentShnarf, _expectedShnarf)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) UnPauseByType(opts *bind.TransactOpts, _pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "unPauseByType", _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UnPauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.UnPauseByType(&_LineaSepoliaVerify.TransactOpts, _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) UnPauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.UnPauseByType(&_LineaSepoliaVerify.TransactOpts, _pauseType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactor) UnsetVerifierAddress(opts *bind.TransactOpts, _proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.contract.Transact(opts, "unsetVerifierAddress", _proofType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifySession) UnsetVerifierAddress(_proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.UnsetVerifierAddress(&_LineaSepoliaVerify.TransactOpts, _proofType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_LineaSepoliaVerify *LineaSepoliaVerifyTransactorSession) UnsetVerifierAddress(_proofType *big.Int) (*types.Transaction, error) {
	return _LineaSepoliaVerify.Contract.UnsetVerifierAddress(&_LineaSepoliaVerify.TransactOpts, _proofType)
}

// LineaSepoliaVerifyAmountUsedInPeriodResetIterator is returned from FilterAmountUsedInPeriodReset and is used to iterate over the raw logs and unpacked data for AmountUsedInPeriodReset events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyAmountUsedInPeriodResetIterator struct {
	Event *LineaSepoliaVerifyAmountUsedInPeriodReset // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyAmountUsedInPeriodResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyAmountUsedInPeriodReset)
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
		it.Event = new(LineaSepoliaVerifyAmountUsedInPeriodReset)
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
func (it *LineaSepoliaVerifyAmountUsedInPeriodResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyAmountUsedInPeriodResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyAmountUsedInPeriodReset represents a AmountUsedInPeriodReset event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyAmountUsedInPeriodReset struct {
	ResettingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAmountUsedInPeriodReset is a free log retrieval operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterAmountUsedInPeriodReset(opts *bind.FilterOpts, resettingAddress []common.Address) (*LineaSepoliaVerifyAmountUsedInPeriodResetIterator, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyAmountUsedInPeriodResetIterator{contract: _LineaSepoliaVerify.contract, event: "AmountUsedInPeriodReset", logs: logs, sub: sub}, nil
}

// WatchAmountUsedInPeriodReset is a free log subscription operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchAmountUsedInPeriodReset(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyAmountUsedInPeriodReset, resettingAddress []common.Address) (event.Subscription, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyAmountUsedInPeriodReset)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
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

// ParseAmountUsedInPeriodReset is a log parse operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseAmountUsedInPeriodReset(log types.Log) (*LineaSepoliaVerifyAmountUsedInPeriodReset, error) {
	event := new(LineaSepoliaVerifyAmountUsedInPeriodReset)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyDataFinalizedV3Iterator is returned from FilterDataFinalizedV3 and is used to iterate over the raw logs and unpacked data for DataFinalizedV3 events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyDataFinalizedV3Iterator struct {
	Event *LineaSepoliaVerifyDataFinalizedV3 // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyDataFinalizedV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyDataFinalizedV3)
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
		it.Event = new(LineaSepoliaVerifyDataFinalizedV3)
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
func (it *LineaSepoliaVerifyDataFinalizedV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyDataFinalizedV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyDataFinalizedV3 represents a DataFinalizedV3 event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyDataFinalizedV3 struct {
	StartBlockNumber    *big.Int
	EndBlockNumber      *big.Int
	Shnarf              [32]byte
	ParentStateRootHash [32]byte
	FinalStateRootHash  [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDataFinalizedV3 is a free log retrieval operation binding the contract event 0xa0262dc79e4ccb71ceac8574ae906311ae338aa4a2044fd4ec4b99fad5ab60cb.
//
// Solidity: event DataFinalizedV3(uint256 indexed startBlockNumber, uint256 indexed endBlockNumber, bytes32 indexed shnarf, bytes32 parentStateRootHash, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterDataFinalizedV3(opts *bind.FilterOpts, startBlockNumber []*big.Int, endBlockNumber []*big.Int, shnarf [][32]byte) (*LineaSepoliaVerifyDataFinalizedV3Iterator, error) {

	var startBlockNumberRule []interface{}
	for _, startBlockNumberItem := range startBlockNumber {
		startBlockNumberRule = append(startBlockNumberRule, startBlockNumberItem)
	}
	var endBlockNumberRule []interface{}
	for _, endBlockNumberItem := range endBlockNumber {
		endBlockNumberRule = append(endBlockNumberRule, endBlockNumberItem)
	}
	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "DataFinalizedV3", startBlockNumberRule, endBlockNumberRule, shnarfRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyDataFinalizedV3Iterator{contract: _LineaSepoliaVerify.contract, event: "DataFinalizedV3", logs: logs, sub: sub}, nil
}

// WatchDataFinalizedV3 is a free log subscription operation binding the contract event 0xa0262dc79e4ccb71ceac8574ae906311ae338aa4a2044fd4ec4b99fad5ab60cb.
//
// Solidity: event DataFinalizedV3(uint256 indexed startBlockNumber, uint256 indexed endBlockNumber, bytes32 indexed shnarf, bytes32 parentStateRootHash, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchDataFinalizedV3(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyDataFinalizedV3, startBlockNumber []*big.Int, endBlockNumber []*big.Int, shnarf [][32]byte) (event.Subscription, error) {

	var startBlockNumberRule []interface{}
	for _, startBlockNumberItem := range startBlockNumber {
		startBlockNumberRule = append(startBlockNumberRule, startBlockNumberItem)
	}
	var endBlockNumberRule []interface{}
	for _, endBlockNumberItem := range endBlockNumber {
		endBlockNumberRule = append(endBlockNumberRule, endBlockNumberItem)
	}
	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "DataFinalizedV3", startBlockNumberRule, endBlockNumberRule, shnarfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyDataFinalizedV3)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "DataFinalizedV3", log); err != nil {
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

// ParseDataFinalizedV3 is a log parse operation binding the contract event 0xa0262dc79e4ccb71ceac8574ae906311ae338aa4a2044fd4ec4b99fad5ab60cb.
//
// Solidity: event DataFinalizedV3(uint256 indexed startBlockNumber, uint256 indexed endBlockNumber, bytes32 indexed shnarf, bytes32 parentStateRootHash, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseDataFinalizedV3(log types.Log) (*LineaSepoliaVerifyDataFinalizedV3, error) {
	event := new(LineaSepoliaVerifyDataFinalizedV3)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "DataFinalizedV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyDataSubmittedV3Iterator is returned from FilterDataSubmittedV3 and is used to iterate over the raw logs and unpacked data for DataSubmittedV3 events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyDataSubmittedV3Iterator struct {
	Event *LineaSepoliaVerifyDataSubmittedV3 // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyDataSubmittedV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyDataSubmittedV3)
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
		it.Event = new(LineaSepoliaVerifyDataSubmittedV3)
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
func (it *LineaSepoliaVerifyDataSubmittedV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyDataSubmittedV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyDataSubmittedV3 represents a DataSubmittedV3 event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyDataSubmittedV3 struct {
	ParentShnarf       [32]byte
	Shnarf             [32]byte
	FinalStateRootHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDataSubmittedV3 is a free log retrieval operation binding the contract event 0x55f4c645c36aa5cd3f443d6be44d7a7a5df9d2100d7139dfc69d4289ee072319.
//
// Solidity: event DataSubmittedV3(bytes32 parentShnarf, bytes32 indexed shnarf, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterDataSubmittedV3(opts *bind.FilterOpts, shnarf [][32]byte) (*LineaSepoliaVerifyDataSubmittedV3Iterator, error) {

	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "DataSubmittedV3", shnarfRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyDataSubmittedV3Iterator{contract: _LineaSepoliaVerify.contract, event: "DataSubmittedV3", logs: logs, sub: sub}, nil
}

// WatchDataSubmittedV3 is a free log subscription operation binding the contract event 0x55f4c645c36aa5cd3f443d6be44d7a7a5df9d2100d7139dfc69d4289ee072319.
//
// Solidity: event DataSubmittedV3(bytes32 parentShnarf, bytes32 indexed shnarf, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchDataSubmittedV3(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyDataSubmittedV3, shnarf [][32]byte) (event.Subscription, error) {

	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "DataSubmittedV3", shnarfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyDataSubmittedV3)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "DataSubmittedV3", log); err != nil {
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

// ParseDataSubmittedV3 is a log parse operation binding the contract event 0x55f4c645c36aa5cd3f443d6be44d7a7a5df9d2100d7139dfc69d4289ee072319.
//
// Solidity: event DataSubmittedV3(bytes32 parentShnarf, bytes32 indexed shnarf, bytes32 finalStateRootHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseDataSubmittedV3(log types.Log) (*LineaSepoliaVerifyDataSubmittedV3, error) {
	event := new(LineaSepoliaVerifyDataSubmittedV3)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "DataSubmittedV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyFallbackOperatorAddressSetIterator is returned from FilterFallbackOperatorAddressSet and is used to iterate over the raw logs and unpacked data for FallbackOperatorAddressSet events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyFallbackOperatorAddressSetIterator struct {
	Event *LineaSepoliaVerifyFallbackOperatorAddressSet // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyFallbackOperatorAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyFallbackOperatorAddressSet)
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
		it.Event = new(LineaSepoliaVerifyFallbackOperatorAddressSet)
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
func (it *LineaSepoliaVerifyFallbackOperatorAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyFallbackOperatorAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyFallbackOperatorAddressSet represents a FallbackOperatorAddressSet event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyFallbackOperatorAddressSet struct {
	Caller           common.Address
	FallbackOperator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFallbackOperatorAddressSet is a free log retrieval operation binding the contract event 0x1f82add12d98b5eaed4d6a6d5f74cfc7a85e5c90c335ab5562f77f220ed45d5f.
//
// Solidity: event FallbackOperatorAddressSet(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterFallbackOperatorAddressSet(opts *bind.FilterOpts, caller []common.Address, fallbackOperator []common.Address) (*LineaSepoliaVerifyFallbackOperatorAddressSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fallbackOperatorRule []interface{}
	for _, fallbackOperatorItem := range fallbackOperator {
		fallbackOperatorRule = append(fallbackOperatorRule, fallbackOperatorItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "FallbackOperatorAddressSet", callerRule, fallbackOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyFallbackOperatorAddressSetIterator{contract: _LineaSepoliaVerify.contract, event: "FallbackOperatorAddressSet", logs: logs, sub: sub}, nil
}

// WatchFallbackOperatorAddressSet is a free log subscription operation binding the contract event 0x1f82add12d98b5eaed4d6a6d5f74cfc7a85e5c90c335ab5562f77f220ed45d5f.
//
// Solidity: event FallbackOperatorAddressSet(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchFallbackOperatorAddressSet(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyFallbackOperatorAddressSet, caller []common.Address, fallbackOperator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fallbackOperatorRule []interface{}
	for _, fallbackOperatorItem := range fallbackOperator {
		fallbackOperatorRule = append(fallbackOperatorRule, fallbackOperatorItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "FallbackOperatorAddressSet", callerRule, fallbackOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyFallbackOperatorAddressSet)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "FallbackOperatorAddressSet", log); err != nil {
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

// ParseFallbackOperatorAddressSet is a log parse operation binding the contract event 0x1f82add12d98b5eaed4d6a6d5f74cfc7a85e5c90c335ab5562f77f220ed45d5f.
//
// Solidity: event FallbackOperatorAddressSet(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseFallbackOperatorAddressSet(log types.Log) (*LineaSepoliaVerifyFallbackOperatorAddressSet, error) {
	event := new(LineaSepoliaVerifyFallbackOperatorAddressSet)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "FallbackOperatorAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator is returned from FilterFallbackOperatorRoleGranted and is used to iterate over the raw logs and unpacked data for FallbackOperatorRoleGranted events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator struct {
	Event *LineaSepoliaVerifyFallbackOperatorRoleGranted // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyFallbackOperatorRoleGranted)
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
		it.Event = new(LineaSepoliaVerifyFallbackOperatorRoleGranted)
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
func (it *LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyFallbackOperatorRoleGranted represents a FallbackOperatorRoleGranted event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyFallbackOperatorRoleGranted struct {
	Caller           common.Address
	FallbackOperator common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFallbackOperatorRoleGranted is a free log retrieval operation binding the contract event 0x9fc8868f8577b31b805ee65bb52325782b5e2708dbdb7f04c7467c6785fccb30.
//
// Solidity: event FallbackOperatorRoleGranted(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterFallbackOperatorRoleGranted(opts *bind.FilterOpts, caller []common.Address, fallbackOperator []common.Address) (*LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fallbackOperatorRule []interface{}
	for _, fallbackOperatorItem := range fallbackOperator {
		fallbackOperatorRule = append(fallbackOperatorRule, fallbackOperatorItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "FallbackOperatorRoleGranted", callerRule, fallbackOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyFallbackOperatorRoleGrantedIterator{contract: _LineaSepoliaVerify.contract, event: "FallbackOperatorRoleGranted", logs: logs, sub: sub}, nil
}

// WatchFallbackOperatorRoleGranted is a free log subscription operation binding the contract event 0x9fc8868f8577b31b805ee65bb52325782b5e2708dbdb7f04c7467c6785fccb30.
//
// Solidity: event FallbackOperatorRoleGranted(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchFallbackOperatorRoleGranted(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyFallbackOperatorRoleGranted, caller []common.Address, fallbackOperator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fallbackOperatorRule []interface{}
	for _, fallbackOperatorItem := range fallbackOperator {
		fallbackOperatorRule = append(fallbackOperatorRule, fallbackOperatorItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "FallbackOperatorRoleGranted", callerRule, fallbackOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyFallbackOperatorRoleGranted)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "FallbackOperatorRoleGranted", log); err != nil {
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

// ParseFallbackOperatorRoleGranted is a log parse operation binding the contract event 0x9fc8868f8577b31b805ee65bb52325782b5e2708dbdb7f04c7467c6785fccb30.
//
// Solidity: event FallbackOperatorRoleGranted(address indexed caller, address indexed fallbackOperator)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseFallbackOperatorRoleGranted(log types.Log) (*LineaSepoliaVerifyFallbackOperatorRoleGranted, error) {
	event := new(LineaSepoliaVerifyFallbackOperatorRoleGranted)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "FallbackOperatorRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyInitializedIterator struct {
	Event *LineaSepoliaVerifyInitialized // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyInitialized)
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
		it.Event = new(LineaSepoliaVerifyInitialized)
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
func (it *LineaSepoliaVerifyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyInitialized represents a Initialized event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterInitialized(opts *bind.FilterOpts) (*LineaSepoliaVerifyInitializedIterator, error) {

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyInitializedIterator{contract: _LineaSepoliaVerify.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyInitialized) (event.Subscription, error) {

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyInitialized)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseInitialized(log types.Log) (*LineaSepoliaVerifyInitialized, error) {
	event := new(LineaSepoliaVerifyInitialized)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyL2MerkleRootAddedIterator is returned from FilterL2MerkleRootAdded and is used to iterate over the raw logs and unpacked data for L2MerkleRootAdded events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyL2MerkleRootAddedIterator struct {
	Event *LineaSepoliaVerifyL2MerkleRootAdded // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyL2MerkleRootAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyL2MerkleRootAdded)
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
		it.Event = new(LineaSepoliaVerifyL2MerkleRootAdded)
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
func (it *LineaSepoliaVerifyL2MerkleRootAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyL2MerkleRootAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyL2MerkleRootAdded represents a L2MerkleRootAdded event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyL2MerkleRootAdded struct {
	L2MerkleRoot [32]byte
	TreeDepth    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterL2MerkleRootAdded is a free log retrieval operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterL2MerkleRootAdded(opts *bind.FilterOpts, l2MerkleRoot [][32]byte, treeDepth []*big.Int) (*LineaSepoliaVerifyL2MerkleRootAddedIterator, error) {

	var l2MerkleRootRule []interface{}
	for _, l2MerkleRootItem := range l2MerkleRoot {
		l2MerkleRootRule = append(l2MerkleRootRule, l2MerkleRootItem)
	}
	var treeDepthRule []interface{}
	for _, treeDepthItem := range treeDepth {
		treeDepthRule = append(treeDepthRule, treeDepthItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "L2MerkleRootAdded", l2MerkleRootRule, treeDepthRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyL2MerkleRootAddedIterator{contract: _LineaSepoliaVerify.contract, event: "L2MerkleRootAdded", logs: logs, sub: sub}, nil
}

// WatchL2MerkleRootAdded is a free log subscription operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchL2MerkleRootAdded(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyL2MerkleRootAdded, l2MerkleRoot [][32]byte, treeDepth []*big.Int) (event.Subscription, error) {

	var l2MerkleRootRule []interface{}
	for _, l2MerkleRootItem := range l2MerkleRoot {
		l2MerkleRootRule = append(l2MerkleRootRule, l2MerkleRootItem)
	}
	var treeDepthRule []interface{}
	for _, treeDepthItem := range treeDepth {
		treeDepthRule = append(treeDepthRule, treeDepthItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "L2MerkleRootAdded", l2MerkleRootRule, treeDepthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyL2MerkleRootAdded)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "L2MerkleRootAdded", log); err != nil {
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

// ParseL2MerkleRootAdded is a log parse operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseL2MerkleRootAdded(log types.Log) (*LineaSepoliaVerifyL2MerkleRootAdded, error) {
	event := new(LineaSepoliaVerifyL2MerkleRootAdded)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "L2MerkleRootAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyL2MessagingBlockAnchoredIterator is returned from FilterL2MessagingBlockAnchored and is used to iterate over the raw logs and unpacked data for L2MessagingBlockAnchored events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyL2MessagingBlockAnchoredIterator struct {
	Event *LineaSepoliaVerifyL2MessagingBlockAnchored // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyL2MessagingBlockAnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyL2MessagingBlockAnchored)
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
		it.Event = new(LineaSepoliaVerifyL2MessagingBlockAnchored)
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
func (it *LineaSepoliaVerifyL2MessagingBlockAnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyL2MessagingBlockAnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyL2MessagingBlockAnchored represents a L2MessagingBlockAnchored event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyL2MessagingBlockAnchored struct {
	L2Block *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterL2MessagingBlockAnchored is a free log retrieval operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterL2MessagingBlockAnchored(opts *bind.FilterOpts, l2Block []*big.Int) (*LineaSepoliaVerifyL2MessagingBlockAnchoredIterator, error) {

	var l2BlockRule []interface{}
	for _, l2BlockItem := range l2Block {
		l2BlockRule = append(l2BlockRule, l2BlockItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "L2MessagingBlockAnchored", l2BlockRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyL2MessagingBlockAnchoredIterator{contract: _LineaSepoliaVerify.contract, event: "L2MessagingBlockAnchored", logs: logs, sub: sub}, nil
}

// WatchL2MessagingBlockAnchored is a free log subscription operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchL2MessagingBlockAnchored(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyL2MessagingBlockAnchored, l2Block []*big.Int) (event.Subscription, error) {

	var l2BlockRule []interface{}
	for _, l2BlockItem := range l2Block {
		l2BlockRule = append(l2BlockRule, l2BlockItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "L2MessagingBlockAnchored", l2BlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyL2MessagingBlockAnchored)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "L2MessagingBlockAnchored", log); err != nil {
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

// ParseL2MessagingBlockAnchored is a log parse operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseL2MessagingBlockAnchored(log types.Log) (*LineaSepoliaVerifyL2MessagingBlockAnchored, error) {
	event := new(LineaSepoliaVerifyL2MessagingBlockAnchored)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "L2MessagingBlockAnchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyLimitAmountChangedIterator is returned from FilterLimitAmountChanged and is used to iterate over the raw logs and unpacked data for LimitAmountChanged events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyLimitAmountChangedIterator struct {
	Event *LineaSepoliaVerifyLimitAmountChanged // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyLimitAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyLimitAmountChanged)
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
		it.Event = new(LineaSepoliaVerifyLimitAmountChanged)
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
func (it *LineaSepoliaVerifyLimitAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyLimitAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyLimitAmountChanged represents a LimitAmountChanged event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyLimitAmountChanged struct {
	AmountChangeBy           common.Address
	Amount                   *big.Int
	AmountUsedLoweredToLimit bool
	UsedAmountResetToZero    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLimitAmountChanged is a free log retrieval operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterLimitAmountChanged(opts *bind.FilterOpts, amountChangeBy []common.Address) (*LineaSepoliaVerifyLimitAmountChangedIterator, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyLimitAmountChangedIterator{contract: _LineaSepoliaVerify.contract, event: "LimitAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLimitAmountChanged is a free log subscription operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchLimitAmountChanged(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyLimitAmountChanged, amountChangeBy []common.Address) (event.Subscription, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyLimitAmountChanged)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
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

// ParseLimitAmountChanged is a log parse operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseLimitAmountChanged(log types.Log) (*LineaSepoliaVerifyLimitAmountChanged, error) {
	event := new(LineaSepoliaVerifyLimitAmountChanged)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyLineaRollupVersionChangedIterator is returned from FilterLineaRollupVersionChanged and is used to iterate over the raw logs and unpacked data for LineaRollupVersionChanged events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyLineaRollupVersionChangedIterator struct {
	Event *LineaSepoliaVerifyLineaRollupVersionChanged // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyLineaRollupVersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyLineaRollupVersionChanged)
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
		it.Event = new(LineaSepoliaVerifyLineaRollupVersionChanged)
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
func (it *LineaSepoliaVerifyLineaRollupVersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyLineaRollupVersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyLineaRollupVersionChanged represents a LineaRollupVersionChanged event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyLineaRollupVersionChanged struct {
	PreviousVersion [8]byte
	NewVersion      [8]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLineaRollupVersionChanged is a free log retrieval operation binding the contract event 0x2f8492a7a430cf917798dfb60bc5af634f68e6c40287947df0ea6f7ec0669bd8.
//
// Solidity: event LineaRollupVersionChanged(bytes8 indexed previousVersion, bytes8 indexed newVersion)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterLineaRollupVersionChanged(opts *bind.FilterOpts, previousVersion [][8]byte, newVersion [][8]byte) (*LineaSepoliaVerifyLineaRollupVersionChangedIterator, error) {

	var previousVersionRule []interface{}
	for _, previousVersionItem := range previousVersion {
		previousVersionRule = append(previousVersionRule, previousVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "LineaRollupVersionChanged", previousVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyLineaRollupVersionChangedIterator{contract: _LineaSepoliaVerify.contract, event: "LineaRollupVersionChanged", logs: logs, sub: sub}, nil
}

// WatchLineaRollupVersionChanged is a free log subscription operation binding the contract event 0x2f8492a7a430cf917798dfb60bc5af634f68e6c40287947df0ea6f7ec0669bd8.
//
// Solidity: event LineaRollupVersionChanged(bytes8 indexed previousVersion, bytes8 indexed newVersion)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchLineaRollupVersionChanged(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyLineaRollupVersionChanged, previousVersion [][8]byte, newVersion [][8]byte) (event.Subscription, error) {

	var previousVersionRule []interface{}
	for _, previousVersionItem := range previousVersion {
		previousVersionRule = append(previousVersionRule, previousVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "LineaRollupVersionChanged", previousVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyLineaRollupVersionChanged)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "LineaRollupVersionChanged", log); err != nil {
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

// ParseLineaRollupVersionChanged is a log parse operation binding the contract event 0x2f8492a7a430cf917798dfb60bc5af634f68e6c40287947df0ea6f7ec0669bd8.
//
// Solidity: event LineaRollupVersionChanged(bytes8 indexed previousVersion, bytes8 indexed newVersion)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseLineaRollupVersionChanged(log types.Log) (*LineaSepoliaVerifyLineaRollupVersionChanged, error) {
	event := new(LineaSepoliaVerifyLineaRollupVersionChanged)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "LineaRollupVersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyMessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyMessageClaimedIterator struct {
	Event *LineaSepoliaVerifyMessageClaimed // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyMessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyMessageClaimed)
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
		it.Event = new(LineaSepoliaVerifyMessageClaimed)
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
func (it *LineaSepoliaVerifyMessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyMessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyMessageClaimed represents a MessageClaimed event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyMessageClaimed struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*LineaSepoliaVerifyMessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyMessageClaimedIterator{contract: _LineaSepoliaVerify.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyMessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyMessageClaimed)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseMessageClaimed(log types.Log) (*LineaSepoliaVerifyMessageClaimed, error) {
	event := new(LineaSepoliaVerifyMessageClaimed)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyMessageSentIterator struct {
	Event *LineaSepoliaVerifyMessageSent // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyMessageSent)
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
		it.Event = new(LineaSepoliaVerifyMessageSent)
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
func (it *LineaSepoliaVerifyMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyMessageSent represents a MessageSent event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyMessageSent struct {
	From        common.Address
	To          common.Address
	Fee         *big.Int
	Value       *big.Int
	Nonce       *big.Int
	Calldata    []byte
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*LineaSepoliaVerifyMessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyMessageSentIterator{contract: _LineaSepoliaVerify.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyMessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyMessageSent)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseMessageSent(log types.Log) (*LineaSepoliaVerifyMessageSent, error) {
	event := new(LineaSepoliaVerifyMessageSent)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyPauseTypeRoleSetIterator is returned from FilterPauseTypeRoleSet and is used to iterate over the raw logs and unpacked data for PauseTypeRoleSet events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyPauseTypeRoleSetIterator struct {
	Event *LineaSepoliaVerifyPauseTypeRoleSet // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyPauseTypeRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyPauseTypeRoleSet)
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
		it.Event = new(LineaSepoliaVerifyPauseTypeRoleSet)
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
func (it *LineaSepoliaVerifyPauseTypeRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyPauseTypeRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyPauseTypeRoleSet represents a PauseTypeRoleSet event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyPauseTypeRoleSet struct {
	PauseType uint8
	Role      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPauseTypeRoleSet is a free log retrieval operation binding the contract event 0x33aa8fd1ce49e1761bc8d27fd53414bfefc45d690feed0ce55019d7d3aec6091.
//
// Solidity: event PauseTypeRoleSet(uint8 indexed pauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterPauseTypeRoleSet(opts *bind.FilterOpts, pauseType []uint8, role [][32]byte) (*LineaSepoliaVerifyPauseTypeRoleSetIterator, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "PauseTypeRoleSet", pauseTypeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyPauseTypeRoleSetIterator{contract: _LineaSepoliaVerify.contract, event: "PauseTypeRoleSet", logs: logs, sub: sub}, nil
}

// WatchPauseTypeRoleSet is a free log subscription operation binding the contract event 0x33aa8fd1ce49e1761bc8d27fd53414bfefc45d690feed0ce55019d7d3aec6091.
//
// Solidity: event PauseTypeRoleSet(uint8 indexed pauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchPauseTypeRoleSet(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyPauseTypeRoleSet, pauseType []uint8, role [][32]byte) (event.Subscription, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "PauseTypeRoleSet", pauseTypeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyPauseTypeRoleSet)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "PauseTypeRoleSet", log); err != nil {
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

// ParsePauseTypeRoleSet is a log parse operation binding the contract event 0x33aa8fd1ce49e1761bc8d27fd53414bfefc45d690feed0ce55019d7d3aec6091.
//
// Solidity: event PauseTypeRoleSet(uint8 indexed pauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParsePauseTypeRoleSet(log types.Log) (*LineaSepoliaVerifyPauseTypeRoleSet, error) {
	event := new(LineaSepoliaVerifyPauseTypeRoleSet)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "PauseTypeRoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyPausedIterator struct {
	Event *LineaSepoliaVerifyPaused // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyPaused)
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
		it.Event = new(LineaSepoliaVerifyPaused)
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
func (it *LineaSepoliaVerifyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyPaused represents a Paused event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyPaused struct {
	MessageSender common.Address
	PauseType     uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x534f879afd40abb4e39f8e1b77a316be4c8e3521d9cf5a3a3db8959d574d4559.
//
// Solidity: event Paused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterPaused(opts *bind.FilterOpts, pauseType []uint8) (*LineaSepoliaVerifyPausedIterator, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "Paused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyPausedIterator{contract: _LineaSepoliaVerify.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x534f879afd40abb4e39f8e1b77a316be4c8e3521d9cf5a3a3db8959d574d4559.
//
// Solidity: event Paused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyPaused, pauseType []uint8) (event.Subscription, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "Paused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyPaused)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x534f879afd40abb4e39f8e1b77a316be4c8e3521d9cf5a3a3db8959d574d4559.
//
// Solidity: event Paused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParsePaused(log types.Log) (*LineaSepoliaVerifyPaused, error) {
	event := new(LineaSepoliaVerifyPaused)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyRateLimitInitializedIterator is returned from FilterRateLimitInitialized and is used to iterate over the raw logs and unpacked data for RateLimitInitialized events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRateLimitInitializedIterator struct {
	Event *LineaSepoliaVerifyRateLimitInitialized // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyRateLimitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyRateLimitInitialized)
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
		it.Event = new(LineaSepoliaVerifyRateLimitInitialized)
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
func (it *LineaSepoliaVerifyRateLimitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyRateLimitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyRateLimitInitialized represents a RateLimitInitialized event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRateLimitInitialized struct {
	PeriodInSeconds  *big.Int
	LimitInWei       *big.Int
	CurrentPeriodEnd *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRateLimitInitialized is a free log retrieval operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterRateLimitInitialized(opts *bind.FilterOpts) (*LineaSepoliaVerifyRateLimitInitializedIterator, error) {

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "RateLimitInitialized")
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyRateLimitInitializedIterator{contract: _LineaSepoliaVerify.contract, event: "RateLimitInitialized", logs: logs, sub: sub}, nil
}

// WatchRateLimitInitialized is a free log subscription operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchRateLimitInitialized(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyRateLimitInitialized) (event.Subscription, error) {

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "RateLimitInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyRateLimitInitialized)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RateLimitInitialized", log); err != nil {
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

// ParseRateLimitInitialized is a log parse operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseRateLimitInitialized(log types.Log) (*LineaSepoliaVerifyRateLimitInitialized, error) {
	event := new(LineaSepoliaVerifyRateLimitInitialized)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RateLimitInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleAdminChangedIterator struct {
	Event *LineaSepoliaVerifyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyRoleAdminChanged)
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
		it.Event = new(LineaSepoliaVerifyRoleAdminChanged)
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
func (it *LineaSepoliaVerifyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyRoleAdminChanged represents a RoleAdminChanged event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LineaSepoliaVerifyRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyRoleAdminChangedIterator{contract: _LineaSepoliaVerify.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyRoleAdminChanged)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseRoleAdminChanged(log types.Log) (*LineaSepoliaVerifyRoleAdminChanged, error) {
	event := new(LineaSepoliaVerifyRoleAdminChanged)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleGrantedIterator struct {
	Event *LineaSepoliaVerifyRoleGranted // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyRoleGranted)
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
		it.Event = new(LineaSepoliaVerifyRoleGranted)
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
func (it *LineaSepoliaVerifyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyRoleGranted represents a RoleGranted event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LineaSepoliaVerifyRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyRoleGrantedIterator{contract: _LineaSepoliaVerify.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyRoleGranted)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseRoleGranted(log types.Log) (*LineaSepoliaVerifyRoleGranted, error) {
	event := new(LineaSepoliaVerifyRoleGranted)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleRevokedIterator struct {
	Event *LineaSepoliaVerifyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyRoleRevoked)
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
		it.Event = new(LineaSepoliaVerifyRoleRevoked)
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
func (it *LineaSepoliaVerifyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyRoleRevoked represents a RoleRevoked event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LineaSepoliaVerifyRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyRoleRevokedIterator{contract: _LineaSepoliaVerify.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyRoleRevoked)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseRoleRevoked(log types.Log) (*LineaSepoliaVerifyRoleRevoked, error) {
	event := new(LineaSepoliaVerifyRoleRevoked)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyRollingHashUpdatedIterator is returned from FilterRollingHashUpdated and is used to iterate over the raw logs and unpacked data for RollingHashUpdated events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRollingHashUpdatedIterator struct {
	Event *LineaSepoliaVerifyRollingHashUpdated // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyRollingHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyRollingHashUpdated)
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
		it.Event = new(LineaSepoliaVerifyRollingHashUpdated)
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
func (it *LineaSepoliaVerifyRollingHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyRollingHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyRollingHashUpdated represents a RollingHashUpdated event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyRollingHashUpdated struct {
	MessageNumber *big.Int
	RollingHash   [32]byte
	MessageHash   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollingHashUpdated is a free log retrieval operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterRollingHashUpdated(opts *bind.FilterOpts, messageNumber []*big.Int, rollingHash [][32]byte, messageHash [][32]byte) (*LineaSepoliaVerifyRollingHashUpdatedIterator, error) {

	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var rollingHashRule []interface{}
	for _, rollingHashItem := range rollingHash {
		rollingHashRule = append(rollingHashRule, rollingHashItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "RollingHashUpdated", messageNumberRule, rollingHashRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyRollingHashUpdatedIterator{contract: _LineaSepoliaVerify.contract, event: "RollingHashUpdated", logs: logs, sub: sub}, nil
}

// WatchRollingHashUpdated is a free log subscription operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchRollingHashUpdated(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyRollingHashUpdated, messageNumber []*big.Int, rollingHash [][32]byte, messageHash [][32]byte) (event.Subscription, error) {

	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var rollingHashRule []interface{}
	for _, rollingHashItem := range rollingHash {
		rollingHashRule = append(rollingHashRule, rollingHashItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "RollingHashUpdated", messageNumberRule, rollingHashRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyRollingHashUpdated)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RollingHashUpdated", log); err != nil {
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

// ParseRollingHashUpdated is a log parse operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseRollingHashUpdated(log types.Log) (*LineaSepoliaVerifyRollingHashUpdated, error) {
	event := new(LineaSepoliaVerifyRollingHashUpdated)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "RollingHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyUnPauseTypeRoleSetIterator is returned from FilterUnPauseTypeRoleSet and is used to iterate over the raw logs and unpacked data for UnPauseTypeRoleSet events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyUnPauseTypeRoleSetIterator struct {
	Event *LineaSepoliaVerifyUnPauseTypeRoleSet // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyUnPauseTypeRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyUnPauseTypeRoleSet)
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
		it.Event = new(LineaSepoliaVerifyUnPauseTypeRoleSet)
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
func (it *LineaSepoliaVerifyUnPauseTypeRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyUnPauseTypeRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyUnPauseTypeRoleSet represents a UnPauseTypeRoleSet event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyUnPauseTypeRoleSet struct {
	UnPauseType uint8
	Role        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnPauseTypeRoleSet is a free log retrieval operation binding the contract event 0xe7bf4b8dc0c17a52dc9e52323a3ab61cb2079db35f969125b1f8a3d984c6f6c2.
//
// Solidity: event UnPauseTypeRoleSet(uint8 indexed unPauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterUnPauseTypeRoleSet(opts *bind.FilterOpts, unPauseType []uint8, role [][32]byte) (*LineaSepoliaVerifyUnPauseTypeRoleSetIterator, error) {

	var unPauseTypeRule []interface{}
	for _, unPauseTypeItem := range unPauseType {
		unPauseTypeRule = append(unPauseTypeRule, unPauseTypeItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "UnPauseTypeRoleSet", unPauseTypeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyUnPauseTypeRoleSetIterator{contract: _LineaSepoliaVerify.contract, event: "UnPauseTypeRoleSet", logs: logs, sub: sub}, nil
}

// WatchUnPauseTypeRoleSet is a free log subscription operation binding the contract event 0xe7bf4b8dc0c17a52dc9e52323a3ab61cb2079db35f969125b1f8a3d984c6f6c2.
//
// Solidity: event UnPauseTypeRoleSet(uint8 indexed unPauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchUnPauseTypeRoleSet(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyUnPauseTypeRoleSet, unPauseType []uint8, role [][32]byte) (event.Subscription, error) {

	var unPauseTypeRule []interface{}
	for _, unPauseTypeItem := range unPauseType {
		unPauseTypeRule = append(unPauseTypeRule, unPauseTypeItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "UnPauseTypeRoleSet", unPauseTypeRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyUnPauseTypeRoleSet)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "UnPauseTypeRoleSet", log); err != nil {
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

// ParseUnPauseTypeRoleSet is a log parse operation binding the contract event 0xe7bf4b8dc0c17a52dc9e52323a3ab61cb2079db35f969125b1f8a3d984c6f6c2.
//
// Solidity: event UnPauseTypeRoleSet(uint8 indexed unPauseType, bytes32 indexed role)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseUnPauseTypeRoleSet(log types.Log) (*LineaSepoliaVerifyUnPauseTypeRoleSet, error) {
	event := new(LineaSepoliaVerifyUnPauseTypeRoleSet)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "UnPauseTypeRoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyUnPausedIterator struct {
	Event *LineaSepoliaVerifyUnPaused // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyUnPaused)
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
		it.Event = new(LineaSepoliaVerifyUnPaused)
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
func (it *LineaSepoliaVerifyUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyUnPaused represents a UnPaused event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyUnPaused struct {
	MessageSender common.Address
	PauseType     uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xd071d2b85dec4489435b541d2f0e2570db09b09db9efd8703948d44a433df65a.
//
// Solidity: event UnPaused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterUnPaused(opts *bind.FilterOpts, pauseType []uint8) (*LineaSepoliaVerifyUnPausedIterator, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "UnPaused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyUnPausedIterator{contract: _LineaSepoliaVerify.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xd071d2b85dec4489435b541d2f0e2570db09b09db9efd8703948d44a433df65a.
//
// Solidity: event UnPaused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyUnPaused, pauseType []uint8) (event.Subscription, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "UnPaused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyUnPaused)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xd071d2b85dec4489435b541d2f0e2570db09b09db9efd8703948d44a433df65a.
//
// Solidity: event UnPaused(address messageSender, uint8 indexed pauseType)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseUnPaused(log types.Log) (*LineaSepoliaVerifyUnPaused, error) {
	event := new(LineaSepoliaVerifyUnPaused)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LineaSepoliaVerifyVerifierAddressChangedIterator is returned from FilterVerifierAddressChanged and is used to iterate over the raw logs and unpacked data for VerifierAddressChanged events raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyVerifierAddressChangedIterator struct {
	Event *LineaSepoliaVerifyVerifierAddressChanged // Event containing the contract specifics and raw log

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
func (it *LineaSepoliaVerifyVerifierAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LineaSepoliaVerifyVerifierAddressChanged)
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
		it.Event = new(LineaSepoliaVerifyVerifierAddressChanged)
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
func (it *LineaSepoliaVerifyVerifierAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LineaSepoliaVerifyVerifierAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LineaSepoliaVerifyVerifierAddressChanged represents a VerifierAddressChanged event raised by the LineaSepoliaVerify contract.
type LineaSepoliaVerifyVerifierAddressChanged struct {
	VerifierAddress    common.Address
	ProofType          *big.Int
	VerifierSetBy      common.Address
	OldVerifierAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerifierAddressChanged is a free log retrieval operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) FilterVerifierAddressChanged(opts *bind.FilterOpts, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (*LineaSepoliaVerifyVerifierAddressChangedIterator, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.FilterLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return &LineaSepoliaVerifyVerifierAddressChangedIterator{contract: _LineaSepoliaVerify.contract, event: "VerifierAddressChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierAddressChanged is a free log subscription operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) WatchVerifierAddressChanged(opts *bind.WatchOpts, sink chan<- *LineaSepoliaVerifyVerifierAddressChanged, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (event.Subscription, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _LineaSepoliaVerify.contract.WatchLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LineaSepoliaVerifyVerifierAddressChanged)
				if err := _LineaSepoliaVerify.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
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

// ParseVerifierAddressChanged is a log parse operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_LineaSepoliaVerify *LineaSepoliaVerifyFilterer) ParseVerifierAddressChanged(log types.Log) (*LineaSepoliaVerifyVerifierAddressChanged, error) {
	event := new(LineaSepoliaVerifyVerifierAddressChanged)
	if err := _LineaSepoliaVerify.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
