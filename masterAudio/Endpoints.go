package masterAudio

import (
	"fmt"

	"github.com/moutend/go-wca/pkg/wca"
)

func GetDefaultAudioDevice() (*AudioEndpoint, error) {
	endpoint := NewAudioDeviceEndpoint()

	enum, err := GetAudioEndpointEnum()
	if err != nil {
		return nil, err
	}
	defer enum.Release()
	err = enum.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &endpoint)
	if err != nil {
		return nil, err
	}
	obj, err := GetAudioEndpointVolAndProps(endpoint)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetAllEndpoints() (*[]*AudioEndpoint, error) {

	coll, err := GetDeviceCollection(wca.DEVICE_STATE_ACTIVE)
	if err != nil {
		fmt.Println("could not get audio devices")
		return nil, err
	}
	num, err := GetDeviceCollectionLength(coll)
	if err != nil {
		fmt.Println("could not get the number of audio devices")
		return nil, err
	}

	endpoitns := []*AudioEndpoint{}
	for deviceIdx := uint32(0); deviceIdx < *num; deviceIdx++ {
		endpoint, err := GetAudioDeviceFromDeviceCollection(coll, deviceIdx)
		if err != nil {
			fmt.Printf("err: %v | and could not get device of index %v\n", err, deviceIdx)
			continue
		}
		endpoitns = append(endpoitns, endpoint)
	}
	def, err := GetDefaultAudioDevice()
	if err != nil {
		fmt.Println("could not get the default audio device")
		return nil, err
	}
	defName := def.Name

	for _, v := range endpoitns {
		name := v.Name
		if name == defName {
			v.IsDefault = true
		}

	}
	return &endpoitns, nil
}
