package main

var validTags = map[string]bool{
	"Assets":                                true,
	"AssetsCurrent":                         true,
	"AssetsNoncurrent":                      true,
	"NoncurrentAssets":                      true,
	"AccountsPayableCurrent":                true,
	"AccountsReceivableNet":                 true,
	"AccountsReceivableNetCurrent":          true,
	"AdvertisingExpense":                    true,
	"Cash":                                  true,
	"CashAndCashEquivalentsAtCarryingValue": true,
	"CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents": true,
	"CostOfRevenue":                                       true,
	"CostOfGoodsAndServicesSold":                          true,
	"GrossProfit":                                         true,
	"ResearchAndDevelopmentExpense":                       true,
	"SellingGeneralAndAdministrativeExpense":              true,
	"SellingAndMarketingExpense":                          true,
	"OperatingExpenses":                                   true,
	"OperatingIncomeLoss":                                 true,
	"NonoperatingIncomeExpense":                           true,
	"IncomeTaxExpenseBenefit":                             true,
	"NetIncomeLoss":                                       true,
	"NetIncomeLossAttributableToNoncontrollingInterest":   true,
	"EarningsPerShareBasic":                               true,
	"EarningsPerShareDiluted":                             true,
	"WeightedAverageNumberOfSharesOutstandingBasic":       true,
	"WeightedAverageNumberOfDilutedSharesOutstanding":     true,
	"MarketableSecuritiesCurrent":                         true,
	"InventoryNet":                                        true,
	"MarketableSecuritiesNoncurrent":                      true,
	"PropertyPlantAndEquipmentNet":                        true,
	"PropertyPlantAndEquipmentGross":                      true,
	"OtherAssetsNoncurrent":                               true,
	"LiabilitiesCurrent":                                  true,
	"OtherLiabilitiesCurrent":                             true,
	"ContractWithCustomerLiabilityCurrent":                true,
	"LongTermDebtCurrent":                                 true,
	"LongTermDebtNoncurrent":                              true,
	"OtherLiabilitiesNoncurrent":                          true,
	"LiabilitiesNoncurrent":                               true,
	"Liabilities":                                         true,
	"StockholdersEquity":                                  true,
	"DepreciationDepletionAndAmortization":                true,
	"NetCashProvidedByUsedInOperatingActivities":          true,
	"NetCashProvidedByUsedInInvestingActivities":          true,
	"NetCashProvidedByUsedInFinancingActivities":          true,
	"EntityCommonStockSharesOutstanding":                  true,
	"NontradeReceivablesCurrent":                          true,
	"PaymentsForRepurchaseOfCommonStock":                  true,
	"PaymentsOfDividends":                                 true,
	"RepaymentsOfLongTermDebt":                            true,
	"RevenueFromContractWithCustomerExcludingAssessedTax": true,
	"Revenues":             true,
	"Goodwill":             true,
	"ShortTermInvestments": true,
}

type stockPrice struct {
	Ticker       string  `json:"ticker"`
	Date         string  `json:"date"`
	Volume       int64   `json:"volume"`
	Open         float64 `json:"open"`
	Close        float64 `json:"close"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Transactions int64   `json:"transactions"`
}

type company struct {
	Cik    int64  `json:"cik"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}

type QueryParamDate struct {
	Ticker     string `json:"ticker"`
	End_date   string `json:"end_date"`
	Start_date string `json:"start_date"`
}

type condensedFacts struct {
	Tag   string  `json:"tag"`
	Date  string  `json:"date"`
	Value float64 `json:"value"`
	Qtrs  int8    `json:"qtrs"`
}

type articleInfo struct {
	PublisherName string `json:"publisher_name"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	ArticleURL    string `json:"article_url"`
}

type factFilingParam struct {
	dateParam QueryParamDate
	tag       string
	yearly    bool
}

func (q QueryParamDate) getTicker() string {
	return q.Ticker
}

type basePageResponse struct {
	Facts       []condensedFacts `json:"facts"`
	Articles    []articleInfo    `json:"articles"`
	StockPrices []stockPrice     `json:"stock_prices"`
}
