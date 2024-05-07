import (
	"fmt"
)

func declinations(c int) string {
	twoInt := c % 100 // последние две цифры числа
	oneInt := c % 10  // последняя цифра числа

	var result string
	switch {
	case oneInt >= 11 && twoInt <= 19:
		result = "компьютеров"
	case oneInt == 1:
		result = "компьютер"
	case oneInt >= 2 && oneInt <= 4:
		result = "компьютера"
	default:
		result = "компьютеров"
	}
	return fmt.Sprintf("%d %s", c, result)
}

func main() {
	fmt.Println(declinations(5))
}
