package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type defibrillatorInfo struct {
	ID        int
	Name      string
	Address   string
	Phone     string
	Longitude float64
	Latitude  float64
}

func main() {

	//parse input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	longitude := AsciiFloatToFloat(scanner.Text())
	scanner.Scan()
	latitude := AsciiFloatToFloat(scanner.Text())
	var n int
	scanner.Scan()
	fmt.Sscanln(scanner.Text(), &n)
	defibrillatorList := make([]defibrillatorInfo, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		defibrillatorList[i] = GetDefibrillatorInfoFromString(scanner.Text())
	}
	//perform logic
	outputIndex := 0
	distance := GetDistanceBetweenHumanAndAID(longitude, latitude, defibrillatorList[0])
	for i := 1; i < n; i++ {
		newDistance := GetDistanceBetweenHumanAndAID(longitude, latitude, defibrillatorList[i])
		if newDistance < distance {
			outputIndex = i
			distance = newDistance
		}
	}
	//output
	fmt.Print(defibrillatorList[outputIndex].Name)

}

func GetDefibrillatorInfoFromString(input string) (output defibrillatorInfo) {
	info := strings.Split(input, ";")
	var err error
	output.ID, err = strconv.Atoi(info[0])
	output.Name = info[1]
	output.Address = info[2]
	output.Phone = info[3]
	output.Longitude = AsciiFloatToFloat(info[4])
	output.Latitude = AsciiFloatToFloat(info[5])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func AsciiFloatToFloat(input string) (output float64) {
	input = strings.Replace(input, ",", ".", -1)
	fmt.Sscanf(input, "%f", &output)
	return
}

func GetDistanceBetweenHumanAndAID(longitude, lattitude float64, info defibrillatorInfo) float64 {
	x := (info.Longitude - longitude) * math.Cos((lattitude+info.Latitude)/2.0)
	y := info.Latitude - lattitude
	d := math.Sqrt(x*x+y*y) * 6371
	return d
}
