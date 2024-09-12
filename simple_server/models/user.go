package models

type User struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ProfilePicUrl string `json:"profile_pic_url"`
}
