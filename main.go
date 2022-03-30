package main

import (
	controller "TestAPI/controller"
	"TestAPI/model"
	storage "TestAPI/storage"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {

	url := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

	e := echo.New()
	storage.NewDB()

	db := storage.GetDBInstance()

	dbcheckrate := model.Rate{}
	dbcheck := db.Take(&dbcheckrate)

	if errors.Is(dbcheck.Error, gorm.ErrRecordNotFound) {
		controller.CrawlIfEmpty(url)
	} else {
		controller.UpdateDB(url)
	}

	e.GET("/", hello)
	e.GET("/rates/analyze", controller.GetRatesAnalyze)
	e.GET("/rates/latest", controller.GetRatesLatest)
	e.GET("/rates/:date", controller.GetRatesBydate)
	e.Logger.Fatal(e.Start(":1323"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
