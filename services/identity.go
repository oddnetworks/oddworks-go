package services

import (
	"github.com/oddnetworks/oddworks-go/types"
)

type IdentityService struct{}

func (service *IdentityService) GetChannel(id string) types.Channel {
	return types.Channel{}
}

func (service *IdentityService) GetPlatform(id string, channel types.Channel) types.Platform {
	return types.Platform{}
}
