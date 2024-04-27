package nestedmap

import (
	"fmt"
	"strings"
)

func generateNestedMap(dict map[string]int) map[string]interface{} {
	result := make(map[string]interface{}, 0)

	for k, v := range dict {
		k_arr := strings.Split(k, "_")

		if len(k_arr) == 1 {
			result[k_arr[0]] = v
		} else {
			nested := result
			for i := 0; i < len(k_arr); i++ {
				if len(k_arr)-i == 1 {
					nested[k_arr[i]] = v
				} else {
					if _, ok := nested[k_arr[i]]; !ok {
						nested[k_arr[i]] = make(map[string]interface{}, 0)
					}
					nested = nested[k_arr[i]].(map[string]interface{})
				}
			}
		}
	}

	fmt.Printf("%+v\n", result)

	return result
}
