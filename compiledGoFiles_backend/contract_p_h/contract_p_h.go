// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_p_h

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SC_P_HA_1Dispute is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1Dispute struct {
	DisputeID                 *big.Int
	PID                       *big.Int
	HID                       *big.Int
	FinalBillID               *big.Int
	TimestampOfRaisingDispute *big.Int
	TimestampOfSolvingDispute *big.Int
	FinalCostChanged          bool
}

// SC_P_HA_1EstimatedCheckUpCost is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1EstimatedCheckUpCost struct {
	EstimatedCostBillID            *big.Int
	PID                            *big.Int
	HID                            *big.Int
	EstimatedCost                  *big.Int
	TimestampOfEstimate            *big.Int
	TimestampOfLockingByHospital   *big.Int
	TimestampOfLockingByPatient    *big.Int
	TimestampOfCheckUpStart        *big.Int
	TimestampOfUnlockingByPatient  *big.Int
	TimestampOfUnlockingByHospital *big.Int
}

// SC_P_HA_1Excess is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1Excess struct {
	ExcessID                            *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	ExcessAmount                        *big.Int
	TimestampOfRefundingExcessToPatient *big.Int
}

// SC_P_HA_1FileProperties is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1FileProperties struct {
	NoOfInputGates *big.Int
	FileChunkSize  *big.Int
	Depth          *big.Int
	MaxLinesToGate *big.Int
}

// SC_P_HA_1FinalCheckUpCost is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1FinalCheckUpCost struct {
	FinalCostBillID                     *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	FinalCost                           *big.Int
	TimestampOfFinalBilling             *big.Int
	ExcessID                            *big.Int
	TimeStampOfPatientFinalBillConsent  *big.Int
	TimeStampOfFinalConsent             *big.Int
	TimeStampOfComplain                 *big.Int
	TimeStampOfUnlockingByHospitalFinal *big.Int
	TimestampOfUnlockingByPatientFinal  *big.Int
}

// SC_P_HA_1KeyAndExchange is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1KeyAndExchange struct {
	Key                  [32]byte
	KeyHash              [32]byte
	TimeStampOfKeyReveal *big.Int
}

// SC_P_HA_1MultiSigOnMedicalData is an auto generated low-level Go binding around an user-defined struct.
type SC_P_HA_1MultiSigOnMedicalData struct {
	MultiSigID                                           *big.Int
	PatientID                                            *big.Int
	HospitalID                                           *big.Int
	HashOfData                                           [32]byte
	HashOfEncryptedData                                  [32]byte
	HashOfPatientIdDateHashEncFile                       [32]byte
	SignOnFileHashByHospital                             []byte
	SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile []byte
	TimeStampOfSignByHospital                            *big.Int
	TimeStampOfVerificationByPatient                     *big.Int
	Verified                                             bool
	TimeStampOfUnlockingByPatientMulti                   *big.Int
	TimeStampOfUnlockingByHospitalMulti                  *big.Int
}

// SCPHA1ABI is the input ABI used to generate the binding from.
const SCPHA1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr_SC_P_HA_2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finalCost\",\"type\":\"uint256\"}],\"name\":\"DischargeAndGenerateFinalCostBill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MultiSignIdToEncFileProperties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_noOfInputGates\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLinesToGate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MultiSignIdToKeyAndExchange\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfKeyReveal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addr_SC_P_HA_2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"consentFinalBill_Patient\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dispute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"disputeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRaisingDispute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfSolvingDispute\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalCostChanged\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimatedBillIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"estimatedBillIDToFinalBillID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"estimatedCheckUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"estimatedCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfEstimate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfCheckUpStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByHospital\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"estimatedCostBillIDToMultiSigOnMedicalDataID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"excess\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRefundingExcessToPatient\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excessIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalBillIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalBillIDToDisputeID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalCheckUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfFinalBilling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfPatientFinalBillConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfFinalConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfComplain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_final\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient_final\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedCost\",\"type\":\"uint256\"}],\"name\":\"generateEstimatedCostBill\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getDisputeStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"disputeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRaisingDispute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfSolvingDispute\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalCostChanged\",\"type\":\"bool\"}],\"internalType\":\"structSC_P_HA_1.Dispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getEncFileProperties\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_noOfInputGates\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLinesToGate\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.FileProperties\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getEstimatedCheckUpCostStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"estimatedCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfEstimate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfCheckUpStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByHospital\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.EstimatedCheckUpCost\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getExcessStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRefundingExcessToPatient\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.Excess\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getFinalBillStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"finalCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfFinalBilling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfPatientFinalBillConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfFinalConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfComplain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_final\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient_final\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.FinalCheckUpCost\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getKeyAndExchangeStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfKeyReveal\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.KeyAndExchange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getMultiSigOnMedicalDataStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"multiSigID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"patientID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hospitalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfEncryptedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfPatientIdDateHashEncFile\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signOnFileHashByHospital\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfSignByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfVerificationByPatient\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByPatient_multi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_multi\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.MultiSigOnMedicalData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseTheTimeBy10seconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRefundingExcessToPatient\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.Excess\",\"name\":\"excess_object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_finalId\",\"type\":\"uint256\"}],\"name\":\"incrementAndSetExcessObject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hashOfData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hashOfEncryptedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hashOfPatientIdDateHashEncFile\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signOnFileHashByHospital\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_noOfInputGates\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fileChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxInputLinesToAGate\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"keepSignedHashToBlockchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pId\",\"type\":\"uint256\"}],\"name\":\"keyReveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"lockEstimatedAmount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiSigIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiSigOnMedicalData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"multiSigID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"patientID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hospitalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfEncryptedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfPatientIdDateHashEncFile\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signOnFileHashByHospital\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfSignByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfVerificationByPatient\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByPatient_multi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_multi\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"patientId_hospitalId_estimateBillId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"bytes32[][]\",\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedVectors_ThisGate\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gateIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"patient_Complain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"patient_final_consent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"raiseDispute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"disputeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRaisingDispute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfSolvingDispute\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalCostChanged\",\"type\":\"bool\"}],\"internalType\":\"structSC_P_HA_1.Dispute\",\"name\":\"dispute_object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setDisputeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"estimatedCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfEstimate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfLockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfCheckUpStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByHospital\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.EstimatedCheckUpCost\",\"name\":\"estimatedCheckUpCost_Object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setEstimatedCheckUpCostStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfRefundingExcessToPatient\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.Excess\",\"name\":\"excess_object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setExcessObject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"finalCostBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfFinalBilling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfPatientFinalBillConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfFinalConsent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfComplain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_final\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampOfUnlockingByPatient_final\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.FinalCheckUpCost\",\"name\":\"finalCheckUpCost_Object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setFinalCheckUpCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"multiSigID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"patientID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hospitalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfEncryptedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfPatientIdDateHashEncFile\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signOnFileHashByHospital\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfSignByHospital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfVerificationByPatient\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByPatient_multi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStampOfUnlockingByHospital_multi\",\"type\":\"uint256\"}],\"internalType\":\"structSC_P_HA_1.MultiSigOnMedicalData\",\"name\":\"multiSigOnMedicalData_Object\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setMultiSigOnMedicalData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedCostBillID\",\"type\":\"uint256\"}],\"name\":\"startTreatment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForAcceptingOrProducingComplain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForGeneratingFillBill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForKeyRevealing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForPatientLocking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForRaisingDispute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForSigningFinalBill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForSolvingDispute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForStartingOperation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForVerifyingSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferEtherFromThisContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_multiSigID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_HashOf_PatientId_DateOfReport_HashOfEncryptedFile\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_hashOfEncryptedData\",\"type\":\"bytes32\"}],\"name\":\"verifyAndGiveConsent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SCPHA1FuncSigs maps the 4-byte function signature to its string representation.
var SCPHA1FuncSigs = map[string]string{
	"9e331c86": "DischargeAndGenerateFinalCostBill(address,uint256,uint256,uint256)",
	"a6237310": "MultiSignIdToEncFileProperties(uint256)",
	"544193e3": "MultiSignIdToKeyAndExchange(uint256)",
	"44d4fd19": "Now()",
	"82dab224": "addr_SC_P_HA_2()",
	"5d617c6e": "consentFinalBill_Patient(address,uint256,uint256)",
	"86d6282c": "dispute(uint256)",
	"427d517e": "disputeIDGenerator()",
	"d215ea8b": "estimatedBillIDGenerator()",
	"8e2d5531": "estimatedBillIDToFinalBillID(uint256)",
	"59e04b3a": "estimatedCheckUpCost(uint256)",
	"b90d4235": "estimatedCostBillIDToMultiSigOnMedicalDataID(uint256)",
	"8d19b2de": "excess(uint256)",
	"40841aa5": "excessIDGenerator()",
	"b642a283": "finalBillIDGenerator()",
	"0d84410d": "finalBillIDToDisputeID(uint256)",
	"0c89bf44": "finalCheckUpCost(uint256)",
	"fceaffda": "generateEstimatedCostBill(address,uint256,uint256)",
	"01a2a657": "getDisputeStruct(uint256)",
	"1082b93b": "getEncFileProperties(uint256)",
	"3544ff50": "getEstimatedCheckUpCostStruct(uint256)",
	"6b83401d": "getExcessStruct(uint256)",
	"0f6df082": "getFinalBillStruct(uint256)",
	"4f5a4e1c": "getKeyAndExchangeStruct(uint256)",
	"8275370a": "getMultiSigOnMedicalDataStruct(uint256)",
	"118bedad": "increaseTheTimeBy10seconds()",
	"0dcfa2b3": "incrementAndSetExcessObject((uint256,uint256,uint256,uint256,uint256),uint256)",
	"f5fc313e": "keepSignedHashToBlockchain(address,uint256,uint256,bytes32,bytes32,bytes32,bytes,bytes,uint256,uint256,uint256,bytes32)",
	"46770a9b": "keyReveal(address,bytes32,uint256,uint256)",
	"a212af75": "lockEstimatedAmount(address,uint256,uint256)",
	"ec948a38": "multiSigIDGenerator()",
	"c50eaa14": "multiSigOnMedicalData(uint256)",
	"e77cc243": "patientId_hospitalId_estimateBillId(uint256,uint256)",
	"29954842": "patient_Complain(address,bytes32[][],bytes[],uint256[],uint256,uint256)",
	"fbdfac1c": "patient_final_consent(address,uint256,uint256)",
	"e0230f03": "raiseDispute(address,uint256,uint256)",
	"0cde8571": "setDisputeStruct((uint256,uint256,uint256,uint256,uint256,uint256,bool),uint256)",
	"c6ede1c2": "setEstimatedCheckUpCostStruct((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256)",
	"4fbca6f2": "setExcessObject((uint256,uint256,uint256,uint256,uint256),uint256)",
	"066997c3": "setFinalCheckUpCost((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256)",
	"d50f1fb6": "setMultiSigOnMedicalData((uint256,uint256,uint256,bytes32,bytes32,bytes32,bytes,bytes,uint256,uint256,bool,uint256,uint256),uint256)",
	"8cef6685": "startTreatment(address,uint256,uint256)",
	"8513f5e5": "timeLimitForAcceptingOrProducingComplain()",
	"672e7b21": "timeLimitForGeneratingFillBill()",
	"361ff9b3": "timeLimitForKeyRevealing()",
	"25e3bb73": "timeLimitForPatientLocking()",
	"b17c31fb": "timeLimitForRaisingDispute()",
	"fdcd9147": "timeLimitForSigningFinalBill()",
	"77e711ba": "timeLimitForSolvingDispute()",
	"ef1b3f78": "timeLimitForStartingOperation()",
	"ba7e5638": "timeLimitForVerifyingSignature()",
	"2c7ad37c": "transferEtherFromThisContract(address,uint256)",
	"bb6bbf35": "verifyAndGiveConsent(address,uint256,uint256,bytes32,bytes32)",
}

