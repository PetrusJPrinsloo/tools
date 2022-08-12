package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

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
	// 00000001100001000000011101000000000000000000
	//skip first 5 characters because they are not used
	//hexString = hexString[5:]

	//for _, v := range hexString {
	//	fmt.Printf("%s", Hex2Bin(v))
	//}

	toPrint, _ := Hex2Bin(hexString)

	fmt.Println(Hex2Bin(toPrint))
}

//func Hex2Bin(in rune) string {
//	var out []rune
//	for i := 3; i >= 0; i-- {
//		b := in >> uint(i)
//		out = append(out, (b%2)+48)
//	}
//	return string(out)
//}

func Hex2Bin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016b", ui), nil
}
