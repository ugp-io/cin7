package cin7

import (
	"fmt"
	"time"
)

var url = `https://inventory.dearsystems.com/externalapi/v2/`

type ISO8601Time struct {
	time.Time
}

var iso8601Layouts = []string{
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05.000Z",
}

func (ct *ISO8601Time) UnmarshalJSON(b []byte) error {

	s := string(b)
	s = s[1 : len(s)-1]

	var err error
	for _, layout := range iso8601Layouts {
		var t time.Time
		t, err = time.Parse(layout, s)
		if err == nil {
			ct.Time = t
			return nil
		}
	}

	return fmt.Errorf("unable to parse date: %v", err)
}

//////////////////////////////////////////
/*			CREATE REQUESTS				*/
//////////////////////////////////////////

type CreateProduct struct {
	SKU                          string                         `json:"SKU"`
	Name                         string                         `json:"Name"` // maximum length is 100 characters
	Category                     string                         `json:"Category"`
	Type                         string                         `json:"Type"`
	CostingMethod                string                         `json:"CostingMethod"`                // Valid values: FIFO, Special - Batch, Special - Serial Number, FIFO - Serial Number, FIFO - Batch, FEFO - Batch, FEFO - Serial Number
	UOM                          string                         `json:"UOM"`                          // A unit of measure with this name must exist in reference books under units
	QuantityToProduce            float64                        `json:"QuantityToProduce"`            // Required if BillOfMaterial is true
	AssemblyCostEstimationMethod string                         `json:"AssemblyCostEstimationMethod"` // Required if BillOfMaterial is true
	Status                       *string                        `json:"Status"`                       // Valid values: Active and Deprecated
	Brand                        *string                        `json:"Brand"`
	DropShipMode                 *string                        `json:"DropShipMode"`    // Valid values: No Drop Ship, Optional Drop Ship, Always Drop Ship. Default value is No Drop Ship
	DefaultLocation              *string                        `json:"DefaultLocation"` // Location with this name must exist in reference books under locations
	Length                       *float64                       `json:"Length"`
	Width                        *float64                       `json:"Width"`
	Height                       *float64                       `json:"Height"`
	Weight                       *float64                       `json:"Weight"`
	CartonLength                 *float64                       `json:"CartonLength"`
	CartonWidth                  *float64                       `json:"CartonWidth"`
	CartonHeight                 *float64                       `json:"CartonHeight"`
	CartonQuantity               *float64                       `json:"CartonQuantity"`
	CartonInnerQuantity          *float64                       `json:"CartonInnerQuantity"`
	WeightUnits                  *string                        `json:"WeightUnits"` // Valid values: g, kg, oz, lb
	DimensionsUnits              *string                        `json:"DimensionsUnits"`
	Barcode                      *string                        `json:"Barcode"`
	MinimumBeforeReorder         *float64                       `json:"MinimumBeforeReorder"`
	ReorderQuantity              *float64                       `json:"ReorderQuantity"`
	PriceTier1                   *float64                       `json:"PriceTier1"`
	PriceTier2                   *float64                       `json:"PriceTier2"`
	PriceTier3                   *float64                       `json:"PriceTier3"`
	PriceTier4                   *float64                       `json:"PriceTier4"`
	PriceTier5                   *float64                       `json:"PriceTier5"`
	PriceTier6                   *float64                       `json:"PriceTier6"`
	PriceTier7                   *float64                       `json:"PriceTier7"`
	PriceTier8                   *float64                       `json:"PriceTier8"`
	PriceTier9                   *float64                       `json:"PriceTier9"`
	PriceTier10                  *float64                       `json:"PriceTier10"`
	ShortDescription             *string                        `json:"ShortDescription"`
	Description                  *string                        `json:"Description"`
	InternalNote                 *string                        `json:"InternalNote"`
	AdditionalAttribute1         *string                        `json:"AdditionalAttribute1"`
	AdditionalAttribute2         *string                        `json:"AdditionalAttribute2"`
	AdditionalAttribute3         *string                        `json:"AdditionalAttribute3"`
	AdditionalAttribute4         *string                        `json:"AdditionalAttribute4"`
	AdditionalAttribute5         *string                        `json:"AdditionalAttribute5"`
	AdditionalAttribute6         *string                        `json:"AdditionalAttribute6"`
	AdditionalAttribute7         *string                        `json:"AdditionalAttribute7"`
	AdditionalAttribute8         *string                        `json:"AdditionalAttribute8"`
	AdditionalAttribute9         *string                        `json:"AdditionalAttribute9"`
	AdditionalAttribute10        *string                        `json:"AdditionalAttribute10"`
	AttributeSet                 *string                        `json:"AttributeSet"`
	DiscountRule                 *string                        `json:"DiscountRule"` // discount with this name must exist in Product Discounts reference book and should be active
	Tags                         *string                        `json:"Tags"`         // Comma delimited list
	StockLocator                 *string                        `json:"StockLocator"`
	COGSAccount                  *string                        `json:"COGSAccount"`      // Account code of an active EXPENSE class account from Chart Of Accounts
	RevenueAccount               *string                        `json:"RevenueAccount"`   // Account code of an active REVENUE class account from Chart Of Accounts
	ExpenseAccount               *string                        `json:"ExpenseAccount"`   // Account code of an active EXPENSE class account from Chart Of Accounts
	InventoryAccount             *string                        `json:"InventoryAccount"` // Account code of an active ASSET class account from Chart Of Accounts with Type not equal to FIXED or BANK
	PurchaseTaxRule              *string                        `json:"PurchaseTaxRule"`  // Tax Rule with this name must exist in Taxation Rules reference book and should be Active and have For Purchases flag set to “true” in API
	SaleTaxRule                  *string                        `json:"SaleTaxRule"`      // Tax Rule with this name must exist in Taxation Rules reference book and should be Active and have For Sales flag set to “true” in API
	Sellable                     *bool                          `json:"Sellable"`         // Default value is true
	PickZones                    *string                        `json:"PickZones"`
	BillOfMaterial               *bool                          `json:"BillOfMaterial"`
	AutoAssembly                 *bool                          `json:"AutoAssembly"`
	AutoDisassembly              *bool                          `json:"AutoDisassembly"`
	AssemblyInstructionURL       *string                        `json:"AssemblyInstructionURL"`
	HSCode                       *string                        `json:"HSCode"`
	CountryOfOrigin              *string                        `json:"CountryOfOrigin"`
	CountryOfOriginCode          *string                        `json:"CountryOfOriginCode"`
	Suppliers                    []ProductSupplier              `json:"Suppliers"`
	ReorderLevels                []ProductReorderLevel          `json:"ReorderLevels"`
	BillOfMaterialsProducts      []ProductBillOfMaterialProduct `json:"BillOfMaterialsProducts"`
	BillOfMaterialsServices      []ProductBillOfMaterialService `json:"BillOfMaterialsServices"`
	Movements                    []ProductMovement              `json:"Movements"`
	Attachments                  []AttachmentLine               `json:"Attachments"`
	CustomPrices                 []CustomPrice                  `json:"CustomPrices"`
	// PriceTiers
}

type CreatePurchase struct {
	Approach             string                   `json:"Approach"` // INVOICE for Invoice First, or STOCK for Stock First
	Location             string                   `json:"Location"`
	ID                   *string                  `json:"ID,omitempty,omitempty"` // Required for PUT`
	SupplierID           *string                  `json:"SupplierID,omitempty"`   // Required if Supplier not provided
	Supplier             *string                  `json:"Supplier,omitempty"`     // Required if SupplierID not provided
	Contact              *string                  `json:"Contact,omitempty"`
	Phone                *string                  `json:"Phone,omitempty"`
	BlindReceipt         *bool                    `json:"BlindReceipt,omitempty"` // True if there is no order in the purchase. False is default
	TaxRule              *string                  `json:"TaxRule,omitempty"`
	TaxCalculation       *string                  `json:"TaxCalculation,omitempty"` // Inclusive or Exclusive
	Terms                *string                  `json:"Terms,omitempty"`
	RequiredBy           *string                  `json:"RequiredBy,omitempty"`
	Note                 *string                  `json:"Note,omitempty"`
	CurrencyRate         *float64                 `json:"CurrencyRate,omitempty"`
	AdditionalAttributes map[string]string        `json:"AdditionalAttributes,omitempty"`
	BillingAddress       *Address                 `json:"BillingAddress,omitempty"`
	ShippingAddress      *PurchaseShippingAddress `json:"ShippingAddress,omitempty"`
}

