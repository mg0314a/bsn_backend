// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package produce

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

// ProduceProduct is an auto generated low-level Go binding around an user-defined struct.
type ProduceProduct struct {
	Owner           common.Address
	Producer        common.Address
	CreatedAt       *big.Int
	Batch           *big.Int
	MaterialBatches []*big.Int
	Sold            bool
}

// ProduceABI is the input ABI used to generate the binding from.
const ProduceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_materialTypeCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"productID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"EvtProductCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"EvtProductOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"details\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"materialBatches\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structProduce.Product\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"}],\"name\":\"getMyProducts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"myProductIDs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"}],\"name\":\"getProductPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"materialBatches\",\"type\":\"uint256[]\"}],\"name\":\"registerProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_materialContract\",\"type\":\"address\"}],\"name\":\"setMaterialContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentContract\",\"type\":\"address\"}],\"name\":\"setPaymentContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"materialBatchNum\",\"type\":\"uint256\"}],\"name\":\"trace\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"transferProducts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateProductPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProduceBin is the compiled bytecode used for deploying new contracts.
var ProduceBin = "0x60806040523480156200001157600080fd5b506040516200202f3803806200202f83398181016040526200003791908101906200013f565b6000620000496200013760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060088190555050506200017a565b600033905090565b6000806040838503121562000152578182fd5b825160018060a01b038116811462000168578283fd5b80925050602083015190509250929050565b611ea5806200018a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638da5cb5b1161008c578063a005ec7a11610066578063a005ec7a14610200578063a68a4fed14610230578063b26f324614610260578063f2fde38b1461027c576100cf565b80638da5cb5b146101aa5780639c3bddb2146101c85780639f96eb9d146101e4576100cf565b8063346b77f1146100d457806336a714f614610104578063715018a6146101205780637acc0b201461012a57806381ee6c8c1461015e5780638abb20cc1461018e575b600080fd5b6100ee60048036036100e99190810190611933565b610298565b6040516100fb9190611dee565b60405180910390f35b61011e600480360361011991908101906118c9565b6102f3565b005b6101286103cd565b005b610144600480360361013f9190810190611989565b610522565b604051610155959493929190611b1c565b60405180910390f35b61017860048036036101739190810190611989565b6105a5565b6040516101859190611b57565b60405180910390f35b6101a860048036036101a391908101906118c9565b610783565b005b6101b261085d565b6040516101bf9190611ac6565b60405180910390f35b6101e260048036036101dd91908101906119a4565b610886565b005b6101fe60048036036101f991908101906119c8565b6109c9565b005b61021a60048036036102159190810190611989565b610da0565b6040516102279190611d87565b60405180910390f35b61024a60048036036102459190810190611989565b610fb2565b6040516102579190611b57565b60405180910390f35b61027a600480360361027591908101906118ec565b6110a8565b005b610296600480360361029191908101906118c9565b611317565b005b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b6102fb6114db565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610389576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038090611c52565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103d56114db565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a90611c52565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060050160009054906101000a900460ff16905085565b60606105ff600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206114e3565b60405190808252806020026020018201604052801561062d5781602001602082028038833980820191505090505b50905060008090506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000206001015490505b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000206002015481101561077957600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060000160008281526020019081526020016000205483838060010194508151811061076057fe5b602002602001018181525050808060010191505061068e565b5081915050919050565b61078b6114db565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610819576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081090611c52565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a09409a5336040518263ffffffff1660e01b81526004016108e19190611adf565b60206040518083038186803b1580156108f957600080fd5b505afa15801561090d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610931919081019061195f565b610970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096790611d49565b60405180910390fd5b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020819055505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a09409a5336040518263ffffffff1660e01b8152600401610a249190611adf565b60206040518083038186803b158015610a3c57600080fd5b505afa158015610a50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a74919081019061195f565b610ab3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aaa90611d49565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166005600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4f90611bb0565b60405180910390fd5b6040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001428152602001838152602001828152602001600015158152506005600085815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004019080519060200190610c859291906117d3565b5060a08201518160050160006101000a81548160ff021916908315150217905550905050610d0b83600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002061150990919063ffffffff16565b60008090505b8151811015610d5f57610d528460096000858581518110610d2e57fe5b6020026020010151815260200190815260200160002061150990919063ffffffff16565b8080600101915050610d11565b50827f7f73f7750119b7ad6911a0a4f5394ab05c7261658625bec67841058aaa682f088533604051610d92929190611dff565b60405180910390a250505050565b610da8611820565b600073ffffffffffffffffffffffffffffffffffffffff166005600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610e4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4590611d0b565b60405180910390fd5b610e56611820565b600560008481526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201805480602002602001604051908101604052809291908181526020018280548015610f8357602002820191906000526020600020905b815481526020019060010190808311610f6f575b505050505081526020016005820160009054906101000a900460ff161515151581525050905080915050919050565b6060610fcf600960008481526020019081526020016000206114e3565b604051908082528060200260200182016040528015610ffd5781602001602082028038833980820191505090505b5090506000600960008481526020019081526020016000206001015490505b600960008481526020019081526020016000206002015481101561109f576009600084815260200190815260200160002060000160008281526020019081526020016000205482600960008681526020019081526020016000206001015483038151811061108657fe5b602002602001018181525050808060010191505061101c565b50809050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad4a296d336040518263ffffffff1660e01b81526004016111039190611adf565b60206040518083038186803b15801561111b57600080fd5b505afa15801561112f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611153919081019061195f565b611192576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118990611ccd565b60405180910390fd5b6111ea600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206114e3565b81111561122c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122390611c8f565b60405180910390fd5b60008090505b81811015611310576000611294600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020611551565b90506112f881600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002061150990919063ffffffff16565b61130281866115b9565b508080600101915050611232565b5050505050565b61131f6114db565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146113ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a490611c52565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561141d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141490611bee565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b60008160020154826001015411156114f757fe5b81600101548260020154039050919050565b81600201548260010154111561151b57fe5b80826000016000846002015481526020019081526020016000208190555081600201600081548092919060010191905055505050565b6000816002015482600101541061156457fe5b8160000160008360010154815260200190815260200160002054905081600001600083600101548152602001908152602001600020600090558160010160008154809291906001019190505550809050919050565b600073ffffffffffffffffffffffffffffffffffffffff166005600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561165f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161165690611d0b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166005600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611704576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116fb90611b72565b60405180910390fd5b60006005600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816005600085815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550827f0a04c007c049cdb8eb2c578fc38a48849cf66acce185c47f8fcf11a2957b620c82846040516117c6929190611af8565b60405180910390a2505050565b82805482825590600052602060002090810192821561180f579160200282015b8281111561180e5782518255916020019190600101906117f3565b5b50905061181c9190611884565b5090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081526020016000151581525090565b6118a691905b808211156118a257600081600090555060010161188a565b5090565b90565b60008135905060018060a01b03811681146118c357600080fd5b92915050565b6000602082840312156118da578081fd5b6118e483836118a9565b905092915050565b60008060008060808587031215611901578283fd5b61190b86866118a9565b935061191a86602087016118a9565b9250604085013591506060850135905092959194509250565b60008060408385031215611945578182fd5b61194f84846118a9565b9150602083013590509250929050565b600060208284031215611970578081fd5b8151801515811461197f578182fd5b8091505092915050565b60006020828403121561199a578081fd5b8135905092915050565b600080604083850312156119b6578182fd5b82359150602083013590509250929050565b600080600080608085870312156119dd578384fd5b84359350602080860135935060408601359250606086013567ffffffffffffffff811115611a09578283fd5b80870188601f820112611a1a578384fd5b80359150611a2f611a2a83611e4c565b611e1f565b8083825284820191508483018b868787028601011115611a4d578687fd5b8693505b84841015611a72578035835285830192508581019050600184019350611a51565b50809550505050505092959194509250565b6000815180845260208401935060208301825b82811015611aba5781518652602086019550602082019150600181019050611a97565b50505082905092915050565b600060208201905060018060a01b038316825292915050565b600060208201905060018060a01b038316825292915050565b600060408201905060018060a01b0380851683528084166020840152509392505050565b600060a08201905060018060a01b03808816835280871660208401525084604083015283606083015282151560808301529695505050505050565b600060208252611b6a6020830184611a84565b905092915050565b600060208252601b60208301527f73656c662d7472616e7366657220697320646973616c6c6f77656400000000006040830152606082019050919050565b600060208252601660208301527f70726f6475637420616c726561647920657869737473000000000000000000006040830152606082019050919050565b600060208252602660208301527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408301527f64647265737300000000000000000000000000000000000000000000000000006060830152608082019050919050565b6000602082526020808301527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726040830152606082019050919050565b600060208252601460208301527f696e73756666696369656e742070726f647563740000000000000000000000006040830152606082019050919050565b600060208252601060208301527f6f6e6c7920666f72207061796d656e74000000000000000000000000000000006040830152606082019050919050565b600060208252601660208301527f70726f6475637420646f6573206e6f74206578697374000000000000000000006040830152606082019050919050565b600060208252601960208301527f6f6e6c7920666f722070726f647563742070726f6475636572000000000000006040830152606082019050919050565b60006020825260018060a01b03808451166020840152806020850151166040840152506040830151606083015260608301516080830152608083015160c060a0840152611dd760e0840182611a84565b60a0850151151560c0850152809250505092915050565b600060208201905082825292915050565b600060408201905083825260018060a01b03831660208301529392505050565b6000604051905081810181811067ffffffffffffffff82111715611e4257600080fd5b8060405250919050565b600067ffffffffffffffff821115611e62578081fd5b602080830201905091905056fea2646970667358221220b1156132f9a5aae02b70e7dfaa7244403c0a7f10f1478398600e68b3e0e7ff6264736f6c63430006000033"

