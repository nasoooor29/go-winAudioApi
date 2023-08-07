package winAudioApi

import (
	"ZAZA/winAudioApi/audioSession"
	"ZAZA/winAudioApi/masterAudio"
	"fmt"
)

type EndpointSession struct {
	*masterAudio.AudioEndpoint
	Apps []*audioSession.AppAudio
}

func NewEndpointSession(endpoint *masterAudio.AudioEndpoint) (*EndpointSession, error) {
	apps, err := audioSession.GetAllAppsFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	es := EndpointSession{
		AudioEndpoint: endpoint,
		Apps:          *apps,
	}
	return &es, nil
}

func GetAllEndpointSessions() ([]*EndpointSession, error) {
	endpoints, err := masterAudio.GetAllEndpoints()
	if err != nil {
		fmt.Printf("could not get the endpoints err: %v\n", err)
		return nil, err
	}
	allEndpointSessions := []*EndpointSession{}
	for _, endpoint := range *endpoints {
		es, err := NewEndpointSession(endpoint)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		allEndpointSessions = append(allEndpointSessions, es)
	}
	return allEndpointSessions, nil
}

func GetEndpointSessionByName(name string) (*EndpointSession, error) {
	endpoints, err := GetAllEndpointSessions()
	if err != nil {
		return nil, err
	}
	for _, e := range endpoints {
		if name == e.Name {
			return e, nil
		}
	}
	return nil, fmt.Errorf("could not find endpoint session %v", name)
}

func GetDefaultEndpointSession() (*EndpointSession, error) {
	device, err := masterAudio.GetDefaultAudioDevice()
	if err != nil {
		return nil, err
	}
	endpoint, err := NewEndpointSession(device)
	if err != nil {
		return nil, err
	}
	return endpoint, err
}
