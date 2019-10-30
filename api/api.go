package api

import (
	"context"
	"crypto/ecdsa"
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/insight-chain/inb-go/accounts/keystore"
	"github.com/insight-chain/inb-go/common"
	"github.com/insight-chain/inb-go/common/hexutil"
	"github.com/insight-chain/inb-go/core/state"
	"github.com/insight-chain/inb-go/core/types"
	"github.com/insight-chain/inb-go/crypto"
	"github.com/insight-chain/inb-sdk-go/ethclient"
	"github.com/insight-chain/inb-sdk-go/sdk-types"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"runtime"
	"strings"
)

var Client *ethclient.Client
var Conf *sdk_types.Configure
var txType types.TxType

func init() {
	Conf = InitConfig()
	Client = InitClient(Conf.Url)
}

//get current path
func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	path := strings.Split(file, "src")
	return path[0] + "src/github.com/insight-chain/inb-sdk-go/conf.json"
}

//Initialize client
func InitClient(url string) *ethclient.Client {
	Client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	return Client
}

//Initialize configuration
func InitConfig() *sdk_types.Configure {
	conf_path := CurrentFile()
	file, err := os.Open(conf_path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	conf := sdk_types.Configure{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return &conf
}

/*//Read the configuration file and convert the json file in the keystore into a private key.
func KeystoreToPrivateKey(conf *sdk_types.Configure) (string, string, error) {
	keyjson, err := ioutil.ReadFile(conf.PrivateKeyFile)
	if err != nil {
		fmt.Println("读取keyjson失败：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, conf.Password)

	if err != nil {

		return "", "", err

	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}*/

//Get privateKey by parameters
func KeystoreToPrivateKey2(privateKeyFile, password string) (string, string, error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)

	if err != nil {

		return "", "", err

	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}

//Create a new keystore
func CreateKeystore(filepath string) *keystore.KeyStore {
	ks := keystore.NewKeyStore(filepath, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks
}

//Create a new Account
func CreateAccount(ks *keystore.KeyStore, passWord string) string {
	account2, err := ks.NewAccount(passWord)
	if err != nil {
		log.Fatal(err)
	}
	return account2.Address.Hex()
}

//Read the json file from the keystore and convert it to an address
func KeystoreImport(file string, ks *keystore.KeyStore, passWord string) string {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	account, err := ks.Import(jsonBytes, passWord, passWord)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address.Hex()
}

//GenPrivKey returns a privateKey and corresponding account,preparing for transfer
func GenPrivKey() (string, string) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
	if err != nil {
		return "", ""
	}
	priv := hex.EncodeToString(privateKeyECDSA.D.Bytes())
	addr := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	return priv, addr.String()
}

//CreateTransaction returns a transaction type ,preparing for sending
func CreateTransaction(privateKey, fromAddress, toAddress, data string, value int, txtype types.TxType) (*types.Transaction, error) {
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, toAddress, amount, data, txtype)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return tx, nil
}

//Ordinary is an ordinary transfer
func Ordinary(privKeyFile, password, toAddress string, value int) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	data := ""
	txType = 1
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, toAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//fmt.Println("tx.Data:", tx.Data())
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//create a raw transaction
func NewRawTransaction(privKeyFile, password, toAddress, resourcePayer string, value int) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	data := ""
	txType = 1
	tx, err := Client.NewRawTx(chainID, nonce, privateKey, toAddress, resourcePayer, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tx, nil
}

//Sign a payment transaction
func SignPaymentTransaction(rawTxHex string, resPayerPrivFile, password string) (hexutil.Bytes, error) {
	var resourcePayerPriv string
	if len(resPayerPrivFile) == 64 {
		resourcePayerPriv = resPayerPrivFile
	} else {
		resourcePayerPriv, _, _ = KeystoreToPrivateKey2(resPayerPrivFile, password)
	}
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx, err := Client.SignPaymentTx(chainID, rawTxHex, resourcePayerPriv)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return tx.Raw, nil
}
func SendRawTransaction(rawTx string) (string, error) {
	txHash, err := Client.SendRawTx(rawTx)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return txHash, nil
}

//Send a mortage Staking,the value must > 1000000
func Staking(privKeyFile, password string, value int) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}

	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	data := ""
	txType = 2
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//Initiate Unstaking resource application
func UnStaking(privKeyFile, password string, value int) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	data := ""
	txType = 4
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//Stake on a regular basis to earn income, the following block heights are optional:
//days(30、90、180、360、720、1080、1800、3600) * 86400 / 2 (block production interval)
func TimeLimitedStaking(privKeyFile, password string, value int, data string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//data1 := "days:" + data
	txType = 3
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//vote for candidateNodes,auto use all staking in your account,
// you can vote for at least one node and at most 10 nodes.
func Vote(privKeyFile, password, toAddress string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	value := 0
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txType = 5
	data := toAddress
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//Receiving Unstaking amount , delay three days after unstaking of the application
func Receive(privKeyFile, password string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	value := 0
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txType = 7
	data := ""
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//When you have a lock record and have voted ,also the last time you received the award is more than seven days,
//you can send a transaction with the current lock record hash.
func ReceiveLockedAward(privKeyFile, password string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	accountInfo, err := GetAccountInfo(fromAddress)
	if err != nil {
		panic(err)
	}
	data := accountInfo.Stakings[0].Hash
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	value := 0
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txType = 8
	//data := data
	//data1 := "ReceiveLockedAward:" + data
	tx, err := Client.NewTransactionForRLA(chainID, nonce, privateKey, fromAddress, amount, data[:], txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//When you have voted and the last time you received the award is more than seven days,
//you can send the transaction
func ReceiveVoteAward(privKeyFile, password, toAddress string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	value := 0
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txType = 9
	data := ""
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, toAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//Receive the consumed resources once a day
func Reset(privKeyFile, password string) (string, error) {
	var privateKey, fromAddress string
	if len(privKeyFile) == 64 {
		privateKey, fromAddress = privKeyFile, password
	} else {
		privateKey, fromAddress, _ = KeystoreToPrivateKey2(privKeyFile, password)
	}
	nonce, err := Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	value := 0
	amount := big.NewInt(int64(value))
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txType = 6
	data := ""
	tx, err := Client.NewTransaction(chainID, nonce, privateKey, fromAddress, amount, data, txType)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	txHash, err := Client.SdkSendTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	return txHash, nil
}

//GetBalance returns the balance of an account
func GetBalance(address string) string {
	balance, err := Client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance.String()
}

//GetBlockNum returns the latest blockNumber
func GetBlockNum() string {
	block, err := Client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return block.Number().String()

}

//GetBlockTransactions returns all transactions in a block
func GetBlockTransactions(num int) []string {
	blockNumber := big.NewInt(int64(num))
	block, err := Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	transactions := block.Transactions()
	var txs []string
	for _, tx := range transactions {
		txs = append(txs, tx.Hash().Hex())
	}
	return txs
}

//GetTransactionByHash returns transaction details by transction hash
func GetTransactionByHash(txHash string) (tx *types.Transaction, isPending bool, err error) {
	tx, isPending, err = Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
		return nil, false, err
	}
	return tx, isPending, nil
}

//GetAccountInfo returns account related information
func GetAccountInfo(address string) (*state.Account, error) {
	fromAddr := common.HexToAddress(address)
	accountInfo, err := Client.AccountInfo(context.Background(), fromAddr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return accountInfo, nil
}