// SCPHA1Bin is the compiled bytecode used for deploying new contracts.
var SCPHA1Bin = "0x608060405260006003556000600555600060075560006009556000600b553480156200002a57600080fd5b5060405162005ca538038062005ca58339810160408190526200004d9162000073565b600180546001600160a01b0319166001600160a01b0392909216919091179055620000a3565b60006020828403121562000085578081fd5b81516001600160a01b03811681146200009c578182fd5b9392505050565b615bf280620000b36000396000f3fe60806040526004361061031a5760003560e01c80638275370a116101ab578063ba7e5638116100f7578063e77cc24311610095578063f5fc313e1161006f578063f5fc313e146108eb578063fbdfac1c1461090b578063fceaffda1461092b578063fdcd91471461048a5761031a565b8063e77cc243146108b6578063ec948a38146108d6578063ef1b3f781461048a5761031a565b8063c6ede1c2116100d1578063c6ede1c214610841578063d215ea8b14610861578063d50f1fb614610876578063e0230f03146108965761031a565b8063ba7e56381461048a578063bb6bbf35146107e8578063c50eaa14146108085761031a565b80638e2d553111610164578063a62373101161013e578063a623731014610783578063b17c31fb1461048a578063b642a283146107b3578063b90d4235146107c85761031a565b80638e2d55311461073d5780639e331c861461075d578063a212af75146107705761031a565b80638275370a1461066a57806382dab224146106975780638513f5e51461048a57806386d6282c146106b95780638cef6685146106ec5780638d19b2de1461070c5761031a565b8063361ff9b31161026a5780634fbca6f2116102235780635d617c6e116101fd5780635d617c6e1461062a578063672e7b211461048a5780636b83401d1461063d57806377e711ba1461048a5761031a565b80634fbca6f2146105a5578063544193e3146105c557806359e04b3a146105f45761031a565b8063361ff9b31461048a57806340841aa514610519578063427d517e1461052e57806344d4fd191461054357806346770a9b146105585780634f5a4e1c146105785761031a565b80630f6df082116102d757806325e3bb73116102b157806325e3bb731461048a578063299548421461049f5780632c7ad37c146104cc5780633544ff50146104ec5761031a565b80630f6df0821461041b5780631082b93b14610448578063118bedad146104755761031a565b806301a2a6571461031f578063066997c3146103555780630c89bf44146103775780630cde8571146103ae5780630d84410d146103ce5780630dcfa2b3146103fb575b600080fd5b34801561032b57600080fd5b5061033f61033a36600461561c565b61093e565b60405161034c91906156f2565b60405180910390f35b34801561036157600080fd5b50610375610370366004615463565b6109c2565b005b34801561038357600080fd5b5061039761039236600461561c565b610a73565b60405161034c9b9a99989796959493929190615abf565b3480156103ba57600080fd5b506103756103c93660046152e6565b610ada565b3480156103da57600080fd5b506103ee6103e936600461561c565b610b6b565b60405161034c9190615980565b34801561040757600080fd5b506103756104163660046153ff565b610b7d565b34801561042757600080fd5b5061043b61043636600461561c565b610cb1565b60405161034c9190615818565b34801561045457600080fd5b5061046861046336600461561c565b610d58565b60405161034c91906157ed565b34801561048157600080fd5b50610375610da5565b34801561049657600080fd5b506103ee610dab565b3480156104ab57600080fd5b506104bf6104ba366004614fdc565b610db0565b60405161034c91906156b3565b3480156104d857600080fd5b506103756104e7366004614eba565b611436565b3480156104f857600080fd5b5061050c61050736600461561c565b61148b565b60405161034c9190615742565b34801561052557600080fd5b506103ee611528565b34801561053a57600080fd5b506103ee61152e565b34801561054f57600080fd5b506103ee611534565b34801561056457600080fd5b50610375610573366004615160565b61153a565b34801561058457600080fd5b5061059861059336600461561c565b611829565b60405161034c9190615895565b3480156105b157600080fd5b506103756105c03660046153ff565b61186a565b3480156105d157600080fd5b506105e56105e036600461561c565b6118dc565b60405161034c939291906156be565b34801561060057600080fd5b5061061461060f36600461561c565b6118fd565b60405161034c9a99989796959493929190615a7a565b61037561063836600461519a565b61195e565b34801561064957600080fd5b5061065d61065836600461561c565b611cb3565b60405161034c91906157b3565b34801561067657600080fd5b5061068a61068536600461561c565b611d1d565b60405161034c91906158b6565b3480156106a357600080fd5b506106ac611eff565b60405161034c919061569f565b3480156106c557600080fd5b506106d96106d436600461561c565b611f0e565b60405161034c9796959493929190615a48565b3480156106f857600080fd5b5061037561070736600461519a565b611f5d565b34801561071857600080fd5b5061072c61072736600461561c565b6120f0565b60405161034c959493929190615a25565b34801561074957600080fd5b506103ee61075836600461561c565b61212e565b6103ee61076b366004615160565b612140565b61037561077e36600461519a565b6128eb565b34801561078f57600080fd5b506107a361079e36600461561c565b612ade565b60405161034c9493929190615a0a565b3480156107bf57600080fd5b506103ee612b05565b3480156107d457600080fd5b506103ee6107e336600461561c565b612b0b565b3480156107f457600080fd5b506103756108033660046151ce565b612b1d565b34801561081457600080fd5b5061082861082336600461561c565b612f04565b60405161034c9d9c9b9a99989796959493929190615989565b34801561084d57600080fd5b5061037561085c366004615368565b61309d565b34801561086d57600080fd5b506103ee613143565b34801561088257600080fd5b50610375610891366004615505565b613149565b3480156108a257600080fd5b506103ee6108b136600461519a565b613242565b3480156108c257600080fd5b506103ee6108d136600461564c565b6135cd565b3480156108e257600080fd5b506103ee6135ea565b3480156108f757600080fd5b50610375610906366004615211565b6135f0565b34801561091757600080fd5b5061037561092636600461519a565b613780565b61037561093936600461519a565b613932565b610946614acd565b600a600183038154811061095657fe5b60009182526020918290206040805160e0810182526007909302909101805483526001810154938301939093526002830154908201526003820154606082015260048201546080820152600582015460a082015260069091015460ff16151560c082015290505b919050565b60015433906001600160a01b031681146109db57600080fd5b82600660018403815481106109ec57fe5b90600052602060002090600b0201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015561014082015181600a0155905050505050565b60068181548110610a8057fe5b90600052602060002090600b020160009150905080600001549080600101549080600201549080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a015490508b565b60015433906001600160a01b03168114610af357600080fd5b82600a6001840381548110610b0457fe5b600091825260209182902083516007929092020190815590820151600182015560408201516002820155606082015160038201556080820151600482015560a0820151600582015560c0909101516006909101805460ff1916911515919091179055505050565b60106020526000908152604090205481565b60015433906001600160a01b03168114610b9657600080fd5b6009805460010190819055808452600680546000198501908110610bb657fe5b600091825260208083206005600b90930201820193909355600880546001810182559252855191027ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3810191909155908401517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee482015560408401517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee582015560608401517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee68201556080909301517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee7909301929092555050565b610cb9614b0c565b60066001830381548110610cc957fe5b90600052602060002090600b02016040518061016001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a820154815250509050919050565b610d60614b66565b506000908152600e6020908152604091829020825160808101845281548152600182015492810192909252600281015492820192909252600390910154606082015290565b42600055565b606481565b6000828152600d6020908152604080832054600c90925280832054905163fd4cc07960e01b8152606092849283928392908d906001600160a01b0382169063fd4cc07990610e02908c90600401615980565b60006040518083038186803b158015610e1a57600080fd5b505afa158015610e2e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e569190810190614ee5565b5050604051639836b82560e01b81529094506001600160a01b03821690639836b82590610e8790339060040161569f565b60006040518083038186803b158015610e9f57600080fd5b505afa158015610eb3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610edb9190810190614f3d565b5050600280549399508d9450926000198701925082109050610ef957fe5b90600052602060002090600d02016002015414610f1557600080fd5b8460026001840381548110610f2657fe5b90600052602060002090600d02016001015414610f4257600080fd5b60066001840381548110610f5257fe5b90600052602060002090600b020160070154600014610f7057600080fd5b60066001840381548110610f8057fe5b90600052602060002090600b020160080154600014610f9e57600080fd5b6000828152600f6020526040902060020154606442919091031115610fc257600080fd5b4260066001850381548110610fd357fe5b90600052602060002090600b0201600801819055508c518c511461105c57836001600160a01b03166108fc6006600186038154811061100e57fe5b90600052602060002090600b0201600301549081150290604051600060405180830381858888f1935050505015801561104b573d6000803e3d6000fd5b50600097505050505050505061142c565b60005b8c51811015611125578d818151811061107457fe5b602002602001015160008151811061108857fe5b60200260200101516110ac8e838151811061109f57fe5b6020026020010151613c59565b1461111d57846001600160a01b03166108fc600660018703815481106110ce57fe5b90600052602060002090600b0201600301549081150290604051600060405180830381858888f1935050505015801561110b573d6000803e3d6000fd5b5060009850505050505050505061142c565b60010161105f565b506000828152600e60205260409020600301548b516001909101101561116257836001600160a01b03166108fc6006600186038154811061100e57fe5b60005b8b51811015611200576111d78e828151811061117d57fe5b60200260200101518d838151811061119157fe5b6020026020010151600260018703815481106111a957fe5b90600052602060002090600d020160040154600e600088815260200190815260200160002060020154613ceb565b6111f857846001600160a01b03166108fc600660018703815481106110ce57fe5b600101611165565b5060018d510360405190808252806020026020018201604052801561123957816020015b60608152602001906001900390816112245790505b50965060005b60018c51038110156112fa578c816001018151811061125a57fe5b6020026020010151516040519080825280601f01601f19166020018201604052801561128d576020820181803883390190505b5088828151811061129a57fe5b60200260200101819052506112db8d82600101815181106112b757fe5b6020026020010151600f600086815260200190815260200160002060000154613ed4565b8882815181106112e757fe5b602090810291909101015260010161123f565b50611352878d60008151811061130c57fe5b60200260200101516002600186038154811061132457fe5b90600052602060002090600d020160030154600f600087815260200190815260200160002060000154613fb6565b955085156113c557336001600160a01b03166108fc6006600186038154811061137757fe5b90600052602060002090600b0201600301549081150290604051600060405180830381858888f193505050501580156113b4573d6000803e3d6000fd5b50600197505050505050505061142c565b836001600160a01b03166108fc600660018603815481106113e257fe5b90600052602060002090600b0201600301549081150290604051600060405180830381858888f1935050505015801561141f573d6000803e3d6000fd5b5060009750505050505050505b9695505050505050565b60015433906001600160a01b0316811461144f57600080fd5b6040516001600160a01b0384169083156108fc029084906000818181858888f19350505050158015611485573d6000803e3d6000fd5b50505050565b611493614b8e565b600460018303815481106114a357fe5b90600052602060002090600a0201604051806101400160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152602001600882015481526020016009820154815250509050919050565b60095481565b600b5481565b60005481565b6000828152600c6020908152604080832054600d909252808320549051621264bd60e81b815287906001600160a01b03821690631264bd009061158190339060040161569f565b60006040518083038186803b15801561159957600080fd5b505afa1580156115ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115d59190810190614ee5565b909150508094505083600660018503815481106115ee57fe5b90600052602060002090600b0201600201541461160a57600080fd5b846006600185038154811061161b57fe5b90600052602060002090600b0201600101541461163757600080fd5b60646006600185038154811061164957fe5b90600052602060002090600b0201600601544203111561166857600080fd5b6000828152600f60205260409020600201541561168457600080fd5b6040516304a4a6a360e31b81526000906001600160a01b038316906325253518906116b3908990600401615980565b60006040518083038186803b1580156116cb57600080fd5b505afa1580156116df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117079190810190614f3d565b5050604080516020808252818301909252949550600094606094509250905060208201818038833901905050905060005b602081101561177c578a816020811061174d57fe5b1a60f81b82828151811061175d57fe5b60200101906001600160f81b031916908160001a905350600101611738565b5080805190602001209150600f600086815260200190815260200160002060010154821461180457826001600160a01b03166108fc600660018903815481106117c157fe5b90600052602060002090600b0201600301549081150290604051600060405180830381858888f193505050501580156117fe573d6000803e3d6000fd5b5061181c565b6000858152600f602052604090204260028201558a90555b5050505050505050505050565b611831614be1565b506000908152600f6020908152604091829020825160608101845281548152600182015492810192909252600201549181019190915290565b60015433906001600160a01b0316811461188357600080fd5b826008600184038154811061189457fe5b90600052602060002090600502016000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b600f6020526000908152604090208054600182015460029092015490919083565b6004818154811061190a57fe5b90600052602060002090600a0201600091509050806000015490806001015490806002015490806003015490806004015490806005015490806006015490806007015490806008015490806009015490508a565b604051639836b82560e01b815260001983019060009085906001600160a01b03821690639836b8259061199590339060040161569f565b60006040518083038186803b1580156119ad57600080fd5b505afa1580156119c1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119e99190810190614f3d565b505060405163fd4cc07960e01b815292955060009350506001600160a01b0384169163fd4cc0799150611a20908890600401615980565b60006040518083038186803b158015611a3857600080fd5b505afa158015611a4c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a749190810190614ee5565b905050809150508560068581548110611a8957fe5b90600052602060002090600b02016000015414611aa557600080fd5b8260068581548110611ab357fe5b90600052602060002090600b02016001015414611acf57600080fd5b8460068581548110611add57fe5b90600052602060002090600b02016002015414611af957600080fd5b4260068581548110611b0757fe5b90600052602060002090600b0201600401541115611b2457600080fd5b60068481548110611b3157fe5b90600052602060002090600b020160060154600014611b4f57600080fd5b600060068581548110611b5e57fe5b90600052602060002090600b020160050154905080600014611c58578060086001830381548110611b8b57fe5b90600052602060002090600502016000015414611ba757600080fd5b8360086001830381548110611bb857fe5b90600052602060002090600502016001015414611bd457600080fd5b8560086001830381548110611be557fe5b90600052602060002090600502016002015414611c0157600080fd5b600060086001830381548110611c1357fe5b9060005260206000209060050201600301541115611c58574260086001830381548110611c3c57fe5b90600052602060002090600502016004015410611c5857600080fd5b606460068681548110611c6757fe5b90600052602060002090600b02016004015442031115611c8657600080fd5b4260068681548110611c9457fe5b90600052602060002090600b0201600601819055505050505050505050565b611cbb614c01565b60086001830381548110611ccb57fe5b90600052602060002090600502016040518060a0016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815250509050919050565b611d25614c30565b60026001830381548110611d3557fe5b90600052602060002090600d0201604051806101a0016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e245780601f10611df957610100808354040283529160200191611e24565b820191906000526020600020905b815481529060010190602001808311611e0757829003601f168201915b505050918352505060078201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015611eb85780601f10611e8d57610100808354040283529160200191611eb8565b820191906000526020600020905b815481529060010190602001808311611e9b57829003601f168201915b50505091835250506008820154602082015260098201546040820152600a82015460ff1615156060820152600b8201546080820152600c9091015460a09091015292915050565b6001546001600160a01b031681565b600a8181548110611f1b57fe5b60009182526020909120600790910201805460018201546002830154600384015460048501546005860154600690960154949650929491939092919060ff1687565b604051621264bd60e81b815260009084906001600160a01b03821690631264bd0090611f8d90339060040161569f565b60006040518083038186803b158015611fa557600080fd5b505afa158015611fb9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611fe19190810190614ee5565b506004805491945060001986019250859183908110611ffc57fe5b90600052602060002090600a0201600001541461201857600080fd5b826004828154811061202657fe5b90600052602060002090600a0201600201541461204257600080fd5b846004828154811061205057fe5b90600052602060002090600a0201600101541461206c57600080fd5b6004818154811061207957fe5b90600052602060002090600a02016006015442101561209757600080fd5b6064600482815481106120a657fe5b90600052602060002090600a020160060154420311156120c557600080fd5b42600482815481106120d357fe5b90600052602060002090600a020160070181905550505050505050565b600881815481106120fd57fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401549294509092909185565b600c6020526000908152604090205481565b604051621264bd60e81b81526000908190819087906001600160a01b03821690631264bd009061217490339060040161569f565b60006040518083038186803b15801561218c57600080fd5b505afa1580156121a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121c89190810190614ee5565b506040516304a4a6a360e31b8152909450600091506001600160a01b038316906325253518906121fc908a90600401615980565b60006040518083038186803b15801561221457600080fd5b505afa158015612228573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122509190810190614f3d565b50505060008b8152600d6020526040902054600480549197509394506000198c01938c935091508390811061228157fe5b90600052602060002090600a0201600001541461229d57600080fd5b87600482815481106122ab57fe5b90600052602060002090600a020160010154146122c757600080fd5b84600482815481106122d557fe5b90600052602060002090600a020160020154146122f157600080fd5b600481815481106122fe57fe5b90600052602060002090600a020160080154600014801561233d57506004818154811061232757fe5b90600052602060002090600a0201600901546000145b61234657600080fd5b6004818154811061235357fe5b90600052602060002090600a0201600701546000141561237257600080fd5b6004818154811061237f57fe5b90600052602060002090600a02016003015487111561239d57600080fd5b83600260018603815481106123ae57fe5b90600052602060002090600d020160000154146123ca57600080fd5b87600260018603815481106123db57fe5b90600052602060002090600d020160010154146123f757600080fd5b846002600186038154811061240857fe5b90600052602060002090600d0201600201541461242457600080fd5b6002600185038154811061243457fe5b90600052602060002090600d0201600901546000141561245357600080fd5b60646002600186038154811061246557fe5b90600052602060002090600d0201600901544203111561248457600080fd5b6002600185038154811061249457fe5b60009182526020909120600a600d90920201015460ff1615156001146124b957600080fd5b6007805460010190556124ca614b0c565b6007548152602081018990526040810186905260608101889052426080820152600060a082015260048054839081106124ff57fe5b90600052602060002090600a0201600301548810156126eb57612520614c01565b60006004848154811061252f57fe5b90600052602060002090600a0201600301548a10156126e857896004858154811061255657fe5b600091825260208083206003600a909302019190910154600980546001019081905586529085018e905260408086018c9052929003606085018190524260808601529151919250339183156108fc0291849190818181858888f193505050501580156125c6573d6000803e3d6000fd5b506040516001600160a01b0386169082156108fc029083906000818181858888f193505050501580156125fd573d6000803e3d6000fd5b506008805460018101825560009190915282517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360059092029182015560208301517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee482015560408301517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee582015560608301517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee682015560808301517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee79091015560095460a08401525b50505b600060c082018181526006805460018101825590835283517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f600b909202918201556020808501517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d408301556040808601517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4184015560608601517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4284015560808601517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4384015560a08601517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4484015592517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4583015560e08501517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d468301556101008501517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d478301556101208501517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d48830155610140909401517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d49909101556007548c8352600c9093529020819055955050505050505b949350505050565b604051639836b82560e01b815260009084906001600160a01b03821690639836b8259061291c90339060040161569f565b60006040518083038186803b15801561293457600080fd5b505afa158015612948573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526129709190810190614f3d565b505060048054939650600019880194508793909250849150811061299057fe5b90600052602060002090600a020160000154146129ac57600080fd5b82600482815481106129ba57fe5b90600052602060002090600a020160010154146129d657600080fd5b84600482815481106129e457fe5b90600052602060002090600a02016002015414612a0057600080fd5b60048181548110612a0d57fe5b90600052602060002090600a02016005015460001415612a2c57600080fd5b606460048281548110612a3b57fe5b90600052602060002090600a02016005015442031115612a5a57600080fd5b60048181548110612a6757fe5b90600052602060002090600a020160090154600014612a8557600080fd5b3460048281548110612a9357fe5b90600052602060002090600a0201600301541115612ab057600080fd5b4260046001860381548110612ac157fe5b90600052602060002090600a020160060181905550505050505050565b600e6020526000908152604090208054600182015460028301546003909301549192909184565b60075481565b600d6020526000908152604090205481565b604051639836b82560e01b815260009086906001600160a01b03821690639836b82590612b4e90339060040161569f565b60006040518083038186803b158015612b6657600080fd5b505afa158015612b7a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612ba29190810190614f3d565b505060405163fd4cc07960e01b815292955060009350506001600160a01b0384169163fd4cc0799150612bd9908990600401615980565b60006040518083038186803b158015612bf157600080fd5b505afa158015612c05573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612c2d9190810190614ee5565b50506002805491925060001989019189919083908110612c4957fe5b90600052602060002090600d02016000015414612c6557600080fd5b8360028281548110612c7357fe5b90600052602060002090600d02016001015414612c8f57600080fd5b8660028281548110612c9d57fe5b90600052602060002090600d02016002015414612cb957600080fd5b8460028281548110612cc757fe5b90600052602060002090600d02016004015414612ce357600080fd5b612d95828760028481548110612cf557fe5b60009182526020918290206007600d9092020101805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015612d8b5780601f10612d6057610100808354040283529160200191612d8b565b820191906000526020600020905b815481529060010190602001808311612d6e57829003601f168201915b505050505061402e565b1515600114612da357600080fd5b612e3e8260028381548110612db457fe5b90600052602060002090600d02016003015460028481548110612dd357fe5b60009182526020918290206006600d9092020101805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015612d8b5780601f10612d6057610100808354040283529160200191612d8b565b1515600114612e4c57600080fd5b60028181548110612e5957fe5b90600052602060002090600d020160080154421015612e7757600080fd5b606460028281548110612e8657fe5b90600052602060002090600d02016008015442031115612ea557600080fd5b4260028281548110612eb357fe5b90600052602060002090600d020160090181905550600160028281548110612ed757fe5b60009182526020909120600d90910201600a01805460ff1916911515919091179055505050505050505050565b60028181548110612f1157fe5b90600052602060002090600d0201600091509050806000015490806001015490806002015490806003015490806004015490806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612fe15780601f10612fb657610100808354040283529160200191612fe1565b820191906000526020600020905b815481529060010190602001808311612fc457829003601f168201915b5050505060078301805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529495949350908301828280156130715780601f1061304657610100808354040283529160200191613071565b820191906000526020600020905b81548152906001019060200180831161305457829003601f168201915b50505060088401546009850154600a860154600b870154600c909701549596929591945060ff1692508d565b60015433906001600160a01b031681146130b657600080fd5b82600460018403815481106130c757fe5b90600052602060002090600a0201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015561010082015181600801556101208201518160090155905050505050565b60055481565b60015433906001600160a01b0316811461316257600080fd5b826002600184038154811061317357fe5b90600052602060002090600d0201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060190805190602001906131d9929190614ca3565b5060e082015180516131f5916007840191602090910190614ca3565b5061010082015160088201556101208201516009820155610140820151600a8201805460ff1916911515919091179055610160820151600b82015561018090910151600c90910155505050565b604051639836b82560e01b8152600090600019840190829086906001600160a01b03821690639836b8259061327b90339060040161569f565b60006040518083038186803b15801561329357600080fd5b505afa1580156132a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526132cf9190810190614f3d565b505060405163fd4cc07960e01b815292955060009350506001600160a01b0384169163fd4cc0799150613306908990600401615980565b60006040518083038186803b15801561331e57600080fd5b505afa158015613332573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261335a9190810190614ee5565b50506006805491925088918690811061336f57fe5b90600052602060002090600b0201600001541461338b57600080fd5b826006858154811061339957fe5b90600052602060002090600b020160010154146133b557600080fd5b85600685815481106133c357fe5b90600052602060002090600b020160020154146133df57600080fd5b42600685815481106133ed57fe5b90600052602060002090600b0201600401541061340957600080fd5b60646006858154811061341857fe5b90600052602060002090600b0201600401544203111561343757600080fd5b600b80546001019055613448614acd565b50506040805160e081018252600b805482526020808301958652828401898152606084018b81524260808601908152600060a0870181815260c08801828152600a805460018101825590845298517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a86007909a02998a01559a517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a989015593517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2aa88015591517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2ab870155517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2ac86015590517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2ad85015595517fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2ae909301805460ff1916931515939093179092555488855260109091529220829055509150505b9392505050565b601160209081526000928352604080842090915290825290205481565b60035481565b604051621264bd60e81b81526000908d906001600160a01b03821690631264bd009061362090339060040161569f565b60006040518083038186803b15801561363857600080fd5b505afa15801561364c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526136749190810190614ee5565b5092505060018d108015906136f95750806001600160a01b031663c784e1146040518163ffffffff1660e01b815260040160206040518083038186803b1580156136bd57600080fd5b505afa1580156136d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506136f59190810190615634565b8d11155b61370257600080fd5b8b61370c57600080fd5b60008c8152600d60205260409020541561372557600080fd5b6003805460010190819055613741908e848e8e8e8e8e42614101565b61374f600354878787614310565b61375b60035484614371565b505060035460009a8b52600d6020526040909a20999099555050505050505050505050565b6000828152600c6020908152604080832054600d909252808320549051639836b82560e01b815286906001600160a01b03821690639836b825906137c890339060040161569f565b60006040518083038186803b1580156137e057600080fd5b505afa1580156137f4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261381c9190810190614f3d565b50506006805493985089945092600019880192508210905061383a57fe5b90600052602060002090600b0201600201541461385657600080fd5b836006600185038154811061386757fe5b90600052602060002090600b0201600101541461388357600080fd5b6006600184038154811061389357fe5b90600052602060002090600b0201600701546000146138b157600080fd5b600660018403815481106138c157fe5b90600052602060002090600b0201600801546000146138df57600080fd5b6000828152600f602052604090206002015460644291909103111561390357600080fd5b426006600185038154811061391457fe5b90600052602060002090600b02016007018190555050505050505050565b604051621264bd60e81b815260009084906001600160a01b03821690631264bd009061396290339060040161569f565b60006040518083038186803b15801561397a57600080fd5b505afa15801561398e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526139b69190810190614ee5565b5092505060018410801590613a3b5750806001600160a01b031663c784e1146040518163ffffffff1660e01b815260040160206040518083038186803b1580156139ff57600080fd5b505afa158015613a13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613a379190810190615634565b8411155b613a4457600080fd5b82341015613a5157600080fd5b600580546001019055613a62614b8e565b60058054825260208083018781526040808501878152606086019889524260808701818152600060c0890181815260a08a0193845260e08a018281526101008b018381526101208c01848152600480546001810182559086529c517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b600a909e029d8e015598517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c8d015595517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d8c01559c517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19e8b015591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19f8a015591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a0890155517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a188015598517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a2870155517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a386015590517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd1a490940193909355905495855260118152818520938552929092525090205550565b600080606083516040519080825280601f01601f191660200182016040528015613c8a576020820181803883390190505b50905060005b8451811015613cdb57848181518110613ca557fe5b602001015160f81c60f81b828281518110613cbc57fe5b60200101906001600160f81b031916908160001a905350600101613c90565b5080516020909101209392505050565b60008085600081518110613cfb57fe5b60209081029190910181015160408051818152606080820183529294506000938493929082018180388339019050509050600091505b60018603821015613ec657600180831b8916831c1415613e055760005b6020811015613da857898360010181518110613d6657fe5b60200260200101518160208110613d7957fe5b1a60f81b828281518110613d8957fe5b60200101906001600160f81b031916908160001a905350600101613d4e565b5060005b6020811015613df357848160208110613dc157fe5b1a60f81b828260200181518110613dd457fe5b60200101906001600160f81b031916908160001a905350600101613dac565b50805160208201209350839250613ebb565b60005b6020811015613e4c57848160208110613e1d57fe5b1a60f81b828281518110613e2d57fe5b60200101906001600160f81b031916908160001a905350600101613e08565b5060005b6020811015613ead57898360010181518110613e6857fe5b60200260200101518160208110613e7b57fe5b1a60f81b828260200181518110613e8e57fe5b60200101906001600160f81b031916908160001a905350600101613e50565b508051602082012093508392505b600190910190613d31565b505050909214949350505050565b606080606084516040519080825280601f01601f191660200182016040528015613f05576020820181803883390190505b509050613f138486516143a8565b905084516040519080825280601f01601f191660200182016040528015613f41576020820181803883390190505b50915060005b8551811015613faa57858181518110613f5c57fe5b602001015160f81c60f81b828281518110613f7357fe5b602001015160f81c60f81b18838281518110613f8b57fe5b60200101906001600160f81b031916908160001a905350600101613f47565b50909150505b92915050565b60008060006060613fc78887614632565b925086516040519080825280601f01601f191660200182016040528015613ff5576020820181803883390190505b5090506140028786613ed4565b905061400d816148f1565b915081831461402257600193505050506128e3565b600093505050506128e3565b600080600080845160411461404957600093505050506135c6565b50505060208201516040830151606084015160001a601b81101561406b57601b015b8060ff16601b1415801561408357508060ff16601c14155b1561409457600093505050506135c6565b866001600160a01b0316600187838686604051600081526020016040526040516140c194939291906156d4565b6020604051602081039080840390855afa1580156140e3573d6000803e3d6000fd5b505050602060405103516001600160a01b03161493505050506135c6565b614109614c30565b50604080516101a0810182528a815260208082018b81529282018a8152606083018a8152608084018a815260a085018a815260c086018a815260e087018a9052426101008801526000610120880181905261014088018190526101608801819052610180880181905260028054600181018255915287517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace600d90920291820190815598517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf82015594517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad086015592517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad185015590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad2840155517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad38301555180519394859490936142a0937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad401920190614ca3565b5060e082015180516142bc916007840191602090910190614ca3565b5061010082015160088201556101208201516009820155610140820151600a8201805460ff1916911515919091179055610160820151600b82015561018090910151600c9091015550505050505050505050565b614318614b66565b8381526020810183905261432e6002850261490f565b60019081016040808401918252606084019485526000978852600e6020908152972083518155969092015190860155516002850155516003909301929092555050565b614379614be1565b60208082019283526000938452600f905260409283902081518155915160018301559190910151600290910155565b60608060606000846040519080825280601f01601f1916602001820160405280156143da576020820181803883390190505b5060408051602080825281830190925291945060208201818038833901905050915060005b60208604811015614627578061452a57604080516021808252606082810190935260208201818038833901905050905060005b60208110156144765788816020811061444757fe5b1a60f81b82828151811061445757fe5b60200101906001600160f81b031916908160001a905350600101614432565b50600060f81b8160208151811061448957fe5b60200101906001600160f81b031916908160001a90535080516020820120925060005b6020811015614523578381602081106144c157fe5b1a60f81b8582815181106144d157fe5b60200101906001600160f81b031916908160001a9053508381602081106144f457fe5b1a60f81b86828151811061450457fe5b60200101906001600160f81b031916908160001a9053506001016144ac565b505061461f565b604080516020808252818301909252606094509060208201818038833901905050925060005b60208110156145945782816020811061456557fe5b1a60f81b84828151811061457557fe5b60200101906001600160f81b031916908160001a905350600101614550565b5082516020840120915060005b602081101561461d578281602081106145b657fe5b1a60f81b8482815181106145c657fe5b60200101906001600160f81b031916908160001a9053508281602081106145e957fe5b1a60f81b85828460200201815181106145fe57fe5b60200101906001600160f81b031916908160001a9053506001016145a1565b505b6001016143ff565b509195945050505050565b600080606081600260608760008151811061464957fe5b6020026020010151516002026040519080825280601f01601f19166020018201604052801561467f576020820181803883390190505b50905060005b8860008151811061469257fe5b6020026020010151518110156146f857886000815181106146af57fe5b602002602001015181815181106146c257fe5b602001015160f81c60f81b8282815181106146d957fe5b60200101906001600160f81b031916908160001a905350600101614685565b5060005b8860018151811061470957fe5b602002602001015151811015614786578860018151811061472657fe5b6020026020010151818151811061473957fe5b602001015160f81c60f81b828a60008151811061475257fe5b60200260200101515183018151811061476757fe5b60200101906001600160f81b031916908160001a9053506001016146fc565b5080516020909101209450613fb09350505050565b60200101906001600160f81b031916908160001a9053506147bb856148f1565b9650613fb095505050505050565b60208160ff161015614807578160f81b868260ff16815181106147e857fe5b60200101906001600160f81b031916908160001a9053506001016147c9565b50600060f81b8560008151811061479b57fe5b8960008151811061482757fe5b60200260200101515181101561488d578960008151811061484457fe5b6020026020010151818151811061485757fe5b602001015160f81c60f81b83828151811061486e57fe5b60200101906001600160f81b031916908160001a90535060010161481a565b50815160208301209550878614156148b0578060f81b8560008151811061479b57fe5b60005b60208160ff161015614807578160f81b868260ff16815181106148d257fe5b60200101906001600160f81b031916908160001a9053506001016148b3565b805160009082906149065750600090506109bd565b50506020015190565b604080517ff8f9cbfae6cc78fbefe7cdc3a1793dfcf4f0e8bbd8cec470b6a28a7a5a3e1efd81527ff5ecf1b3e9debc68e1d9cfabc5997135bfb7a7a3938b7b606b5b4b3f2f1f0ffe60208201527ff6e4ed9ff2d6b458eadcdf97bd91692de2d4da8fd2d0ac50c6ae9a8272523616818301527fc8c0b887b0a8a4489c948c7f847c6125746c645c544c444038302820181008ff60608201527ff7cae577eec2a03cf3bad76fb589591debb2dd67e0aa9834bea6925f6a4a2e0e60808201527fe39ed557db96902cd38ed14fad815115c786af479b7e8324736353433727170760a08201527fc976c13bb96e881cb166a933a55e490d9d56952b8d4e801485467d236242260660c08201527f753a6d1b65325d0c552a4d1345224105391a310b29122104190a11030902010060e0820152610100808201909252600160f81b6001600160801b68010000000000000000640100000000620100006010600460026000198c019081041790810417908104178881041790810417908104179081041790810417017e818283848586878898a8b8c8d8e8f929395969799a9b9d9e9faaeb6bedeeff0281900460ff039091015104600160ff1b909211020190565b6040518060e001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b60405180610160016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b604051806101400160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516060810182526000808252602082018190529181019190915290565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b604051806101a001604052806000815260200160008152602001600081526020016000801916815260200160008019168152602001600080191681526020016060815260200160608152602001600081526020016000815260200160001515815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614ce457805160ff1916838001178555614d11565b82800160010185558215614d11579182015b82811115614d11578251825591602001919060010190614cf6565b50614d1d929150614d21565b5090565b614d3b91905b80821115614d1d5760008155600101614d27565b90565b8035613fb081615ba4565b600082601f830112614d59578081fd5b8135614d6c614d6782615b34565b615b0d565b818152915060208083019084810160005b84811015614da657614d94888484358a0101614e1f565b84529282019290820190600101614d7d565b505050505092915050565b600082601f830112614dc1578081fd5b8135614dcf614d6782615b34565b818152915060208083019084810181840286018201871015614df057600080fd5b60005b84811015614da657813584529282019290820190600101614df3565b80358015158114613fb057600080fd5b600082601f830112614e2f578081fd5b8135614e3d614d6782615b54565b9150808252836020828501011115614e5457600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112614e7d578081fd5b8151614e8b614d6782615b54565b9150808252836020828501011115614ea257600080fd5b614eb3816020840160208601615b78565b5092915050565b60008060408385031215614ecc578182fd5b8235614ed781615ba4565b946020939093013593505050565b600080600060608486031215614ef9578081fd5b8351614f0481615ba4565b60208501516040860151919450925067ffffffffffffffff811115614f27578182fd5b614f3386828701614e6d565b9150509250925092565b60008060008060008060c08789031215614f55578384fd5b8651614f6081615ba4565b60208801516040890151919750955067ffffffffffffffff80821115614f84578586fd5b614f908a838b01614e6d565b95506060890151915060ff82168214614fa7578384fd5b608089015160a08a0151929550935080821115614fc2578283fd5b50614fcf89828a01614e6d565b9150509295509295509295565b60008060008060008060c08789031215614ff4578384fd5b614ffe8735615ba4565b8635955067ffffffffffffffff6020880135111561501a578384fd5b6020870135870188601f82011261502f578485fd5b61503c614d678235615b34565b81358152602080820191908301875b84358110156150ec578c603f833587010112615065578889fd5b615078614d676020843588010135615b34565b8060208435880101358252602082019150604084358801018f604060208088358c0101350287358b01010111156150ad578b8cfd5b8b5b602086358a0101358110156150d45781358452602093840193909101906001016150af565b5050855250602093840193919091019060010161504b565b505080975050505067ffffffffffffffff6040880135111561510c578384fd5b61511c8860408901358901614d49565b935067ffffffffffffffff60608801351115615136578182fd5b6151468860608901358901614db1565b92506080870135915060a087013590509295509295509295565b60008060008060808587031215615175578182fd5b843561518081615ba4565b966020860135965060408601359560600135945092505050565b6000806000606084860312156151ae578081fd5b83356151b981615ba4565b95602085013595506040909401359392505050565b600080600080600060a086880312156151e5578283fd5b85356151f081615ba4565b97602087013597506040870135966060810135965060800135945092505050565b6000806000806000806000806000806000806101808d8f03121561523357898afd5b61523d8e8e614d3e565b9b5060208d01359a5060408d0135995060608d0135985060808d0135975060a08d0135965067ffffffffffffffff60c08e0135111561527a578586fd5b61528a8e60c08f01358f01614e1f565b955067ffffffffffffffff60e08e013511156152a4578485fd5b6152b48e60e08f01358f01614e1f565b94506101008d013593506101208d013592506101408d013591506101608d013590509295989b509295989b509295989b565b6000808284036101008112156152fa578283fd5b60e0811215615307578283fd5b5061531260e0615b0d565b833581526020840135602082015260408401356040820152606084013560608201526080840135608082015260a084013560a08201526153558560c08601614e0f565b60c08201529460e0939093013593505050565b60008082840361016081121561537c578283fd5b6101408082121561538b578384fd5b61539481615b0d565b853581526020808701359082015260408087013590820152606080870135908201526080808701359082015260a0808701359082015260c0808701359082015260e0808701359082015261010080870135908201526101208087013590820152969401359450505050565b60008082840360c0811215615412578283fd5b60a081121561541f578283fd5b5061542a60a0615b0d565b83358152602080850135908201526040808501359082015260608085013590820152608080850135908201529460a09093013593505050565b600080828403610180811215615477578283fd5b61016080821215615486578384fd5b61548f81615b0d565b853581526020808701359082015260408087013590820152606080870135908201526080808701359082015260a0808701359082015260c0808701359082015260e08087013590820152610100808701359082015261012080870135908201526101408087013590820152969401359450505050565b60008060408385031215615517578182fd5b823567ffffffffffffffff8082111561552e578384fd5b6101a0918501808703831315615542578485fd5b61554b83615b0d565b813581526020820135602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c0820135935082841115615594578586fd5b6155a088858401614e1f565b60c082015260e08201359350828411156155b8578586fd5b6155c488858401614e1f565b60e08201526101008281013590820152610120808301359082015261014093506155f088858401614e0f565b938101939093526101608181013590840152610180908101359083015250946020939093013593505050565b60006020828403121561562d578081fd5b5035919050565b600060208284031215615645578081fd5b5051919050565b6000806040838503121561565e578182fd5b50508035926020909101359150565b15159052565b6000815180845261568b816020860160208601615b78565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b901515815260200190565b9283526020830191909152604082015260600190565b93845260ff9290921660208401526040830152606082015260800190565b600060e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c0830151151560c083015292915050565b600061014082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151818401525061012080840151818401525092915050565b600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b8151815260208083015190820152604080830151908201526060918201519181019190915260800190565b600061016082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010080840151818401525061012080840151818401525061014080840151818401525092915050565b81518152602080830151908201526040918201519181019190915260600190565b60006020825282516020830152602083015160408301526040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101a08060e08501526159116101c0850183615673565b60e08601519250610100601f1986830301818701526159308285615673565b878201516101208881019190915288015161014080890191909152880151945061016092506159618388018661566d565b9187015161018087810191909152909601519190940152509092915050565b90815260200190565b60006101a08f83528e60208401528d60408401528c60608401528b60808401528a60a08401528060c08401526159c18184018b615673565b83810360e08501526159d3818b615673565b6101008501999099525050506101208101949094529115156101408401526101608301526101809091015298975050505050505050565b93845260208401929092526040830152606082015260800190565b948552602085019390935260408401919091526060830152608082015260a00190565b968752602087019590955260408601939093526060850191909152608084015260a0830152151560c082015260e00190565b998a5260208a019890985260408901969096526060880194909452608087019290925260a086015260c085015260e08401526101008301526101208201526101400190565b9a8b5260208b019990995260408a01979097526060890195909552608088019390935260a087019190915260c086015260e08501526101008401526101208301526101408201526101600190565b60405181810167ffffffffffffffff81118282101715615b2c57600080fd5b604052919050565b600067ffffffffffffffff821115615b4a578081fd5b5060209081020190565b600067ffffffffffffffff821115615b6a578081fd5b50601f01601f191660200190565b60005b83811015615b93578181015183820152602001615b7b565b838111156114855750506000910152565b6001600160a01b0381168114615bb957600080fd5b5056fea2646970667358221220517fb28e70ae0f68a2ec3f3124870eb379dc181cc01d20d92404cbdc2d07574a64736f6c63430006020033"

