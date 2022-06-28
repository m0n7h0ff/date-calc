package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {

		/* e2 := Employee{
		FIO: "Перехода",
		start: Date(2022,6,7),
	}

	e3 := Employee{
		FIO: "Тюшняков",
		start: Date(2022,6,1),
	}

	e4 := Employee{
		FIO: "Дубинин",
		start: Date(2022,6,5),
	}

	fmt.Println(e1,e2,e3,e4) */

	e1 := Employee{
		FIO: "Монахов",
		start: Date(2022,6,3),
	}

	tNow := e1.start
	days := Subs(tNow)
	_, ost := math.Modf(float64(days) / 8.0)
	fmt.Printf("Раздница в днях: %v, остаток %v", days, ost)
	fmt.Println()

	printResult(ost)
}
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Subs(startDate time.Time) int {
	

	input := readStr("Введите число:")

	grade, err := time.Parse("2006-01-02", input)
	if err != nil {
		log.Fatal(err)
	}

	t1 := Date(startDate.Year(), int(startDate.Month()), startDate.Day())
	t2 := Date(grade.Year(), int(grade.Month()), grade.Day())

	days := t2.Sub(t1).Hours() / 24
	return int(days)
}

func readStr(s string) string{
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}

func printResult(ost float64){
	switch ost {
	case 0:
		fmt.Println("1-я смена, в день")
	case 0.125:
		fmt.Println("2-я смена, в день")
	case 0.75, 0.25:
		fmt.Println("1й выходной")
	case 0.375, 0.875:
		fmt.Println("2й выходной")
	case 0.5:
		fmt.Println("1-я смена, в ночь")
	case 0.625:
		fmt.Println("2-я смена, в ночь")
	default:
		fmt.Println("Fail/////")
	}
}

