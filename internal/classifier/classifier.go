package classifier

import (
	"strings"
)

type Classifier struct {
}

func (c *Classifier) Classify(transactionTitle string) string {

	if transactionTitle == "" {
		return "unknown"
	}
	tags := map[string][]string{
		"food":          {"restaurant", "cafe", "dominos", "swiggy", "food", "zomato", "staples enterprise", "master harsh mitruka", "icecream"},
		"travel":        {"flight", "hotel", "uber", "ola", "rapido", "travel", "irctc", "makemytrip", "goibibo", "ontrot", "bus"},
		"clothes":       {"myntra", "zara", "levis", "gap", "espanshe"},
		"investment":    {"mutual fund", "stock", "investment", "sip", "groww", "zerodha"},
		"grocery":       {"bigbasket", "grofers", "grocery", "zepto", "hyper mart", "bakery", "supermarket", "dairy"},
		"entertainment": {"bookmyshow", "netflix", "prime", "hotstar", "spotify", "gaana", "music", "entertainment"},
		"shopping":      {"amazon", "flipkart", "snapdeal", "shopping", "shop", "store"},
		"selfcare":      {"salon", "spa", "selfcare", "gym", "fitness", "health", "medical"},
		"bills":         {"electricity", "water", "bills", "bill", "dth", "internet", "phone", "recharge", "gas"},
	}

	transactionTitleLower := transactionTitle

	for tag, keywords := range tags {
		for _, keyword := range keywords {
			if containsCaseInsensitive(transactionTitleLower, keyword) {
				return tag
			}
		}
	}

	return "miscellaneous"
}

func containsCaseInsensitive(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func NewClassifier() Classifier {
	return Classifier{}
}
