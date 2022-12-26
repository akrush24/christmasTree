package main

import (
	"fmt"
	"math/rand"
	"time"
)

func christmasTree(branchMax int, levelMax int) {
	level := 1
	branchMiddle := branchMax / 2
	grovUp := true
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
			} else if (branch >= branchMiddle-1 && branch <= branchMiddle+1) && (level >= levelMax-3) {
				fmt.Print("🪵")
			} else if branch < branchMiddle-level+1 || branch > branchMiddle+level-1 || level > levelMax-3 {
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
		if !grovUp {
			grovUp = true
			level++
		} else if level == levelMax {
			level++
		} else {
			grovUp = false
		}
	}

	happyString := " HAPPY NEW 2023 YEAR! "
	firstSets := branchMax*2 - len(happyString)
	w := 0
	happyStringStar := ""
	for w < firstSets {
		happyStringStar = happyStringStar + "💫"
		w += len("💫")
	}
	happyString = happyStringStar + happyString + happyStringStar
	fmt.Print(happyString)
	fmt.Print("\n")
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
	christmasTree(31, 15)
}
