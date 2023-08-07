package utils

import (
	"github.com/shirou/gopsutil/v3/process"
)

func GetAllProcesses() (map[uint32]*process.Process, error) {
	allps := map[uint32]*process.Process{}
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, process := range processes {
		allps[uint32(process.Pid)] = process
	}
	return allps, nil
}