// DeploySCPHA1 deploys a new Ethereum contract, binding an instance of SCPHA1 to it.
func DeploySCPHA1(auth *bind.TransactOpts, backend bind.ContractBackend, _addr_SC_P_HA_2 common.Address) (common.Address, *types.Transaction, *SCPHA1, error) {
	parsed, err := abi.JSON(strings.NewReader(SCPHA1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SCPHA1Bin), backend, _addr_SC_P_HA_2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SCPHA1{SCPHA1Caller: SCPHA1Caller{contract: contract}, SCPHA1Transactor: SCPHA1Transactor{contract: contract}, SCPHA1Filterer: SCPHA1Filterer{contract: contract}}, nil
}

// SCPHA1 is an auto generated Go binding around an Ethereum contract.
type SCPHA1 struct {
	SCPHA1Caller     // Read-only binding to the contract
	SCPHA1Transactor // Write-only binding to the contract
	SCPHA1Filterer   // Log filterer for contract events
}

// SCPHA1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SCPHA1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SCPHA1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCPHA1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCPHA1Session struct {
	Contract     *SCPHA1           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCPHA1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCPHA1CallerSession struct {
	Contract *SCPHA1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SCPHA1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCPHA1TransactorSession struct {
	Contract     *SCPHA1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCPHA1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SCPHA1Raw struct {
	Contract *SCPHA1 // Generic contract binding to access the raw methods on
}

// SCPHA1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCPHA1CallerRaw struct {
	Contract *SCPHA1Caller // Generic read-only contract binding to access the raw methods on
}

// SCPHA1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCPHA1TransactorRaw struct {
	Contract *SCPHA1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSCPHA1 creates a new instance of SCPHA1, bound to a specific deployed contract.
func NewSCPHA1(address common.Address, backend bind.ContractBackend) (*SCPHA1, error) {
	contract, err := bindSCPHA1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCPHA1{SCPHA1Caller: SCPHA1Caller{contract: contract}, SCPHA1Transactor: SCPHA1Transactor{contract: contract}, SCPHA1Filterer: SCPHA1Filterer{contract: contract}}, nil
}

// NewSCPHA1Caller creates a new read-only instance of SCPHA1, bound to a specific deployed contract.
func NewSCPHA1Caller(address common.Address, caller bind.ContractCaller) (*SCPHA1Caller, error) {
	contract, err := bindSCPHA1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCPHA1Caller{contract: contract}, nil
}

// NewSCPHA1Transactor creates a new write-only instance of SCPHA1, bound to a specific deployed contract.
func NewSCPHA1Transactor(address common.Address, transactor bind.ContractTransactor) (*SCPHA1Transactor, error) {
	contract, err := bindSCPHA1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCPHA1Transactor{contract: contract}, nil
}

// NewSCPHA1Filterer creates a new log filterer instance of SCPHA1, bound to a specific deployed contract.
func NewSCPHA1Filterer(address common.Address, filterer bind.ContractFilterer) (*SCPHA1Filterer, error) {
	contract, err := bindSCPHA1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCPHA1Filterer{contract: contract}, nil
}

// bindSCPHA1 binds a generic wrapper to an already deployed contract.
func bindSCPHA1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SCPHA1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCPHA1 *SCPHA1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCPHA1.Contract.SCPHA1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCPHA1 *SCPHA1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCPHA1.Contract.SCPHA1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCPHA1 *SCPHA1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCPHA1.Contract.SCPHA1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCPHA1 *SCPHA1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCPHA1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCPHA1 *SCPHA1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCPHA1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCPHA1 *SCPHA1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCPHA1.Contract.contract.Transact(opts, method, params...)
}

// MultiSignIdToEncFileProperties is a free data retrieval call binding the contract method 0xa6237310.
//
// Solidity: function MultiSignIdToEncFileProperties(uint256 ) constant returns(uint256 _noOfInputGates, uint256 fileChunkSize, uint256 depth, uint256 maxLinesToGate)
func (_SCPHA1 *SCPHA1Caller) MultiSignIdToEncFileProperties(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NoOfInputGates *big.Int
	FileChunkSize  *big.Int
	Depth          *big.Int
	MaxLinesToGate *big.Int
}, error) {
	ret := new(struct {
		NoOfInputGates *big.Int
		FileChunkSize  *big.Int
		Depth          *big.Int
		MaxLinesToGate *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "MultiSignIdToEncFileProperties", arg0)
	return *ret, err
}

// MultiSignIdToEncFileProperties is a free data retrieval call binding the contract method 0xa6237310.
//
// Solidity: function MultiSignIdToEncFileProperties(uint256 ) constant returns(uint256 _noOfInputGates, uint256 fileChunkSize, uint256 depth, uint256 maxLinesToGate)
func (_SCPHA1 *SCPHA1Session) MultiSignIdToEncFileProperties(arg0 *big.Int) (struct {
	NoOfInputGates *big.Int
	FileChunkSize  *big.Int
	Depth          *big.Int
	MaxLinesToGate *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSignIdToEncFileProperties(&_SCPHA1.CallOpts, arg0)
}

// MultiSignIdToEncFileProperties is a free data retrieval call binding the contract method 0xa6237310.
//
// Solidity: function MultiSignIdToEncFileProperties(uint256 ) constant returns(uint256 _noOfInputGates, uint256 fileChunkSize, uint256 depth, uint256 maxLinesToGate)
func (_SCPHA1 *SCPHA1CallerSession) MultiSignIdToEncFileProperties(arg0 *big.Int) (struct {
	NoOfInputGates *big.Int
	FileChunkSize  *big.Int
	Depth          *big.Int
	MaxLinesToGate *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSignIdToEncFileProperties(&_SCPHA1.CallOpts, arg0)
}

// MultiSignIdToKeyAndExchange is a free data retrieval call binding the contract method 0x544193e3.
//
// Solidity: function MultiSignIdToKeyAndExchange(uint256 ) constant returns(bytes32 key, bytes32 keyHash, uint256 timeStampOfKeyReveal)
func (_SCPHA1 *SCPHA1Caller) MultiSignIdToKeyAndExchange(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Key                  [32]byte
	KeyHash              [32]byte
	TimeStampOfKeyReveal *big.Int
}, error) {
	ret := new(struct {
		Key                  [32]byte
		KeyHash              [32]byte
		TimeStampOfKeyReveal *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "MultiSignIdToKeyAndExchange", arg0)
	return *ret, err
}

// MultiSignIdToKeyAndExchange is a free data retrieval call binding the contract method 0x544193e3.
//
// Solidity: function MultiSignIdToKeyAndExchange(uint256 ) constant returns(bytes32 key, bytes32 keyHash, uint256 timeStampOfKeyReveal)
func (_SCPHA1 *SCPHA1Session) MultiSignIdToKeyAndExchange(arg0 *big.Int) (struct {
	Key                  [32]byte
	KeyHash              [32]byte
	TimeStampOfKeyReveal *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSignIdToKeyAndExchange(&_SCPHA1.CallOpts, arg0)
}

// MultiSignIdToKeyAndExchange is a free data retrieval call binding the contract method 0x544193e3.
//
// Solidity: function MultiSignIdToKeyAndExchange(uint256 ) constant returns(bytes32 key, bytes32 keyHash, uint256 timeStampOfKeyReveal)
func (_SCPHA1 *SCPHA1CallerSession) MultiSignIdToKeyAndExchange(arg0 *big.Int) (struct {
	Key                  [32]byte
	KeyHash              [32]byte
	TimeStampOfKeyReveal *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSignIdToKeyAndExchange(&_SCPHA1.CallOpts, arg0)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "Now")
	return *ret0, err
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) Now() (*big.Int, error) {
	return _SCPHA1.Contract.Now(&_SCPHA1.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) Now() (*big.Int, error) {
	return _SCPHA1.Contract.Now(&_SCPHA1.CallOpts)
}

// AddrSCPHA2 is a free data retrieval call binding the contract method 0x82dab224.
//
// Solidity: function addr_SC_P_HA_2() constant returns(address)
func (_SCPHA1 *SCPHA1Caller) AddrSCPHA2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "addr_SC_P_HA_2")
	return *ret0, err
}

// AddrSCPHA2 is a free data retrieval call binding the contract method 0x82dab224.
//
// Solidity: function addr_SC_P_HA_2() constant returns(address)
func (_SCPHA1 *SCPHA1Session) AddrSCPHA2() (common.Address, error) {
	return _SCPHA1.Contract.AddrSCPHA2(&_SCPHA1.CallOpts)
}

// AddrSCPHA2 is a free data retrieval call binding the contract method 0x82dab224.
//
// Solidity: function addr_SC_P_HA_2() constant returns(address)
func (_SCPHA1 *SCPHA1CallerSession) AddrSCPHA2() (common.Address, error) {
	return _SCPHA1.Contract.AddrSCPHA2(&_SCPHA1.CallOpts)
}

// Dispute is a free data retrieval call binding the contract method 0x86d6282c.
//
// Solidity: function dispute(uint256 ) constant returns(uint256 disputeID, uint256 pID, uint256 hID, uint256 finalBillID, uint256 timestampOfRaisingDispute, uint256 timestampOfSolvingDispute, bool finalCostChanged)
func (_SCPHA1 *SCPHA1Caller) Dispute(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DisputeID                 *big.Int
	PID                       *big.Int
	HID                       *big.Int
	FinalBillID               *big.Int
	TimestampOfRaisingDispute *big.Int
	TimestampOfSolvingDispute *big.Int
	FinalCostChanged          bool
}, error) {
	ret := new(struct {
		DisputeID                 *big.Int
		PID                       *big.Int
		HID                       *big.Int
		FinalBillID               *big.Int
		TimestampOfRaisingDispute *big.Int
		TimestampOfSolvingDispute *big.Int
		FinalCostChanged          bool
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "dispute", arg0)
	return *ret, err
}

// Dispute is a free data retrieval call binding the contract method 0x86d6282c.
//
// Solidity: function dispute(uint256 ) constant returns(uint256 disputeID, uint256 pID, uint256 hID, uint256 finalBillID, uint256 timestampOfRaisingDispute, uint256 timestampOfSolvingDispute, bool finalCostChanged)
func (_SCPHA1 *SCPHA1Session) Dispute(arg0 *big.Int) (struct {
	DisputeID                 *big.Int
	PID                       *big.Int
	HID                       *big.Int
	FinalBillID               *big.Int
	TimestampOfRaisingDispute *big.Int
	TimestampOfSolvingDispute *big.Int
	FinalCostChanged          bool
}, error) {
	return _SCPHA1.Contract.Dispute(&_SCPHA1.CallOpts, arg0)
}

// Dispute is a free data retrieval call binding the contract method 0x86d6282c.
//
// Solidity: function dispute(uint256 ) constant returns(uint256 disputeID, uint256 pID, uint256 hID, uint256 finalBillID, uint256 timestampOfRaisingDispute, uint256 timestampOfSolvingDispute, bool finalCostChanged)
func (_SCPHA1 *SCPHA1CallerSession) Dispute(arg0 *big.Int) (struct {
	DisputeID                 *big.Int
	PID                       *big.Int
	HID                       *big.Int
	FinalBillID               *big.Int
	TimestampOfRaisingDispute *big.Int
	TimestampOfSolvingDispute *big.Int
	FinalCostChanged          bool
}, error) {
	return _SCPHA1.Contract.Dispute(&_SCPHA1.CallOpts, arg0)
}

// DisputeIDGenerator is a free data retrieval call binding the contract method 0x427d517e.
//
// Solidity: function disputeIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) DisputeIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "disputeIDGenerator")
	return *ret0, err
}

// DisputeIDGenerator is a free data retrieval call binding the contract method 0x427d517e.
//
// Solidity: function disputeIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) DisputeIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.DisputeIDGenerator(&_SCPHA1.CallOpts)
}

// DisputeIDGenerator is a free data retrieval call binding the contract method 0x427d517e.
//
// Solidity: function disputeIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) DisputeIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.DisputeIDGenerator(&_SCPHA1.CallOpts)
}

// EstimatedBillIDGenerator is a free data retrieval call binding the contract method 0xd215ea8b.
//
// Solidity: function estimatedBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) EstimatedBillIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "estimatedBillIDGenerator")
	return *ret0, err
}

// EstimatedBillIDGenerator is a free data retrieval call binding the contract method 0xd215ea8b.
//
// Solidity: function estimatedBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) EstimatedBillIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedBillIDGenerator(&_SCPHA1.CallOpts)
}

// EstimatedBillIDGenerator is a free data retrieval call binding the contract method 0xd215ea8b.
//
// Solidity: function estimatedBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) EstimatedBillIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedBillIDGenerator(&_SCPHA1.CallOpts)
}

// EstimatedBillIDToFinalBillID is a free data retrieval call binding the contract method 0x8e2d5531.
//
// Solidity: function estimatedBillIDToFinalBillID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) EstimatedBillIDToFinalBillID(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "estimatedBillIDToFinalBillID", arg0)
	return *ret0, err
}

// EstimatedBillIDToFinalBillID is a free data retrieval call binding the contract method 0x8e2d5531.
//
// Solidity: function estimatedBillIDToFinalBillID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) EstimatedBillIDToFinalBillID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedBillIDToFinalBillID(&_SCPHA1.CallOpts, arg0)
}

// EstimatedBillIDToFinalBillID is a free data retrieval call binding the contract method 0x8e2d5531.
//
// Solidity: function estimatedBillIDToFinalBillID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) EstimatedBillIDToFinalBillID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedBillIDToFinalBillID(&_SCPHA1.CallOpts, arg0)
}

// EstimatedCheckUpCost is a free data retrieval call binding the contract method 0x59e04b3a.
//
// Solidity: function estimatedCheckUpCost(uint256 ) constant returns(uint256 estimatedCostBillID, uint256 pID, uint256 hID, uint256 estimatedCost, uint256 timestampOfEstimate, uint256 timestampOfLockingByHospital, uint256 timestampOfLockingByPatient, uint256 timestampOfCheckUpStart, uint256 timestampOfUnlockingByPatient, uint256 timestampOfUnlockingByHospital)
func (_SCPHA1 *SCPHA1Caller) EstimatedCheckUpCost(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EstimatedCostBillID            *big.Int
	PID                            *big.Int
	HID                            *big.Int
	EstimatedCost                  *big.Int
	TimestampOfEstimate            *big.Int
	TimestampOfLockingByHospital   *big.Int
	TimestampOfLockingByPatient    *big.Int
	TimestampOfCheckUpStart        *big.Int
	TimestampOfUnlockingByPatient  *big.Int
	TimestampOfUnlockingByHospital *big.Int
}, error) {
	ret := new(struct {
		EstimatedCostBillID            *big.Int
		PID                            *big.Int
		HID                            *big.Int
		EstimatedCost                  *big.Int
		TimestampOfEstimate            *big.Int
		TimestampOfLockingByHospital   *big.Int
		TimestampOfLockingByPatient    *big.Int
		TimestampOfCheckUpStart        *big.Int
		TimestampOfUnlockingByPatient  *big.Int
		TimestampOfUnlockingByHospital *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "estimatedCheckUpCost", arg0)
	return *ret, err
}

// EstimatedCheckUpCost is a free data retrieval call binding the contract method 0x59e04b3a.
//
// Solidity: function estimatedCheckUpCost(uint256 ) constant returns(uint256 estimatedCostBillID, uint256 pID, uint256 hID, uint256 estimatedCost, uint256 timestampOfEstimate, uint256 timestampOfLockingByHospital, uint256 timestampOfLockingByPatient, uint256 timestampOfCheckUpStart, uint256 timestampOfUnlockingByPatient, uint256 timestampOfUnlockingByHospital)
func (_SCPHA1 *SCPHA1Session) EstimatedCheckUpCost(arg0 *big.Int) (struct {
	EstimatedCostBillID            *big.Int
	PID                            *big.Int
	HID                            *big.Int
	EstimatedCost                  *big.Int
	TimestampOfEstimate            *big.Int
	TimestampOfLockingByHospital   *big.Int
	TimestampOfLockingByPatient    *big.Int
	TimestampOfCheckUpStart        *big.Int
	TimestampOfUnlockingByPatient  *big.Int
	TimestampOfUnlockingByHospital *big.Int
}, error) {
	return _SCPHA1.Contract.EstimatedCheckUpCost(&_SCPHA1.CallOpts, arg0)
}

// EstimatedCheckUpCost is a free data retrieval call binding the contract method 0x59e04b3a.
//
// Solidity: function estimatedCheckUpCost(uint256 ) constant returns(uint256 estimatedCostBillID, uint256 pID, uint256 hID, uint256 estimatedCost, uint256 timestampOfEstimate, uint256 timestampOfLockingByHospital, uint256 timestampOfLockingByPatient, uint256 timestampOfCheckUpStart, uint256 timestampOfUnlockingByPatient, uint256 timestampOfUnlockingByHospital)
func (_SCPHA1 *SCPHA1CallerSession) EstimatedCheckUpCost(arg0 *big.Int) (struct {
	EstimatedCostBillID            *big.Int
	PID                            *big.Int
	HID                            *big.Int
	EstimatedCost                  *big.Int
	TimestampOfEstimate            *big.Int
	TimestampOfLockingByHospital   *big.Int
	TimestampOfLockingByPatient    *big.Int
	TimestampOfCheckUpStart        *big.Int
	TimestampOfUnlockingByPatient  *big.Int
	TimestampOfUnlockingByHospital *big.Int
}, error) {
	return _SCPHA1.Contract.EstimatedCheckUpCost(&_SCPHA1.CallOpts, arg0)
}

// EstimatedCostBillIDToMultiSigOnMedicalDataID is a free data retrieval call binding the contract method 0xb90d4235.
//
// Solidity: function estimatedCostBillIDToMultiSigOnMedicalDataID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) EstimatedCostBillIDToMultiSigOnMedicalDataID(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "estimatedCostBillIDToMultiSigOnMedicalDataID", arg0)
	return *ret0, err
}

// EstimatedCostBillIDToMultiSigOnMedicalDataID is a free data retrieval call binding the contract method 0xb90d4235.
//
// Solidity: function estimatedCostBillIDToMultiSigOnMedicalDataID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) EstimatedCostBillIDToMultiSigOnMedicalDataID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedCostBillIDToMultiSigOnMedicalDataID(&_SCPHA1.CallOpts, arg0)
}

// EstimatedCostBillIDToMultiSigOnMedicalDataID is a free data retrieval call binding the contract method 0xb90d4235.
//
// Solidity: function estimatedCostBillIDToMultiSigOnMedicalDataID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) EstimatedCostBillIDToMultiSigOnMedicalDataID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.EstimatedCostBillIDToMultiSigOnMedicalDataID(&_SCPHA1.CallOpts, arg0)
}

// Excess is a free data retrieval call binding the contract method 0x8d19b2de.
//
// Solidity: function excess(uint256 ) constant returns(uint256 excessID, uint256 pID, uint256 hID, uint256 excessAmount, uint256 timestampOfRefundingExcessToPatient)
func (_SCPHA1 *SCPHA1Caller) Excess(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ExcessID                            *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	ExcessAmount                        *big.Int
	TimestampOfRefundingExcessToPatient *big.Int
}, error) {
	ret := new(struct {
		ExcessID                            *big.Int
		PID                                 *big.Int
		HID                                 *big.Int
		ExcessAmount                        *big.Int
		TimestampOfRefundingExcessToPatient *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "excess", arg0)
	return *ret, err
}

// Excess is a free data retrieval call binding the contract method 0x8d19b2de.
//
// Solidity: function excess(uint256 ) constant returns(uint256 excessID, uint256 pID, uint256 hID, uint256 excessAmount, uint256 timestampOfRefundingExcessToPatient)
func (_SCPHA1 *SCPHA1Session) Excess(arg0 *big.Int) (struct {
	ExcessID                            *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	ExcessAmount                        *big.Int
	TimestampOfRefundingExcessToPatient *big.Int
}, error) {
	return _SCPHA1.Contract.Excess(&_SCPHA1.CallOpts, arg0)
}

// Excess is a free data retrieval call binding the contract method 0x8d19b2de.
//
// Solidity: function excess(uint256 ) constant returns(uint256 excessID, uint256 pID, uint256 hID, uint256 excessAmount, uint256 timestampOfRefundingExcessToPatient)
func (_SCPHA1 *SCPHA1CallerSession) Excess(arg0 *big.Int) (struct {
	ExcessID                            *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	ExcessAmount                        *big.Int
	TimestampOfRefundingExcessToPatient *big.Int
}, error) {
	return _SCPHA1.Contract.Excess(&_SCPHA1.CallOpts, arg0)
}

// ExcessIDGenerator is a free data retrieval call binding the contract method 0x40841aa5.
//
// Solidity: function excessIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) ExcessIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "excessIDGenerator")
	return *ret0, err
}

// ExcessIDGenerator is a free data retrieval call binding the contract method 0x40841aa5.
//
// Solidity: function excessIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) ExcessIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.ExcessIDGenerator(&_SCPHA1.CallOpts)
}

// ExcessIDGenerator is a free data retrieval call binding the contract method 0x40841aa5.
//
// Solidity: function excessIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) ExcessIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.ExcessIDGenerator(&_SCPHA1.CallOpts)
}

// FinalBillIDGenerator is a free data retrieval call binding the contract method 0xb642a283.
//
// Solidity: function finalBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) FinalBillIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "finalBillIDGenerator")
	return *ret0, err
}

// FinalBillIDGenerator is a free data retrieval call binding the contract method 0xb642a283.
//
// Solidity: function finalBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) FinalBillIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.FinalBillIDGenerator(&_SCPHA1.CallOpts)
}

// FinalBillIDGenerator is a free data retrieval call binding the contract method 0xb642a283.
//
// Solidity: function finalBillIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) FinalBillIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.FinalBillIDGenerator(&_SCPHA1.CallOpts)
}

// FinalBillIDToDisputeID is a free data retrieval call binding the contract method 0x0d84410d.
//
// Solidity: function finalBillIDToDisputeID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) FinalBillIDToDisputeID(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "finalBillIDToDisputeID", arg0)
	return *ret0, err
}

// FinalBillIDToDisputeID is a free data retrieval call binding the contract method 0x0d84410d.
//
// Solidity: function finalBillIDToDisputeID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) FinalBillIDToDisputeID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.FinalBillIDToDisputeID(&_SCPHA1.CallOpts, arg0)
}

// FinalBillIDToDisputeID is a free data retrieval call binding the contract method 0x0d84410d.
//
// Solidity: function finalBillIDToDisputeID(uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) FinalBillIDToDisputeID(arg0 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.FinalBillIDToDisputeID(&_SCPHA1.CallOpts, arg0)
}

// FinalCheckUpCost is a free data retrieval call binding the contract method 0x0c89bf44.
//
// Solidity: function finalCheckUpCost(uint256 ) constant returns(uint256 finalCostBillID, uint256 pID, uint256 hID, uint256 finalCost, uint256 timestampOfFinalBilling, uint256 excessID, uint256 timeStampOfPatientFinalBillConsent, uint256 timeStampOfFinalConsent, uint256 timeStampOfComplain, uint256 timeStampOfUnlockingByHospital_final, uint256 timestampOfUnlockingByPatient_final)
func (_SCPHA1 *SCPHA1Caller) FinalCheckUpCost(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FinalCostBillID                     *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	FinalCost                           *big.Int
	TimestampOfFinalBilling             *big.Int
	ExcessID                            *big.Int
	TimeStampOfPatientFinalBillConsent  *big.Int
	TimeStampOfFinalConsent             *big.Int
	TimeStampOfComplain                 *big.Int
	TimeStampOfUnlockingByHospitalFinal *big.Int
	TimestampOfUnlockingByPatientFinal  *big.Int
}, error) {
	ret := new(struct {
		FinalCostBillID                     *big.Int
		PID                                 *big.Int
		HID                                 *big.Int
		FinalCost                           *big.Int
		TimestampOfFinalBilling             *big.Int
		ExcessID                            *big.Int
		TimeStampOfPatientFinalBillConsent  *big.Int
		TimeStampOfFinalConsent             *big.Int
		TimeStampOfComplain                 *big.Int
		TimeStampOfUnlockingByHospitalFinal *big.Int
		TimestampOfUnlockingByPatientFinal  *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "finalCheckUpCost", arg0)
	return *ret, err
}

// FinalCheckUpCost is a free data retrieval call binding the contract method 0x0c89bf44.
//
// Solidity: function finalCheckUpCost(uint256 ) constant returns(uint256 finalCostBillID, uint256 pID, uint256 hID, uint256 finalCost, uint256 timestampOfFinalBilling, uint256 excessID, uint256 timeStampOfPatientFinalBillConsent, uint256 timeStampOfFinalConsent, uint256 timeStampOfComplain, uint256 timeStampOfUnlockingByHospital_final, uint256 timestampOfUnlockingByPatient_final)
func (_SCPHA1 *SCPHA1Session) FinalCheckUpCost(arg0 *big.Int) (struct {
	FinalCostBillID                     *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	FinalCost                           *big.Int
	TimestampOfFinalBilling             *big.Int
	ExcessID                            *big.Int
	TimeStampOfPatientFinalBillConsent  *big.Int
	TimeStampOfFinalConsent             *big.Int
	TimeStampOfComplain                 *big.Int
	TimeStampOfUnlockingByHospitalFinal *big.Int
	TimestampOfUnlockingByPatientFinal  *big.Int
}, error) {
	return _SCPHA1.Contract.FinalCheckUpCost(&_SCPHA1.CallOpts, arg0)
}

// FinalCheckUpCost is a free data retrieval call binding the contract method 0x0c89bf44.
//
// Solidity: function finalCheckUpCost(uint256 ) constant returns(uint256 finalCostBillID, uint256 pID, uint256 hID, uint256 finalCost, uint256 timestampOfFinalBilling, uint256 excessID, uint256 timeStampOfPatientFinalBillConsent, uint256 timeStampOfFinalConsent, uint256 timeStampOfComplain, uint256 timeStampOfUnlockingByHospital_final, uint256 timestampOfUnlockingByPatient_final)
func (_SCPHA1 *SCPHA1CallerSession) FinalCheckUpCost(arg0 *big.Int) (struct {
	FinalCostBillID                     *big.Int
	PID                                 *big.Int
	HID                                 *big.Int
	FinalCost                           *big.Int
	TimestampOfFinalBilling             *big.Int
	ExcessID                            *big.Int
	TimeStampOfPatientFinalBillConsent  *big.Int
	TimeStampOfFinalConsent             *big.Int
	TimeStampOfComplain                 *big.Int
	TimeStampOfUnlockingByHospitalFinal *big.Int
	TimestampOfUnlockingByPatientFinal  *big.Int
}, error) {
	return _SCPHA1.Contract.FinalCheckUpCost(&_SCPHA1.CallOpts, arg0)
}

// GetDisputeStruct is a free data retrieval call binding the contract method 0x01a2a657.
//
// Solidity: function getDisputeStruct(uint256 id) constant returns(SC_P_HA_1Dispute)
func (_SCPHA1 *SCPHA1Caller) GetDisputeStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1Dispute, error) {
	var (
		ret0 = new(SC_P_HA_1Dispute)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getDisputeStruct", id)
	return *ret0, err
}

// GetDisputeStruct is a free data retrieval call binding the contract method 0x01a2a657.
//
// Solidity: function getDisputeStruct(uint256 id) constant returns(SC_P_HA_1Dispute)
func (_SCPHA1 *SCPHA1Session) GetDisputeStruct(id *big.Int) (SC_P_HA_1Dispute, error) {
	return _SCPHA1.Contract.GetDisputeStruct(&_SCPHA1.CallOpts, id)
}

// GetDisputeStruct is a free data retrieval call binding the contract method 0x01a2a657.
//
// Solidity: function getDisputeStruct(uint256 id) constant returns(SC_P_HA_1Dispute)
func (_SCPHA1 *SCPHA1CallerSession) GetDisputeStruct(id *big.Int) (SC_P_HA_1Dispute, error) {
	return _SCPHA1.Contract.GetDisputeStruct(&_SCPHA1.CallOpts, id)
}

// GetEncFileProperties is a free data retrieval call binding the contract method 0x1082b93b.
//
// Solidity: function getEncFileProperties(uint256 id) constant returns(SC_P_HA_1FileProperties)
func (_SCPHA1 *SCPHA1Caller) GetEncFileProperties(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1FileProperties, error) {
	var (
		ret0 = new(SC_P_HA_1FileProperties)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getEncFileProperties", id)
	return *ret0, err
}

// GetEncFileProperties is a free data retrieval call binding the contract method 0x1082b93b.
//
// Solidity: function getEncFileProperties(uint256 id) constant returns(SC_P_HA_1FileProperties)
func (_SCPHA1 *SCPHA1Session) GetEncFileProperties(id *big.Int) (SC_P_HA_1FileProperties, error) {
	return _SCPHA1.Contract.GetEncFileProperties(&_SCPHA1.CallOpts, id)
}

// GetEncFileProperties is a free data retrieval call binding the contract method 0x1082b93b.
//
// Solidity: function getEncFileProperties(uint256 id) constant returns(SC_P_HA_1FileProperties)
func (_SCPHA1 *SCPHA1CallerSession) GetEncFileProperties(id *big.Int) (SC_P_HA_1FileProperties, error) {
	return _SCPHA1.Contract.GetEncFileProperties(&_SCPHA1.CallOpts, id)
}

// GetEstimatedCheckUpCostStruct is a free data retrieval call binding the contract method 0x3544ff50.
//
// Solidity: function getEstimatedCheckUpCostStruct(uint256 id) constant returns(SC_P_HA_1EstimatedCheckUpCost)
func (_SCPHA1 *SCPHA1Caller) GetEstimatedCheckUpCostStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1EstimatedCheckUpCost, error) {
	var (
		ret0 = new(SC_P_HA_1EstimatedCheckUpCost)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getEstimatedCheckUpCostStruct", id)
	return *ret0, err
}

// GetEstimatedCheckUpCostStruct is a free data retrieval call binding the contract method 0x3544ff50.
//
// Solidity: function getEstimatedCheckUpCostStruct(uint256 id) constant returns(SC_P_HA_1EstimatedCheckUpCost)
func (_SCPHA1 *SCPHA1Session) GetEstimatedCheckUpCostStruct(id *big.Int) (SC_P_HA_1EstimatedCheckUpCost, error) {
	return _SCPHA1.Contract.GetEstimatedCheckUpCostStruct(&_SCPHA1.CallOpts, id)
}

// GetEstimatedCheckUpCostStruct is a free data retrieval call binding the contract method 0x3544ff50.
//
// Solidity: function getEstimatedCheckUpCostStruct(uint256 id) constant returns(SC_P_HA_1EstimatedCheckUpCost)
func (_SCPHA1 *SCPHA1CallerSession) GetEstimatedCheckUpCostStruct(id *big.Int) (SC_P_HA_1EstimatedCheckUpCost, error) {
	return _SCPHA1.Contract.GetEstimatedCheckUpCostStruct(&_SCPHA1.CallOpts, id)
}

// GetExcessStruct is a free data retrieval call binding the contract method 0x6b83401d.
//
// Solidity: function getExcessStruct(uint256 id) constant returns(SC_P_HA_1Excess)
func (_SCPHA1 *SCPHA1Caller) GetExcessStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1Excess, error) {
	var (
		ret0 = new(SC_P_HA_1Excess)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getExcessStruct", id)
	return *ret0, err
}

// GetExcessStruct is a free data retrieval call binding the contract method 0x6b83401d.
//
// Solidity: function getExcessStruct(uint256 id) constant returns(SC_P_HA_1Excess)
func (_SCPHA1 *SCPHA1Session) GetExcessStruct(id *big.Int) (SC_P_HA_1Excess, error) {
	return _SCPHA1.Contract.GetExcessStruct(&_SCPHA1.CallOpts, id)
}

// GetExcessStruct is a free data retrieval call binding the contract method 0x6b83401d.
//
// Solidity: function getExcessStruct(uint256 id) constant returns(SC_P_HA_1Excess)
func (_SCPHA1 *SCPHA1CallerSession) GetExcessStruct(id *big.Int) (SC_P_HA_1Excess, error) {
	return _SCPHA1.Contract.GetExcessStruct(&_SCPHA1.CallOpts, id)
}

// GetFinalBillStruct is a free data retrieval call binding the contract method 0x0f6df082.
//
// Solidity: function getFinalBillStruct(uint256 id) constant returns(SC_P_HA_1FinalCheckUpCost)
func (_SCPHA1 *SCPHA1Caller) GetFinalBillStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1FinalCheckUpCost, error) {
	var (
		ret0 = new(SC_P_HA_1FinalCheckUpCost)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getFinalBillStruct", id)
	return *ret0, err
}

// GetFinalBillStruct is a free data retrieval call binding the contract method 0x0f6df082.
//
// Solidity: function getFinalBillStruct(uint256 id) constant returns(SC_P_HA_1FinalCheckUpCost)
func (_SCPHA1 *SCPHA1Session) GetFinalBillStruct(id *big.Int) (SC_P_HA_1FinalCheckUpCost, error) {
	return _SCPHA1.Contract.GetFinalBillStruct(&_SCPHA1.CallOpts, id)
}

// GetFinalBillStruct is a free data retrieval call binding the contract method 0x0f6df082.
//
// Solidity: function getFinalBillStruct(uint256 id) constant returns(SC_P_HA_1FinalCheckUpCost)
func (_SCPHA1 *SCPHA1CallerSession) GetFinalBillStruct(id *big.Int) (SC_P_HA_1FinalCheckUpCost, error) {
	return _SCPHA1.Contract.GetFinalBillStruct(&_SCPHA1.CallOpts, id)
}

// GetKeyAndExchangeStruct is a free data retrieval call binding the contract method 0x4f5a4e1c.
//
// Solidity: function getKeyAndExchangeStruct(uint256 id) constant returns(SC_P_HA_1KeyAndExchange)
func (_SCPHA1 *SCPHA1Caller) GetKeyAndExchangeStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1KeyAndExchange, error) {
	var (
		ret0 = new(SC_P_HA_1KeyAndExchange)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getKeyAndExchangeStruct", id)
	return *ret0, err
}

// GetKeyAndExchangeStruct is a free data retrieval call binding the contract method 0x4f5a4e1c.
//
// Solidity: function getKeyAndExchangeStruct(uint256 id) constant returns(SC_P_HA_1KeyAndExchange)
func (_SCPHA1 *SCPHA1Session) GetKeyAndExchangeStruct(id *big.Int) (SC_P_HA_1KeyAndExchange, error) {
	return _SCPHA1.Contract.GetKeyAndExchangeStruct(&_SCPHA1.CallOpts, id)
}

// GetKeyAndExchangeStruct is a free data retrieval call binding the contract method 0x4f5a4e1c.
//
// Solidity: function getKeyAndExchangeStruct(uint256 id) constant returns(SC_P_HA_1KeyAndExchange)
func (_SCPHA1 *SCPHA1CallerSession) GetKeyAndExchangeStruct(id *big.Int) (SC_P_HA_1KeyAndExchange, error) {
	return _SCPHA1.Contract.GetKeyAndExchangeStruct(&_SCPHA1.CallOpts, id)
}

// GetMultiSigOnMedicalDataStruct is a free data retrieval call binding the contract method 0x8275370a.
//
// Solidity: function getMultiSigOnMedicalDataStruct(uint256 id) constant returns(SC_P_HA_1MultiSigOnMedicalData)
func (_SCPHA1 *SCPHA1Caller) GetMultiSigOnMedicalDataStruct(opts *bind.CallOpts, id *big.Int) (SC_P_HA_1MultiSigOnMedicalData, error) {
	var (
		ret0 = new(SC_P_HA_1MultiSigOnMedicalData)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "getMultiSigOnMedicalDataStruct", id)
	return *ret0, err
}

// GetMultiSigOnMedicalDataStruct is a free data retrieval call binding the contract method 0x8275370a.
//
// Solidity: function getMultiSigOnMedicalDataStruct(uint256 id) constant returns(SC_P_HA_1MultiSigOnMedicalData)
func (_SCPHA1 *SCPHA1Session) GetMultiSigOnMedicalDataStruct(id *big.Int) (SC_P_HA_1MultiSigOnMedicalData, error) {
	return _SCPHA1.Contract.GetMultiSigOnMedicalDataStruct(&_SCPHA1.CallOpts, id)
}

// GetMultiSigOnMedicalDataStruct is a free data retrieval call binding the contract method 0x8275370a.
//
// Solidity: function getMultiSigOnMedicalDataStruct(uint256 id) constant returns(SC_P_HA_1MultiSigOnMedicalData)
func (_SCPHA1 *SCPHA1CallerSession) GetMultiSigOnMedicalDataStruct(id *big.Int) (SC_P_HA_1MultiSigOnMedicalData, error) {
	return _SCPHA1.Contract.GetMultiSigOnMedicalDataStruct(&_SCPHA1.CallOpts, id)
}

// MultiSigIDGenerator is a free data retrieval call binding the contract method 0xec948a38.
//
// Solidity: function multiSigIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) MultiSigIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "multiSigIDGenerator")
	return *ret0, err
}

// MultiSigIDGenerator is a free data retrieval call binding the contract method 0xec948a38.
//
// Solidity: function multiSigIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) MultiSigIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.MultiSigIDGenerator(&_SCPHA1.CallOpts)
}

