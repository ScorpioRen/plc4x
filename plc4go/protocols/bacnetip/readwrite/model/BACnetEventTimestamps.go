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

// BACnetEventTimestamps is the corresponding interface of BACnetEventTimestamps
type BACnetEventTimestamps interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() BACnetTimeStamp
	// GetToFault returns ToFault (property field)
	GetToFault() BACnetTimeStamp
	// GetToNormal returns ToNormal (property field)
	GetToNormal() BACnetTimeStamp
	// IsBACnetEventTimestamps is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventTimestamps()
	// CreateBuilder creates a BACnetEventTimestampsBuilder
	CreateBACnetEventTimestampsBuilder() BACnetEventTimestampsBuilder
}

// _BACnetEventTimestamps is the data-structure of this message
type _BACnetEventTimestamps struct {
	ToOffnormal BACnetTimeStamp
	ToFault     BACnetTimeStamp
	ToNormal    BACnetTimeStamp
}

var _ BACnetEventTimestamps = (*_BACnetEventTimestamps)(nil)

// NewBACnetEventTimestamps factory function for _BACnetEventTimestamps
func NewBACnetEventTimestamps(toOffnormal BACnetTimeStamp, toFault BACnetTimeStamp, toNormal BACnetTimeStamp) *_BACnetEventTimestamps {
	if toOffnormal == nil {
		panic("toOffnormal of type BACnetTimeStamp for BACnetEventTimestamps must not be nil")
	}
	if toFault == nil {
		panic("toFault of type BACnetTimeStamp for BACnetEventTimestamps must not be nil")
	}
	if toNormal == nil {
		panic("toNormal of type BACnetTimeStamp for BACnetEventTimestamps must not be nil")
	}
	return &_BACnetEventTimestamps{ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventTimestampsBuilder is a builder for BACnetEventTimestamps
type BACnetEventTimestampsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(toOffnormal BACnetTimeStamp, toFault BACnetTimeStamp, toNormal BACnetTimeStamp) BACnetEventTimestampsBuilder
	// WithToOffnormal adds ToOffnormal (property field)
	WithToOffnormal(BACnetTimeStamp) BACnetEventTimestampsBuilder
	// WithToOffnormalBuilder adds ToOffnormal (property field) which is build by the builder
	WithToOffnormalBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder
	// WithToFault adds ToFault (property field)
	WithToFault(BACnetTimeStamp) BACnetEventTimestampsBuilder
	// WithToFaultBuilder adds ToFault (property field) which is build by the builder
	WithToFaultBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder
	// WithToNormal adds ToNormal (property field)
	WithToNormal(BACnetTimeStamp) BACnetEventTimestampsBuilder
	// WithToNormalBuilder adds ToNormal (property field) which is build by the builder
	WithToNormalBuilder(func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder
	// Build builds the BACnetEventTimestamps or returns an error if something is wrong
	Build() (BACnetEventTimestamps, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventTimestamps
}

// NewBACnetEventTimestampsBuilder() creates a BACnetEventTimestampsBuilder
func NewBACnetEventTimestampsBuilder() BACnetEventTimestampsBuilder {
	return &_BACnetEventTimestampsBuilder{_BACnetEventTimestamps: new(_BACnetEventTimestamps)}
}

type _BACnetEventTimestampsBuilder struct {
	*_BACnetEventTimestamps

	err *utils.MultiError
}

var _ (BACnetEventTimestampsBuilder) = (*_BACnetEventTimestampsBuilder)(nil)

func (m *_BACnetEventTimestampsBuilder) WithMandatoryFields(toOffnormal BACnetTimeStamp, toFault BACnetTimeStamp, toNormal BACnetTimeStamp) BACnetEventTimestampsBuilder {
	return m.WithToOffnormal(toOffnormal).WithToFault(toFault).WithToNormal(toNormal)
}

func (m *_BACnetEventTimestampsBuilder) WithToOffnormal(toOffnormal BACnetTimeStamp) BACnetEventTimestampsBuilder {
	m.ToOffnormal = toOffnormal
	return m
}

func (m *_BACnetEventTimestampsBuilder) WithToOffnormalBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder {
	builder := builderSupplier(m.ToOffnormal.CreateBACnetTimeStampBuilder())
	var err error
	m.ToOffnormal, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return m
}

func (m *_BACnetEventTimestampsBuilder) WithToFault(toFault BACnetTimeStamp) BACnetEventTimestampsBuilder {
	m.ToFault = toFault
	return m
}

func (m *_BACnetEventTimestampsBuilder) WithToFaultBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder {
	builder := builderSupplier(m.ToFault.CreateBACnetTimeStampBuilder())
	var err error
	m.ToFault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return m
}

func (m *_BACnetEventTimestampsBuilder) WithToNormal(toNormal BACnetTimeStamp) BACnetEventTimestampsBuilder {
	m.ToNormal = toNormal
	return m
}

func (m *_BACnetEventTimestampsBuilder) WithToNormalBuilder(builderSupplier func(BACnetTimeStampBuilder) BACnetTimeStampBuilder) BACnetEventTimestampsBuilder {
	builder := builderSupplier(m.ToNormal.CreateBACnetTimeStampBuilder())
	var err error
	m.ToNormal, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTimeStampBuilder failed"))
	}
	return m
}

func (m *_BACnetEventTimestampsBuilder) Build() (BACnetEventTimestamps, error) {
	if m.ToOffnormal == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'toOffnormal' not set"))
	}
	if m.ToFault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'toFault' not set"))
	}
	if m.ToNormal == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'toNormal' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEventTimestamps.deepCopy(), nil
}

