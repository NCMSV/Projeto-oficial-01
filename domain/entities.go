package domain

type Person struct {
	Username string `json: "username"`
	Password string `json: "password"`
	Fullname string `json: "fullname"`
	ID       int    `json: "id"`
}

type People struct {
	People []Person `json: "people"`
}
