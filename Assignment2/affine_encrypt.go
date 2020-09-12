package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func gcd(a int, b int) int {
  var gcd int;
  for i:= 1; i <= b && i <= a; i++ {
    if(a % i == 0 && b % i == 0) {
      gcd = i
    }
  }
  return gcd
}

func encrypt(inp string, a int, b int) string {
  var cipher string
  cipher = ""
  for _, letter := range inp {
    if(letter == ' ' || letter == '\n') {
      cipher = cipher + string(letter)
      continue
    }
    cipher = cipher + string((((a * (int(letter) - int('A')) + b) % 26) + int('A')))
  }
  return cipher 
}

func main() {
  var a int;
  var b int;
  a = 2;
  fmt.Println("Enter message to be encrypted:")
  
  in := bufio.NewReader(os.Stdin)
  input, _ := in.ReadString('\n')
  input = strings.ToUpper(input)

  for gcd(a, 26) != 1 {
    fmt.Println("Enter a (coprime with 26)");
    fmt.Scanln(&a);
  }
  fmt.Println("Enter b")
  fmt.Scanln(&b);
  encrypted_inp := encrypt(input, a, b)
  fmt.Println("Encrypted text: \n")
  fmt.Println(encrypted_inp)
}