// MultiSigIDGenerator is a free data retrieval call binding the contract method 0xec948a38.
//
// Solidity: function multiSigIDGenerator() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) MultiSigIDGenerator() (*big.Int, error) {
	return _SCPHA1.Contract.MultiSigIDGenerator(&_SCPHA1.CallOpts)
}

// MultiSigOnMedicalData is a free data retrieval call binding the contract method 0xc50eaa14.
//
// Solidity: function multiSigOnMedicalData(uint256 ) constant returns(uint256 multiSigID, uint256 patientID, uint256 hospitalID, bytes32 hashOfData, bytes32 hashOfEncryptedData, bytes32 hashOfPatientIdDateHashEncFile, bytes signOnFileHashByHospital, bytes signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 timeStampOfSignByHospital, uint256 timeStampOfVerificationByPatient, bool verified, uint256 timeStampOfUnlockingByPatient_multi, uint256 timeStampOfUnlockingByHospital_multi)
func (_SCPHA1 *SCPHA1Caller) MultiSigOnMedicalData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MultiSigID                                           *big.Int
	PatientID                                            *big.Int
	HospitalID                                           *big.Int
	HashOfData                                           [32]byte
	HashOfEncryptedData                                  [32]byte
	HashOfPatientIdDateHashEncFile                       [32]byte
	SignOnFileHashByHospital                             []byte
	SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile []byte
	TimeStampOfSignByHospital                            *big.Int
	TimeStampOfVerificationByPatient                     *big.Int
	Verified                                             bool
	TimeStampOfUnlockingByPatientMulti                   *big.Int
	TimeStampOfUnlockingByHospitalMulti                  *big.Int
}, error) {
	ret := new(struct {
		MultiSigID                                           *big.Int
		PatientID                                            *big.Int
		HospitalID                                           *big.Int
		HashOfData                                           [32]byte
		HashOfEncryptedData                                  [32]byte
		HashOfPatientIdDateHashEncFile                       [32]byte
		SignOnFileHashByHospital                             []byte
		SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile []byte
		TimeStampOfSignByHospital                            *big.Int
		TimeStampOfVerificationByPatient                     *big.Int
		Verified                                             bool
		TimeStampOfUnlockingByPatientMulti                   *big.Int
		TimeStampOfUnlockingByHospitalMulti                  *big.Int
	})
	out := ret
	err := _SCPHA1.contract.Call(opts, out, "multiSigOnMedicalData", arg0)
	return *ret, err
}

// MultiSigOnMedicalData is a free data retrieval call binding the contract method 0xc50eaa14.
//
// Solidity: function multiSigOnMedicalData(uint256 ) constant returns(uint256 multiSigID, uint256 patientID, uint256 hospitalID, bytes32 hashOfData, bytes32 hashOfEncryptedData, bytes32 hashOfPatientIdDateHashEncFile, bytes signOnFileHashByHospital, bytes signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 timeStampOfSignByHospital, uint256 timeStampOfVerificationByPatient, bool verified, uint256 timeStampOfUnlockingByPatient_multi, uint256 timeStampOfUnlockingByHospital_multi)
func (_SCPHA1 *SCPHA1Session) MultiSigOnMedicalData(arg0 *big.Int) (struct {
	MultiSigID                                           *big.Int
	PatientID                                            *big.Int
	HospitalID                                           *big.Int
	HashOfData                                           [32]byte
	HashOfEncryptedData                                  [32]byte
	HashOfPatientIdDateHashEncFile                       [32]byte
	SignOnFileHashByHospital                             []byte
	SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile []byte
	TimeStampOfSignByHospital                            *big.Int
	TimeStampOfVerificationByPatient                     *big.Int
	Verified                                             bool
	TimeStampOfUnlockingByPatientMulti                   *big.Int
	TimeStampOfUnlockingByHospitalMulti                  *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSigOnMedicalData(&_SCPHA1.CallOpts, arg0)
}

// MultiSigOnMedicalData is a free data retrieval call binding the contract method 0xc50eaa14.
//
// Solidity: function multiSigOnMedicalData(uint256 ) constant returns(uint256 multiSigID, uint256 patientID, uint256 hospitalID, bytes32 hashOfData, bytes32 hashOfEncryptedData, bytes32 hashOfPatientIdDateHashEncFile, bytes signOnFileHashByHospital, bytes signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 timeStampOfSignByHospital, uint256 timeStampOfVerificationByPatient, bool verified, uint256 timeStampOfUnlockingByPatient_multi, uint256 timeStampOfUnlockingByHospital_multi)
func (_SCPHA1 *SCPHA1CallerSession) MultiSigOnMedicalData(arg0 *big.Int) (struct {
	MultiSigID                                           *big.Int
	PatientID                                            *big.Int
	HospitalID                                           *big.Int
	HashOfData                                           [32]byte
	HashOfEncryptedData                                  [32]byte
	HashOfPatientIdDateHashEncFile                       [32]byte
	SignOnFileHashByHospital                             []byte
	SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile []byte
	TimeStampOfSignByHospital                            *big.Int
	TimeStampOfVerificationByPatient                     *big.Int
	Verified                                             bool
	TimeStampOfUnlockingByPatientMulti                   *big.Int
	TimeStampOfUnlockingByHospitalMulti                  *big.Int
}, error) {
	return _SCPHA1.Contract.MultiSigOnMedicalData(&_SCPHA1.CallOpts, arg0)
}

// PatientIdHospitalIdEstimateBillId is a free data retrieval call binding the contract method 0xe77cc243.
//
// Solidity: function patientId_hospitalId_estimateBillId(uint256 , uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) PatientIdHospitalIdEstimateBillId(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "patientId_hospitalId_estimateBillId", arg0, arg1)
	return *ret0, err
}

// PatientIdHospitalIdEstimateBillId is a free data retrieval call binding the contract method 0xe77cc243.
//
// Solidity: function patientId_hospitalId_estimateBillId(uint256 , uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) PatientIdHospitalIdEstimateBillId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.PatientIdHospitalIdEstimateBillId(&_SCPHA1.CallOpts, arg0, arg1)
}

// PatientIdHospitalIdEstimateBillId is a free data retrieval call binding the contract method 0xe77cc243.
//
// Solidity: function patientId_hospitalId_estimateBillId(uint256 , uint256 ) constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) PatientIdHospitalIdEstimateBillId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _SCPHA1.Contract.PatientIdHospitalIdEstimateBillId(&_SCPHA1.CallOpts, arg0, arg1)
}

// TimeLimitForAcceptingOrProducingComplain is a free data retrieval call binding the contract method 0x8513f5e5.
//
// Solidity: function timeLimitForAcceptingOrProducingComplain() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForAcceptingOrProducingComplain(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForAcceptingOrProducingComplain")
	return *ret0, err
}

// TimeLimitForAcceptingOrProducingComplain is a free data retrieval call binding the contract method 0x8513f5e5.
//
// Solidity: function timeLimitForAcceptingOrProducingComplain() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForAcceptingOrProducingComplain() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForAcceptingOrProducingComplain(&_SCPHA1.CallOpts)
}

// TimeLimitForAcceptingOrProducingComplain is a free data retrieval call binding the contract method 0x8513f5e5.
//
// Solidity: function timeLimitForAcceptingOrProducingComplain() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForAcceptingOrProducingComplain() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForAcceptingOrProducingComplain(&_SCPHA1.CallOpts)
}

// TimeLimitForGeneratingFillBill is a free data retrieval call binding the contract method 0x672e7b21.
//
// Solidity: function timeLimitForGeneratingFillBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForGeneratingFillBill(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForGeneratingFillBill")
	return *ret0, err
}

// TimeLimitForGeneratingFillBill is a free data retrieval call binding the contract method 0x672e7b21.
//
// Solidity: function timeLimitForGeneratingFillBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForGeneratingFillBill() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForGeneratingFillBill(&_SCPHA1.CallOpts)
}

// TimeLimitForGeneratingFillBill is a free data retrieval call binding the contract method 0x672e7b21.
//
// Solidity: function timeLimitForGeneratingFillBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForGeneratingFillBill() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForGeneratingFillBill(&_SCPHA1.CallOpts)
}

// TimeLimitForKeyRevealing is a free data retrieval call binding the contract method 0x361ff9b3.
//
// Solidity: function timeLimitForKeyRevealing() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForKeyRevealing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForKeyRevealing")
	return *ret0, err
}

// TimeLimitForKeyRevealing is a free data retrieval call binding the contract method 0x361ff9b3.
//
// Solidity: function timeLimitForKeyRevealing() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForKeyRevealing() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForKeyRevealing(&_SCPHA1.CallOpts)
}

// TimeLimitForKeyRevealing is a free data retrieval call binding the contract method 0x361ff9b3.
//
// Solidity: function timeLimitForKeyRevealing() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForKeyRevealing() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForKeyRevealing(&_SCPHA1.CallOpts)
}

// TimeLimitForPatientLocking is a free data retrieval call binding the contract method 0x25e3bb73.
//
// Solidity: function timeLimitForPatientLocking() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForPatientLocking(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForPatientLocking")
	return *ret0, err
}

// TimeLimitForPatientLocking is a free data retrieval call binding the contract method 0x25e3bb73.
//
// Solidity: function timeLimitForPatientLocking() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForPatientLocking() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForPatientLocking(&_SCPHA1.CallOpts)
}

// TimeLimitForPatientLocking is a free data retrieval call binding the contract method 0x25e3bb73.
//
// Solidity: function timeLimitForPatientLocking() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForPatientLocking() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForPatientLocking(&_SCPHA1.CallOpts)
}

// TimeLimitForRaisingDispute is a free data retrieval call binding the contract method 0xb17c31fb.
//
// Solidity: function timeLimitForRaisingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForRaisingDispute(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForRaisingDispute")
	return *ret0, err
}

// TimeLimitForRaisingDispute is a free data retrieval call binding the contract method 0xb17c31fb.
//
// Solidity: function timeLimitForRaisingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForRaisingDispute() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForRaisingDispute(&_SCPHA1.CallOpts)
}

// TimeLimitForRaisingDispute is a free data retrieval call binding the contract method 0xb17c31fb.
//
// Solidity: function timeLimitForRaisingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForRaisingDispute() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForRaisingDispute(&_SCPHA1.CallOpts)
}

// TimeLimitForSigningFinalBill is a free data retrieval call binding the contract method 0xfdcd9147.
//
// Solidity: function timeLimitForSigningFinalBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForSigningFinalBill(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForSigningFinalBill")
	return *ret0, err
}

// TimeLimitForSigningFinalBill is a free data retrieval call binding the contract method 0xfdcd9147.
//
// Solidity: function timeLimitForSigningFinalBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForSigningFinalBill() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForSigningFinalBill(&_SCPHA1.CallOpts)
}

// TimeLimitForSigningFinalBill is a free data retrieval call binding the contract method 0xfdcd9147.
//
// Solidity: function timeLimitForSigningFinalBill() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForSigningFinalBill() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForSigningFinalBill(&_SCPHA1.CallOpts)
}

// TimeLimitForSolvingDispute is a free data retrieval call binding the contract method 0x77e711ba.
//
// Solidity: function timeLimitForSolvingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForSolvingDispute(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForSolvingDispute")
	return *ret0, err
}

// TimeLimitForSolvingDispute is a free data retrieval call binding the contract method 0x77e711ba.
//
// Solidity: function timeLimitForSolvingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForSolvingDispute() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForSolvingDispute(&_SCPHA1.CallOpts)
}

// TimeLimitForSolvingDispute is a free data retrieval call binding the contract method 0x77e711ba.
//
// Solidity: function timeLimitForSolvingDispute() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForSolvingDispute() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForSolvingDispute(&_SCPHA1.CallOpts)
}

// TimeLimitForStartingOperation is a free data retrieval call binding the contract method 0xef1b3f78.
//
// Solidity: function timeLimitForStartingOperation() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForStartingOperation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForStartingOperation")
	return *ret0, err
}

// TimeLimitForStartingOperation is a free data retrieval call binding the contract method 0xef1b3f78.
//
// Solidity: function timeLimitForStartingOperation() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForStartingOperation() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForStartingOperation(&_SCPHA1.CallOpts)
}

// TimeLimitForStartingOperation is a free data retrieval call binding the contract method 0xef1b3f78.
//
// Solidity: function timeLimitForStartingOperation() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForStartingOperation() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForStartingOperation(&_SCPHA1.CallOpts)
}

// TimeLimitForVerifyingSignature is a free data retrieval call binding the contract method 0xba7e5638.
//
// Solidity: function timeLimitForVerifyingSignature() constant returns(uint256)
func (_SCPHA1 *SCPHA1Caller) TimeLimitForVerifyingSignature(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCPHA1.contract.Call(opts, out, "timeLimitForVerifyingSignature")
	return *ret0, err
}

// TimeLimitForVerifyingSignature is a free data retrieval call binding the contract method 0xba7e5638.
//
// Solidity: function timeLimitForVerifyingSignature() constant returns(uint256)
func (_SCPHA1 *SCPHA1Session) TimeLimitForVerifyingSignature() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForVerifyingSignature(&_SCPHA1.CallOpts)
}

// TimeLimitForVerifyingSignature is a free data retrieval call binding the contract method 0xba7e5638.
//
// Solidity: function timeLimitForVerifyingSignature() constant returns(uint256)
func (_SCPHA1 *SCPHA1CallerSession) TimeLimitForVerifyingSignature() (*big.Int, error) {
	return _SCPHA1.Contract.TimeLimitForVerifyingSignature(&_SCPHA1.CallOpts)
}

// DischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0x9e331c86.
//
// Solidity: function DischargeAndGenerateFinalCostBill(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _pID, uint256 _finalCost) returns(uint256)
func (_SCPHA1 *SCPHA1Transactor) DischargeAndGenerateFinalCostBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, _estimatedBillID *big.Int, _pID *big.Int, _finalCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "DischargeAndGenerateFinalCostBill", System_Users_Info_address, _estimatedBillID, _pID, _finalCost)
}

// DischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0x9e331c86.
//
// Solidity: function DischargeAndGenerateFinalCostBill(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _pID, uint256 _finalCost) returns(uint256)
func (_SCPHA1 *SCPHA1Session) DischargeAndGenerateFinalCostBill(System_Users_Info_address common.Address, _estimatedBillID *big.Int, _pID *big.Int, _finalCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.DischargeAndGenerateFinalCostBill(&_SCPHA1.TransactOpts, System_Users_Info_address, _estimatedBillID, _pID, _finalCost)
}

// DischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0x9e331c86.
//
// Solidity: function DischargeAndGenerateFinalCostBill(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _pID, uint256 _finalCost) returns(uint256)
func (_SCPHA1 *SCPHA1TransactorSession) DischargeAndGenerateFinalCostBill(System_Users_Info_address common.Address, _estimatedBillID *big.Int, _pID *big.Int, _finalCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.DischargeAndGenerateFinalCostBill(&_SCPHA1.TransactOpts, System_Users_Info_address, _estimatedBillID, _pID, _finalCost)
}

// ConsentFinalBillPatient is a paid mutator transaction binding the contract method 0x5d617c6e.
//
// Solidity: function consentFinalBill_Patient(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1Transactor) ConsentFinalBillPatient(opts *bind.TransactOpts, System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "consentFinalBill_Patient", System_Users_Info_address, _finalBillID, _hID)
}

// ConsentFinalBillPatient is a paid mutator transaction binding the contract method 0x5d617c6e.
//
// Solidity: function consentFinalBill_Patient(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1Session) ConsentFinalBillPatient(System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.ConsentFinalBillPatient(&_SCPHA1.TransactOpts, System_Users_Info_address, _finalBillID, _hID)
}

// ConsentFinalBillPatient is a paid mutator transaction binding the contract method 0x5d617c6e.
//
// Solidity: function consentFinalBill_Patient(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1TransactorSession) ConsentFinalBillPatient(System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.ConsentFinalBillPatient(&_SCPHA1.TransactOpts, System_Users_Info_address, _finalBillID, _hID)
}

// GenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0xfceaffda.
//
// Solidity: function generateEstimatedCostBill(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCost) returns()
func (_SCPHA1 *SCPHA1Transactor) GenerateEstimatedCostBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, _pID *big.Int, _estimatedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "generateEstimatedCostBill", System_Users_Info_address, _pID, _estimatedCost)
}

// GenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0xfceaffda.
//
// Solidity: function generateEstimatedCostBill(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCost) returns()
func (_SCPHA1 *SCPHA1Session) GenerateEstimatedCostBill(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.GenerateEstimatedCostBill(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCost)
}

// GenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0xfceaffda.
//
// Solidity: function generateEstimatedCostBill(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCost) returns()
func (_SCPHA1 *SCPHA1TransactorSession) GenerateEstimatedCostBill(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.GenerateEstimatedCostBill(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCost)
}

// IncreaseTheTimeBy10seconds is a paid mutator transaction binding the contract method 0x118bedad.
//
// Solidity: function increaseTheTimeBy10seconds() returns()
func (_SCPHA1 *SCPHA1Transactor) IncreaseTheTimeBy10seconds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "increaseTheTimeBy10seconds")
}

// IncreaseTheTimeBy10seconds is a paid mutator transaction binding the contract method 0x118bedad.
//
// Solidity: function increaseTheTimeBy10seconds() returns()
func (_SCPHA1 *SCPHA1Session) IncreaseTheTimeBy10seconds() (*types.Transaction, error) {
	return _SCPHA1.Contract.IncreaseTheTimeBy10seconds(&_SCPHA1.TransactOpts)
}

// IncreaseTheTimeBy10seconds is a paid mutator transaction binding the contract method 0x118bedad.
//
// Solidity: function increaseTheTimeBy10seconds() returns()
func (_SCPHA1 *SCPHA1TransactorSession) IncreaseTheTimeBy10seconds() (*types.Transaction, error) {
	return _SCPHA1.Contract.IncreaseTheTimeBy10seconds(&_SCPHA1.TransactOpts)
}

// IncrementAndSetExcessObject is a paid mutator transaction binding the contract method 0x0dcfa2b3.
//
// Solidity: function incrementAndSetExcessObject(SC_P_HA_1Excess excess_object, uint256 _finalId) returns()
func (_SCPHA1 *SCPHA1Transactor) IncrementAndSetExcessObject(opts *bind.TransactOpts, excess_object SC_P_HA_1Excess, _finalId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "incrementAndSetExcessObject", excess_object, _finalId)
}

// IncrementAndSetExcessObject is a paid mutator transaction binding the contract method 0x0dcfa2b3.
//
// Solidity: function incrementAndSetExcessObject(SC_P_HA_1Excess excess_object, uint256 _finalId) returns()
func (_SCPHA1 *SCPHA1Session) IncrementAndSetExcessObject(excess_object SC_P_HA_1Excess, _finalId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.IncrementAndSetExcessObject(&_SCPHA1.TransactOpts, excess_object, _finalId)
}

// IncrementAndSetExcessObject is a paid mutator transaction binding the contract method 0x0dcfa2b3.
//
// Solidity: function incrementAndSetExcessObject(SC_P_HA_1Excess excess_object, uint256 _finalId) returns()
func (_SCPHA1 *SCPHA1TransactorSession) IncrementAndSetExcessObject(excess_object SC_P_HA_1Excess, _finalId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.IncrementAndSetExcessObject(&_SCPHA1.TransactOpts, excess_object, _finalId)
}

// KeepSignedHashToBlockchain is a paid mutator transaction binding the contract method 0xf5fc313e.
//
// Solidity: function keepSignedHashToBlockchain(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID, bytes32 _hashOfData, bytes32 _hashOfEncryptedData, bytes32 _hashOfPatientIdDateHashEncFile, bytes _signOnFileHashByHospital, bytes _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 _noOfInputGates, uint256 _fileChunkSize, uint256 _maxInputLinesToAGate, bytes32 _keyHash) returns()
func (_SCPHA1 *SCPHA1Transactor) KeepSignedHashToBlockchain(opts *bind.TransactOpts, System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int, _hashOfData [32]byte, _hashOfEncryptedData [32]byte, _hashOfPatientIdDateHashEncFile [32]byte, _signOnFileHashByHospital []byte, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile []byte, _noOfInputGates *big.Int, _fileChunkSize *big.Int, _maxInputLinesToAGate *big.Int, _keyHash [32]byte) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "keepSignedHashToBlockchain", System_Users_Info_address, _pID, _estimatedCostBillID, _hashOfData, _hashOfEncryptedData, _hashOfPatientIdDateHashEncFile, _signOnFileHashByHospital, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, _noOfInputGates, _fileChunkSize, _maxInputLinesToAGate, _keyHash)
}

// KeepSignedHashToBlockchain is a paid mutator transaction binding the contract method 0xf5fc313e.
//
// Solidity: function keepSignedHashToBlockchain(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID, bytes32 _hashOfData, bytes32 _hashOfEncryptedData, bytes32 _hashOfPatientIdDateHashEncFile, bytes _signOnFileHashByHospital, bytes _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 _noOfInputGates, uint256 _fileChunkSize, uint256 _maxInputLinesToAGate, bytes32 _keyHash) returns()
func (_SCPHA1 *SCPHA1Session) KeepSignedHashToBlockchain(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int, _hashOfData [32]byte, _hashOfEncryptedData [32]byte, _hashOfPatientIdDateHashEncFile [32]byte, _signOnFileHashByHospital []byte, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile []byte, _noOfInputGates *big.Int, _fileChunkSize *big.Int, _maxInputLinesToAGate *big.Int, _keyHash [32]byte) (*types.Transaction, error) {
	return _SCPHA1.Contract.KeepSignedHashToBlockchain(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCostBillID, _hashOfData, _hashOfEncryptedData, _hashOfPatientIdDateHashEncFile, _signOnFileHashByHospital, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, _noOfInputGates, _fileChunkSize, _maxInputLinesToAGate, _keyHash)
}

// KeepSignedHashToBlockchain is a paid mutator transaction binding the contract method 0xf5fc313e.
//
// Solidity: function keepSignedHashToBlockchain(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID, bytes32 _hashOfData, bytes32 _hashOfEncryptedData, bytes32 _hashOfPatientIdDateHashEncFile, bytes _signOnFileHashByHospital, bytes _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, uint256 _noOfInputGates, uint256 _fileChunkSize, uint256 _maxInputLinesToAGate, bytes32 _keyHash) returns()
func (_SCPHA1 *SCPHA1TransactorSession) KeepSignedHashToBlockchain(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int, _hashOfData [32]byte, _hashOfEncryptedData [32]byte, _hashOfPatientIdDateHashEncFile [32]byte, _signOnFileHashByHospital []byte, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile []byte, _noOfInputGates *big.Int, _fileChunkSize *big.Int, _maxInputLinesToAGate *big.Int, _keyHash [32]byte) (*types.Transaction, error) {
	return _SCPHA1.Contract.KeepSignedHashToBlockchain(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCostBillID, _hashOfData, _hashOfEncryptedData, _hashOfPatientIdDateHashEncFile, _signOnFileHashByHospital, _signOnHashOf_PatientId_DateOfReport_HashOfEncryptedFile, _noOfInputGates, _fileChunkSize, _maxInputLinesToAGate, _keyHash)
}

// KeyReveal is a paid mutator transaction binding the contract method 0x46770a9b.
//
// Solidity: function keyReveal(address System_Users_Info_address, bytes32 _key, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA1 *SCPHA1Transactor) KeyReveal(opts *bind.TransactOpts, System_Users_Info_address common.Address, _key [32]byte, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "keyReveal", System_Users_Info_address, _key, _estimatedBillID, _pId)
}

// KeyReveal is a paid mutator transaction binding the contract method 0x46770a9b.
//
// Solidity: function keyReveal(address System_Users_Info_address, bytes32 _key, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA1 *SCPHA1Session) KeyReveal(System_Users_Info_address common.Address, _key [32]byte, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.KeyReveal(&_SCPHA1.TransactOpts, System_Users_Info_address, _key, _estimatedBillID, _pId)
}

// KeyReveal is a paid mutator transaction binding the contract method 0x46770a9b.
//
// Solidity: function keyReveal(address System_Users_Info_address, bytes32 _key, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA1 *SCPHA1TransactorSession) KeyReveal(System_Users_Info_address common.Address, _key [32]byte, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.KeyReveal(&_SCPHA1.TransactOpts, System_Users_Info_address, _key, _estimatedBillID, _pId)
}

