//
//  text.go.go
//  utils
//
//  Created by d-exclaimation on 1:38 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package utils

import "strings"

// CamelCase a string
func CamelCase(strs []string) string {
	res := make([]string, len(strs))
	for i, word := range strs {
		res[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(res, "")
}

// SnakeCase a string
func SnakeCase(strs []string) string {
	res := make([]string, len(strs))
	for i, word := range strs {
		res[i] = strings.ToLower(word)
	}
	return strings.Join(res, "_")
}
