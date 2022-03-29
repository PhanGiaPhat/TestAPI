package controller

import (
	"TestAPI/model"
	"TestAPI/storage"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"time"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

func GetRatesLatest(c echo.Context) error {
	rates, _ := GetRepoRatesLates()
	newrates, _ := FormatRate(rates)
	m, _ := SortRate(newrates)
	return c.JSON(http.StatusOK, m)
}

func SortRate(rates model.LatestFormatted) (map[string]interface{}, error) {

	m := structs.Map(rates)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return m, nil
}

func FormatRate(rates model.Rates) (model.LatestFormatted, error) {
	formattedrates := model.LatestFormatted{Base: "EUR"}

	e := reflect.ValueOf(&rates).Elem()

	for i := 1; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()
		val := fmt.Sprintf("%v", varValue)
		switch varName {
		case "USD":
			formattedrates.Rates.USD = val
		case "JPY":
			formattedrates.Rates.JPY = val
		case "BGN":
			formattedrates.Rates.BGN = val
		case "CZK":
			formattedrates.Rates.CZK = val
		case "DKK":
			formattedrates.Rates.DKK = val
		case "GBP":
			formattedrates.Rates.GBP = val
		case "HUF":
			formattedrates.Rates.HUF = val
		case "PLN":
			formattedrates.Rates.PLN = val
		case "RON":
			formattedrates.Rates.RON = val
		case "SEK":
			formattedrates.Rates.SEK = val
		case "CHF":
			formattedrates.Rates.CHF = val
		case "ISK":
			formattedrates.Rates.ISK = val
		case "NOK":
			formattedrates.Rates.NOK = val
		case "HRK":
			formattedrates.Rates.HRK = val
		case "RUB":
			formattedrates.Rates.RUB = val
		case "TRY":
			formattedrates.Rates.TRY = val
		case "AUD":
			formattedrates.Rates.AUD = val
		case "BRL":
			formattedrates.Rates.BRL = val
		case "CAD":
			formattedrates.Rates.CAD = val
		case "CNY":
			formattedrates.Rates.CNY = val
		case "HKD":
			formattedrates.Rates.HKD = val
		case "IDR":
			formattedrates.Rates.IDR = val
		case "ILS":
			formattedrates.Rates.ILS = val
		case "INR":
			formattedrates.Rates.INR = val
		case "KRW":
			formattedrates.Rates.KRW = val
		case "MXN":
			formattedrates.Rates.MXN = val
		case "MYR":
			formattedrates.Rates.MYR = val
		case "NZD":
			formattedrates.Rates.NZD = val
		case "PHP":
			formattedrates.Rates.PHP = val
		case "SGD":
			formattedrates.Rates.SGD = val
		case "THB":
			formattedrates.Rates.THB = val
		case "ZAR":
			formattedrates.Rates.ZAR = val
		}

	}
	return formattedrates, nil
}

func GetRepoRatesLates() (model.Rates, error) {
	db := storage.GetDBInstance()
	rates := model.Rates{}
	db.Last(&rates)
	return rates, nil
}

func GetRatesBydate(c echo.Context) error {
	date := c.Param("date")

	datechecked, err := time.Parse("2006-01-02", date)
	if err != nil {
		return c.String(http.StatusOK, "Wrong Parameter")
	}

	rates, _ := GetRepoRatesBydate(datechecked.String())

	newrates, _ := FormatRate(rates)
	m, _ := SortRate(newrates)
	return c.JSON(http.StatusOK, m)
}

func GetRepoRatesBydate(date string) (model.Rates, error) {
	db := storage.GetDBInstance()
	rates := model.Rates{}

	if err := db.First(&rates, "date = ?", date).Error; err != nil {
		return rates, nil
	}
	db.First(&rates, "date = ?", date)
	return rates, nil
}

func GetRatesAnalyze(c echo.Context) error {
	rates, _ := GetRepoRatesAnalyze()
	m, _ := SortRateAnalyze(rates)
	return c.JSON(http.StatusOK, m)
}

