// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TagFilter - A tag filter defined by the tags key and value.
type TagFilter struct {
	// The key of the tag.
	Key string `json:"key"`
	// The value of the tag.
	Value *string `json:"value,omitempty"`
}

func (o *TagFilter) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *TagFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}