// LockEstimatedAmount is a paid mutator transaction binding the contract method 0xa212af75.
//
// Solidity: function lockEstimatedAmount(address System_Users_Info_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA1 *SCPHA1Transactor) LockEstimatedAmount(opts *bind.TransactOpts, System_Users_Info_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "lockEstimatedAmount", System_Users_Info_address, _hID, _estimatedBillID)
}

// LockEstimatedAmount is a paid mutator transaction binding the contract method 0xa212af75.
//
// Solidity: function lockEstimatedAmount(address System_Users_Info_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA1 *SCPHA1Session) LockEstimatedAmount(System_Users_Info_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.LockEstimatedAmount(&_SCPHA1.TransactOpts, System_Users_Info_address, _hID, _estimatedBillID)
}

// LockEstimatedAmount is a paid mutator transaction binding the contract method 0xa212af75.
//
// Solidity: function lockEstimatedAmount(address System_Users_Info_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA1 *SCPHA1TransactorSession) LockEstimatedAmount(System_Users_Info_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.LockEstimatedAmount(&_SCPHA1.TransactOpts, System_Users_Info_address, _hID, _estimatedBillID)
}

// PatientComplain is a paid mutator transaction binding the contract method 0x29954842.
//
// Solidity: function patient_Complain(address System_Users_Info_address, bytes32[][] complaint, bytes[] encodedVectors_ThisGate, uint256[] _gateIndexes, uint256 _estimatedBillID, uint256 _hID) returns(bool)
func (_SCPHA1 *SCPHA1Transactor) PatientComplain(opts *bind.TransactOpts, System_Users_Info_address common.Address, complaint [][][32]byte, encodedVectors_ThisGate [][]byte, _gateIndexes []*big.Int, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "patient_Complain", System_Users_Info_address, complaint, encodedVectors_ThisGate, _gateIndexes, _estimatedBillID, _hID)
}

// PatientComplain is a paid mutator transaction binding the contract method 0x29954842.
//
// Solidity: function patient_Complain(address System_Users_Info_address, bytes32[][] complaint, bytes[] encodedVectors_ThisGate, uint256[] _gateIndexes, uint256 _estimatedBillID, uint256 _hID) returns(bool)
func (_SCPHA1 *SCPHA1Session) PatientComplain(System_Users_Info_address common.Address, complaint [][][32]byte, encodedVectors_ThisGate [][]byte, _gateIndexes []*big.Int, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.PatientComplain(&_SCPHA1.TransactOpts, System_Users_Info_address, complaint, encodedVectors_ThisGate, _gateIndexes, _estimatedBillID, _hID)
}

// PatientComplain is a paid mutator transaction binding the contract method 0x29954842.
//
// Solidity: function patient_Complain(address System_Users_Info_address, bytes32[][] complaint, bytes[] encodedVectors_ThisGate, uint256[] _gateIndexes, uint256 _estimatedBillID, uint256 _hID) returns(bool)
func (_SCPHA1 *SCPHA1TransactorSession) PatientComplain(System_Users_Info_address common.Address, complaint [][][32]byte, encodedVectors_ThisGate [][]byte, _gateIndexes []*big.Int, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.PatientComplain(&_SCPHA1.TransactOpts, System_Users_Info_address, complaint, encodedVectors_ThisGate, _gateIndexes, _estimatedBillID, _hID)
}

// PatientFinalConsent is a paid mutator transaction binding the contract method 0xfbdfac1c.
//
// Solidity: function patient_final_consent(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1Transactor) PatientFinalConsent(opts *bind.TransactOpts, System_Users_Info_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "patient_final_consent", System_Users_Info_address, _estimatedBillID, _hID)
}

// PatientFinalConsent is a paid mutator transaction binding the contract method 0xfbdfac1c.
//
// Solidity: function patient_final_consent(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1Session) PatientFinalConsent(System_Users_Info_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.PatientFinalConsent(&_SCPHA1.TransactOpts, System_Users_Info_address, _estimatedBillID, _hID)
}

// PatientFinalConsent is a paid mutator transaction binding the contract method 0xfbdfac1c.
//
// Solidity: function patient_final_consent(address System_Users_Info_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA1 *SCPHA1TransactorSession) PatientFinalConsent(System_Users_Info_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.PatientFinalConsent(&_SCPHA1.TransactOpts, System_Users_Info_address, _estimatedBillID, _hID)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0xe0230f03.
//
// Solidity: function raiseDispute(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns(uint256)
func (_SCPHA1 *SCPHA1Transactor) RaiseDispute(opts *bind.TransactOpts, System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "raiseDispute", System_Users_Info_address, _finalBillID, _hID)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0xe0230f03.
//
// Solidity: function raiseDispute(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns(uint256)
func (_SCPHA1 *SCPHA1Session) RaiseDispute(System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.RaiseDispute(&_SCPHA1.TransactOpts, System_Users_Info_address, _finalBillID, _hID)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0xe0230f03.
//
// Solidity: function raiseDispute(address System_Users_Info_address, uint256 _finalBillID, uint256 _hID) returns(uint256)
func (_SCPHA1 *SCPHA1TransactorSession) RaiseDispute(System_Users_Info_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.RaiseDispute(&_SCPHA1.TransactOpts, System_Users_Info_address, _finalBillID, _hID)
}

// SetDisputeStruct is a paid mutator transaction binding the contract method 0x0cde8571.
//
// Solidity: function setDisputeStruct(SC_P_HA_1Dispute dispute_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Transactor) SetDisputeStruct(opts *bind.TransactOpts, dispute_object SC_P_HA_1Dispute, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "setDisputeStruct", dispute_object, id)
}

// SetDisputeStruct is a paid mutator transaction binding the contract method 0x0cde8571.
//
// Solidity: function setDisputeStruct(SC_P_HA_1Dispute dispute_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Session) SetDisputeStruct(dispute_object SC_P_HA_1Dispute, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetDisputeStruct(&_SCPHA1.TransactOpts, dispute_object, id)
}

// SetDisputeStruct is a paid mutator transaction binding the contract method 0x0cde8571.
//
// Solidity: function setDisputeStruct(SC_P_HA_1Dispute dispute_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1TransactorSession) SetDisputeStruct(dispute_object SC_P_HA_1Dispute, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetDisputeStruct(&_SCPHA1.TransactOpts, dispute_object, id)
}

// SetEstimatedCheckUpCostStruct is a paid mutator transaction binding the contract method 0xc6ede1c2.
//
// Solidity: function setEstimatedCheckUpCostStruct(SC_P_HA_1EstimatedCheckUpCost estimatedCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Transactor) SetEstimatedCheckUpCostStruct(opts *bind.TransactOpts, estimatedCheckUpCost_Object SC_P_HA_1EstimatedCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "setEstimatedCheckUpCostStruct", estimatedCheckUpCost_Object, id)
}

// SetEstimatedCheckUpCostStruct is a paid mutator transaction binding the contract method 0xc6ede1c2.
//
// Solidity: function setEstimatedCheckUpCostStruct(SC_P_HA_1EstimatedCheckUpCost estimatedCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Session) SetEstimatedCheckUpCostStruct(estimatedCheckUpCost_Object SC_P_HA_1EstimatedCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetEstimatedCheckUpCostStruct(&_SCPHA1.TransactOpts, estimatedCheckUpCost_Object, id)
}

// SetEstimatedCheckUpCostStruct is a paid mutator transaction binding the contract method 0xc6ede1c2.
//
// Solidity: function setEstimatedCheckUpCostStruct(SC_P_HA_1EstimatedCheckUpCost estimatedCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1TransactorSession) SetEstimatedCheckUpCostStruct(estimatedCheckUpCost_Object SC_P_HA_1EstimatedCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetEstimatedCheckUpCostStruct(&_SCPHA1.TransactOpts, estimatedCheckUpCost_Object, id)
}

// SetExcessObject is a paid mutator transaction binding the contract method 0x4fbca6f2.
//
// Solidity: function setExcessObject(SC_P_HA_1Excess excess_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Transactor) SetExcessObject(opts *bind.TransactOpts, excess_object SC_P_HA_1Excess, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "setExcessObject", excess_object, id)
}

// SetExcessObject is a paid mutator transaction binding the contract method 0x4fbca6f2.
//
// Solidity: function setExcessObject(SC_P_HA_1Excess excess_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Session) SetExcessObject(excess_object SC_P_HA_1Excess, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetExcessObject(&_SCPHA1.TransactOpts, excess_object, id)
}

// SetExcessObject is a paid mutator transaction binding the contract method 0x4fbca6f2.
//
// Solidity: function setExcessObject(SC_P_HA_1Excess excess_object, uint256 id) returns()
func (_SCPHA1 *SCPHA1TransactorSession) SetExcessObject(excess_object SC_P_HA_1Excess, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetExcessObject(&_SCPHA1.TransactOpts, excess_object, id)
}

// SetFinalCheckUpCost is a paid mutator transaction binding the contract method 0x066997c3.
//
// Solidity: function setFinalCheckUpCost(SC_P_HA_1FinalCheckUpCost finalCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Transactor) SetFinalCheckUpCost(opts *bind.TransactOpts, finalCheckUpCost_Object SC_P_HA_1FinalCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "setFinalCheckUpCost", finalCheckUpCost_Object, id)
}

// SetFinalCheckUpCost is a paid mutator transaction binding the contract method 0x066997c3.
//
// Solidity: function setFinalCheckUpCost(SC_P_HA_1FinalCheckUpCost finalCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Session) SetFinalCheckUpCost(finalCheckUpCost_Object SC_P_HA_1FinalCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetFinalCheckUpCost(&_SCPHA1.TransactOpts, finalCheckUpCost_Object, id)
}

// SetFinalCheckUpCost is a paid mutator transaction binding the contract method 0x066997c3.
//
// Solidity: function setFinalCheckUpCost(SC_P_HA_1FinalCheckUpCost finalCheckUpCost_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1TransactorSession) SetFinalCheckUpCost(finalCheckUpCost_Object SC_P_HA_1FinalCheckUpCost, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetFinalCheckUpCost(&_SCPHA1.TransactOpts, finalCheckUpCost_Object, id)
}

// SetMultiSigOnMedicalData is a paid mutator transaction binding the contract method 0xd50f1fb6.
//
// Solidity: function setMultiSigOnMedicalData(SC_P_HA_1MultiSigOnMedicalData multiSigOnMedicalData_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Transactor) SetMultiSigOnMedicalData(opts *bind.TransactOpts, multiSigOnMedicalData_Object SC_P_HA_1MultiSigOnMedicalData, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "setMultiSigOnMedicalData", multiSigOnMedicalData_Object, id)
}

// SetMultiSigOnMedicalData is a paid mutator transaction binding the contract method 0xd50f1fb6.
//
// Solidity: function setMultiSigOnMedicalData(SC_P_HA_1MultiSigOnMedicalData multiSigOnMedicalData_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1Session) SetMultiSigOnMedicalData(multiSigOnMedicalData_Object SC_P_HA_1MultiSigOnMedicalData, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetMultiSigOnMedicalData(&_SCPHA1.TransactOpts, multiSigOnMedicalData_Object, id)
}

// SetMultiSigOnMedicalData is a paid mutator transaction binding the contract method 0xd50f1fb6.
//
// Solidity: function setMultiSigOnMedicalData(SC_P_HA_1MultiSigOnMedicalData multiSigOnMedicalData_Object, uint256 id) returns()
func (_SCPHA1 *SCPHA1TransactorSession) SetMultiSigOnMedicalData(multiSigOnMedicalData_Object SC_P_HA_1MultiSigOnMedicalData, id *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.SetMultiSigOnMedicalData(&_SCPHA1.TransactOpts, multiSigOnMedicalData_Object, id)
}

// StartTreatment is a paid mutator transaction binding the contract method 0x8cef6685.
//
// Solidity: function startTreatment(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID) returns()
func (_SCPHA1 *SCPHA1Transactor) StartTreatment(opts *bind.TransactOpts, System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "startTreatment", System_Users_Info_address, _pID, _estimatedCostBillID)
}

// StartTreatment is a paid mutator transaction binding the contract method 0x8cef6685.
//
// Solidity: function startTreatment(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID) returns()
func (_SCPHA1 *SCPHA1Session) StartTreatment(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.StartTreatment(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCostBillID)
}

// StartTreatment is a paid mutator transaction binding the contract method 0x8cef6685.
//
// Solidity: function startTreatment(address System_Users_Info_address, uint256 _pID, uint256 _estimatedCostBillID) returns()
func (_SCPHA1 *SCPHA1TransactorSession) StartTreatment(System_Users_Info_address common.Address, _pID *big.Int, _estimatedCostBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.StartTreatment(&_SCPHA1.TransactOpts, System_Users_Info_address, _pID, _estimatedCostBillID)
}

// TransferEtherFromThisContract is a paid mutator transaction binding the contract method 0x2c7ad37c.
//
// Solidity: function transferEtherFromThisContract(address _addr, uint256 _amount) returns()
func (_SCPHA1 *SCPHA1Transactor) TransferEtherFromThisContract(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "transferEtherFromThisContract", _addr, _amount)
}

// TransferEtherFromThisContract is a paid mutator transaction binding the contract method 0x2c7ad37c.
//
// Solidity: function transferEtherFromThisContract(address _addr, uint256 _amount) returns()
func (_SCPHA1 *SCPHA1Session) TransferEtherFromThisContract(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.TransferEtherFromThisContract(&_SCPHA1.TransactOpts, _addr, _amount)
}

// TransferEtherFromThisContract is a paid mutator transaction binding the contract method 0x2c7ad37c.
//
// Solidity: function transferEtherFromThisContract(address _addr, uint256 _amount) returns()
func (_SCPHA1 *SCPHA1TransactorSession) TransferEtherFromThisContract(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SCPHA1.Contract.TransferEtherFromThisContract(&_SCPHA1.TransactOpts, _addr, _amount)
}

// VerifyAndGiveConsent is a paid mutator transaction binding the contract method 0xbb6bbf35.
//
// Solidity: function verifyAndGiveConsent(address System_Users_Info_address, uint256 _multiSigID, uint256 _hID, bytes32 _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, bytes32 _hashOfEncryptedData) returns()
func (_SCPHA1 *SCPHA1Transactor) VerifyAndGiveConsent(opts *bind.TransactOpts, System_Users_Info_address common.Address, _multiSigID *big.Int, _hID *big.Int, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile [32]byte, _hashOfEncryptedData [32]byte) (*types.Transaction, error) {
	return _SCPHA1.contract.Transact(opts, "verifyAndGiveConsent", System_Users_Info_address, _multiSigID, _hID, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, _hashOfEncryptedData)
}

// VerifyAndGiveConsent is a paid mutator transaction binding the contract method 0xbb6bbf35.
//
// Solidity: function verifyAndGiveConsent(address System_Users_Info_address, uint256 _multiSigID, uint256 _hID, bytes32 _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, bytes32 _hashOfEncryptedData) returns()
func (_SCPHA1 *SCPHA1Session) VerifyAndGiveConsent(System_Users_Info_address common.Address, _multiSigID *big.Int, _hID *big.Int, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile [32]byte, _hashOfEncryptedData [32]byte) (*types.Transaction, error) {
	return _SCPHA1.Contract.VerifyAndGiveConsent(&_SCPHA1.TransactOpts, System_Users_Info_address, _multiSigID, _hID, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, _hashOfEncryptedData)
}

// VerifyAndGiveConsent is a paid mutator transaction binding the contract method 0xbb6bbf35.
//
// Solidity: function verifyAndGiveConsent(address System_Users_Info_address, uint256 _multiSigID, uint256 _hID, bytes32 _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, bytes32 _hashOfEncryptedData) returns()
func (_SCPHA1 *SCPHA1TransactorSession) VerifyAndGiveConsent(System_Users_Info_address common.Address, _multiSigID *big.Int, _hID *big.Int, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile [32]byte, _hashOfEncryptedData [32]byte) (*types.Transaction, error) {
	return _SCPHA1.Contract.VerifyAndGiveConsent(&_SCPHA1.TransactOpts, System_Users_Info_address, _multiSigID, _hID, _HashOf_PatientId_DateOfReport_HashOfEncryptedFile, _hashOfEncryptedData)
}

// SCPHA2ABI is the input ABI used to generate the binding from.
const SCPHA2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"consentFinalBill_AfterDisputeRaise_Patient\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"hospital_leave_DischargeAndGenerateFinalCostBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"hospital_leave_generateEstimatedCostBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"hospital_leave_keepHashToBlockchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pId\",\"type\":\"uint256\"}],\"name\":\"hospital_leave_keyReveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"patient_leave_consentFinalBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"patient_leave_keepHashToBlockchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"patient_leave_lockEstimatedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"}],\"name\":\"patient_leave_verifyAndGiveConsent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"P_HA_1_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_estimatedBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finalBillID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_modifiedCost\",\"type\":\"uint256\"}],\"name\":\"solveDisputeAndModifiedFinalBill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SCPHA2FuncSigs maps the 4-byte function signature to its string representation.
var SCPHA2FuncSigs = map[string]string{
	"9c61ab02": "consentFinalBill_AfterDisputeRaise_Patient(address,address,uint256,uint256)",
	"f86e8d5c": "hospital_leave_DischargeAndGenerateFinalCostBill(address,address,uint256,uint256)",
	"32707436": "hospital_leave_generateEstimatedCostBill(address,address,uint256,uint256)",
	"9c8662f5": "hospital_leave_keepHashToBlockchain(address,address,uint256,uint256)",
	"fdd54eca": "hospital_leave_keyReveal(address,address,uint256,uint256)",
	"1c3f24b5": "patient_leave_consentFinalBill(address,address,uint256,uint256)",
	"232b4868": "patient_leave_keepHashToBlockchain(address,address,uint256,uint256)",
	"2b5294db": "patient_leave_lockEstimatedAmount(address,address,uint256,uint256)",
	"05c29d34": "patient_leave_verifyAndGiveConsent(address,address,uint256,uint256)",
	"9057947b": "solveDisputeAndModifiedFinalBill(address,address,uint256,uint256,uint256,uint256,uint256)",
}

