package result

type Resp struct {
	Code  int         `json:"code"`  // 错误码
	Msg   string      `json:"msg"`   // 错误描述
	Data  interface{} `json:"data"`  // 返回数据
	Error interface{} `json:"error"` //参数错误
	Time  int         `json:"-"`     //浏览器缓存时间 秒
}

func response(code int, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (this *Resp) WithTime(time int) *Resp {
	this.Time = time
	return this
}

func (this *Resp) WithData(data interface{}) *Resp {
	return &Resp{
		Code: this.Code,
		Msg:  this.Msg,
		Data: data,
	}
}

func (this *Resp) WithParamError(err interface{}) *Resp {
	this.Error = err
	return this
}
