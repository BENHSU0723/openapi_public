package models

import "time"

// It's self-defined struct for 5glan multicast used
// SMF can subs UDM to get 5glan related group config data changed

type Vn5gGroupConfigSubscription struct {

	// Internal group id
	GroupId string `json:"groupId,omitempty"`

	ExternalGroupId string `json:"externalGroupId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	SingleNssai Snssai `json:"singleNssai,omitempty"`

	// - record the current vn group members who have build up Group Used PDU session, not all members
	// - it's used for 5glan multicast to find SM context
	CurrMembers map[string]string // supi as key, the ref string of group used SM context as value

	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`

	// String providing an URI formatted according to RFC 3986.
	OriginalCallbackReference string `json:"originalCallbackReference,omitempty"`

	MonitoredResourceUris []string `json:"monitoredResourceUris"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	UniqueSubscription bool `json:"uniqueSubscription,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
