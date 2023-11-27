package main

import (
	"fmt"
	"log"
	"math"
)

type Cinema struct {
	numRows       int
	seatsPerRow   int
	seatsFilled   int
	totalIncome   int
	currentIncome int
	layout        [][]string
}

func NewCinema(numRows int, seatsPerRow int) Cinema {
	c := Cinema{numRows, seatsPerRow, 0, 0, 0, make([][]string, numRows)}
	for i := range c.layout {
		c.layout[i] = make([]string, seatsPerRow)
	}
	c.initializeLayout()
	c.calculateTotalPossibleIncome()
	return c
}

func (c *Cinema) initializeLayout() {
	for row := range c.layout {
		for col := range c.layout[row] {
			c.layout[row][col] = "S"
		}
	}
}

func (c *Cinema) calculateTotalPossibleIncome() {
	totalSeats := c.numRows * c.seatsPerRow
	if totalSeats <= 60 {
		c.totalIncome = totalSeats * 10
	} else {
		halfwayRow := int(math.Floor(float64(c.numRows) / float64(2)))
		c.totalIncome = c.seatsPerRow * ((halfwayRow)*10 + (c.numRows-halfwayRow)*8)
	}
}

func (c *Cinema) IsSeatAvailable(rowNum int, colNum int) bool {
	return c.layout[rowNum-1][colNum-1] != "B"
}

func (c *Cinema) fillSeat(rowNum int, colNum int) {
	c.layout[rowNum-1][colNum-1] = "B"
	c.seatsFilled += 1
}

func (c *Cinema) openSeat(rowNum int, colNum int) {
	c.layout[rowNum-1][colNum-1] = "S"
	c.seatsFilled -= 1
}

func (c *Cinema) getTicketPrice(rowNum int) int {
	numSeatsTotal := c.numRows * c.seatsPerRow
	halfwayRow := int(math.Floor(float64(c.numRows) / float64(2)))
	if numSeatsTotal <= 60 || rowNum <= halfwayRow {
		return 10
	} else {
		return 8
	}
}

func (c *Cinema) isSeatSelectionValid(rowNum int, colNum int) bool {
	return rowNum > 0 && rowNum <= c.numRows && colNum > 0 && colNum <= c.seatsPerRow
}

func (c *Cinema) SelectSeat(rowNum int, colNum int) {
	fmt.Println()
	if !c.isSeatSelectionValid(rowNum, colNum) {
		fmt.Println("Wrong input!")
	} else if c.IsSeatAvailable(rowNum, colNum) {
		ticketPrice := c.getTicketPrice(rowNum)
		fmt.Printf("Ticket price: $%d\n", ticketPrice)
		c.fillSeat(rowNum, colNum)
		c.currentIncome += ticketPrice
	} else {
		fmt.Println("That ticket has already been purchased!")
	}
	fmt.Println()
}

func (c *Cinema) PrintStats() {
	percentSeatsFilled := float64(100) * (float64(c.seatsFilled) / float64(c.numRows*c.seatsPerRow))
	fmt.Printf("Number of purchased tickets: %d\n", c.seatsFilled)
	fmt.Printf("Percentage: %.2f%%\n", percentSeatsFilled)
	fmt.Printf("Current income: %d\n", c.currentIncome)
	fmt.Printf("Total income: %d\n", c.totalIncome)
	fmt.Println()
}

func (c *Cinema) PrintCinema() {
	fmt.Println("Cinema:")
	fmt.Print(" ")
	for i := 0; i < c.seatsPerRow; i++ {
		fmt.Printf(" %d", i+1)
	}
	fmt.Println()
	for j := 0; j < c.numRows; j++ {
		fmt.Printf("%d", j+1)
		for seat := range c.layout[j] {
			fmt.Printf(" %s", c.layout[j][seat])
		}
		fmt.Println()
	}
	fmt.Println()
}

func getUserInput(msg string) int {
	fmt.Println(msg)
	var n int64
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

func cinemaMenu(cinema Cinema) {
	displayMenu := true

	for displayMenu {
		userChoice := getUserInput("1. Show the seats\n2. Buy a ticket \n3. Statistics \n0. Exit")
		fmt.Println()
		if userChoice == 0 {
			displayMenu = false
		} else if userChoice == 1 {
			cinema.PrintCinema()
		} else if userChoice == 2 {
			chosenRow := getUserInput("Enter a row number:")
			chosenCol := getUserInput("Enter a seat number in that row:")
			cinema.SelectSeat(chosenRow, chosenCol)
		} else if userChoice == 3 {
			cinema.PrintStats()
		}
	}

}

func main() {
	numRows := getUserInput("Enter the number of rows:")
	numSeats := getUserInput("Enter the number of seats in each row:")
	cinema := NewCinema(numRows, numSeats)
	fmt.Println()
	cinemaMenu(cinema)
}
