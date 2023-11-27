package cinema

import (
	"fmt"
	"math"
)

type Cinema struct {
	numRows     int
	seatsPerRow int
	layout      [][]string
}

func NewCinema(numRows int, seatsPerRow int) Cinema {
	c := Cinema{numRows, seatsPerRow, make([][]string, numRows)}
	for i := range c.layout {
		c.layout[i] = make([]string, seatsPerRow)
	}
	c.initializeLayout()
	return c
}

func (c Cinema) initializeLayout() {
	for row := range c.layout {
		for col := range c.layout[row] {
			c.layout[row][col] = "S"
		}
	}
}

func (c Cinema) fillSeat(rowNum int, colNum int) {
	c.layout[rowNum][colNum] = "B"
}

func (c Cinema) openSeat(rowNum int, colNum int) {
	c.layout[rowNum][colNum] = "S"
}

func (c Cinema) printTicketPrice(rowNum int) {
	numSeatsTotal := c.numRows * c.seatsPerRow
	halfwayRow := int(math.Floor(float64(c.numRows) / float64(2)))
	if numSeatsTotal > 60 && rowNum <= halfwayRow {
		fmt.Println("Ticket price: $10")
	} else {
		fmt.Println("Ticket price: $8")
	}

}

func (c Cinema) SelectSeat(rowNum int, colNum int) {
	c.printTicketPrice(rowNum)
	c.fillSeat(rowNum, colNum)
}

func (c Cinema) PrintCinema() {
	fmt.Println("Cinema:")
	fmt.Print(" ")
	for i := 0; i < c.seatsPerRow; i++ {
		fmt.Printf(" %d", i+1)
	}
	fmt.Print("\n")
	for j := 0; j < c.numRows; j++ {
		fmt.Printf("%d", j+1)
		for seat := range c.layout[j] {
			fmt.Printf(" %s", c.layout[j][seat])
		}
		fmt.Print("\n")
	}
}
