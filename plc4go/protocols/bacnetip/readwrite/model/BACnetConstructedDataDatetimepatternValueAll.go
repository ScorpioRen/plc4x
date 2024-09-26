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

// BACnetConstructedDataDatetimepatternValueAll is the corresponding interface of BACnetConstructedDataDatetimepatternValueAll
type BACnetConstructedDataDatetimepatternValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataDatetimepatternValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDatetimepatternValueAll()
	// CreateBuilder creates a BACnetConstructedDataDatetimepatternValueAllBuilder
	CreateBACnetConstructedDataDatetimepatternValueAllBuilder() BACnetConstructedDataDatetimepatternValueAllBuilder
}

// _BACnetConstructedDataDatetimepatternValueAll is the data-structure of this message
type _BACnetConstructedDataDatetimepatternValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataDatetimepatternValueAll = (*_BACnetConstructedDataDatetimepatternValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDatetimepatternValueAll)(nil)

// NewBACnetConstructedDataDatetimepatternValueAll factory function for _BACnetConstructedDataDatetimepatternValueAll
func NewBACnetConstructedDataDatetimepatternValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatetimepatternValueAll {
	_result := &_BACnetConstructedDataDatetimepatternValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDatetimepatternValueAllBuilder is a builder for BACnetConstructedDataDatetimepatternValueAll
type BACnetConstructedDataDatetimepatternValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataDatetimepatternValueAllBuilder
	// Build builds the BACnetConstructedDataDatetimepatternValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataDatetimepatternValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDatetimepatternValueAll
}

// NewBACnetConstructedDataDatetimepatternValueAllBuilder() creates a BACnetConstructedDataDatetimepatternValueAllBuilder
func NewBACnetConstructedDataDatetimepatternValueAllBuilder() BACnetConstructedDataDatetimepatternValueAllBuilder {
	return &_BACnetConstructedDataDatetimepatternValueAllBuilder{_BACnetConstructedDataDatetimepatternValueAll: new(_BACnetConstructedDataDatetimepatternValueAll)}
}

type _BACnetConstructedDataDatetimepatternValueAllBuilder struct {
	*_BACnetConstructedDataDatetimepatternValueAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataDatetimepatternValueAllBuilder) = (*_BACnetConstructedDataDatetimepatternValueAllBuilder)(nil)

func (m *_BACnetConstructedDataDatetimepatternValueAllBuilder) WithMandatoryFields() BACnetConstructedDataDatetimepatternValueAllBuilder {
	return m
}

func (m *_BACnetConstructedDataDatetimepatternValueAllBuilder) Build() (BACnetConstructedDataDatetimepatternValueAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataDatetimepatternValueAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAllBuilder) MustBuild() BACnetConstructedDataDatetimepatternValueAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataDatetimepatternValueAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataDatetimepatternValueAllBuilder()
}

// CreateBACnetConstructedDataDatetimepatternValueAllBuilder creates a BACnetConstructedDataDatetimepatternValueAllBuilder
func (m *_BACnetConstructedDataDatetimepatternValueAll) CreateBACnetConstructedDataDatetimepatternValueAllBuilder() BACnetConstructedDataDatetimepatternValueAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataDatetimepatternValueAllBuilder()
	}
	return &_BACnetConstructedDataDatetimepatternValueAllBuilder{_BACnetConstructedDataDatetimepatternValueAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIMEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatetimepatternValueAll(structType any) BACnetConstructedDataDatetimepatternValueAll {
	if casted, ok := structType.(BACnetConstructedDataDatetimepatternValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatetimepatternValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetTypeName() string {
	return "BACnetConstructedDataDatetimepatternValueAll"
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDatetimepatternValueAll BACnetConstructedDataDatetimepatternValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatetimepatternValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatetimepatternValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatetimepatternValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatetimepatternValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatetimepatternValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatetimepatternValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatetimepatternValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatetimepatternValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) IsBACnetConstructedDataDatetimepatternValueAll() {
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) deepCopy() *_BACnetConstructedDataDatetimepatternValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDatetimepatternValueAllCopy := &_BACnetConstructedDataDatetimepatternValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDatetimepatternValueAllCopy
}

func (m *_BACnetConstructedDataDatetimepatternValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
