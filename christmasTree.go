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

	for level <= levelMax {
		branch := 0
		for branch < branchMax {
			if level == levelMax {
				if rand.Intn(2) == 0 {
					fmt.Print("🌲")
				} else {
					fmt.Print("🎄")
				}
			} else if (branch >= branchMiddle-1 && branch <= branchMiddle+1) && (level >= levelMax-4) {
				fmt.Print("🪵")
				groveUp = false
			} else if branch < branchMiddle-level+1 || branch > branchMiddle+level-1 || level > levelMax-4 {
				salut := rand.Intn(50)
				if level < 10 && (branch < branchMiddle-level-2 || branch > branchMiddle+level+2) {
					switch salut {
					case 1:
						fmt.Print("✨")
					case 2:
						fmt.Print("💥")
					default:
						fmt.Print("  ")
					}
				} else {
					fmt.Print("  ")
				}
			} else if branch == branchMiddle && level == 1 {
				fmt.Print("🌟")
			} else {
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

	happyString := " " + brandString + " A HAPPY NEW 2023 YEAR!!! "
	for len(happyString)%4 != 0 {
		if len(happyString)%2 == 0 {
			happyString += " "
		} else {
			happyString = " " + happyString
		}
	}
	firstSets := branchMax*2 - len(happyString)
	w := 0
	happyStringStar := ""
	for w < firstSets {
		if rand.Intn(2) == 0 {
			happyStringStar = happyStringStar + "💥"
			w += len("💥")
		} else {
			happyStringStar = happyStringStar + "🎊"
			w += len("🎊")
		}
	}
	happyString = happyStringStar + happyString + happyStringStar
	fmt.Print(happyString + "\n")

	branch := 0
	for branch < branchMax {
		if rand.Intn(2) == 0 {
			fmt.Print("🌲")
		} else {
			fmt.Print("🎄")
		}
		branch++
	}
	fmt.Print("\n")
}

func main() {
	args := os.Args[1:]
	var brandString string = ""
	if len(args) > 0 {
		brandString = args[0]
	}
	for true {
		christmasTree(30, 17, brandString)
		time.Sleep(time.Second)
		fmt.Printf("\x1bc")
	}
}
