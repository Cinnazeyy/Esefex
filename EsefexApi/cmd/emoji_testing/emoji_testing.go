package main

import (
	"esefexapi/sounddb"
	"fmt"
)

func main() {
	text := "<:emoji:630819109726191617>🀄🆘🧌🤡🆘"

	icon, err := sounddb.ExtractIcon(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(icon)
}