// SCPHA2Bin is the compiled bytecode used for deploying new contracts.
var SCPHA2Bin = "0x608060405234801561001057600080fd5b50613431806100206000396000f3fe6080604052600436106100915760003560e01c80639057947b116100595780639057947b146101385780639c61ab02146101585780639c8662f51461016b578063f86e8d5c1461018b578063fdd54eca146101ab57610091565b806305c29d34146100965780631c3f24b5146100b8578063232b4868146100d85780632b5294db146100f85780633270743614610118575b600080fd5b3480156100a257600080fd5b506100b66100b1366004612cf0565b6101cb565b005b3480156100c457600080fd5b506100b66100d3366004612cf0565b610622565b3480156100e457600080fd5b506100b66100f3366004612cf0565b610a37565b34801561010457600080fd5b506100b6610113366004612cf0565b610e14565b34801561012457600080fd5b506100b6610133366004612cf0565b6110df565b34801561014457600080fd5b506100b6610153366004612d35565b611308565b6100b6610166366004612cf0565b611b5b565b34801561017757600080fd5b506100b6610186366004612cf0565b611f87565b34801561019757600080fd5b506100b66101a6366004612cf0565b6122c0565b3480156101b757600080fd5b506100b66101c6366004612cf0565b612607565b604051639836b82560e01b815260009081908190879087906001600160a01b03831690639836b8259061020290339060040161310a565b60006040518083038186803b15801561021a57600080fd5b505afa15801561022e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102569190810190612c51565b505060405163b90d423560e01b8152929850506001600160a01b038416925063b90d42359161028a91508990600401613383565b60206040518083038186803b1580156102a257600080fd5b505afa1580156102b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102da91908101906130c0565b604051638e2d553160e01b81529094506001600160a01b03821690638e2d553190610309908990600401613383565b60206040518083038186803b15801561032157600080fd5b505afa158015610335573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061035991908101906130c0565b92506103636129cd565b60405163413a9b8560e11b81526001600160a01b03831690638275370a9061038f908890600401613383565b60006040518083038186803b1580156103a757600080fd5b505afa1580156103bb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103e39190810190612fb1565b90506103ed612a40565b6040516303544ff560e41b81526001600160a01b03841690633544ff5090610419908b90600401613383565b6101406040518083038186803b15801561043257600080fd5b505afa158015610446573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061046a9190810190612e00565b90508882604001511461047c57600080fd5b8682602001511461048c57600080fd5b610120820151158015906104a4575042826101200151105b6104ad57600080fd5b610160820151156104bd57600080fd5b84156104c857600080fd5b826001600160a01b031663672e7b216040518163ffffffff1660e01b815260040160206040518083038186803b15801561050157600080fd5b505afa158015610515573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061053991908101906130c0565b82610120015142031161054b57600080fd5b426101608301526060810151604051630b1eb4df60e21b81526001600160a01b03851691632c7ad37c9161058391339160040161311e565b600060405180830381600087803b15801561059d57600080fd5b505af11580156105b1573d6000803e3d6000fd5b5050604051636a878fdb60e11b81526001600160a01b038616925063d50f1fb691506105e39085908a906004016132b4565b600060405180830381600087803b1580156105fd57600080fd5b505af1158015610611573d6000803e3d6000fd5b505050505050505050505050505050565b604051639836b82560e01b815260009081908190879087906001600160a01b03831690639836b8259061065990339060040161310a565b60006040518083038186803b15801561067157600080fd5b505afa158015610685573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106ad9190810190612c51565b5050604051638e2d553160e01b8152929850506001600160a01b0384169250638e2d5531916106e191508a90600401613383565b60206040518083038186803b1580156106f957600080fd5b505afa15801561070d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061073191908101906130c0565b60405163b90d423560e01b81529094506001600160a01b0382169063b90d423590610760908a90600401613383565b60206040518083038186803b15801561077857600080fd5b505afa15801561078c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107b091908101906130c0565b92506107ba612a93565b6040516307b6f84160e11b81526001600160a01b03831690630f6df082906107e6908890600401613383565b6101606040518083038186803b1580156107ff57600080fd5b505afa158015610813573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108379190810190612edc565b9050610841612aed565b6040516313d6938760e21b81526001600160a01b03841690634f5a4e1c9061086d908890600401613383565b60606040518083038186803b15801561088557600080fd5b505afa158015610899573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108bd9190810190612f74565b9050878260400151146108cf57600080fd5b868260200151146108df57600080fd5b60c0820151158015906108f55750428260c00151105b6108fe57600080fd5b6101408201511561090e57600080fd5b60408101511561091d57600080fd5b826001600160a01b031663b17c31fb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561095657600080fd5b505afa15801561096a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061098e91908101906130c0565b8260c0015142031161099f57600080fd5b426101408301526060820151604051630b1eb4df60e21b81526001600160a01b03851691632c7ad37c916109d791339160040161311e565b600060405180830381600087803b1580156109f157600080fd5b505af1158015610a05573d6000803e3d6000fd5b505060405163066997c360e01b81526001600160a01b038616925063066997c391506105e39085908a90600401613239565b604051639836b82560e01b81526000908190869086906001600160a01b03831690639836b82590610a6c90339060040161310a565b60006040518083038186803b158015610a8457600080fd5b505afa158015610a98573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ac09190810190612c51565b50929750610ad29350612a4092505050565b6040516303544ff560e41b81526001600160a01b03831690633544ff5090610afe908990600401613383565b6101406040518083038186803b158015610b1757600080fd5b505afa158015610b2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b4f9190810190612e00565b60405163b90d423560e01b81529091506001600160a01b0383169063b90d423590610b7e908990600401613383565b60206040518083038186803b158015610b9657600080fd5b505afa158015610baa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610bce91908101906130c0565b9350610bd86129cd565b60405163413a9b8560e11b81526001600160a01b03841690638275370a90610c04908890600401613383565b60006040518083038186803b158015610c1c57600080fd5b505afa158015610c30573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c589190810190612fb1565b905087816040015114610c6a57600080fd5b85816020015114610c7a57600080fd5b61010081015115801590610c92575042816101000151105b610c9b57600080fd5b61016081015115610cab57600080fd5b61012081015115610cbb57600080fd5b826001600160a01b031663ba7e56386040518163ffffffff1660e01b815260040160206040518083038186803b158015610cf457600080fd5b505afa158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610d2c91908101906130c0565b816101000151420310610d3e57600080fd5b426101608201526060820151604051630b1eb4df60e21b81526001600160a01b03851691632c7ad37c91610d7691339160040161311e565b600060405180830381600087803b158015610d9057600080fd5b505af1158015610da4573d6000803e3d6000fd5b5050604051636a878fdb60e11b81526001600160a01b038616925063d50f1fb69150610dd690849089906004016132b4565b600060405180830381600087803b158015610df057600080fd5b505af1158015610e04573d6000803e3d6000fd5b5050505050505050505050505050565b604051639836b82560e01b8152600090859085906001600160a01b03831690639836b82590610e4790339060040161310a565b60006040518083038186803b158015610e5f57600080fd5b505afa158015610e73573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e9b9190810190612c51565b50929650610ead9350612a4092505050565b6040516303544ff560e41b81526001600160a01b03831690633544ff5090610ed9908890600401613383565b6101406040518083038186803b158015610ef257600080fd5b505afa158015610f06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f2a9190810190612e00565b80519091508514610f3a57600080fd5b85816040015114610f4a57600080fd5b83816020015114610f5a57600080fd5b60e081015115610f6957600080fd5b61010081015115610f7957600080fd5b61012081015115610f8957600080fd5b816001600160a01b031663ef1b3f786040518163ffffffff1660e01b815260040160206040518083038186803b158015610fc257600080fd5b505afa158015610fd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610ffa91908101906130c0565b8160c0015142031161100b57600080fd5b426101008201526060810151604051630b1eb4df60e21b81526001600160a01b03841691632c7ad37c9161104391339160040161311e565b600060405180830381600087803b15801561105d57600080fd5b505af1158015611071573d6000803e3d6000fd5b5050604051636376f0e160e11b81526001600160a01b038516925063c6ede1c291506110a3908490899060040161318f565b600060405180830381600087803b1580156110bd57600080fd5b505af11580156110d1573d6000803e3d6000fd5b505050505050505050505050565b604051621264bd60e81b8152600090859085906001600160a01b03831690631264bd009061111190339060040161310a565b60006040518083038186803b15801561112957600080fd5b505afa15801561113d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111659190810190612bf9565b5093506111729050612a40565b6040516303544ff560e41b81526001600160a01b03831690633544ff509061119e908890600401613383565b6101406040518083038186803b1580156111b757600080fd5b505afa1580156111cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506111ef9190810190612e00565b805190915085146111ff57600080fd5b8381604001511461120f57600080fd5b8581602001511461121f57600080fd5b60c08101511561122e57600080fd5b816001600160a01b03166325e3bb736040518163ffffffff1660e01b815260040160206040518083038186803b15801561126757600080fd5b505afa15801561127b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061129f91908101906130c0565b8160a001514203116112b057600080fd5b610100810151156112c057600080fd5b610120810151156112d057600080fd5b426101208201526060810151604051630b1eb4df60e21b81526001600160a01b03841691632c7ad37c9161104391339160040161311e565b604051621264bd60e81b81526000908190899089906001600160a01b03831690631264bd009061133c90339060040161310a565b60006040518083038186803b15801561135457600080fd5b505afa158015611368573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113909190810190612bf9565b506040516304a4a6a360e31b8152909550600091506001600160a01b038416906325253518906113c4908a90600401613383565b60006040518083038186803b1580156113dc57600080fd5b505afa1580156113f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114189190810190612c51565b5093945061142a9350612b0d92505050565b6040516301a2a65760e01b81526001600160a01b038416906301a2a65790611456908e90600401613383565b60e06040518083038186803b15801561146e57600080fd5b505afa158015611482573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506114a69190810190612d96565b90506114b0612a93565b6040516307b6f84160e11b81526001600160a01b03851690630f6df082906114dc908d90600401613383565b6101606040518083038186803b1580156114f557600080fd5b505afa158015611509573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061152d9190810190612edc565b9050611537612b4c565b82518d1461154457600080fd5b8783604001511461155457600080fd5b8983602001511461156457600080fd5b846001600160a01b03166377e711ba6040518163ffffffff1660e01b815260040160206040518083038186803b15801561159d57600080fd5b505afa1580156115b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506115d591908101906130c0565b8360800151420311156115e757600080fd5b8a8360600151146115f757600080fd5b81518b1461160457600080fd5b8982602001511461161457600080fd5b8782604001511461162457600080fd5b826080015182608001511061163857600080fd5b816060015189111561164957600080fd5b8160600151891015611a6b5761165d612a40565b856001600160a01b0316633544ff508e6040518263ffffffff1660e01b81526004016116899190613383565b6101406040518083038186803b1580156116a257600080fd5b505afa1580156116b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506116da9190810190612e00565b80519091508d146116ea57600080fd5b8a8160200151146116fa57600080fd5b8881604001511461170a57600080fd5b600160c085015260a0830151156119125760a0830151604051636b83401d60e01b81529098506001600160a01b03871690636b83401d9061174f908b90600401613383565b60a06040518083038186803b15801561176757600080fd5b505afa15801561177b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061179f9190810190612e8b565b60a08401518151919350146117b357600080fd5b8a8260200151146117c357600080fd5b888260400151146117d357600080fd5b6060808201518b900383820152830151604051630b1eb4df60e21b81526001600160a01b03881691632c7ad37c91611814913391908f90039060040161311e565b600060405180830381600087803b15801561182e57600080fd5b505af1158015611842573d6000803e3d6000fd5b5050506060840151604051630b1eb4df60e21b81526001600160a01b0389169250632c7ad37c9161187b9189918f90039060040161311e565b600060405180830381600087803b15801561189557600080fd5b505af11580156118a9573d6000803e3d6000fd5b50506040516327de537960e11b81526001600160a01b0389169250634fbca6f291506118db9085908c906004016131ff565b600060405180830381600087803b1580156118f557600080fd5b505af1158015611909573d6000803e3d6000fd5b50505050611a65565b8a8260200181815250508882604001818152505089816060015103826060018181525050856001600160a01b0316630dcfa2b3838e6040518363ffffffff1660e01b81526004016119649291906131ff565b600060405180830381600087803b15801561197e57600080fd5b505af1158015611992573d6000803e3d6000fd5b5050506060840151604051630b1eb4df60e21b81526001600160a01b0389169250632c7ad37c916119cb9133918f90039060040161311e565b600060405180830381600087803b1580156119e557600080fd5b505af11580156119f9573d6000803e3d6000fd5b5050506060840151604051630b1eb4df60e21b81526001600160a01b0389169250632c7ad37c91611a329189918f90039060040161311e565b600060405180830381600087803b158015611a4c57600080fd5b505af1158015611a60573d6000803e3d6000fd5b505050505b50611a73565b600060c08401525b428360a0018181525050846001600160a01b0316630cde8571848f6040518363ffffffff1660e01b8152600401611aab929190613137565b600060405180830381600087803b158015611ac557600080fd5b505af1158015611ad9573d6000803e3d6000fd5b505050506060820189905242608083015260405163066997c360e01b81526001600160a01b0386169063066997c390611b189085908f90600401613239565b600060405180830381600087803b158015611b3257600080fd5b505af1158015611b46573d6000803e3d6000fd5b50505050505050505050505050505050505050565b604051639836b82560e01b8152600090859085906001600160a01b03831690639836b82590611b8e90339060040161310a565b60006040518083038186803b158015611ba657600080fd5b505afa158015611bba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611be29190810190612c51565b505060405163fd4cc07960e01b815292965060009350506001600160a01b0385169163fd4cc0799150611c19908890600401613383565b60006040518083038186803b158015611c3157600080fd5b505afa158015611c45573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c6d9190810190612bf9565b50909150611c7b9050612a93565b6040516307b6f84160e11b81526001600160a01b03841690630f6df08290611ca7908a90600401613383565b6101606040518083038186803b158015611cc057600080fd5b505afa158015611cd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611cf89190810190612edc565b80519091508714611d0857600080fd5b84816020015114611d1857600080fd5b85816040015114611d2857600080fd5b42816080015110611d3857600080fd5b60c081015115611d4757600080fd5b604051630d84410d60e01b81526000906001600160a01b03851690630d84410d90611d76908b90600401613383565b60206040518083038186803b158015611d8e57600080fd5b505afa158015611da2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611dc691908101906130c0565b905080611dd257600080fd5b611dda612b0d565b6040516301a2a65760e01b81526001600160a01b038616906301a2a65790611e06908590600401613383565b60e06040518083038186803b158015611e1e57600080fd5b505afa158015611e32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611e569190810190612d96565b905060008160a0015111611e6957600080fd5b846001600160a01b031663fdcd91476040518163ffffffff1660e01b815260040160206040518083038186803b158015611ea257600080fd5b505afa158015611eb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611eda91908101906130c0565b8160a0015142031115611eec57600080fd5b846001600160a01b0316632c7ad37c8585606001516002026040518363ffffffff1660e01b8152600401611f2192919061311e565b600060405180830381600087803b158015611f3b57600080fd5b505af1158015611f4f573d6000803e3d6000fd5b50504260c0860152505060405163066997c360e01b81526001600160a01b0386169063066997c3906105e39086908d90600401613239565b604051621264bd60e81b81526000908190869086906001600160a01b03831690631264bd0090611fbb90339060040161310a565b60006040518083038186803b158015611fd357600080fd5b505afa158015611fe7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261200f9190810190612bf9565b50945061201c9050612a40565b6040516303544ff560e41b81526001600160a01b03831690633544ff5090612048908990600401613383565b6101406040518083038186803b15801561206157600080fd5b505afa158015612075573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506120999190810190612e00565b60405163b90d423560e01b81529091506001600160a01b0383169063b90d4235906120c8908990600401613383565b60206040518083038186803b1580156120e057600080fd5b505afa1580156120f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061211891908101906130c0565b93506121226129cd565b60405163413a9b8560e11b81526001600160a01b03841690638275370a9061214e908890600401613383565b60006040518083038186803b15801561216657600080fd5b505afa15801561217a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121a29190810190612fb1565b9050858160400151146121b457600080fd5b878160200151146121c457600080fd5b610100810151158015906121dc575042816101000151105b6121e557600080fd5b610180810151156121f557600080fd5b6101208101511561220557600080fd5b826001600160a01b031663ba7e56386040518163ffffffff1660e01b815260040160206040518083038186803b15801561223e57600080fd5b505afa158015612252573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061227691908101906130c0565b81610100015142031161228857600080fd5b426101808201526060820151604051630b1eb4df60e21b81526001600160a01b03851691632c7ad37c91610d7691339160040161311e565b604051621264bd60e81b81526000908190869086906001600160a01b03831690631264bd00906122f490339060040161310a565b60006040518083038186803b15801561230c57600080fd5b505afa158015612320573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526123489190810190612bf9565b50604051638e2d553160e01b81529095506001600160a01b0383169150638e2d553190612379908990600401613383565b60206040518083038186803b15801561239157600080fd5b505afa1580156123a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506123c991908101906130c0565b92506123d3612a93565b6040516307b6f84160e11b81526001600160a01b03831690630f6df082906123ff908790600401613383565b6101606040518083038186803b15801561241857600080fd5b505afa15801561242c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506124509190810190612edc565b90508481604001511461246257600080fd5b8581602001511461247257600080fd5b6101208101511561248257600080fd5b6080810151158015906124985750428160800151105b6124a157600080fd5b60c0810151156124b057600080fd5b816001600160a01b031663fdcd91476040518163ffffffff1660e01b815260040160206040518083038186803b1580156124e957600080fd5b505afa1580156124fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061252191908101906130c0565b816080015142031161253257600080fd5b426101208201526060810151604051630b1eb4df60e21b81526001600160a01b03841691632c7ad37c9161256a91339160040161311e565b600060405180830381600087803b15801561258457600080fd5b505af1158015612598573d6000803e3d6000fd5b505060405163066997c360e01b81526001600160a01b038516925063066997c391506125ca9084908890600401613239565b600060405180830381600087803b1580156125e457600080fd5b505af11580156125f8573d6000803e3d6000fd5b50505050505050505050505050565b604051621264bd60e81b815260009081908190879087906001600160a01b03831690631264bd009061263d90339060040161310a565b60006040518083038186803b15801561265557600080fd5b505afa158015612669573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526126919190810190612bf9565b50604051638e2d553160e01b81529096506001600160a01b0383169150638e2d5531906126c2908a90600401613383565b60206040518083038186803b1580156126da57600080fd5b505afa1580156126ee573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061271291908101906130c0565b60405163b90d423560e01b81529094506001600160a01b0382169063b90d423590612741908a90600401613383565b60206040518083038186803b15801561275957600080fd5b505afa15801561276d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061279191908101906130c0565b925061279b612a93565b6040516307b6f84160e11b81526001600160a01b03831690630f6df082906127c7908890600401613383565b6101606040518083038186803b1580156127e057600080fd5b505afa1580156127f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506128189190810190612edc565b9050612822612aed565b6040516313d6938760e21b81526001600160a01b03841690634f5a4e1c9061284e908890600401613383565b60606040518083038186803b15801561286657600080fd5b505afa15801561287a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061289e9190810190612f74565b9050868260400151146128b057600080fd5b878260200151146128c057600080fd5b610120820151156128d057600080fd5b6040810151158015906128e65750428160400151105b6128ef57600080fd5b60e0820151156128fe57600080fd5b6101008201511561290e57600080fd5b826001600160a01b0316638513f5e56040518163ffffffff1660e01b815260040160206040518083038186803b15801561294757600080fd5b505afa15801561295b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061297f91908101906130c0565b816040015142031161299057600080fd5b4260e08301526060820151604051630b1eb4df60e21b81526001600160a01b03851691632c7ad37c916109d791339160029091029060040161311e565b604051806101a001604052806000815260200160008152602001600081526020016000801916815260200160008019168152602001600080191681526020016060815260200160608152602001600081526020016000815260200160001515815260200160008152602001600081525090565b604051806101400160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b60405180610160016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516060810182526000808252602082018190529181019190915290565b6040518060e001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b80518015158114612b8b57600080fd5b92915050565b600082601f830112612ba1578081fd5b815167ffffffffffffffff811115612bb7578182fd5b612bca601f8201601f191660200161338c565b9150808252836020828501011115612be157600080fd5b612bf28160208401602086016133b3565b5092915050565b600080600060608486031215612c0d578283fd5b8351612c18816133e3565b60208501516040860151919450925067ffffffffffffffff811115612c3b578182fd5b612c4786828701612b91565b9150509250925092565b60008060008060008060c08789031215612c69578384fd5b8651612c74816133e3565b60208801516040890151919750955067ffffffffffffffff80821115612c98578586fd5b612ca48a838b01612b91565b95506060890151915060ff82168214612cbb578384fd5b608089015160a08a0151929550935080821115612cd6578283fd5b50612ce389828a01612b91565b9150509295509295509295565b60008060008060808587031215612d05578182fd5b8435612d10816133e3565b93506020850135612d20816133e3565b93969395505050506040820135916060013590565b600080600080600080600060e0888a031215612d4f578485fd5b8735612d5a816133e3565b96506020880135612d6a816133e3565b96999698505050506040850135946060810135946080820135945060a0820135935060c0909101359150565b600060e08284031215612da7578081fd5b612db160e061338c565b825181526020830151602082015260408301516040820152606083015160608201526080830151608082015260a083015160a0820152612df48460c08501612b7b565b60c08201529392505050565b6000610140808385031215612e13578182fd5b612e1c8161338c565b835181526020840151602082015260408401516040820152606084015160608201526080840151608082015260a084015160a082015260c084015160c082015260e084015160e08201526101009150818401518282015261012091508184015182820152809250505092915050565b600060a08284031215612e9c578081fd5b612ea660a061338c565b82518152602083015160208201526040830151604082015260608301516060820152608083015160808201528091505092915050565b6000610160808385031215612eef578182fd5b612ef88161338c565b835181526020840151602082015260408401516040820152606084015160608201526080840151608082015260a084015160a082015260c084015160c082015260e084015160e0820152610100915081840151828201526101209150818401518282015261014091508184015182820152809250505092915050565b600060608284031215612f85578081fd5b612f8f606061338c565b8251815260208301516020820152604083015160408201528091505092915050565b600060208284031215612fc2578081fd5b815167ffffffffffffffff80821115612fd9578283fd5b6101a0918401808603831315612fed578384fd5b612ff68361338c565b815181526020820151602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015193508284111561303f578485fd5b61304b87858401612b91565b60c082015260e0820151935082841115613063578485fd5b61306f87858401612b91565b60e082015261010082810151908201526101208083015190820152610140935061309b87858401612b7b565b9381019390935261016081810151908401526101809081015190830152509392505050565b6000602082840312156130d1578081fd5b5051919050565b15159052565b600081518084526130f68160208601602086016133b3565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b600061010082019050835182526020840151602083015260408401516040830152606084015160608301526080840151608083015260a084015160a083015260c0840151151560c08301528260e08301529392505050565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c0808401519082015260e08084015190820152610100808401519082015261012092830151928101929092526101408201526101600190565b825181526020808401519082015260408084015190820152606080840151908201526080928301519281019290925260a082015260c00190565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c0808401519082015260e080840151908201526101008084015190820152610120808401519082015261014092830151928101929092526101608201526101800190565b600060408252835160408301526020840151606083015260408401516080830152606084015160a0830152608084015160c083015260a084015160e083015260c08401516101a061010081818601526133116101e08601846130de565b60e08801519350610120603f19878303018188015261333082866130de565b9289015161014088810191909152818a0151610160808a0191909152908a0151955061018093909250613365848901876130d8565b91890151938701939093529601516101c08501525050506020015290565b90815260200190565b60405181810167ffffffffffffffff811182821017156133ab57600080fd5b604052919050565b60005b838110156133ce5781810151838201526020016133b6565b838111156133dd576000848401525b50505050565b6001600160a01b03811681146133f857600080fd5b5056fea26469706673582212201a5c4c07233e6d6c55eb81ecd98668becdccfc86ee7c2c20a38832fcd6a00c5464736f6c63430006020033"

// DeploySCPHA2 deploys a new Ethereum contract, binding an instance of SCPHA2 to it.
func DeploySCPHA2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SCPHA2, error) {
	parsed, err := abi.JSON(strings.NewReader(SCPHA2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SCPHA2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SCPHA2{SCPHA2Caller: SCPHA2Caller{contract: contract}, SCPHA2Transactor: SCPHA2Transactor{contract: contract}, SCPHA2Filterer: SCPHA2Filterer{contract: contract}}, nil
}

// SCPHA2 is an auto generated Go binding around an Ethereum contract.
type SCPHA2 struct {
	SCPHA2Caller     // Read-only binding to the contract
	SCPHA2Transactor // Write-only binding to the contract
	SCPHA2Filterer   // Log filterer for contract events
}

// SCPHA2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SCPHA2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SCPHA2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCPHA2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCPHA2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCPHA2Session struct {
	Contract     *SCPHA2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCPHA2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCPHA2CallerSession struct {
	Contract *SCPHA2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SCPHA2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCPHA2TransactorSession struct {
	Contract     *SCPHA2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCPHA2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SCPHA2Raw struct {
	Contract *SCPHA2 // Generic contract binding to access the raw methods on
}

// SCPHA2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCPHA2CallerRaw struct {
	Contract *SCPHA2Caller // Generic read-only contract binding to access the raw methods on
}

// SCPHA2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCPHA2TransactorRaw struct {
	Contract *SCPHA2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSCPHA2 creates a new instance of SCPHA2, bound to a specific deployed contract.
func NewSCPHA2(address common.Address, backend bind.ContractBackend) (*SCPHA2, error) {
	contract, err := bindSCPHA2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCPHA2{SCPHA2Caller: SCPHA2Caller{contract: contract}, SCPHA2Transactor: SCPHA2Transactor{contract: contract}, SCPHA2Filterer: SCPHA2Filterer{contract: contract}}, nil
}

// NewSCPHA2Caller creates a new read-only instance of SCPHA2, bound to a specific deployed contract.
func NewSCPHA2Caller(address common.Address, caller bind.ContractCaller) (*SCPHA2Caller, error) {
	contract, err := bindSCPHA2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCPHA2Caller{contract: contract}, nil
}

// NewSCPHA2Transactor creates a new write-only instance of SCPHA2, bound to a specific deployed contract.
func NewSCPHA2Transactor(address common.Address, transactor bind.ContractTransactor) (*SCPHA2Transactor, error) {
	contract, err := bindSCPHA2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCPHA2Transactor{contract: contract}, nil
}

// NewSCPHA2Filterer creates a new log filterer instance of SCPHA2, bound to a specific deployed contract.
func NewSCPHA2Filterer(address common.Address, filterer bind.ContractFilterer) (*SCPHA2Filterer, error) {
	contract, err := bindSCPHA2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCPHA2Filterer{contract: contract}, nil
}

// bindSCPHA2 binds a generic wrapper to an already deployed contract.
func bindSCPHA2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SCPHA2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCPHA2 *SCPHA2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCPHA2.Contract.SCPHA2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCPHA2 *SCPHA2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCPHA2.Contract.SCPHA2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCPHA2 *SCPHA2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCPHA2.Contract.SCPHA2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCPHA2 *SCPHA2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCPHA2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCPHA2 *SCPHA2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCPHA2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCPHA2 *SCPHA2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCPHA2.Contract.contract.Transact(opts, method, params...)
}

// ConsentFinalBillAfterDisputeRaisePatient is a paid mutator transaction binding the contract method 0x9c61ab02.
//
// Solidity: function consentFinalBill_AfterDisputeRaise_Patient(address System_Users_Info_address, address P_HA_1_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2Transactor) ConsentFinalBillAfterDisputeRaisePatient(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "consentFinalBill_AfterDisputeRaise_Patient", System_Users_Info_address, P_HA_1_address, _finalBillID, _hID)
}

// ConsentFinalBillAfterDisputeRaisePatient is a paid mutator transaction binding the contract method 0x9c61ab02.
//
// Solidity: function consentFinalBill_AfterDisputeRaise_Patient(address System_Users_Info_address, address P_HA_1_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2Session) ConsentFinalBillAfterDisputeRaisePatient(System_Users_Info_address common.Address, P_HA_1_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.ConsentFinalBillAfterDisputeRaisePatient(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _finalBillID, _hID)
}

// ConsentFinalBillAfterDisputeRaisePatient is a paid mutator transaction binding the contract method 0x9c61ab02.
//
// Solidity: function consentFinalBill_AfterDisputeRaise_Patient(address System_Users_Info_address, address P_HA_1_address, uint256 _finalBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) ConsentFinalBillAfterDisputeRaisePatient(System_Users_Info_address common.Address, P_HA_1_address common.Address, _finalBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.ConsentFinalBillAfterDisputeRaisePatient(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _finalBillID, _hID)
}

// HospitalLeaveDischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0xf86e8d5c.
//
// Solidity: function hospital_leave_DischargeAndGenerateFinalCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pID) returns()
func (_SCPHA2 *SCPHA2Transactor) HospitalLeaveDischargeAndGenerateFinalCostBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "hospital_leave_DischargeAndGenerateFinalCostBill", System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pID)
}

// HospitalLeaveDischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0xf86e8d5c.
//
// Solidity: function hospital_leave_DischargeAndGenerateFinalCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pID) returns()
func (_SCPHA2 *SCPHA2Session) HospitalLeaveDischargeAndGenerateFinalCostBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveDischargeAndGenerateFinalCostBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pID)
}

// HospitalLeaveDischargeAndGenerateFinalCostBill is a paid mutator transaction binding the contract method 0xf86e8d5c.
//
// Solidity: function hospital_leave_DischargeAndGenerateFinalCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) HospitalLeaveDischargeAndGenerateFinalCostBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveDischargeAndGenerateFinalCostBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pID)
}

// HospitalLeaveGenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0x32707436.
//
// Solidity: function hospital_leave_generateEstimatedCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Transactor) HospitalLeaveGenerateEstimatedCostBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "hospital_leave_generateEstimatedCostBill", System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveGenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0x32707436.
//
// Solidity: function hospital_leave_generateEstimatedCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Session) HospitalLeaveGenerateEstimatedCostBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveGenerateEstimatedCostBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveGenerateEstimatedCostBill is a paid mutator transaction binding the contract method 0x32707436.
//
// Solidity: function hospital_leave_generateEstimatedCostBill(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) HospitalLeaveGenerateEstimatedCostBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveGenerateEstimatedCostBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x9c8662f5.
//
// Solidity: function hospital_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Transactor) HospitalLeaveKeepHashToBlockchain(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "hospital_leave_keepHashToBlockchain", System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x9c8662f5.
//
// Solidity: function hospital_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Session) HospitalLeaveKeepHashToBlockchain(System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveKeepHashToBlockchain(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x9c8662f5.
//
// Solidity: function hospital_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _pID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) HospitalLeaveKeepHashToBlockchain(System_Users_Info_address common.Address, P_HA_1_address common.Address, _pID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveKeepHashToBlockchain(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _pID, _estimatedBillID)
}

// HospitalLeaveKeyReveal is a paid mutator transaction binding the contract method 0xfdd54eca.
//
// Solidity: function hospital_leave_keyReveal(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA2 *SCPHA2Transactor) HospitalLeaveKeyReveal(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "hospital_leave_keyReveal", System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pId)
}

// HospitalLeaveKeyReveal is a paid mutator transaction binding the contract method 0xfdd54eca.
//
// Solidity: function hospital_leave_keyReveal(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA2 *SCPHA2Session) HospitalLeaveKeyReveal(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveKeyReveal(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pId)
}

// HospitalLeaveKeyReveal is a paid mutator transaction binding the contract method 0xfdd54eca.
//
// Solidity: function hospital_leave_keyReveal(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _pId) returns()
func (_SCPHA2 *SCPHA2TransactorSession) HospitalLeaveKeyReveal(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _pId *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.HospitalLeaveKeyReveal(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _pId)
}

// PatientLeaveConsentFinalBill is a paid mutator transaction binding the contract method 0x1c3f24b5.
//
// Solidity: function patient_leave_consentFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2Transactor) PatientLeaveConsentFinalBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "patient_leave_consentFinalBill", System_Users_Info_address, P_HA_1_address, _estimatedBillID, _hID)
}

// PatientLeaveConsentFinalBill is a paid mutator transaction binding the contract method 0x1c3f24b5.
//
// Solidity: function patient_leave_consentFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2Session) PatientLeaveConsentFinalBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveConsentFinalBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _hID)
}

// PatientLeaveConsentFinalBill is a paid mutator transaction binding the contract method 0x1c3f24b5.
//
// Solidity: function patient_leave_consentFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _estimatedBillID, uint256 _hID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) PatientLeaveConsentFinalBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _estimatedBillID *big.Int, _hID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveConsentFinalBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _estimatedBillID, _hID)
}

// PatientLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x232b4868.
//
// Solidity: function patient_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Transactor) PatientLeaveKeepHashToBlockchain(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "patient_leave_keepHashToBlockchain", System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x232b4868.
//
// Solidity: function patient_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Session) PatientLeaveKeepHashToBlockchain(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveKeepHashToBlockchain(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveKeepHashToBlockchain is a paid mutator transaction binding the contract method 0x232b4868.
//
// Solidity: function patient_leave_keepHashToBlockchain(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) PatientLeaveKeepHashToBlockchain(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveKeepHashToBlockchain(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveLockEstimatedAmount is a paid mutator transaction binding the contract method 0x2b5294db.
//
// Solidity: function patient_leave_lockEstimatedAmount(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Transactor) PatientLeaveLockEstimatedAmount(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "patient_leave_lockEstimatedAmount", System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveLockEstimatedAmount is a paid mutator transaction binding the contract method 0x2b5294db.
//
// Solidity: function patient_leave_lockEstimatedAmount(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Session) PatientLeaveLockEstimatedAmount(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveLockEstimatedAmount(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveLockEstimatedAmount is a paid mutator transaction binding the contract method 0x2b5294db.
//
// Solidity: function patient_leave_lockEstimatedAmount(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) PatientLeaveLockEstimatedAmount(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveLockEstimatedAmount(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveVerifyAndGiveConsent is a paid mutator transaction binding the contract method 0x05c29d34.
//
// Solidity: function patient_leave_verifyAndGiveConsent(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Transactor) PatientLeaveVerifyAndGiveConsent(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "patient_leave_verifyAndGiveConsent", System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveVerifyAndGiveConsent is a paid mutator transaction binding the contract method 0x05c29d34.
//
// Solidity: function patient_leave_verifyAndGiveConsent(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2Session) PatientLeaveVerifyAndGiveConsent(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveVerifyAndGiveConsent(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// PatientLeaveVerifyAndGiveConsent is a paid mutator transaction binding the contract method 0x05c29d34.
//
// Solidity: function patient_leave_verifyAndGiveConsent(address System_Users_Info_address, address P_HA_1_address, uint256 _hID, uint256 _estimatedBillID) returns()
func (_SCPHA2 *SCPHA2TransactorSession) PatientLeaveVerifyAndGiveConsent(System_Users_Info_address common.Address, P_HA_1_address common.Address, _hID *big.Int, _estimatedBillID *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.PatientLeaveVerifyAndGiveConsent(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _hID, _estimatedBillID)
}

// SolveDisputeAndModifiedFinalBill is a paid mutator transaction binding the contract method 0x9057947b.
//
// Solidity: function solveDisputeAndModifiedFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _disputeID, uint256 _estimatedBillID, uint256 _finalBillID, uint256 _pID, uint256 _modifiedCost) returns()
func (_SCPHA2 *SCPHA2Transactor) SolveDisputeAndModifiedFinalBill(opts *bind.TransactOpts, System_Users_Info_address common.Address, P_HA_1_address common.Address, _disputeID *big.Int, _estimatedBillID *big.Int, _finalBillID *big.Int, _pID *big.Int, _modifiedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA2.contract.Transact(opts, "solveDisputeAndModifiedFinalBill", System_Users_Info_address, P_HA_1_address, _disputeID, _estimatedBillID, _finalBillID, _pID, _modifiedCost)
}

// SolveDisputeAndModifiedFinalBill is a paid mutator transaction binding the contract method 0x9057947b.
//
// Solidity: function solveDisputeAndModifiedFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _disputeID, uint256 _estimatedBillID, uint256 _finalBillID, uint256 _pID, uint256 _modifiedCost) returns()
func (_SCPHA2 *SCPHA2Session) SolveDisputeAndModifiedFinalBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _disputeID *big.Int, _estimatedBillID *big.Int, _finalBillID *big.Int, _pID *big.Int, _modifiedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.SolveDisputeAndModifiedFinalBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _disputeID, _estimatedBillID, _finalBillID, _pID, _modifiedCost)
}

// SolveDisputeAndModifiedFinalBill is a paid mutator transaction binding the contract method 0x9057947b.
//
// Solidity: function solveDisputeAndModifiedFinalBill(address System_Users_Info_address, address P_HA_1_address, uint256 _disputeID, uint256 _estimatedBillID, uint256 _finalBillID, uint256 _pID, uint256 _modifiedCost) returns()
func (_SCPHA2 *SCPHA2TransactorSession) SolveDisputeAndModifiedFinalBill(System_Users_Info_address common.Address, P_HA_1_address common.Address, _disputeID *big.Int, _estimatedBillID *big.Int, _finalBillID *big.Int, _pID *big.Int, _modifiedCost *big.Int) (*types.Transaction, error) {
	return _SCPHA2.Contract.SolveDisputeAndModifiedFinalBill(&_SCPHA2.TransactOpts, System_Users_Info_address, P_HA_1_address, _disputeID, _estimatedBillID, _finalBillID, _pID, _modifiedCost)
}