// POST method will return exception if Order status is not - DRAFT or NOT AVAILABLE
type CreatePurchaseOrder struct {
	TaskID                   string                      `json:"TaskID"`
	CombineAdditionalCharges bool                        `json:"CombineAdditionalCharges"` // if true then additional charges lines displayed in Lines array
	Memo                     string                      `json:"Memo"`
	Status                   string                      `json:"Status"` // Purchase Order Status. Possible Values are NOT AVAILABLE,DRAFT,AUTHORISED,VOIDED. For POST only DRAFT and AUTHORISED values accepted
	Lines                    []PurchaseOrderLine         `json:"Lines"`
	TotalBeforeTax           *float64                    `json:"TotalBeforeTax"` // Sum of order lines and additional charges without taxes. Not required for POST
	Tax                      *float64                    `json:"Tax"`            // Sum of order lines and additional charges taxes. Not required for POST
	Total                    *float64                    `json:"Total"`          // Sum of order lines and additional charges with taxes. Not required for POST
	AdditionalCharges        *[]PurchaseAdditionalCharge `json:"AdditionalCharges"`
}

type CreateSale struct {
	Location             string               `json:"Location"`
	ID                   *string              `json:"ID,omitempty"` // Required for PUT
	Customer             *string              `json:"Customer"`     // Required if CustomerID not provided
	CustomerID           *string              `json:"CustomerID"`
	CurrencyRate         *float64             `json:"CurrencyRate"` // Required if Customer currency is different from base currency
	Contact              *string              `json:"Contact"`
	Phone                *string              `json:"Phone"`
	Email                *string              `json:"Email"`
	DefaultAccount       *string              `json:"DefaultAccount"`
	SkipQuote            *bool                `json:"SkipQuote"`
	ShippingNotes        *string              `json:"ShippingNotes"`
	TaxRule              *string              `json:"TaxRule"`
	TaxInclusive         *bool                `json:"TaxInclusive"`
	Terms                *string              `json:"Terms"`
	PriceTier            *string              `json:"PriceTier"`
	ShipBy               *ISO8601Time         `json:"ShipBy"`
	SaleOrderDate        *ISO8601Time         `json:"SaleOrderDate"`
	LastModifiedOn       *ISO8601Time         `json:"LastModifiedOn"`
	Note                 *string              `json:"Note"`
	CustomerReference    *string              `json:"CustomerReference"`
	SalesRepresentative  *string              `json:"SalesRepresentative"`
	Carrier              *string              `json:"Carrier"`
	ExternalID           *string              `json:"ExternalID"`
	SaleType             *string              `json:"SaleType"` // Available only for POST method. Type of sale, available values are Simple, Advanced
	AdditionalAttributes map[string]string    `json:"AdditionalAttributes"`
	BillingAddress       *Address             `json:"BillingAddress"`
	ShippingAddress      *SaleShippingAddress `json:"ShippingAddress"`
}

// POST method will return exception if Order status is not - DRAFT or NOT AVAILABLE
// POST method will return exception if Quote status is not AUTHORISED.
// POST method can accept "AutoPickPackShipMode" property in body.
//
//	Default: NOPICK
//	This property affects only Sales with Type = Simple Sale and with no backorder when changing order status to AUTHORISED
//
// Available valus for AutoPickPackShipMode
//
//	NOPICK - Order will be created without picking
//	AUTOPICK - Order sill be created with Pick phase authorised
//	AUTOPICKPACK - Order sill be created with Pick and Pack phases authorised
//	AUTOPICKPACKSHIP - Order sill be created with Pick and Pack and Ship phases authorised
//
// If you add to Order product with quantity more than you have. Cin7 Core silently create backorder record for this product when changing order status to AUTHORISED
type CreateSaleOrder struct {
	SaleID                   *string
	SaleOrderNumber          *string // auto-generated
	CombineAdditionalCharges *bool   // if true then additional charges lines displayed in Lines array
	Memo                     *string
	Status                   *string // Order Status. Possible Values are values. For POST only DRAFT and AUTHORISED values accepted
	Lines                    *[]SaleOrderLine
	AdditionalCharges        *[]SaleAdditionalCharge
	TotalBeforeTax           *float64 // Sum of order lines and additional charges without taxes. Not required for POST
	Tax                      *float64 // Sum of order lines and additional charges taxes. Not required for POST
	Total                    *float64 // Sum of order lines and additional charges with taxes. Not required for POST
}

// POST method will return exception if Quote status is not - DRAFT or NOT AVAILABLE
// POST method will return exception if Sale's SkipQuote parameter is true.
// POST method does not support stand alone Credit Notes
type CreateSaleQuote struct {
	SaleID                   string
	CombineAdditionalCharges bool // if true then additional charges lines displayed in Lines array
	Memo                     string
	Status                   string // For POST only DRAFT and AUTHORISED values accepted
	Lines                    []SaleQuoteLine
	AdditionalCharges        *[]SaleAdditionalCharge
}

type CreateStockAdjustment struct {
	EffectiveDate *ISO8601Time
	Status        *string
	TaskID        *string
	Account       *string
	Reference     *string
	UpdateOnHand  *bool
	Lines         []StockLine
}

type CreateStockTake struct {
	EffectiveDate *ISO8601Time
	Account       *string
	LocationID    *string
	Location      *string
	Reference     *string
	Tags          *[]string
	Categories    *[]IDName
}

type CreateStockTransfer struct {
	Status               *string
	From                 *string
	FromLocation         *string
	To                   *string
	ToLocation           *string
	CostDistributionType *string
	InTransitAccount     *string
	DepartureDate        *ISO8601Time
	CompletionDate       *ISO8601Time
	RequiredByDate       *ISO8601Time
	Reference            *string
	Lines                *[]StockTransferLine
	SkipOrder            *bool
}

//////////////////////////////////////////
/*			BROWSE REQUESTS				*/
//////////////////////////////////////////

type BrowseProductRequest struct {
	Page                 *string
	Limit                *string
	ID                   *string
	Name                 *string
	Sku                  *string
	ModifiedSince        *string
	IncludeDeprecated    *string
	IncludeBOM           *string
	IncludeSuppliers     *string
	IncludeMovements     *string
	IncludeAttachments   *string
	IncludeReorderLevels *string
	IncludeCustomPrices  *string
}

type BrowseProductAvailabilityRequest struct {
	Page     *string
	Limit    *string
	ID       *string
	Name     *string
	Sku      *string
	Location *string
	Batch    *string
	Category *string
}

type BrowsePurchaseRequest struct {
	Page                  *string
	Limit                 *string
	Search                *string
	RequiredBy            *string
	UpdatedSince          *string
	OrderStatus           *string
	RestockReceivedStatus *string
	InvoiceStatus         *string
	CreditNoteStatus      *string
	UnstockStatus         *string
	Status                *string
	DropShipTaskID        *string
}

type ReadPurchaseRequest struct {
	ID                       string
	CombineAdditionalCharges *string
}

type ReadPurchaseOrderRequest struct {
	TaskID                   string
	CombineAdditionalCharges *string
}

type ReadSaleRequest struct {
	ID                       string
	CombineAdditionalCharges *string
	HideInventoryMovements   *string
	IncludeTransactions      *string
	CountryFormat            *string
}

type ReadSaleOrderRequest struct {
	SaleID                   string
	CombineAdditionalCharges *string
	IncludeProductInfo       *string
}

type ReadSaleQuoteRequest struct {
	SaleID                   string
	CombineAdditionalCharges *string
	IncludeProductInfo       *string
}

type BrowseLocationRequest struct {
	Page       *string
	Limit      *string
	ID         *string
	Deprecated *string
	Name       *string
}

type BrowseStockTransfer struct {
	Page   *string
	Limit  *string
	Status *string
	Search *string
}

//////////////////////////////////////////
/*			RESPONSE STRUCTS			*/
//////////////////////////////////////////

type ProductResponse struct {
	Total    *int       `json:"Total"`
	Page     *int       `json:"Page"`
	Products []*Product `json:"Products"`
}

type ProductAvailabilityResponse struct {
	Total                 *int                   `json:"Total"`
	Page                  *int                   `json:"Page"`
	ProductAvailabilities []*ProductAvailability `json:"ProductAvailabilityList"`
}

type PurchaseListResponse struct {
	Total     *int            `json:"Total"`
	Page      *int            `json:"Page"`
	Purchases []*PurchaseList `json:"PurchaseList"`
}

