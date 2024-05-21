// 1. Calculate Average
// Objective: Write a function that takes an array of numbers and returns the average. Use loops and basic arithmetic.
// Function signature:
package main
import "fmt"
func calculateAverage(numbers []int) float64 {
    // Write your code here to calculate and return the average of the array elements.
	sum:=0
	for i:=0;i<len(numbers);i++{
		sum+=numbers[i]
	}
	var avg float64=float64(sum)/float64(len(numbers))
	return avg

}
func main(){
	fmt.Printf("the average is %v",calculateAverage([]int{10,20,30,40,50}))
}
 
// Example usage:
console.log(calculateAverage([10, 20, 30, 40, 50])); // Expected output: 30
 
// 2. Check Age
// Write a function that takes an age and prints whether the person is a minor, a young adult, or an adult.
// Use conditional statements.
// Function signature:
func checkAge(age int) string {
    // Write your code here to determine and print if the age corresponds to a minor, a young adult, or an adult.
	if(age>150){
		return "out of age"
	}
	if(age<0){
		return "no negative age"
	}
	if(age<13){
		return "minor"
	}else if(age<21){
		return "young adult"
	}else {
		return "adult"
	}

}
func main(){
	fmt.Printf("the age is %v",checkAge(23))
}
 
// Example usage:
checkAge(25); // Possible output: young adult
 
// 3. Reverse String
// Objective: Create a function that reverses a string. This will demonstrate basic string manipulation and for loops.
// Function signature:
func reverseString(str string) string {
    // Write your code here to reverse and return the string.
	var reverse=""
	for i:=len(str)-1;i>=0;i--{
		reverse+=string(str[i])
	}
	return reverse
}
func main(){
	fmt.Printf("the reverseString is %v",reverseString("hello"))
}
 
// Example usage:
console.log(reverseString("hello")); // Expected output: "olleh"
 
// 4. Find Largest Number
// Objective: Write a function that takes an array of numbers and returns the largest number.
// Use loops and conditional statements to solve the problem.
// Function signature:
func findLargestNumber(number []int) int {
    // Write your code here to find and return the largest number in the array.
	higherNumber:=number[0]
	for i:=1;i<len(number);i++{
		if(number[i]>higherNumber){
			higherNumber=number[i]
		}
	}
	return higherNumber
}
func main(){
	fmt.Printf("the largest number is %v",findLargestNumber([]int{20,45,33,67,12,7}))
}
 
// Example usage:
console.log(findLargestNumber([10, 20, 30, 40, 50])); // Expected output: 50
 
// 5. Simple Counter
// Create an object that acts as a counter with methods to add, subtract, and reset the count.
// Demonstrate object methods and the use of this.
// Object definition:
package main
import "fmt"

type Counter struct{
    count int
}

func (c Counter) add() Counter{
    c.count++
    return c
}
func (c Counter) substract() Counter{
    c.count--
    return c
}
func (c Counter) reset() Counter{
    c.count=0
    return c
}
func (c Counter) display(){
    fmt.Println(c.count)
}

func main() {
    a:=Counter{count:0}
    a=a.add()
    a.display()
    a=a.substract()
    a.display()
}
 
// Example usage:
counter.add();
counter.add();
counter.display();  // Output: 2
counter.subtract();
counter.display();  // Output: 1
counter.reset();
counter.display();  // Output: 0
 