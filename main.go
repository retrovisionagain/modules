package main

import (
	"fmt"
	"smartgopher/modules"
	"smartgopher/storage/schedule"
)

func main() {
	var date_now string
	fmt.Println("hey yo, it", modules.Check_now(date_now))
	schedule.Schedule_to_json()
}
