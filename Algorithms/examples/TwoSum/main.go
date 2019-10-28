package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
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

// TwoSum Returns number of times three items sums equal 0
func TwoSum(arr []int) int {
	count := 0
	N := len(arr)
	sort.Ints(arr)

	for j := 0; j < N; j++ {
		valIdx := sort.Search(N, func(i int) bool {
			return arr[i] >= -arr[j]
		})

		if valIdx < N && arr[valIdx] == -arr[j] {
			count++
		}
	}
	return count / 2
}

// TwoSum Returns number of times three items sums equal 0
func ThreeSumFast(arr []int) int {
	count := 0
	N := len(arr)
	sort.Ints(arr)

	for j := 0; j < N; j++ {
		for k := j + 1; k < N; k++ {
			valIdx := sort.Search(N, func(i int) bool {
				return arr[i] >= (-arr[j] - arr[k])
			})

			if valIdx < N && arr[valIdx] == -arr[j] {
				count++
			}
		}
	}

	return count / 2
}

func main() {
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/1Kints.txt")
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/2Kints.txt")
	// numArr := readLines("https://algs4.cs.princeton.edu/14analysis/4Kints.txt")
	numArr := readLines("https://algs4.cs.princeton.edu/14analysis/8Kints.txt")

	// numArr := []int{-3, 4, 9, -4, 6, 3}
	// Timed tests
	start := time.Now()

	fmt.Println("\nNumber of doubles equaling 0: ", ThreeSumFast(numArr))

	elapsed := time.Since(start)

	log.Printf("\n\nMain.go Executed in: %s\n\n", elapsed)

}
