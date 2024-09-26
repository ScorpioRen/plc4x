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

// BACnetPrescale is the corresponding interface of BACnetPrescale
type BACnetPrescale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetMultiplier returns Multiplier (property field)
	GetMultiplier() BACnetContextTagUnsignedInteger
	// GetModuloDivide returns ModuloDivide (property field)
	GetModuloDivide() BACnetContextTagUnsignedInteger
	// IsBACnetPrescale is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPrescale()
	// CreateBuilder creates a BACnetPrescaleBuilder
	CreateBACnetPrescaleBuilder() BACnetPrescaleBuilder
}

// _BACnetPrescale is the data-structure of this message
type _BACnetPrescale struct {
	Multiplier   BACnetContextTagUnsignedInteger
	ModuloDivide BACnetContextTagUnsignedInteger
}

var _ BACnetPrescale = (*_BACnetPrescale)(nil)

// NewBACnetPrescale factory function for _BACnetPrescale
func NewBACnetPrescale(multiplier BACnetContextTagUnsignedInteger, moduloDivide BACnetContextTagUnsignedInteger) *_BACnetPrescale {
	if multiplier == nil {
		panic("multiplier of type BACnetContextTagUnsignedInteger for BACnetPrescale must not be nil")
	}
	if moduloDivide == nil {
		panic("moduloDivide of type BACnetContextTagUnsignedInteger for BACnetPrescale must not be nil")
	}
	return &_BACnetPrescale{Multiplier: multiplier, ModuloDivide: moduloDivide}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPrescaleBuilder is a builder for BACnetPrescale
type BACnetPrescaleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(multiplier BACnetContextTagUnsignedInteger, moduloDivide BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder
	// WithMultiplier adds Multiplier (property field)
	WithMultiplier(BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder
	// WithMultiplierBuilder adds Multiplier (property field) which is build by the builder
	WithMultiplierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPrescaleBuilder
	// WithModuloDivide adds ModuloDivide (property field)
	WithModuloDivide(BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder
	// WithModuloDivideBuilder adds ModuloDivide (property field) which is build by the builder
	WithModuloDivideBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPrescaleBuilder
	// Build builds the BACnetPrescale or returns an error if something is wrong
	Build() (BACnetPrescale, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPrescale
}

// NewBACnetPrescaleBuilder() creates a BACnetPrescaleBuilder
func NewBACnetPrescaleBuilder() BACnetPrescaleBuilder {
	return &_BACnetPrescaleBuilder{_BACnetPrescale: new(_BACnetPrescale)}
}

type _BACnetPrescaleBuilder struct {
	*_BACnetPrescale

	err *utils.MultiError
}

var _ (BACnetPrescaleBuilder) = (*_BACnetPrescaleBuilder)(nil)

func (m *_BACnetPrescaleBuilder) WithMandatoryFields(multiplier BACnetContextTagUnsignedInteger, moduloDivide BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder {
	return m.WithMultiplier(multiplier).WithModuloDivide(moduloDivide)
}

func (m *_BACnetPrescaleBuilder) WithMultiplier(multiplier BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder {
	m.Multiplier = multiplier
	return m
}

func (m *_BACnetPrescaleBuilder) WithMultiplierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPrescaleBuilder {
	builder := builderSupplier(m.Multiplier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.Multiplier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetPrescaleBuilder) WithModuloDivide(moduloDivide BACnetContextTagUnsignedInteger) BACnetPrescaleBuilder {
	m.ModuloDivide = moduloDivide
	return m
}

func (m *_BACnetPrescaleBuilder) WithModuloDivideBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPrescaleBuilder {
	builder := builderSupplier(m.ModuloDivide.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ModuloDivide, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetPrescaleBuilder) Build() (BACnetPrescale, error) {
	if m.Multiplier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'multiplier' not set"))
	}
	if m.ModuloDivide == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'moduloDivide' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPrescale.deepCopy(), nil
}

func (m *_BACnetPrescaleBuilder) MustBuild() BACnetPrescale {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPrescaleBuilder) DeepCopy() any {
	return m.CreateBACnetPrescaleBuilder()
}

// CreateBACnetPrescaleBuilder creates a BACnetPrescaleBuilder
func (m *_BACnetPrescale) CreateBACnetPrescaleBuilder() BACnetPrescaleBuilder {
	if m == nil {
		return NewBACnetPrescaleBuilder()
	}
	return &_BACnetPrescaleBuilder{_BACnetPrescale: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPrescale) GetMultiplier() BACnetContextTagUnsignedInteger {
	return m.Multiplier
}

func (m *_BACnetPrescale) GetModuloDivide() BACnetContextTagUnsignedInteger {
	return m.ModuloDivide
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPrescale(structType any) BACnetPrescale {
	if casted, ok := structType.(BACnetPrescale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPrescale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPrescale) GetTypeName() string {
	return "BACnetPrescale"
}

func (m *_BACnetPrescale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (multiplier)
	lengthInBits += m.Multiplier.GetLengthInBits(ctx)

	// Simple field (moduloDivide)
	lengthInBits += m.ModuloDivide.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPrescale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPrescaleParse(ctx context.Context, theBytes []byte) (BACnetPrescale, error) {
	return BACnetPrescaleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPrescaleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
		return BACnetPrescaleParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetPrescaleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPrescale, error) {
	v, err := (&_BACnetPrescale{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPrescale) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetPrescale BACnetPrescale, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPrescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPrescale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	multiplier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "multiplier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'multiplier' field"))
	}
	m.Multiplier = multiplier

	moduloDivide, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "moduloDivide", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moduloDivide' field"))
	}
	m.ModuloDivide = moduloDivide

	if closeErr := readBuffer.CloseContext("BACnetPrescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPrescale")
	}

	return m, nil
}

func (m *_BACnetPrescale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPrescale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPrescale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPrescale")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "multiplier", m.GetMultiplier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'multiplier' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "moduloDivide", m.GetModuloDivide(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'moduloDivide' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPrescale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPrescale")
	}
	return nil
}

func (m *_BACnetPrescale) IsBACnetPrescale() {}

func (m *_BACnetPrescale) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPrescale) deepCopy() *_BACnetPrescale {
	if m == nil {
		return nil
	}
	_BACnetPrescaleCopy := &_BACnetPrescale{
		m.Multiplier.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ModuloDivide.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	return _BACnetPrescaleCopy
}

func (m *_BACnetPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
