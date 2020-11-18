package data

type Group struct {
	Dat Dat_group
	Err string
}

type Dat_group struct {
	Group_list []Group_list
	Total      int
}

type Group_list struct {
	Id          int
	Ident       string
	Name        string
	Mgmt        int
	Admin_objs  []Admin_objs
	Member_objs []Member_objs
}

type Member_objs struct {
	Id       int
	Username string
	Disname  string
	Phone    string
	Email    string
	Im       string
	Is_root  int
}

type Admin_objs struct {
	Id       int
	Username string
	Disname  string
	Phone    string
	Email    string
	Im       string
	Is_root  int
}
