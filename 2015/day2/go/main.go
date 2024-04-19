// --- Day 2: I Was Told There Would Be No Math ---
// The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.
//
// Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.
//
// For example:
//
// A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
// A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
// All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part 1
type Dimensions struct {
	length int
	width  int
	height int
}

func boxSurfaceArea(dimensions Dimensions) int {
	surfaceArea := 2*dimensions.length*dimensions.width + 2*dimensions.width*dimensions.height + 2*dimensions.height*dimensions.length

	// Calculate the areas of each pair of dimensions
	area1 := dimensions.length * dimensions.width
	area2 := dimensions.width * dimensions.height
	area3 := dimensions.height * dimensions.length

	// Find the smallest area among the three
	smallestArea := area1
	if area2 < smallestArea {
		smallestArea = area2
	}
	if area3 < smallestArea {
		smallestArea = area3
	}

	totalArea := surfaceArea + smallestArea
	return totalArea
}

func totalWrappingPaper() {
	// Open file, close it in case of error
	file, err := os.Open("dimensoes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scanner to read the file
	scanner := bufio.NewScanner(file)

	var totalWrappingPaper int

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		dims := strings.Split(line, "x") // Each dimension will be separated by 'x'
		if len(dims) != 3 {
			fmt.Println("Invalid dimensions format", line)
			continue
		}

		// Convert str to int, in this case length will be converted to int on the file
		length, err := strconv.Atoi(dims[0])
		if err != nil {
			fmt.Println("Error parsing length:", err)
			continue
		}

		width, err := strconv.Atoi(dims[1])
		if err != nil {
			fmt.Println("Error parsing width:", err)
			continue
		}

		height, err := strconv.Atoi(dims[2])
		if err != nil {
			fmt.Println("Error parsing height:", err)
			continue
		}

		dimensions := Dimensions{length: length, width: width, height: height}
		wrappingPaperNeeded := boxSurfaceArea(dimensions)
		totalWrappingPaper += wrappingPaperNeeded
	}
	fmt.Printf("Total wrapping paper needed: %d square feet\n", totalWrappingPaper)
}

// --- Part Two ---
// The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.
//
// The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.
//
// For example:
//
// A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
// A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
// How many total feet of ribbon should they order?

func calculateRibbon() {
	// Open file, close it in case of error
	file, err := os.Open("dimensoes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scanner to read the file
	scanner := bufio.NewScanner(file)

	var totalRibbon int

	// Read and process lines
	for scanner.Scan() {
		line := scanner.Text()
		dims := strings.Split(line, "x") // Each dimension will be separated by 'x'
		if len(dims) != 3 {
			fmt.Println("Invalid dimensions format", line)
			continue
		}

		// Convert str to int
		length, err := strconv.Atoi(dims[0])
		if err != nil {
			fmt.Println("Error parsing length:", err)
			continue
		}

		width, err := strconv.Atoi(dims[1])
		if err != nil {
			fmt.Println("Error parsing width:", err)
			continue
		}

		height, err := strconv.Atoi(dims[2])
		if err != nil {
			fmt.Println("Error parsing height:", err)
			continue
		}

		dimensions := Dimensions{length: length, width: width, height: height}
		ribbonNeeded := calculateRibbonForOnePresent(dimensions)
		totalRibbon += ribbonNeeded
	}
	fmt.Printf("They should order %d feet of ribbon\n", totalRibbon)
}

func calculateRibbonForOnePresent(dimensions Dimensions) int {
	// Calculate the smallest perimeter among the three faces
	smallestPerimeter := min(min(2*dimensions.length+2*dimensions.width, 2*dimensions.width+2*dimensions.height), 2*dimensions.height+2*dimensions.length)

	// Calculate the ribbon needed for the bow
	bow := dimensions.length * dimensions.width * dimensions.height

	// Total feet of ribbon needed
	totalFeetOfRibbon := smallestPerimeter + bow
	return totalFeetOfRibbon
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	totalWrappingPaper()
	calculateRibbon()
}
