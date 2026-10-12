package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10 
	var sngNum float64 = -3

	fmt.Println(intNum > 5)
	fmt.Println(intNum >= 5 && intOther < 11)
	fmt.Println(sngNum != -1 || intOther < 0)
	fmt.Println((!(intNum > 3) || intNum <= 5 ))
	fmt.Println((!(intOther >= intNum) ))
	fmt.Println((0 - sngNum > 0 ))
	fmt.Println((4 / 2 == intOther / intNum))
	fmt.Println((intOther % 2 == 0 ))
	fmt.Println((intOther + 2 * intNum != 30 || !(sngNum > 0) ))
	fmt.Println((intOther > 0 && intNum > 0 || sngNum > 0 ))
	fmt.Println((sngNum > 0 || (intNum >= 0 && -1 * intOther == -10) ))
	fmt.Println((intNum == 5))
	fmt.Println((intNum > 0 || (sngNum <= 0 && intOther == 13) ))
	fmt.Println((!(!(!(!(intNum > 0)))) ))
}