package audioSession

import (
	"fmt"

	"github.com/moutend/go-wca/pkg/wca"
	"github.com/shirou/gopsutil/v3/process"
)

type AppAudio struct {
	AppName   string
	ProcessId uint32
	Ps        *process.Process

	Session *AudioSession
}

func (a *AppAudio) GetVolume() (int, error) {
	var vol float32
	if err := a.Session.SimpleAudio.GetMasterVolume(&vol); err != nil {
		return -1, nil
	}
	return int(vol * 100), nil
}

func (a *AppAudio) SetVolume(volume int) (bool, error) {
	floatVol := float32(volume) / 100
	if err := a.Session.SimpleAudio.SetMasterVolume(floatVol, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (a *AppAudio) GetMuteState() (bool, error) {
	var state bool
	if err := a.Session.SimpleAudio.GetMute(&state); err != nil {
		return false, err
	}
	return state, nil
}

func (a *AppAudio) SetMuteState(state bool) (bool, error) {
	if err := a.Session.SimpleAudio.SetMute(state, nil); err != nil {
		return false, err
	}
	return true, nil
}

func NewAppAudio(sc *wca.IAudioSessionControl2, processes map[uint32]*process.Process) (*AppAudio, error) {
	pid, err := GetSessionPID(sc)
	if err != nil {
		return nil, err
	}
	if *pid == 0 {
		return nil, nil
	}
	as, err := NewAudioSession(sc)
	if err != nil {
		return nil, err
	}
	ps, found := processes[*pid]
	if !found {
		return nil, fmt.Errorf("could not find the process id from the given processes")
	}
	psName, err := ps.Name()
	if err != nil {
		return nil, err
	}

	return &AppAudio{
		AppName:   psName,
		ProcessId: *pid,
		Ps:        ps,
		Session:   as,
	}, nil
}

func (a *AppAudio) Close() {
	a.Session.Close()
}