// DeployProduce deploys a new FiscoBcos contract, binding an instance of Produce to it.
func DeployProduce(auth *bind.TransactOpts, backend bind.ContractBackend, accessContractAddress common.Address, _materialTypeCount *big.Int) (common.Address, *types.Transaction, *Produce, error) {
	parsed, err := abi.JSON(strings.NewReader(ProduceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProduceBin), backend, accessContractAddress, _materialTypeCount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Produce{ProduceCaller: ProduceCaller{contract: contract}, ProduceTransactor: ProduceTransactor{contract: contract}, ProduceFilterer: ProduceFilterer{contract: contract}, ABI: parsed}, nil
}

// Produce is an auto generated Go binding around an FiscoBcos contract.
type Produce struct {
	ProduceCaller             // Read-only binding to the contract
	ProduceTransactor         // Write-only binding to the contract
	ProduceFilterer           // Log filterer for contract events
	ABI               abi.ABI // contract abi
}

// ProduceCaller is an auto generated read-only Go binding around an FiscoBcos contract.
type ProduceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProduceTransactor is an auto generated write-only Go binding around an FiscoBcos contract.
type ProduceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProduceFilterer is an auto generated log filtering Go binding around an FiscoBcos contract events.
type ProduceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProduceSession is an auto generated Go binding around an FiscoBcos contract,
// with pre-set call and transact options.
type ProduceSession struct {
	Contract     *Produce          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProduceCallerSession is an auto generated read-only Go binding around an FiscoBcos contract,
// with pre-set call options.
type ProduceCallerSession struct {
	Contract *ProduceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProduceTransactorSession is an auto generated write-only Go binding around an FiscoBcos contract,
// with pre-set transact options.
type ProduceTransactorSession struct {
	Contract     *ProduceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProduceRaw is an auto generated low-level Go binding around an FiscoBcos contract.
type ProduceRaw struct {
	Contract *Produce // Generic contract binding to access the raw methods on
}

// ProduceCallerRaw is an auto generated low-level read-only Go binding around an FiscoBcos contract.
type ProduceCallerRaw struct {
	Contract *ProduceCaller // Generic read-only contract binding to access the raw methods on
}

// ProduceTransactorRaw is an auto generated low-level write-only Go binding around an FiscoBcos contract.
type ProduceTransactorRaw struct {
	Contract *ProduceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProduce creates a new instance of Produce, bound to a specific deployed contract.
func NewProduce(address common.Address, backend bind.ContractBackend) (*Produce, error) {
	contract, err := bindProduce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(ProduceABI))
	if err != nil {
		return nil, err
	}
	return &Produce{ProduceCaller: ProduceCaller{contract: contract}, ProduceTransactor: ProduceTransactor{contract: contract}, ProduceFilterer: ProduceFilterer{contract: contract}, ABI: parsed}, nil
}

// NewProduceCaller creates a new read-only instance of Produce, bound to a specific deployed contract.
func NewProduceCaller(address common.Address, caller bind.ContractCaller) (*ProduceCaller, error) {
	contract, err := bindProduce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProduceCaller{contract: contract}, nil
}

// NewProduceTransactor creates a new write-only instance of Produce, bound to a specific deployed contract.
func NewProduceTransactor(address common.Address, transactor bind.ContractTransactor) (*ProduceTransactor, error) {
	contract, err := bindProduce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProduceTransactor{contract: contract}, nil
}

// NewProduceFilterer creates a new log filterer instance of Produce, bound to a specific deployed contract.
func NewProduceFilterer(address common.Address, filterer bind.ContractFilterer) (*ProduceFilterer, error) {
	contract, err := bindProduce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProduceFilterer{contract: contract}, nil
}

// bindProduce binds a generic wrapper to an already deployed contract.
func bindProduce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProduceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Produce *ProduceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Produce.Contract.ProduceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Produce *ProduceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Produce.Contract.ProduceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Produce *ProduceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Produce.Contract.ProduceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Produce *ProduceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Produce.Contract.contract.Call(opts, result, method, params...)
}

func (_Produce *ProduceCallerRaw) ReadCall(result interface{}, method string, output []byte) error {
	return _Produce.Contract.contract.ReadCall(result, method, output)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Produce *ProduceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Produce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Produce *ProduceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Produce.Contract.contract.Transact(opts, method, params...)
}

// Details is a free data retrieval call binding the contract method 0xa005ec7a.
//
// Solidity: function details(uint256 id) view returns(ProduceProduct)
func (_Produce *ProduceCaller) Details(opts *bind.CallOpts, id *big.Int) (ProduceProduct, error) {
	var (
		ret0 = new(ProduceProduct)
	)
	out := ret0
	err := _Produce.contract.Call(opts, out, "details", id)
	return *ret0, err
}

// Details is a free data retrieval call binding the contract method 0xa005ec7a.
//
// Solidity: function details(uint256 id) view returns(ProduceProduct)
func (_Produce *ProduceSession) Details(id *big.Int) (ProduceProduct, error) {
	return _Produce.Contract.Details(&_Produce.CallOpts, id)
}

// Details is a free data retrieval call binding the contract method 0xa005ec7a.
//
// Solidity: function details(uint256 id) view returns(ProduceProduct)
func (_Produce *ProduceCallerSession) Details(id *big.Int) (ProduceProduct, error) {
	return _Produce.Contract.Details(&_Produce.CallOpts, id)
}

// GetMyProducts is a free data retrieval call binding the contract method 0x81ee6c8c.
//
// Solidity: function getMyProducts(uint256 productType) view returns(uint256[] myProductIDs)
func (_Produce *ProduceCaller) GetMyProducts(opts *bind.CallOpts, productType *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Produce.contract.Call(opts, out, "getMyProducts", productType)
	return *ret0, err
}

// GetMyProducts is a free data retrieval call binding the contract method 0x81ee6c8c.
//
// Solidity: function getMyProducts(uint256 productType) view returns(uint256[] myProductIDs)
func (_Produce *ProduceSession) GetMyProducts(productType *big.Int) ([]*big.Int, error) {
	return _Produce.Contract.GetMyProducts(&_Produce.CallOpts, productType)
}

// GetMyProducts is a free data retrieval call binding the contract method 0x81ee6c8c.
//
// Solidity: function getMyProducts(uint256 productType) view returns(uint256[] myProductIDs)
func (_Produce *ProduceCallerSession) GetMyProducts(productType *big.Int) ([]*big.Int, error) {
	return _Produce.Contract.GetMyProducts(&_Produce.CallOpts, productType)
}

// GetProductPrice is a free data retrieval call binding the contract method 0x346b77f1.
//
// Solidity: function getProductPrice(address _to, uint256 productType) view returns(uint256 price)
func (_Produce *ProduceCaller) GetProductPrice(opts *bind.CallOpts, _to common.Address, productType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Produce.contract.Call(opts, out, "getProductPrice", _to, productType)
	return *ret0, err
}

// GetProductPrice is a free data retrieval call binding the contract method 0x346b77f1.
//
// Solidity: function getProductPrice(address _to, uint256 productType) view returns(uint256 price)
func (_Produce *ProduceSession) GetProductPrice(_to common.Address, productType *big.Int) (*big.Int, error) {
	return _Produce.Contract.GetProductPrice(&_Produce.CallOpts, _to, productType)
}

// GetProductPrice is a free data retrieval call binding the contract method 0x346b77f1.
//
// Solidity: function getProductPrice(address _to, uint256 productType) view returns(uint256 price)
func (_Produce *ProduceCallerSession) GetProductPrice(_to common.Address, productType *big.Int) (*big.Int, error) {
	return _Produce.Contract.GetProductPrice(&_Produce.CallOpts, _to, productType)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Produce *ProduceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Produce.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Produce *ProduceSession) Owner() (common.Address, error) {
	return _Produce.Contract.Owner(&_Produce.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Produce *ProduceCallerSession) Owner() (common.Address, error) {
	return _Produce.Contract.Owner(&_Produce.CallOpts)
}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(address owner, address producer, uint256 createdAt, uint256 batch, bool sold)
func (_Produce *ProduceCaller) Products(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner     common.Address
	Producer  common.Address
	CreatedAt *big.Int
	Batch     *big.Int
	Sold      bool
}, error) {
	ret := new(struct {
		Owner     common.Address
		Producer  common.Address
		CreatedAt *big.Int
		Batch     *big.Int
		Sold      bool
	})
	out := ret
	err := _Produce.contract.Call(opts, out, "products", arg0)
	return *ret, err
}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(address owner, address producer, uint256 createdAt, uint256 batch, bool sold)
func (_Produce *ProduceSession) Products(arg0 *big.Int) (struct {
	Owner     common.Address
	Producer  common.Address
	CreatedAt *big.Int
	Batch     *big.Int
	Sold      bool
}, error) {
	return _Produce.Contract.Products(&_Produce.CallOpts, arg0)
}

// Products is a free data retrieval call binding the contract method 0x7acc0b20.
//
// Solidity: function products(uint256 ) view returns(address owner, address producer, uint256 createdAt, uint256 batch, bool sold)
func (_Produce *ProduceCallerSession) Products(arg0 *big.Int) (struct {
	Owner     common.Address
	Producer  common.Address
	CreatedAt *big.Int
	Batch     *big.Int
	Sold      bool
}, error) {
	return _Produce.Contract.Products(&_Produce.CallOpts, arg0)
}

// Trace is a free data retrieval call binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 materialBatchNum) view returns(uint256[] ids)
func (_Produce *ProduceCaller) Trace(opts *bind.CallOpts, materialBatchNum *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Produce.contract.Call(opts, out, "trace", materialBatchNum)
	return *ret0, err
}

// Trace is a free data retrieval call binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 materialBatchNum) view returns(uint256[] ids)
func (_Produce *ProduceSession) Trace(materialBatchNum *big.Int) ([]*big.Int, error) {
	return _Produce.Contract.Trace(&_Produce.CallOpts, materialBatchNum)
}

// Trace is a free data retrieval call binding the contract method 0xa68a4fed.
//
// Solidity: function trace(uint256 materialBatchNum) view returns(uint256[] ids)
func (_Produce *ProduceCallerSession) Trace(materialBatchNum *big.Int) ([]*big.Int, error) {
	return _Produce.Contract.Trace(&_Produce.CallOpts, materialBatchNum)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0x9f96eb9d.
//
// Solidity: function registerProduct(uint256 productType, uint256 id, uint256 batchNumber, uint256[] materialBatches) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadRegisterProduct(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "registerProduct", hexutil.MustDecode(output))
	return err
}

// RegisterProduct is a free data retrieval call binding the contract method 0x9f96eb9d.
//
// Solidity: function registerProduct(uint256 productType, uint256 id, uint256 batchNumber, uint256[] materialBatches) returns()
func (_Produce *ProduceSession) ReadRegisterProduct(output string) error {
	return _Produce.Contract.ReadRegisterProduct(output)
}

// RegisterProduct is a free data retrieval call binding the contract method 0x9f96eb9d.
//
// Solidity: function registerProduct(uint256 productType, uint256 id, uint256 batchNumber, uint256[] materialBatches) returns()
func (_Produce *ProduceCallerSession) ReadRegisterProduct(output string) error {
	return _Produce.Contract.ReadRegisterProduct(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) RegisterProduct(opts *bind.TransactOpts, productType *big.Int, id *big.Int, batchNumber *big.Int, materialBatches []*big.Int) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "registerProduct", productType, id, batchNumber, materialBatches)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0x9f96eb9d.
//
// Solidity: function registerProduct(uint256 productType, uint256 id, uint256 batchNumber, uint256[] materialBatches) returns()
func (_Produce *ProduceSession) RegisterProduct(productType *big.Int, id *big.Int, batchNumber *big.Int, materialBatches []*big.Int) (*types.Transaction, error) {
	return _Produce.Contract.RegisterProduct(&_Produce.TransactOpts, productType, id, batchNumber, materialBatches)
}

// RegisterProduct is a paid mutator transaction binding the contract method 0x9f96eb9d.
//
// Solidity: function registerProduct(uint256 productType, uint256 id, uint256 batchNumber, uint256[] materialBatches) returns()
func (_Produce *ProduceTransactorSession) RegisterProduct(productType *big.Int, id *big.Int, batchNumber *big.Int, materialBatches []*big.Int) (*types.Transaction, error) {
	return _Produce.Contract.RegisterProduct(&_Produce.TransactOpts, productType, id, batchNumber, materialBatches)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadRenounceOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "renounceOwnership", hexutil.MustDecode(output))
	return err
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Produce *ProduceSession) ReadRenounceOwnership(output string) error {
	return _Produce.Contract.ReadRenounceOwnership(output)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Produce *ProduceCallerSession) ReadRenounceOwnership(output string) error {
	return _Produce.Contract.ReadRenounceOwnership(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Produce *ProduceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Produce.Contract.RenounceOwnership(&_Produce.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Produce *ProduceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Produce.Contract.RenounceOwnership(&_Produce.TransactOpts)
}

// SetMaterialContract is a paid mutator transaction binding the contract method 0x36a714f6.
//
// Solidity: function setMaterialContract(address _materialContract) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadSetMaterialContract(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "setMaterialContract", hexutil.MustDecode(output))
	return err
}

// SetMaterialContract is a free data retrieval call binding the contract method 0x36a714f6.
//
// Solidity: function setMaterialContract(address _materialContract) returns()
func (_Produce *ProduceSession) ReadSetMaterialContract(output string) error {
	return _Produce.Contract.ReadSetMaterialContract(output)
}

// SetMaterialContract is a free data retrieval call binding the contract method 0x36a714f6.
//
// Solidity: function setMaterialContract(address _materialContract) returns()
func (_Produce *ProduceCallerSession) ReadSetMaterialContract(output string) error {
	return _Produce.Contract.ReadSetMaterialContract(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) SetMaterialContract(opts *bind.TransactOpts, _materialContract common.Address) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "setMaterialContract", _materialContract)
}

// SetMaterialContract is a paid mutator transaction binding the contract method 0x36a714f6.
//
// Solidity: function setMaterialContract(address _materialContract) returns()
func (_Produce *ProduceSession) SetMaterialContract(_materialContract common.Address) (*types.Transaction, error) {
	return _Produce.Contract.SetMaterialContract(&_Produce.TransactOpts, _materialContract)
}

// SetMaterialContract is a paid mutator transaction binding the contract method 0x36a714f6.
//
// Solidity: function setMaterialContract(address _materialContract) returns()
func (_Produce *ProduceTransactorSession) SetMaterialContract(_materialContract common.Address) (*types.Transaction, error) {
	return _Produce.Contract.SetMaterialContract(&_Produce.TransactOpts, _materialContract)
}

// SetPaymentContract is a paid mutator transaction binding the contract method 0x8abb20cc.
//
// Solidity: function setPaymentContract(address _paymentContract) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadSetPaymentContract(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "setPaymentContract", hexutil.MustDecode(output))
	return err
}

// SetPaymentContract is a free data retrieval call binding the contract method 0x8abb20cc.
//
// Solidity: function setPaymentContract(address _paymentContract) returns()
func (_Produce *ProduceSession) ReadSetPaymentContract(output string) error {
	return _Produce.Contract.ReadSetPaymentContract(output)
}

// SetPaymentContract is a free data retrieval call binding the contract method 0x8abb20cc.
//
// Solidity: function setPaymentContract(address _paymentContract) returns()
func (_Produce *ProduceCallerSession) ReadSetPaymentContract(output string) error {
	return _Produce.Contract.ReadSetPaymentContract(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) SetPaymentContract(opts *bind.TransactOpts, _paymentContract common.Address) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "setPaymentContract", _paymentContract)
}

