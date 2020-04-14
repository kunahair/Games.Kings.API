package memory

type KingsCard struct {
	Card     Card   `json:"card"`
	Rule     string `json:"rule"`
	Holdable bool   `json:"holdable"`
}

func InitKingsCard(card Card, rules RulesTopLevel) KingsCard {
	kingsCard := KingsCard{
		Card: card,
	}

	for _, val := range rules.TopLevel {
		if val.Code == card.Code {
			kingsCard.Rule = val.Rule
			kingsCard.Holdable = val.Holdable
			break
		}
	}

	return kingsCard
}
