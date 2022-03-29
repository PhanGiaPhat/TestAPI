package model

import (
	"time"
)

type Rates struct {
	Date time.Time
	USD  string
	JPY  string
	BGN  string
	CZK  string
	DKK  string
	GBP  string
	HUF  string
	PLN  string
	RON  string
	SEK  string
	CHF  string
	ISK  string
	NOK  string
	HRK  string
	RUB  string
	TRY  string
	AUD  string
	BRL  string
	CAD  string
	CNY  string
	HKD  string
	IDR  string
	ILS  string
	INR  string
	KRW  string
	MXN  string
	MYR  string
	NZD  string
	PHP  string
	SGD  string
	THB  string
	ZAR  string
}

type LatestFormatted struct {
	Base  string
	Rates Ratechild
}

type Ratechild struct {
	USD string
	JPY string
	BGN string
	CZK string
	DKK string
	GBP string
	HUF string
	PLN string
	RON string
	SEK string
	CHF string
	ISK string
	NOK string
	HRK string
	RUB string
	TRY string
	AUD string
	BRL string
	CAD string
	CNY string
	HKD string
	IDR string
	ILS string
	INR string
	KRW string
	MXN string
	MYR string
	NZD string
	PHP string
	SGD string
	THB string
	ZAR string
}

type AnalyzeFormatted struct {
	Base          string
	Rates_Analyze AnalyzeRatechild
}

type AnalyzeRatechild struct {
	USD AnalyzeRatechildDetail
	JPY AnalyzeRatechildDetail
	BGN AnalyzeRatechildDetail
	CZK AnalyzeRatechildDetail
	DKK AnalyzeRatechildDetail
	GBP AnalyzeRatechildDetail
	HUF AnalyzeRatechildDetail
	PLN AnalyzeRatechildDetail
	RON AnalyzeRatechildDetail
	SEK AnalyzeRatechildDetail
	CHF AnalyzeRatechildDetail
	ISK AnalyzeRatechildDetail
	NOK AnalyzeRatechildDetail
	HRK AnalyzeRatechildDetail
	RUB AnalyzeRatechildDetail
	TRY AnalyzeRatechildDetail
	AUD AnalyzeRatechildDetail
	BRL AnalyzeRatechildDetail
	CAD AnalyzeRatechildDetail
	CNY AnalyzeRatechildDetail
	HKD AnalyzeRatechildDetail
	IDR AnalyzeRatechildDetail
	ILS AnalyzeRatechildDetail
	INR AnalyzeRatechildDetail
	KRW AnalyzeRatechildDetail
	MXN AnalyzeRatechildDetail
	MYR AnalyzeRatechildDetail
	NZD AnalyzeRatechildDetail
	PHP AnalyzeRatechildDetail
	SGD AnalyzeRatechildDetail
	THB AnalyzeRatechildDetail
	ZAR AnalyzeRatechildDetail
}

type AnalyzeRatechildDetail struct {
	Min string
	Max string
	Avg string
}
