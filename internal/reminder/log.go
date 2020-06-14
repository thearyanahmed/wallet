package reminder

import (
	"fmt"
	"reflect"
)

const PRODUCTION = false

func Remind(reminders interface{}) {
	if PRODUCTION == false {
		switch reflect.TypeOf(reminders).Kind() {
		case reflect.Slice, reflect.Array, reflect.Int, reflect.String, reflect.Interface :
			s := reflect.ValueOf(reminders)
			for i := 0; i < s.Len(); i++ {
				fmt.Println(s.Index(i))
			}
		}
	}
}