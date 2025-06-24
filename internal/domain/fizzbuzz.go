package domain

type FizzBuzzParams struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// will I need this? I don't think so
type StatsResponse struct {
	Request FizzBuzzParams `json:"request"`
	Count   int            `json:"count"`
}
