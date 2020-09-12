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
func decrypt(cipher_matrix [3][1]int, inverse_keymatrix [3][3]float32, messageVector [3][1]int) [3][1]int {
  for i:= 0; i < 3; i++ {
    for j:=0; j < 1; j++ {
      messageVector[i][j] = 0;
      for x:=0; x < 3; x++ {
        messageVector[i][j] += int(inverse_keymatrix[i][x] * float32(cipher_matrix[x][j]))
      }
      messageVector[i][j] = messageVector[i][j] % 26
    }
  }
  return messageVector 
}
func inverse_matrix(a [3][3]int) [3][3]float32 {
  var determinant float32
  var inverse [3][3]float32
  determinant = 0 
  for i:=0; i < 3; i++ {
    determinant = determinant + float32((a[0][i]*(a[1][(i+1)%3]*a[2][(i+2)%3] - a[1][(i+2)%3]*a[2][(i+1)%3])))
  }

  for i:= 0; i < 3; i++ {
    for j:= 0; j < 3; j++ {
      inverse[i][j] = float32((a[(i+1)%3][(j+1)%3] * a[(i+2)%3][(j+2)%3]) - (a[(i+1)%3][(j+2)%3]*a[(i+2)%3][(j+1)%3]))
    }
  }
  return inverse 
}

func decrypt_handler(inp string, key string) string {
  var keymatrix [3][3]int
  var inverse [3][3]float32
  keymatrix = get_key_matrix(key, keymatrix)
  inverse = inverse_matrix(keymatrix) 
  
  var cipher_text [3][1]int
  for i:=0; i < 3; i++ {
    cipher_text[i][0] = int(inp[i]) % 65
  }

  var message_text [3][1]int
  message_text = decrypt(cipher_text, inverse, message_text)

  output := ""
  for i := 0; i < 3; i++ {
    output += string(message_text[i][0] + 65)
  }
  return output
}

func main() {
  var key string
  fmt.Println("Enter message to be decrypted:")
  
  in := bufio.NewReader(os.Stdin)
  input, _ := in.ReadString('\n')
  input = strings.ToUpper(input)

  fmt.Println("Enter key");
  //fmt.Scanln(&key);
  key = "helloworld"

  decrypted_inp := decrypt_handler(input, key)
  fmt.Println("Decrypted text: \n")
  fmt.Println(decrypted_inp)
}
