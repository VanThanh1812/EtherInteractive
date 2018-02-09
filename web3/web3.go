package web3

import (
	"github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
	"log"
	"fmt"
)

const (
	apiUrl string = "http://localhost:8545"
)

func Call (method string, params string) (*simplejson.Json){

	postBody := `{"jsonrpc":"2.0","method":"`+method+`","params":`+params+`,"id":22}`

	fmt.Println(postBody)

	_, body, errs := gorequest.New().Post(apiUrl).Send(postBody).End()

	if errs != nil {
		panic(errs)
	}

	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		log.Fatalln(err)
	}
	return js
}
