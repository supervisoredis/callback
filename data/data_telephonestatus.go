package data

//回调端返回的状态碼
type Telephonestatus struct {
	Status string //状态
	Msg    string //信息
	Result ALARM
}
type ALARM struct {
	Endpoint   string
	Sname      string
	Event_type string
	Phone      string
}
