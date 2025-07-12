package contact

import (
	"bufio"
	"fmt"
	"os"
)

func SayHello() {
	fmt.Println("Hello from contact.go!")
	var name string
	var age int
	var day, month, year int
	fmt.Print("Enter your name, age and date (YYYY-MM-DD): ")
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Scanf("%d-%d-%d", &year, &month, &day)
	fmt.Printf("Your name is %s, your age is %d and your date is %d-%d-%d\n", name, age, year, month, day)

	// Practice using Buffio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your favorite color: ")
	color, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Printf("Your favorite color is %s\n", color)
	fmt.Println(os.Args)
}