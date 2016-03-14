package datasourcesportscheck

// Product struct.
type Product struct {
	UID          string `json:"uid"`
	EAN          string `json:"ean"`
	ProductID    string `json:"productid"`
	ProductName  string `json:"productname"`
	ProductShort string `json:"productshort"`
	ProductLong  string `json:"productlong"`

	ProductCategory string `json:"productcategory"`
	Colour          string `json:"colour"`
	Size            string `json:"size"`
	Brand           string `json:"brand"`
	Price           string `json:"price"`
	PriceOld        string `json:"priceold"`

	ImageMain string `json:"imagemain"`
	Image01   string `json:"image01"`
	Image02   string `json:"image02"`
	Image03   string `json:"image03"`
	Image04   string `json:"image04"`
}


// ProductRaw is the data as it comes from the ERP system.
// We need to transform it in Product type.
type ProductRaw struct {
	EAN              string `json:"ean"`
	ProductID        string `json:"productid"`
	StyleID          string `json:"styleid"`
	Produktname      string `json:"produktname"`
	ProductShortCopy string `json:"productshortcopy"`
	ProductCategory  string `json:"productcategory"`
	Geschlecht       string `json:"geschlecht"`
	Marke            string `json:"marke"`
	Farbe            string `json:"farbe"`
	Größe            string `json:"größe"`
    LongCopy       string `json:"longcopy"`
    BulletPoints        string `json:"bulletpoints"`
    BildURL            string `json:"bildurl"`
    ZusätzlicherBildlink2   string `json:"zusätzlicherbildlink2"`
    ZusätzlicherBildlink3   string `json:"zusätzlicherbildlink3"`
    ZusätzlicherBildlink4   string `json:"zusätzlicherbildlink4"`
    ZusätzlicherBildlink5   string `json:"zusätzlicherbildlink5"`
    ZusätzlicherBildlink6   string `json:"zusätzlicherbildlink6"`
    ProduktURL   string `json:"produkturl"`
}
