package structs

type ShopOwner struct {
	ShopName  string `json:"shop_name"`
	OwnerName string `json:"owner_name"`
	RegDate   string `json:"reg_date"`
	PhNo      string `json:"ph_no"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
}

type PartnerDetails struct {
	FullName    string `json:"full_name"`
	CompanyName string `json:"company_name"`
	RegDate     string `json:"reg_date"`
	PhNo        string `json:"ph_no"`
	Address     string `json:"address"`
	IsActive    bool   `json:"is_active"`
}

type ShopDetails struct {
	ShopName string `json:"shop_name"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

type BalanceDetails struct {
	GoldDue   float64 `json:"gold_due"`
	SilverDue float64 `json:"silver_due"`
	CashDue   float64 `json:"cash_due"`
}

func NewBalanceDetails() BalanceDetails {
	return BalanceDetails{
		GoldDue:   0,
		SilverDue: 0,
		CashDue:   0,
	}
}
