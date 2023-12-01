/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.91.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the VirtualAccountParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VirtualAccountParameters{}

// VirtualAccountParameters struct for VirtualAccountParameters
type VirtualAccountParameters struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	MinAmount NullableFloat64 `json:"min_amount,omitempty"`
	MaxAmount NullableFloat64 `json:"max_amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ChannelCode VirtualAccountChannelCode `json:"channel_code"`
	ChannelProperties VirtualAccountChannelProperties `json:"channel_properties"`
	// For payments in Vietnam only, alternative display requested for the virtual account
	AlternativeDisplayTypes []string `json:"alternative_display_types,omitempty"`
}

// NewVirtualAccountParameters instantiates a new VirtualAccountParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAccountParameters(channelCode VirtualAccountChannelCode, channelProperties VirtualAccountChannelProperties) *VirtualAccountParameters {
	this := VirtualAccountParameters{}
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	return &this
}

// NewVirtualAccountParametersWithDefaults instantiates a new VirtualAccountParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAccountParametersWithDefaults() *VirtualAccountParameters {
	this := VirtualAccountParameters{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountParameters) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *VirtualAccountParameters) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *VirtualAccountParameters) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *VirtualAccountParameters) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *VirtualAccountParameters) UnsetAmount() {
	o.Amount.Unset()
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountParameters) GetMinAmount() float64 {
	if o == nil || utils.IsNil(o.MinAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MinAmount.Get()
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountParameters) GetMinAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinAmount.Get(), o.MinAmount.IsSet()
}

// HasMinAmount returns a boolean if a field has been set.
func (o *VirtualAccountParameters) HasMinAmount() bool {
	if o != nil && o.MinAmount.IsSet() {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given NullableFloat64 and assigns it to the MinAmount field.
func (o *VirtualAccountParameters) SetMinAmount(v float64) {
	o.MinAmount.Set(&v)
}
// SetMinAmountNil sets the value for MinAmount to be an explicit nil
func (o *VirtualAccountParameters) SetMinAmountNil() {
	o.MinAmount.Set(nil)
}

// UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil
func (o *VirtualAccountParameters) UnsetMinAmount() {
	o.MinAmount.Unset()
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAccountParameters) GetMaxAmount() float64 {
	if o == nil || utils.IsNil(o.MaxAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.MaxAmount.Get()
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAccountParameters) GetMaxAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAmount.Get(), o.MaxAmount.IsSet()
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *VirtualAccountParameters) HasMaxAmount() bool {
	if o != nil && o.MaxAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given NullableFloat64 and assigns it to the MaxAmount field.
func (o *VirtualAccountParameters) SetMaxAmount(v float64) {
	o.MaxAmount.Set(&v)
}
// SetMaxAmountNil sets the value for MaxAmount to be an explicit nil
func (o *VirtualAccountParameters) SetMaxAmountNil() {
	o.MaxAmount.Set(nil)
}

// UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
func (o *VirtualAccountParameters) UnsetMaxAmount() {
	o.MaxAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *VirtualAccountParameters) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountParameters) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *VirtualAccountParameters) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *VirtualAccountParameters) SetCurrency(v string) {
	o.Currency = &v
}

// GetChannelCode returns the ChannelCode field value
func (o *VirtualAccountParameters) GetChannelCode() VirtualAccountChannelCode {
	if o == nil {
		var ret VirtualAccountChannelCode
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *VirtualAccountParameters) GetChannelCodeOk() (*VirtualAccountChannelCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *VirtualAccountParameters) SetChannelCode(v VirtualAccountChannelCode) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *VirtualAccountParameters) GetChannelProperties() VirtualAccountChannelProperties {
	if o == nil {
		var ret VirtualAccountChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *VirtualAccountParameters) GetChannelPropertiesOk() (*VirtualAccountChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *VirtualAccountParameters) SetChannelProperties(v VirtualAccountChannelProperties) {
	o.ChannelProperties = v
}

// GetAlternativeDisplayTypes returns the AlternativeDisplayTypes field value if set, zero value otherwise.
func (o *VirtualAccountParameters) GetAlternativeDisplayTypes() []string {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		var ret []string
		return ret
	}
	return o.AlternativeDisplayTypes
}

// GetAlternativeDisplayTypesOk returns a tuple with the AlternativeDisplayTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualAccountParameters) GetAlternativeDisplayTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeDisplayTypes) {
		return nil, false
	}
	return o.AlternativeDisplayTypes, true
}

// HasAlternativeDisplayTypes returns a boolean if a field has been set.
func (o *VirtualAccountParameters) HasAlternativeDisplayTypes() bool {
	if o != nil && !utils.IsNil(o.AlternativeDisplayTypes) {
		return true
	}

	return false
}

// SetAlternativeDisplayTypes gets a reference to the given []string and assigns it to the AlternativeDisplayTypes field.
func (o *VirtualAccountParameters) SetAlternativeDisplayTypes(v []string) {
	o.AlternativeDisplayTypes = v
}

func (o VirtualAccountParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualAccountParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
    }
	if o.MinAmount.IsSet() {
		toSerialize["min_amount"] = o.MinAmount.Get()
    }
	if o.MaxAmount.IsSet() {
		toSerialize["max_amount"] = o.MaxAmount.Get()
    }
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties
	if !utils.IsNil(o.AlternativeDisplayTypes) {
		toSerialize["alternative_display_types"] = o.AlternativeDisplayTypes
	}
	return toSerialize, nil
}

type NullableVirtualAccountParameters struct {
	value *VirtualAccountParameters
	isSet bool
}

func (v NullableVirtualAccountParameters) Get() *VirtualAccountParameters {
	return v.value
}

func (v *NullableVirtualAccountParameters) Set(val *VirtualAccountParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAccountParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAccountParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAccountParameters(val *VirtualAccountParameters) *NullableVirtualAccountParameters {
	return &NullableVirtualAccountParameters{value: val, isSet: true}
}

func (v NullableVirtualAccountParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAccountParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


