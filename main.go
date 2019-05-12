package main

import (
	"fmt"
	"hn-reader/hn"
)

func main() {

	list, _ := hn.TopItems()
	for _, v := range list[:10] {
		item, _ := hn.GetItem(v)
		fmt.Println(hn.ItemFormat(item))
	}

}
