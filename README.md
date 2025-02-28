# Figure 8 Cursor Movement

A simple Go application that moves your cursor in a figure 8 pattern.

## Overview

This Go application uses Windows system calls to programmatically control the mouse cursor position, moving it in a figure 8 pattern (lemniscate). It automatically detects your screen resolution and scales the pattern accordingly.

## Features

- Automatic screen resolution detection
- Smooth figure 8 cursor movement
- Stop functionality by pressing the 'F' key or with CTRL+C
- Uses the Cobra CLI framework for command structure

## Requirements

- Windows operating system
- Go 1.16 or higher
- [Cobra CLI library](https://github.com/spf13/cobra)

## Installation

1. Clone this repository:

```
git clone https://github.com/BiodigitalJaz/figure8.git
cd figure8
```

2. Install dependencies:

```
go get github.com/spf13/cobra
```

3. Build the application:

```
go build -o figure8.exe
```

## Usage

Run the application from command line:

```
figure8.exe
```

To stop the cursor movement:
- Press the 'F' key
- Or use CTRL+C in the terminal

## How It Works

The application uses parametric equations based on the lemniscate of Bernoulli to create a mathematically accurate figure 8 pattern. It calculates cursor positions using:

```go
denom := 1.0 + math.Pow(math.Sin(t), 2)
xOffset := radiusX * math.Cos(t) / denom
yOffset := radiusY * math.Sin(t) * math.Cos(t) / denom
```

These equations generate a perfect figure 8 shape that is properly scaled to your screen dimensions.

## Customization

You can modify the following parameters in the code to customize the behavior:

- `radiusX` and `radiusY` - Control the size of the figure 8 (currently set to 20% and 15% of screen dimensions)
- The `t += 0.02` value - Controls the speed of movement (smaller values = slower movement)
- Ticker interval - Currently set to 10ms, controls how frequently the cursor updates