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
// AchDetails struct for AchDetails
type AchDetails struct {
	// The bank account number (without separators).
	BankAccountNumber string `json:"bankAccountNumber"`
	// The bank routing number of the account. The field value is `nil` in most cases.
	BankLocationId string `json:"bankLocationId,omitempty"`
	// Encrypted bank account number. The bank account number (without separators).
	EncryptedBankAccountNumber string `json:"encryptedBankAccountNumber,omitempty"`
	// Encrypted location id. The bank routing number of the account. The field value is `nil` in most cases.
	EncryptedBankLocationId string `json:"encryptedBankLocationId,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName string `json:"ownerName,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	RecurringDetailReference string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId string `json:"storedPaymentMethodId,omitempty"`
	// **ach**
	Type string `json:"type,omitempty"`
}