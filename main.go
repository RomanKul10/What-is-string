package main

import (
	"fmt"
	"strings"
)

func main() {
	/* fmt.Println()
	name := "Hello world!"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()

	name = "Привіт світ"
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()
	fmt.Println("Three ways to concatenate strings") // три способи об'єднати рядки

	h := "Hello, "
	w := "world."

	//using + not very efficient
	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	//using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	//using stringbulder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[10:13])

	//indexing

	coursName := "Learn Go for Beginners Crash Course"

	fmt.Println(coursName[0])         // number
	fmt.Println(string(coursName[0])) //leter
	fmt.Println(string(coursName[6]))

	for i := 0; i <= 21; i++ {
		fmt.Print(string(coursName[i]))
	}
	fmt.Println()
	for i := 13; i <= 21; i++ {
		fmt.Print(string(coursName[i]))
	}

	fmt.Println()

	fmt.Println("Length of courseName is", len(coursName))

	var mySlise []string
	mySlise = append(mySlise, "one")
	mySlise = append(mySlise, "two")
	mySlise = append(mySlise, "three")

	fmt.Println("mySlise has", len(mySlise), "elements")
	fmt.Println("the last element mySlise is", mySlise[len(mySlise)-1]) */

	
	//strings package
	courses := []string{
		"Learn Go for Beginners Crash Course"
		"Learn Java for Beginners Crash Course"
		"Learn Python for Beginners Crash Course"
		"Learn C for Beginners Crash Course"
	}

	for _, x := rage course{
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in",x, "and index is", strings.Index(x,"Go"))
		}
	}

	newString := "Go is a great programing language. Go for it"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Python")) 
	fmt.Println(strings.LastIndex(newString, "Go"))
}
