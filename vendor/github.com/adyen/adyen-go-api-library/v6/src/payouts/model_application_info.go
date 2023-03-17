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

// ApplicationInfo struct for ApplicationInfo
type ApplicationInfo struct {
	AdyenLibrary             *CommonField              `json:"adyenLibrary,omitempty"`
	AdyenPaymentSource       *CommonField              `json:"adyenPaymentSource,omitempty"`
	ExternalPlatform         *ExternalPlatform         `json:"externalPlatform,omitempty"`
	MerchantApplication      *CommonField              `json:"merchantApplication,omitempty"`
	MerchantDevice           *MerchantDevice           `json:"merchantDevice,omitempty"`
	ShopperInteractionDevice *ShopperInteractionDevice `json:"shopperInteractionDevice,omitempty"`
}
