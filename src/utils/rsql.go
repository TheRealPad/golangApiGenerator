package utils

import (
	"httpServer/src/initialisation"
	"strings"
)

func Rsql(data string, lst []initialisation.Field) []initialisation.Field {
	splitStr := strings.Split(data, " and ")
	for _, s := range splitStr {
		if strings.Contains(s, "==") {
			data := strings.Split(s, "==")
			isValid := func(n initialisation.Field) bool {
				return n[data[0]].GetData().(string) == data[1]
			}
			lst = Filter(lst, isValid)
		}
	}
	return lst
}
