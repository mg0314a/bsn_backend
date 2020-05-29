// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payment

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Payer     common.Address
	Producer  common.Address
	Amount    *big.Int
	Count     *big.Int
	OrderType *big.Int
	CreatedAt *big.Int
	Status    uint8
}

// PaymentABI is the input ABI used to generate the binding from.
const PaymentABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cancelCompensate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount2Payer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount2Producer\",\"type\":\"uint256\"}],\"name\":\"EvtCancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"EvtConfirmOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cnt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EvtMakeOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EvtMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"confirmOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isMaterial\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"makeOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"setMaterialProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"setProductProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentBin is the compiled bytecode used for deploying new contracts.
var PaymentBin = "0x6080604052600060035560016004553480156200001b57600080fd5b5060405162002b8538038062002b85833981810160405262000041919081019062000107565b600062000053620000ff60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350806007819055505062000123565b600033905090565b60006020828403121562000119578081fd5b8151905092915050565b612a5280620001336000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063715018a611610071578063715018a6146101895780638ac7d79c146101935780638da5cb5b146101af5780639dc29fac146101cd578063d09ef241146101e9578063f2fde38b14610219576100b4565b806317903fb2146100b9578063218c1ad0146100d557806340c10f191461010557806346f6c13414610121578063514fcac71461013d57806370a0823114610159575b600080fd5b6100d360048036036100ce919081019061235a565b610235565b005b6100ef60048036036100ea91908101906123a9565b61030f565b6040516100fc91906129bc565b60405180910390f35b61011f600480360361011a919081019061237d565b610990565b005b61013b6004803603610136919081019061235a565b610b68565b005b6101576004803603610152919081019061240e565b610c42565b005b610173600480360361016e919081019061235a565b611397565b60405161018091906129bc565b60405180910390f35b6101916113e0565b005b6101ad60048036036101a8919081019061240e565b611535565b005b6101b7611b82565b6040516101c49190612444565b60405180910390f35b6101e760048036036101e2919081019061237d565b611bab565b005b61020360048036036101fe919081019061240e565b611dd4565b604051610210919061295f565b60405180910390f35b610233600480360361022e919081019061235a565b611f98565b005b61023d61215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c290612789565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415610380576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103779061258c565b60405180910390fd5b60008090506000915086156104eb57600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041890612803565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663449e815d87876040518363ffffffff1660e01b815260040161047e9291906124af565b60206040518083038186803b15801561049657600080fd5b505afa1580156104aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104ce9190810190612429565b905060035491506002600360008282540192505081905550610643565b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561057d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610574906126e8565b60405180910390fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663346b77f187876040518363ffffffff1660e01b81526004016105da9291906124af565b60206040518083038186803b1580156105f257600080fd5b505afa158015610606573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061062a9190810190612429565b9050600454915060026004600082825401925050819055505b80831015610686576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067d906126aa565b60405180910390fd5b600061069b858361216490919063ffffffff16565b905080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561071f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107169061266c565b60405180910390fd5b61077181600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546121d490919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506107bc6122ce565b6040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152602001838152602001878152602001888152602001428152602001600060ff168152509050806002600086815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff021916908360ff160217905550905050806020015173ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff16857fc396d3bd9db25f3dfb08141f250e40f2abdbafeb64064462fa7c7da4c07967168a8a8a604051610979939291906129fd565b60405180910390a483935050505095945050505050565b61099861215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1d90612789565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8d90612921565b60405180910390fd5b610ae881600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461221e90919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f1c33017a2c796f5ddf2292779d4d6ac19bde011ddd5466223fff9260195d056b8282604051610b5c9291906124af565b60405180910390a15050565b610b7061215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf590612789565b60405180910390fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610ce8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cdf906128e3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480610db957503373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610df8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610def90612528565b60405180910390fd5b60006002600083815260200190815260200160002060060160009054906101000a900460ff1660ff1614610e61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5890612867565b60405180910390fd5b600280600083815260200190815260200160002060060160006101000a81548160ff021916908360ff1602179055506002600082815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561106757610f986002600083815260200190815260200160002060020154600160006002600086815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461221e90919063ffffffff16565b600160006002600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550807fb6fdaefbdd5c945965ddcfafee4b401d1c386e5121a253cbcfc98525811359206002600084815260200190815260200160002060020154600060405161105a9291906129cd565b60405180910390a26112fc565b600060646007546002600085815260200190815260200160002060020154028161108d57fe5b04905060006110bb8260026000868152602001908152602001600020600201546121d490919063ffffffff16565b905061114582600160006002600088815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461221e90919063ffffffff16565b600160006002600087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061124681600160006002600088815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461221e90919063ffffffff16565b600160006002600087815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550827fb6fdaefbdd5c945965ddcfafee4b401d1c386e5121a253cbcfc985258113592082846040516112f19291906129e5565b60405180910390a250505b60026000828152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560028201600090556003820160009055600482016000905560058201600090556006820160006101000a81549060ff0219169055505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6113e861215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161146d90612789565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600073ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156115db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115d2906128e3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461167f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611676906127c6565b60405180910390fd5b60006002600083815260200190815260200160002060060160009054906101000a900460ff1660ff16146116e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116df90612867565b60405180910390fd5b60016002600083815260200190815260200160002060060160006101000a81548160ff021916908360ff160217905550600060018216141561185457600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f45f2ee06002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008681526020019081526020016000206004015460026000878152602001908152602001600020600301546040518563ffffffff1660e01b815260040161181d949392919061247d565b600060405180830381600087803b15801561183757600080fd5b505af115801561184b573d6000803e3d6000fd5b50505050611980565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b26f32466002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008681526020019081526020016000206004015460026000878152602001908152602001600020600301546040518563ffffffff1660e01b815260040161194d949392919061247d565b600060405180830381600087803b15801561196757600080fd5b505af115801561197b573d6000803e3d6000fd5b505050505b611a1e6002600083815260200190815260200160002060020154600160006002600086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461221e90919063ffffffff16565b600160006002600085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550807f8b93f1d6231a764f9d97601f8eb5b8374d5f314d9a905ae331bbc5106ef29016336002600085815260200190815260200160002060040154604051611adf92919061245d565b60405180910390a260026000828152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560028201600090556003820160009055600482016000905560058201600090556006820160006101000a81549060ff0219169055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611bb361215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611c41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3890612789565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611cb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ca8906128a5565b60405180910390fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115611d3b57600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b611d8d81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546121d490919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b611ddc6122ce565b600073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611e82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e79906128e3565b60405180910390fd5b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff1660ff1660ff16815250509050919050565b611fa061215c565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461202e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161202590612789565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561209e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612095906125ca565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b60008083141561217757600090506121ce565b600082840290508284828161218857fe5b04146121c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121c090612725565b60405180910390fd5b809150505b92915050565b600061221683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612273565b905092915050565b600080828401905083811015612269576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122609061262e565b60405180910390fd5b8091505092915050565b60008383111582906122bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b291906124cf565b60405180910390fd5b5060008385039050809150509392505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160008152602001600060ff1681525090565b60008135905060018060a01b038116811461235457600080fd5b92915050565b60006020828403121561236b578081fd5b612375838361233a565b905092915050565b6000806040838503121561238f578081fd5b612399848461233a565b9150602083013590509250929050565b600080600080600060a086880312156123c0578081fd5b853580151581146123cf578182fd5b80955050602086013560018060a01b03811681146123eb578182fd5b809450506040860135925060608601359150608086013590509295509295909350565b60006020828403121561241f578081fd5b8135905092915050565b60006020828403121561243a578081fd5b8151905092915050565b600060208201905060018060a01b038316825292915050565b600060408201905060018060a01b03841682528260208301529392505050565b600060808201905060018060a01b03808716835280861660208401525083604083015282606083015295945050505050565b600060408201905060018060a01b03841682528260208301529392505050565b6000602082528251806020840152815b818110156125005760208186010151604082860101526020810190506124df565b818111156125115782604083860101525b506040601f19601f83011684010191505092915050565b600060208252602c60208301527f6f6e6c792074686520706179657220616e642070726f64756365722063616e2060408301527f63616e63656c206f7264657200000000000000000000000000000000000000006060830152608082019050919050565b600060208252601d60208301527f63616e206e6f74206d616b65206f7264657220746f206e6f626f64792e0000006040830152606082019050919050565b600060208252602660208301527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408301527f64647265737300000000000000000000000000000000000000000000000000006060830152608082019050919050565b600060208252601b60208301527f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006040830152606082019050919050565b600060208252601460208301527f496e73756666696369656e742062616c616e63650000000000000000000000006040830152606082019050919050565b600060208252600e60208301527f7072696365206d69736d617463680000000000000000000000000000000000006040830152606082019050919050565b6000602082526020808301527f70726f6475636520636f6e74726163742061646472657373206e6f74207365746040830152606082019050919050565b600060208252602160208301527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60408301527f77000000000000000000000000000000000000000000000000000000000000006060830152608082019050919050565b6000602082526020808301527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726040830152606082019050919050565b6000602082526020808301527f6f6e6c79207468652070617965722063616e20636f6e6669726d206f726465726040830152606082019050919050565b600060208252602160208301527f6d6174657269616c20636f6e74726163742061646472657373206e6f7420736560408301527f74000000000000000000000000000000000000000000000000000000000000006060830152608082019050919050565b600060208252601260208301527f6f72646572207374617475732077726f6e6700000000000000000000000000006040830152606082019050919050565b600060208252601a60208301527f6275726e2066726f6d20746865207a65726f20616464726573730000000000006040830152606082019050919050565b600060208252601360208301527f6f7264657220646f6573206e6f742065786973000000000000000000000000006040830152606082019050919050565b600060208252601860208301527f6d696e7420746f20746865207a65726f206164647265737300000000000000006040830152606082019050919050565b600060e08201905060018060a01b038084511683528060208501511660208401525060408301516040830152606083015160608301526080830151608083015260a083015160a083015260ff60c08401511660c083015292915050565b600060208201905082825292915050565b60006040820190508382528260208301529392505050565b60006040820190508382528260208301529392505050565b600060608201905084825283602083015282604083015294935050505056fea2646970667358221220af6b2a3e1a7a5ba2b1b9dd3277d85aec9b247544a4c595bab8cfbe28f5b6ac0764736f6c63430006000033"

