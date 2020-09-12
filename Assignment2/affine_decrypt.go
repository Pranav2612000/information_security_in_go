package main

import (
  "fmt"
  "bufio"
  "os"
)

func get_inverse(a int) int {
  var a_inv int
  var flag int
  flag = 0
  for i:= 1; i < 26; i++ {
    flag = (a*i) % 26
    if(flag == 1) {
      a_inv = i
    }
  }
  return a_inv
}
func decrypt(cipher string, a int, b int) string {
  var message string;
  a_inv := get_inverse(a)
  for _, letter := range cipher {
    if(letter == ' ' || letter == '\n') {
      message = message + string(letter)
      continue
    }
    message = message + string((a_inv*(int(letter) + int('A') - b) % 26) + int('A')) 
  }
  return message
}

func main() {
  var a int;
  var b int;
  a = 2;
  fmt.Println("Enter message to be decrypted:")
  
  in := bufio.NewReader(os.Stdin)
  input, _ := in.ReadString('\n')

  fmt.Println("Enter a ");
  fmt.Scanln(&a);

  fmt.Println("Enter b")
  fmt.Scanln(&b);

  decrypted_inp := decrypt(input, a, b)

  fmt.Println("Decrypted text: \n")
  fmt.Println(decrypted_inp)
}
