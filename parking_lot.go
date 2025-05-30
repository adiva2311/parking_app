package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Slot struct {
	Number           int
	RegistrationNumb string
	Color            string
	Occupied         bool
}

type ParkingLot struct {
	Slots map[int]*Slot
	Size  int
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{Slots: make(map[int]*Slot)}
}

func (p *ParkingLot) Create(sizeStr string) {
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		fmt.Println("Invalid size")
		return
	}
	p.Size = size
	for i := 1; i <= size; i++ {
		p.Slots[i] = &Slot{Number: i}
	}
	fmt.Println("Created parking lot with", size, "slots")
}

func (p *ParkingLot) Park(regNum, color string) {
	for i := 1; i <= p.Size; i++ {
		slot := p.Slots[i]
		if !slot.Occupied {
			slot.Occupied = true
			slot.RegistrationNumb = regNum
			slot.Color = color
			fmt.Printf("|PARK| %s is Parking in slot number: %d\n", slot.RegistrationNumb, slot.Number)
			return
		}
	}
	fmt.Printf("Cannot Park %s || Sorry, parking lot is full\n", regNum)
}

func (p *ParkingLot) FindByColor(color string) {
	isFound := false
	for _, slot := range p.Slots {
		if slot.Occupied && strings.EqualFold(slot.Color, color) {
			fmt.Printf("%s Car is founded with RegNumber %s\n", strings.ToUpper(slot.Color), slot.RegistrationNumb)
			isFound = true
		}
	}
	if !isFound {
		fmt.Println("Not found")
	}
}

func (p *ParkingLot) Leave(regNum string, hoursStr string) {
	hours, err := strconv.Atoi(hoursStr)
	if err != nil || hours <= 0 {
		fmt.Println("Invalid hours")
		return
	}
	for _, slot := range p.Slots {
		if slot.Occupied && slot.RegistrationNumb == regNum {
			slot.Occupied = false
			slot.RegistrationNumb = ""
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}
			fmt.Printf("|LEAVE| %s with Slot Number %d is Leaving => $%d\n", regNum, slot.Number, charge)
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", regNum)
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. | Registration No.")
	keys := make([]int, 0, len(p.Slots))
	for k := range p.Slots {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, i := range keys {
		slot := p.Slots[i]
		if slot.Occupied {
			fmt.Printf("Slot %d | %s\n", slot.Number, slot.RegistrationNumb)
		}
	}
}
