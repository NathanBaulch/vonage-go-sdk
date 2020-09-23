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
// EndpointWebsocket Connect to a Websocket
type EndpointWebsocket struct {
	// The type of connection. Must be `websocket`
	Type string `json:"type"`
	Uri string `json:"uri,omitempty"`
	ContentType string `json:"content-type"`
	Headers EndpointWebsocketHeaders `json:"headers,omitempty"`
}
