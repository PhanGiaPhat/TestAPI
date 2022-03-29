package main

import (
	controller "TestAPI/controller"
	"TestAPI/model"
	storage "TestAPI/storage"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	storage.NewDB()
	DB := storage.GetDBInstance()

	url := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

	response, error := http.Get(url)
	checkError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	checkError(error)

	doc.Find("Cube[time]").Each(func(i int, s *goquery.Selection) {

		cubetime, exists := s.Attr("time")
		if !exists {
			cubetime = ""
		}
		newtime, err := time.Parse("2006-01-02", cubetime)
		if err != nil {
		}

		rates := model.Rates{}

		USD, _ := s.Find("Cube[currency='USD']").Attr("rate")
		JPY, _ := s.Find("Cube[currency='JPY']").Attr("rate")
		BGN, _ := s.Find("Cube[currency='BGN']").Attr("rate")
		CZK, _ := s.Find("Cube[currency='CZK']").Attr("rate")
		DKK, _ := s.Find("Cube[currency='DKK']").Attr("rate")
		GBP, _ := s.Find("Cube[currency='GBP']").Attr("rate")
		HUF, _ := s.Find("Cube[currency='HUF']").Attr("rate")
		PLN, _ := s.Find("Cube[currency='PLN']").Attr("rate")
		RON, _ := s.Find("Cube[currency='RON']").Attr("rate")
		SEK, _ := s.Find("Cube[currency='SEK']").Attr("rate")
		CHF, _ := s.Find("Cube[currency='CHF']").Attr("rate")
		ISK, _ := s.Find("Cube[currency='ISK']").Attr("rate")
		NOK, _ := s.Find("Cube[currency='NOK']").Attr("rate")
		HRK, _ := s.Find("Cube[currency='HRK']").Attr("rate")
		RUB, _ := s.Find("Cube[currency='RUB']").Attr("rate")
		TRY, _ := s.Find("Cube[currency='TRY']").Attr("rate")
		AUD, _ := s.Find("Cube[currency='AUD']").Attr("rate")
		BRL, _ := s.Find("Cube[currency='BRL']").Attr("rate")
		CAD, _ := s.Find("Cube[currency='CAD']").Attr("rate")
		CNY, _ := s.Find("Cube[currency='CNY']").Attr("rate")
		HKD, _ := s.Find("Cube[currency='HKD']").Attr("rate")
		IDR, _ := s.Find("Cube[currency='IDR']").Attr("rate")
		ILS, _ := s.Find("Cube[currency='ILS']").Attr("rate")
		INR, _ := s.Find("Cube[currency='INR']").Attr("rate")
		KRW, _ := s.Find("Cube[currency='KRW']").Attr("rate")
		MXN, _ := s.Find("Cube[currency='MXN']").Attr("rate")
		MYR, _ := s.Find("Cube[currency='MYR']").Attr("rate")
		NZD, _ := s.Find("Cube[currency='NZD']").Attr("rate")
		PHP, _ := s.Find("Cube[currency='PHP']").Attr("rate")
		SGD, _ := s.Find("Cube[currency='SGD']").Attr("rate")
		THB, _ := s.Find("Cube[currency='THB']").Attr("rate")
		ZAR, _ := s.Find("Cube[currency='ZAR']").Attr("rate")

		rates.Date = newtime
		rates.USD = USD
		rates.JPY = JPY
		rates.BGN = BGN
		rates.CZK = CZK
		rates.DKK = DKK
		rates.GBP = GBP
		rates.HUF = HUF
		rates.PLN = PLN
		rates.RON = RON
		rates.SEK = SEK
		rates.CHF = CHF
		rates.ISK = ISK
		rates.NOK = NOK
		rates.HRK = HRK
		rates.RUB = RUB
		rates.TRY = TRY
		rates.AUD = AUD
		rates.BRL = BRL
		rates.CAD = CAD
		rates.CNY = CNY
		rates.HKD = HKD
		rates.IDR = IDR
		rates.ILS = ILS
		rates.INR = INR
		rates.KRW = KRW
		rates.MXN = MXN
		rates.MYR = MYR
		rates.NZD = NZD
		rates.PHP = PHP
		rates.SGD = SGD
		rates.THB = THB
		rates.ZAR = ZAR

		DB.Create(rates)

	})

	e.GET("/", hello)
	e.GET("/rates/analyze", controller.GetRatesAnalyze)
	e.GET("/rates/latest", controller.GetRatesLatest)
	e.GET("/rates/:date", controller.GetRatesBydate)
	e.Logger.Fatal(e.Start(":1323"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
