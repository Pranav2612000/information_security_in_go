package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func removeSpaces(str string) string {
      return strings.Join(strings.Fields(str), "")
}
func preprocess(str string) string {
  if(len(str) % 2 != 0) {
    str = str + string('z')
  }
  return str
}
func generateKeyTable(key string, keyT [5][5]string) [5][5]string {
  var dicty [26]int
  for i:=0; i < 26; i++ {
    dicty[i] = 0
  }
  for _, letter := range key {
    if(letter != 'j') {
      dicty[int(letter) - 97] = 2
    }
  }
  dicty[int('j') - 97] = 1

  j := 0
  k := 0
  for i:=0; i < len(key); i++ {
    if(dicty[key[i] - 97] == 2) {
      dicty[key[i] - 97] -= 1;
      keyT[j][k] = string(key[i])
      k++;
      if(k==5) {
        j++;
        k = 0;
      }
    }
  }

  for i:=0; i < 26; i++ {
    if(dicty[i] == 0) {
      keyT[j][k] = string(i + 97)
      k++;
      if(k==5) {
        j++;
        k = 0;
      }
    }
  }
  return keyT
}

func search(keyT [5][5]string, a string, b string, arr [4]int) [4]int {
  if(a == string('j')) {
    a = string('i')
  }
  if(b == string('j')) {
    b = string('i')
  }
  for i:=0; i < 5; i++ {
    for j:=0; j < 5; j++ {
      if(keyT[i][j] == a) {
        arr[0] = i;
        arr[1] = j;
      }
      if(keyT[i][j] == b) {
        arr[2] = i;
        arr[3] = j;
      }
    }
  }
  return arr
}

func decrypt(input string, keyT [5][5]string) string {
  var arr [4]int
  output := "" 
  for i:=0; i < len(input); i+=2 {
    arr = search(keyT, string(input[i]), string(input[i+1]), arr) 
    
    if(arr[0] == arr[2]) {
      output += keyT[arr[0]][((arr[1] - 1) + 5) %5]
      output += keyT[arr[0]][((arr[3] - 1) + 5) %5]
    } else if (arr[1] == arr[3]) {
      output += keyT[((arr[0] - 1) + 5) % 5][arr[1]]
      output += keyT[((arr[2] - 1) + 5) % 5][arr[1]]
    } else {
      output += keyT[arr[0]][arr[3]]
      output += keyT[arr[2]][arr[1]]
    }
  }
  return output
}

func decrypt_handler(input string, key string) string {
  var keyT [5][5]string
  key = strings.ToLower(key)
  input = strings.ToLower(input)
  input = removeSpaces(input)
  //preprocessed_inp := preprocess(input)
  keyT = generateKeyTable(key, keyT)
  /*
  for i:=0; i < 5; i++ {
    for j:= 0; j < 5; j++ {
      fmt.Printf("%s", keyT[i][j])
    }
    fmt.Println("")
  }
  */
  return decrypt(input, keyT)
  //return preprocessed_inp 
}

func main() {
  var key string
  fmt.Println("Enter message to be decrypted")
  in := bufio.NewReader(os.Stdin)
  input, _ := in.ReadString('\n')
  input = strings.ToLower(input) 
  input = strings.TrimSpace(input)
  fmt.Println("Enter key")
  fmt.Scanln(&key)
  key = strings.ToLower(key) 
  key = strings.TrimSpace(key)
  decrypted_inp := decrypt_handler(input, key)
  fmt.Println("Decrypted Text: ")
  fmt.Println(decrypted_inp)
}