func (m *_BACnetEventTimestampsBuilder) MustBuild() BACnetEventTimestamps {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventTimestampsBuilder) DeepCopy() any {
	return m.CreateBACnetEventTimestampsBuilder()
}

// CreateBACnetEventTimestampsBuilder creates a BACnetEventTimestampsBuilder
func (m *_BACnetEventTimestamps) CreateBACnetEventTimestampsBuilder() BACnetEventTimestampsBuilder {
	if m == nil {
		return NewBACnetEventTimestampsBuilder()
	}
	return &_BACnetEventTimestampsBuilder{_BACnetEventTimestamps: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTimestamps) GetToOffnormal() BACnetTimeStamp {
	return m.ToOffnormal
}

func (m *_BACnetEventTimestamps) GetToFault() BACnetTimeStamp {
	return m.ToFault
}

func (m *_BACnetEventTimestamps) GetToNormal() BACnetTimeStamp {
	return m.ToNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventTimestamps(structType any) BACnetEventTimestamps {
	if casted, ok := structType.(BACnetEventTimestamps); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTimestamps); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTimestamps) GetTypeName() string {
	return "BACnetEventTimestamps"
}

func (m *_BACnetEventTimestamps) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits(ctx)

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits(ctx)

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventTimestamps) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTimestampsParse(ctx context.Context, theBytes []byte) (BACnetEventTimestamps, error) {
	return BACnetEventTimestampsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventTimestampsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
		return BACnetEventTimestampsParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventTimestampsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
	v, err := (&_BACnetEventTimestamps{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventTimestamps) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventTimestamps BACnetEventTimestamps, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTimestamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTimestamps")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	toOffnormal, err := ReadSimpleField[BACnetTimeStamp](ctx, "toOffnormal", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormal' field"))
	}
	m.ToOffnormal = toOffnormal

	toFault, err := ReadSimpleField[BACnetTimeStamp](ctx, "toFault", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFault' field"))
	}
	m.ToFault = toFault

	toNormal, err := ReadSimpleField[BACnetTimeStamp](ctx, "toNormal", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormal' field"))
	}
	m.ToNormal = toNormal

	if closeErr := readBuffer.CloseContext("BACnetEventTimestamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTimestamps")
	}

	return m, nil
}

func (m *_BACnetEventTimestamps) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTimestamps) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventTimestamps"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTimestamps")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toOffnormal", m.GetToOffnormal(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toOffnormal' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toFault", m.GetToFault(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toFault' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toNormal", m.GetToNormal(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTimestamps"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTimestamps")
	}
	return nil
}

func (m *_BACnetEventTimestamps) IsBACnetEventTimestamps() {}

func (m *_BACnetEventTimestamps) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventTimestamps) deepCopy() *_BACnetEventTimestamps {
	if m == nil {
		return nil
	}
	_BACnetEventTimestampsCopy := &_BACnetEventTimestamps{
		m.ToOffnormal.DeepCopy().(BACnetTimeStamp),
		m.ToFault.DeepCopy().(BACnetTimeStamp),
		m.ToNormal.DeepCopy().(BACnetTimeStamp),
	}
	return _BACnetEventTimestampsCopy
}

func (m *_BACnetEventTimestamps) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
