package helpers

import (
	"fmt"
	"log"
	"math"
	"net/smtp"
	"os"
	"pood/models"
)

func SendEmail(order models.UserOrder) error {
	// finish the logic send email to the client postbox

	log.Println(order)
	var orderString = fmt.Sprintf("\nOrder Id - %v\n", order.OrderId)
	for _, position := range order.Positions {
		orderString += template(position)
	}
	log.Println(orderString)
	from := os.Getenv("POOD_ADMIN_EMAIL")
	password := os.Getenv("POOD_EMAIL_PASSWORD")

	toEmailAddress := "vejlevas@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Danja pidor lovi zakaz svoj\n"
	body := "Thanks for your order! Your order is being processed. For any information please reach out the Daniil Batjkovichj by email" + orderString
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err2 := smtp.SendMail(address, auth, from, to, message)
	if err2 != nil {
		log.Println(err2)
		return err2
	}
	return nil
}

func template(pos *models.Product) string {
	var template = fmt.Sprintf("\n\tPrice - %v,\n\tArticle - %v,\n\tSupplier - %v,\n\tSupplier Price Num - %v,\n\tBrand - %v,\n\tCurrency - %v,\n\tCurrency rate - %v,\n\tDelivery days - %v,\n\tWeight - %v,\n\tQuantity - %v,\n\tQuantity Price - %v\n",
		roundFloat(pos.Price, 2), pos.Article, pos.Supplier, pos.SupplierPriceNum, pos.Brand, pos.Currency, pos.CurrencyRate, pos.Delivery, roundFloat(pos.Weight, 3), pos.Quantity, roundFloat(pos.ProductQuantityPrice, 2))
	return template
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
