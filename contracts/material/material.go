// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package material

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

// materialRawMaterial is an auto generated low-level Go binding around an user-defined struct.
type materialRawMaterial struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}

// MaterialABI is the input ABI used to generate the binding from.
const MaterialABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"EvtMaterialConsumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"EvtMaterialCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"EvtMaterialTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EvtPriceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batchInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keptNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"consumeMaterial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"}],\"name\":\"getMaterialProducer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"}],\"name\":\"getMyMaterial\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"keptMaterials\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keptNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"newMaterial\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keptNum\",\"type\":\"uint256\"}],\"internalType\":\"structmaterial.RawMaterial\",\"name\":\"m\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchID\",\"type\":\"uint256\"}],\"name\":\"showBatchInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keptNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"transferMaterial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"updateProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedBatchIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MaterialBin is the compiled bytecode used for deploying new contracts.
var MaterialBin = "0x608060405234801561001057600080fd5b5060405161206a38038061206a83398181016040526100329190810190610079565b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100a9565b60006020828403121561008a578081fd5b815160018060a01b038116811461009f578182fd5b8091505092915050565b611fb2806100b86000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063b00351d511610071578063b00351d5146101ca578063c4e30a2d146101fa578063dc8215621461022f578063e0159f8e1461024b578063f45f2ee014610280578063f7d975771461029c576100b4565b806314c7b2f8146100b9578063449e815d146100ee57806348640e331461011e578063612e324a1461013a57806373cbde551461016a5780638c971e361461019a575b600080fd5b6100d360048036036100ce9190810190611bcd565b6102b8565b6040516100e596959493929190611cf0565b60405180910390f35b61010860048036036101039190810190611b42565b6103a5565b6040516101159190611f6b565b60405180910390f35b61013860048036036101339190810190611be8565b610400565b005b610154600480360361014f9190810190611b42565b610456565b6040516101619190611f6b565b60405180910390f35b610184600480360361017f9190810190611bcd565b61047b565b6040516101919190611f6b565b60405180910390f35b6101b460048036036101af9190810190611bcd565b610647565b6040516101c19190611c65565b60405180910390f35b6101e460048036036101df9190810190611c38565b610684565b6040516101f19190611f1f565b60405180910390f35b610214600480360361020f9190810190611bcd565b610985565b60405161022696959493929190611cf0565b60405180910390f35b61024960048036036102449190810190611c14565b6109e1565b005b61026560048036036102609190810190611b6e565b610e24565b60405161027796959493929190611cf0565b60405180910390f35b61029a60048036036102959190810190611afb565b610ea7565b005b6102b660048036036102b19190810190611c14565b6117c5565b005b6000806000806000806102c9611a8f565b600360008981526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152505090508060000151816020015182604001518360600151846080015160008090509650965096509650965096505091939550919395565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b806004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000602052816000526040600020602052806000526040600020600091509150505481565b60006060600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156105a957838290600052602060002090600602016040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481525050815260200190600101906104ef565b50505050905060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000205490505b815181101561063d5761062e8383838151811061061d57fe5b602002602001015160a00151611995565b92508080600101915050610604565b8292505050919050565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b61068c611a8f565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638f4d9565336040518263ffffffff1660e01b81526004016106e79190611c7e565b60206040518083038186803b1580156106ff57600080fd5b505afa158015610713573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107379190810190611ba3565b610776576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076d90611e01565b60405180910390fd5b6040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001428152602001838152602001858152602001848152602001848152509050806003600084815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050155905050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015550507fe1706a40c6f84da3bd18ed43b695c0d62afa92863e79c212de31eda5b77583d833858560405161097393929190611c97565b60405180910390a18090509392505050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154908060040154908060050154905086565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020805490506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000205410610ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abc90611e3f565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000205481548110610b8657fe5b906000526020600020906006020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0690611d85565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000205490506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000208281548110610cc057fe5b9060005260206000209060060201905082816005015410610d6357610ce98160050154846119ea565b8160050181905550600081600501541415610d5e576000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600081548092919060010191905055505b610de3565b6000816005015484039050600082600501819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060008154809291906001019190505550610de185826109e1565b505b7fbb29a9e0a07ed37fc0d005182fb0c11a1fae309a397dfb43b0ec7af31caa90dc338585604051610e1693929190611c97565b60405180910390a150505050565b60016020528260005260406000206020528160005260406000208181548110610e4957fe5b906000526020600020906006020160009250925050508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154908060040154908060050154905086565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad4a296d336040518263ffffffff1660e01b8152600401610f029190611c7e565b60206040518083038186803b158015610f1a57600080fd5b505afa158015610f2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f529190810190611ba3565b610f91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8890611ee1565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611000576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff790611e7d565b60405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020805490506000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054106110e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110db90611e3f565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000205490506000600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020828154811061119557fe5b906000526020600020906006020190506000600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002080549050141561137b57611209611a8f565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050816001015481602001818152505081600201548160400181815250508160030154816060018181525050816004015481608001818152505060008160a0018181525050600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a082015181600501555050505b600060018060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000878152602001908152602001600020805490500390506000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000878152602001908152602001600020828154811061143257fe5b90600052602060002090600602019050826002015481600201541461163957611459611a8f565b88816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050836001015481602001818152505083600201548160400181815250508360030154816060018181525050836004015481608001818152505060008160a0018181525050600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600088815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a082015181600501555050600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000888152602001908152602001600020600184018154811061162757fe5b90600052602060002090600602019150505b848360050154106116e2576116528360050154866119ea565b8360050181905550611668816005015486611995565b81600501819055506000836005015414156116dd576000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000878152602001908152602001600020600081548092919060010191905055505b61177e565b60008360050154860390506116ff82600501548560050154611995565b8260050181905550600084600501819055506000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008881526020019081526020016000206000815480929190600101919050555061177c89898984610ea7565b505b7f70d9b7ac16a9bb81749f2f28832acbf5d2a70bc09695163d1de16338b34bca10888888886040516117b39493929190611cbe565b60405180910390a15050505050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638f4d9565336040518263ffffffff1660e01b81526004016118209190611c7e565b60206040518083038186803b15801561183857600080fd5b505afa15801561184c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506118709190810190611ba3565b6118af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118a690611e01565b60405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002081905550336004600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3061c2e7e7cfec0710fd23c1e44a860a3319b77dc3bbf3283a2a6448f372e60d33838360405161198993929190611c97565b60405180910390a15050565b6000808284019050838110156119e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d790611dc3565b60405180910390fd5b8091505092915050565b6000611a2c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611a34565b905092915050565b6000838311158290611a7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a739190611d2c565b60405180910390fd5b5060008385039050809150509392505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160008152602001600081525090565b60008135905060018060a01b0381168114611af557600080fd5b92915050565b60008060008060808587031215611b10578384fd5b611b1a8686611adb565b9350611b298660208701611adb565b9250604085013591506060850135905092959194509250565b60008060408385031215611b54578182fd5b611b5e8484611adb565b9150602083013590509250929050565b600080600060608486031215611b82578283fd5b611b8c8585611adb565b925060208401359150604084013590509250925092565b600060208284031215611bb4578081fd5b81518015158114611bc3578182fd5b8091505092915050565b600060208284031215611bde578081fd5b8135905092915050565b60008060408385031215611bfa578182fd5b82359150611c0b8460208501611adb565b90509250929050565b60008060408385031215611c26578182fd5b82359150602083013590509250929050565b600080600060608486031215611c4c578283fd5b8335925060208401359150604084013590509250925092565b600060208201905060018060a01b038316825292915050565b600060208201905060018060a01b038316825292915050565b600060608201905060018060a01b0385168252836020830152826040830152949350505050565b600060808201905060018060a01b03808716835280861660208401525083604083015282606083015295945050505050565b600060c08201905060018060a01b03881682528660208301528560408301528460608301528360808301528260a0830152979650505050505050565b6000602082528251806020840152815b81811015611d5d576020818601015160408286010152602081019050611d3c565b81811115611d6e5782604083860101525b506040601f19601f83011684010191505092915050565b600060208252601b60208301527f70726f64756365722063616e6e6f7420626520636f6e73756d657200000000006040830152606082019050919050565b600060208252601b60208301527f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006040830152606082019050919050565b600060208252601a60208301527f6f6e6c7920666f72206d6174657269616c2070726f64756365720000000000006040830152606082019050919050565b600060208252601760208301527f696e73756666696369656e74206d6174657269616c732e0000000000000000006040830152606082019050919050565b600060208252602460208301527f7472616e7366657220746f20612073616d652067757920697320666f7262696460408301527f64656e2e000000000000000000000000000000000000000000000000000000006060830152608082019050919050565b600060208252601060208301527f6f6e6c7920666f72207061796d656e74000000000000000000000000000000006040830152606082019050919050565b600060c08201905060018060a01b0383511682526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015292915050565b60006020820190508282529291505056fea26469706673582212203feaf3f21fa3316c0e24d9980e9afc764634350dec73cd8b90c2e04e06a9966f64736f6c63430006000033"

