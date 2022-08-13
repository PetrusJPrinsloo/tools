package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var m map[string]string

func main() {

	m = make(map[string]string)
	m["0"] = "0000"
	m["1"] = "0001"
	m["2"] = "0010"
	m["3"] = "0011"
	m["4"] = "0100"
	m["5"] = "0101"
	m["6"] = "0110"
	m["7"] = "0111"
	m["8"] = "1000"
	m["9"] = "1001"
	m["A"] = "1010"
	m["B"] = "1011"
	m["C"] = "1100"
	m["D"] = "1101"
	m["E"] = "1110"
	m["F"] = "1111"

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				checkFile()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	fmt.Println("Press the ENTER key to stop")
	fmt.Scanln()
	fmt.Println("done")
}

func checkFile() {
	count := 0

	f, err := os.Open("D:\\SteamLibrary\\steamapps\\common\\Half-Life 2\\hl2\\gamestate.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var text string

	for scanner.Scan() {
		count++

		if count == 217 {
			text = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	found := false
	hexString := ""
	for _, v := range text {

		if v == 'x' && !found {
			found = true
			continue
		}

		if found && v != '"' {
			hexString += string(v)
		}
	}

	fmt.Println(Hex2Bin(hexString))
}

func Hex2Bin(hex string) string {
	var bin string
	for _, v := range hex {
		bin += m[string(v)]
	}
	return bin
}
