// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package access

import (
	"math/big"
	"strings"

	fiscobcos "github.com/chislab/go-fiscobcos"
	"github.com/chislab/go-fiscobcos/accounts/abi"
	"github.com/chislab/go-fiscobcos/accounts/abi/bind"
	"github.com/chislab/go-fiscobcos/common"
	"github.com/chislab/go-fiscobcos/common/hexutil"
	"github.com/chislab/go-fiscobcos/core/types"
	"github.com/chislab/go-fiscobcos/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = fiscobcos.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccessABI is the input ABI used to generate the binding from.
const AccessABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantMaterialProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantProductProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMaterialProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isProductProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessBin is the compiled bytecode used for deploying new contracts.
var AccessBin = "0x60806040526001600260006101000a81548160ff021916908360ff16021790555060028060016101000a81548160ff021916908360ff16021790555060046002806101000a81548160ff021916908360ff160217905550600061006661010960201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350610111565b600033905090565b610b99806101206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80639e20794b116100665780639e20794b1461010c578063a09409a514610128578063ad4a296d14610158578063afacca2e14610188578063f2fde38b146101a457610093565b80633448840b14610098578063715018a6146100b45780638da5cb5b146100be5780638f4d9565146100dc575b600080fd5b6100b260048036036100ad9190810190610a66565b6101c0565b005b6100bc610316565b005b6100c661046b565b6040516100d39190610a96565b60405180910390f35b6100f660048036036100f19190810190610a66565b610494565b6040516101039190610aaf565b60405180910390f35b61012660048036036101219190810190610a66565b610506565b005b610142600480360361013d9190810190610a66565b61065d565b60405161014f9190610aaf565b60405180910390f35b610172600480360361016d9190810190610a66565b6106d0565b60405161017f9190610aaf565b60405180910390f35b6101a2600480360361019d9190810190610a66565b610743565b005b6101be60048036036101b99190810190610a66565b61089a565b005b6101c8610a5e565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024d90610b26565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506002809054906101000a900460ff168118905080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505050565b61031e610a5e565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a390610b26565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905060006002809054906101000a900460ff16821660ff161415915050919050565b61050e610a5e565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461059c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059390610b26565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050600260009054906101000a900460ff168118905080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505050565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506000600260019054906101000a900460ff16821660ff161415915050919050565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690506000600260009054906101000a900460ff16821660ff161415915050919050565b61074b610a5e565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d090610b26565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050600260019054906101000a900460ff168118905080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505050565b6108a2610a5e565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610930576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092790610b26565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099790610ac2565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600060208284031215610a77578081fd5b813560018060a01b0381168114610a8c578182fd5b8091505092915050565b600060208201905060018060a01b038316825292915050565b6000602082019050821515825292915050565b600060208252602660208301527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408301527f64647265737300000000000000000000000000000000000000000000000000006060830152608082019050919050565b6000602082526020808301527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604083015260608201905091905056fea2646970667358221220fbcb310cdfeb5de3b5830ee69d0b81eb442365b7f39e3aec87c845abcc55303864736f6c63430006000033"

// DeployAccess deploys a new FiscoBcos contract, binding an instance of Access to it.
func DeployAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Access, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}, ABI: parsed}, nil
}

// Access is an auto generated Go binding around an FiscoBcos contract.
type Access struct {
	AccessCaller             // Read-only binding to the contract
	AccessTransactor         // Write-only binding to the contract
	AccessFilterer           // Log filterer for contract events
	ABI              abi.ABI // contract abi
}

