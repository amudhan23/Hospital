// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_rc_dbo

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

// SCRCDBOABI is the input ABI used to generate the binding from.
const SCRCDBOABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_researchDataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hashOfMedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sigOnHashOfMedData\",\"type\":\"bytes\"}],\"name\":\"provideMedicalData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"System_Users_Info_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataBaseOwner\",\"type\":\"address\"}],\"name\":\"requestMedicalData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"researchCommunityId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"researchCommunityIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"researchCommunity_databaseOwner_researchDataId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"researchDataDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rcAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dataBaseOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeStamp_researchCommunityRequestData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStamp_dataBaseOwnerSendData\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfMedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sigOnHashOfMedData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"researchDataDetailsId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLimitForGivingData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SCRCDBOFuncSigs maps the 4-byte function signature to its string representation.
var SCRCDBOFuncSigs = map[string]string{
	"47bbf530": "provideMedicalData(address,uint256,bytes32,bytes)",
	"5ca57322": "requestMedicalData(address,address)",
	"a33ac38b": "researchCommunityId()",
	"83401587": "researchCommunityIdToAddress(uint256)",
	"a94717b8": "researchCommunity_databaseOwner_researchDataId(address,address)",
	"b2d8e6a0": "researchDataDetails(uint256)",
	"af55a453": "researchDataDetailsId()",
	"6e8361be": "timeLimitForGivingData()",
}

// SCRCDBOBin is the compiled bytecode used for deploying new contracts.
var SCRCDBOBin = "0x608060405260006002556000600355606460045534801561001f57600080fd5b5061098a8061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a33ac38b1161005b578063a33ac38b146100f3578063a94717b8146100fb578063af55a4531461010e578063b2d8e6a01461011657610088565b806347bbf5301461008d5780635ca57322146100a25780636e8361be146100b557806383401587146100d3575b600080fd5b6100a061009b366004610777565b61013c565b005b6100a06100b0366004610743565b610232565b6100bd610507565b6040516100ca9190610867565b60405180910390f35b6100e66100e1366004610823565b61050d565b6040516100ca9190610853565b6100bd610528565b6100bd610109366004610743565b61052e565b6100bd61054b565b610129610124366004610823565b610551565b6040516100ca9796959493929190610870565b336001600160a01b031660018085038154811061015557fe5b60009182526020909120600260079092020101546001600160a01b03161461017c57600080fd5b60045460018085038154811061018e57fe5b906000526020600020906007020160030154420311156101ad57600080fd5b816001808503815481106101bd57fe5b906000526020600020906007020160050181905550806001808503815481106101e257fe5b90600052602060002090600702016006019080519060200190610206929190610639565b504260018085038154811061021757fe5b90600052602060002090600702016004018190555050505050565b6040516309169d7560e01b815282906001600160a01b038216906309169d7590610260903390600401610853565b60206040518083038186803b15801561027857600080fd5b505afa15801561028c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102b0919081019061083b565b6102b957600080fd5b60405163076443f160e01b81526001600160a01b0382169063076443f1906102e5908590600401610853565b60206040518083038186803b1580156102fd57600080fd5b505afa158015610311573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610335919081019061083b565b61033e57600080fd5b60028054600101905561034f6106b7565b60025481523360208083019182526001600160a01b03858116604085019081524260608601908152600180548082018255600091909152865160079091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810191825595517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7870180549186166001600160a01b031992831617905592517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf8870180549190951693169290921790925590517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf984015560808401517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfa84015560a08401517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfb84015560c08401518051859492936104d1937fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfc01920190610639565b50506002543360009081526006602090815260408083206001600160a01b03989098168352969052949094209390935550505050565b60045481565b6005602052600090815260409020546001600160a01b031681565b60035481565b600660209081526000928352604080842090915290825290205481565b60025481565b6001818154811061055e57fe5b6000918252602091829020600790910201805460018083015460028085015460038601546004870154600588015460068901805460408051601f6000199b841615610100029b909b01909216979097049889018c90048c0281018c01909652878652979a506001600160a01b039586169993909516979196909591939290919083018282801561062f5780601f106106045761010080835404028352916020019161062f565b820191906000526020600020905b81548152906001019060200180831161061257829003601f168201915b5050505050905087565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061067a57805160ff19168380011785556106a7565b828001600101855582156106a7579182015b828111156106a757825182559160200191906001019061068c565b506106b3929150610709565b5090565b6040518060e001604052806000815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000815260200160008019168152602001606081525090565b61072391905b808211156106b3576000815560010161070f565b90565b80356001600160a01b038116811461073d57600080fd5b92915050565b60008060408385031215610755578182fd5b61075f8484610726565b915061076e8460208501610726565b90509250929050565b6000806000806080858703121561078c578182fd5b84356001600160a01b03811681146107a2578283fd5b93506020850135925060408501359150606085013567ffffffffffffffff8111156107cb578182fd5b80860187601f8201126107dc578283fd5b803591506107f16107ec83610924565b6108fd565b828152886020848401011115610805578384fd5b610816836020830160208501610948565b9598949750929550505050565b600060208284031215610834578081fd5b5035919050565b60006020828403121561084c578081fd5b5051919050565b6001600160a01b0391909116815260200190565b90815260200190565b6000888252602060018060a01b03808a16828501528089166040850152508660608401528560808401528460a084015260e060c084015283518060e0850152825b818110156108ce57858101830151858201610100015282016108b1565b818111156108e0578361010083870101525b50601f01601f191692909201610100019998505050505050505050565b60405181810167ffffffffffffffff8111828210171561091c57600080fd5b604052919050565b600067ffffffffffffffff82111561093a578081fd5b50601f01601f191660200190565b8281833750600091015256fea2646970667358221220e27fa9c0ba963f04db6b16f171bac857599392603201b4ac4f051e50de4c72bf64736f6c63430006020033"

