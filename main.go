package main

import (
	"encoding/json"
	"finnomenaTest/endpoint"
	"finnomenaTest/model"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := flag.String("range", "1D", "time range")
	flag.Parse()
	if !validateflag(*r) {
		log.Fatal("range incorrect")
	}
	res, err := endpoint.GetFundsByRange(*r)
	if err != nil {
		log.Printf("getfundsbyrange Error :: %v", err)
	}
	val, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Printf("json marchal Error :: %v", err)
	}
	fmt.Println(string(val))
}

func validateflag(r string) bool {
	switch strings.ToUpper(r) {
	case model.Day:
		return true
	case model.Weekly:
		return true
	case model.Month:
		return true
	case model.Year:
		return true
	}
	return false
}
