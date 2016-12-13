package memory

import "github.com/oddnetworks/oddworks-go/types"

const NAME = "vimeo"

type Store struct {
	db map[string]interface{}
}

func New() *Store {
	return &Store{
		db: make(map[string]interface{}),
	}
}

func (store *Store) Get(id string, channel types.Channel) (interface{}, error) {
	return store.db[channel.ID+":"+id], nil
}

func (store *Store) Set(id string, channel types.Channel, resource interface{}) error {
	store.db[channel.ID+":"+id] = resource
	return nil
}

func (store *Store) Del(id string, channel types.Channel) error {
	delete(store.db, channel.ID+":"+id)
	return nil
}