type PurchaseResponse struct {
	ID                      *string                          `json:"ID"`
	SupplierID              *string                          `json:"SupplierID"`
	Supplier                *string                          `json:"Supplier"`
	Contact                 *string                          `json:"Contact"`
	Phone                   *string                          `json:"Phone"`
	InventoryAccount        *string                          `json:"InventoryAccount"`
	BlindReceipt            *bool                            `json:"BlindReceipt"`
	Approach                *string                          `json:"Approach"`
	BaseCurrency            *string                          `json:"BaseCurrency"`
	SupplierCurrency        *string                          `json:"SupplierCurrency"`
	TaxRule                 *string                          `json:"TaxRule"`
	TaxCalculation          *string                          `json:"TaxCalculation"`
	Terms                   *string                          `json:"Terms"`
	RequiredBy              *ISO8601Time                     `json:"RequiredBy"`
	Location                *string                          `json:"Location"`
	Note                    *string                          `json:"Note"`
	OrderNumber             *string                          `json:"OrderNumber"`
	OrderDate               *ISO8601Time                     `json:"OrderDate"`
	Status                  *string                          `json:"Status"`
	RelatedDropShipSaleTask *string                          `json:"RelatedDropShipSaleTask"`
	CurrencyRate            *float64                         `json:"CurrencyRate"`
	LastUpdatedDate         *ISO8601Time                     `json:"LastUpdatedDate"`
	BillingAddress          *Address                         `json:"BillingAddress"`
	ShippingAddress         *PurchaseShippingAddress         `json:"ShippingAddress"`
	Order                   *PurchaseOrder                   `json:"Order"`
	StockReceived           *PurchaseStockReceived           `json:"StockReceived"`
	Invoice                 *PurchaseInvoice                 `json:"Invoice"`
	CreditNote              *PurchaseCreditNote              `json:"CreditNote"`
	ManualJournals          *PurchaseManualJournal           `json:"ManualJournals"`
	AdditionalAttributes    map[string]string                `json:"AdditionalAttributes"`
	Attachments             *[]PurchaseAttachmentLine        `json:"Attachments"`
	InventoryMovements      *[]PurchaseInventoryMovementLine `json:"InventoryMovements"`
}

type PurchaseOrderResponse struct {
	TaskID                   *string                     `json:"TaskID"`
	CombineAdditionalCharges *bool                       `json:"CombineAdditionalCharges"`
	Memo                     *string                     `json:"Memo"`
	Status                   *string                     `json:"Status"`
	Lines                    *[]PurchaseOrderLine        `json:"Lines"`
	AdditionalCharges        *[]PurchaseAdditionalCharge `json:"AdditionalCharges"`
	TotalBeforeTax           *float64                    `json:"TotalBeforeTax"`
	Tax                      *float64                    `json:"Tax"`
	Total                    *float64                    `json:"Total"`
}

type SaleResponse struct {
	ID                      *string                  `json:"ID"`
	Customer                *string                  `json:"Customer"`
	CustomerID              *string                  `json:"CustomerID"`
	Contact                 *string                  `json:"Contact"`
	Phone                   *string                  `json:"Phone"`
	Email                   *string                  `json:"Email"`
	DefaultAccount          *string                  `json:"DefaultAccount"`
	SkipQuote               *bool                    `json:"SkipQuote"`
	ShippingNotes           *string                  `json:"ShippingNotes"`
	BaseCurrency            *string                  `json:"BaseCurrency"`
	CustomerCurrency        *string                  `json:"CustomerCurrency"`
	TaxRule                 *string                  `json:"TaxRule"`
	TaxCalculation          *string                  `json:"TaxCalculation"`
	Terms                   *string                  `json:"Terms"`
	PriceTier               *string                  `json:"PriceTier"`
	ShipBy                  *ISO8601Time             `json:"ShipBy"`
	Location                *string                  `json:"Location"`
	SaleOrderDate           *ISO8601Time             `json:"SaleOrderDate"`
	LastModifiedOn          *ISO8601Time             `json:"LastModifiedOn"`
	Note                    *string                  `json:"Note"`
	CustomerReference       *string                  `json:"CustomerReference"`
	COGSAmount              *float64                 `json:"COGSAmount"`
	Status                  *string                  `json:"Status"`
	CombinedPickingStatus   *string                  `json:"CombinedPickingStatus"`
	CombinedPackingStatus   *string                  `json:"CombinedPackingStatus"`
	CombinedShippingStatus  *string                  `json:"CombinedShippingStatus"`
	FulFilmentStatus        *string                  `json:"FulFilmentStatus"`
	CombinedInvoiceStatus   *string                  `json:"CombinedInvoiceStatus"`
	CombinedPaymentStatus   *string                  `json:"CombinedPaymentStatus"`
	CombinedTrackingNumbers *string                  `json:"CombinedTrackingNumbers"`
	Carrier                 *string                  `json:"Carrier"`
	CurrencyRate            *float64                 `json:"CurrencyRate"`
	SalesRepresentative     *string                  `json:"SalesRepresentative"`
	ServiceOnly             *bool                    `json:"ServiceOnly"`
	Type                    *string                  `json:"Type"`
	SourceChannel           *string                  `json:"SourceChannel"`
	ExternalID              *string                  `json:"ExternalID"`
	AdditionalAttributes    map[string]string        `json:"AdditionalAttributes"`
	BillingAddress          *Address                 `json:"BillingAddress"`
	ShippingAddress         *SaleShippingAddress     `json:"ShippingAddress"`
	Quote                   *SaleQuote               `json:"Quote"`
	Order                   *SaleOrder               `json:"Order"`
	ManualJournals          *SaleManualJournal       `json:"ManualJournals"`
	Fulfilments             *[]SaleFulfillment       `json:"Fulfilments"`
	Invoices                *[]SaleInvoice           `json:"Invoices"`
	CreditNotes             *[]SaleCreditNote        `json:"CreditNotes"`
	Attachments             *[]AttachmentLine        `json:"Attachments"`
	InventoryMovements      *[]InventoryMovementLine `json:"InventoryMovements"`
	Transactions            *[]SaleTransactionLine   `json:"Transactions"`
}

type SaleOrderResponse struct {
	SaleID                   *string                 `json:"SaleID"`
	SaleOrderNumber          *string                 `json:"SaleOrderNumber"`
	CombineAdditionalCharges *bool                   `json:"CombineAdditionalCharges,omitempty"`
	Memo                     *string                 `json:"Memo"`
	Status                   *string                 `json:"Status"`
	Lines                    *[]SaleOrderLine        `json:"Lines"`
	AdditionalCharges        *[]SaleAdditionalCharge `json:"AdditionalCharges,omitempty"`
	TotalBeforeTax           *float64                `json:"TotalBeforeTax"`
	Tax                      *float64                `json:"Tax"`
	Total                    *float64                `json:"Total"`
}

type SaleQuoteResponse struct {
	SaleID                   string                  `json:"SaleID"`
	CombineAdditionalCharges bool                    `json:"CombineAdditionalCharges,omitempty"`
	Memo                     string                  `json:"Memo"`
	Status                   string                  `json:"Status"`
	TotalBeforeTax           float64                 `json:"TotalBeforeTax"`
	Tax                      float64                 `json:"Tax"`
	Total                    float64                 `json:"Total"`
	Lines                    []SaleQuoteLine         `json:"Lines"`
	AdditionalCharges        *[]SaleAdditionalCharge `json:"AdditionalCharges,omitempty"`
}

type StockAdjustmentResponse struct {
	TaskID             *string                 `json:"TaskID"`
	EffectiveDate      *ISO8601Time            `json:"EffectiveDate"`
	StocktakeNumber    *string                 `json:"StocktakeNumber"`
	Status             *string                 `json:"Status"`
	Account            *string                 `json:"Account"`
	Reference          *string                 `json:"Reference"`
	ExistingStockLines *[]StockLine            `json:"ExistingStockLines"`
	NewStockLines      *[]StockLine            `json:"NewStockLines"`
	Transactions       *[]TransactionStockLine `json:"Transactions"`
}

type StockTakeResponse struct {
	TaskID                     *string                 `json:"TaskID"`
	EffectiveDate              *ISO8601Time            `json:"EffectiveDate"`
	StocktakeNumber            *string                 `json:"StocktakeNumber"`
	Status                     *string                 `json:"Status"`
	Account                    *string                 `json:"Account"`
	LocationID                 *string                 `json:"LocationID"`
	Location                   *string                 `json:"Location"`
	Reference                  *string                 `json:"Reference"`
	Tags                       *[]string               `json:"Tags"`
	PickZones                  *[]string               `json:"PickZones"`
	StockLocators              *[]string               `json:"StockLocators"`
	Categories                 *[]IDName               `json:"Categories"`
	Brands                     *[]IDName               `json:"Brands"`
	Bins                       *[]IDName               `json:"Bins"`
	NonZeroStockOnHandProducts *[]StockLine            `json:"NonZeroStockOnHandProducts"`
	ZeroStockOnHandProducts    *[]StockLine            `json:"ZeroStockOnHandProducts"`
	Transactions               *[]TransactionStockLine `json:"Transactions"`
}

