package entity

type AppInfo struct {
	AppId   string `json:"appId"`
	AppName string `json:"appName"`
}

func CreateAppInfo() []AppInfo {
	appInfos := []AppInfo{
		AppInfo{AppId: "123", AppName: "Spotify"},
		AppInfo{AppId: "234", AppName: "Jio"},
	}
	return appInfos
}
