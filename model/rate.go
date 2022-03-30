package model

import "time"

type Rate struct {
	Create_At time.Time
	Currency  string
	Rate      string
}

type RateFormatted struct {
	Base  string
	Rates map[string]string
}

type AnalyzeDetail struct {
	Currency string
	Min      string
	Max      string
	Avg      string
}

type RateDetail struct {
	Min string
	Max string
	Avg string
}

type RateAnalyze struct {
	Base          string
	Rates_Analyze map[string]RateDetail
}
