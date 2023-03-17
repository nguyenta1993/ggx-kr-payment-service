/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v64/authorise ```
 *
 * API version: 64
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments
// AdditionalDataRiskStandalone struct for AdditionalDataRiskStandalone
type AdditionalDataRiskStandalone struct {
	// Shopper's country of residence in the form of ISO standard 3166 2-character country codes.
	PayPalCountryCode string `json:"PayPal.CountryCode,omitempty"`
	// Shopper's email.
	PayPalEmailId string `json:"PayPal.EmailId,omitempty"`
	// Shopper's first name.
	PayPalFirstName string `json:"PayPal.FirstName,omitempty"`
	// Shopper's last name.
	PayPalLastName string `json:"PayPal.LastName,omitempty"`
	// Unique PayPal Customer Account identification number. Character length and limitations: 13 single-byte alphanumeric characters.
	PayPalPayerId string `json:"PayPal.PayerId,omitempty"`
	// Shopper's phone number.
	PayPalPhone string `json:"PayPal.Phone,omitempty"`
	// Allowed values: * **Eligible** — Merchant is protected by PayPal's Seller Protection Policy for Unauthorized Payments and Item Not Received.  * **PartiallyEligible** — Merchant is protected by PayPal's Seller Protection Policy for Item Not Received.  * **Ineligible** — Merchant is not protected under the Seller Protection Policy.
	PayPalProtectionEligibility string `json:"PayPal.ProtectionEligibility,omitempty"`
	// Unique transaction ID of the payment.
	PayPalTransactionId string `json:"PayPal.TransactionId,omitempty"`
	// Raw AVS result received from the acquirer, where available. Example: D
	AvsResultRaw string `json:"avsResultRaw,omitempty"`
	// The Bank Identification Number of a credit card, which is the first six digits of a card number. Required for [tokenized card request](https://docs.adyen.com/risk-management/standalone-risk#tokenised-pan-request).
	Bin string `json:"bin,omitempty"`
	// Raw CVC result received from the acquirer, where available. Example: 1
	CvcResultRaw string `json:"cvcResultRaw,omitempty"`
	// Unique identifier or token for the shopper's card details.
	RiskToken string `json:"riskToken,omitempty"`
	// A Boolean value indicating whether 3DS authentication was completed on this payment. Example: true
	ThreeDAuthenticated string `json:"threeDAuthenticated,omitempty"`
	// A Boolean value indicating whether 3DS was offered for this payment. Example: true
	ThreeDOffered string `json:"threeDOffered,omitempty"`
	// Required for PayPal payments only. The only supported value is: **paypal**.
	TokenDataType string `json:"tokenDataType,omitempty"`
}