type StockTransferResponse struct {
	Total             *int             `json:"Total"`
	Page              *int             `json:"Page"`
	StockTransferList []*StockTransfer `json:"StockTransferList"`
}

type LocationResponse struct {
	Total     *int        `json:"Total"`
	Page      *int        `json:"Page"`
	Locations []*Location `json:"LocationList"`
}

//////////////////////////////////////////
/*			OTHER STRUCTS				*/
//////////////////////////////////////////

type Product struct {
	ID                           *string                         `json:"ID,omitempty"`
	SKU                          *string                         `json:"SKU,omitempty"`
	Name                         *string                         `json:"Name,omitempty"`
	Category                     *string                         `json:"Category,omitempty"`
	UOM                          *string                         `json:"UOM,omitempty"`
	Status                       *string                         `json:"Status,omitempty"`
	Type                         *string                         `json:"Type,omitempty"`
	CostingMethod                *string                         `json:"CostingMethod,omitempty"`
	QuantityToProduce            *float64                        `json:"QuantityToProduce,omitempty"`
	AssemblyCostEstimationMethod *string                         `json:"AssemblyCostEstimationMethod,omitempty"`
	Brand                        *string                         `json:"Brand,omitempty"`
	DropShipMode                 *string                         `json:"DropShipMode,omitempty"`
	DefaultLocation              *string                         `json:"DefaultLocation,omitempty"`
	Length                       *float64                        `json:"Length,omitempty"`
	Width                        *float64                        `json:"Width,omitempty"`
	Height                       *float64                        `json:"Height,omitempty"`
	Weight                       *float64                        `json:"Weight,omitempty"`
	WeightUnits                  *string                         `json:"WeightUnits,omitempty"`
	DimensionsUnits              *string                         `json:"DimensionsUnits,omitempty"`
	Barcode                      *string                         `json:"Barcode,omitempty"`
	MinimumBeforeReorder         *float64                        `json:"MinimumBeforeReorder,omitempty"`
	ReorderQuantity              *float64                        `json:"ReorderQuantity,omitempty"`
	PriceTier1                   *float64                        `json:"PriceTier1,omitempty"`
	PriceTier2                   *float64                        `json:"PriceTier2,omitempty"`
	PriceTier3                   *float64                        `json:"PriceTier3,omitempty"`
	PriceTier4                   *float64                        `json:"PriceTier4,omitempty"`
	PriceTier5                   *float64                        `json:"PriceTier5,omitempty"`
	PriceTier6                   *float64                        `json:"PriceTier6,omitempty"`
	PriceTier7                   *float64                        `json:"PriceTier7,omitempty"`
	PriceTier8                   *float64                        `json:"PriceTier8,omitempty"`
	PriceTier9                   *float64                        `json:"PriceTier9,omitempty"`
	PriceTier10                  *float64                        `json:"PriceTier10,omitempty"`
	PriceTiers                   map[string]*float64             `json:"PriceTiers,omitempty"`
	AverageCost                  *float64                        `json:"AverageCost,omitempty"`
	ShortDescription             *string                         `json:"ShortDescription,omitempty"`
	Description                  *string                         `json:"Description,omitempty"`
	InternalNote                 *string                         `json:"InternalNote,omitempty"`
	AdditionalAttribute1         *string                         `json:"AdditionalAttribute1,omitempty"`
	AdditionalAttribute2         *string                         `json:"AdditionalAttribute2,omitempty"`
	AdditionalAttribute3         *string                         `json:"AdditionalAttribute3,omitempty"`
	AdditionalAttribute4         *string                         `json:"AdditionalAttribute4,omitempty"`
	AdditionalAttribute5         *string                         `json:"AdditionalAttribute5,omitempty"`
	AdditionalAttribute6         *string                         `json:"AdditionalAttribute6,omitempty"`
	AdditionalAttribute7         *string                         `json:"AdditionalAttribute7,omitempty"`
	AdditionalAttribute8         *string                         `json:"AdditionalAttribute8,omitempty"`
	AdditionalAttribute9         *string                         `json:"AdditionalAttribute9,omitempty"`
	AdditionalAttribute10        *string                         `json:"AdditionalAttribute10,omitempty"`
	AlwaysShowQuantity           *float64                        `json:"AlwaysShowQuantity,omitempty"`
	BOMType                      *string                         `json:"BOMType,omitempty"`
	WarrantyName                 *string                         `json:"WarrantyName,omitempty"`
	AttributeSet                 *string                         `json:"AttributeSet,omitempty"`
	DiscountRule                 *string                         `json:"DiscountRule,omitempty"`
	Tags                         *string                         `json:"Tags,omitempty"`
	StockLocator                 *string                         `json:"StockLocator,omitempty"`
	COGSAccount                  *string                         `json:"COGSAccount,omitempty"`
	RevenueAccount               *string                         `json:"RevenueAccount,omitempty"`
	ExpenseAccount               *string                         `json:"ExpenseAccount,omitempty"`
	InventoryAccount             *string                         `json:"InventoryAccount,omitempty"`
	PurchaseTaxRule              *string                         `json:"PurchaseTaxRule,omitempty"`
	SaleTaxRule                  *string                         `json:"SaleTaxRule,omitempty"`
	LastModifiedOn               *ISO8601Time                    `json:"LastModifiedOn,omitempty"`
	Sellable                     *bool                           `json:"Sellable,omitempty"`
	PickZones                    *string                         `json:"PickZones,omitempty"`
	BillOfMaterial               *bool                           `json:"BillOfMaterial,omitempty"`
	AutoAssembly                 *bool                           `json:"AutoAssembly,omitempty"`
	AutoDisassembly              *bool                           `json:"AutoDisassembly,omitempty"`
	AssemblyInstructionURL       *string                         `json:"AssemblyInstructionURL,omitempty"`
	Suppliers                    *[]ProductSupplier              `json:"Suppliers,omitempty"`
	ReorderLevels                *[]ProductReorderLevel          `json:"ReorderLevels,omitempty"`
	BillOfMaterialsProducts      *[]ProductBillOfMaterialProduct `json:"BillOfMaterialsProducts,omitempty"`
	BillOfMaterialsServices      *[]ProductBillOfMaterialService `json:"BillOfMaterialsServices,omitempty"`
	Movements                    *[]ProductMovement              `json:"Movements,omitempty"`
	Attachments                  *[]ProductAttachment            `json:"Attachments,omitempty"`
	CustomPrices                 *[]CustomPrice                  `json:"CustomPrices,omitempty"`
	CartonHeight                 *float64                        `json:"CartonHeight,omitempty"`
	CartonWidth                  *float64                        `json:"CartonWidth,omitempty"`
	CartonLength                 *float64                        `json:"CartonLength,omitempty"`
	CartonQuantity               *float64                        `json:"CartonQuantity,omitempty"`
	CartonInnerQuantity          *float64                        `json:"CartonInnerQuantity,omitempty"`
	HSCode                       *string                         `json:"HSCode,omitempty"`
	CountryOfOrigin              *string                         `json:"CountryOfOrigin,omitempty"`
	CountryOfOriginCode          *string                         `json:"CountryOfOriginCode,omitempty"`
}

type ProductAvailability struct {
	ID          *string      `json:"ID"`
	SKU         *string      `json:"SKU"`
	Name        *string      `json:"Name"`
	Barcode     *string      `json:"Barcode"`
	Location    *string      `json:"Location"`
	Bin         *string      `json:"Bin"`
	Batch       *string      `json:"Batch"`
	ExpiryDate  *ISO8601Time `json:"ExpiryDate"`
	OnHand      *float64     `json:"OnHand"`
	Allocated   *float64     `json:"Allocated"`
	Available   *float64     `json:"Available"`
	OnOrder     *float64     `json:"OnOrder"`
	StockOnHand *float64     `json:"StockOnHand"`
	InTransit   *float64     `json:"InTransit"`
}

