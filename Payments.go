package api

type LabeledPrice struct {
	// This object represents a portion of the price for goods or services.
	label  string //Portion label
	amount int    //Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

type Invoice struct {
	// This object contains basic information about an invoice.
	title           string //Product name
	description     string //Product description
	start_parameter string //Unique bot deep-linking parameter that can be used to generate this invoice
	currency        string //Three-letter ISO 4217 currency code
	total_amount    int    //Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

type ShippingAddress struct {
	// This object represents a shipping address.
	country_code string //Two-letter ISO 3166-1 alpha-2 country code
	state        string //State, if applicable
	city         string //City
	street_line1 string //First line for the address
	street_line2 string //Second line for the address
	post_code    string //Address post code
}

type OrderInfo struct {
	// This object represents information about an order.
	name             string          //Optional. User name
	phone_number     string          //Optional. User's phone number
	email            string          //Optional. User email
	shipping_address ShippingAddress //Optional. User shipping address
}

type ShippingOption struct {
	// This object represents one shipping option.
	id     string         //Shipping option identifier
	title  string         //Option title
	prices []LabeledPrice //List of price portions
}

type SuccessfulPayment struct {
	// This object contains basic information about a successful payment.
	currency                   string    //Three-letter ISO 4217 currency code
	total_amount               int       //Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	invoice_payload            string    //Bot specified invoice payload
	shipping_option_id         string    //Optional. Identifier of the shipping option chosen by the user
	order_info                 OrderInfo //Optional. Order information provided by the user
	telegram_payment_charge_id string    //Telegram payment identifier
	provider_payment_charge_id string    //Provider payment identifier
}