// AccessCaller is an auto generated read-only Go binding around an FiscoBcos contract.
type AccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessTransactor is an auto generated write-only Go binding around an FiscoBcos contract.
type AccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessFilterer is an auto generated log filtering Go binding around an FiscoBcos contract events.
type AccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessSession is an auto generated Go binding around an FiscoBcos contract,
// with pre-set call and transact options.
type AccessSession struct {
	Contract     *Access           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessCallerSession is an auto generated read-only Go binding around an FiscoBcos contract,
// with pre-set call options.
type AccessCallerSession struct {
	Contract *AccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AccessTransactorSession is an auto generated write-only Go binding around an FiscoBcos contract,
// with pre-set transact options.
type AccessTransactorSession struct {
	Contract     *AccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessRaw is an auto generated low-level Go binding around an FiscoBcos contract.
type AccessRaw struct {
	Contract *Access // Generic contract binding to access the raw methods on
}

// AccessCallerRaw is an auto generated low-level read-only Go binding around an FiscoBcos contract.
type AccessCallerRaw struct {
	Contract *AccessCaller // Generic read-only contract binding to access the raw methods on
}

// AccessTransactorRaw is an auto generated low-level write-only Go binding around an FiscoBcos contract.
type AccessTransactorRaw struct {
	Contract *AccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccess creates a new instance of Access, bound to a specific deployed contract.
func NewAccess(address common.Address, backend bind.ContractBackend) (*Access, error) {
	contract, err := bindAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return nil, err
	}
	return &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}, ABI: parsed}, nil
}

// NewAccessCaller creates a new read-only instance of Access, bound to a specific deployed contract.
func NewAccessCaller(address common.Address, caller bind.ContractCaller) (*AccessCaller, error) {
	contract, err := bindAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessCaller{contract: contract}, nil
}

// NewAccessTransactor creates a new write-only instance of Access, bound to a specific deployed contract.
func NewAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessTransactor, error) {
	contract, err := bindAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessTransactor{contract: contract}, nil
}

// NewAccessFilterer creates a new log filterer instance of Access, bound to a specific deployed contract.
func NewAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessFilterer, error) {
	contract, err := bindAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessFilterer{contract: contract}, nil
}

// bindAccess binds a generic wrapper to an already deployed contract.
func bindAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Access.Contract.AccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Access.Contract.contract.Call(opts, result, method, params...)
}

func (_Access *AccessCallerRaw) ReadCall(result interface{}, method string, output []byte) error {
	return _Access.Contract.contract.ReadCall(result, method, output)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.contract.Transact(opts, method, params...)
}

// IsMaterialProducer is a free data retrieval call binding the contract method 0x8f4d9565.
//
// Solidity: function isMaterialProducer(address account) view returns(bool)
func (_Access *AccessCaller) IsMaterialProducer(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Access.contract.Call(opts, out, "isMaterialProducer", account)
	return *ret0, err
}

// IsMaterialProducer is a free data retrieval call binding the contract method 0x8f4d9565.
//
// Solidity: function isMaterialProducer(address account) view returns(bool)
func (_Access *AccessSession) IsMaterialProducer(account common.Address) (bool, error) {
	return _Access.Contract.IsMaterialProducer(&_Access.CallOpts, account)
}

// IsMaterialProducer is a free data retrieval call binding the contract method 0x8f4d9565.
//
// Solidity: function isMaterialProducer(address account) view returns(bool)
func (_Access *AccessCallerSession) IsMaterialProducer(account common.Address) (bool, error) {
	return _Access.Contract.IsMaterialProducer(&_Access.CallOpts, account)
}

// IsPayment is a free data retrieval call binding the contract method 0xad4a296d.
//
// Solidity: function isPayment(address account) view returns(bool)
func (_Access *AccessCaller) IsPayment(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Access.contract.Call(opts, out, "isPayment", account)
	return *ret0, err
}

// IsPayment is a free data retrieval call binding the contract method 0xad4a296d.
//
// Solidity: function isPayment(address account) view returns(bool)
func (_Access *AccessSession) IsPayment(account common.Address) (bool, error) {
	return _Access.Contract.IsPayment(&_Access.CallOpts, account)
}

// IsPayment is a free data retrieval call binding the contract method 0xad4a296d.
//
// Solidity: function isPayment(address account) view returns(bool)
func (_Access *AccessCallerSession) IsPayment(account common.Address) (bool, error) {
	return _Access.Contract.IsPayment(&_Access.CallOpts, account)
}

