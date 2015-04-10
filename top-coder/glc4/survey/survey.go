package survey

import (
	"bytes"
	"log"
	"net"
	"strconv"
	"strings"
	//
)

func SurveyServer() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Printf("Failed to open TCP Listener\n %v\n", err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Server: Failed to Accept incoming TCP connection\n")
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	buffer := make([]byte, 1000)
	_, err := conn.Read(buffer)

	buffer = bytes.Trim(buffer, "\x00")

	if err != nil {
		log.Fatalf("Server: Error reading from buffer: %v", err)
	}

	defer conn.Close()

	result := CountBigShoppers(decodeRequest(buffer))

	_, err = conn.Write([]byte(strconv.Itoa(result) + "\n"))
	if err != nil {
		log.Fatal("Server: Failed to write response\n Error: %v", err)
	}
}

func decodeRequest(request []byte) (int, []int) {

	strRequest := string(request[:len(request)])
	strRequest = strings.TrimSpace(strRequest)

	requestPieces := strings.Split(strRequest, "#")
	shoppers, err := strconv.Atoi(requestPieces[0])

	strPurchases := strings.Split(requestPieces[1], ",")

	purchases := make([]int, len(strPurchases))

	for index, value := range strPurchases {
		purchases[index], err = strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Failed to convert value to int: %v\n%v", value, err)
		}
	}

	return shoppers, purchases
}

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
