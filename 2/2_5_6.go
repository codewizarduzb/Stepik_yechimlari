package main

import (
 "fmt"
 "unicode"
 "unicode/utf8"
)

func main_2_5_62() {

 var text string
 fmt.Scan(&text)
 isValid := true

 code := []rune(text)

 if utf8.RuneCountInString(string(code)) >= 5 {
  for i := range code {
   if !unicode.IsDigit(code[i]) && !unicode.Is(unicode.Latin, code[i]) {
    isValid = false
   }
  }
 } else {
	isValid = false
 }
 if isValid {
  fmt.Println("Ok")
 } else {
  fmt.Println("Wrong password")
 }
}