package memory

type Player struct {
	Name	string		`json:"name"`
	Cards	[]KingsCard	`json:"cards"`
}
