package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumLoss' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER_ARRAY price as parameter.
 */

func minimumLoss(price []int64) int32 {
	// the length should be based on input's first line, but this function is
	// prepared like that in the original issue, so just calculate it
	pricesLen := len(price)
	// also, the constaints say that 1 < price[i] < 10^16, so it's safe to convert it to int
	intPrices := []int{}
	for _, val := range price {
		intPrices = append(intPrices, int(val))
	}
	// then just calculate stuff
	minLoss := int(10e16)
	for i := 0; i < (pricesLen - 1); i++ {
		for j := 1; j < (pricesLen - i); j++ {
			currentPrice := intPrices[i]
			newPrice := intPrices[i+j]
			if currentPrice < newPrice {
				continue
			}
			loss := currentPrice - newPrice
			if loss < minLoss {
				minLoss = loss
			}
		}
	}
	return int32(minLoss)
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
