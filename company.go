package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	amount int
}

type TimeAndMaterial struct {
	projectName string
	hours int
	rate int
}

type Advertisement struct {
	projectName string
	costPerClick  int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.amount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.hours * tm.rate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.costPerClick * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.projectName
}

func calculateNetIncome(income []Income) {
	var netIncome int = 0
	for _, v := range income {
		fmt.Printf("Income from %s is %d \n", v.source(), v.calculate())
		netIncome += v.calculate()
	}
	fmt.Printf("Net income %d", netIncome)
}

func main() {
	project1 := FixedBilling{
		projectName: "Project1",
		amount:      10000,
	}
	project2 := FixedBilling{
		projectName: "Project2",
		amount:      20000,
	}
	project3 := TimeAndMaterial{
		projectName: "Project3",
		hours:       120,
		rate:        39,
	}

	project4 := Advertisement{
		projectName:  "Project4",
		costPerClick: 190,
		noOfClicks:   2300,
	}

	incomeStreams := []Income{project1, project2, project3, project4}
	calculateNetIncome(incomeStreams)
}
