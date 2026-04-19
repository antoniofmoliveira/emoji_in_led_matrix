package main

import (
	"machine"
)

func main() {

	// PC2 = ADC2 = D16 ... PC5 = ADC5 = D19
	rows := []machine.Pin{machine.D4, machine.PC2, machine.D7, machine.D5, machine.D13, machine.D8, machine.D12, machine.D6}
	cols := []machine.Pin{machine.PC5, machine.D11, machine.D10, machine.D3, machine.D9, machine.D2, machine.PC4, machine.PC3}

	for i := range rows {
		rows[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
		cols[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// !!! [cols][rows]
	emoji := [][]bool{
		{true, true, false, false, false, false, true, true},
		{true, false, true, true, true, true, false, true},
		{false, true, false, true, true, false, true, false},
		{false, true, true, true, true, true, true, false},
		{false, true, false, true, true, false, true, false},
		{false, true, true, false, false, true, true, false},
		{true, false, true, true, true, true, false, true},
		{true, true, false, false, false, false, true, true},
	}
	clearMatrix(rows, cols)

	for {
		for c := range cols {
			cols[c].High()
			for l := range rows {
				rows[l].Set(emoji[l][c])
			}
			clearMatrix(rows, cols)

		}

	}

	// for {
	// 	cols[0].High()
	// 	rows[0].High()
	// 	rows[1].High()
	// 	rows[2].Low()
	// 	rows[3].Low()
	// 	rows[4].Low()
	// 	rows[5].Low()
	// 	rows[6].High()
	// 	rows[7].High()

	// 	clearMatrix(rows, cols)

	// 	cols[1].High()
	// 	rows[0].High()
	// 	rows[1].Low()
	// 	rows[2].High()
	// 	rows[3].High()
	// 	rows[4].High()
	// 	rows[5].High()
	// 	rows[6].Low()
	// 	rows[7].High()
	// 	clearMatrix(rows, cols)

	// 	cols[2].High()
	// 	rows[0].Low()
	// 	rows[1].High()
	// 	rows[2].Low()
	// 	rows[3].High()
	// 	rows[4].Low()
	// 	rows[5].High()
	// 	rows[6].High()
	// 	rows[7].Low()
	// 	clearMatrix(rows, cols)

	// 	cols[3].High()
	// 	rows[0].Low()
	// 	rows[1].High()
	// 	rows[2].High()
	// 	rows[3].High()
	// 	rows[4].High()
	// 	rows[5].Low()
	// 	rows[6].High()
	// 	rows[7].Low()
	// 	clearMatrix(rows, cols)

	// 	cols[4].High()
	// 	rows[0].Low()
	// 	rows[1].High()
	// 	rows[2].High()
	// 	rows[3].High()
	// 	rows[4].High()
	// 	rows[5].Low()
	// 	rows[6].High()
	// 	rows[7].Low()
	// 	clearMatrix(rows, cols)

	// 	cols[5].High()
	// 	rows[0].Low()
	// 	rows[1].High()
	// 	rows[2].Low()
	// 	rows[3].High()
	// 	rows[4].Low()
	// 	rows[5].High()
	// 	rows[6].High()
	// 	rows[7].Low()
	// 	clearMatrix(rows, cols)

	// 	cols[6].High()
	// 	rows[0].High()
	// 	rows[1].Low()
	// 	rows[2].High()
	// 	rows[3].High()
	// 	rows[4].High()
	// 	rows[5].High()
	// 	rows[6].Low()
	// 	rows[7].High()
	// 	clearMatrix(rows, cols)

	// 	cols[7].High()
	// 	rows[0].High()
	// 	rows[1].High()
	// 	rows[2].Low()
	// 	rows[3].Low()
	// 	rows[4].Low()
	// 	rows[5].Low()
	// 	rows[6].High()
	// 	rows[7].High()
	// 	clearMatrix(rows, cols)

	// }
}

func clearMatrix(rows []machine.Pin, cols []machine.Pin) {
	for i := range rows {
		cols[i].Low()
		rows[i].High()
	}

}
