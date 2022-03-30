package controller

import (
	"TestAPI/model"
	"TestAPI/storage"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func checkBool(err bool) {
	if !err {
		log.Println(err)
	}
}

func GetRatesLatest(c echo.Context) error {
	rates, err := GetRepoRatesLates()
	checkError(err)
	ratesfomatted, err := FormatRate(rates)
	checkError(err)
	finalresult, err := SortRate(ratesfomatted)
	checkError(err)
	return c.JSON(http.StatusOK, finalresult)
}

func GetRepoRatesLates() ([]model.Rate, error) {
	db := storage.GetDBInstance()
	rates := []model.Rate{}
	db.Where("Create_at = (?)", db.Table("Rates").Select("MAX(Create_at)")).Find(&rates)
	return rates, nil
}

func GetRatesBydate(c echo.Context) error {
	date := c.Param("date")
	datechecked, err := time.Parse("2006-01-02", date)
	checkError(err)
	rates, err := GetRepoRatesBydate(datechecked.String())
	checkError(err)
	ratesfomatted, err := FormatRate(rates)
	checkError(err)
	finalresult, err := SortRate(ratesfomatted)
	checkError(err)
	return c.JSON(http.StatusOK, finalresult)
}

func GetRepoRatesBydate(date string) ([]model.Rate, error) {
	db := storage.GetDBInstance()
	rates := []model.Rate{}
	db.Where("Create_at = (?)", date).Find(&rates)
	return rates, nil
}

func GetRatesAnalyze(c echo.Context) error {
	rates, err := GetRepoRatesAnalyze()
	checkError(err)
	ratesfomatted, err := AnalyzeFormatRate(rates)
	checkError(err)
	return c.JSON(http.StatusOK, ratesfomatted)
}

func GetRepoRatesAnalyze() ([]model.AnalyzeDetail, error) {
	db := storage.GetDBInstance()
	rates := []model.AnalyzeDetail{}
	db.Table("Rates").Select("Currency, min(rate) as min, max(rate) as max, avg(rate) as avg").Group("Currency").Scan(&rates)
	return rates, nil
}

func SortRate(rates model.RateFormatted) (map[string]interface{}, error) {
	m := structs.Map(rates)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return m, nil
}

func FormatRate(rates []model.Rate) (model.RateFormatted, error) {
	formattedrates := model.RateFormatted{Base: "EUR"}
	ratemap := make(map[string]string)
	for _, s := range rates {
		ratemap[s.Currency] = s.Rate
	}
	formattedrates.Rates = ratemap
	return formattedrates, nil
}

func AnalyzeFormatRate(rates []model.AnalyzeDetail) (model.RateAnalyze, error) {
	formattedrates := model.RateAnalyze{Base: "EUR"}
	ratemap := make(map[string]model.RateDetail)
	for _, s := range rates {
		ratemap[s.Currency] = model.RateDetail{Min: s.Min, Max: s.Max, Avg: s.Avg}
	}
	formattedrates.Rates_Analyze = ratemap
	return formattedrates, nil
}

func CrawlIfEmpty(url string) {
	response, err := http.Get(url)
	checkError(err)
	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)
	DB := storage.GetDBInstance()
	doc.Find("Cube[time]").Each(func(i int, s *goquery.Selection) {
		cubetime, exist := s.Attr("time")
		checkBool(exist)
		create_at, err := time.Parse("2006-01-02", cubetime)
		checkError(err)
		rate := model.Rate{}
		rate.Create_At = create_at
		s.Find("Cube").EachWithBreak(func(j int, t *goquery.Selection) bool {
			time, exist := t.Attr("time")
			checkBool(exist)
			if time != "" {
				return false
			}
			val_currency, exist := t.Attr("currency")
			checkBool(exist)
			val_rate, exist := t.Attr("rate")
			checkBool(exist)
			rate.Currency = val_currency
			rate.Rate = val_rate
			DB.Create(&rate)
			return true
		})
	})
}

func UpdateDB(url string) {
	DB := storage.GetDBInstance()
	var latestDate time.Time
	DB.Table("Rates").Select("MAX(Create_at)").Scan(&latestDate)
	formattedDate := latestDate.Format("2006-01-02")
	response, err := http.Get(url)
	checkError(err)
	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)
	doc.Find("Cube[time]").EachWithBreak(func(i int, s *goquery.Selection) bool {
		cubetime, exist := s.Attr("time")
		checkBool(exist)
		if formattedDate == cubetime {
			return false
		} else {
			create_at, err := time.Parse("2006-01-02", cubetime)
			checkError(err)
			rate := model.Rate{}
			rate.Create_At = create_at
			s.Find("Cube").EachWithBreak(func(j int, t *goquery.Selection) bool {
				time, exist := t.Attr("time")
				checkBool(exist)
				if time != "" {
					return false
				}
				val_currency, exist := t.Attr("currency")
				checkBool(exist)
				val_rate, exist := t.Attr("rate")
				checkBool(exist)
				rate.Currency = val_currency
				rate.Rate = val_rate
				DB.Create(&rate)
				return true
			})
		}
		return true
	})
}
