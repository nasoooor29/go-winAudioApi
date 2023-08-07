package masterAudio

import (
	"fmt"

	"github.com/moutend/go-wca/pkg/wca"
)

// this part for IMMDevice and creating blank IMMDevice object so Windows Core API can fill it with data

func NewAudioDeviceEndpoint() *wca.IMMDevice {
	var endpoint *wca.IMMDevice
	return endpoint
}

// This part is to get the properties of the audio device
// to demonstrate you can see them by going to "sound control panle" -> playback || recordings -> properties
// then you will see the same data ising the the property store of the IMMDevice.

func NewPropertyValue() *wca.PROPVARIANT {
	return &wca.PROPVARIANT{}
}

func (adp *AudioEndpoint) GetAudioDeviceName() (string, error) {
	val, err := adp.GetPropertyValue(&wca.PKEY_Device_FriendlyName)
	if err != nil {
		return "", err
	}
	return val, nil
}

// &wca.PKEY_DeviceInterface_FriendlyName
// &wca.PKEY_Device_DeviceDesc
// &wca.PKEY_Device_FriendlyName
// &wca.PKEY_AudioEndpoint_Association
// &wca.PKEY_AudioEndpoint_GUID
// &wca.PKEY_AudioEndpoint_JackSubType

func (adp *AudioEndpoint) GetPropertyValue(key *wca.PROPERTYKEY) (string, error) {
	val := NewPropertyValue()
	if err := adp.PropertyStore.GetValue(key, val); err != nil {
		return "", err
	}
	stringVal := val.String()
	return stringVal, nil
}

func GetAudioProperties(IMMDevice *wca.IMMDevice) (*wca.IPropertyStore, error) {
	var ps *wca.IPropertyStore

	if err := IMMDevice.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		fmt.Printf("faild to get the property storage of the audio device []")
		return nil, err
	}

	return ps, nil
}

// this part of the file is the methods that control Audio Device Volume endpoint
// using the *wca.IAudioEndpointVolume

func GetAudioEndpointVolume(IMMDevice *wca.IMMDevice) (*wca.IAudioEndpointVolume, error) {

	var aev *wca.IAudioEndpointVolume
	if err := IMMDevice.Activate(wca.IID_IAudioEndpointVolume, wca.CLSCTX_ALL, nil, &aev); err != nil {
		fmt.Printf("faild to get endpoint volume object of audio device []")
		return nil, err
	}

	return aev, nil
}

func GetAudioEndpointVolAndProps(Device *wca.IMMDevice) (*AudioEndpoint, error) {
	props, err := GetAudioProperties(Device)
	if err != nil {
		return nil, err
	}
	vol, err := GetAudioEndpointVolume(Device)
	if err != nil {
		return nil, err
	}
	ae := AudioEndpoint{
		PropertyStore: props,
		Volume:        vol,
		Device:        Device,
	}
	name, err := ae.GetAudioDeviceName()
	if err != nil {
		fmt.Println("could not get the device name")
		return nil, err
	}
	ae.Name = name
	return &ae, nil
}
