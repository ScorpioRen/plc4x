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

// BACnetConstructedDataDatepatternValueAll is the corresponding interface of BACnetConstructedDataDatepatternValueAll
type BACnetConstructedDataDatepatternValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataDatepatternValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDatepatternValueAll()
	// CreateBuilder creates a BACnetConstructedDataDatepatternValueAllBuilder
	CreateBACnetConstructedDataDatepatternValueAllBuilder() BACnetConstructedDataDatepatternValueAllBuilder
}

// _BACnetConstructedDataDatepatternValueAll is the data-structure of this message
type _BACnetConstructedDataDatepatternValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataDatepatternValueAll = (*_BACnetConstructedDataDatepatternValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDatepatternValueAll)(nil)

// NewBACnetConstructedDataDatepatternValueAll factory function for _BACnetConstructedDataDatepatternValueAll
func NewBACnetConstructedDataDatepatternValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatepatternValueAll {
	_result := &_BACnetConstructedDataDatepatternValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDatepatternValueAllBuilder is a builder for BACnetConstructedDataDatepatternValueAll
type BACnetConstructedDataDatepatternValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataDatepatternValueAllBuilder
	// Build builds the BACnetConstructedDataDatepatternValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataDatepatternValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDatepatternValueAll
}

// NewBACnetConstructedDataDatepatternValueAllBuilder() creates a BACnetConstructedDataDatepatternValueAllBuilder
func NewBACnetConstructedDataDatepatternValueAllBuilder() BACnetConstructedDataDatepatternValueAllBuilder {
	return &_BACnetConstructedDataDatepatternValueAllBuilder{_BACnetConstructedDataDatepatternValueAll: new(_BACnetConstructedDataDatepatternValueAll)}
}

type _BACnetConstructedDataDatepatternValueAllBuilder struct {
	*_BACnetConstructedDataDatepatternValueAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataDatepatternValueAllBuilder) = (*_BACnetConstructedDataDatepatternValueAllBuilder)(nil)

func (m *_BACnetConstructedDataDatepatternValueAllBuilder) WithMandatoryFields() BACnetConstructedDataDatepatternValueAllBuilder {
	return m
}

func (m *_BACnetConstructedDataDatepatternValueAllBuilder) Build() (BACnetConstructedDataDatepatternValueAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataDatepatternValueAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataDatepatternValueAllBuilder) MustBuild() BACnetConstructedDataDatepatternValueAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataDatepatternValueAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataDatepatternValueAllBuilder()
}

// CreateBACnetConstructedDataDatepatternValueAllBuilder creates a BACnetConstructedDataDatepatternValueAllBuilder
func (m *_BACnetConstructedDataDatepatternValueAll) CreateBACnetConstructedDataDatepatternValueAllBuilder() BACnetConstructedDataDatepatternValueAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataDatepatternValueAllBuilder()
	}
	return &_BACnetConstructedDataDatepatternValueAllBuilder{_BACnetConstructedDataDatepatternValueAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatepatternValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATEPATTERN_VALUE
}

func (m *_BACnetConstructedDataDatepatternValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatepatternValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatepatternValueAll(structType any) BACnetConstructedDataDatepatternValueAll {
	if casted, ok := structType.(BACnetConstructedDataDatepatternValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatepatternValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatepatternValueAll) GetTypeName() string {
	return "BACnetConstructedDataDatepatternValueAll"
}

func (m *_BACnetConstructedDataDatepatternValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataDatepatternValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDatepatternValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDatepatternValueAll BACnetConstructedDataDatepatternValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatepatternValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatepatternValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatepatternValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatepatternValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDatepatternValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatepatternValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatepatternValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatepatternValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatepatternValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatepatternValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatepatternValueAll) IsBACnetConstructedDataDatepatternValueAll() {}

func (m *_BACnetConstructedDataDatepatternValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDatepatternValueAll) deepCopy() *_BACnetConstructedDataDatepatternValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDatepatternValueAllCopy := &_BACnetConstructedDataDatepatternValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDatepatternValueAllCopy
}

func (m *_BACnetConstructedDataDatepatternValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
