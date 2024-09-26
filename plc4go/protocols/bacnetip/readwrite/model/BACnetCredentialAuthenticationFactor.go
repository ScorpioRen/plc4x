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

// BACnetCredentialAuthenticationFactor is the corresponding interface of BACnetCredentialAuthenticationFactor
type BACnetCredentialAuthenticationFactor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDisable returns Disable (property field)
	GetDisable() BACnetAccessAuthenticationFactorDisableTagged
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed
	// IsBACnetCredentialAuthenticationFactor is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCredentialAuthenticationFactor()
	// CreateBuilder creates a BACnetCredentialAuthenticationFactorBuilder
	CreateBACnetCredentialAuthenticationFactorBuilder() BACnetCredentialAuthenticationFactorBuilder
}

// _BACnetCredentialAuthenticationFactor is the data-structure of this message
type _BACnetCredentialAuthenticationFactor struct {
	Disable              BACnetAccessAuthenticationFactorDisableTagged
	AuthenticationFactor BACnetAuthenticationFactorEnclosed
}

var _ BACnetCredentialAuthenticationFactor = (*_BACnetCredentialAuthenticationFactor)(nil)

// NewBACnetCredentialAuthenticationFactor factory function for _BACnetCredentialAuthenticationFactor
func NewBACnetCredentialAuthenticationFactor(disable BACnetAccessAuthenticationFactorDisableTagged, authenticationFactor BACnetAuthenticationFactorEnclosed) *_BACnetCredentialAuthenticationFactor {
	if disable == nil {
		panic("disable of type BACnetAccessAuthenticationFactorDisableTagged for BACnetCredentialAuthenticationFactor must not be nil")
	}
	if authenticationFactor == nil {
		panic("authenticationFactor of type BACnetAuthenticationFactorEnclosed for BACnetCredentialAuthenticationFactor must not be nil")
	}
	return &_BACnetCredentialAuthenticationFactor{Disable: disable, AuthenticationFactor: authenticationFactor}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCredentialAuthenticationFactorBuilder is a builder for BACnetCredentialAuthenticationFactor
type BACnetCredentialAuthenticationFactorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(disable BACnetAccessAuthenticationFactorDisableTagged, authenticationFactor BACnetAuthenticationFactorEnclosed) BACnetCredentialAuthenticationFactorBuilder
	// WithDisable adds Disable (property field)
	WithDisable(BACnetAccessAuthenticationFactorDisableTagged) BACnetCredentialAuthenticationFactorBuilder
	// WithDisableBuilder adds Disable (property field) which is build by the builder
	WithDisableBuilder(func(BACnetAccessAuthenticationFactorDisableTaggedBuilder) BACnetAccessAuthenticationFactorDisableTaggedBuilder) BACnetCredentialAuthenticationFactorBuilder
	// WithAuthenticationFactor adds AuthenticationFactor (property field)
	WithAuthenticationFactor(BACnetAuthenticationFactorEnclosed) BACnetCredentialAuthenticationFactorBuilder
	// WithAuthenticationFactorBuilder adds AuthenticationFactor (property field) which is build by the builder
	WithAuthenticationFactorBuilder(func(BACnetAuthenticationFactorEnclosedBuilder) BACnetAuthenticationFactorEnclosedBuilder) BACnetCredentialAuthenticationFactorBuilder
	// Build builds the BACnetCredentialAuthenticationFactor or returns an error if something is wrong
	Build() (BACnetCredentialAuthenticationFactor, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCredentialAuthenticationFactor
}

// NewBACnetCredentialAuthenticationFactorBuilder() creates a BACnetCredentialAuthenticationFactorBuilder
func NewBACnetCredentialAuthenticationFactorBuilder() BACnetCredentialAuthenticationFactorBuilder {
	return &_BACnetCredentialAuthenticationFactorBuilder{_BACnetCredentialAuthenticationFactor: new(_BACnetCredentialAuthenticationFactor)}
}

type _BACnetCredentialAuthenticationFactorBuilder struct {
	*_BACnetCredentialAuthenticationFactor

	err *utils.MultiError
}

var _ (BACnetCredentialAuthenticationFactorBuilder) = (*_BACnetCredentialAuthenticationFactorBuilder)(nil)

func (m *_BACnetCredentialAuthenticationFactorBuilder) WithMandatoryFields(disable BACnetAccessAuthenticationFactorDisableTagged, authenticationFactor BACnetAuthenticationFactorEnclosed) BACnetCredentialAuthenticationFactorBuilder {
	return m.WithDisable(disable).WithAuthenticationFactor(authenticationFactor)
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) WithDisable(disable BACnetAccessAuthenticationFactorDisableTagged) BACnetCredentialAuthenticationFactorBuilder {
	m.Disable = disable
	return m
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) WithDisableBuilder(builderSupplier func(BACnetAccessAuthenticationFactorDisableTaggedBuilder) BACnetAccessAuthenticationFactorDisableTaggedBuilder) BACnetCredentialAuthenticationFactorBuilder {
	builder := builderSupplier(m.Disable.CreateBACnetAccessAuthenticationFactorDisableTaggedBuilder())
	var err error
	m.Disable, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetAccessAuthenticationFactorDisableTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) WithAuthenticationFactor(authenticationFactor BACnetAuthenticationFactorEnclosed) BACnetCredentialAuthenticationFactorBuilder {
	m.AuthenticationFactor = authenticationFactor
	return m
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) WithAuthenticationFactorBuilder(builderSupplier func(BACnetAuthenticationFactorEnclosedBuilder) BACnetAuthenticationFactorEnclosedBuilder) BACnetCredentialAuthenticationFactorBuilder {
	builder := builderSupplier(m.AuthenticationFactor.CreateBACnetAuthenticationFactorEnclosedBuilder())
	var err error
	m.AuthenticationFactor, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetAuthenticationFactorEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) Build() (BACnetCredentialAuthenticationFactor, error) {
	if m.Disable == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'disable' not set"))
	}
	if m.AuthenticationFactor == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'authenticationFactor' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetCredentialAuthenticationFactor.deepCopy(), nil
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) MustBuild() BACnetCredentialAuthenticationFactor {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetCredentialAuthenticationFactorBuilder) DeepCopy() any {
	return m.CreateBACnetCredentialAuthenticationFactorBuilder()
}