// DeploySCRCDBO deploys a new Ethereum contract, binding an instance of SCRCDBO to it.
func DeploySCRCDBO(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SCRCDBO, error) {
	parsed, err := abi.JSON(strings.NewReader(SCRCDBOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SCRCDBOBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SCRCDBO{SCRCDBOCaller: SCRCDBOCaller{contract: contract}, SCRCDBOTransactor: SCRCDBOTransactor{contract: contract}, SCRCDBOFilterer: SCRCDBOFilterer{contract: contract}}, nil
}

// SCRCDBO is an auto generated Go binding around an Ethereum contract.
type SCRCDBO struct {
	SCRCDBOCaller     // Read-only binding to the contract
	SCRCDBOTransactor // Write-only binding to the contract
	SCRCDBOFilterer   // Log filterer for contract events
}

// SCRCDBOCaller is an auto generated read-only Go binding around an Ethereum contract.
type SCRCDBOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRCDBOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SCRCDBOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRCDBOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCRCDBOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRCDBOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCRCDBOSession struct {
	Contract     *SCRCDBO          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCRCDBOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCRCDBOCallerSession struct {
	Contract *SCRCDBOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SCRCDBOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCRCDBOTransactorSession struct {
	Contract     *SCRCDBOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SCRCDBORaw is an auto generated low-level Go binding around an Ethereum contract.
type SCRCDBORaw struct {
	Contract *SCRCDBO // Generic contract binding to access the raw methods on
}

// SCRCDBOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCRCDBOCallerRaw struct {
	Contract *SCRCDBOCaller // Generic read-only contract binding to access the raw methods on
}

// SCRCDBOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCRCDBOTransactorRaw struct {
	Contract *SCRCDBOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSCRCDBO creates a new instance of SCRCDBO, bound to a specific deployed contract.
func NewSCRCDBO(address common.Address, backend bind.ContractBackend) (*SCRCDBO, error) {
	contract, err := bindSCRCDBO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCRCDBO{SCRCDBOCaller: SCRCDBOCaller{contract: contract}, SCRCDBOTransactor: SCRCDBOTransactor{contract: contract}, SCRCDBOFilterer: SCRCDBOFilterer{contract: contract}}, nil
}

// NewSCRCDBOCaller creates a new read-only instance of SCRCDBO, bound to a specific deployed contract.
func NewSCRCDBOCaller(address common.Address, caller bind.ContractCaller) (*SCRCDBOCaller, error) {
	contract, err := bindSCRCDBO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCRCDBOCaller{contract: contract}, nil
}

// NewSCRCDBOTransactor creates a new write-only instance of SCRCDBO, bound to a specific deployed contract.
func NewSCRCDBOTransactor(address common.Address, transactor bind.ContractTransactor) (*SCRCDBOTransactor, error) {
	contract, err := bindSCRCDBO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCRCDBOTransactor{contract: contract}, nil
}

// NewSCRCDBOFilterer creates a new log filterer instance of SCRCDBO, bound to a specific deployed contract.
func NewSCRCDBOFilterer(address common.Address, filterer bind.ContractFilterer) (*SCRCDBOFilterer, error) {
	contract, err := bindSCRCDBO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCRCDBOFilterer{contract: contract}, nil
}

// bindSCRCDBO binds a generic wrapper to an already deployed contract.
func bindSCRCDBO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SCRCDBOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCRCDBO *SCRCDBORaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCRCDBO.Contract.SCRCDBOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCRCDBO *SCRCDBORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRCDBO.Contract.SCRCDBOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCRCDBO *SCRCDBORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCRCDBO.Contract.SCRCDBOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCRCDBO *SCRCDBOCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SCRCDBO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCRCDBO *SCRCDBOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRCDBO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCRCDBO *SCRCDBOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCRCDBO.Contract.contract.Transact(opts, method, params...)
}

// ResearchCommunityId is a free data retrieval call binding the contract method 0xa33ac38b.
//
// Solidity: function researchCommunityId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCaller) ResearchCommunityId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCRCDBO.contract.Call(opts, out, "researchCommunityId")
	return *ret0, err
}

// ResearchCommunityId is a free data retrieval call binding the contract method 0xa33ac38b.
//
// Solidity: function researchCommunityId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOSession) ResearchCommunityId() (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchCommunityId(&_SCRCDBO.CallOpts)
}

// ResearchCommunityId is a free data retrieval call binding the contract method 0xa33ac38b.
//
// Solidity: function researchCommunityId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCallerSession) ResearchCommunityId() (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchCommunityId(&_SCRCDBO.CallOpts)
}

