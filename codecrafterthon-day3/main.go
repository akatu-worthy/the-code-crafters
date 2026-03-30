package main

import (
 "fmt"
 "strings"
 "unicode"
)

func main() {
 fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
 fmt.Println("──────────────────────────────────────")

 for {
  var cmd string
  var text string

  fmt.Print("> ")

  fmt.Scan(&cmd)
  cmd = strings.ToLower(cmd)

  if cmd == "exit" {
   fmt.Println("Shutting down String Transformer. Goodbye.")
   return
  }

  fmt.Scanln(&text)

  text = strings.TrimSpace(text)

  if text == "" {
   fmt.Println("No text provided.")
   continue
  }

  var result string
  valid := true

  switch cmd {
  case "upper":
   result = strings.ToUpper(text)

  case "lower":
   result = strings.ToLower(text)

  case "cap":
   result = capitalize(text)

  case "title":
   result = titleCase(text)

  case "snake":
   result = snakeCase(text)

  case "reverse":
   result = reverseWords(text)

  default:
   valid = false
   fmt.Println("Unknown command:")
   fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, exit")
  }

  if valid {
   fmt.Println("→", result)
  }
 }
}


func toUpper(s string) string {
 return strings.ToUpper(s)
}

func toLower(s string) string {
 return strings.ToLower(s)
}

func capitalize(s string) string {
 words := strings.Fields(s)

 for i, w := range words {
  if len(w) > 0 {
   words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
  }
 }

 return strings.Join(words, " ")
}

func titleCase(s string) string {
 small := map[string]bool{
  "a": true, "an": true, "the": true, "and": true,
  "but": true, "or": true, "for": true, "nor": true,
  "on": true, "at": true, "to": true, "by": true,
  "in": true, "of": true, "up": true, "as": true,
  "is": true, "it": true,
 }

 words := strings.Fields(s)

 for i, w := range words {
  lw := strings.ToLower(w)

  if i == 0 || !small[lw] {
   words[i] = strings.ToUpper(string(lw[0])) + lw[1:]
  } else {
   words[i] = lw
  }
 }

 return strings.Join(words, " ")
}

func snakeCase(s string) string {
 var result []rune

 for _, r := range strings.ToLower(s) {
  if unicode.IsLetter(r) || unicode.IsDigit(r) {
   result = append(result, r)
  } else if unicode.IsSpace(r) {
   result = append(result, '_')
  }
 }

 out := strings.ReplaceAll(string(result), "__", "_")
 return strings.Trim(out, "_")
}

func reverseWords(s string) string {
 words := strings.Fields(s)

 for i, w := range words {
  runes := []rune(w)
  for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
   runes[l], runes[r] = runes[r], runes[l]
  }
  words[i] = string(runes)
 }

 return strings.Join(words, " ")
}
