package main

import (
	"machine"
	"time"
)

func main() {

	// PC2 = ADC2 = D16 ... PC5 = ADC5 = D19
	rows := []machine.Pin{machine.D4, machine.PC2, machine.D7, machine.D5, machine.D13, machine.D8, machine.D12, machine.D6}
	cols := []machine.Pin{machine.PC5, machine.D11, machine.D10, machine.D3, machine.D9, machine.D2, machine.PC4, machine.PC3}

	for i := range rows {
		rows[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
		cols[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	emoji1 := [8][8]bool{}
	emoji2 := [8][8]bool{}

	emoji1 = convertToMatrix(&[8]string{
		"11000011",
		"10111101",
		"01011010",
		"01111110",
		"01011010",
		"01100110",
		"10111101",
		"11000011"})

	emoji2 = convertToMatrix(&[8]string{
		"11000011",
		"10111101",
		"01011010",
		"01111110",
		"01100110",
		"01011010",
		"10111101",
		"11000010"})

	for {
		clearMatrix(&rows, &cols)

		nowMore2Secs := time.Now().Add(2 * time.Second)
		for time.Now().Before(nowMore2Secs) {
			show(&cols, &rows, &emoji1)
		}

		clearMatrix(&rows, &cols)

		nowMore2Secs2 := time.Now().Add(2 * time.Second)
		for time.Now().Before(nowMore2Secs2) {
			show(&cols, &rows, &emoji2)
		}
	}

}

func show(cols *[]machine.Pin, rows *[]machine.Pin, matrix *[8][8]bool) {
	for c := range *cols {
		(*cols)[c].High()
		for r := range *rows {
			(*rows)[r].Set(matrix[r][c])
		}
		clearMatrix(rows, cols)
	}
}

func clearMatrix(rows *[]machine.Pin, cols *[]machine.Pin) {
	for i := range *rows {
		(*cols)[i].Low()
		(*rows)[i].High()
	}
}

func convertToMatrix(t *[8]string) [8][8]bool {
	response := [8][8]bool{}
	for row := range 8 {
		for col := range 8 {
			response[row][col] = t[row][col] == '1'
		}
	}
	return response
}