// ResearchCommunityIdToAddress is a free data retrieval call binding the contract method 0x83401587.
//
// Solidity: function researchCommunityIdToAddress(uint256 ) constant returns(address)
func (_SCRCDBO *SCRCDBOCaller) ResearchCommunityIdToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SCRCDBO.contract.Call(opts, out, "researchCommunityIdToAddress", arg0)
	return *ret0, err
}

// ResearchCommunityIdToAddress is a free data retrieval call binding the contract method 0x83401587.
//
// Solidity: function researchCommunityIdToAddress(uint256 ) constant returns(address)
func (_SCRCDBO *SCRCDBOSession) ResearchCommunityIdToAddress(arg0 *big.Int) (common.Address, error) {
	return _SCRCDBO.Contract.ResearchCommunityIdToAddress(&_SCRCDBO.CallOpts, arg0)
}

// ResearchCommunityIdToAddress is a free data retrieval call binding the contract method 0x83401587.
//
// Solidity: function researchCommunityIdToAddress(uint256 ) constant returns(address)
func (_SCRCDBO *SCRCDBOCallerSession) ResearchCommunityIdToAddress(arg0 *big.Int) (common.Address, error) {
	return _SCRCDBO.Contract.ResearchCommunityIdToAddress(&_SCRCDBO.CallOpts, arg0)
}

// ResearchCommunityDatabaseOwnerResearchDataId is a free data retrieval call binding the contract method 0xa94717b8.
//
// Solidity: function researchCommunity_databaseOwner_researchDataId(address , address ) constant returns(uint256)
func (_SCRCDBO *SCRCDBOCaller) ResearchCommunityDatabaseOwnerResearchDataId(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCRCDBO.contract.Call(opts, out, "researchCommunity_databaseOwner_researchDataId", arg0, arg1)
	return *ret0, err
}

// ResearchCommunityDatabaseOwnerResearchDataId is a free data retrieval call binding the contract method 0xa94717b8.
//
// Solidity: function researchCommunity_databaseOwner_researchDataId(address , address ) constant returns(uint256)
func (_SCRCDBO *SCRCDBOSession) ResearchCommunityDatabaseOwnerResearchDataId(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchCommunityDatabaseOwnerResearchDataId(&_SCRCDBO.CallOpts, arg0, arg1)
}

// ResearchCommunityDatabaseOwnerResearchDataId is a free data retrieval call binding the contract method 0xa94717b8.
//
// Solidity: function researchCommunity_databaseOwner_researchDataId(address , address ) constant returns(uint256)
func (_SCRCDBO *SCRCDBOCallerSession) ResearchCommunityDatabaseOwnerResearchDataId(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchCommunityDatabaseOwnerResearchDataId(&_SCRCDBO.CallOpts, arg0, arg1)
}