// DeployPayment deploys a new FiscoBcos contract, binding an instance of Payment to it.
func DeployPayment(auth *bind.TransactOpts, backend bind.ContractBackend, _cancelCompensate *big.Int) (common.Address, *types.Transaction, *Payment, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentBin), backend, _cancelCompensate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}, ABI: parsed}, nil
}

// Payment is an auto generated Go binding around an FiscoBcos contract.
type Payment struct {
	PaymentCaller             // Read-only binding to the contract
	PaymentTransactor         // Write-only binding to the contract
	PaymentFilterer           // Log filterer for contract events
	ABI               abi.ABI // contract abi
}

// PaymentCaller is an auto generated read-only Go binding around an FiscoBcos contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an FiscoBcos contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an FiscoBcos contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an FiscoBcos contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an FiscoBcos contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an FiscoBcos contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an FiscoBcos contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an FiscoBcos contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an FiscoBcos contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}, ABI: parsed}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

func (_Payment *PaymentCallerRaw) ReadCall(result interface{}, method string, output []byte) error {
	return _Payment.Contract.contract.ReadCall(result, method, output)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 amount)
func (_Payment *PaymentCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 amount)
func (_Payment *PaymentSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Payment.Contract.BalanceOf(&_Payment.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 amount)
func (_Payment *PaymentCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Payment.Contract.BalanceOf(&_Payment.CallOpts, account)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 id) view returns(Order order)
func (_Payment *PaymentCaller) GetOrder(opts *bind.CallOpts, id *big.Int) (Order, error) {
	var (
		ret0 = new(Order)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "getOrder", id)
	return *ret0, err
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 id) view returns(Order order)
func (_Payment *PaymentSession) GetOrder(id *big.Int) (Order, error) {
	return _Payment.Contract.GetOrder(&_Payment.CallOpts, id)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 id) view returns(Order order)
func (_Payment *PaymentCallerSession) GetOrder(id *big.Int) (Order, error) {
	return _Payment.Contract.GetOrder(&_Payment.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Payment.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCallerSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadBurn(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "burn", hexutil.MustDecode(output))
	return err
}

// Burn is a free data retrieval call binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Payment *PaymentSession) ReadBurn(output string) error {
	return _Payment.Contract.ReadBurn(output)
}

// Burn is a free data retrieval call binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Payment *PaymentCallerSession) ReadBurn(output string) error {
	return _Payment.Contract.ReadBurn(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Payment *PaymentSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Burn(&_Payment.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Payment *PaymentTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Burn(&_Payment.TransactOpts, account, amount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadCancelOrder(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "cancelOrder", hexutil.MustDecode(output))
	return err
}

// CancelOrder is a free data retrieval call binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Payment *PaymentSession) ReadCancelOrder(output string) error {
	return _Payment.Contract.ReadCancelOrder(output)
}

// CancelOrder is a free data retrieval call binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Payment *PaymentCallerSession) ReadCancelOrder(output string) error {
	return _Payment.Contract.ReadCancelOrder(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) CancelOrder(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "cancelOrder", id)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Payment *PaymentSession) CancelOrder(id *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.CancelOrder(&_Payment.TransactOpts, id)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Payment *PaymentTransactorSession) CancelOrder(id *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.CancelOrder(&_Payment.TransactOpts, id)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 id) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadConfirmOrder(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "confirmOrder", hexutil.MustDecode(output))
	return err
}

