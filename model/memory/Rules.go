package memory

type RulesTopLevel struct {
	TopLevel	[]RulesV2	`json:"rules"`
}

type RulesV2 struct {
	Code	string	`json:"code"`
	Rule	string	`json:"rule"`
}
