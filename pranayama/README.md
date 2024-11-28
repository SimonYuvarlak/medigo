# Pranayama

Pranayama is a breath control practice in yoga that enhances physical and mental well-being. This project focuses on providing guidance and tools for practicing various pranayama techniques.

## Features

- Time based breathing exercises
- Timer for each breath
- Cycles for each practice

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/SimonYuvarlak/medigo
cd pranayama
```

2. Install dependencies:
```bash
go mod download
```

## Usage

### Basic Run
```bash
go run main.go
```

### Running with Different Practices

go run main.go -intervals=<intervals> -cycles=<# of cycles>
For example:
```bash
go run main.go -intervals=4,4,4,4 -cycles=1
```
- Means, run the gong sound in every for minutes, 4 times.
- This will be one cycle.
- If you want more cycles you can say go run main.go -intervals=4,4,4,4 -cycles=3
- If you want infinite cycles you can say go run main.go -intervals=4,4,4,4 -cycles=0
- If you want different practices you can say go run main.go -intervals=4, 7, 8 -cycles=<# of cycles>
- This will be 4, 7, 8 seconds breath with gong sound for each breath.