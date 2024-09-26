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

// BACnetProcessIdSelection is the corresponding interface of BACnetProcessIdSelection
type BACnetProcessIdSelection interface {
	BACnetProcessIdSelectionContract
	BACnetProcessIdSelectionRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetProcessIdSelection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProcessIdSelection()
	// CreateBuilder creates a BACnetProcessIdSelectionBuilder
	CreateBACnetProcessIdSelectionBuilder() BACnetProcessIdSelectionBuilder
}

// BACnetProcessIdSelectionContract provides a set of functions which can be overwritten by a sub struct
type BACnetProcessIdSelectionContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetProcessIdSelection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProcessIdSelection()
	// CreateBuilder creates a BACnetProcessIdSelectionBuilder
	CreateBACnetProcessIdSelectionBuilder() BACnetProcessIdSelectionBuilder
}

// BACnetProcessIdSelectionRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetProcessIdSelectionRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetProcessIdSelection is the data-structure of this message
type _BACnetProcessIdSelection struct {
	_SubType        BACnetProcessIdSelection
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetProcessIdSelectionContract = (*_BACnetProcessIdSelection)(nil)

// NewBACnetProcessIdSelection factory function for _BACnetProcessIdSelection
func NewBACnetProcessIdSelection(peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelection {
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetProcessIdSelection must not be nil")
	}
	return &_BACnetProcessIdSelection{PeekedTagHeader: peekedTagHeader}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProcessIdSelectionBuilder is a builder for BACnetProcessIdSelection
type BACnetProcessIdSelectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetProcessIdSelectionBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetProcessIdSelectionBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProcessIdSelectionBuilder
	// Build builds the BACnetProcessIdSelection or returns an error if something is wrong
	Build() (BACnetProcessIdSelectionContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProcessIdSelectionContract
}

// NewBACnetProcessIdSelectionBuilder() creates a BACnetProcessIdSelectionBuilder
func NewBACnetProcessIdSelectionBuilder() BACnetProcessIdSelectionBuilder {
	return &_BACnetProcessIdSelectionBuilder{_BACnetProcessIdSelection: new(_BACnetProcessIdSelection)}
}

type _BACnetProcessIdSelectionBuilder struct {
	*_BACnetProcessIdSelection

	err *utils.MultiError
}

var _ (BACnetProcessIdSelectionBuilder) = (*_BACnetProcessIdSelectionBuilder)(nil)

func (m *_BACnetProcessIdSelectionBuilder) WithMandatoryFields(peekedTagHeader BACnetTagHeader) BACnetProcessIdSelectionBuilder {
	return m.WithPeekedTagHeader(peekedTagHeader)
}

func (m *_BACnetProcessIdSelectionBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetProcessIdSelectionBuilder {
	m.PeekedTagHeader = peekedTagHeader
	return m
}

func (m *_BACnetProcessIdSelectionBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProcessIdSelectionBuilder {
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

func (m *_BACnetProcessIdSelectionBuilder) Build() (BACnetProcessIdSelectionContract, error) {
	if m.PeekedTagHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetProcessIdSelection.deepCopy(), nil
}

func (m *_BACnetProcessIdSelectionBuilder) MustBuild() BACnetProcessIdSelectionContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetProcessIdSelectionBuilder) DeepCopy() any {
	return m.CreateBACnetProcessIdSelectionBuilder()
}

// CreateBACnetProcessIdSelectionBuilder creates a BACnetProcessIdSelectionBuilder
func (m *_BACnetProcessIdSelection) CreateBACnetProcessIdSelectionBuilder() BACnetProcessIdSelectionBuilder {
	if m == nil {
		return NewBACnetProcessIdSelectionBuilder()
	}
	return &_BACnetProcessIdSelectionBuilder{_BACnetProcessIdSelection: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelection) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetProcessIdSelection) GetPeekedTagNumber() uint8 {
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
func CastBACnetProcessIdSelection(structType any) BACnetProcessIdSelection {
	if casted, ok := structType.(BACnetProcessIdSelection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelection) GetTypeName() string {
	return "BACnetProcessIdSelection"
}

func (m *_BACnetProcessIdSelection) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetProcessIdSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetProcessIdSelectionParse[T BACnetProcessIdSelection](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetProcessIdSelectionParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetProcessIdSelectionParseWithBufferProducer[T BACnetProcessIdSelection]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetProcessIdSelectionParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetProcessIdSelectionParseWithBuffer[T BACnetProcessIdSelection](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetProcessIdSelection{}).parse(ctx, readBuffer)
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

func (m *_BACnetProcessIdSelection) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetProcessIdSelection BACnetProcessIdSelection, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelection")
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
	var _child BACnetProcessIdSelection
	switch {
	case peekedTagNumber == uint8(0): // BACnetProcessIdSelectionNull
		if _child, err = new(_BACnetProcessIdSelectionNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetProcessIdSelectionNull for type-switch of BACnetProcessIdSelection")
		}
	case 0 == 0: // BACnetProcessIdSelectionValue
		if _child, err = new(_BACnetProcessIdSelectionValue).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetProcessIdSelectionValue for type-switch of BACnetProcessIdSelection")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelection")
	}

	return _child, nil
}

func (pm *_BACnetProcessIdSelection) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetProcessIdSelection, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProcessIdSelection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelection")
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

	if popErr := writeBuffer.PopContext("BACnetProcessIdSelection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelection")
	}
	return nil
}

func (m *_BACnetProcessIdSelection) IsBACnetProcessIdSelection() {}

func (m *_BACnetProcessIdSelection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProcessIdSelection) deepCopy() *_BACnetProcessIdSelection {
	if m == nil {
		return nil
	}
	_BACnetProcessIdSelectionCopy := &_BACnetProcessIdSelection{
		nil, // will be set by child
		m.PeekedTagHeader.DeepCopy().(BACnetTagHeader),
	}
	return _BACnetProcessIdSelectionCopy
}
