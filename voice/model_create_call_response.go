/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.3.2
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice
// CreateCallResponse struct for CreateCallResponse
type CreateCallResponse struct {
	// The unique identifier for this call leg. The UUID is created when your call request is accepted by Nexmo. You use the UUID in all requests for individual live calls
	Uuid string `json:"uuid,omitempty"`
	// The status of the call. [See possible values](https://developer.nexmo.com/voice/voice-api/guides/call-flow#event-objects)
	Status string `json:"status,omitempty"`
	Direction Direction `json:"direction,omitempty"`
	// The unique identifier for the conversation this call leg is part of.
	ConversationUuid string `json:"conversation_uuid,omitempty"`
}
