package v1

type JsonRequest struct {
	AppName    string `json:"appName"`
    Creator    string `json:"creator"`
	RobotKeys  string `json:"robotKeys"`
}

type FormRequest struct {
	AppName    string `form:"appName"`
	Creator    string `form:"creator"`
	RobotKeys  string `form:"robotKeys"`
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

// Ops get 数据响应结构体

type AppInfo struct {
	YingyongName        string `json:"yingyong_name"`
	YingyongFabudingyue string `json:"yingyong_fabudingyue"`
	LastModifyTime int `json:"modify_at"`
	Id string `json:"_id"`
}

type GetOpsData struct {
	List  []AppInfo `json:"list"`
	Total int   `json:"total"`
}

type GetOpsResponse struct {
	Code int        `json:"code"`
	Data GetOpsData `json:"data"`
	Msg  string     `json:"msg"`
}


// Ops 添加数据请求结构体

type AddOpsData struct {
	YingyongIdentify        string   `json:"yingyong_identify"`
	YingyongName            string   `json:"yingyong_name"`
	YingyongCeShiK8sjiqun string `json:"yingyong_ceshik8sjiqun"`
	YingYongCangKuDizhi   string `json:"yingyong_cangkudizhi"`
	YingYongBuShuK8s      string `json:"yingyong_bushuk8s"`
	YingYongJinSiQueLeiXing string   `json:"yingyong_jinsiqueleixing"`
	YingYongFaBuDingYue     string   `json:"yingyong_fabudingyue"`
}


type AddOpsRequest struct {
	ModelIdentify string  `json:"modelIdentify"`
	Data AddOpsData `json:"data"`
	LinkData   []string `json:"LinkData"`
}


// Ops 修改数据请求结构体


type UpdateOpsData struct {
	YingyongIdentify       string `json:"yingyong_identify"`
	YingyongName           string `json:"yingyong_name"`
	YingyongCeShiK8sjiqun string `json:"yingyong_ceshik8sjiqun"`
	YingyongCangKuDizhi    string `json:"yingyong_cangkudizhi"`
	YingyongBuShuK8s       string `json:"yingyong_bushuk8s"`
	YingyongJinSiQueLeixing string `json:"yingyong_jinsiqueleixing"`
	YingYongFaBuDingYue    string `json:"yingyong_fabudingyue"`
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
