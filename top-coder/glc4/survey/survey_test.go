package survey

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestBigShoppers(t *testing.T) {

	shoppers := []int{5, 100, 10, 7, 5}
	survey := [][]int{{3, 3}, {97}, {9, 9, 9, 9, 9}, {1, 2, 3}, {3, 3, 3}}
	expected := []int{1, 97, 5, 0, 0}

	var fail bool = false
	for index, _ := range shoppers {
		result := CountBigShoppers(shoppers[index], survey[index])
		if expected[index] != result {

			t.Logf("Failed test %v \t Expected: %v \t Actual: %v \n",
				index,
				expected[index],
				result)
			fail = true
		} else {
			t.Logf("Passed test %v", index)
		}

		if fail {
			t.Fail()
		}
	}
}

func TestNetBigShoppers(t *testing.T) {

	shoppers := []int{5, 100, 10, 7, 5}
	survey := [][]int{{3, 3}, {97}, {9, 9, 9, 9, 9}, {1, 2, 3}, {3, 3, 3}}
	//expected := []int{1, 97, 5, 0, 0}

	go SurveyServer()

	time.Sleep(1000)

	for x := 0; x < len(shoppers); x++ {
		conn, err := net.Dial("tcp", ":8080")
		writer := bufio.NewWriter(conn)

		if err != nil {
			t.Fatalf("Failed to connect to server. \t%v", err)
		}

		request := buildRequest(shoppers[x], survey[x])

		fmt.Printf("Client: Sending Request: %v\n", request)
		_, err = writer.WriteString(request)
		writer.Flush()
		if err != nil {
			t.Fatal("Client: Failed to write request to network")
		}

		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			t.Fatal("Client: Failed to read response from network")
		}

		fmt.Printf("Recieved Result: %v\n", buf)

		/*
			result, _ := strconv.Atoi(line)
			if expected[x] != result {
				t.Fatalf("Incorrect response for test %v. \t Expected: %v\t Actual:%v\n",
					x, expected[x], line)
			}
		*/
	}
}

func buildRequest(shoppers int, survey []int) string {
	var request string
	request += strconv.Itoa(shoppers)
	request += "#"
	for index, value := range survey {
		request += strconv.Itoa(value)
		if index < len(survey)-1 {
			request += ","
		}
	}

	return request
}
