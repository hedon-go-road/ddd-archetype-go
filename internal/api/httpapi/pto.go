package apihttp

type HelloReq struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

type HelloRsp struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
	Server  string `json:"server"`
}
