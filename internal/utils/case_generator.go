package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var TitleCaser cases.Caser

func InitCaser() {
	TitleCaser = cases.Title(language.English)
}
