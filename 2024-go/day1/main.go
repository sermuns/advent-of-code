package main

import (
  "bufio"
  "fmt"
  "sort"
//  "strings"
//  "strconv"
  "os"
)

func readNums(fileName string)([]int, []int){
  file, err := os.Open(fileName)
  defer file.Close()

  if err != nil{
    fmt.Println(err)
    os.Exit(1)
  }

  leftNums := []int{}
  rightNums := []int{}

  reader := bufio.NewReader(file)

  for i := 0; true; i++{

    line, err := reader.ReadString('\n')
    if err != nil {
      break // was unable to read the line
    }

    var leftNum, rightNum int
    fmt.Sscanf(line, "%d   %d", &leftNum, &rightNum)
    leftNums = append(leftNums, leftNum)
    rightNums = append(rightNums, rightNum)
  }
  return leftNums, rightNums
}

func main(){
  leftNums, rightNums := readNums("input")
  sort.Ints(leftNums)
  fmt.Println(leftNums, rightNums)
}
