package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func readLines(url string) []int {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	s := strings.Split(string(data), "\n")
	numArr := []int{}

	for i := range s {
		stringNum := strings.TrimSpace(s[i])
		num, err := strconv.Atoi(stringNum)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n ", err)
		}
		numArr = append(numArr, num)
	}

	return numArr
}

// ThreeSum Returns number of times three items sums equal 0
func ThreeSum(arr []int) int {
	count := 0
	N := len(arr)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/1Kints.txt")
	numArr := readLines("https://algs4.cs.princeton.edu/14analysis/2Kints.txt")
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/4Kints.txt")
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/8Kints.txt")

	// Timed tests
	start := time.Now()

	fmt.Println("\nNumber of triples equaling 0: ", ThreeSum(numArr))

	elapsed := time.Since(start)

	log.Printf("\n\nMain.go Executed in: %s\n\n", elapsed)

}
