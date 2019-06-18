package presist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
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

func save(item interface{}) {
	// Must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())

}
