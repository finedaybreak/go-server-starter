package repo

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QueryOption func(*gorm.DB) *gorm.DB

func ApplyQueryOptions(db *gorm.DB, opts ...QueryOption) *gorm.DB {
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}

// Order specify order when retrieving records from database
//
//	db.Order("name DESC")
//	db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
//	db.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
//		{Column: clause.Column{Name: "name"}, Desc: true},
//		{Column: clause.Column{Name: "age"}, Desc: true},
//	}})
func Order(order any) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Where add conditions
//
// See the [docs] for details on the various formats that where clauses can take. By default, where clauses chain with AND.
//
//	// Find the first user with name jinzhu
//	db.Where("name = ?", "jinzhu").First(&user)
//	// Find the first user with name jinzhu and age 20
//	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
//	// Find the first user with name jinzhu and age not equal to 20
//	db.Where("name = ?", "jinzhu").Where("age <> ?", "20").First(&user)
//
// [docs]: https://gorm.io/docs/query.html#Conditions
func Where(query any, args ...any) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

// WhereIf conditionally adds a where clause based on a predicate
// Use this for complex conditions where you control the logic
//
//	WhereIf(status != nil, "status = ?", *status)
func WhereIf(condition bool, query string, args ...any) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if condition {
			return db.Where(query, args...)
		}
		return db
	}
}

// WherePtr adds where clause if pointer is not nil (works with any pointer type)
//
//	var status *int = nil  // skipped
//	var name *string = ptr("test")  // applied
//	WherePtr("status = ?", status)
//	WherePtr("name = ?", name)
func WherePtr[T any](query string, arg *T) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if arg != nil {
			return db.Where(query, *arg)
		}
		return db
	}
}

// WherePtrNonEmpty adds where clause if pointer is not nil and value is non-empty
// Only works with comparable types, skips zero values
//
//	var name *string = ptr("")  // skipped (empty string)
//	var count *int = ptr(0)     // skipped (zero value)
func WherePtrNonEmpty[T comparable](query string, arg *T) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		var zero T
		if arg != nil && *arg != zero {
			return db.Where(query, *arg)
		}
		return db
	}
}

// WhereAutoLike ( name LIKE ?  %string% )
func WhereAutoLike(name string, arg *string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if arg != nil && *arg != "" {
			return db.Where(fmt.Sprintf("%s LIKE ?", name), "%"+*arg+"%")
		}
		return db
	}
}

// WhereAutoLikePrefix ( name LIKE ?  string% )
func WhereAutoLikePrefix(name string, arg *string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if arg != nil && *arg != "" {
			return db.Where(fmt.Sprintf("%s LIKE ?", name), *arg+"%")
		}
		return db
	}
}

// WhereAutoLikeSuffix ( name LIKE ?  %string )
func WhereAutoLikeSuffix(name string, arg *string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if arg != nil && *arg != "" {
			return db.Where(fmt.Sprintf("%s LIKE ?", name), "%"+*arg)
		}
		return db
	}
}

// Preload preload associations with given conditions
//
//	// get all users, and preload all non-cancelled orders
//	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func Preload(query string, args ...any) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(query, args...)
	}
}

// PreloadBatch preload multiple associations
//
//	db.PreloadBatch("Orders", "Addresses")
func PreloadBatch(query ...string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		for _, q := range query {
			db = db.Preload(q)
		}
		return db
	}
}

// Paginate paginate records with page and pageSize
//
//	db.Paginate(1, 10)
func Paginate(page, pageSize int) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

// ForUpdate add FOR UPDATE clause to the query
func ForUpdate() QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
}
