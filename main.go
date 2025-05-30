package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	parking := NewParkingLot()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "create_parking_lot":
			if len(words) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			parking.Create(words[1])
		case "park":
			if len(words) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			parking.Park(words[1], words[2])
		case "find_by_color":
			if len(words) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			parking.FindByColor(words[1])
		case "leave":
			if len(words) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			parking.Leave(words[1], words[2])
		case "status":
			parking.Status()
		default:
			fmt.Println("Unknown command:", words[0])
		}
	}
}
