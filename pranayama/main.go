package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	soundFile = "sounds/gong.wav"
)

// example run: go run main.go -intervals=4,4,4,4 -cycles=1
func main() {
	// Parse command line arguments
	intervalsStr := flag.String("intervals", "4,4,4,4",
		"Comma-separated intervals in seconds")
	cycles := flag.Int("cycles", 1, "Number of cycles (0 for infinite)")
	flag.Parse()

	// Convert intervals string to []int
	intervalStrs := strings.Split(*intervalsStr, ",")
	intervals := make([]int, len(intervalStrs))
	for i, s := range intervalStrs {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Invalid interval value: %s\n", s)
			os.Exit(1)
		}
		if val <= 0 {
			fmt.Printf("Interval must be positive: %s\n", s)
			os.Exit(1)
		}
		intervals[i] = val
	}

	// Initialize audio
	f, err := os.Open(soundFile)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", soundFile, err)
		os.Exit(1)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		fmt.Printf("Error decoding %s: %v\n", soundFile, err)
		os.Exit(1)
	}
	defer streamer.Close()

	// Create buffer once before the loop
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Start practice
	fmt.Println("Starting pranayama practice...")
	fmt.Printf("Intervals: %v\n", intervals)
	fmt.Printf("Cycles: %d\n", *cycles)

	cycleCount := 0
	for *cycles == 0 || cycleCount < *cycles {
		fmt.Printf("\nCycle %d\n", cycleCount+1)
		for i, interval := range intervals {
			// Play sound using the pre-created buffer
			sound := buffer.Streamer(0, buffer.Len())
			speaker.Play(sound)

			fmt.Printf("Gong played at: %s\n", time.Now().Format("15:04:05"))
			fmt.Printf("Interval %d: waiting %d seconds...\n", i+1, interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
		cycleCount++
	}
	fmt.Println("\nPranayama practice completed!")
}
