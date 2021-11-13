package endpoint

import (
	"encoding/json"
	"finnomenaTest/model"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func GetFundsByRange(r string, s string) (model.Response, error) {
	var (
		response model.Response
		data     model.Data
		next     time.Time
	)
	url := viper.GetString("finnomena.baseurl")
	res, statusCode, err := Get(url)
	if statusCode != http.StatusOK || err != nil {
		return response, err
	}
	sort.Slice(res.Data, func(i, j int) bool {
		return res.Data[i].NavDate.Before(res.Data[j].NavDate)
	})
	for i, v := range res.Data {
		if i == 0 {
			data.Date = v.NavDate
			data.Value = []model.Funds{}
			next = addDate(data.Date, r)
		} else if !v.NavDate.Before(next) || i == len(res.Data)-1 {
			response.Data = append(response.Data, data)
			data.Date = next
			data.Value = []model.Funds{}
			next = addDate(data.Date, r)
		}
		if v.NavDate.Before(next) {
			data.Value = append(data.Value, model.Funds{
				Name:        v.ThailandFundCode,
				RankOfFund:  v.NavReturn,
				UpdatedDate: v.NavDate,
				Performance: v.NavReturn,
				Price:       v.Nav,
			})
		}
	}
	//add sort performance feature
	if s != "" {
		for _, v := range response.Data {
			sort.Slice(v.Value, func(i, j int) bool {
				if strings.ToUpper(s) == model.Min {
					return v.Value[i].Performance < v.Value[j].Performance
				}
				return v.Value[i].Performance > v.Value[j].Performance
			})
		}
	}
	return response, nil
}

func Get(path string) (body model.FinomenaResponse, statusCode int, err error) {
	res, err := http.Get(path)
	if err != nil {
		return body, statusCode, err
	}
	buf, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(buf, &body); err != nil {
		return body, statusCode, err
	}
	return body, res.StatusCode, err
}

func addDate(date time.Time, t string) time.Time {
	switch t {
	case model.Weekly:
		return date.AddDate(0, 0, 6)
	case model.Month:
		return date.AddDate(0, 1, -1)
	case model.Year:
		return date.AddDate(1, 0, -1)
	default:
		return date.AddDate(0, 0, 1)
	}
}