// SystemUsersInfoABI is the input ABI used to generate the binding from.
const SystemUsersInfoABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dataBaseOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dboAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dboID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dboName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dboAddressTodboID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dboAddress\",\"type\":\"address\"}],\"name\":\"dboByAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dboID\",\"type\":\"uint256\"}],\"name\":\"dboById\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dboIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_dboAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_dboName\",\"type\":\"string\"}],\"name\":\"dboRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callingPatientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_IC\",\"type\":\"uint256\"}],\"name\":\"deny_permission_patient_To_IC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hospitalAdd\",\"type\":\"address\"}],\"name\":\"getHospitalbyAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hospitalID\",\"type\":\"uint256\"}],\"name\":\"getHospitalbyID\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_insuranceAdd\",\"type\":\"address\"}],\"name\":\"getInsuranceCompanybyAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_insuranceID\",\"type\":\"uint256\"}],\"name\":\"getInsuranceCompanybyID\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_patientAdd\",\"type\":\"address\"}],\"name\":\"getPatientbyAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_patientID\",\"type\":\"uint256\"}],\"name\":\"getPatientbyID\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govtAuthority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callingPatientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hID\",\"type\":\"uint256\"}],\"name\":\"grant_permission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callingPatientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_IC\",\"type\":\"uint256\"}],\"name\":\"grant_permission_patient_To_IC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hospital\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"haddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"hospitalID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hospitalIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"hospitalRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"icAddressToicId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"icaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"insuranceCompanyID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_icaddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"insuranceCompanyRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insuranceIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"isEmptyString\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"isEqual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mapFromAddToId_P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"numDigits\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"patient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"paddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"patientID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"mob\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"add\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"patientIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mob\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_add\",\"type\":\"string\"}],\"name\":\"patientRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rcAddressTorcID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_rcAddress\",\"type\":\"address\"}],\"name\":\"rcByAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rcID\",\"type\":\"uint256\"}],\"name\":\"rcById\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rcIDGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_rcAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_rcName\",\"type\":\"string\"}],\"name\":\"rcRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"read_permission_to_H\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"read_permission_to_IC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"researchCommunity\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rcID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"rcName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"write_permission_to_H\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemUsersInfoFuncSigs maps the 4-byte function signature to its string representation.
var SystemUsersInfoFuncSigs = map[string]string{
	"6055cd46": "dataBaseOwner(uint256)",
	"076443f1": "dboAddressTodboID(address)",
	"63c5ee17": "dboByAddress(address)",
	"3de79ecd": "dboById(uint256)",
	"1f4f7b49": "dboIDGenerator()",
	"119ab101": "dboRegistration(address,string)",
	"672d06c3": "deny_permission_patient_To_IC(address,uint256)",
	"1264bd00": "getHospitalbyAddress(address)",
	"fd4cc079": "getHospitalbyID(uint256)",
	"59611726": "getInsuranceCompanybyAddress(address)",
	"8a547f66": "getInsuranceCompanybyID(uint256)",
	"9836b825": "getPatientbyAddress(address)",
	"25253518": "getPatientbyID(uint256)",
	"6424d58b": "govtAuthority()",
	"dd822159": "grant_permission(address,uint256)",
	"184e035f": "grant_permission_patient_To_IC(address,uint256)",
	"dca06da7": "hospital(uint256)",
	"aceaf16d": "hospitalIDGenerator()",
	"7fe309c2": "hospitalRegistration(string)",
	"1741fc8e": "icAddressToicId(address)",
	"9700e19e": "insurance(uint256)",
	"484ea50d": "insuranceCompanyRegistration(address,string)",
	"bffb0d50": "insuranceIDGenerator()",
	"a45f379e": "isEmptyString(string)",
	"465c4105": "isEqual(string,string)",
	"b1738999": "mapFromAddToId_P(address)",
	"db9d28d5": "numDigits(uint256)",
	"74607d91": "patient(uint256)",
	"c784e114": "patientIDGenerator()",
	"94361d5c": "patientRegistration(string,uint256,uint256,string)",
	"09169d75": "rcAddressTorcID(address)",
	"3da77dfe": "rcByAddress(address)",
	"88b047f9": "rcById(uint256)",
	"566b1787": "rcIDGenerator()",
	"ef437b3e": "rcRegistration(address,string)",
	"8d2aa67d": "read_permission_to_H(address,address)",
	"fb9ef618": "read_permission_to_IC(address,address)",
	"81a0458d": "researchCommunity(uint256)",
	"e14e1dac": "write_permission_to_H(address,address)",
}

// SystemUsersInfoBin is the compiled bytecode used for deploying new contracts.
var SystemUsersInfoBin = "0x60806040526000600155600060045560006007556000600a556000600d5534801561002957600080fd5b50600080546001600160a01b031916331790556120628061004b6000396000f3fe608060405234801561001057600080fd5b506004361061023d5760003560e01c80637fe309c21161013b578063b1738999116100b8578063dd8221591161007c578063dd822159146104e6578063e14e1dac146104f9578063ef437b3e1461050c578063fb9ef6181461051f578063fd4cc079146105325761023d565b8063b173899914610490578063bffb0d50146104a3578063c784e114146104ab578063db9d28d5146104b3578063dca06da7146104d35761023d565b806394361d5c116100ff57806394361d5c1461043c5780639700e19e1461044f5780639836b82514610462578063a45f379e14610475578063aceaf16d146104885761023d565b80637fe309c2146103dd57806381a0458d146103f057806388b047f9146104035780638a547f66146104165780638d2aa67d146104295761023d565b80633de79ecd116101c95780636055cd461161018d5780636055cd461461037c57806363c5ee171461038f5780636424d58b146103a2578063672d06c3146103b757806374607d91146103ca5761023d565b80633de79ecd1461031b578063465c41051461032e578063484ea50d1461034e578063566b17871461036157806359611726146103695761023d565b80631741fc8e116102105780631741fc8e146102b5578063184e035f146102c85780631f4f7b49146102db57806325253518146102e35780633da77dfe146103085761023d565b8063076443f11461024257806309169d751461026b578063119ab1011461027e5780631264bd0014610293575b600080fd5b610255610250366004611d1c565b610545565b6040516102629190611ffd565b60405180910390f35b610255610279366004611d1c565b610557565b61029161028c366004611d3f565b610569565b005b6102a66102a1366004611d1c565b6106ba565b60405161026293929190611f6d565b6102556102c3366004611d1c565b610801565b6102916102d6366004611dc5565b610813565b610255610858565b6102f66102f1366004611ef6565b61085e565b60405161026296959493929190611f9d565b6102a6610316366004611d1c565b610a9d565b6102a6610329366004611ef6565b610be5565b61034161033c366004611e2b565b610d13565b6040516102629190611ff2565b61029161035c366004611d3f565b610d29565b610255610e60565b6102a6610377366004611d1c565b610e66565b6102a661038a366004611ef6565b610eff565b6102a661039d366004611d1c565b610fc7565b6103aa611062565b6040516102629190611f59565b6102916103c5366004611dc5565b611071565b6102f66103d8366004611ef6565b6110b3565b6102916103eb366004611df0565b61121b565b6102a66103fe366004611ef6565b611347565b6102a6610411366004611ef6565b611354565b6102a6610424366004611ef6565b6113d6565b610341610437366004611d8d565b611456565b61029161044a366004611e82565b611476565b6102a661045d366004611ef6565b61164a565b6102f6610470366004611d1c565b611657565b610341610483366004611df0565b6118b1565b6102556118d1565b61025561049e366004611d1c565b6118d7565b6102556118e9565b6102556118ef565b6104c66104c1366004611ef6565b6118f5565b6040516102629190612006565b6102a66104e1366004611ef6565b611913565b6102916104f4366004611dc5565b611920565b610341610507366004611d8d565b611986565b61029161051a366004611d3f565b6119a6565b61034161052d366004611d8d565b611af7565b6102a6610540366004611ef6565b611b17565b600c6020526000908152604090205481565b600f6020526000908152604090205481565b60005433906001600160a01b0316811461058257600080fd5b6001600160a01b0383166000908152600c6020526040902054156105a557600080fd5b6105ae826118b1565b156105b857600080fd5b600a805460010190556105c9611b97565b6001600160a01b038481168252600a54602080840191825260408401868152600b805460018101825560009190915285517f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9600390920291820180546001600160a01b0319169190961617855592517f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01dba840155518051859493610692937f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01dbb909101920190611bc1565b5050600a546001600160a01b039095166000908152600c602052604090209490945550505050565b6001600160a01b0381166000908152600660205260408120548190606090600181108015906106eb57506004548111155b6106f457600080fd5b60006001820390506005818154811061070957fe5b6000918252602090912060039091020154600580546001600160a01b03909216918390811061073457fe5b9060005260206000209060030201600101546005838154811061075357fe5b600091825260209182902060026003909202018101805460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835290928391908301828280156107eb5780601f106107c0576101008083540402835291602001916107eb565b820191906000526020600020905b8154815290600101906020018083116107ce57829003601f168201915b5050505050905094509450945050509193909250565b60096020526000908152604090205481565b600061081e826113d6565b50506001600160a01b03938416600090815260126020908152604080832096909316825294909452909220805460ff191660011790555050565b600a5481565b600080606060008060606001871015801561087b57506001548711155b61088457600080fd5b60006001880390506002818154811061089957fe5b6000918252602090912060069091020154600280546001600160a01b0390921691839081106108c457fe5b906000526020600020906006020160010154600283815481106108e357fe5b90600052602060002090600602016002016002848154811061090157fe5b906000526020600020906006020160030160009054906101000a900460ff166002858154811061092d57fe5b9060005260206000209060060201600401546002868154811061094c57fe5b9060005260206000209060060201600501838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109f25780601f106109c7576101008083540402835291602001916109f2565b820191906000526020600020905b8154815290600101906020018083116109d557829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815295995086945092508401905082828015610a805780601f10610a5557610100808354040283529160200191610a80565b820191906000526020600020905b815481529060010190602001808311610a6357829003601f168201915b505050505090509650965096509650965096505091939550919395565b6001600160a01b0381166000908152600f6020526040812054819060609060018110801590610ace5750600d548111155b610ad757600080fd5b600e6001820381548110610ae757fe5b6000918252602090912060039091020154600e80546001600160a01b03909216916000198401908110610b1657fe5b906000526020600020906003020160010154600e6001840381548110610b3857fe5b600091825260209182902060026003909202018101805460408051601f600019610100600186161502019093169490940491820185900485028401850190528083529092839190830182828015610bd05780601f10610ba557610100808354040283529160200191610bd0565b820191906000526020600020905b815481529060010190602001808311610bb357829003601f168201915b50505050509050935093509350509193909250565b600080606060018410158015610bfd5750600a548411155b610c0657600080fd5b600b6001850381548110610c1657fe5b6000918252602090912060039091020154600b80546001600160a01b03909216916000198701908110610c4557fe5b906000526020600020906003020160010154600b6001870381548110610c6757fe5b600091825260209182902060026003909202018101805460408051601f600019610100600186161502019093169490940491820185900485028401850190528083529092839190830182828015610cff5780601f10610cd457610100808354040283529160200191610cff565b820191906000526020600020905b815481529060010190602001808311610ce257829003601f168201915b505050505090509250925092509193909250565b8051602091820120825192909101919091201490565b6001600160a01b03821660009081526009602052604090205415610d4c57600080fd5b610d55816118b1565b15610d5f57600080fd5b600780546001019055610d70611b97565b6001600160a01b0383811682526007546020808401918252604084018581526008805460018101825560009190915285517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3600390920291820180546001600160a01b0319169190961617855592517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4840155518051859493610e39937ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee5909101920190611bc1565b50506007546001600160a01b03909416600090815260096020526040902093909355505050565b600d5481565b6001600160a01b038116600090815260096020526040812054819060609060018110801590610e9757506007548111155b610ea057600080fd5b600060018203905060088181548110610eb557fe5b6000918252602090912060039091020154600880546001600160a01b039092169183908110610ee057fe5b9060005260206000209060030201600101546008838154811061075357fe5b600b8181548110610f0c57fe5b6000918252602091829020600391909102018054600180830154600280850180546040805161010096831615969096026000190190911692909204601f81018890048802850188019092528184526001600160a01b0390941696509094919291830182828015610fbd5780601f10610f9257610100808354040283529160200191610fbd565b820191906000526020600020905b815481529060010190602001808311610fa057829003601f168201915b5050505050905083565b6001600160a01b0381166000908152600c6020526040812054819060609060018110801590610ff85750600a548111155b61100157600080fd5b600b600182038154811061101157fe5b6000918252602090912060039091020154600b80546001600160a01b0390921691600019840190811061104057fe5b906000526020600020906003020160010154600b6001840381548110610b3857fe5b6000546001600160a01b031681565b600061107c826113d6565b50506001600160a01b03938416600090815260126020908152604080832096909316825294909452909220805460ff191690555050565b600281815481106110c057fe5b6000918252602091829020600691909102018054600180830154600280850180546040805161010096831615969096026000190190911692909204601f81018890048802850188019092528184526001600160a01b03909416965090949192918301828280156111715780601f1061114657610100808354040283529160200191611171565b820191906000526020600020905b81548152906001019060200180831161115457829003601f168201915b505050506003830154600484015460058501805460408051602060026101006001861615026000190190941693909304601f8101849004840282018401909252818152969760ff90951696939550908301828280156112115780601f106111e657610100808354040283529160200191611211565b820191906000526020600020905b8154815290600101906020018083116111f457829003601f168201915b5050505050905086565b336000908152600660205260409020541561123557600080fd5b61123e816118b1565b1561124857600080fd5b60048054600101908190553360009081526006602052604090205561126b611b97565b506040805160608101825233815260045460208083019182529282018481526005805460018101825560009190915283517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0600390920291820180546001600160a01b0319166001600160a01b0390921691909117815592517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db18201559051805193948594611340937f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db2019290910190611bc1565b5050505050565b600e8181548110610f0c57fe5b60008060606001841015801561136c5750600d548411155b61137557600080fd5b600e600185038154811061138557fe5b6000918252602090912060039091020154600e80546001600160a01b039092169160001987019081106113b457fe5b906000526020600020906003020160010154600e6001870381548110610c6757fe5b6000806060600184101580156113ee57506007548411155b6113f757600080fd5b60006001850390506008818154811061140c57fe5b6000918252602090912060039091020154600880546001600160a01b03909216918390811061143757fe5b90600052602060002090600302016001015460088381548110610b3857fe5b601060209081526000928352604080842090915290825290205460ff1681565b336000908152600360205260409020541561149057600080fd5b611499846118b1565b156114a357600080fd5b6000831180156114b4575060648311155b6114bd57600080fd5b6000821180156114d857506114d1826118f5565b60ff16600a145b6114e157600080fd5b6114ea816118b1565b156114f457600080fd5b6001805481019081905533600090815260036020526040902055611516611c3f565b506040805160c08101825233815260018054602080840191825293830188815260ff881660608501526080840187905260a08401869052600280549384018155600052835160069093027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810180546001600160a01b03959095166001600160a01b031990951694909417845591517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf83015551805193948594611600937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad0019290910190611bc1565b50606082015160038201805460ff191660ff9092169190911790556080820151600482015560a08201518051611640916005840191602090910190611bc1565b5050505050505050565b60088181548110610f0c57fe5b6001600160a01b03811660009081526003602052604081205481906060908290819083906001811080159061168e57506001548111155b61169757600080fd5b6000600182039050600281815481106116ac57fe5b6000918252602090912060069091020154600280546001600160a01b0390921691839081106116d757fe5b906000526020600020906006020160010154600283815481106116f657fe5b90600052602060002090600602016002016002848154811061171457fe5b906000526020600020906006020160030160009054906101000a900460ff166002858154811061174057fe5b9060005260206000209060060201600401546002868154811061175f57fe5b9060005260206000209060060201600501838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118055780601f106117da57610100808354040283529160200191611805565b820191906000526020600020905b8154815290600101906020018083116117e857829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959950869450925084019050828280156118935780601f1061186857610100808354040283529160200191611893565b820191906000526020600020905b81548152906001019060200180831161187657829003601f168201915b50505050509050975097509750975097509750505091939550919395565b805160009082906118c65760019150506118cc565b60009150505b919050565b60045481565b60036020526000908152604090205481565b60075481565b60015481565b6000805b821561190d57600a830492506001016118f9565b92915050565b60058181548110610f0c57fe5b600061192b82611b17565b50506001600160a01b039384166000818152601060209081526040808320949097168083529381528682208054600160ff19918216811790925593835260118252878320948352939052949094208054909416179092555050565b601160209081526000928352604080842090915290825290205460ff1681565b60005433906001600160a01b031681146119bf57600080fd5b6001600160a01b0383166000908152600f6020526040902054156119e257600080fd5b6119eb826118b1565b156119f557600080fd5b600d80546001019055611a06611b97565b6001600160a01b038481168252600d54602080840191825260408401868152600e805460018101825560009190915285517fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd600390920291820180546001600160a01b0319169190961617855592517fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fe840155518051859493611acf937fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3ff909101920190611bc1565b5050600d546001600160a01b039095166000908152600f602052604090209490945550505050565b601260209081526000928352604080842090915290825290205460ff1681565b600080606060018410158015611b2f57506004548411155b611b3857600080fd5b600060018503905060058181548110611b4d57fe5b6000918252602090912060039091020154600580546001600160a01b039092169183908110611b7857fe5b90600052602060002090600302016001015460058381548110610b3857fe5b604051806060016040528060006001600160a01b0316815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c0257805160ff1916838001178555611c2f565b82800160010185558215611c2f579182015b82811115611c2f578251825591602001919060010190611c14565b50611c3b929150611c81565b5090565b6040518060c0016040528060006001600160a01b031681526020016000815260200160608152602001600060ff16815260200160008152602001606081525090565b611c9b91905b80821115611c3b5760008155600101611c87565b90565b600082601f830112611cae578081fd5b813567ffffffffffffffff80821115611cc5578283fd5b604051601f8301601f191681016020018281118282101715611ce5578485fd5b604052828152925082848301602001861015611d0057600080fd5b8260208601602083013760006020848301015250505092915050565b600060208284031215611d2d578081fd5b8135611d3881612014565b9392505050565b60008060408385031215611d51578081fd5b8235611d5c81612014565b9150602083013567ffffffffffffffff811115611d77578182fd5b611d8385828601611c9e565b9150509250929050565b60008060408385031215611d9f578182fd5b8235611daa81612014565b91506020830135611dba81612014565b809150509250929050565b60008060408385031215611dd7578182fd5b8235611de281612014565b946020939093013593505050565b600060208284031215611e01578081fd5b813567ffffffffffffffff811115611e17578182fd5b611e2384828501611c9e565b949350505050565b60008060408385031215611e3d578182fd5b823567ffffffffffffffff80821115611e54578384fd5b611e6086838701611c9e565b93506020850135915080821115611e75578283fd5b50611d8385828601611c9e565b60008060008060808587031215611e97578182fd5b843567ffffffffffffffff80821115611eae578384fd5b611eba88838901611c9e565b955060208701359450604087013593506060870135915080821115611edd578283fd5b50611eea87828801611c9e565b91505092959194509250565b600060208284031215611f07578081fd5b5035919050565b60008151808452815b81811015611f3357602081850181015186830182015201611f17565b81811115611f445782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b600060018060a01b038516825283602083015260606040830152611f946060830184611f0e565b95945050505050565b600060018060a01b038816825286602083015260c06040830152611fc460c0830187611f0e565b60ff8616606084015284608084015282810360a0840152611fe58185611f0e565b9998505050505050505050565b901515815260200190565b90815260200190565b60ff91909116815260200190565b6001600160a01b038116811461202957600080fd5b5056fea2646970667358221220c002ffb38120dfbe2ad95d7425535edaa9cff4fadb7f4e97838db0d5a82114d464736f6c63430006020033"

// DeploySystemUsersInfo deploys a new Ethereum contract, binding an instance of SystemUsersInfo to it.
func DeploySystemUsersInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemUsersInfo, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemUsersInfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemUsersInfoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemUsersInfo{SystemUsersInfoCaller: SystemUsersInfoCaller{contract: contract}, SystemUsersInfoTransactor: SystemUsersInfoTransactor{contract: contract}, SystemUsersInfoFilterer: SystemUsersInfoFilterer{contract: contract}}, nil
}

// SystemUsersInfo is an auto generated Go binding around an Ethereum contract.
type SystemUsersInfo struct {
	SystemUsersInfoCaller     // Read-only binding to the contract
	SystemUsersInfoTransactor // Write-only binding to the contract
	SystemUsersInfoFilterer   // Log filterer for contract events
}

// SystemUsersInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemUsersInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemUsersInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemUsersInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemUsersInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemUsersInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemUsersInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemUsersInfoSession struct {
	Contract     *SystemUsersInfo  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemUsersInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemUsersInfoCallerSession struct {
	Contract *SystemUsersInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SystemUsersInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemUsersInfoTransactorSession struct {
	Contract     *SystemUsersInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SystemUsersInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemUsersInfoRaw struct {
	Contract *SystemUsersInfo // Generic contract binding to access the raw methods on
}

// SystemUsersInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemUsersInfoCallerRaw struct {
	Contract *SystemUsersInfoCaller // Generic read-only contract binding to access the raw methods on
}

// SystemUsersInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemUsersInfoTransactorRaw struct {
	Contract *SystemUsersInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemUsersInfo creates a new instance of SystemUsersInfo, bound to a specific deployed contract.
func NewSystemUsersInfo(address common.Address, backend bind.ContractBackend) (*SystemUsersInfo, error) {
	contract, err := bindSystemUsersInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemUsersInfo{SystemUsersInfoCaller: SystemUsersInfoCaller{contract: contract}, SystemUsersInfoTransactor: SystemUsersInfoTransactor{contract: contract}, SystemUsersInfoFilterer: SystemUsersInfoFilterer{contract: contract}}, nil
}

// NewSystemUsersInfoCaller creates a new read-only instance of SystemUsersInfo, bound to a specific deployed contract.
func NewSystemUsersInfoCaller(address common.Address, caller bind.ContractCaller) (*SystemUsersInfoCaller, error) {
	contract, err := bindSystemUsersInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemUsersInfoCaller{contract: contract}, nil
}

// NewSystemUsersInfoTransactor creates a new write-only instance of SystemUsersInfo, bound to a specific deployed contract.
func NewSystemUsersInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemUsersInfoTransactor, error) {
	contract, err := bindSystemUsersInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemUsersInfoTransactor{contract: contract}, nil
}

// NewSystemUsersInfoFilterer creates a new log filterer instance of SystemUsersInfo, bound to a specific deployed contract.
func NewSystemUsersInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemUsersInfoFilterer, error) {
	contract, err := bindSystemUsersInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemUsersInfoFilterer{contract: contract}, nil
}

// bindSystemUsersInfo binds a generic wrapper to an already deployed contract.
func bindSystemUsersInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemUsersInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemUsersInfo *SystemUsersInfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemUsersInfo.Contract.SystemUsersInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemUsersInfo *SystemUsersInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.SystemUsersInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemUsersInfo *SystemUsersInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.SystemUsersInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemUsersInfo *SystemUsersInfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemUsersInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemUsersInfo *SystemUsersInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemUsersInfo *SystemUsersInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.contract.Transact(opts, method, params...)
}

// DataBaseOwner is a free data retrieval call binding the contract method 0x6055cd46.
//
// Solidity: function dataBaseOwner(uint256 ) constant returns(address dboAddress, uint256 dboID, string dboName)
func (_SystemUsersInfo *SystemUsersInfoCaller) DataBaseOwner(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DboAddress common.Address
	DboID      *big.Int
	DboName    string
}, error) {
	ret := new(struct {
		DboAddress common.Address
		DboID      *big.Int
		DboName    string
	})
	out := ret
	err := _SystemUsersInfo.contract.Call(opts, out, "dataBaseOwner", arg0)
	return *ret, err
}

// DataBaseOwner is a free data retrieval call binding the contract method 0x6055cd46.
//
// Solidity: function dataBaseOwner(uint256 ) constant returns(address dboAddress, uint256 dboID, string dboName)
func (_SystemUsersInfo *SystemUsersInfoSession) DataBaseOwner(arg0 *big.Int) (struct {
	DboAddress common.Address
	DboID      *big.Int
	DboName    string
}, error) {
	return _SystemUsersInfo.Contract.DataBaseOwner(&_SystemUsersInfo.CallOpts, arg0)
}

// DataBaseOwner is a free data retrieval call binding the contract method 0x6055cd46.
//
// Solidity: function dataBaseOwner(uint256 ) constant returns(address dboAddress, uint256 dboID, string dboName)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) DataBaseOwner(arg0 *big.Int) (struct {
	DboAddress common.Address
	DboID      *big.Int
	DboName    string
}, error) {
	return _SystemUsersInfo.Contract.DataBaseOwner(&_SystemUsersInfo.CallOpts, arg0)
}

// DboAddressTodboID is a free data retrieval call binding the contract method 0x076443f1.
//
// Solidity: function dboAddressTodboID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) DboAddressTodboID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "dboAddressTodboID", arg0)
	return *ret0, err
}

// DboAddressTodboID is a free data retrieval call binding the contract method 0x076443f1.
//
// Solidity: function dboAddressTodboID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) DboAddressTodboID(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.DboAddressTodboID(&_SystemUsersInfo.CallOpts, arg0)
}

// DboAddressTodboID is a free data retrieval call binding the contract method 0x076443f1.
//
// Solidity: function dboAddressTodboID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) DboAddressTodboID(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.DboAddressTodboID(&_SystemUsersInfo.CallOpts, arg0)
}

// DboByAddress is a free data retrieval call binding the contract method 0x63c5ee17.
//
// Solidity: function dboByAddress(address _dboAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) DboByAddress(opts *bind.CallOpts, _dboAddress common.Address) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "dboByAddress", _dboAddress)
	return *ret0, *ret1, *ret2, err
}

// DboByAddress is a free data retrieval call binding the contract method 0x63c5ee17.
//
// Solidity: function dboByAddress(address _dboAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) DboByAddress(_dboAddress common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.DboByAddress(&_SystemUsersInfo.CallOpts, _dboAddress)
}

// DboByAddress is a free data retrieval call binding the contract method 0x63c5ee17.
//
// Solidity: function dboByAddress(address _dboAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) DboByAddress(_dboAddress common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.DboByAddress(&_SystemUsersInfo.CallOpts, _dboAddress)
}

// DboById is a free data retrieval call binding the contract method 0x3de79ecd.
//
// Solidity: function dboById(uint256 _dboID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) DboById(opts *bind.CallOpts, _dboID *big.Int) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "dboById", _dboID)
	return *ret0, *ret1, *ret2, err
}

// DboById is a free data retrieval call binding the contract method 0x3de79ecd.
//
// Solidity: function dboById(uint256 _dboID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) DboById(_dboID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.DboById(&_SystemUsersInfo.CallOpts, _dboID)
}

// DboById is a free data retrieval call binding the contract method 0x3de79ecd.
//
// Solidity: function dboById(uint256 _dboID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) DboById(_dboID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.DboById(&_SystemUsersInfo.CallOpts, _dboID)
}

// DboIDGenerator is a free data retrieval call binding the contract method 0x1f4f7b49.
//
// Solidity: function dboIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) DboIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "dboIDGenerator")
	return *ret0, err
}

// DboIDGenerator is a free data retrieval call binding the contract method 0x1f4f7b49.
//
// Solidity: function dboIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) DboIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.DboIDGenerator(&_SystemUsersInfo.CallOpts)
}

// DboIDGenerator is a free data retrieval call binding the contract method 0x1f4f7b49.
//
// Solidity: function dboIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) DboIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.DboIDGenerator(&_SystemUsersInfo.CallOpts)
}

// GetHospitalbyAddress is a free data retrieval call binding the contract method 0x1264bd00.
//
// Solidity: function getHospitalbyAddress(address _hospitalAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetHospitalbyAddress(opts *bind.CallOpts, _hospitalAdd common.Address) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getHospitalbyAddress", _hospitalAdd)
	return *ret0, *ret1, *ret2, err
}

// GetHospitalbyAddress is a free data retrieval call binding the contract method 0x1264bd00.
//
// Solidity: function getHospitalbyAddress(address _hospitalAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetHospitalbyAddress(_hospitalAdd common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetHospitalbyAddress(&_SystemUsersInfo.CallOpts, _hospitalAdd)
}

// GetHospitalbyAddress is a free data retrieval call binding the contract method 0x1264bd00.
//
// Solidity: function getHospitalbyAddress(address _hospitalAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetHospitalbyAddress(_hospitalAdd common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetHospitalbyAddress(&_SystemUsersInfo.CallOpts, _hospitalAdd)
}

// GetHospitalbyID is a free data retrieval call binding the contract method 0xfd4cc079.
//
// Solidity: function getHospitalbyID(uint256 _hospitalID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetHospitalbyID(opts *bind.CallOpts, _hospitalID *big.Int) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getHospitalbyID", _hospitalID)
	return *ret0, *ret1, *ret2, err
}

// GetHospitalbyID is a free data retrieval call binding the contract method 0xfd4cc079.
//
// Solidity: function getHospitalbyID(uint256 _hospitalID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetHospitalbyID(_hospitalID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetHospitalbyID(&_SystemUsersInfo.CallOpts, _hospitalID)
}

// GetHospitalbyID is a free data retrieval call binding the contract method 0xfd4cc079.
//
// Solidity: function getHospitalbyID(uint256 _hospitalID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetHospitalbyID(_hospitalID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetHospitalbyID(&_SystemUsersInfo.CallOpts, _hospitalID)
}

// GetInsuranceCompanybyAddress is a free data retrieval call binding the contract method 0x59611726.
//
// Solidity: function getInsuranceCompanybyAddress(address _insuranceAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetInsuranceCompanybyAddress(opts *bind.CallOpts, _insuranceAdd common.Address) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getInsuranceCompanybyAddress", _insuranceAdd)
	return *ret0, *ret1, *ret2, err
}

// GetInsuranceCompanybyAddress is a free data retrieval call binding the contract method 0x59611726.
//
// Solidity: function getInsuranceCompanybyAddress(address _insuranceAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetInsuranceCompanybyAddress(_insuranceAdd common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetInsuranceCompanybyAddress(&_SystemUsersInfo.CallOpts, _insuranceAdd)
}

// GetInsuranceCompanybyAddress is a free data retrieval call binding the contract method 0x59611726.
//
// Solidity: function getInsuranceCompanybyAddress(address _insuranceAdd) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetInsuranceCompanybyAddress(_insuranceAdd common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetInsuranceCompanybyAddress(&_SystemUsersInfo.CallOpts, _insuranceAdd)
}