type ProductReorderLevel struct {
	LocationID           string  `json:"LocationID"`   // Required if LocationName is empty
	LocationName         string  `json:"LocationName"` // Required if LocationID is empty
	PickZones            string  `json:"PickZones"`
	MinimumBeforeReorder *int    `json:"MinimumBeforeReorder"`
	ReorderQuantity      *int    `json:"ReorderQuantity"`
	StockLocator         *string `json:"StockLocator"`
}

type ProductBillOfMaterialProduct struct {
	ComponentProductID string   `json:"ComponentProductID"` // Required if ProductCode is empty
	ProductCode        string   `json:"ProductCode"`        // Required if ComponentProductID is empty
	Quantity           float64  `json:"Quantity"`
	Name               *string  `json:"Name"`
	WastagePercent     *float64 `json:"WastagePercent"`  // WastagePercent and WastageQuantity are mutually exclusive
	WastageQuantity    *float64 `json:"WastageQuantity"` // WastagePercent and WastageQuantity are mutually exclusive
	CostPercentage     *float64 `json:"CostPercentage"`  // If left as zero equal distribution will be used between components
}

type ProductBillOfMaterialService struct {
	ComponentProductID string  `json:"ComponentProductID"` // Required if Name is empty
	Name               string  `json:"Name"`               // Required if ComponentProductID is empty
	Quantity           float64 `json:"Quantity"`
	ExpenseAccount     *string `json:"ExpenseAccount,omitempty"`
	PriceTier          *int    `json:"PriceTier"`
}

type ProductMovement struct {
	TaskID     *string      `json:"TaskID"`
	Type       *string      `json:"Type"`
	Date       *ISO8601Time `json:"Date"`
	Number     *string      `json:"Number"`
	Status     *int         `json:"Status"`
	Quantity   *float64     `json:"Quantity"`
	Amount     *float64     `json:"Amount"`
	Location   *float64     `json:"Location"`
	BatchSN    *float64     `json:"BatchSN"`
	ExpiryDate *ISO8601Time `json:"ExpiryDate"`
	FromTo     *string      `json:"FromTo"`
}

type ProductAttachment struct {
	ID          *string `json:"ID"`
	ContentType *string `json:"ContentType"`
	IsDefault   *bool   `json:"IsDefault"`
	FileName    *string `json:"FileName"`
	DownloadUrl *string `json:"DownloadUrl"`
}

type CustomPrice struct {
	ProductID    string  `json:"ProductID"`    // Required if ProductSKU is empty
	ProductSKU   string  `json:"ProductSKU"`   // Required if ProductID is empty
	CustomerID   string  `json:"CustomerID"`   // Required if CustomerName is empty
	CustomerName string  `json:"CustomerName"` // Required if CustomerID is empty
	Price        float64 `json:"Price"`
	ProductName  *string `json:"ProductName"`
}

type ProductSupplier struct {
	SupplierID             *string                  `json:"SupplierID,omitempty"`   // Required if SupplierName is empty
	SupplierName           *string                  `json:"SupplierName,omitempty"` // Required if SupplierID is empty
	ProductID              *string                  `json:"ProductID,omitempty"`    // Required when not nested within the Product and ProductSKU is empty
	ProductSKU             *string                  `json:"ProductSKU,omitempty"`   // Required when not nested within the Product and ProductID is empty
	ProductSupplierID      *string                  `json:"ProductSupplierID,omitempty"`
	SupplierInventoryCode  *string                  `json:"SupplierInventoryCode,omitempty"`
	SupplierProductName    *string                  `json:"SupplierProductName,omitempty"`
	Cost                   *float64                 `json:"Cost,omitempty"`
	FixedCost              *float64                 `json:"FixedCost,omitempty"`
	Currency               *string                  `json:"Currency,omitempty"`
	DropShip               *bool                    `json:"DropShip,omitempty"`
	SupplierProductURL     *string                  `json:"SupplierProductURL,omitempty"`
	LastSupplied           *ISO8601Time             `json:"LastSupplied,omitempty"`
	ProductSupplierOptions *[]ProductSupplierOption `json:"ProductSupplierOptions,omitempty"`
}

type ProductSupplierOption struct {
	ID               string                            `json:"ID"`
	LocationID       string                            `json:"LocationID"`   // Required if LocationName is empty
	LocationName     string                            `json:"LocationName"` // Required if LocationID is empty
	ReorderQuantity  *float64                          `json:"ReorderQuantity"`
	Lead             *int                              `json:"Lead"`
	Safety           *int                              `json:"Safety"`
	MinimumToReorder *float64                          `json:"MinimumToReorder"`
	SupplyIntervals  *[]ProductSupplierOptionsInterval `json:"SupplyIntervals"`
}

type ProductSupplierOptionsInterval struct {
	ID                string       `json:"ID"`
	DeliveryMethod    string       `json:"DeliveryMethod"`    // can be 'Fixed' and 'Interval'
	IntervalDays      *int         `json:"IntervalDays"`      // required when DeliveryMethod = Interval
	IntervalStartDate *ISO8601Time `json:"IntervalStartDate"` // required when DeliveryMethod = Interval
	IsMonday          *bool        `json:"IsMonday"`          // Required when DeliveryMethod = Fixed
	IsTuesday         *bool        `json:"IsTuesday"`         // Required when DeliveryMethod = Fixed
	IsWednesday       *bool        `json:"IsWednesday"`       // Required when DeliveryMethod = Fixed
	IsThursday        *bool        `json:"IsThursday"`        // Required when DeliveryMethod = Fixed
	IsFriday          *bool        `json:"IsFriday"`          // Required when DeliveryMethod = Fixed
	IsSaturday        *bool        `json:"IsSaturday"`        // Required when DeliveryMethod = Fixed
	IsSunday          *bool        `json:"IsSunday"`          // Required when DeliveryMethod = Fixed
}

type Address struct {
	ID                  *string `json:"ID,omitempty"`        //applicable for sales only, POST/PUT methods
	DisplayAddressLine1 *string `json:"DisplayAddressLine1"` // Address Line 1 as displayed on Sale form. = Line1 + Line2
	DisplayAddressLine2 *string `json:"DisplayAddressLine2"` // Address Line 2 as displayed on Sale form. = City + State/Region + Zip/Postcode + Country
	Line1               *string `json:"Line1"`
	Line2               *string `json:"Line2"`
	City                *string `json:"City"`
	State               *string `json:"State"`
	Postcode            *string `json:"Postcode"`
	Country             *string `json:"Country"`
}

type SaleOrder struct {
	SaleOrderNumber   *string                 `json:"SaleOrderNumber"`
	Memo              *string                 `json:"Memo"`
	Status            *string                 `json:"Status"`
	Lines             *[]SaleOrderLine        `json:"Lines"`
	AdditionalCharges *[]SaleAdditionalCharge `json:"AdditionalCharges"`
	TotalBeforeTax    *float64                `json:"TotalBeforeTax"`
	Tax               *float64                `json:"Tax"`
	Total             *float64                `json:"Total"`
}

type SaleOrderLine struct {
	ProductID         *string  `json:"ProductID"`
	SKU               *string  `json:"Sku"`
	Name              *string  `json:"Name"`
	Quantity          *float64 `json:"Quantity"`
	Price             *float64 `json:"Price"`
	Discount          *float64 `json:"Discount"`
	Tax               *float64 `json:"Tax"`
	AverageCost       *float64 `json:"AverageCost"`
	TaxRule           *string  `json:"TaxRule"`
	Comment           *string  `json:"Comment"`
	DropShip          *bool    `json:"DropShip"`
	BackorderQuantity *float64 `json:"BackorderQuantity"`
	Total             *float64 `json:"Total"`
}

type SaleShippingAddress struct {
	ID                  *string `json:"ID"`
	DisplayAddressLine1 *string `json:"DisplayAddressLine1"`
	DisplayAddressLine2 *string `json:"DisplayAddressLine2"`
	Line1               *string `json:"Line1"`
	Line2               *string `json:"Line2"`
	City                *string `json:"City"`
	State               *string `json:"State"`
	Postcode            *string `json:"Postcode"`
	Country             *string `json:"Country"`
	Company             *string `json:"Company"`
	Contact             *string `json:"Contact"`
	ShipToOther         *bool   `json:"ShipToOther"`
}

type SalePaymentLine struct {
	ID           *string      `json:"ID"`
	Reference    *string      `json:"Reference"`
	Amount       *float64     `json:"Amount"`
	DatePaid     *ISO8601Time `json:"DatePaid"`
	Account      *string      `json:"Account"`
	CurrencyRate *float64     `json:"CurrencyRate"`
	DateCreated  *ISO8601Time `json:"DateCreated"`
}