// ConfirmOrder is a free data retrieval call binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 id) returns()
func (_Payment *PaymentSession) ReadConfirmOrder(output string) error {
	return _Payment.Contract.ReadConfirmOrder(output)
}

// ConfirmOrder is a free data retrieval call binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 id) returns()
func (_Payment *PaymentCallerSession) ReadConfirmOrder(output string) error {
	return _Payment.Contract.ReadConfirmOrder(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) ConfirmOrder(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "confirmOrder", id)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 id) returns()
func (_Payment *PaymentSession) ConfirmOrder(id *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.ConfirmOrder(&_Payment.TransactOpts, id)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x8ac7d79c.
//
// Solidity: function confirmOrder(uint256 id) returns()
func (_Payment *PaymentTransactorSession) ConfirmOrder(id *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.ConfirmOrder(&_Payment.TransactOpts, id)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x218c1ad0.
//
// Solidity: function makeOrder(bool isMaterial, address to, uint256 orderType, uint256 _count, uint256 _price) returns(uint256 id)

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadMakeOrder(output string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payment.contract.ReadCall(out, "makeOrder", hexutil.MustDecode(output))
	return *ret0, err
}

// MakeOrder is a free data retrieval call binding the contract method 0x218c1ad0.
//
// Solidity: function makeOrder(bool isMaterial, address to, uint256 orderType, uint256 _count, uint256 _price) returns(uint256 id)
func (_Payment *PaymentSession) ReadMakeOrder(output string) (*big.Int, error) {
	return _Payment.Contract.ReadMakeOrder(output)
}

// MakeOrder is a free data retrieval call binding the contract method 0x218c1ad0.
//
// Solidity: function makeOrder(bool isMaterial, address to, uint256 orderType, uint256 _count, uint256 _price) returns(uint256 id)
func (_Payment *PaymentCallerSession) ReadMakeOrder(output string) (*big.Int, error) {
	return _Payment.Contract.ReadMakeOrder(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) MakeOrder(opts *bind.TransactOpts, isMaterial bool, to common.Address, orderType *big.Int, _count *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "makeOrder", isMaterial, to, orderType, _count, _price)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x218c1ad0.
//
// Solidity: function makeOrder(bool isMaterial, address to, uint256 orderType, uint256 _count, uint256 _price) returns(uint256 id)
func (_Payment *PaymentSession) MakeOrder(isMaterial bool, to common.Address, orderType *big.Int, _count *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.MakeOrder(&_Payment.TransactOpts, isMaterial, to, orderType, _count, _price)
}

// MakeOrder is a paid mutator transaction binding the contract method 0x218c1ad0.
//
// Solidity: function makeOrder(bool isMaterial, address to, uint256 orderType, uint256 _count, uint256 _price) returns(uint256 id)
func (_Payment *PaymentTransactorSession) MakeOrder(isMaterial bool, to common.Address, orderType *big.Int, _count *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.MakeOrder(&_Payment.TransactOpts, isMaterial, to, orderType, _count, _price)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadMint(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "mint", hexutil.MustDecode(output))
	return err
}

// Mint is a free data retrieval call binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Payment *PaymentSession) ReadMint(output string) error {
	return _Payment.Contract.ReadMint(output)
}

// Mint is a free data retrieval call binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Payment *PaymentCallerSession) ReadMint(output string) error {
	return _Payment.Contract.ReadMint(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Payment *PaymentSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Mint(&_Payment.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Payment *PaymentTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Mint(&_Payment.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadRenounceOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "renounceOwnership", hexutil.MustDecode(output))
	return err
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payment *PaymentSession) ReadRenounceOwnership(output string) error {
	return _Payment.Contract.ReadRenounceOwnership(output)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payment *PaymentCallerSession) ReadRenounceOwnership(output string) error {
	return _Payment.Contract.ReadRenounceOwnership(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payment *PaymentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payment.Contract.RenounceOwnership(&_Payment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payment *PaymentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payment.Contract.RenounceOwnership(&_Payment.TransactOpts)
}

// SetMaterialProducer is a paid mutator transaction binding the contract method 0x46f6c134.
//
// Solidity: function setMaterialProducer(address producer) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadSetMaterialProducer(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "setMaterialProducer", hexutil.MustDecode(output))
	return err
}

// SetMaterialProducer is a free data retrieval call binding the contract method 0x46f6c134.
//
// Solidity: function setMaterialProducer(address producer) returns()
func (_Payment *PaymentSession) ReadSetMaterialProducer(output string) error {
	return _Payment.Contract.ReadSetMaterialProducer(output)
}

// SetMaterialProducer is a free data retrieval call binding the contract method 0x46f6c134.
//
// Solidity: function setMaterialProducer(address producer) returns()
func (_Payment *PaymentCallerSession) ReadSetMaterialProducer(output string) error {
	return _Payment.Contract.ReadSetMaterialProducer(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) SetMaterialProducer(opts *bind.TransactOpts, producer common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setMaterialProducer", producer)
}

// SetMaterialProducer is a paid mutator transaction binding the contract method 0x46f6c134.
//
// Solidity: function setMaterialProducer(address producer) returns()
func (_Payment *PaymentSession) SetMaterialProducer(producer common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetMaterialProducer(&_Payment.TransactOpts, producer)
}

// SetMaterialProducer is a paid mutator transaction binding the contract method 0x46f6c134.
//
// Solidity: function setMaterialProducer(address producer) returns()
func (_Payment *PaymentTransactorSession) SetMaterialProducer(producer common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetMaterialProducer(&_Payment.TransactOpts, producer)
}

// SetProductProducer is a paid mutator transaction binding the contract method 0x17903fb2.
//
// Solidity: function setProductProducer(address producer) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadSetProductProducer(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "setProductProducer", hexutil.MustDecode(output))
	return err
}

// SetProductProducer is a free data retrieval call binding the contract method 0x17903fb2.
//
// Solidity: function setProductProducer(address producer) returns()
func (_Payment *PaymentSession) ReadSetProductProducer(output string) error {
	return _Payment.Contract.ReadSetProductProducer(output)
}

// SetProductProducer is a free data retrieval call binding the contract method 0x17903fb2.
//
// Solidity: function setProductProducer(address producer) returns()
func (_Payment *PaymentCallerSession) ReadSetProductProducer(output string) error {
	return _Payment.Contract.ReadSetProductProducer(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) SetProductProducer(opts *bind.TransactOpts, producer common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setProductProducer", producer)
}

// SetProductProducer is a paid mutator transaction binding the contract method 0x17903fb2.
//
// Solidity: function setProductProducer(address producer) returns()
func (_Payment *PaymentSession) SetProductProducer(producer common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetProductProducer(&_Payment.TransactOpts, producer)
}

// SetProductProducer is a paid mutator transaction binding the contract method 0x17903fb2.
//
// Solidity: function setProductProducer(address producer) returns()
func (_Payment *PaymentTransactorSession) SetProductProducer(producer common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetProductProducer(&_Payment.TransactOpts, producer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()

/******************************************************************************************************************************/
func (_Payment *PaymentCaller) ReadTransferOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Payment.contract.ReadCall(out, "transferOwnership", hexutil.MustDecode(output))
	return err
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentSession) ReadTransferOwnership(output string) error {
	return _Payment.Contract.ReadTransferOwnership(output)
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentCallerSession) ReadTransferOwnership(output string) error {
	return _Payment.Contract.ReadTransferOwnership(output)
}

/******************************************************************************************************************************/

func (_Payment *PaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// PaymentEvtCancelOrderIterator is returned from FilterEvtCancelOrder and is used to iterate over the raw logs and unpacked data for EvtCancelOrder events raised by the Payment contract.
type PaymentEvtCancelOrderIterator struct {
	Event *PaymentEvtCancelOrder // Event containing the contract specifics and raw log

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
func (it *PaymentEvtCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEvtCancelOrder)
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
		it.Event = new(PaymentEvtCancelOrder)
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
func (it *PaymentEvtCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEvtCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEvtCancelOrder represents a EvtCancelOrder event raised by the Payment contract.
type PaymentEvtCancelOrder struct {
	Id              *big.Int
	Amount2Payer    *big.Int
	Amount2Producer *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEvtCancelOrder is a free log retrieval operation binding the contract event 0xb6fdaefbdd5c945965ddcfafee4b401d1c386e5121a253cbcfc9852581135920.
//
// Solidity: event EvtCancelOrder(uint256 indexed id, uint256 amount2Payer, uint256 amount2Producer)
func (_Payment *PaymentFilterer) FilterEvtCancelOrder(opts *bind.FilterOpts, id []*big.Int) (*PaymentEvtCancelOrderIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EvtCancelOrder", idRule)
	if err != nil {
		return nil, err
	}
	return &PaymentEvtCancelOrderIterator{contract: _Payment.contract, event: "EvtCancelOrder", logs: logs, sub: sub}, nil
}

// WatchEvtCancelOrder is a free log subscription operation binding the contract event 0xb6fdaefbdd5c945965ddcfafee4b401d1c386e5121a253cbcfc9852581135920.
//
// Solidity: event EvtCancelOrder(uint256 indexed id, uint256 amount2Payer, uint256 amount2Producer)
func (_Payment *PaymentFilterer) WatchEvtCancelOrder(opts *bind.WatchOpts, sink chan<- *PaymentEvtCancelOrder, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EvtCancelOrder", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEvtCancelOrder)
				if err := _Payment.contract.UnpackLog(event, "EvtCancelOrder", log); err != nil {
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

// ParseEvtCancelOrder is a log parse operation binding the contract event 0xb6fdaefbdd5c945965ddcfafee4b401d1c386e5121a253cbcfc9852581135920.
//
// Solidity: event EvtCancelOrder(uint256 indexed id, uint256 amount2Payer, uint256 amount2Producer)
func (_Payment *PaymentFilterer) ParseEvtCancelOrder(log types.Log) (*PaymentEvtCancelOrder, error) {
	event := new(PaymentEvtCancelOrder)
	if err := _Payment.contract.UnpackLog(event, "EvtCancelOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentEvtConfirmOrderIterator is returned from FilterEvtConfirmOrder and is used to iterate over the raw logs and unpacked data for EvtConfirmOrder events raised by the Payment contract.
type PaymentEvtConfirmOrderIterator struct {
	Event *PaymentEvtConfirmOrder // Event containing the contract specifics and raw log

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
func (it *PaymentEvtConfirmOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEvtConfirmOrder)
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
		it.Event = new(PaymentEvtConfirmOrder)
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
func (it *PaymentEvtConfirmOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEvtConfirmOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEvtConfirmOrder represents a EvtConfirmOrder event raised by the Payment contract.
type PaymentEvtConfirmOrder struct {
	From      common.Address
	OrderType *big.Int
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvtConfirmOrder is a free log retrieval operation binding the contract event 0x8b93f1d6231a764f9d97601f8eb5b8374d5f314d9a905ae331bbc5106ef29016.
//
// Solidity: event EvtConfirmOrder(address from, uint256 orderType, uint256 indexed id)
func (_Payment *PaymentFilterer) FilterEvtConfirmOrder(opts *bind.FilterOpts, id []*big.Int) (*PaymentEvtConfirmOrderIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EvtConfirmOrder", idRule)
	if err != nil {
		return nil, err
	}
	return &PaymentEvtConfirmOrderIterator{contract: _Payment.contract, event: "EvtConfirmOrder", logs: logs, sub: sub}, nil
}

// WatchEvtConfirmOrder is a free log subscription operation binding the contract event 0x8b93f1d6231a764f9d97601f8eb5b8374d5f314d9a905ae331bbc5106ef29016.
//
// Solidity: event EvtConfirmOrder(address from, uint256 orderType, uint256 indexed id)
func (_Payment *PaymentFilterer) WatchEvtConfirmOrder(opts *bind.WatchOpts, sink chan<- *PaymentEvtConfirmOrder, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EvtConfirmOrder", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEvtConfirmOrder)
				if err := _Payment.contract.UnpackLog(event, "EvtConfirmOrder", log); err != nil {
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

// ParseEvtConfirmOrder is a log parse operation binding the contract event 0x8b93f1d6231a764f9d97601f8eb5b8374d5f314d9a905ae331bbc5106ef29016.
//
// Solidity: event EvtConfirmOrder(address from, uint256 orderType, uint256 indexed id)
func (_Payment *PaymentFilterer) ParseEvtConfirmOrder(log types.Log) (*PaymentEvtConfirmOrder, error) {
	event := new(PaymentEvtConfirmOrder)
	if err := _Payment.contract.UnpackLog(event, "EvtConfirmOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentEvtMakeOrderIterator is returned from FilterEvtMakeOrder and is used to iterate over the raw logs and unpacked data for EvtMakeOrder events raised by the Payment contract.
type PaymentEvtMakeOrderIterator struct {
	Event *PaymentEvtMakeOrder // Event containing the contract specifics and raw log

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
func (it *PaymentEvtMakeOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEvtMakeOrder)
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
		it.Event = new(PaymentEvtMakeOrder)
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
func (it *PaymentEvtMakeOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEvtMakeOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEvtMakeOrder represents a EvtMakeOrder event raised by the Payment contract.
type PaymentEvtMakeOrder struct {
	OrderType *big.Int
	Id        *big.Int
	Payer     common.Address
	Producer  common.Address
	Cnt       *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvtMakeOrder is a free log retrieval operation binding the contract event 0xc396d3bd9db25f3dfb08141f250e40f2abdbafeb64064462fa7c7da4c0796716.
//
// Solidity: event EvtMakeOrder(uint256 orderType, uint256 indexed id, address indexed payer, address indexed producer, uint256 cnt, uint256 price)
func (_Payment *PaymentFilterer) FilterEvtMakeOrder(opts *bind.FilterOpts, id []*big.Int, payer []common.Address, producer []common.Address) (*PaymentEvtMakeOrderIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var producerRule []interface{}
	for _, producerItem := range producer {
		producerRule = append(producerRule, producerItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EvtMakeOrder", idRule, payerRule, producerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentEvtMakeOrderIterator{contract: _Payment.contract, event: "EvtMakeOrder", logs: logs, sub: sub}, nil
}

// WatchEvtMakeOrder is a free log subscription operation binding the contract event 0xc396d3bd9db25f3dfb08141f250e40f2abdbafeb64064462fa7c7da4c0796716.
//
// Solidity: event EvtMakeOrder(uint256 orderType, uint256 indexed id, address indexed payer, address indexed producer, uint256 cnt, uint256 price)
func (_Payment *PaymentFilterer) WatchEvtMakeOrder(opts *bind.WatchOpts, sink chan<- *PaymentEvtMakeOrder, id []*big.Int, payer []common.Address, producer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var producerRule []interface{}
	for _, producerItem := range producer {
		producerRule = append(producerRule, producerItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EvtMakeOrder", idRule, payerRule, producerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEvtMakeOrder)
				if err := _Payment.contract.UnpackLog(event, "EvtMakeOrder", log); err != nil {
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

// ParseEvtMakeOrder is a log parse operation binding the contract event 0xc396d3bd9db25f3dfb08141f250e40f2abdbafeb64064462fa7c7da4c0796716.
//
// Solidity: event EvtMakeOrder(uint256 orderType, uint256 indexed id, address indexed payer, address indexed producer, uint256 cnt, uint256 price)
func (_Payment *PaymentFilterer) ParseEvtMakeOrder(log types.Log) (*PaymentEvtMakeOrder, error) {
	event := new(PaymentEvtMakeOrder)
	if err := _Payment.contract.UnpackLog(event, "EvtMakeOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentEvtMintIterator is returned from FilterEvtMint and is used to iterate over the raw logs and unpacked data for EvtMint events raised by the Payment contract.
type PaymentEvtMintIterator struct {
	Event *PaymentEvtMint // Event containing the contract specifics and raw log

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
func (it *PaymentEvtMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentEvtMint)
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
		it.Event = new(PaymentEvtMint)
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
func (it *PaymentEvtMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentEvtMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentEvtMint represents a EvtMint event raised by the Payment contract.
type PaymentEvtMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvtMint is a free log retrieval operation binding the contract event 0x1c33017a2c796f5ddf2292779d4d6ac19bde011ddd5466223fff9260195d056b.
//
// Solidity: event EvtMint(address account, uint256 amount)
func (_Payment *PaymentFilterer) FilterEvtMint(opts *bind.FilterOpts) (*PaymentEvtMintIterator, error) {

	logs, sub, err := _Payment.contract.FilterLogs(opts, "EvtMint")
	if err != nil {
		return nil, err
	}
	return &PaymentEvtMintIterator{contract: _Payment.contract, event: "EvtMint", logs: logs, sub: sub}, nil
}

// WatchEvtMint is a free log subscription operation binding the contract event 0x1c33017a2c796f5ddf2292779d4d6ac19bde011ddd5466223fff9260195d056b.
//
// Solidity: event EvtMint(address account, uint256 amount)
func (_Payment *PaymentFilterer) WatchEvtMint(opts *bind.WatchOpts, sink chan<- *PaymentEvtMint) (event.Subscription, error) {

	logs, sub, err := _Payment.contract.WatchLogs(opts, "EvtMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentEvtMint)
				if err := _Payment.contract.UnpackLog(event, "EvtMint", log); err != nil {
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

// ParseEvtMint is a log parse operation binding the contract event 0x1c33017a2c796f5ddf2292779d4d6ac19bde011ddd5466223fff9260195d056b.
//
// Solidity: event EvtMint(address account, uint256 amount)
func (_Payment *PaymentFilterer) ParseEvtMint(log types.Log) (*PaymentEvtMint, error) {
	event := new(PaymentEvtMint)
	if err := _Payment.contract.UnpackLog(event, "EvtMint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Payment contract.
type PaymentOwnershipTransferredIterator struct {
	Event *PaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentOwnershipTransferred)
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
		it.Event = new(PaymentOwnershipTransferred)
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
func (it *PaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Payment contract.
type PaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payment *PaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentOwnershipTransferredIterator{contract: _Payment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payment *PaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentOwnershipTransferred)
				if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Payment *PaymentFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentOwnershipTransferred, error) {
	event := new(PaymentOwnershipTransferred)
	if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
