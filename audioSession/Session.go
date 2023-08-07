package audioSession

import (
	"unsafe"

	"github.com/moutend/go-wca/pkg/wca"
)

type AudioSession struct {
	SessionControl *wca.IAudioSessionControl2
	SimpleAudio    *wca.ISimpleAudioVolume
}


func NewAudioSession(sc *wca.IAudioSessionControl2) (*AudioSession, error) {
  sAudio, err := NewSimpleAudioVolume(sc)
  if err != nil {
	return nil, err
  }
  return &AudioSession{
	SessionControl: sc,
	SimpleAudio: sAudio,	
  }, nil
}

func (as *AudioSession) Close() {
	as.SessionControl.Release()
	as.SimpleAudio.Release()
}

func GetSessionFromSessionEnum(
	sessionIdx int,
	sessionEnum *wca.IAudioSessionEnumerator,
) (*wca.IAudioSessionControl2, error) {
	var audioSessionControl *wca.IAudioSessionControl
	if err := sessionEnum.GetSession(sessionIdx, &audioSessionControl); err != nil {
		return nil, err
	}
	dispatch, err := audioSessionControl.QueryInterface(wca.IID_IAudioSessionControl2)
	if err != nil {
		return nil, err
	}
	audioSessionControl2 := (*wca.IAudioSessionControl2)(unsafe.Pointer(dispatch))
	return audioSessionControl2, nil
}

func GetSessionPID(sessionControl *wca.IAudioSessionControl2) (*uint32, error) {
	var pid uint32
	if err := sessionControl.GetProcessId(&pid); err != nil {
		return nil, err
	}
	return &pid, nil
}

func NewSimpleAudioVolume(sessionControl *wca.IAudioSessionControl2) (*wca.ISimpleAudioVolume, error) {
	dispatch, err := sessionControl.QueryInterface(wca.IID_ISimpleAudioVolume)
	if err != nil {
		return nil, err
	}
	simpleAudioVolume := (*wca.ISimpleAudioVolume)(unsafe.Pointer(dispatch))
	return simpleAudioVolume, nil
}
