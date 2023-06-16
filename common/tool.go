package common

import (
	"fmt"
)

func PrintScan(varName string, name ...any) {
	if len(name) == 0 {
		fmt.Printf("%s\n", varName)
	} else {
		switch value := name[0].(type) {
		case *string:
			fmt.Printf("%s: ", varName)
			fmt.Scanln(value)
		case *int:
			fmt.Printf("%s: ", varName)
			fmt.Scanln(value)
		case nil:
			fmt.Printf("\n%s ", varName)
			fmt.Scanln()
		}
	}
}

func PrintError(err error) {
	if err != nil {
		println(err.Error())
	}
}
