/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v69/payments ```  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=69) to find out what changed in this version!
 *
 * API version: 69
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// AdditionalDataRatepay struct for AdditionalDataRatepay
type AdditionalDataRatepay struct {
	// Amount the customer has to pay each month.
	RatepayInstallmentAmount string `json:"ratepay.installmentAmount,omitempty"`
	// Interest rate of this installment.
	RatepayInterestRate string `json:"ratepay.interestRate,omitempty"`
	// Amount of the last installment.
	RatepayLastInstallmentAmount string `json:"ratepay.lastInstallmentAmount,omitempty"`
	// Calendar day of the first payment.
	RatepayPaymentFirstday string `json:"ratepay.paymentFirstday,omitempty"`
	// Date the merchant delivered the goods to the customer.
	RatepaydataDeliveryDate string `json:"ratepaydata.deliveryDate,omitempty"`
	// Date by which the customer must settle the payment.
	RatepaydataDueDate string `json:"ratepaydata.dueDate,omitempty"`
	// Invoice date, defined by the merchant. If not included, the invoice date is set to the delivery date.
	RatepaydataInvoiceDate string `json:"ratepaydata.invoiceDate,omitempty"`
	// Identification name or number for the invoice, defined by the merchant.
	RatepaydataInvoiceId string `json:"ratepaydata.invoiceId,omitempty"`
}
