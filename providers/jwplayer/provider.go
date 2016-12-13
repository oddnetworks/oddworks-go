package jwplayer

import (
	"github.com/oddnetworks/oddworks-go/types"
)

const NAME = "jwplayer"

type Client struct{}

func New() *Client {
	return &Client{}
}

func (client *Client) GetCollections(channel types.Channel) []types.Collection {
	return []types.Collection{
		types.Collection{ID: NAME + "-1", Title: "One"},
		types.Collection{ID: NAME + "-2", Title: "Two"},
		types.Collection{ID: NAME + "-3", Title: "Three"},
		types.Collection{ID: NAME + "-4", Title: "Four"},
	}
}

func (client *Client) GetCollection(id string, channel types.Channel) types.Collection {
	return types.Collection{ID: NAME + "-1", Title: "One"}
}

func (client *Client) GetVideo(id string, channel types.Channel) types.Video {
	return types.Video{ID: NAME + "-1", Title: "One"}
}
