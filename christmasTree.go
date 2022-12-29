package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func christmasTree(branchMax int, levelMax int, brandString string) {
	level := 1
	branchMiddle := branchMax / 2
	groveUp := true
	rand.Seed(time.Now().UnixNano())

	var x, y int = 0, 0
	for ; y < 5; y++ {
		for ; x < branchMax; x++ {
			switch rand.Intn(50) {
			case 1, 2, 3:
				fmt.Print("✨")
			case 4:
				fmt.Print("💥")
			default:
				fmt.Print("  ")
			}
		}
		x = 0
		fmt.Println("")
	}

	for level <= levelMax {
		branch := 0
		for branch < branchMax {
			if level == levelMax {
				// bottom
				if rand.Intn(2) == 0 {
					fmt.Print("🌲")
				} else {
					fmt.Print("🎄")
				}
			} else if (branch >= branchMiddle-1 && branch <= branchMiddle+1) && (level >= levelMax-4) {
				// trunk
				fmt.Print("🪵")
				groveUp = false
			} else if branch < branchMiddle-level+1 || branch > branchMiddle+level-1 || (level > levelMax-2 || level < 2) {
				// sky
				if level < levelMax-5 && (branch < branchMiddle-level-1 || branch > branchMiddle+level+1) {
					switch rand.Intn(50) {
					case 1:
						fmt.Print("✨")
					case 2, 3, 4, 5, 6:
						fmt.Print("💥")
					default:
						fmt.Print("  ")
					}
				} else if branch == branchMiddle {
					fmt.Print("🌟")
				} else {
					fmt.Print("  ")
				}
			} else {
				// decorations of the tree
				decor := rand.Intn(12)
				switch decor {
				case 0:
					fmt.Print("🔵")
				case 1:
					fmt.Print("🎁")
				case 2:
					fmt.Print("🟠")
				case 3:
					fmt.Print("🔵")
				case 4:
					fmt.Print("🔔")
				case 5:
					fmt.Print("🍍")
				case 6:
					fmt.Print("🔴")
				case 7:
					fmt.Print("🎀")
				case 8:
					fmt.Print("⭐")
				case 9:
					fmt.Print("🔴")
				default:
					fmt.Print("🟢")
				}
			}
			branch++
		}
		fmt.Print("\n")

		if !groveUp {
			groveUp = true
			level++
		} else if level == levelMax {
			level++
		} else {
			groveUp = false
		}
	}
	if brandString != "" {
		brandString = " " + brandString
	}
	happyString := brandString + " HAPPY NEW 2023 YEAR!!! "
	for len(happyString)%4 != 0 {
		if len(happyString)%2 == 0 {
			happyString += " "
		} else {
			happyString = " " + happyString
		}
	}

	firstSets := branchMax*2 - len(happyString)
	w := 0
	happyStringStart := ""
	for w < firstSets {
		if rand.Intn(2) == 0 {
			happyStringStart = happyStringStart + "💥"
			w += len("💥")
		} else {
			happyStringStart = happyStringStart + "🎊"
			w += len("🎊")
		}
	}

	w = 0
	happyStringEnd := ""
	for w < firstSets-4 {
		if rand.Intn(2) == 0 {
			happyStringEnd = happyStringEnd + "💥"
			w += len("💥")
		} else {
			happyStringEnd = happyStringEnd + "🎊"
			w += len("🎊")
		}
	}

	happyString = happyStringStart + happyString + happyStringEnd

	fmt.Print(happyString + "\n")
	branch := 0
	bottomLine2 := ""
	for branch < branchMax {
		if rand.Intn(2) == 0 {
			bottomLine2 += "🌲"
		} else {
			bottomLine2 += "🎄"
		}
		branch++
	}
	fmt.Println(bottomLine2)
	fmt.Print("\n")
}

func main() {
	args := os.Args[1:]
	var brandString string = ""
	if len(args) > 0 {
		brandString = args[0]
	}
	// clear console before first print
	fmt.Printf("\x1b[2J")
	for true {
		christmasTree(31, 17, brandString)
		time.Sleep(time.Second)
		// clear console
		fmt.Printf("\x1b[2J")
	}
}