// DeployMaterial deploys a new FiscoBcos contract, binding an instance of Material to it.
func DeployMaterial(auth *bind.TransactOpts, backend bind.ContractBackend, accessContractAddress common.Address) (common.Address, *types.Transaction, *Material, error) {
	parsed, err := abi.JSON(strings.NewReader(MaterialABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MaterialBin), backend, accessContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Material{MaterialCaller: MaterialCaller{contract: contract}, MaterialTransactor: MaterialTransactor{contract: contract}, MaterialFilterer: MaterialFilterer{contract: contract}, ABI: parsed}, nil
}

// Material is an auto generated Go binding around an FiscoBcos contract.
type Material struct {
	MaterialCaller             // Read-only binding to the contract
	MaterialTransactor         // Write-only binding to the contract
	MaterialFilterer           // Log filterer for contract events
	ABI                abi.ABI // contract abi
}

// MaterialCaller is an auto generated read-only Go binding around an FiscoBcos contract.
type MaterialCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaterialTransactor is an auto generated write-only Go binding around an FiscoBcos contract.
type MaterialTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaterialFilterer is an auto generated log filtering Go binding around an FiscoBcos contract events.
type MaterialFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaterialSession is an auto generated Go binding around an FiscoBcos contract,
// with pre-set call and transact options.
type MaterialSession struct {
	Contract     *Material         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaterialCallerSession is an auto generated read-only Go binding around an FiscoBcos contract,
// with pre-set call options.
type MaterialCallerSession struct {
	Contract *MaterialCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MaterialTransactorSession is an auto generated write-only Go binding around an FiscoBcos contract,
// with pre-set transact options.
type MaterialTransactorSession struct {
	Contract     *MaterialTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MaterialRaw is an auto generated low-level Go binding around an FiscoBcos contract.
type MaterialRaw struct {
	Contract *Material // Generic contract binding to access the raw methods on
}

// MaterialCallerRaw is an auto generated low-level read-only Go binding around an FiscoBcos contract.
type MaterialCallerRaw struct {
	Contract *MaterialCaller // Generic read-only contract binding to access the raw methods on
}

// MaterialTransactorRaw is an auto generated low-level write-only Go binding around an FiscoBcos contract.
type MaterialTransactorRaw struct {
	Contract *MaterialTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaterial creates a new instance of Material, bound to a specific deployed contract.
func NewMaterial(address common.Address, backend bind.ContractBackend) (*Material, error) {
	contract, err := bindMaterial(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(MaterialABI))
	if err != nil {
		return nil, err
	}
	return &Material{MaterialCaller: MaterialCaller{contract: contract}, MaterialTransactor: MaterialTransactor{contract: contract}, MaterialFilterer: MaterialFilterer{contract: contract}, ABI: parsed}, nil
}

// NewMaterialCaller creates a new read-only instance of Material, bound to a specific deployed contract.
func NewMaterialCaller(address common.Address, caller bind.ContractCaller) (*MaterialCaller, error) {
	contract, err := bindMaterial(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaterialCaller{contract: contract}, nil
}

// NewMaterialTransactor creates a new write-only instance of Material, bound to a specific deployed contract.
func NewMaterialTransactor(address common.Address, transactor bind.ContractTransactor) (*MaterialTransactor, error) {
	contract, err := bindMaterial(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaterialTransactor{contract: contract}, nil
}

// NewMaterialFilterer creates a new log filterer instance of Material, bound to a specific deployed contract.
func NewMaterialFilterer(address common.Address, filterer bind.ContractFilterer) (*MaterialFilterer, error) {
	contract, err := bindMaterial(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaterialFilterer{contract: contract}, nil
}

// bindMaterial binds a generic wrapper to an already deployed contract.
func bindMaterial(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MaterialABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Material *MaterialRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Material.Contract.MaterialCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Material *MaterialRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Material.Contract.MaterialTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Material *MaterialRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Material.Contract.MaterialTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Material *MaterialCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Material.Contract.contract.Call(opts, result, method, params...)
}

func (_Material *MaterialCallerRaw) ReadCall(result interface{}, method string, output []byte) error {
	return _Material.Contract.contract.ReadCall(result, method, output)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Material *MaterialTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Material.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Material *MaterialTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Material.Contract.contract.Transact(opts, method, params...)
}

// BatchInfos is a free data retrieval call binding the contract method 0xc4e30a2d.
//
// Solidity: function batchInfos(uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCaller) BatchInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	ret := new(struct {
		Producer     common.Address
		CreatedAt    *big.Int
		BatchID      *big.Int
		MaterialType *big.Int
		TotalNum     *big.Int
		KeptNum      *big.Int
	})
	out := ret
	err := _Material.contract.Call(opts, out, "batchInfos", arg0)
	return *ret, err
}

// BatchInfos is a free data retrieval call binding the contract method 0xc4e30a2d.
//
// Solidity: function batchInfos(uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialSession) BatchInfos(arg0 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.BatchInfos(&_Material.CallOpts, arg0)
}

// BatchInfos is a free data retrieval call binding the contract method 0xc4e30a2d.
//
// Solidity: function batchInfos(uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCallerSession) BatchInfos(arg0 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.BatchInfos(&_Material.CallOpts, arg0)
}

// GetMaterialProducer is a free data retrieval call binding the contract method 0x8c971e36.
//
// Solidity: function getMaterialProducer(uint256 materialType) view returns(address)
func (_Material *MaterialCaller) GetMaterialProducer(opts *bind.CallOpts, materialType *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Material.contract.Call(opts, out, "getMaterialProducer", materialType)
	return *ret0, err
}

// GetMaterialProducer is a free data retrieval call binding the contract method 0x8c971e36.
//
// Solidity: function getMaterialProducer(uint256 materialType) view returns(address)
func (_Material *MaterialSession) GetMaterialProducer(materialType *big.Int) (common.Address, error) {
	return _Material.Contract.GetMaterialProducer(&_Material.CallOpts, materialType)
}

// GetMaterialProducer is a free data retrieval call binding the contract method 0x8c971e36.
//
// Solidity: function getMaterialProducer(uint256 materialType) view returns(address)
func (_Material *MaterialCallerSession) GetMaterialProducer(materialType *big.Int) (common.Address, error) {
	return _Material.Contract.GetMaterialProducer(&_Material.CallOpts, materialType)
}

// GetMyMaterial is a free data retrieval call binding the contract method 0x73cbde55.
//
// Solidity: function getMyMaterial(uint256 materialType) view returns(uint256 num)
func (_Material *MaterialCaller) GetMyMaterial(opts *bind.CallOpts, materialType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Material.contract.Call(opts, out, "getMyMaterial", materialType)
	return *ret0, err
}

// GetMyMaterial is a free data retrieval call binding the contract method 0x73cbde55.
//
// Solidity: function getMyMaterial(uint256 materialType) view returns(uint256 num)
func (_Material *MaterialSession) GetMyMaterial(materialType *big.Int) (*big.Int, error) {
	return _Material.Contract.GetMyMaterial(&_Material.CallOpts, materialType)
}

// GetMyMaterial is a free data retrieval call binding the contract method 0x73cbde55.
//
// Solidity: function getMyMaterial(uint256 materialType) view returns(uint256 num)
func (_Material *MaterialCallerSession) GetMyMaterial(materialType *big.Int) (*big.Int, error) {
	return _Material.Contract.GetMyMaterial(&_Material.CallOpts, materialType)
}

// GetPrice is a free data retrieval call binding the contract method 0x449e815d.
//
// Solidity: function getPrice(address to, uint256 materialType) view returns(uint256)
func (_Material *MaterialCaller) GetPrice(opts *bind.CallOpts, to common.Address, materialType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Material.contract.Call(opts, out, "getPrice", to, materialType)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x449e815d.
//
// Solidity: function getPrice(address to, uint256 materialType) view returns(uint256)
func (_Material *MaterialSession) GetPrice(to common.Address, materialType *big.Int) (*big.Int, error) {
	return _Material.Contract.GetPrice(&_Material.CallOpts, to, materialType)
}

// GetPrice is a free data retrieval call binding the contract method 0x449e815d.
//
// Solidity: function getPrice(address to, uint256 materialType) view returns(uint256)
func (_Material *MaterialCallerSession) GetPrice(to common.Address, materialType *big.Int) (*big.Int, error) {
	return _Material.Contract.GetPrice(&_Material.CallOpts, to, materialType)
}

// KeptMaterials is a free data retrieval call binding the contract method 0xe0159f8e.
//
// Solidity: function keptMaterials(address , uint256 , uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCaller) KeptMaterials(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	ret := new(struct {
		Producer     common.Address
		CreatedAt    *big.Int
		BatchID      *big.Int
		MaterialType *big.Int
		TotalNum     *big.Int
		KeptNum      *big.Int
	})
	out := ret
	err := _Material.contract.Call(opts, out, "keptMaterials", arg0, arg1, arg2)
	return *ret, err
}

// KeptMaterials is a free data retrieval call binding the contract method 0xe0159f8e.
//
// Solidity: function keptMaterials(address , uint256 , uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialSession) KeptMaterials(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.KeptMaterials(&_Material.CallOpts, arg0, arg1, arg2)
}