// SetPaymentContract is a paid mutator transaction binding the contract method 0x8abb20cc.
//
// Solidity: function setPaymentContract(address _paymentContract) returns()
func (_Produce *ProduceSession) SetPaymentContract(_paymentContract common.Address) (*types.Transaction, error) {
	return _Produce.Contract.SetPaymentContract(&_Produce.TransactOpts, _paymentContract)
}

// SetPaymentContract is a paid mutator transaction binding the contract method 0x8abb20cc.
//
// Solidity: function setPaymentContract(address _paymentContract) returns()
func (_Produce *ProduceTransactorSession) SetPaymentContract(_paymentContract common.Address) (*types.Transaction, error) {
	return _Produce.Contract.SetPaymentContract(&_Produce.TransactOpts, _paymentContract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadTransferOwnership(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "transferOwnership", hexutil.MustDecode(output))
	return err
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Produce *ProduceSession) ReadTransferOwnership(output string) error {
	return _Produce.Contract.ReadTransferOwnership(output)
}

// TransferOwnership is a free data retrieval call binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Produce *ProduceCallerSession) ReadTransferOwnership(output string) error {
	return _Produce.Contract.ReadTransferOwnership(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Produce *ProduceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Produce.Contract.TransferOwnership(&_Produce.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Produce *ProduceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Produce.Contract.TransferOwnership(&_Produce.TransactOpts, newOwner)
}

// TransferProducts is a paid mutator transaction binding the contract method 0xb26f3246.
//
// Solidity: function transferProducts(address from, address to, uint256 productType, uint256 count) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadTransferProducts(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "transferProducts", hexutil.MustDecode(output))
	return err
}

// TransferProducts is a free data retrieval call binding the contract method 0xb26f3246.
//
// Solidity: function transferProducts(address from, address to, uint256 productType, uint256 count) returns()
func (_Produce *ProduceSession) ReadTransferProducts(output string) error {
	return _Produce.Contract.ReadTransferProducts(output)
}

// TransferProducts is a free data retrieval call binding the contract method 0xb26f3246.
//
// Solidity: function transferProducts(address from, address to, uint256 productType, uint256 count) returns()
func (_Produce *ProduceCallerSession) ReadTransferProducts(output string) error {
	return _Produce.Contract.ReadTransferProducts(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) TransferProducts(opts *bind.TransactOpts, from common.Address, to common.Address, productType *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "transferProducts", from, to, productType, count)
}

// TransferProducts is a paid mutator transaction binding the contract method 0xb26f3246.
//
// Solidity: function transferProducts(address from, address to, uint256 productType, uint256 count) returns()
func (_Produce *ProduceSession) TransferProducts(from common.Address, to common.Address, productType *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Produce.Contract.TransferProducts(&_Produce.TransactOpts, from, to, productType, count)
}

// TransferProducts is a paid mutator transaction binding the contract method 0xb26f3246.
//
// Solidity: function transferProducts(address from, address to, uint256 productType, uint256 count) returns()
func (_Produce *ProduceTransactorSession) TransferProducts(from common.Address, to common.Address, productType *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Produce.Contract.TransferProducts(&_Produce.TransactOpts, from, to, productType, count)
}

// UpdateProductPrice is a paid mutator transaction binding the contract method 0x9c3bddb2.
//
// Solidity: function updateProductPrice(uint256 productType, uint256 newPrice) returns()

/******************************************************************************************************************************/
func (_Produce *ProduceCaller) ReadUpdateProductPrice(output string) error {
	var ()
	out := &[]interface{}{}
	err := _Produce.contract.ReadCall(out, "updateProductPrice", hexutil.MustDecode(output))
	return err
}

// UpdateProductPrice is a free data retrieval call binding the contract method 0x9c3bddb2.
//
// Solidity: function updateProductPrice(uint256 productType, uint256 newPrice) returns()
func (_Produce *ProduceSession) ReadUpdateProductPrice(output string) error {
	return _Produce.Contract.ReadUpdateProductPrice(output)
}

// UpdateProductPrice is a free data retrieval call binding the contract method 0x9c3bddb2.
//
// Solidity: function updateProductPrice(uint256 productType, uint256 newPrice) returns()
func (_Produce *ProduceCallerSession) ReadUpdateProductPrice(output string) error {
	return _Produce.Contract.ReadUpdateProductPrice(output)
}

/******************************************************************************************************************************/

func (_Produce *ProduceTransactor) UpdateProductPrice(opts *bind.TransactOpts, productType *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Produce.contract.Transact(opts, "updateProductPrice", productType, newPrice)
}

// UpdateProductPrice is a paid mutator transaction binding the contract method 0x9c3bddb2.
//
// Solidity: function updateProductPrice(uint256 productType, uint256 newPrice) returns()
func (_Produce *ProduceSession) UpdateProductPrice(productType *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Produce.Contract.UpdateProductPrice(&_Produce.TransactOpts, productType, newPrice)
}

// UpdateProductPrice is a paid mutator transaction binding the contract method 0x9c3bddb2.
//
// Solidity: function updateProductPrice(uint256 productType, uint256 newPrice) returns()
func (_Produce *ProduceTransactorSession) UpdateProductPrice(productType *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Produce.Contract.UpdateProductPrice(&_Produce.TransactOpts, productType, newPrice)
}

// ProduceEvtProductCreatedIterator is returned from FilterEvtProductCreated and is used to iterate over the raw logs and unpacked data for EvtProductCreated events raised by the Produce contract.
type ProduceEvtProductCreatedIterator struct {
	Event *ProduceEvtProductCreated // Event containing the contract specifics and raw log

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
func (it *ProduceEvtProductCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProduceEvtProductCreated)
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
		it.Event = new(ProduceEvtProductCreated)
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
func (it *ProduceEvtProductCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProduceEvtProductCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProduceEvtProductCreated represents a EvtProductCreated event raised by the Produce contract.
type ProduceEvtProductCreated struct {
	ProductType *big.Int
	ProductID   *big.Int
	Creator     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEvtProductCreated is a free log retrieval operation binding the contract event 0x7f73f7750119b7ad6911a0a4f5394ab05c7261658625bec67841058aaa682f08.
//
// Solidity: event EvtProductCreated(uint256 productType, uint256 indexed productID, address creator)
func (_Produce *ProduceFilterer) FilterEvtProductCreated(opts *bind.FilterOpts, productID []*big.Int) (*ProduceEvtProductCreatedIterator, error) {

	var productIDRule []interface{}
	for _, productIDItem := range productID {
		productIDRule = append(productIDRule, productIDItem)
	}

	logs, sub, err := _Produce.contract.FilterLogs(opts, "EvtProductCreated", productIDRule)
	if err != nil {
		return nil, err
	}
	return &ProduceEvtProductCreatedIterator{contract: _Produce.contract, event: "EvtProductCreated", logs: logs, sub: sub}, nil
}

// WatchEvtProductCreated is a free log subscription operation binding the contract event 0x7f73f7750119b7ad6911a0a4f5394ab05c7261658625bec67841058aaa682f08.
//
// Solidity: event EvtProductCreated(uint256 productType, uint256 indexed productID, address creator)
func (_Produce *ProduceFilterer) WatchEvtProductCreated(opts *bind.WatchOpts, sink chan<- *ProduceEvtProductCreated, productID []*big.Int) (event.Subscription, error) {

	var productIDRule []interface{}
	for _, productIDItem := range productID {
		productIDRule = append(productIDRule, productIDItem)
	}

	logs, sub, err := _Produce.contract.WatchLogs(opts, "EvtProductCreated", productIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProduceEvtProductCreated)
				if err := _Produce.contract.UnpackLog(event, "EvtProductCreated", log); err != nil {
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

// ParseEvtProductCreated is a log parse operation binding the contract event 0x7f73f7750119b7ad6911a0a4f5394ab05c7261658625bec67841058aaa682f08.
//
// Solidity: event EvtProductCreated(uint256 productType, uint256 indexed productID, address creator)
func (_Produce *ProduceFilterer) ParseEvtProductCreated(log types.Log) (*ProduceEvtProductCreated, error) {
	event := new(ProduceEvtProductCreated)
	if err := _Produce.contract.UnpackLog(event, "EvtProductCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProduceEvtProductOwnerChangedIterator is returned from FilterEvtProductOwnerChanged and is used to iterate over the raw logs and unpacked data for EvtProductOwnerChanged events raised by the Produce contract.
type ProduceEvtProductOwnerChangedIterator struct {
	Event *ProduceEvtProductOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProduceEvtProductOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProduceEvtProductOwnerChanged)
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
		it.Event = new(ProduceEvtProductOwnerChanged)
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
func (it *ProduceEvtProductOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProduceEvtProductOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProduceEvtProductOwnerChanged represents a EvtProductOwnerChanged event raised by the Produce contract.
type ProduceEvtProductOwnerChanged struct {
	Id       *big.Int
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvtProductOwnerChanged is a free log retrieval operation binding the contract event 0x0a04c007c049cdb8eb2c578fc38a48849cf66acce185c47f8fcf11a2957b620c.
//
// Solidity: event EvtProductOwnerChanged(uint256 indexed id, address oldOwner, address newOwner)
func (_Produce *ProduceFilterer) FilterEvtProductOwnerChanged(opts *bind.FilterOpts, id []*big.Int) (*ProduceEvtProductOwnerChangedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Produce.contract.FilterLogs(opts, "EvtProductOwnerChanged", idRule)
	if err != nil {
		return nil, err
	}
	return &ProduceEvtProductOwnerChangedIterator{contract: _Produce.contract, event: "EvtProductOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchEvtProductOwnerChanged is a free log subscription operation binding the contract event 0x0a04c007c049cdb8eb2c578fc38a48849cf66acce185c47f8fcf11a2957b620c.
//
// Solidity: event EvtProductOwnerChanged(uint256 indexed id, address oldOwner, address newOwner)
func (_Produce *ProduceFilterer) WatchEvtProductOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProduceEvtProductOwnerChanged, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Produce.contract.WatchLogs(opts, "EvtProductOwnerChanged", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProduceEvtProductOwnerChanged)
				if err := _Produce.contract.UnpackLog(event, "EvtProductOwnerChanged", log); err != nil {
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

// ParseEvtProductOwnerChanged is a log parse operation binding the contract event 0x0a04c007c049cdb8eb2c578fc38a48849cf66acce185c47f8fcf11a2957b620c.
//
// Solidity: event EvtProductOwnerChanged(uint256 indexed id, address oldOwner, address newOwner)
func (_Produce *ProduceFilterer) ParseEvtProductOwnerChanged(log types.Log) (*ProduceEvtProductOwnerChanged, error) {
	event := new(ProduceEvtProductOwnerChanged)
	if err := _Produce.contract.UnpackLog(event, "EvtProductOwnerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProduceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Produce contract.
type ProduceOwnershipTransferredIterator struct {
	Event *ProduceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProduceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProduceOwnershipTransferred)
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
		it.Event = new(ProduceOwnershipTransferred)
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
func (it *ProduceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProduceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProduceOwnershipTransferred represents a OwnershipTransferred event raised by the Produce contract.
type ProduceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Produce *ProduceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProduceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Produce.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProduceOwnershipTransferredIterator{contract: _Produce.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Produce *ProduceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProduceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Produce.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProduceOwnershipTransferred)
				if err := _Produce.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Produce *ProduceFilterer) ParseOwnershipTransferred(log types.Log) (*ProduceOwnershipTransferred, error) {
	event := new(ProduceOwnershipTransferred)
	if err := _Produce.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