type SaleQuote struct {
	Memo              *string                 `json:"Memo"`
	Status            *string                 `json:"Status"`
	Prepayments       *[]SalePaymentLine      `json:"Prepayments"`
	Lines             *[]SaleQuoteLine        `json:"Lines"`
	AdditionalCharges *[]SaleAdditionalCharge `json:"AdditionalCharges"`
	TotalBeforeTax    *float64                `json:"TotalBeforeTax"`
	Tax               *float64                `json:"Tax"`
	Total             *float64                `json:"Total"`
}

type SaleAdditionalCharge struct {
	Description *string  `json:"Description"`
	Price       *float64 `json:"Price"`
	Quantity    *float64 `json:"Quantity"`
	Discount    *float64 `json:"Discount"`
	Tax         *float64 `json:"Tax"`
	Total       *float64 `json:"Total"`
	TaxRule     *string  `json:"TaxRule"`
	Comment     *string  `json:"Comment"`
}

type SaleInvoiceAdditionalCharge struct {
	Description *string  `json:"Description"`
	Price       *float64 `json:"Price"`
	Quantity    *float64 `json:"Quantity"`
	Discount    *float64 `json:"Discount"`
	Tax         *float64 `json:"Tax"`
	Total       *float64 `json:"Total"`
	TaxRule     *string  `json:"TaxRule"`
	Account     *string  `json:"Account"`
	Comment     *string  `json:"Comment"`
}

type SaleQuoteLine struct {
	Name        string   `json:"Name"`
	Quantity    float64  `json:"Quantity"`
	Price       float64  `json:"Price"`
	Tax         float64  `json:"Tax"`
	AverageCost float64  `json:"AverageCost"`
	TaxRule     string   `json:"TaxRule"`
	Comment     string   `json:"Comment"`
	ProductID   *string  `json:"ProductID"` // Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
	SKU         *string  `json:"SKU"`       // Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
	Discount    *float64 `json:"Discount"`
	Total       *float64 `json:"Total"`
}

type SaleFulfillment struct {
	TaskID              *string                  `json:"TaskID"`
	FulfillmentNumber   *int                     `json:"FulfillmentNumber"`
	LinkedInvoiceNumber *string                  `json:"LinkedInvoiceNumber"`
	FulfillmentStatus   *string                  `json:"FulfilmentStatus"`
	Pick                *SaleFulfillmentPickPack `json:"Pick"`
	Pack                *SaleFulfillmentPickPack `json:"Pack"`
	Ship                *SaleFulfillmentShip     `json:"Ship"`
}

type SaleInvoice struct {
	TaskID                  *string                        `json:"TaskID"`
	InvoiceNumber           *string                        `json:"InvoiceNumber"`
	Memo                    *string                        `json:"Memo"`
	Status                  *string                        `json:"Status"`
	InvoiceDate             *ISO8601Time                   `json:"InvoiceDate"`
	InvoiceDueDate          *ISO8601Time                   `json:"InvoiceDueDate"`
	CurrencyConversionRate  *float64                       `json:"CurrencyConversionRate"`
	BillingAddressLine1     *string                        `json:"BillingAddressLine1"`
	BillingAddressLine2     *string                        `json:"BillingAddressLine2"`
	LinkedFulfillmentNumber *int                           `json:"LinkedFulfillmentNumber"`
	Lines                   *[]SaleInvoiceLine             `json:"Lines"`
	AdditionalCharges       *[]SaleInvoiceAdditionalCharge `json:"AdditionalCharges"`
	Payments                *[]SalePaymentLine             `json:"Payments"`
	TotalBeforeTax          *float64                       `json:"TotalBeforeTax"`
	Tax                     *float64                       `json:"Tax"`
	Total                   *float64                       `json:"Total"`
	Paid                    *float64                       `json:"Paid"`
}

type SaleInvoiceLine struct {
	ProductID   *string  `json:"ProductID"`
	SKU         *string  `json:"SKU"`
	Name        *string  `json:"Name"`
	Quantity    *float64 `json:"Quantity"`
	Price       *float64 `json:"Price"`
	Discount    *float64 `json:"Discount"`
	Tax         *float64 `json:"Tax"`
	Total       *float64 `json:"Total"`
	AverageCost *float64 `json:"AverageCost"`
	TaxRule     *string  `json:"TaxRule"`
	Account     *string  `json:"Account"`
	Comment     *string  `json:"Comment"`
	AdditionalFields
}

type SaleFulfillmentPickPack struct {
	Status *string                        `json:"Status"`
	Lines  *[]SaleFulfillmentPickPackLine `json:"Lines"`
}

type SaleFulfillmentPickPackLine struct {
	ProductID                  *string      `json:"ProductID"`
	SKU                        *string      `json:"SKU"`
	Name                       *string      `json:"Name"`
	Location                   *string      `json:"Location"`
	LocationID                 *string      `json:"LocationID"`
	Quantity                   *float64     `json:"Quantity"`
	BatchSN                    *string      `json:"BatchSN"`
	ExpiryDate                 *ISO8601Time `json:"ExpiryDate"`
	Box                        *string      `json:"Box"`
	NonInventory               *bool        `json:"NonInventory"`
	WarrantyRegistrationNumber *string      `json:"WarrantyRegistrationNumber"`
	RestockLocation            *string      `json:"RestockLocation"`
	RestockLocationID          *string      `json:"RestockLocationID"`
	AdditionalFields
}

type SaleFulfillmentShip struct {
	Status          *string                    `json:"Status"`
	RequireBy       *ISO8601Time               `json:"RequireBy"`
	ShippingAddress *SaleShippingAddress       `json:"ShippingAddress"`
	ShippingNotes   *string                    `json:"ShippingNotes"`
	Lines           *[]SaleFulfillmentShipLine `json:"Lines"`
}

type SaleFulfillmentShipLine struct {
	ID             *string      `json:"ID"`
	ShipmentDate   *ISO8601Time `json:"ShipmentDate"`
	Carrier        *string      `json:"Carrier"`
	Boxes          *string      `json:"Boxes"`
	TrackingNumber *string      `json:"TrackingNumber"`
	TrackingURL    *string      `json:"TrackingURL"`
	IsShipped      *bool        `json:"IsShipped"`
}

type SaleCreditNote struct {
	TaskID                   *string                        `json:"TaskID"`
	CreditNoteInvoiceNumber  *string                        `json:"CreditNoteInvoiceNumber"`
	Memo                     *string                        `json:"Memo"`
	Status                   *string                        `json:"Status"`
	CreditNoteDate           *ISO8601Time                   `json:"CreditNoteDate"`
	CreditNoteNumber         *string                        `json:"CreditNoteNumber"`
	CreditNoteConversionRate *float64                       `json:"CreditNoteConversionRate"`
	TotalBeforeTax           *float64                       `json:"TotalBeforeTax"`
	Tax                      *float64                       `json:"Tax"`
	Total                    *float64                       `json:"Total"`
	RestockStatus            *string                        `json:"RestockStatus"`
	Lines                    *[]SaleInvoiceLine             `json:"Lines"`
	AdditionalCharges        *[]SaleInvoiceAdditionalCharge `json:"AdditionalCharges"`
	Refunds                  *[]SalePaymentLine             `json:"Refunds"`
	Restock                  *[]SaleFulfillmentPickPackLine `json:"Restock"`
}

type SaleManualJournal struct {
	Status *string `json:"Status"`
	Lines  *[]SaleManualJournalLine
}

type SaleManualJournalLine struct {
	Reference *string      `json:"Reference"`
	Amount    *float64     `json:"Amount"`
	Date      *ISO8601Time `json:"Date"`
	Debit     *string      `json:"Debit"`
	Credit    *string      `json:"Credit"`
}

type SaleTransactionLine struct {
	TaskID        *string      `json:"TaskID"`
	TransactionID *string      `json:"TransactionID"`
	Debit         *string      `json:"Debit"`
	Credit        *string      `json:"Credit"`
	Description   *string      `json:"Description"`
	Amount        *float64     `json:"Amount"`
	EffectiveDate *ISO8601Time `json:"EffectiveDate"`
}

