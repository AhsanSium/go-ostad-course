package main
import "fmt"

func main () {
	// string ---> sequence of bytes
	// byte ----> uint8  -----> can represent 256 value or symbol or letter

	var textValue string = "Hello World"
	
	banglaText := "বাংলা টাইপ করুন"

	fmt.Println(textValue[0], string(textValue[0]))
	fmt.Println(banglaText[0], string(banglaText[0]))

	fmt.Printf("Type of Bangla text [0]: %T \n", banglaText[0])

	banglaTextRune := []rune("বাংলা টাইপ করুন")

	fmt.Println(banglaTextRune)
	fmt.Println(banglaTextRune[0], string(banglaTextRune[0]))
	fmt.Printf("Type of Rune Bangla text [0]: %T \n", banglaTextRune[0])
}