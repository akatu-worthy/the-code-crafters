package main

import (
 "fmt"
 "strconv"
)

func main() {
 var input string
 var base int

 fmt.Println("Number Base Converter")
 fmt.Println("---------------------")
 fmt.Println("Choose input base:")
 fmt.Println("1 = Binary")
 fmt.Println("2 = Decimal")
 fmt.Println("3 = Hexadecimal")

 fmt.Print("Enter choice: ")
 fmt.Scan(&base)

 fmt.Print("Enter the number: ")
 fmt.Scan(&input)


 if input == "" {
  fmt.Println("Error: Input cannot be empty")
  return
 }


 switch base {
 case 1:
  for _, ch := range input {
   if ch != '0' && ch != '1' {
    fmt.Println("Error: Invalid binary number")
    return
   }
  }
 case 2:
  if input[0] == '-' && len(input) > 1 {
   input = input[1:]
  }
  for _, ch := range input {
   if ch < '0' || ch > '9' {
    fmt.Println("Error: Invalid decimal number")
    return
   }
  }
 case 3:
  for _, ch := range input {
   if !(ch >= '0' && ch <= '9' ||
    ch >= 'a' && ch <= 'f' ||
    ch >= 'A' && ch <= 'F') {
    fmt.Println("Error: Invalid hexadecimal number")
    return
   }
  }
 default:
  fmt.Println("Invalid base choice")
  return
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
  fmt.Println("Error:", err)
  return
 }

 fmt.Println("\nConversions:")
 fmt.Println("Binary: ", strconv.FormatInt(decimal, 2))
 fmt.Println("Decimal: ", strconv.FormatInt(decimal, 10))
 fmt.Println("Hexadecimal: ", strconv.FormatInt(decimal, 16))
}