type PurchaseList struct {
	ID                      *string      `json:"ID"`
	BlindReceipt            *bool        `json:"BlindReceipt"`
	OrderNumber             *string      `json:"OrderNumber"`
	Status                  *string      `json:"Status"`
	OrderDate               *ISO8601Time `json:"OrderDate"`
	InvoiceDate             *ISO8601Time `json:"InvoiceDate"`
	Supplier                *string      `json:"Supplier"`
	SupplierID              *string      `json:"SupplierID"`
	InvoiceNumber           *string      `json:"InvoiceNumber"`
	InvoiceAmount           *float64     `json:"InvoiceAmount"`
	PaidAmount              *float64     `json:"PaidAmount"`
	InvoiceDueDate          *ISO8601Time `json:"InvoiceDueDate"`
	RequiredBy              *ISO8601Time `json:"RequiredBy"`
	BaseCurrency            *string      `json:"BaseCurrency"`
	SupplierCurrency        *string      `json:"SupplierCurrency"`
	CreditNoteNumber        *string      `json:"CreditNoteNumber"`
	OrderStatus             *string      `json:"OrderStatus"`
	StockReceivedStatus     *string      `json:"StockReceivedStatus"`
	UnstockStatus           *string      `json:"UnstockStatus"`
	InvoiceStatus           *string      `json:"InvoiceStatus"`
	CreditNoteStatus        *string      `json:"CreditNoteStatus"`
	LastUpdatedDate         *ISO8601Time `json:"LastUpdatedDate"`
	CombinedReceivingStatus *string      `json:"CombinedReceivingStatus"`
	CombinedInvoiceStatus   *string      `json:"CombinedInvoiceStatus"`
	CombinedPaymentStatus   *string      `json:"CombinedPaymentStatus"`
	Type                    *string      `json:"Type"`
	IsServiceOnly           *bool        `json:"IsServiceOnly"`
	DropShipTaskID          *string      `json:"DropShipTaskID"`
}

type PurchaseOrder struct {
	Memo              *string                     `json:"Memo"`
	Status            *string                     `json:"Status"`
	Lines             *[]PurchaseOrderLine        `json:"Lines"`
	AdditionalCharges *[]PurchaseAdditionalCharge `json:"AdditionalCharges"`
	Prepayments       *[]SalePaymentLine          `json:"Prepayments"`
	TotalBeforeTax    *float64                    `json:"TotalBeforeTax"`
	Tax               *float64                    `json:"Tax"`
	Total             *float64                    `json:"Total"`
}

type PurchaseShippingAddress struct {
	DisplayAddressLine1 *string `json:"DisplayAddressLine1"` // Address Line 1 as displayed on Sale form. = Line1 + Line2
	DisplayAddressLine2 *string `json:"DisplayAddressLine2"` // Address Line 2 as displayed on Sale form. = City + State/Region + Zip/Postcode + Country
	Line1               *string `json:"Line1"`
	Line2               *string `json:"Line2"`
	City                *string `json:"City"`
	State               *string `json:"State"`
	Postcode            *string `json:"Postcode"`
	Country             *string `json:"Country"`
	ShipToOther         *bool   `json:"ShipToOther"`
	Company             *string `json:"Company"` // Company name when ShipToOther is set to True
}

type PurchaseOrderLine struct {
	Name        string   `json:"Name"`
	Quantity    float64  `json:"Quantity"`
	Price       float64  `json:"Price"`
	Tax         float64  `json:"Tax"`
	TaxRule     string   `json:"TaxRule"`
	Total       float64  `json:"Total"`
	ProductID   *string  `json:"ProductID"` // Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
	SKU         *string  `json:"SKU"`       // Required If CombineAdditionalCharges param exist for this endpoint and it have values = false
	Discount    *float64 `json:"Discount"`
	SupplierSKU *string  `json:"SupplierSKU"`
	Comment     *string  `json:"Comment"`
	AdditionalFields
}

type PurchaseAdditionalCharge struct {
	Description *string  `json:"Description"`
	Reference   *string  `json:"Reference"`
	Price       *float64 `json:"Price"`
	Quantity    *float64 `json:"Quantity"`
	Discount    *float64 `json:"Discount"`
	Tax         *float64 `json:"Tax"`
	Total       *float64 `json:"Total"`
	TaxRule     *string  `json:"TaxRule"`
}

type PurchaseStockReceived struct {
	Status *string              `json:"Status"`
	Lines  *[]PurchaseStockLine `json:"Lines"`
}

type PurchaseStockLine struct {
	Date        *ISO8601Time `json:"Date"`
	Quantity    *float64     `json:"Quantity"`
	ProductID   *string      `json:"ProductID"`
	SKU         *string      `json:"SKU"`
	Name        *string      `json:"Name"`
	Location    *string      `json:"Location"`
	LocationID  *string      `json:"LocationID"`
	Received    *bool        `json:"Received"`
	BatchSN     *string      `json:"BatchSN"`
	SupplierSKU *string      `json:"SupplierSKU"`
	ExpiryDate  *ISO8601Time `json:"ExpiryDate"`
	CardID      *string      `json:"CardID"`
	AdditionalFields
}

type PurchaseInvoice struct {
	InvoiceDate       *ISO8601Time                       `json:"InvoiceDate"`
	InvoiceDueDate    *ISO8601Time                       `json:"InvoiceDueDate"`
	InvoiceNumber     *string                            `json:"InvoiceNumber"`
	Status            *string                            `json:"Status"`
	Lines             *[]PurchaseInvoiceLine             `json:"Lines"`
	AdditionalCharges *[]PurchaseInvoiceAdditionalCharge `json:"AdditionalCharges"`
	Payments          *[]SalePaymentLine                 `json:"Payments"`
	TotalBeforeTax    *float64                           `json:"TotalBeforeTax"`
	Tax               *float64                           `json:"Tax"`
	Total             *float64                           `json:"Total"`
	Paid              *float64                           `json:"Paid"`
}

type PurchaseInvoiceLine struct {
	ProductID *string  `json:"ProductID"`
	SKU       *string  `json:"SKU"`
	Name      *string  `json:"Name"`
	Quantity  *float64 `json:"Quantity"`
	Price     *float64 `json:"Price"`
	Discount  *float64 `json:"Discount"`
	Tax       *float64 `json:"Tax"`
	TaxRule   *string  `json:"TaxRule"`
	Account   *string  `json:"Account"`
	Comment   *string  `json:"Comment"`
	Total     *float64 `json:"Total"`
	AdditionalFields
}

type PurchaseInvoiceAdditionalCharge struct {
	Description *string  `json:"Description"`
	Reference   *string  `json:"Reference"`
	Quantity    *float64 `json:"Quantity"`
	Price       *float64 `json:"Price"`
	Discount    *float64 `json:"Discount"`
	Tax         *float64 `json:"Tax"`
	TaxRule     *string  `json:"TaxRule"`
	Account     *string  `json:"Account"`
	Total       *float64 `json:"Total"`
}

type PurchaseCreditNote struct {
	CreditNoteNumber  *string                            `json:"CreditNoteNumber"`
	CreditNoteDate    *ISO8601Time                       `json:"CreditNoteDate"`
	Status            *string                            `json:"Status"`
	Lines             *[]PurchaseInvoiceLine             `json:"Lines"`
	AdditionalCharges *[]PurchaseInvoiceAdditionalCharge `json:"AdditionalCharges"`
	Refunds           *[]SalePaymentLine                 `json:"Refunds"`
	Unstock           *[]interface{}                     `json:"Unstock"`
	TotalBeforeTax    *float64                           `json:"TotalBeforeTax"`
	Tax               *float64                           `json:"Tax"`
	Total             *float64                           `json:"Total"`
}

type PurchaseUnstockLine struct {
	CardID     *string      `json:"CardID"`
	Date       *ISO8601Time `json:"Date"`
	Quantity   *float64     `json:"Quantity"`
	ProductID  *string      `json:"ProductID"`
	SKU        *string      `json:"SKU"`
	Name       *string      `json:"Name"`
	Location   *string      `json:"Location"`
	BatchSN    *string      `json:"BatchSN"`
	ExpiryDate *ISO8601Time `json:"ExpiryDate"`
	AdditionalFields
}

type PurchaseManualJournal struct {
	Status *string                      `json:"Status"`
	Lines  *[]PurchaseManualJournalLine `json:"Lines"`
}

type PurchaseManualJournalLine struct {
	Reference *string      `json:"Reference"`
	Amount    *float64     `json:"Amount"`
	Date      *ISO8601Time `json:"Date"`
	Debit     *string      `json:"Debit"`
	Credit    *string      `json:"Credit"`
	IsSystem  *bool        `json:"IsSystem"`
}

type PurchaseAttachmentLine struct {
	ID          *string `json:"ID"`
	ContentType *string `json:"ContentType"`
	IsDefault   *bool   `json:"IsDefault"`
	FileName    *string `json:"FileName"`
	DownloadUrl *string `json:"DownloadUrl"`
}

