package persist

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.4/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.4/crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
