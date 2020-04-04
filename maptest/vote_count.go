package maptest

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// VoteCount 测试
func VoteCount() {

	lines, err := readFile()

	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}

	for _, line := range lines {
		counts[line]++
	}

	for k, v := range counts {
		fmt.Println("vote name is ", k, "count is", v)
	}
}

// VoteCountUseSort use sort
func VoteCountUseSort() {

	lines, err := readFile()

	if err != nil {
		log.Fatal(err)
	}

	countsMap := map[string]int{}
	var votes []string

	// fill map
	for _, line := range lines {
		countsMap[line]++
	}

	// fill slice
	for k, _ := range countsMap {
		votes = append(votes, k)
	}

	// 排序
	sort.Strings(votes)

	// 打印
	for _, v := range votes {
		fmt.Println("vote name is ", v, "count is", countsMap[v])
	}
}

//  readFile 从文件中读取数据
func readFile() ([]string, error) {

	var lines []string

	file, err := os.Open("votes.txt")

	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = scanner.Err()

	if err != nil {
		return lines, err
	}
	return lines, nil

}
