package actioninfo

import "fmt"

type DataParser interface {
	ActionInfo() (string, error)
	Parse(datastring string) (err error)
}

func Info(dataset []string, dp DataParser) {
	for _, el := range dataset {
		err := dp.Parse(el)
		if err != nil {
			fmt.Println(err)
			continue
		}

		str, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(str)
	}
}
