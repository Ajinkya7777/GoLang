package main
import "fmt"

func calcAvg(arr [6] int,size int)int{
	var sum int
	for _,value := range arr{
		sum += value
	}
	return sum / size
}

func main() {
	scores := [6]int{50,50,50,50,50,50}

	avg := calcAvg(scores,len(scores))
	fmt.Println("Avg:",avg)
}