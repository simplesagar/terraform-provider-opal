// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestedItem - # Requested Item Object
// ### Description
// The `RequestedItem` object is used to represent a request target item.
//
// ### Usage Example
// Returned from the `GET Requests` endpoint.
type RequestedItem struct {
	// The name of the access level requested.
	AccessLevelName *string `json:"access_level_name,omitempty"`
	// The ID of the access level requested on the remote system.
	AccessLevelRemoteID *string `json:"access_level_remote_id,omitempty"`
	// The ID of the group requested.
	GroupID *string `json:"group_id,omitempty"`
	// The name of the target.
	Name *string `json:"name,omitempty"`
	// The ID of the resource requested.
	ResourceID *string `json:"resource_id,omitempty"`
}

func (o *RequestedItem) GetAccessLevelName() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelName
}

func (o *RequestedItem) GetAccessLevelRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.AccessLevelRemoteID
}

func (o *RequestedItem) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}

func (o *RequestedItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RequestedItem) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}