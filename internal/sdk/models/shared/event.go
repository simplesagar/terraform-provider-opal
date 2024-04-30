// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"time"
)

// # Event Object
// ### Description
// The `Event` object is used to represent an event.
//
// ### Usage Example
// Fetch from the `LIST Events` endpoint.
type Event struct {
	// The email of the actor user.
	ActorEmail *string `json:"actor_email,omitempty"`
	// The IP address of the event actor.
	ActorIPAddress *string `json:"actor_ip_address,omitempty"`
	// The name of the actor user.
	ActorName interface{} `json:"actor_name"`
	// The ID of the actor user.
	ActorUserID string `json:"actor_user_id"`
	// The name of the API token used to create the event.
	APITokenName *string `json:"api_token_name,omitempty"`
	// The preview of the API token used to create the event.
	APITokenPreview *string `json:"api_token_preview,omitempty"`
	// The day and time the event was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the event.
	EventID string `json:"event_id"`
	// The event type.
	EventType string     `json:"event_type"`
	SubEvents []SubEvent `json:"sub_events,omitempty"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Event) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Event) GetActorEmail() *string {
	if o == nil {
		return nil
	}
	return o.ActorEmail
}

func (o *Event) GetActorIPAddress() *string {
	if o == nil {
		return nil
	}
	return o.ActorIPAddress
}

func (o *Event) GetActorName() interface{} {
	if o == nil {
		return nil
	}
	return o.ActorName
}

func (o *Event) GetActorUserID() string {
	if o == nil {
		return ""
	}
	return o.ActorUserID
}

func (o *Event) GetAPITokenName() *string {
	if o == nil {
		return nil
	}
	return o.APITokenName
}

func (o *Event) GetAPITokenPreview() *string {
	if o == nil {
		return nil
	}
	return o.APITokenPreview
}

func (o *Event) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Event) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

func (o *Event) GetEventType() string {
	if o == nil {
		return ""
	}
	return o.EventType
}

func (o *Event) GetSubEvents() []SubEvent {
	if o == nil {
		return nil
	}
	return o.SubEvents
}