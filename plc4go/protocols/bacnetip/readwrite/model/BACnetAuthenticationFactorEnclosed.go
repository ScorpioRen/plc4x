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

// BACnetAuthenticationFactorEnclosed is the corresponding interface of BACnetAuthenticationFactorEnclosed
type BACnetAuthenticationFactorEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactor
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetAuthenticationFactorEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthenticationFactorEnclosed()
	// CreateBuilder creates a BACnetAuthenticationFactorEnclosedBuilder
	CreateBACnetAuthenticationFactorEnclosedBuilder() BACnetAuthenticationFactorEnclosedBuilder
}

// _BACnetAuthenticationFactorEnclosed is the data-structure of this message
type _BACnetAuthenticationFactorEnclosed struct {
	OpeningTag           BACnetOpeningTag
	AuthenticationFactor BACnetAuthenticationFactor
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetAuthenticationFactorEnclosed = (*_BACnetAuthenticationFactorEnclosed)(nil)

// NewBACnetAuthenticationFactorEnclosed factory function for _BACnetAuthenticationFactorEnclosed
func NewBACnetAuthenticationFactorEnclosed(openingTag BACnetOpeningTag, authenticationFactor BACnetAuthenticationFactor, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAuthenticationFactorEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetAuthenticationFactorEnclosed must not be nil")
	}
	if authenticationFactor == nil {
		panic("authenticationFactor of type BACnetAuthenticationFactor for BACnetAuthenticationFactorEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetAuthenticationFactorEnclosed must not be nil")
	}
	return &_BACnetAuthenticationFactorEnclosed{OpeningTag: openingTag, AuthenticationFactor: authenticationFactor, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAuthenticationFactorEnclosedBuilder is a builder for BACnetAuthenticationFactorEnclosed
type BACnetAuthenticationFactorEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, authenticationFactor BACnetAuthenticationFactor, closingTag BACnetClosingTag) BACnetAuthenticationFactorEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetAuthenticationFactorEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAuthenticationFactorEnclosedBuilder
	// WithAuthenticationFactor adds AuthenticationFactor (property field)
	WithAuthenticationFactor(BACnetAuthenticationFactor) BACnetAuthenticationFactorEnclosedBuilder
	// WithAuthenticationFactorBuilder adds AuthenticationFactor (property field) which is build by the builder
	WithAuthenticationFactorBuilder(func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetAuthenticationFactorEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAuthenticationFactorEnclosedBuilder
	// Build builds the BACnetAuthenticationFactorEnclosed or returns an error if something is wrong
	Build() (BACnetAuthenticationFactorEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAuthenticationFactorEnclosed
}

// NewBACnetAuthenticationFactorEnclosedBuilder() creates a BACnetAuthenticationFactorEnclosedBuilder
func NewBACnetAuthenticationFactorEnclosedBuilder() BACnetAuthenticationFactorEnclosedBuilder {
	return &_BACnetAuthenticationFactorEnclosedBuilder{_BACnetAuthenticationFactorEnclosed: new(_BACnetAuthenticationFactorEnclosed)}
}

type _BACnetAuthenticationFactorEnclosedBuilder struct {
	*_BACnetAuthenticationFactorEnclosed

	err *utils.MultiError
}

var _ (BACnetAuthenticationFactorEnclosedBuilder) = (*_BACnetAuthenticationFactorEnclosedBuilder)(nil)

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, authenticationFactor BACnetAuthenticationFactor, closingTag BACnetClosingTag) BACnetAuthenticationFactorEnclosedBuilder {
	return m.WithOpeningTag(openingTag).WithAuthenticationFactor(authenticationFactor).WithClosingTag(closingTag)
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetAuthenticationFactorEnclosedBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAuthenticationFactorEnclosedBuilder {
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

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithAuthenticationFactor(authenticationFactor BACnetAuthenticationFactor) BACnetAuthenticationFactorEnclosedBuilder {
	m.AuthenticationFactor = authenticationFactor
	return m
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithAuthenticationFactorBuilder(builderSupplier func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorEnclosedBuilder {
	builder := builderSupplier(m.AuthenticationFactor.CreateBACnetAuthenticationFactorBuilder())
	var err error
	m.AuthenticationFactor, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetAuthenticationFactorBuilder failed"))
	}
	return m
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetAuthenticationFactorEnclosedBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAuthenticationFactorEnclosedBuilder {
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

func (m *_BACnetAuthenticationFactorEnclosedBuilder) Build() (BACnetAuthenticationFactorEnclosed, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.AuthenticationFactor == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'authenticationFactor' not set"))
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
	return m._BACnetAuthenticationFactorEnclosed.deepCopy(), nil
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) MustBuild() BACnetAuthenticationFactorEnclosed {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetAuthenticationFactorEnclosedBuilder) DeepCopy() any {
	return m.CreateBACnetAuthenticationFactorEnclosedBuilder()
}

// CreateBACnetAuthenticationFactorEnclosedBuilder creates a BACnetAuthenticationFactorEnclosedBuilder
func (m *_BACnetAuthenticationFactorEnclosed) CreateBACnetAuthenticationFactorEnclosedBuilder() BACnetAuthenticationFactorEnclosedBuilder {
	if m == nil {
		return NewBACnetAuthenticationFactorEnclosedBuilder()
	}
	return &_BACnetAuthenticationFactorEnclosedBuilder{_BACnetAuthenticationFactorEnclosed: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAuthenticationFactorEnclosed) GetAuthenticationFactor() BACnetAuthenticationFactor {
	return m.AuthenticationFactor
}

func (m *_BACnetAuthenticationFactorEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorEnclosed(structType any) BACnetAuthenticationFactorEnclosed {
	if casted, ok := structType.(BACnetAuthenticationFactorEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorEnclosed) GetTypeName() string {
	return "BACnetAuthenticationFactorEnclosed"
}

func (m *_BACnetAuthenticationFactorEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationFactorEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAuthenticationFactorEnclosed, error) {
	return BACnetAuthenticationFactorEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAuthenticationFactorEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorEnclosed, error) {
		return BACnetAuthenticationFactorEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAuthenticationFactorEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAuthenticationFactorEnclosed, error) {
	v, err := (&_BACnetAuthenticationFactorEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthenticationFactorEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetAuthenticationFactorEnclosed BACnetAuthenticationFactorEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	authenticationFactor, err := ReadSimpleField[BACnetAuthenticationFactor](ctx, "authenticationFactor", ReadComplex[BACnetAuthenticationFactor](BACnetAuthenticationFactorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationFactor' field"))
	}
	m.AuthenticationFactor = authenticationFactor

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorEnclosed")
	}

	return m, nil
}

func (m *_BACnetAuthenticationFactorEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetAuthenticationFactor](ctx, "authenticationFactor", m.GetAuthenticationFactor(), WriteComplex[BACnetAuthenticationFactor](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'authenticationFactor' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationFactorEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAuthenticationFactorEnclosed) IsBACnetAuthenticationFactorEnclosed() {}

func (m *_BACnetAuthenticationFactorEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthenticationFactorEnclosed) deepCopy() *_BACnetAuthenticationFactorEnclosed {
	if m == nil {
		return nil
	}
	_BACnetAuthenticationFactorEnclosedCopy := &_BACnetAuthenticationFactorEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.AuthenticationFactor.DeepCopy().(BACnetAuthenticationFactor),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetAuthenticationFactorEnclosedCopy
}

func (m *_BACnetAuthenticationFactorEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
