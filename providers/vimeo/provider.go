package vimeo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/oddnetworks/oddworks-go/types"
)

const NAME = "vimeo"

type Provider struct {
	client *http.Client
}

func formatID(uri string) {
	id := strings.Split(uri, "/")
	fmt.Println(id)
}

func New() *Provider {
	return &Provider{
		client: &http.Client{},
	}
}

func (p *Provider) GetCollections(channel types.Channel) []types.Collection {
	return []types.Collection{
		types.Collection{ID: NAME + "-1", Title: "One"},
		types.Collection{ID: NAME + "-2", Title: "Two"},
		types.Collection{ID: NAME + "-3", Title: "Three"},
		types.Collection{ID: NAME + "-4", Title: "Four"},
	}
}

func (p *Provider) GetCollection(id string, channel types.Channel) types.Collection {
	req, err := http.NewRequest("GET", "https://api.vimeo.com/me/albums/"+id+"?fields=uri,name,description,pictures,metadata.connections.videos.total", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+channel.Secrets[NAME])

	res, err := p.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json := string(body)

	return types.Collection{
		ID:    NAME + "-" + gjson.Get(json, "uri").String(),
		Title: "One",
	}
}

func (p *Provider) GetVideo(id string, channel types.Channel) types.Video {
	return types.Video{ID: NAME + "-1", Title: "One"}
}
