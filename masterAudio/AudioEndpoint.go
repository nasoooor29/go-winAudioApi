package masterAudio

import (
	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

type AudioEndpoint struct {
	Name          string
	IsDefault     bool
	PropertyStore *wca.IPropertyStore
	Volume        *wca.IAudioEndpointVolume
	Device        *wca.IMMDevice
	Enum          *wca.IMMDeviceEnumerator
}

func (ae AudioEndpoint) GetMuteState() (bool, error) {
	var state bool
	if err := ae.Volume.GetMute(&state); err != nil {
		return false, err
	}
	return state, nil
}

func (ae AudioEndpoint) SetMuteState(state bool) (bool, error) {

	if err := ae.Volume.SetMute(state, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (ae AudioEndpoint) GetVolume() (int, error) {
	var vol float32
	if err := ae.Volume.GetMasterVolumeLevelScalar(&vol); err != nil {
		return -1, err
	}
	return int(vol * 100), nil
}
func (ae AudioEndpoint) SetVolume(volume int) (bool, error) {
	floatVol := float32(volume) / 100
	if err := ae.Volume.SetMasterVolumeLevelScalar(floatVol, nil); err != nil {
		return false, err
	}
	return true, nil
}

func CloseEndpoint(ae AudioEndpoint) (bool, error) {
	ole.CoUninitialize()
	ae.Device.Release()
	ae.PropertyStore.Release()
	ae.Volume.Release()
	return true, nil
}
