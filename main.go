package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type PhoneInfo struct {
	PhoneNum string
	Province string
	City     string
	CardType string
	AreaZone string
	ZipCode  string
}

func main() {
	file, err := os.Open("Mobile0701.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recordMap := make(map[string]PhoneInfo)
	for _, value := range allRecords {
		var info PhoneInfo
		info.PhoneNum = value[1]
		info.Province = value[2]
		info.City = value[3]
		info.CardType = value[4]
		info.AreaZone = value[5]
		info.ZipCode = value[6]
		recordMap[info.PhoneNum] = info
	}

	rec, ok := recordMap["1392462"]
	if ok {
		fmt.Println(rec)
	} else {
		fmt.Println("not found")
	}
}
