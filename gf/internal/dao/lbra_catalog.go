// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"pdmaner-test/internal/dao/internal"
)

// internalLbraCatalogDao is internal type for wrapping internal DAO implements.
type internalLbraCatalogDao = *internal.LbraCatalogDao

// lbraCatalogDao is the data access object for table lbra.
// You can define custom methods on it to extend its functionality as you wish.
type lbraCatalogDao struct {
	internalLbraCatalogDao
}

var (
	// LbraCatalog is globally public accessible object for table lbra operations.
	LbraCatalog = lbraCatalogDao{
		internal.NewLbraCatalogDao(),
	}
)

// Fill with you ideas below.
