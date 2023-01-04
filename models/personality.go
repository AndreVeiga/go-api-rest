package models

type Personality struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var personalities []Personality

func GetAll() []Personality {
	var p1 = Personality{Id: 1, Nome: "Nome 1", Historia: "Historia 1"}
	personalities = append(personalities, p1)

	var p2 = Personality{Id: 2, Nome: "Nome 2", Historia: "Historia 2"}
	personalities = append(personalities, p2)

	return personalities
}
