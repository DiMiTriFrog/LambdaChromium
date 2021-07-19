package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
	"strings"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"	
	"github.com/go-rod/rod/lib/devices"
	"sync"
)



func randDevices() devices.Device{
	randgen := []devices.Device {devices.GalaxySIII,devices.GalaxyS5,devices.IPadMini,devices.IPad,devices.IPadPro,devices.Nexus10,devices.Nexus7,devices.GalaxyNote3,devices.GalaxyNoteII,devices.Nexus5X,devices.Nexus5,devices.Nexus4,devices.IPhoneX,devices.IPhone6or7or8Plus,devices.IPhone6or7or8,devices.IPhone5orSE}
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(randgen))
	choice := randgen[randIdx]
	return choice
}

func start_search()(string) {
	url := launcher.New().
	Proxy("yourproxy").
	//The lambda layers contents use /opt PATH 
	Bin("/opt/headless-chromium").
	Set("--headless").
	Set("--single-process").
	Set("--v=99").
	Set("--enable-webgl").
	Set("--disable-dev-shm-usage").
	Set("--ignore-gpu-blacklist").
	Set("--ignore-certificate-errors").
	Set("--allow-running-insecure-content").
	Set("-–disable-extensions").
	Set("--user-data-dir=/tmp/user-data").
	Set("--data-path=/tmp/data-path").
	Set("--disable-dev-shm-usage").
	Set("--homedir=/tmp").
	Set("--disk-cache-dir=/tmp/cache-dir").
	Set("--no-sandbox").
	Set("--use-gl=osmesa").
	Set("--window-size=312,512").
	MustLaunch()
	
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()
	browser.MustIgnoreCertErrors(true)
	page := browser.MustPage()
	page.MustEmulate(randDevices())
	
	//CODE of search start here...

	//CODE of search finish here

	new_url := page.MustInfo().URL
	return new_url
}

type Request struct {
	ID float64 `json:"id"`
	value string `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok bool `json:"ok"`
}

func Handler(request Request) (Response, error) {
	//Get the URL of last web loaded
	resp := start_search()
	return Response{
		Message: fmt.Sprintf(resp),
		Ok: true,
	}, nil
}

func main(){
	lambda.Start(Handler) 
}
