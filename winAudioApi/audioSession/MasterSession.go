package audioSession

import (
	"ZAZA/winAudioApi/masterAudio"
	"ZAZA/winAudioApi/utils"
	"fmt"

	"github.com/moutend/go-wca/pkg/wca"
)

type MasterSession struct {
	AudioDevice    *wca.IMMDevice
	SessionManager *wca.IAudioSessionManager2
	SessionEnum    *wca.IAudioSessionEnumerator
}

func (ms *MasterSession) Close() {
	ms.SessionManager.Release()
	ms.SessionEnum.Release()
}

func NewSessionManager(endpoint *masterAudio.AudioEndpoint) (*wca.IAudioSessionManager2, error) {
	AudioDevice := endpoint.Device
	var audioSessionManager2 *wca.IAudioSessionManager2
	if err := AudioDevice.Activate(
		wca.IID_IAudioSessionManager2,
		wca.CLSCTX_ALL,
		nil,
		&audioSessionManager2,
	); err != nil {
		fmt.Println("could not create the session")
		return nil, err
	}

	return audioSessionManager2, nil
}

func NewSessionManagerEnum(sessionManager *wca.IAudioSessionManager2) (*wca.IAudioSessionEnumerator, error) {
	var sessionEnum *wca.IAudioSessionEnumerator

	if err := sessionManager.GetSessionEnumerator(&sessionEnum); err != nil {
		return nil, err
	}
	return sessionEnum, nil
}

func GetMasterSessionLength(sEnum *wca.IAudioSessionEnumerator) (int, error) {
	var sessionLength int

	if err := sEnum.GetCount(&sessionLength); err != nil {
		return -1, err
	}
	return sessionLength, nil
}

func GetAppsSessionEnumAndLength(endpoint *masterAudio.AudioEndpoint) (*wca.IAudioSessionEnumerator, int, error) {

	sManager, err := NewSessionManager(endpoint)
	if err != nil {
		return nil, -1, err
	}

	sessionEnum, err := NewSessionManagerEnum(sManager)
	if err != nil {
		return nil, -1, err
	}
	sLength, err := GetMasterSessionLength(sessionEnum)
	if err != nil {
		return nil, -1, err
	}
	return sessionEnum, sLength, nil
}

func GetAllAppsFromEndpoint(endpoint *masterAudio.AudioEndpoint) (*[]*AppAudio, error) {
	sessionEnum, sLength, err := GetAppsSessionEnumAndLength(endpoint)
	if err != nil {
		return nil, err
	}

	processes, err := utils.GetAllProcesses()
	if err != nil {
		return nil, err
	}
	allApps := []*AppAudio{}
	for sessionIdx := 0; sessionIdx < sLength; sessionIdx++ {
		sessionControl, err := GetSessionFromSessionEnum(sessionIdx, sessionEnum)
		if err != nil {
			return nil, err
		}
		app, err := NewAppAudio(sessionControl, processes)
		if err != nil {
			continue
		}
		allApps = append(allApps, app)
	}
	return &allApps, nil
}
