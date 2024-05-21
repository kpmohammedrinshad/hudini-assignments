package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text:")
	text,_:= reader.ReadString('\n')

	freqWord:=map[string]int{}

	// text = strings.TrimSpace(text)
	// words:=strings.Split(text," ")
	words:=strings.Fields(text)

	for _,word:=range words{
		freqWord[word]++
	}
	for k,v:=range freqWord{
		fmt.Printf("'%s' : %d\n",k,v)
	}

}