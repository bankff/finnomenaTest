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
	s := flag.String("sort", "", "sort by performance")
	flag.Parse()
	if !validateRange(*r) {
		log.Fatal("range invalid")
	}
	if *s != "" && !validateSort(*s) {
		log.Fatal("sort invalid")
	}
	res, err := endpoint.GetFundsByRange(*r, *s)
	if err != nil {
		log.Printf("getfundsbyrange Error :: %v", err)
	}
	val, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Printf("json marchal Error :: %v", err)
	}
	fmt.Println(string(val))
}

func validateRange(r string) bool {
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
func validateSort(s string) bool {
	switch strings.ToUpper(s) {
	case model.Max:
		return true
	case model.Min:
		return true
	}
	return false
}
