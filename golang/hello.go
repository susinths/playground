package main
import "fmt"

func main() {
  var s string = "Hello, World\n"
  fmt.Printf(s)

for index, char := range(s) {
fmt.Printf("%d-%c ", index, char)
}

}