// KeptMaterials is a free data retrieval call binding the contract method 0xe0159f8e.
//
// Solidity: function keptMaterials(address , uint256 , uint256 ) view returns(address producer, uint256 createdAt, uint256 batchID, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCallerSession) KeptMaterials(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchID      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.KeptMaterials(&_Material.CallOpts, arg0, arg1, arg2)
}

// ShowBatchInfo is a free data retrieval call binding the contract method 0x14c7b2f8.
//
// Solidity: function showBatchInfo(uint256 batchID) view returns(address producer, uint256 createdAt, uint256 batchId, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCaller) ShowBatchInfo(opts *bind.CallOpts, batchID *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchId      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	ret := new(struct {
		Producer     common.Address
		CreatedAt    *big.Int
		BatchId      *big.Int
		MaterialType *big.Int
		TotalNum     *big.Int
		KeptNum      *big.Int
	})
	out := ret
	err := _Material.contract.Call(opts, out, "showBatchInfo", batchID)
	return *ret, err
}

// ShowBatchInfo is a free data retrieval call binding the contract method 0x14c7b2f8.
//
// Solidity: function showBatchInfo(uint256 batchID) view returns(address producer, uint256 createdAt, uint256 batchId, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialSession) ShowBatchInfo(batchID *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchId      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.ShowBatchInfo(&_Material.CallOpts, batchID)
}

