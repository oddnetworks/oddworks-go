package providers

import (
	"github.com/oddnetworks/oddworks-go/providers/jwplayer"
	"github.com/oddnetworks/oddworks-go/providers/vimeo"
	"github.com/oddnetworks/oddworks-go/types"
)

type Provider interface {
	GetCollections(channel types.Channel) []types.Collection
	GetCollection(id string, channel types.Channel) types.Collection
	GetVideo(id string, channel types.Channel) types.Video
}

var Providers map[string]Provider = make(map[string]Provider)

func init() {
	Providers[vimeo.NAME] = vimeo.New()
	Providers[jwplayer.NAME] = jwplayer.New()
}
