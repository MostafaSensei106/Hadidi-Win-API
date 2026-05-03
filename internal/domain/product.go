package domain

import (
	"context"
	"errors"
)

type ProductType string
type Material string
type Brand string
type GlassType string
type Color string
type DoorType string
type WindowType string
type InsectScreenType string

const (
	ProductTypeSimple       ProductType = "simple"
	ProductTypeConfigurable ProductType = "configurable"
)

const (
	MaterialAluminum Material = "aluminum"
)

const (
	BrandPolywin Brand = "polywin"
	BrandKimit   Brand = "kimit"
)

const (
	GlassSingle GlassType = "single"
	GlassDouble GlassType = "double"
	GlassFrost  GlassType = "frost"
)

const (
	ColorBlack Color = "black"
	ColorWhite Color = "white"
	ColorSand  Color = "sand"
	ColorMetal Color = "metal"
)

const (
	TiltWindow    WindowType = "tilt"
	SlidingWindow WindowType = "sliding"
	WoodWindow    WindowType = "wood"
	HingedWindow  WindowType = "Hinged"
)

const (
	HingedDorr  DoorType = "Hinged"
	SlidingDoor DoorType = "Sliding"
	SwingDoor   DoorType = "Swing"
)

type Product struct {
	BaseEntity
	ID          string
	Name        string
	Description string
	Price       int64
	Type        ProductType
}

type ProductConfigSchema struct {
	ProductID          string
	AvailableBrands    []Brand
	AvailableMaterials []Material
	AvailableGlass     []GlassType
	AvailableColors    []string
	SupportsWindows    bool
	SupportsDoors      bool
}

type ProductConfig struct {
	ProductID  string
	Brand      Brand
	Material   Material
	GlassType  GlassType
	WindowType WindowType
	DoorType   DoorType
	InsectType InsectScreenType
	Width      float32
	Height     float32
	Color      string
	Qty        uint32
}

func (c *ProductConfig) Validate(schema *ProductConfigSchema) error {
	if c == nil {
		return errors.New("config is required")
	}

	if c.Width <= 0 || c.Height <= 0 {
		return errors.New("invalid dimensions")
	}

	if c.Qty == 0 {
		return errors.New("quantity must be > 0")
	}

	valid := false
	for _, b := range schema.AvailableBrands {
		if c.Brand == b {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("invalid brand")
	}

	return nil
}

func CalculatePrice(config *ProductConfig) int64 {
	if config == nil {
		return 0
	}
	var price int64 = 1000
	area := config.Width * config.Height
	price += int64(area * 10)
	switch config.GlassType {
	case GlassDouble:
		price += 200
	case GlassFrost:
		price += 150
	}

	if config.Material == MaterialAluminum {
		price += 100
	}

	return price * int64(config.Qty)
}

type ProductRepository interface {
	Get(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	Create(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]*Product, error)
}

type ConfigSchemaRepository interface {
	GetByProductID(ctx context.Context, productID string) (*ProductConfigSchema, error)
}

type ProductUsecase interface {
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context) ([]*Product, error)
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type ConfigUsecase interface {
	GetSchema(ctx context.Context, productID string) (*ProductConfigSchema, error)
	ValidateConfig(ctx context.Context, config *ProductConfig) error
	Calculate(ctx context.Context, config *ProductConfig) int64
}
