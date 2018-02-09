package main

import (
	"fmt"
	"SmartContract/web3"
	"bufio"
	"os"
	"encoding/json"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main(){
	account := coinbase()
	fmt.Println("account: ", account)

	fmt.Println("---\n")
	listacc := accounts()
	fmt.Println("list account: ", listacc)

	fmt.Println("---\n")

	newAcc := newAccount()
	fmt.Println("new addr: ", newAcc)

	fmt.Println("---\n")

	res := sendEthFromCoinbase()
	fmt.Println("---\n")
	fmt.Println("Transaction: ",res)
}
func newAccount() interface{} {
	fmt.Println("Input pass for new account")
	pass:= "12345678"
	res := web3.Call("personal_newAccount", `["`+pass+`"]`).Get("result").MustString()
	return res
}

func sendEthFromCoinbase () string {
	// from: coinbase
	// to: addr
	// value: 0x....
	// optinal ...
	fmt.Print("Input to addr: ")
	toAddr,_ := reader.ReadString('\n')
	toAddr = strings.Replace(toAddr, "\n", "", -1)

	fmt.Print("Input value: ")
	value,_ := reader.ReadString('\n')
	value = strings.Replace(value, "\n", "", -1)

	param := createStructTransaction(coinbase(), toAddr, value)
	res2, _ := web3.Call("personal_sendTransaction", param).Encode()
	return string(res2)
}

func accounts() interface{} {
	res := web3.Call("eth_accounts", "[]").Get("result").MustArray()
	return res
}
func coinbase() string {
	res := web3.Call("eth_coinbase", "[]").Get("result").MustString()
	return res
}

func createStructTransaction (from string, to string, value string) string {
	item := ItemTransation{
		From: from,
		To: to,
		Value: value,
	}
	item_json, _ := json.Marshal(item)
	fmt.Println("item: ", string(item_json))
	params := `[`+string(item_json)+`,"thanhpv1234"]`
	return params
}

type ItemTransation struct {
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
}