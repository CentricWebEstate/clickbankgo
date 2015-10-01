package clickbank

type EncryptedNotification struct {
	IV           string `json:"iv"`
	Notification string `json:"notification"`
}

type ClickbankLineItem struct {
	ItemNo        string  `json:"itemNo"`
	ProductTitle  string  `json:"productTitle"`
	Shippable     bool    `json:"shippable"`
	Recurring     bool    `json:"recurring"`
	AccountAmount float32 `json:"accountAmount"`
	Quantity      int     `json:"quantity"`
	DownloadUrl   string  `json:"downloadUrl"`
}

type ClickbankAddress struct {
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Country    string `json:"country"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	County     string `json:"county"`
}

type ClickbankContact struct {
	FirstName   string           `json:"firstName"`
	LastName    string           `json:"lastName"`
	FullName    string           `json:"fullName"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Address     ClickbankAddress `json:"address"`
}

type ClickbankUpsell struct {
	UpsellOriginalReceipt string `json:"upsellOriginalReceipt"`
	UpsellFlowId          int    `json:"upsellFlowId"`
	UpsellSession         string `json:"upsellSession"`
	UpsellPath            string `json:"upsellPath"`
}

type ClickbankHopfeed struct {
	HopfeedClickId           string  `json:"hopfeedClickId"`
	HopfeedApplicationId     int     `json:"hopfeedApplicationId"`
	HopfeedCreativeId        int     `json:"hopfeedCreativeId"`
	HopfeedApplicationPayout float32 `json:"hopfeedApplicationPayout"`
	HopfeedVendorPayout      float32 `json:"hopfeedVendorPayout"`
}

type ClickbankCustomer struct {
	Shipping ClickbankContact `json:"shipping"`
	Billing  ClickbankContact `json:"billing"`
}

type ClickbankNotification struct {
	TransactionTime     string                 `json:"transactionTime"`
	Receipt             string                 `json:"receipt"`
	TransactionType     string                 `json:"transactionType"`
	Vendor              string                 `json:"vendor"`
	Affiliate           string                 `json:"affiliate"`
	Role                string                 `json:"role"`
	TotalAccountAmount  float32                `json:"totalAccountAmount"`
	PaymentMethod       string                 `json:"paymentMethod"`
	TotalOrderAmount    float32                `json:"totalOrderAmount"`
	TotalTaxAmount      float32                `json:"totalTaxAmount"`
	TotalShippingAmount float32                `json:"totalShippingAmount"`
	Currency            string                 `json:"currency"`
	OrderLanguage       string                 `json:"orderLanguage"`
	TrackingCodes       []string               `json:"trackingCodes"`
	LineItems           []ClickbankLineItem    `json:"lineItems"`
	Customer            ClickbankCustomer      `json:"customer"`
	Upsell              ClickbankUpsell        `json:"upsell"`
	Hopfeed             ClickbankHopfeed       `json:"hopfeed"`
	Version             float32                `json:"version"`
	AttemptCount        int                    `json:"attemptCount"`
	VendorVariables     map[string]interface{} `json:"vendorVariables"`
}
