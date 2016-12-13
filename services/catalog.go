package services

import (
	"github.com/oddnetworks/oddworks-go/providers"
	"github.com/oddnetworks/oddworks-go/stores"
	"github.com/oddnetworks/oddworks-go/types"
)

type CatalogService struct {
	Store *stores.Store
}

func (service *CatalogService) GetCollection(id string, channel types.Channel) types.Collection {
	spec := service.GetCollectionSpec(id, channel)
	return providers.Providers[spec.Provider].GetCollection(id, channel)
}

func (service *CatalogService) GetCollectionSpec(id string, channel types.Channel) types.CollectionSpec {
	return types.CollectionSpec{
		ID:       id,
		Provider: "vimeo",
	}
}

func (service *CatalogService) GetVideo(id string, channel types.Channel) types.Video {
	spec := service.GetVideoSpec(id, channel)
	return providers.Providers[spec.Provider].GetVideo(id, channel)
}

func (service *CatalogService) GetVideoSpec(id string, channel types.Channel) types.VideoSpec {
	return types.VideoSpec{
		ID:       id,
		Provider: "vimeo",
	}
}
