package main

import (
	"fmt"

	"github.com/gen2brain/dlgs"
)

func main() {
	// err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	// if err != nil {
	// 	panic(err)
	// }
	// err := beeep.Notify("Title", "Message body", "assets/information.png")
	// if err != nil {
	// 	panic(err)
	// }
	// err := beeep.Alert("Title", "Message body", "assets/warning.png")
	// if err != nil {
	// 	panic(err)
	// }
	// yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(yes)
	// passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(passwd)
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}
