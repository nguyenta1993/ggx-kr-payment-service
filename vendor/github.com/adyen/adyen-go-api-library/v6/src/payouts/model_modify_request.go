/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

// ModifyRequest struct for ModifyRequest
type ModifyRequest struct {
	// This field contains additional data, which may be required for a particular payout request.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The PSP reference received in the `/submitThirdParty` response.
	OriginalReference string `json:"originalReference"`
}
