package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// declaring variables
	fmt.Println("Variables")
	var a = 42
	fmt.Println(a)

	name := "Sagun Shrestha"
	fmt.Println(name) // println automatically adds \n

	// multiple variables in same line
	age, dob := 25, "1996/02/24"
	fmt.Printf("age = %d, dob = %s\n", age, dob)

	// complex numbers
	fmt.Println("\n\nComplex numbers")
	c := 4 + 3i
	d := 4 - 3i
	fmt.Printf("type of c and d is %T and %T\n", c, d)
	fmt.Println("product of", c, "and", d, "is", real(c*d)) // %v for complex

	// arrays
	fmt.Println("\n\nArrays and slices")
	langs := [...]string{"PHP", "Python", "Go", "Java", "English", "Nepali"}
	fmt.Println(len(langs))

	for idx, lang := range langs {
		fmt.Println(idx+1, lang)
	}

	// slices
	dynamic_langs := langs[:2]
	static_langs := langs[2:4]
	natural_langs := langs[4:]
	fmt.Println("dynamic languages =", dynamic_langs, "\nstatic languages =", static_langs, "\nnatural languages =", natural_langs)

	// maps - same as python dictionary
	fmt.Println("\n\nMaps")
	suburbs := map[string]int{
		"Sydney":       2000,
		"Melbourne":    3000,
		"North Sydney": 2060,
		"Parramatta":   2150,
		"Brisbane":     4000,
	}
	for suburb, postcode := range suburbs {
		fmt.Printf("%s is in postcode %d\n", suburb, postcode)
	}

	// if statements
	fmt.Println("\n\nif statements")
	var num int
	fmt.Print("Enter a number:")
	_, err := fmt.Scan(&num) // use _ for unused variables as go compiler throws error if you don't use a variable
	if err != nil {
		panic(err)
	}

	if num > 80 {
		fmt.Println("greater than 80")
	} else if num > 50 {
		fmt.Println("greater than 50 and less than 81")
	} else if num > 25 {
		fmt.Println("greater than 25 and less than 51")
	} else {
		fmt.Println("less than 25")
	}

	// loops
	fmt.Println("\n\nLoops")
	for i := 1; i <= 5; i++ {
		fmt.Printf("square of %d is %d\n", i, i*i)
	}

	// using for loop like while loop
	j := 3
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		var key string
		fmt.Printf("This will run forever unless you press q: ")
		_, err := fmt.Scan(&key)
		if err != nil {
			return
		}

		if key == "q" {
			fmt.Println("breaking out of loop")
			break
		}
	}
}
