/*
 * Application API
 *
 * Nexmo provides an Application API to allow management of your Nexmo Applications.  This API is backwards compatible with version 1. Applications created using version 1 of the API can also be managed using version 2 (this version) of the API. 
 *
 * API version: 2.0.5
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package application
// ApplicationResponseCapabilitiesMessages Messages / Dispatch related configuration
type ApplicationResponseCapabilitiesMessages struct {
	Webhooks ApplicationResponseCapabilitiesMessagesWebhooks `json:"webhooks,omitempty"`
}