// ShowBatchInfo is a free data retrieval call binding the contract method 0x14c7b2f8.
//
// Solidity: function showBatchInfo(uint256 batchID) view returns(address producer, uint256 createdAt, uint256 batchId, uint256 materialType, uint256 totalNum, uint256 keptNum)
func (_Material *MaterialCallerSession) ShowBatchInfo(batchID *big.Int) (struct {
	Producer     common.Address
	CreatedAt    *big.Int
	BatchId      *big.Int
	MaterialType *big.Int
	TotalNum     *big.Int
	KeptNum      *big.Int
}, error) {
	return _Material.Contract.ShowBatchInfo(&_Material.CallOpts, batchID)
}

// UsedBatchIdx is a free data retrieval call binding the contract method 0x612e324a.
//
// Solidity: function usedBatchIdx(address , uint256 ) view returns(uint256)
func (_Material *MaterialCaller) UsedBatchIdx(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Material.contract.Call(opts, out, "usedBatchIdx", arg0, arg1)
	return *ret0, err
}

// UsedBatchIdx is a free data retrieval call binding the contract method 0x612e324a.
//
// Solidity: function usedBatchIdx(address , uint256 ) view returns(uint256)
func (_Material *MaterialSession) UsedBatchIdx(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Material.Contract.UsedBatchIdx(&_Material.CallOpts, arg0, arg1)
}

