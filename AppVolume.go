package winAudioApi

import (
	"fmt"

	"github.com/nasooory/go-winAudioApi/audioSession"
)

func FindApp(es *EndpointSession, appName string) ([]*audioSession.AppAudio, error) {
	retVal := []*audioSession.AppAudio{}
	if len(es.Apps) == 0 {
		return nil, fmt.Errorf("app array on endpoint session length is 0")
	}
	for _, app := range es.Apps {
		if app.AppName == appName {
			retVal = append(retVal, app)
		}
	}
	return retVal, nil
}

func (es *EndpointSession) SetAppVolume(name string, volume int) (bool, error) {
	apps, err := FindApp(es, name)
	if err != nil {
		return false, err
	}
	for _, app := range apps {
		_, err := app.SetVolume(volume)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (es *EndpointSession) GetAppVolume(name string) ([]int, error) {
	apps, err := FindApp(es, name)
	if err != nil {
		return nil, err
	}
	vols := []int{}
	for _, app := range apps {
		vol, err := app.GetVolume()
		if err != nil {
			return nil, err
		}
		vols = append(vols, vol)
	}
	return vols, nil
}

func (es *EndpointSession) SetAppMuteState(name string, state bool) (bool, error) {
	apps, err := FindApp(es, name)
	if err != nil {
		return false, err
	}
	for _, app := range apps {
		_, err := app.SetMuteState(state)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (es *EndpointSession) GetAppMuteState(name string) ([]bool, error) {
	apps, err := FindApp(es, name)
	if err != nil {
		return nil, err
	}
	vols := []bool{}
	for _, app := range apps {
		vol, err := app.GetMuteState()
		if err != nil {
			return nil, err
		}
		vols = append(vols, vol)
	}
	return vols, nil
}
