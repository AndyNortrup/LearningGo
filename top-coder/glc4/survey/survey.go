package survey

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	//
)

func CountBigShoppers(N int, s []int) int {
	var bigShoppers = make([]int, N)
	var result int

	var sum int
	for _, value := range s {
		sum += value
	}

	var shopperIndex int
	for x := 0; x < len(s); x++ {

		for y := 0; y < s[x]; y++ {
			bigShoppers[shopperIndex]++
			if bigShoppers[shopperIndex] == len(s) {
				result++
			}
			shopperIndex++
			shopperIndex = shopperIndex % N
		}

	}

	return result
}

func SurveyServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to open TCP Listener")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Failed to Accept incoming TCP connection")
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input, _ := reader.ReadString('\n')
	survey := strings.Split(input, "#")
	shoppers, _ := strconv.Atoi(survey[0])
	var purchases = make([]int, 0)
	for _, value := range strings.Split(survey[1], ",") {
		numPurchased, _ := strconv.Atoi(value)
		purchases = append(purchases, numPurchased)
	}

	writer.Write(CountBigShoppers(shoppers, purchases))
}