func GetRepoRatesAnalyze() (model.AnalyzeFormatted, error) {

	rates := model.AnalyzeFormatted{Base: "EUR"}

	e := reflect.ValueOf(&rates.Rates_Analyze).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		switch varName {
		case "USD":
			rates.Rates_Analyze.USD.Min = getMin("USD")
			rates.Rates_Analyze.USD.Max = getMax("USD")
			rates.Rates_Analyze.USD.Avg = getAvg("USD")
		case "JPY":
			rates.Rates_Analyze.JPY.Min = getMin("JPY")
			rates.Rates_Analyze.JPY.Max = getMax("JPY")
			rates.Rates_Analyze.JPY.Avg = getAvg("JPY")
		case "BGN":
			rates.Rates_Analyze.BGN.Min = getMin("BGN")
			rates.Rates_Analyze.BGN.Max = getMax("BGN")
			rates.Rates_Analyze.BGN.Avg = getAvg("BGN")
		case "CZK":
			rates.Rates_Analyze.CZK.Min = getMin("CZK")
			rates.Rates_Analyze.CZK.Max = getMax("CZK")
			rates.Rates_Analyze.CZK.Avg = getAvg("CZK")
		case "DKK":
			rates.Rates_Analyze.DKK.Min = getMin("DKK")
			rates.Rates_Analyze.DKK.Max = getMax("DKK")
			rates.Rates_Analyze.DKK.Avg = getAvg("DKK")
		case "GBP":
			rates.Rates_Analyze.GBP.Min = getMin("GBP")
			rates.Rates_Analyze.GBP.Max = getMax("GBP")
			rates.Rates_Analyze.GBP.Avg = getAvg("GBP")
		case "HUF":
			rates.Rates_Analyze.HUF.Min = getMin("HUF")
			rates.Rates_Analyze.HUF.Max = getMax("HUF")
			rates.Rates_Analyze.HUF.Avg = getAvg("HUF")
		case "PLN":
			rates.Rates_Analyze.PLN.Min = getMin("PLN")
			rates.Rates_Analyze.PLN.Max = getMax("PLN")
			rates.Rates_Analyze.PLN.Avg = getAvg("PLN")
		case "RON":
			rates.Rates_Analyze.RON.Min = getMin("RON")
			rates.Rates_Analyze.RON.Max = getMax("RON")
			rates.Rates_Analyze.RON.Avg = getAvg("RON")
		case "SEK":
			rates.Rates_Analyze.SEK.Min = getMin("SEK")
			rates.Rates_Analyze.SEK.Max = getMax("SEK")
			rates.Rates_Analyze.SEK.Avg = getAvg("SEK")
		case "CHF":
			rates.Rates_Analyze.CHF.Min = getMin("CHF")
			rates.Rates_Analyze.CHF.Max = getMax("CHF")
			rates.Rates_Analyze.CHF.Avg = getAvg("CHF")
		case "ISK":
			rates.Rates_Analyze.ISK.Min = getMin("ISK")
			rates.Rates_Analyze.ISK.Max = getMax("ISK")
			rates.Rates_Analyze.ISK.Avg = getAvg("ISK")
		case "NOK":
			rates.Rates_Analyze.NOK.Min = getMin("NOK")
			rates.Rates_Analyze.NOK.Max = getMax("NOK")
			rates.Rates_Analyze.NOK.Avg = getAvg("NOK")
		case "HRK":
			rates.Rates_Analyze.HRK.Min = getMin("HRK")
			rates.Rates_Analyze.HRK.Max = getMax("HRK")
			rates.Rates_Analyze.HRK.Avg = getAvg("HRK")
		case "RUB":
			rates.Rates_Analyze.RUB.Min = getMin("RUB")
			rates.Rates_Analyze.RUB.Max = getMax("RUB")
			rates.Rates_Analyze.RUB.Avg = getAvg("RUB")
		case "TRY":
			rates.Rates_Analyze.TRY.Min = getMin("TRY")
			rates.Rates_Analyze.TRY.Max = getMax("TRY")
			rates.Rates_Analyze.TRY.Avg = getAvg("TRY")
		case "AUD":
			rates.Rates_Analyze.AUD.Min = getMin("AUD")
			rates.Rates_Analyze.AUD.Max = getMax("AUD")
			rates.Rates_Analyze.AUD.Avg = getAvg("AUD")
		case "BRL":
			rates.Rates_Analyze.BRL.Min = getMin("BRL")
			rates.Rates_Analyze.BRL.Max = getMax("BRL")
			rates.Rates_Analyze.BRL.Avg = getAvg("BRL")
		case "CAD":
			rates.Rates_Analyze.CAD.Min = getMin("CAD")
			rates.Rates_Analyze.CAD.Max = getMax("CAD")
			rates.Rates_Analyze.CAD.Avg = getAvg("CAD")
		case "CNY":
			rates.Rates_Analyze.CNY.Min = getMin("CNY")
			rates.Rates_Analyze.CNY.Max = getMax("CNY")
			rates.Rates_Analyze.CNY.Avg = getAvg("CNY")
		case "HKD":
			rates.Rates_Analyze.HKD.Min = getMin("HKD")
			rates.Rates_Analyze.HKD.Max = getMax("HKD")
			rates.Rates_Analyze.HKD.Avg = getAvg("HKD")
		case "IDR":
			rates.Rates_Analyze.IDR.Min = getMin("IDR")
			rates.Rates_Analyze.IDR.Max = getMax("IDR")
			rates.Rates_Analyze.IDR.Avg = getAvg("IDR")
		case "ILS":
			rates.Rates_Analyze.ILS.Min = getMin("ILS")
			rates.Rates_Analyze.ILS.Max = getMax("ILS")
			rates.Rates_Analyze.ILS.Avg = getAvg("ILS")
		case "INR":
			rates.Rates_Analyze.INR.Min = getMin("INR")
			rates.Rates_Analyze.INR.Max = getMax("INR")
			rates.Rates_Analyze.INR.Avg = getAvg("INR")
		case "KRW":
			rates.Rates_Analyze.KRW.Min = getMin("KRW")
			rates.Rates_Analyze.KRW.Max = getMax("KRW")
			rates.Rates_Analyze.KRW.Avg = getAvg("KRW")
		case "MXN":
			rates.Rates_Analyze.MXN.Min = getMin("MXN")
			rates.Rates_Analyze.MXN.Max = getMax("MXN")
			rates.Rates_Analyze.MXN.Avg = getAvg("MXN")
		case "MYR":
			rates.Rates_Analyze.MYR.Min = getMin("MYR")
			rates.Rates_Analyze.MYR.Max = getMax("MYR")
			rates.Rates_Analyze.MYR.Avg = getAvg("MYR")
		case "NZD":
			rates.Rates_Analyze.NZD.Min = getMin("NZD")
			rates.Rates_Analyze.NZD.Max = getMax("NZD")
			rates.Rates_Analyze.NZD.Avg = getAvg("NZD")
		case "PHP":
			rates.Rates_Analyze.PHP.Min = getMin("PHP")
			rates.Rates_Analyze.PHP.Max = getMax("PHP")
			rates.Rates_Analyze.PHP.Avg = getAvg("PHP")
		case "SGD":
			rates.Rates_Analyze.SGD.Min = getMin("SGD")
			rates.Rates_Analyze.SGD.Max = getMax("SGD")
			rates.Rates_Analyze.SGD.Avg = getAvg("SGD")
		case "THB":
			rates.Rates_Analyze.THB.Min = getMin("THB")
			rates.Rates_Analyze.THB.Max = getMax("THB")
			rates.Rates_Analyze.THB.Avg = getAvg("THB")
		case "ZAR":
			rates.Rates_Analyze.ZAR.Min = getMin("ZAR")
			rates.Rates_Analyze.ZAR.Max = getMax("ZAR")
			rates.Rates_Analyze.ZAR.Avg = getAvg("ZAR")
		}

	}

	return rates, nil
}

func SortRateAnalyze(rates model.AnalyzeFormatted) (map[string]interface{}, error) {

	m := structs.Map(rates)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return m, nil
}

func getMin(ratein string) (valout string) {
	db := storage.GetDBInstance()
	col := ratein
	var sql string
	sql = "SELECT MIN(" + col + ") as min FROM rates"
	db.Raw(sql).Scan(&valout)
	return valout
}

func getMax(ratein string) (valout string) {
	db := storage.GetDBInstance()
	col := ratein
	var sql string
	sql = "SELECT Max(" + col + ") as max FROM rates"
	db.Raw(sql).Scan(&valout)
	return valout
}

func getAvg(ratein string) (valout string) {
	db := storage.GetDBInstance()
	col := ratein
	var sql string
	sql = "SELECT AVG(" + col + ") as avg FROM rates"
	db.Raw(sql).Scan(&valout)
	return valout
}
