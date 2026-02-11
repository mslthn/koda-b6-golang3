package lib

import (
	"fmt"
	"strings"
)

func SearchPerson(nameList []string, users string) {
	
	matchUser := []string{}

	for i := range(len(nameList)) {
		if (strings.Contains(strings.ToLower(nameList[i]), strings.ToLower(users) )) {
			matchUser = append(matchUser, nameList[i])
		}
	}
	
	fmt.Println(matchUser)
}