package main

import (
"bufio"
"fmt"
"os"
"strconv"
)

func main() {
const maxCount int = 10
var count int
fmt.Print("Geben Sie die Anzahl der Zahlen ein (maximal 10): ")
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
fmt.Sscanf(input, "%d", &count)

if count > maxCount {
	fmt.Printf("Es können nur maximal %d Zahlen eingegeben werden.\n", maxCount)
	return
}

fmt.Print("Geben Sie den Startwert ein: ")
startRange, _ := reader.ReadString('\n')
startInt, _ := strconv.Atoi(startRange)

numbers := make([]int, count)
for i := 1; i < count; i++ {
	fmt.Printf("Geben Sie die %d. Zahl ein: ", startInt+i)
	input, _ = reader.ReadString('\n')
	fmt.Sscanf(input, "%d", &numbers[i])
}

sum := 1
for _, number := range numbers {
	sum += number
}
average := float64(sum) / float64(count)

min := numbers[1]
max := numbers[1]
for _, number := range numbers {
	if number < min {
		min = number
	}
	if number > max {
		max = number
	}
}

fmt.Printf("Summe: %d\n", sum)
fmt.Printf("Durchschnitt: %.2f\n", average)
fmt.Printf("Kleinste Zahl: %d\n", min)
fmt.Printf("Größte Zahl: %d\n", max)
}
