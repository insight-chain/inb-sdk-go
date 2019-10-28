package main

import (
	"github.com/insight-chain/inb-go/core/types"
	"github.com/insight-chain/inb-sdk-go/api"
	"github.com/insight-chain/inb-sdk-go/sdk-types"
)

//var Client *ethclient.Client
var conf *sdk_types.Configure
var txType types.TxType

func init() {
	conf = api.InitConfig()
}

func main() {
	/*privKey, address, _ := api.KeystoreToPrivateKey(conf)
	fmt.Printf("privKey:%s\naddress:%s\n", privKey, address)*/
	//test privateKey:9d33ea78c111cf8e1af71475c36d85a4267272b6a2813ba94ebb82863e49f1d7
	//test address:0x9591fe43dad2fc1d589cfe8f7e4119bce7f93021
	//test account is to suocang

	//new privateKey:27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381
	//new account:0x953828E2639097f13e52C91339af771B0D0a2a53
	//blockHeader, _ := api.Client.HeaderByNumber(context.Background(), big.NewInt(1))
	//fmt.Println("blockHeader:", blockHeader)
	//pendingTxCount, _ := api.Client.PendingTransactionCount(context.Background())
	//fmt.Println("pendingTxCount", pendingTxCount)
	//fromAddress := "0x952180cc82c09745a6a0dd338cecc9758619dcea"
	//pendingNounceAt, _ := api.Client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
	//fmt.Println("pendingNounceAt:", pendingNounceAt)
	//pendingCodeAt, _ := api.Client.PendingCodeAt(context.Background(), common.HexToAddress(fromAddress))
	//fmt.Println("pendingCodeAt:", pendingCodeAt)
	//pendingNetAt, _ := api.Client.PendingNetAt(context.Background(), common.HexToAddress(fromAddress))
	//fmt.Println("pendingNetAt:", pendingNetAt)
	//netAt, _ := api.Client.NetAt(context.Background(), common.HexToAddress(fromAddress), big.NewInt(77694))
	//fmt.Println("netAt:", netAt)
	//syncProgress, _ := api.Client.SyncProgress(context.Background())
	//fmt.Println("syncProgress:", syncProgress)
	//txReceipt, _ := api.Client.TransactionReceipt(context.Background(), common.HexToHash("0x2372224df2530805f4fdfa35174f648218f66c7284ae94a129c5f59918a0d4c0"))
	//fmt.Println("txReceipt:", txReceipt)
	//txInBlock, _ := api.Client.TransactionInBlock(context.Background(), common.HexToHash(
	//	"0xcc37b18f50029f7669df78b3cfba91dea20f855f49c89bbe1481d1e14e4256af"), 1)
	//fmt.Println("txInBlock:", txInBlock)
	//tx, _, _ := api.Client.TransactionByHash(context.Background(), common.HexToHash("0x57421d1011330ef3dee01083f0074ef82527ba84f71f09c3a7a0baaf5a7bde42"))
	//txSender, _ := api.Client.TransactionSender(context.Background(), tx, common.HexToHash(
	//	"0xcc37b18f50029f7669df78b3cfba91dea20f855f49c89bbe1481d1e14e4256af"), 1)
	//fmt.Println("txSender:", txSender.String())
	//headerByNum, _ := api.Client.HeaderByNumber(context.Background(), big.NewInt(0))
	//fmt.Println("headerByNum:", headerByNum)
	//blockByHah, _ := api.Client.BlockByHash(context.Background(), common.HexToHash("0xcc37b18f50029f7669df78b3cfba91dea20f855f49c89bbe1481d1e14e4256af"))
	//fmt.Println("blockByHah:", blockByHah)

	/*//send normal transaction
	privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd42"
	value := 1
	tx, _ := api.Ordinary(privKey, fromAddress, toAddress, value)
	fmt.Println("tx:", tx)*/
	/*//send normal transaction NEW!
	privKeyFile := "/home/GO_PATH/src/github.com/insight-chain/data5/keystore/UTC--2019-10-23T08-38-18.294495121Z--957370d3a7a3ffb3df1afd67f0078dce8f96873c"
	password := "1"
	toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd43"
	value := 1
	tx, _ := api.Ordinary(privKeyFile, password, toAddress, value)
	fmt.Println("tx:", tx)*/
	/*//send payment signature transaction
	privKey := "2306a2b439e0e843b56c725c99148ac541a96f3a6d69795e757a86f07d4a8614"
	fromAddress := "0x957370D3A7a3fFb3df1aFD67f0078DCe8F96873c"
	toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd42"
	resourcePayment := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	value := big.NewInt(70000)
	nonce, _ := api.Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	chainId, _ := api.Client.NetworkID(context.Background())
	tx, _ := api.Client.SignPaymentTransaction(chainId, nonce, privKey, toAddress, resourcePayment, value, "", 1)
	fmt.Println("tx:", tx.Raw)
	rawTx := "0xf8bd0a949568dbb7f09c38ee29322ccdedf657ec7e4cdd428301117080820731a07e47b804a3b30906b5cbd5c7b87069427f57f887002e98a3a28646fad6b18a94a05d411653c96fdedfd7157f777cdad20046515644e915cc9a0f93940518f33e2f01f85a94953828e2639097f13e52c91339af771b0d0a2a53820732a0aac0ffe29007f86d8df44fcae648f46f219908a59f5f363be1df603422dd37eca02bc1cc4c41f94153a182b4c81fec752c647a57c0949b2f9f4ec3d72f8f784342"
	txHash, _ := api.Client.SendRawTransaction(rawTx)
	fmt.Println("txHash:", txHash)*/

	//create a raw Tx
	//privKey := "2306a2b439e0e843b56c725c99148ac541a96f3a6d69795e757a86f07d4a8614"
	//fromAddress := "0x957370D3A7a3fFb3df1aFD67f0078DCe8F96873c"
	//toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd42"
	//resourcePayment := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	//resPayerAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	//value := 70000
	//rawTx, _ := api.NewRawTransaction(privKey, fromAddress, toAddress, resPayerAddress, value)
	//fmt.Println("rawTx:", rawTx)
	//signRawTx, _ := api.SignPaymentTransaction(rawTx, resourcePayment)
	//fmt.Println("signRawTx:", signRawTx)
	//signRawTx2 := "0xf8bd0d949568dbb7f09c38ee29322ccdedf657ec7e4cdd428301117080820731a0b5cc6ba08a5b1b327351771cc60342769c2fd2cbae97326a162cec36d2397d86a03d31624e14eda4961956a1e73b4901ef59040cd270995a6a4d42d4bf6cd7e2d501f85a94953828e2639097f13e52c91339af771b0d0a2a53820731a0ad04d6285e236d7ff0eba1f8d996ea902604f6d3d30973a689f932a45e6f5e59a018bdbf6d4c3f536c52af35b6bba0bae9001b285aedbfbbced3e50ecbe49360a6"
	//txHash, _ := api.SendRawTransaction(signRawTx2)
	//fmt.Println("txHash:", txHash)

	/*//test new raw transaction
	privKey := "2306a2b439e0e843b56c725c99148ac541a96f3a6d69795e757a86f07d4a8614"
	fromAddress := "0x957370D3A7a3fFb3df1aFD67f0078DCe8F96873c"
	toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd42"
	resourcePayerPri := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	resourcePayerAddr := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	value := big.NewInt(700)
	nonce, _ := api.Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	chainId, _ := api.Client.NetworkID(context.Background())
	rawTxHex, _ := api.Client.NewRawTransaction(chainId, nonce, privKey, toAddress, resourcePayerAddr, value, "", 1)
	tx, _ := api.Client.SignPaymentTransaction(chainId, rawTxHex, resourcePayerPri)
	fmt.Println("rawTx:", tx.Raw)*/
	/*rawTx := "0xf8bb0c949568dbb7f09c38ee29322ccdedf657ec7e4cdd428202bc80820731a01f0f8da2f2f6474dbac50322e80db9dfc8b61192a7f7a952a145e54ff3873c74a0286217bb0a5e027ac760495a1fb1d4fcef525c5647b8a635a445b16612ba960e01f85994953828e2639097f13e52c91339af771b0d0a2a53820732a09673c7569b78888dec520029dce4b8dd7bae39ed2a66ded9c44d3b7ae102eac09f27804e6cebfe51a706666d3437ac616bcc99136f0098e8ee397955f5aad29e"
	txHash, _ := api.Client.SendRawTransaction(rawTx)
	fmt.Println("txHash:", txHash)*/

	/*//test rawTx
	privKey := "2306a2b439e0e843b56c725c99148ac541a96f3a6d69795e757a86f07d4a8614"

	fromAddress := "0x957370D3A7a3fFb3df1aFD67f0078DCe8F96873c"
	toAddress := "0x9568dbb7f09c38ee29322ccdedf657ec7e4cdd42"
	nonce, _ := api.Client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	chainId, _ := api.Client.NetworkID(context.Background())
	//resourcePayment := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	value := big.NewInt(70000)
	signTx, _ := api.Client.RawTransaction(chainId, nonce, privKey, toAddress, value, "", 1)
	ts := types.Transactions{signTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Printf(rawTxHex)
	rawTxBytes2, _ := hex.DecodeString(rawTxHex)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes2, &tx)
	err := api.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())*/

	/*//send mortage transaction
	privKeyFile := "/home/GO_PATH/src/github.com/insight-chain/data5/keystore/UTC--2019-10-23T08-38-18.294495121Z--957370d3a7a3ffb3df1afd67f0078dce8f96873c"
	password := "1"
	value := 1000000
	tx, _ := api.Staking(privKeyFile, password, value)
	fmt.Println("tx:", tx)*/
	/*//send Redeem transaction
	privKey := "16927473f4d24b40529b55cf67ed64de1d2a5edb709b2bcbcdc0b4dee4e69d2c"
	fromAddress := "0x952180cc82c09745a6a0dd338cecc9758619dcea"
	value := 100000
	tx, _ := api.UnStaking(privKey, fromAddress, value)
	fmt.Println("tx:", tx)*/
	/*//send staking transaction
	privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	value := 1000000
	data := "1296000"
	tx, _ := api.TimeLimitedStaking(privKey, fromAddress, value, data)
	fmt.Println("tx:", tx)*/
	/*//send vote transaction
	privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	toAddress := "0x957370D3A7a3fFb3df1aFD67f0078DCe8F96873c"
	tx, _ := api.Vote(privKey, fromAddress, toAddress)
	fmt.Println("tx:", tx)*/
	/*//send Receive transaction
	privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	//toAddress := "0x957c9C2EE36D1eE7B07cFf6aC1Cdd465869D6Ef9"
	tx, _ := api.Receive(privKey, fromAddress)
	fmt.Println("tx:", tx)*/
	// ReceiveLockedAward uncompleted txType=8

	/*privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	//data := "d2acc22e5d2cd5d4eff635610cae6c228c5e90b51d1fb391879e85cb7d75b6a2"
	accountInfo, err := api.GetAccountInfo(fromAddress)
	if err != nil {
		panic(err)
	}
	data:=accountInfo.Stakings[0].Hash
	tx, _ := api.ReceiveLockedAward(privKey, fromAddress, data)
	fmt.Println("tx:", tx)*/
	/*// ReceiveVoteAward  txType=9
	privKey := "27cc9861d41b67b5202ae98da3df62d14a831cf9806ddd1515a9310cf597c381"
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	toAddress := "0x9591fe43dad2fc1d589cfe8f7e4119bce7f93021"
	tx, _ := api.ReceiveVoteAward(privKey, fromAddress, toAddress)
	fmt.Println("tx:", tx)*/
	//send reset transaction
	/*privKey := "16927473f4d24b40529b55cf67ed64de1d2a5edb709b2bcbcdc0b4dee4e69d2c"
	fromAddress := "0x952180cc82c09745a6a0dd338cecc9758619dcea"
	//toAddress := "0x957c9C2EE36D1eE7B07cFf6aC1Cdd465869D6Ef9"
	tx, _ := Reset(privKey, fromAddress)
	fmt.Println("tx:", tx)*/
	//
	/*//getBalance
	balance := GetBalance("0x95e7f9263b6d51f49da1ca309d04ebc7b4aad94c")
	fmt.Println("balance:", balance)*/
	/*//get blockNum
	blockNum := GetBlockNum()
	fmt.Println("blockNum:", blockNum)*/
	/*// get transactions in block
	transactions := GetBlockTransactions(994)
	for _, tx := range transactions {
		fmt.Println(tx)
	}*/
	/*//get Transaction by hash
	tx3, isPending, err := GetTransactionByHash("0xba604bbb2cfa34205ef916c286d52c41311258fc4e224220b82f6111a90e7176")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx3：%v\n", tx3)
	fmt.Println("isPending:", isPending)*/
	/*//create an account
	privKey, fromaddress := api.GenPrivKey()
	fmt.Printf("privKey:%s\naddress:%s\n", privKey, fromaddress)*/

	/*//view blockHeader //err:missing required field 'vdposContext' for Header
	header, err := api.Client.HeaderByNumber(context.Background(), nil)
	fmt.Println("here........")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header)*/

	/*//var accountInfo *Account
	fromAddress := "0x953828E2639097f13e52C91339af771B0D0a2a53"
	accountInfo, err := api.GetAccountInfo(fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("str:=", accountInfo.Stakings[0].Hash.UnmarshalJSON())
	//fmt.Println("str:=", string(accountInfo.Stakings[0].Hash))
	fmt.Println("str:=", accountInfo.Stakings[0].Hash.String())
	fmt.Println("str:=", []byte(accountInfo.Stakings[0].Hash.String()))

	fmt.Println("Balance", accountInfo.Balance)
	fmt.Println("codehash", accountInfo.CodeHash)
	fmt.Println("nonce", accountInfo.Nonce)
	fmt.Println("res", accountInfo.Res)
	fmt.Println("root", accountInfo.Root.String())*/
}
