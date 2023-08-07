package masterAudio

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

// this part to get IMMDevice enum and create device collection and get the length of the collection





func GetAudioEndpointEnum() (*wca.IMMDeviceEnumerator, error) {
	ole.CoUninitialize()
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		fmt.Println("could not initilize")
		return nil, fmt.Errorf("could not initialize error is: [%v]", err)
	}

	var mmde *wca.IMMDeviceEnumerator
	if err := wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		fmt.Println("could not create instance")
		return nil, fmt.Errorf("could not create the enum error is: [%v]", err)
	}


	return mmde, nil
}

func GetDeviceCollectionLength(deviceCollection *wca.IMMDeviceCollection) (*uint32, error) {
	var deviceCount uint32
	err := deviceCollection.GetCount(&deviceCount)
	if err != nil {
		fmt.Println("Could not count the number of devices")
		return nil, err
	}
	return &deviceCount, nil
}

func GetDeviceCollection(stateMask uint32) (*wca.IMMDeviceCollection, error) {
	var deviceCollection *wca.IMMDeviceCollection
	enum, err := GetAudioEndpointEnum()
	if err != nil {
		return nil, err
	}
	defer enum.Release()

	if err := enum.EnumAudioEndpoints(wca.EAll, stateMask, &deviceCollection); err != nil {
		return nil, err
	}

	return deviceCollection, nil
}

func GetAudioDeviceFromDeviceCollection(deviceCollection *wca.IMMDeviceCollection, deviceIdx uint32) (*AudioEndpoint, error) {
	endpoint := NewAudioDeviceEndpoint()
	if err := deviceCollection.Item(deviceIdx, &endpoint); err != nil {
		fmt.Printf("faild to get device of index [%v], %v", deviceIdx, err)
		return nil, err
	}
	obj, err := GetAudioEndpointVolAndProps(endpoint)
	if err != nil {
		return nil, err
	}
	return obj, nil

}
