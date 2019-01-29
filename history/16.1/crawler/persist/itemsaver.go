package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	// 创建一个可以传递任意类型的 channel
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}
