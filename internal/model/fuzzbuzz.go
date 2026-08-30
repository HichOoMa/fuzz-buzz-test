package model

type FuzzBuzzRequest struct {
	Int1  int    `query:"int1"`
	Int2  int    `query:"int2"`
	Limit int    `query:"limit"`
	Str1  string `query:"str1"`
	Str2  string `query:"str2"`
}
