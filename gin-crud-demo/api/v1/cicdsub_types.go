package v1

type JsonRequest struct {
	AppName    string `json:"appName"`
        Creator    string `json:"creator"`
	RobotKeys  string `json:"robotKeys"`
}

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data  interface{} `json:"data"`
}


type JsonGetOpsRequest struct {
	AppName  string `json:"appName"`
	Page     string `json:"page"`
	PageNum  string `json:"pageNum"`
}



type RequestOpsParams struct {
	Method    string `req:"method"`
	Identify  string `req:"identify"`
	KeyWord   string `req:"keyWord"`
	Page      int    `req:"page"`
	PageNum   int    `req:"pageNum"`
}

type dataOpsResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type UpdateOpsRequest struct {
	ModelIdentify string `json:"modelIdentify"`
	ID            string `json:"id"`
	Data          UpdateOpsData   `json:"data"`
	LinkData      []string `json:"LinkData"`
}


type JsonUpdateRequestId struct {
	AppName    string `json:"appName"`
	Id    int64 `json:"id"`
	ModifyBy string `json:"modifyBy"`
	RobotKeys  string `json:"robotKeys"`
}

type GetSubData struct {
	AppName        string `json:"appName"`
	RobotKeys string `json:"robotKeys"`
	LastModifyBy string `json:"lastModifyBy"`
	CreateTime string `json:"create_time"`
	LastModifyTime string `json:"lastModifyTime"`
	ID int64 `json:"id"`
}
type GetCicdSubListResponse struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List  []GetSubData `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}


type JsonDeleteRequestId struct {
	AppName    string `json:"appName"`
	Id    int64 `json:"id"`
}
