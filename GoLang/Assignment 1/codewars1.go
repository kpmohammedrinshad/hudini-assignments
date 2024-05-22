// https://www.codewars.com/kata/586f6741c66d18c22800010a
//1) sum of a sequence
package kata


func SequenceSum(start, end, step int) int {
  sum:=0
  for i:=start;i<=end;i+=step{
    sum+=i
  }
  return sum
}


// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go
//2) get the middle character
package kata

func GetMiddle(s string) string {
  if len(s)==0{
    return ""
  }
  middleIndex:=len(s)/2
  if len(s)%2==0{
    return string(s[middleIndex-1]) + string(s [middleIndex])
  }
  return string(s[middleIndex])
  //Code goes here!
}

//https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
//3)Reversed string
package kata

func Solution(word string) string {
  rev:=""
  for i:=len(word)-1;i>=0;i--{
    rev+=string(word[i])
  }
  return rev
}

//https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/train/go
//4)even or odd
package kata

func EvenOrOdd(number int) string {
  if number%2==0{
    return "Even"
  }
  return "Odd" // your code here
}

//https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
//5)Vowel count
package kata

func GetCount(str string) (count int) {
  // Enter solution here
  for _,char:=range str{
    switch char{
      case 'a','e','i','o','u':
      count++
    }
  }
  return count
}

//https://codewars.com/kata/50654ddff44f800200000004/train/go
//6)multiply
package multiply

func Multiply(a, b int) int {
  return a * b
}

//
//7)