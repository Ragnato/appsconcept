package domain

type FizzBuzzParams struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// need to create this on a dto package
type StatsResponse struct {
	Request FizzBuzzParams `json:"request"`
	Count   int            `json:"count"`
}
