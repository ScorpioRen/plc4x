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

// BACnetRecipientEnclosed is the corresponding interface of BACnetRecipientEnclosed
type BACnetRecipientEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetRecipientEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientEnclosed()
	// CreateBuilder creates a BACnetRecipientEnclosedBuilder
	CreateBACnetRecipientEnclosedBuilder() BACnetRecipientEnclosedBuilder
}

// _BACnetRecipientEnclosed is the data-structure of this message
type _BACnetRecipientEnclosed struct {
	OpeningTag BACnetOpeningTag
	Recipient  BACnetRecipient
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetRecipientEnclosed = (*_BACnetRecipientEnclosed)(nil)

// NewBACnetRecipientEnclosed factory function for _BACnetRecipientEnclosed
func NewBACnetRecipientEnclosed(openingTag BACnetOpeningTag, recipient BACnetRecipient, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetRecipientEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetRecipientEnclosed must not be nil")
	}
	if recipient == nil {
		panic("recipient of type BACnetRecipient for BACnetRecipientEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetRecipientEnclosed must not be nil")
	}
	return &_BACnetRecipientEnclosed{OpeningTag: openingTag, Recipient: recipient, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetRecipientEnclosedBuilder is a builder for BACnetRecipientEnclosed
type BACnetRecipientEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, recipient BACnetRecipient, closingTag BACnetClosingTag) BACnetRecipientEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetRecipientEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetRecipientEnclosedBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipient) BACnetRecipientEnclosedBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientBuilder) BACnetRecipientBuilder) BACnetRecipientEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetRecipientEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetRecipientEnclosedBuilder
	// Build builds the BACnetRecipientEnclosed or returns an error if something is wrong
	Build() (BACnetRecipientEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetRecipientEnclosed
}

// NewBACnetRecipientEnclosedBuilder() creates a BACnetRecipientEnclosedBuilder
func NewBACnetRecipientEnclosedBuilder() BACnetRecipientEnclosedBuilder {
	return &_BACnetRecipientEnclosedBuilder{_BACnetRecipientEnclosed: new(_BACnetRecipientEnclosed)}
}

type _BACnetRecipientEnclosedBuilder struct {
	*_BACnetRecipientEnclosed

	err *utils.MultiError
}

var _ (BACnetRecipientEnclosedBuilder) = (*_BACnetRecipientEnclosedBuilder)(nil)

func (m *_BACnetRecipientEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, recipient BACnetRecipient, closingTag BACnetClosingTag) BACnetRecipientEnclosedBuilder {
	return m.WithOpeningTag(openingTag).WithRecipient(recipient).WithClosingTag(closingTag)
}

func (m *_BACnetRecipientEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetRecipientEnclosedBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetRecipientEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetRecipientEnclosedBuilder {
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

func (m *_BACnetRecipientEnclosedBuilder) WithRecipient(recipient BACnetRecipient) BACnetRecipientEnclosedBuilder {
	m.Recipient = recipient
	return m
}

func (m *_BACnetRecipientEnclosedBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientBuilder) BACnetRecipientBuilder) BACnetRecipientEnclosedBuilder {
	builder := builderSupplier(m.Recipient.CreateBACnetRecipientBuilder())
	var err error
	m.Recipient, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetRecipientBuilder failed"))
	}
	return m
}

func (m *_BACnetRecipientEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetRecipientEnclosedBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetRecipientEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetRecipientEnclosedBuilder {
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

func (m *_BACnetRecipientEnclosedBuilder) Build() (BACnetRecipientEnclosed, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.Recipient == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'recipient' not set"))
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
	return m._BACnetRecipientEnclosed.deepCopy(), nil
}

func (m *_BACnetRecipientEnclosedBuilder) MustBuild() BACnetRecipientEnclosed {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetRecipientEnclosedBuilder) DeepCopy() any {
	return m.CreateBACnetRecipientEnclosedBuilder()
}

// CreateBACnetRecipientEnclosedBuilder creates a BACnetRecipientEnclosedBuilder
func (m *_BACnetRecipientEnclosed) CreateBACnetRecipientEnclosedBuilder() BACnetRecipientEnclosedBuilder {
	if m == nil {
		return NewBACnetRecipientEnclosedBuilder()
	}
	return &_BACnetRecipientEnclosedBuilder{_BACnetRecipientEnclosed: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetRecipientEnclosed) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetRecipientEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetRecipientEnclosed(structType any) BACnetRecipientEnclosed {
	if casted, ok := structType.(BACnetRecipientEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientEnclosed) GetTypeName() string {
	return "BACnetRecipientEnclosed"
}

func (m *_BACnetRecipientEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	return BACnetRecipientEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetRecipientEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientEnclosed, error) {
		return BACnetRecipientEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetRecipientEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	v, err := (&_BACnetRecipientEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetRecipientEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetRecipientEnclosed BACnetRecipientEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	recipient, err := ReadSimpleField[BACnetRecipient](ctx, "recipient", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetRecipientEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientEnclosed")
	}

	return m, nil
}

func (m *_BACnetRecipientEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipientEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetRecipient](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipient](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRecipientEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetRecipientEnclosed) IsBACnetRecipientEnclosed() {}

func (m *_BACnetRecipientEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetRecipientEnclosed) deepCopy() *_BACnetRecipientEnclosed {
	if m == nil {
		return nil
	}
	_BACnetRecipientEnclosedCopy := &_BACnetRecipientEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.Recipient.DeepCopy().(BACnetRecipient),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetRecipientEnclosedCopy
}

func (m *_BACnetRecipientEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