// UsedBatchIdx is a free data retrieval call binding the contract method 0x612e324a.
//
// Solidity: function usedBatchIdx(address , uint256 ) view returns(uint256)
func (_Material *MaterialCallerSession) UsedBatchIdx(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Material.Contract.UsedBatchIdx(&_Material.CallOpts, arg0, arg1)
}

// ConsumeMaterial is a paid mutator transaction binding the contract method 0xdc821562.
//
// Solidity: function consumeMaterial(uint256 materialType, uint256 num) returns()

/******************************************************************************************************************************/
func (_Material *MaterialCaller) ReadConsumeMaterial(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Material.contract.ReadCall(out, "consumeMaterial", hexutil.MustDecode(output))
	return err
}

// ConsumeMaterial is a free data retrieval call binding the contract method 0xdc821562.
//
// Solidity: function consumeMaterial(uint256 materialType, uint256 num) returns()
func (_Material *MaterialSession) ReadConsumeMaterial(output string) error {
	return _Material.Contract.ReadConsumeMaterial(output)
}

// ConsumeMaterial is a free data retrieval call binding the contract method 0xdc821562.
//
// Solidity: function consumeMaterial(uint256 materialType, uint256 num) returns()
func (_Material *MaterialCallerSession) ReadConsumeMaterial(output string) error {
	return _Material.Contract.ReadConsumeMaterial(output)
}

/******************************************************************************************************************************/

func (_Material *MaterialTransactor) ConsumeMaterial(opts *bind.TransactOpts, materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.contract.Transact(opts, "consumeMaterial", materialType, num)
}

// ConsumeMaterial is a paid mutator transaction binding the contract method 0xdc821562.
//
// Solidity: function consumeMaterial(uint256 materialType, uint256 num) returns()
func (_Material *MaterialSession) ConsumeMaterial(materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.Contract.ConsumeMaterial(&_Material.TransactOpts, materialType, num)
}

// ConsumeMaterial is a paid mutator transaction binding the contract method 0xdc821562.
//
// Solidity: function consumeMaterial(uint256 materialType, uint256 num) returns()
func (_Material *MaterialTransactorSession) ConsumeMaterial(materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.Contract.ConsumeMaterial(&_Material.TransactOpts, materialType, num)
}

// NewMaterial is a paid mutator transaction binding the contract method 0xb00351d5.
//
// Solidity: function newMaterial(uint256 materialType, uint256 totalNum, uint256 batchID) returns(materialRawMaterial m)

/******************************************************************************************************************************/
func (_Material *MaterialCaller) ReadNewMaterial(output string) (materialRawMaterial, error) {
	var (
		ret0 = new(materialRawMaterial)
	)
	out := ret0
	err := _Material.contract.ReadCall(out, "newMaterial", hexutil.MustDecode(output))
	return *ret0, err
}

// NewMaterial is a free data retrieval call binding the contract method 0xb00351d5.
//
// Solidity: function newMaterial(uint256 materialType, uint256 totalNum, uint256 batchID) returns(materialRawMaterial m)
func (_Material *MaterialSession) ReadNewMaterial(output string) (materialRawMaterial, error) {
	return _Material.Contract.ReadNewMaterial(output)
}

// NewMaterial is a free data retrieval call binding the contract method 0xb00351d5.
//
// Solidity: function newMaterial(uint256 materialType, uint256 totalNum, uint256 batchID) returns(materialRawMaterial m)
func (_Material *MaterialCallerSession) ReadNewMaterial(output string) (materialRawMaterial, error) {
	return _Material.Contract.ReadNewMaterial(output)
}

/******************************************************************************************************************************/

func (_Material *MaterialTransactor) NewMaterial(opts *bind.TransactOpts, materialType *big.Int, totalNum *big.Int, batchID *big.Int) (*types.Transaction, error) {
	return _Material.contract.Transact(opts, "newMaterial", materialType, totalNum, batchID)
}

// NewMaterial is a paid mutator transaction binding the contract method 0xb00351d5.
//
// Solidity: function newMaterial(uint256 materialType, uint256 totalNum, uint256 batchID) returns(materialRawMaterial m)
func (_Material *MaterialSession) NewMaterial(materialType *big.Int, totalNum *big.Int, batchID *big.Int) (*types.Transaction, error) {
	return _Material.Contract.NewMaterial(&_Material.TransactOpts, materialType, totalNum, batchID)
}

// NewMaterial is a paid mutator transaction binding the contract method 0xb00351d5.
//
// Solidity: function newMaterial(uint256 materialType, uint256 totalNum, uint256 batchID) returns(materialRawMaterial m)
func (_Material *MaterialTransactorSession) NewMaterial(materialType *big.Int, totalNum *big.Int, batchID *big.Int) (*types.Transaction, error) {
	return _Material.Contract.NewMaterial(&_Material.TransactOpts, materialType, totalNum, batchID)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 materialType, uint256 price) returns()

/******************************************************************************************************************************/
func (_Material *MaterialCaller) ReadSetPrice(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Material.contract.ReadCall(out, "setPrice", hexutil.MustDecode(output))
	return err
}

