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

// FraudResult struct for FraudResult
type FraudResult struct {
	// The total fraud score generated by the risk checks.
	AccountScore int32 `json:"accountScore"`
	// The result of the individual risk checks.
	Results *[]FraudCheckResult `json:"results,omitempty"`
}
