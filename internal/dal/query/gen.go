// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q            = new(Query)
	AddressBook  *addressBook
	Category     *category
	Dish         *dish
	DishFlavor   *dishFlavor
	Employee     *employee
	Order        *order
	OrderDetail  *orderDetail
	Setmeal      *setmeal
	SetmealDish  *setmealDish
	ShoppingCart *shoppingCart
	User         *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AddressBook = &Q.AddressBook
	Category = &Q.Category
	Dish = &Q.Dish
	DishFlavor = &Q.DishFlavor
	Employee = &Q.Employee
	Order = &Q.Order
	OrderDetail = &Q.OrderDetail
	Setmeal = &Q.Setmeal
	SetmealDish = &Q.SetmealDish
	ShoppingCart = &Q.ShoppingCart
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		AddressBook:  newAddressBook(db, opts...),
		Category:     newCategory(db, opts...),
		Dish:         newDish(db, opts...),
		DishFlavor:   newDishFlavor(db, opts...),
		Employee:     newEmployee(db, opts...),
		Order:        newOrder(db, opts...),
		OrderDetail:  newOrderDetail(db, opts...),
		Setmeal:      newSetmeal(db, opts...),
		SetmealDish:  newSetmealDish(db, opts...),
		ShoppingCart: newShoppingCart(db, opts...),
		User:         newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AddressBook  addressBook
	Category     category
	Dish         dish
	DishFlavor   dishFlavor
	Employee     employee
	Order        order
	OrderDetail  orderDetail
	Setmeal      setmeal
	SetmealDish  setmealDish
	ShoppingCart shoppingCart
	User         user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		AddressBook:  q.AddressBook.clone(db),
		Category:     q.Category.clone(db),
		Dish:         q.Dish.clone(db),
		DishFlavor:   q.DishFlavor.clone(db),
		Employee:     q.Employee.clone(db),
		Order:        q.Order.clone(db),
		OrderDetail:  q.OrderDetail.clone(db),
		Setmeal:      q.Setmeal.clone(db),
		SetmealDish:  q.SetmealDish.clone(db),
		ShoppingCart: q.ShoppingCart.clone(db),
		User:         q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		AddressBook:  q.AddressBook.replaceDB(db),
		Category:     q.Category.replaceDB(db),
		Dish:         q.Dish.replaceDB(db),
		DishFlavor:   q.DishFlavor.replaceDB(db),
		Employee:     q.Employee.replaceDB(db),
		Order:        q.Order.replaceDB(db),
		OrderDetail:  q.OrderDetail.replaceDB(db),
		Setmeal:      q.Setmeal.replaceDB(db),
		SetmealDish:  q.SetmealDish.replaceDB(db),
		ShoppingCart: q.ShoppingCart.replaceDB(db),
		User:         q.User.replaceDB(db),
	}
}

type queryCtx struct {
	AddressBook  *addressBookDo
	Category     *categoryDo
	Dish         *dishDo
	DishFlavor   *dishFlavorDo
	Employee     *employeeDo
	Order        *orderDo
	OrderDetail  *orderDetailDo
	Setmeal      *setmealDo
	SetmealDish  *setmealDishDo
	ShoppingCart *shoppingCartDo
	User         *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AddressBook:  q.AddressBook.WithContext(ctx),
		Category:     q.Category.WithContext(ctx),
		Dish:         q.Dish.WithContext(ctx),
		DishFlavor:   q.DishFlavor.WithContext(ctx),
		Employee:     q.Employee.WithContext(ctx),
		Order:        q.Order.WithContext(ctx),
		OrderDetail:  q.OrderDetail.WithContext(ctx),
		Setmeal:      q.Setmeal.WithContext(ctx),
		SetmealDish:  q.SetmealDish.WithContext(ctx),
		ShoppingCart: q.ShoppingCart.WithContext(ctx),
		User:         q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
