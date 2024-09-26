/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersChangeOfDiscreteValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValue
type BACnetNotificationParametersChangeOfDiscreteValueNewValue interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValue()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetNotificationParametersChangeOfDiscreteValueNewValueContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValue()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedIsContextTag returns PeekedIsContextTag (discriminator field)
	GetPeekedIsContextTag() bool
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValue struct {
	_SubType        BACnetNotificationParametersChangeOfDiscreteValueNewValue
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)(nil)

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValue factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValue
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfDiscreteValueNewValue must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetNotificationParametersChangeOfDiscreteValueNewValue must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfDiscreteValueNewValue must not be nil")
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder is a builder for BACnetNotificationParametersChangeOfDiscreteValueNewValue
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
	// Build builds the BACnetNotificationParametersChangeOfDiscreteValueNewValue or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder() creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValue: new(_BACnetNotificationParametersChangeOfDiscreteValueNewValue)}
}

type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder)(nil)

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	return m.WithOpeningTag(openingTag).WithPeekedTagHeader(peekedTagHeader).WithClosingTag(closingTag)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	m.PeekedTagHeader = peekedTagHeader
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	builder := builderSupplier(m.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	m.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) Build() (BACnetNotificationParametersChangeOfDiscreteValueNewValueContract, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.PeekedTagHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue.deepCopy(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) MustBuild() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder) DeepCopy() any {
	return m.CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder()
}

// CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder creates a BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder
func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) CreateBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder() BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder {
	if m == nil {
		return NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder()
	}
	return &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBuilder{_BACnetNotificationParametersChangeOfDiscreteValueNewValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (pm *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedIsContextTag() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValue(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValue"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParse[T BACnetNotificationParametersChangeOfDiscreteValueNewValue](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBufferProducer[T BACnetNotificationParametersChangeOfDiscreteValueNewValue](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParseWithBuffer[T BACnetNotificationParametersChangeOfDiscreteValueNewValue](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetNotificationParametersChangeOfDiscreteValueNewValue{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValue BACnetNotificationParametersChangeOfDiscreteValueNewValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	peekedIsContextTag, err := ReadVirtualField[bool](ctx, "peekedIsContextTag", (*bool)(nil), bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedIsContextTag' field"))
	}
	_ = peekedIsContextTag

	// Validation
	if !(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetNotificationParametersChangeOfDiscreteValueNewValue
	switch {
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
		if _child, err = new(_BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime for type-switch of BACnetNotificationParametersChangeOfDiscreteValueNewValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}

	return _child, nil
}

func (pm *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParametersChangeOfDiscreteValueNewValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ = peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) IsBACnetNotificationParametersChangeOfDiscreteValueNewValue() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValue) deepCopy() *_BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfDiscreteValueNewValueCopy := &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
		nil, // will be set by child
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetNotificationParametersChangeOfDiscreteValueNewValueCopy
}
