package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumLoss' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER_ARRAY price as parameter.
 */

// func minimumLoss(price []int64) int32 {
// 	// the length should be based on input's first line, but this function is
// 	// prepared like that in the original issue, so just calculate it
// 	pricesLen := len(price)
// 	// then just calculate stuff
// 	minLoss := int64(10e16)
// 	lastPrice := price[0]
// 	for i := 0; i < (pricesLen - 1); i++ {
// 		currentPrice := price[i]
// 		if currentPrice > lastPrice {
// 			continue
// 		}
// 		for j := 1; j < (pricesLen - i); j++ {
// 			newPrice := price[i+j]
// 			if currentPrice < newPrice {
// 				continue
// 			}
// 			loss := currentPrice - newPrice
// 			if loss < minLoss {
// 				minLoss = loss
// 			}
// 		}
// 		lastPrice = currentPrice
// 	}
// 	return int32(minLoss)
// }

func minimumLoss(price []int64) int32 {
	// the length should be based on input's first line, but this function is
	// prepared like that in the original issue, so just calculate it
	pricesLen := len(price)
	// // then just calculate stuff
	minLoss := int(10e16)

	intPrice := make([]int, len(price))
	sortedIntPrice := make([]int, len(price))

	for i, val := range price {
		intPrice[i] = int(val)
	}

	copy(sortedIntPrice, intPrice)
	sort.Ints(sortedIntPrice)

	currentLen := pricesLen
	for _, val := range intPrice {
		sortedIndex := index(sortedIntPrice, val)
		if sortedIndex > 0 {
			currentMinLoss := sortedIntPrice[sortedIndex] - sortedIntPrice[sortedIndex-1]
			if currentMinLoss < minLoss {
				minLoss = currentMinLoss
			}
		}
		sortedIntPrice = remove(sortedIntPrice, sortedIndex)
		currentLen--
	}
	return int32(minLoss)
}

func index(slice []int, item int) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	priceTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var price []int64

	for i := 0; i < int(n); i++ {
		priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
		checkError(err)
		price = append(price, priceItem)
	}

	result := minimumLoss(price)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
