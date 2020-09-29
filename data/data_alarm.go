package data

//从系统提供的回调接口中获取的一整份的告警信息的结构体表示（JSON格式）
type Alarm struct {
	Id           int64
	Sid          int64
	Sname        string
	Node_path    string
	Nid          int64
	Endpoint     string
	Priority     int
	Event_type   string
	Category     int
	Status       uint16
	Hashid       uint64
	Etime        int64
	Value        string
	Info         string
	Last_updator string
	Created      string
	Groups       []string
	Users        []string
	Detail       []Details
}
type Details struct {
	Metric string
	Tags   tags
	Points []points
}
type tags struct {
	Jmxport string
}
type points struct {
	Timestamp uint64
	Value     int64
	Extra     string
}
