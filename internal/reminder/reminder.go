package reminder

import "fmt"

const PRODUCTION = false

func Remind(reminder string) {
	if PRODUCTION == false {
		fmt.Println(reminder)
	}
}