// IsProductProducer is a free data retrieval call binding the contract method 0xa09409a5.
//
// Solidity: function isProductProducer(address account) view returns(bool)
func (_Access *AccessCaller) IsProductProducer(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Access.contract.Call(opts, out, "isProductProducer", account)
	return *ret0, err
}

// IsProductProducer is a free data retrieval call binding the contract method 0xa09409a5.
//
// Solidity: function isProductProducer(address account) view returns(bool)
func (_Access *AccessSession) IsProductProducer(account common.Address) (bool, error) {
	return _Access.Contract.IsProductProducer(&_Access.CallOpts, account)
}

// IsProductProducer is a free data retrieval call binding the contract method 0xa09409a5.
//
// Solidity: function isProductProducer(address account) view returns(bool)
func (_Access *AccessCallerSession) IsProductProducer(account common.Address) (bool, error) {
	return _Access.Contract.IsProductProducer(&_Access.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Access.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessSession) Owner() (common.Address, error) {
	return _Access.Contract.Owner(&_Access.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessCallerSession) Owner() (common.Address, error) {
	return _Access.Contract.Owner(&_Access.CallOpts)
}

// GrantMaterialProducer is a paid mutator transaction binding the contract method 0x3448840b.
//
// Solidity: function grantMaterialProducer(address account) returns()

/******************************************************************************************************************************/
func (_Access *AccessCaller) ReadGrantMaterialProducer(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Access.contract.ReadCall(out, "grantMaterialProducer", hexutil.MustDecode(output))
	return err
}

// GrantMaterialProducer is a free data retrieval call binding the contract method 0x3448840b.
//
// Solidity: function grantMaterialProducer(address account) returns()
func (_Access *AccessSession) ReadGrantMaterialProducer(output string) error {
	return _Access.Contract.ReadGrantMaterialProducer(output)
}

// GrantMaterialProducer is a free data retrieval call binding the contract method 0x3448840b.
//
// Solidity: function grantMaterialProducer(address account) returns()
func (_Access *AccessCallerSession) ReadGrantMaterialProducer(output string) error {
	return _Access.Contract.ReadGrantMaterialProducer(output)
}

/******************************************************************************************************************************/

func (_Access *AccessTransactor) GrantMaterialProducer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "grantMaterialProducer", account)
}

// GrantMaterialProducer is a paid mutator transaction binding the contract method 0x3448840b.
//
// Solidity: function grantMaterialProducer(address account) returns()
func (_Access *AccessSession) GrantMaterialProducer(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantMaterialProducer(&_Access.TransactOpts, account)
}

// GrantMaterialProducer is a paid mutator transaction binding the contract method 0x3448840b.
//
// Solidity: function grantMaterialProducer(address account) returns()
func (_Access *AccessTransactorSession) GrantMaterialProducer(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantMaterialProducer(&_Access.TransactOpts, account)
}

// GrantPayment is a paid mutator transaction binding the contract method 0x9e20794b.
//
// Solidity: function grantPayment(address account) returns()

/******************************************************************************************************************************/
func (_Access *AccessCaller) ReadGrantPayment(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Access.contract.ReadCall(out, "grantPayment", hexutil.MustDecode(output))
	return err
}

// GrantPayment is a free data retrieval call binding the contract method 0x9e20794b.
//
// Solidity: function grantPayment(address account) returns()
func (_Access *AccessSession) ReadGrantPayment(output string) error {
	return _Access.Contract.ReadGrantPayment(output)
}

// GrantPayment is a free data retrieval call binding the contract method 0x9e20794b.
//
// Solidity: function grantPayment(address account) returns()
func (_Access *AccessCallerSession) ReadGrantPayment(output string) error {
	return _Access.Contract.ReadGrantPayment(output)
}

/******************************************************************************************************************************/

func (_Access *AccessTransactor) GrantPayment(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "grantPayment", account)
}

// GrantPayment is a paid mutator transaction binding the contract method 0x9e20794b.
//
// Solidity: function grantPayment(address account) returns()
func (_Access *AccessSession) GrantPayment(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantPayment(&_Access.TransactOpts, account)
}

