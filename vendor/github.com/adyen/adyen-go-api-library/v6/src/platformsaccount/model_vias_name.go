/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v6/createAccountHolder ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount
// ViasName struct for ViasName
type ViasName struct {
	// The first name.
	FirstName string `json:"firstName"`
	// The gender. >The following values are permitted: `MALE`, `FEMALE`, `UNKNOWN`.
	Gender string `json:"gender"`
	// The name's infix, if applicable. >A maximum length of twenty (20) characters is imposed.
	Infix string `json:"infix,omitempty"`
	// The last name.
	LastName string `json:"lastName"`
}