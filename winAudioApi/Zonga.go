package winAudioApi

import (
	"fmt"
	"time"
)


func Dada() {
	start := time.Now().UnixMilli()
	endpoint, err := GetDefaultEndpointSession()
	if err != nil {
		return
	}
	
	endpoint.SetVolume(100)
	endpoint.SetAppVolume("chrome.exe", 30)
	endpoint.SetAppVolume("discord.exe", 50)

	end := time.Now().UnixMilli()
	fmt.Println("time = ", end-start, "Milliseconds")
}
