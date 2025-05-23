/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

import (
	"encoding/json"
)

// TextMessageV2
// TextMessageV2
// https://developers.line.biz/en/reference/messaging-api/#text-message-v2
type TextMessageV2 struct {
	Message

	/**
	 * Get QuickReply
	 */
	QuickReply *QuickReply `json:"quickReply,omitempty"`

	/**
	 * Get Sender
	 */
	Sender *Sender `json:"sender,omitempty"`

	/**
	 * Get Text
	 */
	Text string `json:"text"`

	/**
	 * A mapping that specifies substitutions for parts enclosed in {} within the `text` field.
	 */
	Substitution map[string]SubstitutionObjectInterface `json:"substitution,omitempty"`

	/**
	 * Quote token of the message you want to quote.
	 */
	QuoteToken string `json:"quoteToken,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the TextMessageV2 struct.
func (r *TextMessageV2) MarshalJSON() ([]byte, error) {

	type Alias TextMessageV2
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "textV2",
	})
}
