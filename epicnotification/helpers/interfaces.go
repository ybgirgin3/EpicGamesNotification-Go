package helpers

import "time"

// CatalogInterface represents the interface for the Catalog data in the JSON response
type CatalogInterface interface {
	GetSearchStore() SearchStoreInterface
}

// SearchStoreInterface represents the interface for the SearchStore data in the JSON response
type SearchStoreInterface interface {
	GetElements() []ElementInterface
}

// ElementInterface represents the interface for the Element data in the JSON response
type ElementInterface interface {
	GetTitle() string
	GetID() string
	GetNamespace() string
	GetDescription() string
	GetEffectiveDate() time.Time
	GetOfferType() string
	GetExpiryDate() interface{}
}

// ResponseInterface represents the interface for the entire JSON response
type ResponseInterface interface {
	GetData() DataInterface
}

// DataInterface represents the interface for the Data data in the JSON response
type DataInterface interface {
	GetCatalog() CatalogInterface
}

// Implement the methods for each interface in the corresponding structs
// ...

// Now you can use these interfaces in your code to work with the JSON response in a more abstract way