type PurchaseInventoryMovementLine struct {
	TaskID    *string      `json:"TaskID"`
	ProductID *string      `json:"ProductID"`
	Date      *ISO8601Time `json:"Date"`
	COGS      *float64     `json:"COGS"`
	AdditionalFields
}

type AttachmentLine struct {
	ID          *string `json:"ID"`
	ContentType *string `json:"ContentType"`
	IsDefault   *bool   `json:"IsDefault"`
	FileName    *string `json:"FileName"`
	DownloadUrl *string `json:"DownloadUrl"`
}

type InventoryMovementLine struct {
	TaskID    string      `json:"TaskID"`
	ProductID string      `json:"ProductID"`
	Date      ISO8601Time `json:"Date"`
	COGS      float64     `json:"COGS"`
	AdditionalFields
}

type AdditionalFields struct {
	ProductHeight        *float64 `json:"ProductHeight"`
	ProductLength        *float64 `json:"ProductLength"`
	ProductWeight        *float64 `json:"ProductWeight"`
	ProductWidth         *float64 `json:"ProductWidth"`
	WeightUnits          *string  `json:"WeightUnits"`
	DimensionsUnits      *string  `json:"DimensionsUnits"`
	ProductCustomField1  *string  `json:"ProductCustomField1"`
	ProductCustomField2  *string  `json:"ProductCustomField2"`
	ProductCustomField3  *string  `json:"ProductCustomField3"`
	ProductCustomField4  *string  `json:"ProductCustomField4"`
	ProductCustomField5  *string  `json:"ProductCustomField5"`
	ProductCustomField6  *string  `json:"ProductCustomField6"`
	ProductCustomField7  *string  `json:"ProductCustomField7"`
	ProductCustomField8  *string  `json:"ProductCustomField8"`
	ProductCustomField9  *string  `json:"ProductCustomField9"`
	ProductCustomField10 *string  `json:"ProductCustomField10"`
}

type StockLine struct {
	Quantity             float64      `json:"Quantity,omitempty"`
	UnitCost             float64      `json:"UnitCost,omitempty"`
	ProductID            *string      `json:"ProductID,omitempty"`
	SKU                  *string      `json:"SKU,omitempty"`
	ProductName          *string      `json:"ProductName,omitempty"`
	LocationID           *string      `json:"LocationID,omitempty"`
	Location             *string      `json:"Location,omitempty"`
	BatchSN              *string      `json:"BatchSN,omitempty"`
	ExpiryDate           *ISO8601Time `json:"ExpiryDate,omitempty"`
	ReceivedDate         *ISO8601Time `json:"ReceivedDate,omitempty"`
	Comments             *string      `json:"Comments,omitempty"`
	ProductLength        *float64     `json:"ProductLength,omitempty"`
	ProductWidth         *float64     `json:"ProductWidth,omitempty"`
	ProductHeight        *float64     `json:"ProductHeight,omitempty"`
	ProductWeight        *float64     `json:"ProductWeight,omitempty"`
	WeightUnits          *string      `json:"WeightUnits,omitempty"`
	DimensionsUnits      *string      `json:"DimensionsUnits,omitempty"`
	ProductCustomField1  *string      `json:"ProductCustomField1,omitempty"`
	ProductCustomField2  *string      `json:"ProductCustomField2,omitempty"`
	ProductCustomField3  *string      `json:"ProductCustomField3,omitempty"`
	ProductCustomField4  *string      `json:"ProductCustomField4,omitempty"`
	ProductCustomField5  *string      `json:"ProductCustomField5,omitempty"`
	ProductCustomField6  *string      `json:"ProductCustomField6,omitempty"`
	ProductCustomField7  *string      `json:"ProductCustomField7,omitempty"`
	ProductCustomField8  *string      `json:"ProductCustomField8,omitempty"`
	ProductCustomField9  *string      `json:"ProductCustomField9,omitempty"`
	ProductCustomField10 *string      `json:"ProductCustomField10,omitempty"`
	Barcode              *string      `json:"Barcode,omitempty"`
	Unit                 *string      `json:"Unit,omitempty"`
	CostingMethod        *string      `json:"CostingMethod,omitempty"`
}

type TransactionStockLine struct {
	ID            *string      `json:"ID"`
	Debit         *string      `json:"Debit"`
	Credit        *string      `json:"Credit"`
	Amount        *float64     `json:"Amount"`
	EffectiveDate *ISO8601Time `json:"EffectiveDate"`
}

type Location struct {
	ID         *string `json:"ID"`
	Name       *string `json:"Name"`
	IsDefault  *bool   `json:"IsDefault"`
	Deprecated *bool   `json:"Deprecated"`
	// Bins                 *[]Bin  `json:"Bins"`
	FixedAssetsLocation  *bool   `json:"FixedAssetsLocation"`
	ParentID             *string `json:"ParentID"`
	ReferenceCount       *int    `json:"ReferenceCount"`
	AddressLine1         *string `json:"AddressLine1"`
	AddressLine2         *string `json:"AddressLine2"`
	AddressCitySuburb    *string `json:"AddressCitySuburb"`
	AddressStateProvince *string `json:"AddressStateProvince"`
	AddressZipPostCode   *string `json:"AddressZipPostCode"`
	AddressCountry       *string `json:"AddressCountry"`
	PickZones            *string `json:"PickZones"`
	IsShopfloor          *bool   `json:"IsShopfloor"`
	IsCoMan              *bool   `json:"IsCoMan"`
	IsStaging            *bool   `json:"IsStaging"`
}

type IDName struct {
	ID   *string `json:"ID"`
	Name *string `json:"Name"`
}

type StockTransfer struct {
	TaskID               *string                  `json:"TaskID"`
	Status               *string                  `json:"Status"`
	From                 *string                  `json:"From"`
	FromLocation         *string                  `json:"FromLocation"`
	To                   *string                  `json:"To"`
	ToLocation           *string                  `json:"ToLocation"`
	Number               *string                  `json:"Number"`
	CostDistributionType *string                  `json:"CostDistributionType"`
	InTransitAccount     *string                  `json:"InTransitAccount"`
	DepartureDate        *ISO8601Time             `json:"DepartureDate"`
	CompletionDate       *ISO8601Time             `json:"CompletionDate"`
	RequiredByDate       *ISO8601Time             `json:"RequiredByDate"`
	Reference            *string                  `json:"Reference"`
	Lines                *[]StockTransferLine     `json:"Lines"`
	ManualJournals       *[]PurchaseManualJournal `json:"ManualJournals"`
	SkipOrder            *bool                    `json:"SkipOrder"`
	Order                *StockTransferOrder      `json:"Order"`
}

type StockTransferLine struct {
	ProductID            *string      `json:"ProductID"`
	SKU                  *string      `json:"SKU"`
	ProductName          *string      `json:"ProductName"`
	QuantityOnHand       *float64     `json:"QuantityOnHand"`
	QuantityAvailable    *float64     `json:"QuantityAvailable"`
	TransferQuantity     *float64     `json:"TransferQuantity"`
	BatchSN              *string      `json:"BatchSN"`
	ExpiryDate           *ISO8601Time `json:"ExpiryDate"`
	Comments             *string      `json:"Comments"`
	ProductLength        *float64     `json:"ProductLength"`
	ProductWidth         *float64     `json:"ProductWidth"`
	ProductHeight        *float64     `json:"ProductHeight"`
	ProductWeight        *float64     `json:"ProductWeight"`
	WeightUnits          *string      `json:"WeightUnits"`
	DimensionsUnits      *string      `json:"DimensionsUnits"`
	ProductCustomField1  *string      `json:"ProductCustomField1"`
	ProductCustomField2  *string      `json:"ProductCustomField2"`
	ProductCustomField3  *string      `json:"ProductCustomField3"`
	ProductCustomField4  *string      `json:"ProductCustomField4"`
	ProductCustomField5  *string      `json:"ProductCustomField5"`
	ProductCustomField6  *string      `json:"ProductCustomField6"`
	ProductCustomField7  *string      `json:"ProductCustomField7"`
	ProductCustomField8  *string      `json:"ProductCustomField8"`
	ProductCustomField9  *string      `json:"ProductCustomField9"`
	ProductCustomField10 *string      `json:"ProductCustomField10"`
}

type StockTransferOrder struct {
	Status *string              `json:"Status"`
	Lines  *[]StockTransferLine `json:"Lines"`
}
