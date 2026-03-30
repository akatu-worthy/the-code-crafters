package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
 var base int

 fmt.Println("\033[36mNumber Base Converter\033[0m")
 fmt.Println("\033[36m---------------------\033[0m")
 fmt.Println("\033[36mChoose input base:\033[0m")
 fmt.Println("\033[36m1 = Binary\033[0m")
 fmt.Println("\033[36m2 = Decimal\033[0m")
 fmt.Println("\033[36m3 = Hexadecimal\033[0m")
 fmt.Println("\033[36m4 = Quit\033[0m")

 fmt.Print("\033[36mEnter choice:\033[0m ")
 fmt.Scan(&base)

 fmt.Print("\033[36mEnter the number:\033[0m ")
 fmt.Scan(&input)


 if input == "" {
  fmt.Println("\033[31mError: Input cannot be empty\033[0m")
  continue
 }


 switch base {
 case 1:
  
	for _, ch := range input {
   if ch != '0' && ch != '1' {
    fmt.Println("\033[31mError: Invalid binary number\033[0m")
    continue
   }
  }
 case 2:
  if input[0] == '-' && len(input) > 1 {
   input = input[1:]
  }
  
  for _, ch := range input {
   if ch < '0' || ch > '9' {
    fmt.Println("\033[31mError: Invalid decimal number\033[0m")
    continue
   }
  }
 case 3:
  
	for _, ch := range input {
   if !(ch >= '0' && ch <= '9' ||
    ch >= 'a' && ch <= 'f' ||
    ch >= 'A' && ch <= 'F') {
    fmt.Println("\033[31mError: Invalid hexadecimal number\033[0m")
    continue
   }
  }
case 4:
	if input == "4" {
	fmt.Println("\033[32mGoodbye...\033[0m")
	return
 }

 default:
  fmt.Println("\033[31mInvalid base choice\033[0m")
  continue
 }

 var decimal int64
 var err error

 switch base {
 case 1:
  decimal, err = strconv.ParseInt(input, 2, 64)
 
case 2:
  decimal, err = strconv.ParseInt(input, 10, 64)

case 3:
  decimal, err = strconv.ParseInt(input, 16, 64)
 }

 if err != nil {
  fmt.Println("\033[31mError:\033[0m", err)
  continue
 }

 fmt.Println("Conversions:")
 fmt.Println("\033[32mBinary:\033[0m ", strconv.FormatInt(decimal, 2))
 fmt.Println("\033[32mDecimal:\033[0m ", strconv.FormatInt(decimal, 10))
 fmt.Println("\033[32mHexadecimal:\033[0m ", strconv.FormatInt(decimal, 16))

 continue
	}
}