// ResearchDataDetails is a free data retrieval call binding the contract method 0xb2d8e6a0.
//
// Solidity: function researchDataDetails(uint256 ) constant returns(uint256 id, address rcAddr, address dataBaseOwner, uint256 timeStamp_researchCommunityRequestData, uint256 timeStamp_dataBaseOwnerSendData, bytes32 hashOfMedData, bytes sigOnHashOfMedData)
func (_SCRCDBO *SCRCDBOCaller) ResearchDataDetails(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                                    *big.Int
	RcAddr                                common.Address
	DataBaseOwner                         common.Address
	TimeStampResearchCommunityRequestData *big.Int
	TimeStampDataBaseOwnerSendData        *big.Int
	HashOfMedData                         [32]byte
	SigOnHashOfMedData                    []byte
}, error) {
	ret := new(struct {
		Id                                    *big.Int
		RcAddr                                common.Address
		DataBaseOwner                         common.Address
		TimeStampResearchCommunityRequestData *big.Int
		TimeStampDataBaseOwnerSendData        *big.Int
		HashOfMedData                         [32]byte
		SigOnHashOfMedData                    []byte
	})
	out := ret
	err := _SCRCDBO.contract.Call(opts, out, "researchDataDetails", arg0)
	return *ret, err
}

// ResearchDataDetails is a free data retrieval call binding the contract method 0xb2d8e6a0.
//
// Solidity: function researchDataDetails(uint256 ) constant returns(uint256 id, address rcAddr, address dataBaseOwner, uint256 timeStamp_researchCommunityRequestData, uint256 timeStamp_dataBaseOwnerSendData, bytes32 hashOfMedData, bytes sigOnHashOfMedData)
func (_SCRCDBO *SCRCDBOSession) ResearchDataDetails(arg0 *big.Int) (struct {
	Id                                    *big.Int
	RcAddr                                common.Address
	DataBaseOwner                         common.Address
	TimeStampResearchCommunityRequestData *big.Int
	TimeStampDataBaseOwnerSendData        *big.Int
	HashOfMedData                         [32]byte
	SigOnHashOfMedData                    []byte
}, error) {
	return _SCRCDBO.Contract.ResearchDataDetails(&_SCRCDBO.CallOpts, arg0)
}

// ResearchDataDetails is a free data retrieval call binding the contract method 0xb2d8e6a0.
//
// Solidity: function researchDataDetails(uint256 ) constant returns(uint256 id, address rcAddr, address dataBaseOwner, uint256 timeStamp_researchCommunityRequestData, uint256 timeStamp_dataBaseOwnerSendData, bytes32 hashOfMedData, bytes sigOnHashOfMedData)
func (_SCRCDBO *SCRCDBOCallerSession) ResearchDataDetails(arg0 *big.Int) (struct {
	Id                                    *big.Int
	RcAddr                                common.Address
	DataBaseOwner                         common.Address
	TimeStampResearchCommunityRequestData *big.Int
	TimeStampDataBaseOwnerSendData        *big.Int
	HashOfMedData                         [32]byte
	SigOnHashOfMedData                    []byte
}, error) {
	return _SCRCDBO.Contract.ResearchDataDetails(&_SCRCDBO.CallOpts, arg0)
}

// ResearchDataDetailsId is a free data retrieval call binding the contract method 0xaf55a453.
//
// Solidity: function researchDataDetailsId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCaller) ResearchDataDetailsId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCRCDBO.contract.Call(opts, out, "researchDataDetailsId")
	return *ret0, err
}

// ResearchDataDetailsId is a free data retrieval call binding the contract method 0xaf55a453.
//
// Solidity: function researchDataDetailsId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOSession) ResearchDataDetailsId() (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchDataDetailsId(&_SCRCDBO.CallOpts)
}

// ResearchDataDetailsId is a free data retrieval call binding the contract method 0xaf55a453.
//
// Solidity: function researchDataDetailsId() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCallerSession) ResearchDataDetailsId() (*big.Int, error) {
	return _SCRCDBO.Contract.ResearchDataDetailsId(&_SCRCDBO.CallOpts)
}

// TimeLimitForGivingData is a free data retrieval call binding the contract method 0x6e8361be.
//
// Solidity: function timeLimitForGivingData() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCaller) TimeLimitForGivingData(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SCRCDBO.contract.Call(opts, out, "timeLimitForGivingData")
	return *ret0, err
}

