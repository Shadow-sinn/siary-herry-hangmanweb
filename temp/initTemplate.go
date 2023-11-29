package initTemplate

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./temp/*.html")
	if errTemp != nil {
		fmt.Printf("Erreur de Templates : %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temp
}
