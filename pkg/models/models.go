package models

type User struct {
	ID       int
	Name     string
	Login    string
	Password string
	LastCig  int
}

type Group struct {
	ID           int
	UsersList    []int
	IntervalTime int
}