// CreateBACnetCredentialAuthenticationFactorBuilder creates a BACnetCredentialAuthenticationFactorBuilder
func (m *_BACnetCredentialAuthenticationFactor) CreateBACnetCredentialAuthenticationFactorBuilder() BACnetCredentialAuthenticationFactorBuilder {
	if m == nil {
		return NewBACnetCredentialAuthenticationFactorBuilder()
	}
	return &_BACnetCredentialAuthenticationFactorBuilder{_BACnetCredentialAuthenticationFactor: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCredentialAuthenticationFactor) GetDisable() BACnetAccessAuthenticationFactorDisableTagged {
	return m.Disable
}

func (m *_BACnetCredentialAuthenticationFactor) GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed {
	return m.AuthenticationFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCredentialAuthenticationFactor(structType any) BACnetCredentialAuthenticationFactor {
	if casted, ok := structType.(BACnetCredentialAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCredentialAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) GetTypeName() string {
	return "BACnetCredentialAuthenticationFactor"
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (disable)
	lengthInBits += m.Disable.GetLengthInBits(ctx)

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCredentialAuthenticationFactorParse(ctx context.Context, theBytes []byte) (BACnetCredentialAuthenticationFactor, error) {
	return BACnetCredentialAuthenticationFactorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCredentialAuthenticationFactorParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCredentialAuthenticationFactor, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCredentialAuthenticationFactor, error) {
		return BACnetCredentialAuthenticationFactorParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCredentialAuthenticationFactorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCredentialAuthenticationFactor, error) {
	v, err := (&_BACnetCredentialAuthenticationFactor{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetCredentialAuthenticationFactor) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetCredentialAuthenticationFactor BACnetCredentialAuthenticationFactor, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCredentialAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCredentialAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	disable, err := ReadSimpleField[BACnetAccessAuthenticationFactorDisableTagged](ctx, "disable", ReadComplex[BACnetAccessAuthenticationFactorDisableTagged](BACnetAccessAuthenticationFactorDisableTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'disable' field"))
	}
	m.Disable = disable

	authenticationFactor, err := ReadSimpleField[BACnetAuthenticationFactorEnclosed](ctx, "authenticationFactor", ReadComplex[BACnetAuthenticationFactorEnclosed](BACnetAuthenticationFactorEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationFactor' field"))
	}
	m.AuthenticationFactor = authenticationFactor

	if closeErr := readBuffer.CloseContext("BACnetCredentialAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCredentialAuthenticationFactor")
	}

	return m, nil
}

func (m *_BACnetCredentialAuthenticationFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCredentialAuthenticationFactor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCredentialAuthenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCredentialAuthenticationFactor")
	}

	if err := WriteSimpleField[BACnetAccessAuthenticationFactorDisableTagged](ctx, "disable", m.GetDisable(), WriteComplex[BACnetAccessAuthenticationFactorDisableTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'disable' field")
	}

	if err := WriteSimpleField[BACnetAuthenticationFactorEnclosed](ctx, "authenticationFactor", m.GetAuthenticationFactor(), WriteComplex[BACnetAuthenticationFactorEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'authenticationFactor' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCredentialAuthenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCredentialAuthenticationFactor")
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) IsBACnetCredentialAuthenticationFactor() {}

func (m *_BACnetCredentialAuthenticationFactor) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCredentialAuthenticationFactor) deepCopy() *_BACnetCredentialAuthenticationFactor {
	if m == nil {
		return nil
	}
	_BACnetCredentialAuthenticationFactorCopy := &_BACnetCredentialAuthenticationFactor{
		m.Disable.DeepCopy().(BACnetAccessAuthenticationFactorDisableTagged),
		m.AuthenticationFactor.DeepCopy().(BACnetAuthenticationFactorEnclosed),
	}
	return _BACnetCredentialAuthenticationFactorCopy
}

func (m *_BACnetCredentialAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
