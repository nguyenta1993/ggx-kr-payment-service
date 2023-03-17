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
// ThreeDS2Result struct for ThreeDS2Result
type ThreeDS2Result struct {
	// The `authenticationValue` value as defined in the 3D Secure 2 specification.
	AuthenticationValue string `json:"authenticationValue,omitempty"`
	// The algorithm used by the ACS to calculate the authentication value, only for Cartes Bancaires integrations.
	CavvAlgorithm string `json:"cavvAlgorithm,omitempty"`
	// Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. For possible values, refer to [3D Secure API reference](https://docs.adyen.com/online-payments/3d-secure/api-reference#mpidata).
	ChallengeCancel string `json:"challengeCancel,omitempty"`
	// Specifies a preference for receiving a challenge from the issuer. Allowed values: * `noPreference` * `requestNoChallenge` * `requestChallenge` * `requestChallengeAsMandate` 
	ChallengeIndicator string `json:"challengeIndicator,omitempty"`
	// The `dsTransID` value as defined in the 3D Secure 2 specification.
	DsTransID string `json:"dsTransID,omitempty"`
	// The `eci` value as defined in the 3D Secure 2 specification.
	Eci string `json:"eci,omitempty"`
	// Indicates the exemption type that was applied by the issuer to the authentication, if exemption applied. Allowed values: * `lowValue` * `secureCorporate` * `trustedBeneficiary` * `transactionRiskAnalysis` 
	ExemptionIndicator string `json:"exemptionIndicator,omitempty"`
	// The `messageVersion` value as defined in the 3D Secure 2 specification.
	MessageVersion string `json:"messageVersion,omitempty"`
	// Risk score calculated by Cartes Bancaires Directory Server (DS).
	RiskScore string `json:"riskScore,omitempty"`
	// The `threeDSServerTransID` value as defined in the 3D Secure 2 specification.
	ThreeDSServerTransID string `json:"threeDSServerTransID,omitempty"`
	// The `timestamp` value of the 3D Secure 2 authentication.
	Timestamp string `json:"timestamp,omitempty"`
	// The `transStatus` value as defined in the 3D Secure 2 specification.
	TransStatus string `json:"transStatus,omitempty"`
	// Provides information on why the `transStatus` field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values).
	TransStatusReason string `json:"transStatusReason,omitempty"`
	// The `whiteListStatus` value as defined in the 3D Secure 2 specification.
	WhiteListStatus string `json:"whiteListStatus,omitempty"`
}