package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)
func get_key_matrix(key string, keymatrix [3][3]int) [3][3]int {
  k := 0
  for i:=0 ; i < 3; i++ {
    for j:=0; j < 3; j++ {
      keymatrix[i][j] = int(key[k]) % 65;
      k++;
    }
  }
  return keymatrix
}
func encrypt(cipher_matrix [3][1]int, keymatrix [3][3]int, messageVector [3][1]int) [3][1]int {
  for i:= 0; i < 3; i++ {
    for j:=0; j < 1; j++ {
      cipher_matrix[i][j] = 0;
      for x:=0; x < 3; x++ {
        cipher_matrix[i][j] += keymatrix[i][x] * messageVector[x][j]
      }
      cipher_matrix[i][j] = cipher_matrix[i][j] % 26
    }
  }
  return cipher_matrix
}

func encrypt_handler(inp string, key string) string {
  var keymatrix [3][3]int
  keymatrix = get_key_matrix(key, keymatrix)
  
  var messageVector [3][1]int
  for i:=0; i < 3; i++ {
    messageVector[i][0] = int(inp[i]) % 65
  }

  var cipher_matrix [3][1]int
  cipher_matrix = encrypt(cipher_matrix, keymatrix, messageVector)

  output := ""
  for i := 0; i < 3; i++ {
    output += string(cipher_matrix[i][0] + 65)
  }
  return output
}

func main() {
  var key string
  fmt.Println("Enter message to be encrypted:")
  
  in := bufio.NewReader(os.Stdin)
  input, _ := in.ReadString('\n')
  input = strings.ToUpper(input)

  fmt.Println("Enter key");
  //fmt.Scanln(&key);
  key = "helloworld"

  encrypted_inp := encrypt_handler(input, key)
  fmt.Println("Encrypted text: \n")
  fmt.Println(encrypted_inp)
}
