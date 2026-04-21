# Arduino Uno SMD - Emoji in 8x8 Led Matrix wiht Golang

Emoji in 8x8 Led Matrix with an arduino uno smd and Golang.

* VS Code
* Golang 1.26.2
* TinyGo 0.41.0 [TinyGo](tinygo.org)
* TinyGo Toolkit for VS Code (amken.tinygo-toolkit)

## TinyGo Toolkit

* Connect board (arduino uno smd)
* View > Command Palette > TinyGo: Initilize TinyGo Project
* View > Command Palette > TinyGo: Configure Intellisense
* Project > Target (arduino)
* Project > Port (/dev/ttyUSB0)
* Project Baud Rate (9600)
* Project > Flash

## Details

* Uses PC2 to PC5 (ADC2 to ADC5) as digital input along D2 to D13
* After setting each column, the program clears the entire matrix, giving the ilusion of all columns being active at the same time
* It is necessary to use pointers to not flick the leds between each loop
