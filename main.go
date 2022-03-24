package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var notepad []string
	var max int64
	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&max)

	run(&notepad, max)
}

func run(notepad *[]string, max int64) {
	input := readInput()
	command := input[0]
	if command == "exit" {
		fmt.Println("[Info] Bye!")
		return
	} else if command == "create" {
		if int64(len(*notepad)) == max {
			fmt.Println("[Error] Notepad is full")
		} else {
			subInput := input[1:]
			if len(subInput) == 0 {
				fmt.Println("[Error] Missing note argument")
			} else {
				note := strings.Trim(strings.Join(subInput, " "), " ")
				if note == "" {
					fmt.Println("[Error] Missing note argument")
				} else {
					*notepad = append(*notepad, note)
					fmt.Println("[OK] The note was successfully created")
				}
			}
		}
	} else if command == "update" {
		if len(input) == 1 {
			fmt.Println("[Error] Missing position argument")
		} else {
			if len(input) == 2 {
				fmt.Println("[Error] Missing note argument")
			} else {
				stringPosition := input[1]
				position, err := strconv.ParseInt(stringPosition, 10, 64)
				notePadLen := int64(len(*notepad))
				if err != nil {
					if strings.Trim(stringPosition, " ") == "" {
						fmt.Println("[Error] Missing position argument")
					} else {
						fmt.Printf("[Error] Invalid position: %s\n", stringPosition)
					}
				} else {
					if position < 1 || position > max {
						fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, max)
					} else {
						if position > notePadLen {
							fmt.Println("[Error] There is nothing to update")
						} else {
							subInput := strings.Trim(strings.Join(input[2:], " "), " ")
							if subInput == "" {
								fmt.Println("[Error] Missing note argument")
							} else {
								(*notepad)[position-1] = subInput
								fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
							}
						}
					}
				}
			}
		}
	} else if command == "delete" {
		if len(input) == 1 {
			fmt.Println("[Error] Missing position argument")
		} else {
			stringPosition := input[1]
			position, err := strconv.ParseInt(stringPosition, 10, 64)
			notePadLen := int64(len(*notepad))
			if err != nil {
				if strings.Trim(stringPosition, " ") == "" {
					fmt.Println("[Error] Missing position argument")
				} else {
					fmt.Printf("[Error] Invalid position: %s\n", stringPosition)
				}
			} else {
				if position < 1 || position > max {
					fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, max)
				} else {
					if position > notePadLen {
						fmt.Println("[Error] There is nothing to delete")
					} else {
						index := position - 1
						var newNotepad []string
						newNotepad = append(newNotepad, (*notepad)[:index]...)
						newNotepad = append(newNotepad, (*notepad)[index+1:]...)
						*notepad = newNotepad
						fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
					}
				}
			}
		}
	} else if command == "clear" {
		*notepad = (*notepad)[:0]
		fmt.Println("[OK] All notes were successfully deleted")
	} else if command == "list" {
		if len(*notepad) == 0 {
			fmt.Println("[Info] Notepad is empty")
		} else {
			for i, v := range *notepad {
				fmt.Printf("[Info] %d: %s\n", i+1, v)
			}
		}
	} else {
		fmt.Println("[Error] Unknown command")
	}

	run(notepad, max)
}

func readInput() []string {
	fmt.Println("Enter command and data:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return strings.Split(text, " ")
}
