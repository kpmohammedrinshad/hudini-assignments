// 1) count the monkey
package kata

import (
	"regexp"
	"strconv"
)

func monkeyCount(n int) []int {
	// Your code here, happy coding!
	a := []int{}
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	return a
}

//2) return negative
package kata

func MakeNegative(x int) int {
  if x >= 0{
    return -x
  }
     return 0
  }
 

//3)find multiples of a number
package kata

func FindMultiples(integer, limit int) []int {
  // Your code here!
  list:=[]int{}
  for i:=integer;i<=limit;i+=integer{
    list=append(list,i)
  }
  return list
}

//4)count by x
package kata


func CountBy(x, n int) []int {
  list:=[]int{}
  for i:=x;i<=n*x;i+=x{
    if(i%x==0){
      list=append(list,i)
    }
  }
  return list
}

//5)power of 2
package kata
import "math"

func PowersOfTwo(n int) []uint64 {
  // your code goes here
  list:=[]uint64{}
  for i:=0;i<=n;i++{
    power:=math.Pow(2,float64(i))
    list=append(list,uint64(power))
  }
  return list
}

//6)total amount of points
package kata
import "strings"

func Points(games []string) int {
  // your code here!
  score:=0
  for i:=0;i<len(games);i++{
    result:=strings.Split(games[i],":")
    if(result[0]>result[1]){
      score=score+3
    }else if(result[0]==result[1]){
      score=score+1
    }else{
     score=score+0
    }
  }
  
  return score
}
//or
package kata

func Points(games []string) int {
  // your code here!
  score:=0
  for i:=0;i<len(games);i++{
//     n := games[i]
    if games[i][0] > games[i][2] {
      score=score+3
    }else if(games[i][0]==games[i][2]){
      score=score+1
    }else{
     score=score+0
    }
  }
  
  return score
}

//7) Beginner Series #3 Sum of Numbers
func GetSum(a,b int) int{
	if(a>b){
		a,b=b,a
	}
	sum:=0
	for i:=a;i<=b;i++{
		sum+=i
	}
	return sum
}

//8)shortest word
import "strings"
func FindShort(s string) int{
	short:=strings.Split(s," ")
	small:=len(short[0])
	for i:=1;i<len(short);i++{
		if(len(short[i])<small){
			small=len(short[i])
		}
	}
	return small
}

//9)sum of integers in string
import (
  "regexp"
  "strconv"
)
func SumOfIntegersInString(strng string) int{
  sum:=0
  for _,str:=range regexp.MustCompile(`\d+`).FindAllString(strng,-1){
  i,_:=strconv.Atoi(str)
  sum+=i
}
return sum
}


// 10)Least larger
func LeastLarger(a []int, i int) int {
	least:=10000
	index:=-1
	for j:=0;j<len(a);j++{
	  if(a[j]>a[i]){
		if a[j]<least{
		  least=a[j]
		  index=j
		}
	  }
	}
	return index
  }

//11)square(n) sum
import "math"

func SquareSum(numbers []int) int {
    // your code here
  sum:=0
  for i:=0;i<len(numbers);i++{
    power:=math.Pow(float64(numbers[i]),2)
    sum=sum+int(power)
  }
  return sum
}

//12)count odd numbers below n
func OddCount(n int) int{
	//your code here
   return n/2
  }
//13)odd-even string sort

func SortMyString(s string) string {
	even,odd := "",""
	for i:=0;i<len(s);i++{
	  if(i%2==0){
		even=even+string(s[i])
	  }else{
		odd=odd+string(s[i])
	  }
	}
	return even +" "+odd
  }
  
