package model

type FuzzBuzzRequest struct {
	Int1  int    `query:"int1" validate:"required,gt=0"`
	Int2  int    `query:"int2" validate:"required,gt=0"`
	Limit int    `query:"limit" validate:"required,gt=0,lte=100"`
	Str1  string `query:"str1" validate:"required"`
	Str2  string `query:"str2" validate:"required"`
}
