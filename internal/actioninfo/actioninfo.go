package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {

		if err := dp.Parse(data); err != nil {
			log.Printf("data parsing error: %v", err)
			continue // Переходим к следующей итерации
		}
		if info, err := dp.ActionInfo(); err != nil {
			log.Printf("error receiving activity information: %v", err)
		} else {
			fmt.Println(info)
		}
	}
}