// GrantPayment is a paid mutator transaction binding the contract method 0x9e20794b.
//
// Solidity: function grantPayment(address account) returns()
func (_Access *AccessTransactorSession) GrantPayment(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantPayment(&_Access.TransactOpts, account)
}

// GrantProductProducer is a paid mutator transaction binding the contract method 0xafacca2e.
//
// Solidity: function grantProductProducer(address account) returns()

/******************************************************************************************************************************/
func (_Access *AccessCaller) ReadGrantProductProducer(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Access.contract.ReadCall(out, "grantProductProducer", hexutil.MustDecode(output))
	return err
}

// GrantProductProducer is a free data retrieval call binding the contract method 0xafacca2e.
//
// Solidity: function grantProductProducer(address account) returns()
func (_Access *AccessSession) ReadGrantProductProducer(output string) error {
	return _Access.Contract.ReadGrantProductProducer(output)
}

// GrantProductProducer is a free data retrieval call binding the contract method 0xafacca2e.
//
// Solidity: function grantProductProducer(address account) returns()
func (_Access *AccessCallerSession) ReadGrantProductProducer(output string) error {
	return _Access.Contract.ReadGrantProductProducer(output)
}

/******************************************************************************************************************************/

func (_Access *AccessTransactor) GrantProductProducer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "grantProductProducer", account)
}

// GrantProductProducer is a paid mutator transaction binding the contract method 0xafacca2e.
//
// Solidity: function grantProductProducer(address account) returns()
func (_Access *AccessSession) GrantProductProducer(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantProductProducer(&_Access.TransactOpts, account)
}

// GrantProductProducer is a paid mutator transaction binding the contract method 0xafacca2e.
//
// Solidity: function grantProductProducer(address account) returns()
func (_Access *AccessTransactorSession) GrantProductProducer(account common.Address) (*types.Transaction, error) {
	return _Access.Contract.GrantProductProducer(&_Access.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()

/******************************************************************************************************************************/
func (_Access *AccessCaller) ReadRenounceOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Access.contract.ReadCall(out, "renounceOwnership", hexutil.MustDecode(output))
	return err
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessSession) ReadRenounceOwnership(output string) error {
	return _Access.Contract.ReadRenounceOwnership(output)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessCallerSession) ReadRenounceOwnership(output string) error {
	return _Access.Contract.ReadRenounceOwnership(output)
}

/******************************************************************************************************************************/

func (_Access *AccessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessSession) RenounceOwnership() (*types.Transaction, error) {
	return _Access.Contract.RenounceOwnership(&_Access.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Access.Contract.RenounceOwnership(&_Access.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()

/******************************************************************************************************************************/
func (_Access *AccessCaller) ReadTransferOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Access.contract.ReadCall(out, "transferOwnership", hexutil.MustDecode(output))
	return err
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessSession) ReadTransferOwnership(output string) error {
	return _Access.Contract.ReadTransferOwnership(output)
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessCallerSession) ReadTransferOwnership(output string) error {
	return _Access.Contract.ReadTransferOwnership(output)
}

/******************************************************************************************************************************/

func (_Access *AccessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Access.Contract.TransferOwnership(&_Access.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Access.Contract.TransferOwnership(&_Access.TransactOpts, newOwner)
}

// AccessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Access contract.
type AccessOwnershipTransferredIterator struct {
	Event *AccessOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log         // Log channel receiving the found contract events
	sub  fiscobcos.Subscription // Subscription for errors, completion and termination
	done bool                   // Whether the subscription completed delivering logs
	fail error                  // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessOwnershipTransferred)
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
		it.Event = new(AccessOwnershipTransferred)
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
func (it *AccessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessOwnershipTransferred represents a OwnershipTransferred event raised by the Access contract.
type AccessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Access *AccessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Access.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccessOwnershipTransferredIterator{contract: _Access.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Access *AccessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Access.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessOwnershipTransferred)
				if err := _Access.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Access *AccessFilterer) ParseOwnershipTransferred(log types.Log) (*AccessOwnershipTransferred, error) {
	event := new(AccessOwnershipTransferred)
	if err := _Access.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
