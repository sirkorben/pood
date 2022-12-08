package db

import (
	"fmt"
	"log"
	"pood/models"
)

func AddProductToShoppingCart(userId int, product models.Product) error {
	orderId, err := GetNonConfirmedOrderId(userId)
	if err != nil {
		return err
	}
	// TODO: check if same order being added one more time - just increase quantity
	_, err = DB.Exec("insert into products_ordered (order_id, price, article, supplier, supplier_price_num, brand, currency, currency_rate, delivery, weight, quantity, product_quantity_price) values (?,?,?,?,?,?,?,?,?,?,?,?)",
		orderId, product.Price, product.Article, product.Supplier, product.SupplierPriceNum, product.Brand, product.Currency, product.CurrencyRate, product.Delivery, product.Weight*float64(product.Quantity), product.Quantity, product.Price*float64(product.Quantity))
	if err != nil {
		log.Println("INSERTING PRODUCT ERROR \t ", err)
		return err
	}
	log.Printf("User with [Id - %v] has added product [article - %v] to order with [oreder_id - %v]", userId, product.Article, orderId)
	return nil
}

func GetProductsUnderNonConfirmedOrderId(userId int) ([]*models.Product, error) {
	orderId, err := GetNonConfirmedOrderId(userId)
	if err != nil {
		return nil, err
	}
	rows, err := DB.Query("SELECT * FROM products_ordered WHERE order_id = ? ORDER BY product_quantity_price DESC", orderId)
	if err != nil {
		fmt.Println("3\t", err)
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err = rows.Scan(&orderId, &product.Price, &product.Article, &product.Supplier, &product.SupplierPriceNum, &product.Brand, &product.Currency, &product.CurrencyRate, &product.Delivery,
			&product.Weight, &product.Quantity, &product.ProductQuantityPrice)
		if err != nil {
			fmt.Println("4\t", err)
			return nil, err
		}
		products = append(products, product)
	}
	if err != nil {
		fmt.Println("t5\t", err)
		return nil, err
	}
	return products, nil
}

func GetOrderedProductsByConfirmedOrderId(orderId string) []models.Product {

	return nil
}
