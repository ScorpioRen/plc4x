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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimeValueAll is the corresponding interface of BACnetConstructedDataTimeValueAll
type BACnetConstructedDataTimeValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataTimeValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeValueAll()
	// CreateBuilder creates a BACnetConstructedDataTimeValueAllBuilder
	CreateBACnetConstructedDataTimeValueAllBuilder() BACnetConstructedDataTimeValueAllBuilder
}

// _BACnetConstructedDataTimeValueAll is the data-structure of this message
type _BACnetConstructedDataTimeValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataTimeValueAll = (*_BACnetConstructedDataTimeValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeValueAll)(nil)

// NewBACnetConstructedDataTimeValueAll factory function for _BACnetConstructedDataTimeValueAll
func NewBACnetConstructedDataTimeValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeValueAll {
	_result := &_BACnetConstructedDataTimeValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimeValueAllBuilder is a builder for BACnetConstructedDataTimeValueAll
type BACnetConstructedDataTimeValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataTimeValueAllBuilder
	// Build builds the BACnetConstructedDataTimeValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeValueAll
}

// NewBACnetConstructedDataTimeValueAllBuilder() creates a BACnetConstructedDataTimeValueAllBuilder
func NewBACnetConstructedDataTimeValueAllBuilder() BACnetConstructedDataTimeValueAllBuilder {
	return &_BACnetConstructedDataTimeValueAllBuilder{_BACnetConstructedDataTimeValueAll: new(_BACnetConstructedDataTimeValueAll)}
}

type _BACnetConstructedDataTimeValueAllBuilder struct {
	*_BACnetConstructedDataTimeValueAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeValueAllBuilder) = (*_BACnetConstructedDataTimeValueAllBuilder)(nil)

func (m *_BACnetConstructedDataTimeValueAllBuilder) WithMandatoryFields() BACnetConstructedDataTimeValueAllBuilder {
	return m
}

func (m *_BACnetConstructedDataTimeValueAllBuilder) Build() (BACnetConstructedDataTimeValueAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataTimeValueAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataTimeValueAllBuilder) MustBuild() BACnetConstructedDataTimeValueAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataTimeValueAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataTimeValueAllBuilder()
}

// CreateBACnetConstructedDataTimeValueAllBuilder creates a BACnetConstructedDataTimeValueAllBuilder
func (m *_BACnetConstructedDataTimeValueAll) CreateBACnetConstructedDataTimeValueAllBuilder() BACnetConstructedDataTimeValueAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataTimeValueAllBuilder()
	}
	return &_BACnetConstructedDataTimeValueAllBuilder{_BACnetConstructedDataTimeValueAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIME_VALUE
}

func (m *_BACnetConstructedDataTimeValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeValueAll(structType any) BACnetConstructedDataTimeValueAll {
	if casted, ok := structType.(BACnetConstructedDataTimeValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeValueAll) GetTypeName() string {
	return "BACnetConstructedDataTimeValueAll"
}

func (m *_BACnetConstructedDataTimeValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeValueAll BACnetConstructedDataTimeValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeValueAll) IsBACnetConstructedDataTimeValueAll() {}

func (m *_BACnetConstructedDataTimeValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeValueAll) deepCopy() *_BACnetConstructedDataTimeValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeValueAllCopy := &_BACnetConstructedDataTimeValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeValueAllCopy
}

func (m *_BACnetConstructedDataTimeValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
