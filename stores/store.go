package stores

import (
	"github.com/oddnetworks/oddworks-go/stores/memory"
	"github.com/oddnetworks/oddworks-go/types"
)

type Store interface {
	Get(id string, channel types.Channel) (interface{}, error)
	Set(id string, channel types.Channel, resource interface{}) error
	Del(id string, channel types.Channel) error
}

var Stores map[string]Store = make(map[string]Store)

func init() {
	Stores[memory.NAME] = memory.New()
}