// GetInsuranceCompanybyID is a free data retrieval call binding the contract method 0x8a547f66.
//
// Solidity: function getInsuranceCompanybyID(uint256 _insuranceID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetInsuranceCompanybyID(opts *bind.CallOpts, _insuranceID *big.Int) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getInsuranceCompanybyID", _insuranceID)
	return *ret0, *ret1, *ret2, err
}

// GetInsuranceCompanybyID is a free data retrieval call binding the contract method 0x8a547f66.
//
// Solidity: function getInsuranceCompanybyID(uint256 _insuranceID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetInsuranceCompanybyID(_insuranceID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetInsuranceCompanybyID(&_SystemUsersInfo.CallOpts, _insuranceID)
}

// GetInsuranceCompanybyID is a free data retrieval call binding the contract method 0x8a547f66.
//
// Solidity: function getInsuranceCompanybyID(uint256 _insuranceID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetInsuranceCompanybyID(_insuranceID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetInsuranceCompanybyID(&_SystemUsersInfo.CallOpts, _insuranceID)
}

// GetPatientbyAddress is a free data retrieval call binding the contract method 0x9836b825.
//
// Solidity: function getPatientbyAddress(address _patientAdd) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetPatientbyAddress(opts *bind.CallOpts, _patientAdd common.Address) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new(*big.Int)
		ret5 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getPatientbyAddress", _patientAdd)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetPatientbyAddress is a free data retrieval call binding the contract method 0x9836b825.
//
// Solidity: function getPatientbyAddress(address _patientAdd) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetPatientbyAddress(_patientAdd common.Address) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetPatientbyAddress(&_SystemUsersInfo.CallOpts, _patientAdd)
}

// GetPatientbyAddress is a free data retrieval call binding the contract method 0x9836b825.
//
// Solidity: function getPatientbyAddress(address _patientAdd) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetPatientbyAddress(_patientAdd common.Address) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetPatientbyAddress(&_SystemUsersInfo.CallOpts, _patientAdd)
}

// GetPatientbyID is a free data retrieval call binding the contract method 0x25253518.
//
// Solidity: function getPatientbyID(uint256 _patientID) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) GetPatientbyID(opts *bind.CallOpts, _patientID *big.Int) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new(*big.Int)
		ret5 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "getPatientbyID", _patientID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetPatientbyID is a free data retrieval call binding the contract method 0x25253518.
//
// Solidity: function getPatientbyID(uint256 _patientID) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) GetPatientbyID(_patientID *big.Int) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetPatientbyID(&_SystemUsersInfo.CallOpts, _patientID)
}

// GetPatientbyID is a free data retrieval call binding the contract method 0x25253518.
//
// Solidity: function getPatientbyID(uint256 _patientID) constant returns(address, uint256, string, uint8, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GetPatientbyID(_patientID *big.Int) (common.Address, *big.Int, string, uint8, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.GetPatientbyID(&_SystemUsersInfo.CallOpts, _patientID)
}

// GovtAuthority is a free data retrieval call binding the contract method 0x6424d58b.
//
// Solidity: function govtAuthority() constant returns(address)
func (_SystemUsersInfo *SystemUsersInfoCaller) GovtAuthority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "govtAuthority")
	return *ret0, err
}

// GovtAuthority is a free data retrieval call binding the contract method 0x6424d58b.
//
// Solidity: function govtAuthority() constant returns(address)
func (_SystemUsersInfo *SystemUsersInfoSession) GovtAuthority() (common.Address, error) {
	return _SystemUsersInfo.Contract.GovtAuthority(&_SystemUsersInfo.CallOpts)
}

// GovtAuthority is a free data retrieval call binding the contract method 0x6424d58b.
//
// Solidity: function govtAuthority() constant returns(address)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) GovtAuthority() (common.Address, error) {
	return _SystemUsersInfo.Contract.GovtAuthority(&_SystemUsersInfo.CallOpts)
}

// Hospital is a free data retrieval call binding the contract method 0xdca06da7.
//
// Solidity: function hospital(uint256 ) constant returns(address haddr, uint256 hospitalID, string name)
func (_SystemUsersInfo *SystemUsersInfoCaller) Hospital(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Haddr      common.Address
	HospitalID *big.Int
	Name       string
}, error) {
	ret := new(struct {
		Haddr      common.Address
		HospitalID *big.Int
		Name       string
	})
	out := ret
	err := _SystemUsersInfo.contract.Call(opts, out, "hospital", arg0)
	return *ret, err
}

// Hospital is a free data retrieval call binding the contract method 0xdca06da7.
//
// Solidity: function hospital(uint256 ) constant returns(address haddr, uint256 hospitalID, string name)
func (_SystemUsersInfo *SystemUsersInfoSession) Hospital(arg0 *big.Int) (struct {
	Haddr      common.Address
	HospitalID *big.Int
	Name       string
}, error) {
	return _SystemUsersInfo.Contract.Hospital(&_SystemUsersInfo.CallOpts, arg0)
}

// Hospital is a free data retrieval call binding the contract method 0xdca06da7.
//
// Solidity: function hospital(uint256 ) constant returns(address haddr, uint256 hospitalID, string name)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) Hospital(arg0 *big.Int) (struct {
	Haddr      common.Address
	HospitalID *big.Int
	Name       string
}, error) {
	return _SystemUsersInfo.Contract.Hospital(&_SystemUsersInfo.CallOpts, arg0)
}

// HospitalIDGenerator is a free data retrieval call binding the contract method 0xaceaf16d.
//
// Solidity: function hospitalIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) HospitalIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "hospitalIDGenerator")
	return *ret0, err
}

// HospitalIDGenerator is a free data retrieval call binding the contract method 0xaceaf16d.
//
// Solidity: function hospitalIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) HospitalIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.HospitalIDGenerator(&_SystemUsersInfo.CallOpts)
}

// HospitalIDGenerator is a free data retrieval call binding the contract method 0xaceaf16d.
//
// Solidity: function hospitalIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) HospitalIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.HospitalIDGenerator(&_SystemUsersInfo.CallOpts)
}

// IcAddressToicId is a free data retrieval call binding the contract method 0x1741fc8e.
//
// Solidity: function icAddressToicId(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) IcAddressToicId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "icAddressToicId", arg0)
	return *ret0, err
}

// IcAddressToicId is a free data retrieval call binding the contract method 0x1741fc8e.
//
// Solidity: function icAddressToicId(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) IcAddressToicId(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.IcAddressToicId(&_SystemUsersInfo.CallOpts, arg0)
}

// IcAddressToicId is a free data retrieval call binding the contract method 0x1741fc8e.
//
// Solidity: function icAddressToicId(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) IcAddressToicId(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.IcAddressToicId(&_SystemUsersInfo.CallOpts, arg0)
}

// Insurance is a free data retrieval call binding the contract method 0x9700e19e.
//
// Solidity: function insurance(uint256 ) constant returns(address icaddr, uint256 insuranceCompanyID, string name)
func (_SystemUsersInfo *SystemUsersInfoCaller) Insurance(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Icaddr             common.Address
	InsuranceCompanyID *big.Int
	Name               string
}, error) {
	ret := new(struct {
		Icaddr             common.Address
		InsuranceCompanyID *big.Int
		Name               string
	})
	out := ret
	err := _SystemUsersInfo.contract.Call(opts, out, "insurance", arg0)
	return *ret, err
}

// Insurance is a free data retrieval call binding the contract method 0x9700e19e.
//
// Solidity: function insurance(uint256 ) constant returns(address icaddr, uint256 insuranceCompanyID, string name)
func (_SystemUsersInfo *SystemUsersInfoSession) Insurance(arg0 *big.Int) (struct {
	Icaddr             common.Address
	InsuranceCompanyID *big.Int
	Name               string
}, error) {
	return _SystemUsersInfo.Contract.Insurance(&_SystemUsersInfo.CallOpts, arg0)
}

// Insurance is a free data retrieval call binding the contract method 0x9700e19e.
//
// Solidity: function insurance(uint256 ) constant returns(address icaddr, uint256 insuranceCompanyID, string name)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) Insurance(arg0 *big.Int) (struct {
	Icaddr             common.Address
	InsuranceCompanyID *big.Int
	Name               string
}, error) {
	return _SystemUsersInfo.Contract.Insurance(&_SystemUsersInfo.CallOpts, arg0)
}

// InsuranceIDGenerator is a free data retrieval call binding the contract method 0xbffb0d50.
//
// Solidity: function insuranceIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) InsuranceIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "insuranceIDGenerator")
	return *ret0, err
}

// InsuranceIDGenerator is a free data retrieval call binding the contract method 0xbffb0d50.
//
// Solidity: function insuranceIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) InsuranceIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.InsuranceIDGenerator(&_SystemUsersInfo.CallOpts)
}

// InsuranceIDGenerator is a free data retrieval call binding the contract method 0xbffb0d50.
//
// Solidity: function insuranceIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) InsuranceIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.InsuranceIDGenerator(&_SystemUsersInfo.CallOpts)
}

// IsEmptyString is a free data retrieval call binding the contract method 0xa45f379e.
//
// Solidity: function isEmptyString(string s) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCaller) IsEmptyString(opts *bind.CallOpts, s string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "isEmptyString", s)
	return *ret0, err
}

// IsEmptyString is a free data retrieval call binding the contract method 0xa45f379e.
//
// Solidity: function isEmptyString(string s) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoSession) IsEmptyString(s string) (bool, error) {
	return _SystemUsersInfo.Contract.IsEmptyString(&_SystemUsersInfo.CallOpts, s)
}

// IsEmptyString is a free data retrieval call binding the contract method 0xa45f379e.
//
// Solidity: function isEmptyString(string s) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) IsEmptyString(s string) (bool, error) {
	return _SystemUsersInfo.Contract.IsEmptyString(&_SystemUsersInfo.CallOpts, s)
}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCaller) IsEqual(opts *bind.CallOpts, a string, b string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "isEqual", a, b)
	return *ret0, err
}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoSession) IsEqual(a string, b string) (bool, error) {
	return _SystemUsersInfo.Contract.IsEqual(&_SystemUsersInfo.CallOpts, a, b)
}

// IsEqual is a free data retrieval call binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string a, string b) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) IsEqual(a string, b string) (bool, error) {
	return _SystemUsersInfo.Contract.IsEqual(&_SystemUsersInfo.CallOpts, a, b)
}

// MapFromAddToIdP is a free data retrieval call binding the contract method 0xb1738999.
//
// Solidity: function mapFromAddToId_P(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) MapFromAddToIdP(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "mapFromAddToId_P", arg0)
	return *ret0, err
}

// MapFromAddToIdP is a free data retrieval call binding the contract method 0xb1738999.
//
// Solidity: function mapFromAddToId_P(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) MapFromAddToIdP(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.MapFromAddToIdP(&_SystemUsersInfo.CallOpts, arg0)
}

// MapFromAddToIdP is a free data retrieval call binding the contract method 0xb1738999.
//
// Solidity: function mapFromAddToId_P(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) MapFromAddToIdP(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.MapFromAddToIdP(&_SystemUsersInfo.CallOpts, arg0)
}

// NumDigits is a free data retrieval call binding the contract method 0xdb9d28d5.
//
// Solidity: function numDigits(uint256 number) constant returns(uint8)
func (_SystemUsersInfo *SystemUsersInfoCaller) NumDigits(opts *bind.CallOpts, number *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "numDigits", number)
	return *ret0, err
}

// NumDigits is a free data retrieval call binding the contract method 0xdb9d28d5.
//
// Solidity: function numDigits(uint256 number) constant returns(uint8)
func (_SystemUsersInfo *SystemUsersInfoSession) NumDigits(number *big.Int) (uint8, error) {
	return _SystemUsersInfo.Contract.NumDigits(&_SystemUsersInfo.CallOpts, number)
}

// NumDigits is a free data retrieval call binding the contract method 0xdb9d28d5.
//
// Solidity: function numDigits(uint256 number) constant returns(uint8)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) NumDigits(number *big.Int) (uint8, error) {
	return _SystemUsersInfo.Contract.NumDigits(&_SystemUsersInfo.CallOpts, number)
}

// Patient is a free data retrieval call binding the contract method 0x74607d91.
//
// Solidity: function patient(uint256 ) constant returns(address paddr, uint256 patientID, string name, uint8 age, uint256 mob, string add)
func (_SystemUsersInfo *SystemUsersInfoCaller) Patient(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Paddr     common.Address
	PatientID *big.Int
	Name      string
	Age       uint8
	Mob       *big.Int
	Add       string
}, error) {
	ret := new(struct {
		Paddr     common.Address
		PatientID *big.Int
		Name      string
		Age       uint8
		Mob       *big.Int
		Add       string
	})
	out := ret
	err := _SystemUsersInfo.contract.Call(opts, out, "patient", arg0)
	return *ret, err
}

// Patient is a free data retrieval call binding the contract method 0x74607d91.
//
// Solidity: function patient(uint256 ) constant returns(address paddr, uint256 patientID, string name, uint8 age, uint256 mob, string add)
func (_SystemUsersInfo *SystemUsersInfoSession) Patient(arg0 *big.Int) (struct {
	Paddr     common.Address
	PatientID *big.Int
	Name      string
	Age       uint8
	Mob       *big.Int
	Add       string
}, error) {
	return _SystemUsersInfo.Contract.Patient(&_SystemUsersInfo.CallOpts, arg0)
}

// Patient is a free data retrieval call binding the contract method 0x74607d91.
//
// Solidity: function patient(uint256 ) constant returns(address paddr, uint256 patientID, string name, uint8 age, uint256 mob, string add)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) Patient(arg0 *big.Int) (struct {
	Paddr     common.Address
	PatientID *big.Int
	Name      string
	Age       uint8
	Mob       *big.Int
	Add       string
}, error) {
	return _SystemUsersInfo.Contract.Patient(&_SystemUsersInfo.CallOpts, arg0)
}

// PatientIDGenerator is a free data retrieval call binding the contract method 0xc784e114.
//
// Solidity: function patientIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) PatientIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "patientIDGenerator")
	return *ret0, err
}

// PatientIDGenerator is a free data retrieval call binding the contract method 0xc784e114.
//
// Solidity: function patientIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) PatientIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.PatientIDGenerator(&_SystemUsersInfo.CallOpts)
}

// PatientIDGenerator is a free data retrieval call binding the contract method 0xc784e114.
//
// Solidity: function patientIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) PatientIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.PatientIDGenerator(&_SystemUsersInfo.CallOpts)
}

// RcAddressTorcID is a free data retrieval call binding the contract method 0x09169d75.
//
// Solidity: function rcAddressTorcID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) RcAddressTorcID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "rcAddressTorcID", arg0)
	return *ret0, err
}

// RcAddressTorcID is a free data retrieval call binding the contract method 0x09169d75.
//
// Solidity: function rcAddressTorcID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) RcAddressTorcID(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.RcAddressTorcID(&_SystemUsersInfo.CallOpts, arg0)
}

// RcAddressTorcID is a free data retrieval call binding the contract method 0x09169d75.
//
// Solidity: function rcAddressTorcID(address ) constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) RcAddressTorcID(arg0 common.Address) (*big.Int, error) {
	return _SystemUsersInfo.Contract.RcAddressTorcID(&_SystemUsersInfo.CallOpts, arg0)
}

// RcByAddress is a free data retrieval call binding the contract method 0x3da77dfe.
//
// Solidity: function rcByAddress(address _rcAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) RcByAddress(opts *bind.CallOpts, _rcAddress common.Address) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "rcByAddress", _rcAddress)
	return *ret0, *ret1, *ret2, err
}

// RcByAddress is a free data retrieval call binding the contract method 0x3da77dfe.
//
// Solidity: function rcByAddress(address _rcAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) RcByAddress(_rcAddress common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.RcByAddress(&_SystemUsersInfo.CallOpts, _rcAddress)
}

// RcByAddress is a free data retrieval call binding the contract method 0x3da77dfe.
//
// Solidity: function rcByAddress(address _rcAddress) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) RcByAddress(_rcAddress common.Address) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.RcByAddress(&_SystemUsersInfo.CallOpts, _rcAddress)
}

// RcById is a free data retrieval call binding the contract method 0x88b047f9.
//
// Solidity: function rcById(uint256 _rcID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCaller) RcById(opts *bind.CallOpts, _rcID *big.Int) (common.Address, *big.Int, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SystemUsersInfo.contract.Call(opts, out, "rcById", _rcID)
	return *ret0, *ret1, *ret2, err
}

// RcById is a free data retrieval call binding the contract method 0x88b047f9.
//
// Solidity: function rcById(uint256 _rcID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoSession) RcById(_rcID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.RcById(&_SystemUsersInfo.CallOpts, _rcID)
}

// RcById is a free data retrieval call binding the contract method 0x88b047f9.
//
// Solidity: function rcById(uint256 _rcID) constant returns(address, uint256, string)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) RcById(_rcID *big.Int) (common.Address, *big.Int, string, error) {
	return _SystemUsersInfo.Contract.RcById(&_SystemUsersInfo.CallOpts, _rcID)
}

// RcIDGenerator is a free data retrieval call binding the contract method 0x566b1787.
//
// Solidity: function rcIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCaller) RcIDGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "rcIDGenerator")
	return *ret0, err
}

// RcIDGenerator is a free data retrieval call binding the contract method 0x566b1787.
//
// Solidity: function rcIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoSession) RcIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.RcIDGenerator(&_SystemUsersInfo.CallOpts)
}

// RcIDGenerator is a free data retrieval call binding the contract method 0x566b1787.
//
// Solidity: function rcIDGenerator() constant returns(uint256)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) RcIDGenerator() (*big.Int, error) {
	return _SystemUsersInfo.Contract.RcIDGenerator(&_SystemUsersInfo.CallOpts)
}

// ReadPermissionToH is a free data retrieval call binding the contract method 0x8d2aa67d.
//
// Solidity: function read_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCaller) ReadPermissionToH(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "read_permission_to_H", arg0, arg1)
	return *ret0, err
}

// ReadPermissionToH is a free data retrieval call binding the contract method 0x8d2aa67d.
//
// Solidity: function read_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoSession) ReadPermissionToH(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.ReadPermissionToH(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// ReadPermissionToH is a free data retrieval call binding the contract method 0x8d2aa67d.
//
// Solidity: function read_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) ReadPermissionToH(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.ReadPermissionToH(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// ReadPermissionToIC is a free data retrieval call binding the contract method 0xfb9ef618.
//
// Solidity: function read_permission_to_IC(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCaller) ReadPermissionToIC(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "read_permission_to_IC", arg0, arg1)
	return *ret0, err
}

// ReadPermissionToIC is a free data retrieval call binding the contract method 0xfb9ef618.
//
// Solidity: function read_permission_to_IC(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoSession) ReadPermissionToIC(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.ReadPermissionToIC(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// ReadPermissionToIC is a free data retrieval call binding the contract method 0xfb9ef618.
//
// Solidity: function read_permission_to_IC(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) ReadPermissionToIC(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.ReadPermissionToIC(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// ResearchCommunity is a free data retrieval call binding the contract method 0x81a0458d.
//
// Solidity: function researchCommunity(uint256 ) constant returns(address rcAddress, uint256 rcID, string rcName)
func (_SystemUsersInfo *SystemUsersInfoCaller) ResearchCommunity(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RcAddress common.Address
	RcID      *big.Int
	RcName    string
}, error) {
	ret := new(struct {
		RcAddress common.Address
		RcID      *big.Int
		RcName    string
	})
	out := ret
	err := _SystemUsersInfo.contract.Call(opts, out, "researchCommunity", arg0)
	return *ret, err
}

// ResearchCommunity is a free data retrieval call binding the contract method 0x81a0458d.
//
// Solidity: function researchCommunity(uint256 ) constant returns(address rcAddress, uint256 rcID, string rcName)
func (_SystemUsersInfo *SystemUsersInfoSession) ResearchCommunity(arg0 *big.Int) (struct {
	RcAddress common.Address
	RcID      *big.Int
	RcName    string
}, error) {
	return _SystemUsersInfo.Contract.ResearchCommunity(&_SystemUsersInfo.CallOpts, arg0)
}

// ResearchCommunity is a free data retrieval call binding the contract method 0x81a0458d.
//
// Solidity: function researchCommunity(uint256 ) constant returns(address rcAddress, uint256 rcID, string rcName)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) ResearchCommunity(arg0 *big.Int) (struct {
	RcAddress common.Address
	RcID      *big.Int
	RcName    string
}, error) {
	return _SystemUsersInfo.Contract.ResearchCommunity(&_SystemUsersInfo.CallOpts, arg0)
}

// WritePermissionToH is a free data retrieval call binding the contract method 0xe14e1dac.
//
// Solidity: function write_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCaller) WritePermissionToH(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemUsersInfo.contract.Call(opts, out, "write_permission_to_H", arg0, arg1)
	return *ret0, err
}

// WritePermissionToH is a free data retrieval call binding the contract method 0xe14e1dac.
//
// Solidity: function write_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoSession) WritePermissionToH(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.WritePermissionToH(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// WritePermissionToH is a free data retrieval call binding the contract method 0xe14e1dac.
//
// Solidity: function write_permission_to_H(address , address ) constant returns(bool)
func (_SystemUsersInfo *SystemUsersInfoCallerSession) WritePermissionToH(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _SystemUsersInfo.Contract.WritePermissionToH(&_SystemUsersInfo.CallOpts, arg0, arg1)
}

// DboRegistration is a paid mutator transaction binding the contract method 0x119ab101.
//
// Solidity: function dboRegistration(address _dboAddress, string _dboName) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) DboRegistration(opts *bind.TransactOpts, _dboAddress common.Address, _dboName string) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "dboRegistration", _dboAddress, _dboName)
}

// DboRegistration is a paid mutator transaction binding the contract method 0x119ab101.
//
// Solidity: function dboRegistration(address _dboAddress, string _dboName) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) DboRegistration(_dboAddress common.Address, _dboName string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.DboRegistration(&_SystemUsersInfo.TransactOpts, _dboAddress, _dboName)
}

// DboRegistration is a paid mutator transaction binding the contract method 0x119ab101.
//
// Solidity: function dboRegistration(address _dboAddress, string _dboName) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) DboRegistration(_dboAddress common.Address, _dboName string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.DboRegistration(&_SystemUsersInfo.TransactOpts, _dboAddress, _dboName)
}

// DenyPermissionPatientToIC is a paid mutator transaction binding the contract method 0x672d06c3.
//
// Solidity: function deny_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) DenyPermissionPatientToIC(opts *bind.TransactOpts, callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "deny_permission_patient_To_IC", callingPatientAddress, _IC)
}

// DenyPermissionPatientToIC is a paid mutator transaction binding the contract method 0x672d06c3.
//
// Solidity: function deny_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) DenyPermissionPatientToIC(callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.DenyPermissionPatientToIC(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _IC)
}

// DenyPermissionPatientToIC is a paid mutator transaction binding the contract method 0x672d06c3.
//
// Solidity: function deny_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) DenyPermissionPatientToIC(callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.DenyPermissionPatientToIC(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _IC)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd822159.
//
// Solidity: function grant_permission(address callingPatientAddress, uint256 _hID) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) GrantPermission(opts *bind.TransactOpts, callingPatientAddress common.Address, _hID *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "grant_permission", callingPatientAddress, _hID)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd822159.
//
// Solidity: function grant_permission(address callingPatientAddress, uint256 _hID) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) GrantPermission(callingPatientAddress common.Address, _hID *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.GrantPermission(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _hID)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd822159.
//
// Solidity: function grant_permission(address callingPatientAddress, uint256 _hID) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) GrantPermission(callingPatientAddress common.Address, _hID *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.GrantPermission(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _hID)
}

// GrantPermissionPatientToIC is a paid mutator transaction binding the contract method 0x184e035f.
//
// Solidity: function grant_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) GrantPermissionPatientToIC(opts *bind.TransactOpts, callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "grant_permission_patient_To_IC", callingPatientAddress, _IC)
}

// GrantPermissionPatientToIC is a paid mutator transaction binding the contract method 0x184e035f.
//
// Solidity: function grant_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) GrantPermissionPatientToIC(callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.GrantPermissionPatientToIC(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _IC)
}

// GrantPermissionPatientToIC is a paid mutator transaction binding the contract method 0x184e035f.
//
// Solidity: function grant_permission_patient_To_IC(address callingPatientAddress, uint256 _IC) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) GrantPermissionPatientToIC(callingPatientAddress common.Address, _IC *big.Int) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.GrantPermissionPatientToIC(&_SystemUsersInfo.TransactOpts, callingPatientAddress, _IC)
}

// HospitalRegistration is a paid mutator transaction binding the contract method 0x7fe309c2.
//
// Solidity: function hospitalRegistration(string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) HospitalRegistration(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "hospitalRegistration", _name)
}

// HospitalRegistration is a paid mutator transaction binding the contract method 0x7fe309c2.
//
// Solidity: function hospitalRegistration(string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) HospitalRegistration(_name string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.HospitalRegistration(&_SystemUsersInfo.TransactOpts, _name)
}

// HospitalRegistration is a paid mutator transaction binding the contract method 0x7fe309c2.
//
// Solidity: function hospitalRegistration(string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) HospitalRegistration(_name string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.HospitalRegistration(&_SystemUsersInfo.TransactOpts, _name)
}

// InsuranceCompanyRegistration is a paid mutator transaction binding the contract method 0x484ea50d.
//
// Solidity: function insuranceCompanyRegistration(address _icaddr, string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) InsuranceCompanyRegistration(opts *bind.TransactOpts, _icaddr common.Address, _name string) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "insuranceCompanyRegistration", _icaddr, _name)
}

// InsuranceCompanyRegistration is a paid mutator transaction binding the contract method 0x484ea50d.
//
// Solidity: function insuranceCompanyRegistration(address _icaddr, string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) InsuranceCompanyRegistration(_icaddr common.Address, _name string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.InsuranceCompanyRegistration(&_SystemUsersInfo.TransactOpts, _icaddr, _name)
}

// InsuranceCompanyRegistration is a paid mutator transaction binding the contract method 0x484ea50d.
//
// Solidity: function insuranceCompanyRegistration(address _icaddr, string _name) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) InsuranceCompanyRegistration(_icaddr common.Address, _name string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.InsuranceCompanyRegistration(&_SystemUsersInfo.TransactOpts, _icaddr, _name)
}

// PatientRegistration is a paid mutator transaction binding the contract method 0x94361d5c.
//
// Solidity: function patientRegistration(string _name, uint256 _age, uint256 _mob, string _add) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) PatientRegistration(opts *bind.TransactOpts, _name string, _age *big.Int, _mob *big.Int, _add string) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "patientRegistration", _name, _age, _mob, _add)
}

// PatientRegistration is a paid mutator transaction binding the contract method 0x94361d5c.
//
// Solidity: function patientRegistration(string _name, uint256 _age, uint256 _mob, string _add) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) PatientRegistration(_name string, _age *big.Int, _mob *big.Int, _add string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.PatientRegistration(&_SystemUsersInfo.TransactOpts, _name, _age, _mob, _add)
}

// PatientRegistration is a paid mutator transaction binding the contract method 0x94361d5c.
//
// Solidity: function patientRegistration(string _name, uint256 _age, uint256 _mob, string _add) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) PatientRegistration(_name string, _age *big.Int, _mob *big.Int, _add string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.PatientRegistration(&_SystemUsersInfo.TransactOpts, _name, _age, _mob, _add)
}

// RcRegistration is a paid mutator transaction binding the contract method 0xef437b3e.
//
// Solidity: function rcRegistration(address _rcAddress, string _rcName) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactor) RcRegistration(opts *bind.TransactOpts, _rcAddress common.Address, _rcName string) (*types.Transaction, error) {
	return _SystemUsersInfo.contract.Transact(opts, "rcRegistration", _rcAddress, _rcName)
}

// RcRegistration is a paid mutator transaction binding the contract method 0xef437b3e.
//
// Solidity: function rcRegistration(address _rcAddress, string _rcName) returns()
func (_SystemUsersInfo *SystemUsersInfoSession) RcRegistration(_rcAddress common.Address, _rcName string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.RcRegistration(&_SystemUsersInfo.TransactOpts, _rcAddress, _rcName)
}

// RcRegistration is a paid mutator transaction binding the contract method 0xef437b3e.
//
// Solidity: function rcRegistration(address _rcAddress, string _rcName) returns()
func (_SystemUsersInfo *SystemUsersInfoTransactorSession) RcRegistration(_rcAddress common.Address, _rcName string) (*types.Transaction, error) {
	return _SystemUsersInfo.Contract.RcRegistration(&_SystemUsersInfo.TransactOpts, _rcAddress, _rcName)
}
