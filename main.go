package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func isFKeyPressed() bool {
	user32 := syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState := user32.NewProc("GetAsyncKeyState")
	keyState, _, _ := getAsyncKeyState.Call(uintptr(0x46)) // 0x46 is 'F'
	return keyState&(1<<15) != 0
}

func moveMouse(x, y int) {
	user32 := syscall.NewLazyDLL("user32.dll")
	setCursorPos := user32.NewProc("SetCursorPos")
	setCursorPos.Call(uintptr(x), uintptr(y))
}

func getScreenMetrics() (int, int) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")

	// SM_CXSCREEN = 0, SM_CYSCREEN = 1
	width, _, _ := getSystemMetrics.Call(0)
	height, _, _ := getSystemMetrics.Call(1)

	return int(width), int(height)
}

func moveMouseFigure8() {
	// Get actual screen dimensions
	screenWidth, screenHeight := getScreenMetrics()

	centerX, centerY := screenWidth/2, screenHeight/2

	// Smaller radii to ensure visibility
	radiusX := float64(screenWidth) * 0.2
	radiusY := float64(screenHeight) * 0.15

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Now do the figure 8
	t := 0.0
	for {
		select {
		case <-stop:
			fmt.Println("\nExiting...")
			return
		case <-ticker.C:
			if isFKeyPressed() {
				fmt.Println("\n'F' key pressed, stopping...")
				return
			}

			// Use a different approach for the figure 8
			// This is a lemniscate of Bernoulli
			denom := 1.0 + math.Pow(math.Sin(t), 2)

			// Calculate x and y offsets from center
			xOffset := radiusX * math.Cos(t) / denom
			yOffset := radiusY * math.Sin(t) * math.Cos(t) / denom

			// Apply offsets to center coordinates
			x := centerX + int(xOffset)
			y := centerY + int(yOffset)

			moveMouse(x, y)
			t += 0.02
		}
	}
}

func main() {
	var figure8Cmd = &cobra.Command{
		Use:   "figure8",
		Short: "Moves the mouse in a figure 8 pattern until exited or 'F' key is pressed",
		Run: func(cmd *cobra.Command, args []string) {
			moveMouseFigure8()
		},
	}

	if err := figure8Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