// TimeLimitForGivingData is a free data retrieval call binding the contract method 0x6e8361be.
//
// Solidity: function timeLimitForGivingData() constant returns(uint256)
func (_SCRCDBO *SCRCDBOSession) TimeLimitForGivingData() (*big.Int, error) {
	return _SCRCDBO.Contract.TimeLimitForGivingData(&_SCRCDBO.CallOpts)
}

// TimeLimitForGivingData is a free data retrieval call binding the contract method 0x6e8361be.
//
// Solidity: function timeLimitForGivingData() constant returns(uint256)
func (_SCRCDBO *SCRCDBOCallerSession) TimeLimitForGivingData() (*big.Int, error) {
	return _SCRCDBO.Contract.TimeLimitForGivingData(&_SCRCDBO.CallOpts)
}

// ProvideMedicalData is a paid mutator transaction binding the contract method 0x47bbf530.
//
// Solidity: function provideMedicalData(address System_Users_Info_address, uint256 _researchDataId, bytes32 _hashOfMedData, bytes _sigOnHashOfMedData) returns()
func (_SCRCDBO *SCRCDBOTransactor) ProvideMedicalData(opts *bind.TransactOpts, System_Users_Info_address common.Address, _researchDataId *big.Int, _hashOfMedData [32]byte, _sigOnHashOfMedData []byte) (*types.Transaction, error) {
	return _SCRCDBO.contract.Transact(opts, "provideMedicalData", System_Users_Info_address, _researchDataId, _hashOfMedData, _sigOnHashOfMedData)
}

// ProvideMedicalData is a paid mutator transaction binding the contract method 0x47bbf530.
//
// Solidity: function provideMedicalData(address System_Users_Info_address, uint256 _researchDataId, bytes32 _hashOfMedData, bytes _sigOnHashOfMedData) returns()
func (_SCRCDBO *SCRCDBOSession) ProvideMedicalData(System_Users_Info_address common.Address, _researchDataId *big.Int, _hashOfMedData [32]byte, _sigOnHashOfMedData []byte) (*types.Transaction, error) {
	return _SCRCDBO.Contract.ProvideMedicalData(&_SCRCDBO.TransactOpts, System_Users_Info_address, _researchDataId, _hashOfMedData, _sigOnHashOfMedData)
}

// ProvideMedicalData is a paid mutator transaction binding the contract method 0x47bbf530.
//
// Solidity: function provideMedicalData(address System_Users_Info_address, uint256 _researchDataId, bytes32 _hashOfMedData, bytes _sigOnHashOfMedData) returns()
func (_SCRCDBO *SCRCDBOTransactorSession) ProvideMedicalData(System_Users_Info_address common.Address, _researchDataId *big.Int, _hashOfMedData [32]byte, _sigOnHashOfMedData []byte) (*types.Transaction, error) {
	return _SCRCDBO.Contract.ProvideMedicalData(&_SCRCDBO.TransactOpts, System_Users_Info_address, _researchDataId, _hashOfMedData, _sigOnHashOfMedData)
}

// RequestMedicalData is a paid mutator transaction binding the contract method 0x5ca57322.
//
// Solidity: function requestMedicalData(address System_Users_Info_address, address _dataBaseOwner) returns()
func (_SCRCDBO *SCRCDBOTransactor) RequestMedicalData(opts *bind.TransactOpts, System_Users_Info_address common.Address, _dataBaseOwner common.Address) (*types.Transaction, error) {
	return _SCRCDBO.contract.Transact(opts, "requestMedicalData", System_Users_Info_address, _dataBaseOwner)
}

// RequestMedicalData is a paid mutator transaction binding the contract method 0x5ca57322.
//
// Solidity: function requestMedicalData(address System_Users_Info_address, address _dataBaseOwner) returns()
func (_SCRCDBO *SCRCDBOSession) RequestMedicalData(System_Users_Info_address common.Address, _dataBaseOwner common.Address) (*types.Transaction, error) {
	return _SCRCDBO.Contract.RequestMedicalData(&_SCRCDBO.TransactOpts, System_Users_Info_address, _dataBaseOwner)
}

// RequestMedicalData is a paid mutator transaction binding the contract method 0x5ca57322.
//
// Solidity: function requestMedicalData(address System_Users_Info_address, address _dataBaseOwner) returns()
func (_SCRCDBO *SCRCDBOTransactorSession) RequestMedicalData(System_Users_Info_address common.Address, _dataBaseOwner common.Address) (*types.Transaction, error) {
	return _SCRCDBO.Contract.RequestMedicalData(&_SCRCDBO.TransactOpts, System_Users_Info_address, _dataBaseOwner)
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
