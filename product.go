package main

import (
	"encoding/json"
	"io"
)

// The structure for our API product
type Product struct {
	CategoryID                  string `json:""`
	ItemID                      string `json:""`
	CustomLabelSKU              string `json:""`
	Title                       string `json:""`
	Format                      string `json:""`
	VariationSet                string `json:""`
	SiteID                      string `json:""`
	Reserved1Donoteditordelete  string `json:""`
	Brand                       string `json:""`
	Size                        string `json:""`
	SizeType                    string `json:""`
	SleeveLength                string `json:""`
	Color                       string `json:""`
	Type                        string `json:""`
	Department                  string `json:""`
	Inseam                      string `json:""`
	Style                       string `json:""`
	OuterShellMaterial          string `json:""`
	USShoeSize                  string `json:""`
	Pattern                     string `json:""`
	Theme                       string `json:""`
	Closure                     string `json:""`
	Accents                     string `json:""`
	Features                    string `json:""`
	FabricType                  string `json:""`
	Material                    string `json:""`
	Fit                         string `json:""`
	CollarStyle                 string `json:""`
	Character                   string `json:""`
	ProductLine                 string `json:""`
	Neckline                    string `json:""`
	Vintage                     string `json:""`
	NeckSize                    string `json:""`
	TopCuffStyle                string `json:""`
	FabricWash                  string `json:""`
	PerformanceActivity         string `json:""`
	WaistSize                   string `json:""`
	LiningMaterial              string `json:""`
	InsulationMaterial          string `json:""`
	SwimBottomStyle             string `json:""`
	JacketCoatLength            string `json:""`
	JacketCut                   string `json:""`
	JacketFrontButtonStyle      string `json:""`
	NumberofPieces              string `json:""`
	JacketLapelStyle            string `json:""`
	Customized                  string `json:""`
	StyleCode                   string `json:""`
	EUShoeSize                  string `json:""`
	ShoeWidth                   string `json:""`
	ShoeShaftStyle              string `json:""`
	Silhouette                  string `json:""`
	UpperMaterial               string `json:""`
	UKShoeSize                  string `json:""`
	AUShoeSize                  string `json:""`
	BagWidth                    string `json:""`
	BagHeight                   string `json:""`
	BagDepth                    string `json:""`
	Occasion                    string `json:""`
	Model                       string `json:""`
	Season                      string `json:""`
	CharacterFamily             string `json:""`
	CountryRegionofManufacture  string `json:""`
	GraphicPrint                string `json:""`
	Personalize                 string `json:""`
	Handmade                    string `json:""`
	GarmentCare                 string `json:""`
	ChestSize                   string `json:""`
	YearManufactured            string `json:""`
	CaliforniaProp65Warning     string `json:""`
	MPN                         string `json:""`
	PersonalizationInstructions string `json:""`
	ShirtLength                 string `json:""`
	PlacketType                 string `json:""`
	SleeveType                  string `json:""`
	Rise                        string `json:""`
	Distressed                  string `json:""`
	PocketType                  string `json:""`
	SuspenderButtonsBeltLoops   string `json:""`
	LegOpening                  string `json:""`
	WarmthWeight                string `json:""`
	KnitStyle                   string `json:""`
	CompressionLevel            string `json:""`
	FrontType                   string `json:""`
	JacketVentStyle             string `json:""`
	JacketSleeveButtonStyle     string `json:""`
	LegStyle                    string `json:""`
	MadetoMeasure               string `json:""`
	InsoleMaterial              string `json:""`
	OutsoleMaterial             string `json:""`
	CleatType                   string `json:""`
	CushioningLevel             string `json:""`
	Signed                      string `json:""`
	MidsoleType                 string `json:""`
	HeeltoToeDrop               string `json:""`
	Pronation                   string `json:""`
	StudShape                   string `json:""`
	ToeShape                    string `json:""`
	CalfWidth                   string `json:""`
	ItemLength                  string `json:""`
	ItemWidth                   string `json:""`
	HardwareMaterial            string `json:""`
	HandleStyle                 string `json:""`
	UPC                         string `json:""`
	Wash                        string `json:""`
	DenimTrends                 string `json:""`
	ManufacturerColor           string `json:""`
	Fastening                   string `json:""`
	Width                       string `json:""`
	EuroSize                    string `json:""`
	FeaturedGraphics            string `json:""`
	Personalized                string `json:""`
	UPS                         string `json:""`
	ShoeHeight                  string `json:""`
	HandleStrapMaterial         string `json:""`
	InsidePockets               string `json:""`
	OutsidePockets              string `json:""`
	HardSoft                    string `json:""`
	AdjustmentType              string `json:""`
	Length                      string `json:""`
	SubStyle                    string `json:""`
	Gender                      string `json:""`
	Collar                      string `json:""`
	CuffStyle                   string `json:""`
	Sleeve                      string `json:""`
	USShoeSizeWomens            string `json:""`
	ExactMaterial               string `json:""`
}

// A collection of Product objects
type Products []*Product

func (p *Product) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// func getProducts() Products {
// 	return productList
// }

// var productList = []*Product{
// 	&Product{
// 		ID:            1,
// 		Description:   "7 Diamonds - Closer To You Fannel - Blue - size M",
// 		PurchaseDate:  "2020-10-03",
// 		SalesDate:     "2021-03-25",
// 		ShipDate:      "2021-03-26",
// 		PurchasePrice: 12.75,
// 		SalesPrice:    49,
// 	},
// 	&Product{
// 		ID:            2,
// 		Description:   "7 For All Mankind - Slimmy Slim - Everglade - size 34",
// 		PurchaseDate:  "2021-01-21",
// 		SalesDate:     "2021-07-30",
// 		ShipDate:      "2021-07-31",
// 		PurchasePrice: 22.75,
// 		SalesPrice:    56.90,
// 	},
// }
