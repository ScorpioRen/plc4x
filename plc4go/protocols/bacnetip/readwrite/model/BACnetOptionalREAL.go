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

// BACnetOptionalREAL is the corresponding interface of BACnetOptionalREAL
type BACnetOptionalREAL interface {
	BACnetOptionalREALContract
	BACnetOptionalREALRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetOptionalREAL is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalREAL()
	// CreateBuilder creates a BACnetOptionalREALBuilder
	CreateBACnetOptionalREALBuilder() BACnetOptionalREALBuilder
}

// BACnetOptionalREALContract provides a set of functions which can be overwritten by a sub struct
type BACnetOptionalREALContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetOptionalREAL is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOptionalREAL()
	// CreateBuilder creates a BACnetOptionalREALBuilder
	CreateBACnetOptionalREALBuilder() BACnetOptionalREALBuilder
}

// BACnetOptionalREALRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetOptionalREALRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetOptionalREAL is the data-structure of this message
type _BACnetOptionalREAL struct {
	_SubType        BACnetOptionalREAL
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetOptionalREALContract = (*_BACnetOptionalREAL)(nil)

// NewBACnetOptionalREAL factory function for _BACnetOptionalREAL
func NewBACnetOptionalREAL(peekedTagHeader BACnetTagHeader) *_BACnetOptionalREAL {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetOptionalREAL must not be nil")
	}
	return &_BACnetOptionalREAL{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetOptionalREALBuilder is a builder for BACnetOptionalREAL
type BACnetOptionalREALBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetOptionalREALBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetOptionalREALBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetOptionalREALBuilder
	// Build builds the BACnetOptionalREAL or returns an error if something is wrong
	Build() (BACnetOptionalREALContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetOptionalREALContract
}

// NewBACnetOptionalREALBuilder() creates a BACnetOptionalREALBuilder
func NewBACnetOptionalREALBuilder() BACnetOptionalREALBuilder {
	return &_BACnetOptionalREALBuilder{_BACnetOptionalREAL: new(_BACnetOptionalREAL)}
}

type _BACnetOptionalREALBuilder struct {
	*_BACnetOptionalREAL

	err *utils.MultiError
}

var _ (BACnetOptionalREALBuilder) = (*_BACnetOptionalREALBuilder)(nil)

func (m *_BACnetOptionalREALBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetOptionalREALBuilder {
	return m.WithPeekedTagHeader(peekedTagHeader)
}

func (m *_BACnetOptionalREALBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetOptionalREALBuilder {
	m.PeekedTagHeader = peekedTagHeader
	return m
}

func (m *_BACnetOptionalREALBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetOptionalREALBuilder {
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

func (m *_BACnetOptionalREALBuilder) Build() (BACnetOptionalREALContract, error) {
	if m.PeekedTagHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetOptionalREAL.deepCopy(), nil
}

func (m *_BACnetOptionalREALBuilder) MustBuild() BACnetOptionalREALContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetOptionalREALBuilder) DeepCopy() any {
	return m.CreateBACnetOptionalREALBuilder()
}

// CreateBACnetOptionalREALBuilder creates a BACnetOptionalREALBuilder
func (m *_BACnetOptionalREAL) CreateBACnetOptionalREALBuilder() BACnetOptionalREALBuilder {
	if m == nil {
		return NewBACnetOptionalREALBuilder()
	}
	return &_BACnetOptionalREALBuilder{_BACnetOptionalREAL: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalREAL) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetOptionalREAL) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOptionalREAL(structType any) BACnetOptionalREAL {
	if casted, ok := structType.(BACnetOptionalREAL); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalREAL); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalREAL) GetTypeName() string {
	return "BACnetOptionalREAL"
}

func (m *_BACnetOptionalREAL) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalREAL) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetOptionalREALParse[T BACnetOptionalREAL](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetOptionalREALParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalREALParseWithBufferProducer[T BACnetOptionalREAL]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetOptionalREALParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetOptionalREALParseWithBuffer[T BACnetOptionalREAL](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetOptionalREAL{}).parse(ctx, readBuffer)
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

func (m *_BACnetOptionalREAL) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetOptionalREAL BACnetOptionalREAL, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalREAL"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalREAL")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetOptionalREAL
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalREALNull
		if _child, err = new(_BACnetOptionalREALNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalREALNull for type-switch of BACnetOptionalREAL")
		}
	case 0 == 0: // BACnetOptionalREALValue
		if _child, err = new(_BACnetOptionalREALValue).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetOptionalREALValue for type-switch of BACnetOptionalREAL")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalREAL"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalREAL")
	}

	return _child, nil
}

func (pm *_BACnetOptionalREAL) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalREAL, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOptionalREAL"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalREAL")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetOptionalREAL"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalREAL")
	}
	return nil
}

func (m *_BACnetOptionalREAL) IsBACnetOptionalREAL() {}

func (m *_BACnetOptionalREAL) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOptionalREAL) deepCopy() *_BACnetOptionalREAL {
	if m == nil {
		return nil
	}
	_BACnetOptionalREALCopy := &_BACnetOptionalREAL{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetOptionalREALCopy
}