// SetPrice is a free data retrieval call binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 materialType, uint256 price) returns()
func (_Material *MaterialSession) ReadSetPrice(output string) error {
	return _Material.Contract.ReadSetPrice(output)
}

// SetPrice is a free data retrieval call binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 materialType, uint256 price) returns()
func (_Material *MaterialCallerSession) ReadSetPrice(output string) error {
	return _Material.Contract.ReadSetPrice(output)
}

/******************************************************************************************************************************/

func (_Material *MaterialTransactor) SetPrice(opts *bind.TransactOpts, materialType *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Material.contract.Transact(opts, "setPrice", materialType, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 materialType, uint256 price) returns()
func (_Material *MaterialSession) SetPrice(materialType *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Material.Contract.SetPrice(&_Material.TransactOpts, materialType, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 materialType, uint256 price) returns()
func (_Material *MaterialTransactorSession) SetPrice(materialType *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Material.Contract.SetPrice(&_Material.TransactOpts, materialType, price)
}

// TransferMaterial is a paid mutator transaction binding the contract method 0xf45f2ee0.
//
// Solidity: function transferMaterial(address _from, address _to, uint256 materialType, uint256 num) returns()

/******************************************************************************************************************************/
func (_Material *MaterialCaller) ReadTransferMaterial(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Material.contract.ReadCall(out, "transferMaterial", hexutil.MustDecode(output))
	return err
}

// TransferMaterial is a free data retrieval call binding the contract method 0xf45f2ee0.
//
// Solidity: function transferMaterial(address _from, address _to, uint256 materialType, uint256 num) returns()
func (_Material *MaterialSession) ReadTransferMaterial(output string) error {
	return _Material.Contract.ReadTransferMaterial(output)
}

// TransferMaterial is a free data retrieval call binding the contract method 0xf45f2ee0.
//
// Solidity: function transferMaterial(address _from, address _to, uint256 materialType, uint256 num) returns()
func (_Material *MaterialCallerSession) ReadTransferMaterial(output string) error {
	return _Material.Contract.ReadTransferMaterial(output)
}

/******************************************************************************************************************************/

func (_Material *MaterialTransactor) TransferMaterial(opts *bind.TransactOpts, _from common.Address, _to common.Address, materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.contract.Transact(opts, "transferMaterial", _from, _to, materialType, num)
}

// TransferMaterial is a paid mutator transaction binding the contract method 0xf45f2ee0.
//
// Solidity: function transferMaterial(address _from, address _to, uint256 materialType, uint256 num) returns()
func (_Material *MaterialSession) TransferMaterial(_from common.Address, _to common.Address, materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.Contract.TransferMaterial(&_Material.TransactOpts, _from, _to, materialType, num)
}

// TransferMaterial is a paid mutator transaction binding the contract method 0xf45f2ee0.
//
// Solidity: function transferMaterial(address _from, address _to, uint256 materialType, uint256 num) returns()
func (_Material *MaterialTransactorSession) TransferMaterial(_from common.Address, _to common.Address, materialType *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Material.Contract.TransferMaterial(&_Material.TransactOpts, _from, _to, materialType, num)
}

// UpdateProducer is a paid mutator transaction binding the contract method 0x48640e33.
//
// Solidity: function updateProducer(uint256 materialType, address producer) returns()

/******************************************************************************************************************************/
func (_Material *MaterialCaller) ReadUpdateProducer(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Material.contract.ReadCall(out, "updateProducer", hexutil.MustDecode(output))
	return err
}

// UpdateProducer is a free data retrieval call binding the contract method 0x48640e33.
//
// Solidity: function updateProducer(uint256 materialType, address producer) returns()
func (_Material *MaterialSession) ReadUpdateProducer(output string) error {
	return _Material.Contract.ReadUpdateProducer(output)
}

// UpdateProducer is a free data retrieval call binding the contract method 0x48640e33.
//
// Solidity: function updateProducer(uint256 materialType, address producer) returns()
func (_Material *MaterialCallerSession) ReadUpdateProducer(output string) error {
	return _Material.Contract.ReadUpdateProducer(output)
}

/******************************************************************************************************************************/

func (_Material *MaterialTransactor) UpdateProducer(opts *bind.TransactOpts, materialType *big.Int, producer common.Address) (*types.Transaction, error) {
	return _Material.contract.Transact(opts, "updateProducer", materialType, producer)
}

// UpdateProducer is a paid mutator transaction binding the contract method 0x48640e33.
//
// Solidity: function updateProducer(uint256 materialType, address producer) returns()
func (_Material *MaterialSession) UpdateProducer(materialType *big.Int, producer common.Address) (*types.Transaction, error) {
	return _Material.Contract.UpdateProducer(&_Material.TransactOpts, materialType, producer)
}

// UpdateProducer is a paid mutator transaction binding the contract method 0x48640e33.
//
// Solidity: function updateProducer(uint256 materialType, address producer) returns()
func (_Material *MaterialTransactorSession) UpdateProducer(materialType *big.Int, producer common.Address) (*types.Transaction, error) {
	return _Material.Contract.UpdateProducer(&_Material.TransactOpts, materialType, producer)
}

// MaterialEvtMaterialConsumedIterator is returned from FilterEvtMaterialConsumed and is used to iterate over the raw logs and unpacked data for EvtMaterialConsumed events raised by the Material contract.
type MaterialEvtMaterialConsumedIterator struct {
	Event *MaterialEvtMaterialConsumed // Event containing the contract specifics and raw log

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
func (it *MaterialEvtMaterialConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaterialEvtMaterialConsumed)
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
		it.Event = new(MaterialEvtMaterialConsumed)
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
func (it *MaterialEvtMaterialConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaterialEvtMaterialConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaterialEvtMaterialConsumed represents a EvtMaterialConsumed event raised by the Material contract.
type MaterialEvtMaterialConsumed struct {
	From         common.Address
	MaterialType *big.Int
	Num          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvtMaterialConsumed is a free log retrieval operation binding the contract event 0xbb29a9e0a07ed37fc0d005182fb0c11a1fae309a397dfb43b0ec7af31caa90dc.
//
// Solidity: event EvtMaterialConsumed(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) FilterEvtMaterialConsumed(opts *bind.FilterOpts) (*MaterialEvtMaterialConsumedIterator, error) {

	logs, sub, err := _Material.contract.FilterLogs(opts, "EvtMaterialConsumed")
	if err != nil {
		return nil, err
	}
	return &MaterialEvtMaterialConsumedIterator{contract: _Material.contract, event: "EvtMaterialConsumed", logs: logs, sub: sub}, nil
}

// WatchEvtMaterialConsumed is a free log subscription operation binding the contract event 0xbb29a9e0a07ed37fc0d005182fb0c11a1fae309a397dfb43b0ec7af31caa90dc.
//
// Solidity: event EvtMaterialConsumed(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) WatchEvtMaterialConsumed(opts *bind.WatchOpts, sink chan<- *MaterialEvtMaterialConsumed) (event.Subscription, error) {

	logs, sub, err := _Material.contract.WatchLogs(opts, "EvtMaterialConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaterialEvtMaterialConsumed)
				if err := _Material.contract.UnpackLog(event, "EvtMaterialConsumed", log); err != nil {
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

// ParseEvtMaterialConsumed is a log parse operation binding the contract event 0xbb29a9e0a07ed37fc0d005182fb0c11a1fae309a397dfb43b0ec7af31caa90dc.
//
// Solidity: event EvtMaterialConsumed(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) ParseEvtMaterialConsumed(log types.Log) (*MaterialEvtMaterialConsumed, error) {
	event := new(MaterialEvtMaterialConsumed)
	if err := _Material.contract.UnpackLog(event, "EvtMaterialConsumed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MaterialEvtMaterialCreatedIterator is returned from FilterEvtMaterialCreated and is used to iterate over the raw logs and unpacked data for EvtMaterialCreated events raised by the Material contract.
type MaterialEvtMaterialCreatedIterator struct {
	Event *MaterialEvtMaterialCreated // Event containing the contract specifics and raw log

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
func (it *MaterialEvtMaterialCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaterialEvtMaterialCreated)
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
		it.Event = new(MaterialEvtMaterialCreated)
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
func (it *MaterialEvtMaterialCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaterialEvtMaterialCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaterialEvtMaterialCreated represents a EvtMaterialCreated event raised by the Material contract.
type MaterialEvtMaterialCreated struct {
	From         common.Address
	MaterialType *big.Int
	Num          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvtMaterialCreated is a free log retrieval operation binding the contract event 0xe1706a40c6f84da3bd18ed43b695c0d62afa92863e79c212de31eda5b77583d8.
//
// Solidity: event EvtMaterialCreated(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) FilterEvtMaterialCreated(opts *bind.FilterOpts) (*MaterialEvtMaterialCreatedIterator, error) {

	logs, sub, err := _Material.contract.FilterLogs(opts, "EvtMaterialCreated")
	if err != nil {
		return nil, err
	}
	return &MaterialEvtMaterialCreatedIterator{contract: _Material.contract, event: "EvtMaterialCreated", logs: logs, sub: sub}, nil
}

// WatchEvtMaterialCreated is a free log subscription operation binding the contract event 0xe1706a40c6f84da3bd18ed43b695c0d62afa92863e79c212de31eda5b77583d8.
//
// Solidity: event EvtMaterialCreated(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) WatchEvtMaterialCreated(opts *bind.WatchOpts, sink chan<- *MaterialEvtMaterialCreated) (event.Subscription, error) {

	logs, sub, err := _Material.contract.WatchLogs(opts, "EvtMaterialCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaterialEvtMaterialCreated)
				if err := _Material.contract.UnpackLog(event, "EvtMaterialCreated", log); err != nil {
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

// ParseEvtMaterialCreated is a log parse operation binding the contract event 0xe1706a40c6f84da3bd18ed43b695c0d62afa92863e79c212de31eda5b77583d8.
//
// Solidity: event EvtMaterialCreated(address from, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) ParseEvtMaterialCreated(log types.Log) (*MaterialEvtMaterialCreated, error) {
	event := new(MaterialEvtMaterialCreated)
	if err := _Material.contract.UnpackLog(event, "EvtMaterialCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MaterialEvtMaterialTransferredIterator is returned from FilterEvtMaterialTransferred and is used to iterate over the raw logs and unpacked data for EvtMaterialTransferred events raised by the Material contract.
type MaterialEvtMaterialTransferredIterator struct {
	Event *MaterialEvtMaterialTransferred // Event containing the contract specifics and raw log

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
func (it *MaterialEvtMaterialTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaterialEvtMaterialTransferred)
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
		it.Event = new(MaterialEvtMaterialTransferred)
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
func (it *MaterialEvtMaterialTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaterialEvtMaterialTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaterialEvtMaterialTransferred represents a EvtMaterialTransferred event raised by the Material contract.
type MaterialEvtMaterialTransferred struct {
	From         common.Address
	To           common.Address
	MaterialType *big.Int
	Num          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvtMaterialTransferred is a free log retrieval operation binding the contract event 0x70d9b7ac16a9bb81749f2f28832acbf5d2a70bc09695163d1de16338b34bca10.
//
// Solidity: event EvtMaterialTransferred(address from, address to, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) FilterEvtMaterialTransferred(opts *bind.FilterOpts) (*MaterialEvtMaterialTransferredIterator, error) {

	logs, sub, err := _Material.contract.FilterLogs(opts, "EvtMaterialTransferred")
	if err != nil {
		return nil, err
	}
	return &MaterialEvtMaterialTransferredIterator{contract: _Material.contract, event: "EvtMaterialTransferred", logs: logs, sub: sub}, nil
}

// WatchEvtMaterialTransferred is a free log subscription operation binding the contract event 0x70d9b7ac16a9bb81749f2f28832acbf5d2a70bc09695163d1de16338b34bca10.
//
// Solidity: event EvtMaterialTransferred(address from, address to, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) WatchEvtMaterialTransferred(opts *bind.WatchOpts, sink chan<- *MaterialEvtMaterialTransferred) (event.Subscription, error) {

	logs, sub, err := _Material.contract.WatchLogs(opts, "EvtMaterialTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaterialEvtMaterialTransferred)
				if err := _Material.contract.UnpackLog(event, "EvtMaterialTransferred", log); err != nil {
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

// ParseEvtMaterialTransferred is a log parse operation binding the contract event 0x70d9b7ac16a9bb81749f2f28832acbf5d2a70bc09695163d1de16338b34bca10.
//
// Solidity: event EvtMaterialTransferred(address from, address to, uint256 materialType, uint256 num)
func (_Material *MaterialFilterer) ParseEvtMaterialTransferred(log types.Log) (*MaterialEvtMaterialTransferred, error) {
	event := new(MaterialEvtMaterialTransferred)
	if err := _Material.contract.UnpackLog(event, "EvtMaterialTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MaterialEvtPriceUpdatedIterator is returned from FilterEvtPriceUpdated and is used to iterate over the raw logs and unpacked data for EvtPriceUpdated events raised by the Material contract.
type MaterialEvtPriceUpdatedIterator struct {
	Event *MaterialEvtPriceUpdated // Event containing the contract specifics and raw log

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
func (it *MaterialEvtPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaterialEvtPriceUpdated)
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
		it.Event = new(MaterialEvtPriceUpdated)
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
func (it *MaterialEvtPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaterialEvtPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaterialEvtPriceUpdated represents a EvtPriceUpdated event raised by the Material contract.
type MaterialEvtPriceUpdated struct {
	From         common.Address
	MaterialType *big.Int
	Price        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvtPriceUpdated is a free log retrieval operation binding the contract event 0x3061c2e7e7cfec0710fd23c1e44a860a3319b77dc3bbf3283a2a6448f372e60d.
//
// Solidity: event EvtPriceUpdated(address from, uint256 materialType, uint256 price)
func (_Material *MaterialFilterer) FilterEvtPriceUpdated(opts *bind.FilterOpts) (*MaterialEvtPriceUpdatedIterator, error) {

	logs, sub, err := _Material.contract.FilterLogs(opts, "EvtPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &MaterialEvtPriceUpdatedIterator{contract: _Material.contract, event: "EvtPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchEvtPriceUpdated is a free log subscription operation binding the contract event 0x3061c2e7e7cfec0710fd23c1e44a860a3319b77dc3bbf3283a2a6448f372e60d.
//
// Solidity: event EvtPriceUpdated(address from, uint256 materialType, uint256 price)
func (_Material *MaterialFilterer) WatchEvtPriceUpdated(opts *bind.WatchOpts, sink chan<- *MaterialEvtPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _Material.contract.WatchLogs(opts, "EvtPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaterialEvtPriceUpdated)
				if err := _Material.contract.UnpackLog(event, "EvtPriceUpdated", log); err != nil {
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

// ParseEvtPriceUpdated is a log parse operation binding the contract event 0x3061c2e7e7cfec0710fd23c1e44a860a3319b77dc3bbf3283a2a6448f372e60d.
//
// Solidity: event EvtPriceUpdated(address from, uint256 materialType, uint256 price)
func (_Material *MaterialFilterer) ParseEvtPriceUpdated(log types.Log) (*MaterialEvtPriceUpdated, error) {
	event := new(MaterialEvtPriceUpdated)
	if err := _Material.contract.UnpackLog(event, "EvtPriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
