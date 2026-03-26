## Go Maps

 Maps are used to store data values in key value pairs.
 A map is an unordered and changeable collection that does not allow duplicates.

 Example:
 package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

Result:
a   map[brand:Ford model:Mustang year:1964]
b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]