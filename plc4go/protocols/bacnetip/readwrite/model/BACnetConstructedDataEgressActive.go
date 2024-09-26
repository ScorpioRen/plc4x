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

// BACnetConstructedDataEgressActive is the corresponding interface of BACnetConstructedDataEgressActive
type BACnetConstructedDataEgressActive interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEgressActive returns EgressActive (property field)
	GetEgressActive() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataEgressActive is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEgressActive()
	// CreateBuilder creates a BACnetConstructedDataEgressActiveBuilder
	CreateBACnetConstructedDataEgressActiveBuilder() BACnetConstructedDataEgressActiveBuilder
}

// _BACnetConstructedDataEgressActive is the data-structure of this message
type _BACnetConstructedDataEgressActive struct {
	BACnetConstructedDataContract
	EgressActive BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataEgressActive = (*_BACnetConstructedDataEgressActive)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEgressActive)(nil)

// NewBACnetConstructedDataEgressActive factory function for _BACnetConstructedDataEgressActive
func NewBACnetConstructedDataEgressActive(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, egressActive BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEgressActive {
	if egressActive == nil {
		panic("egressActive of type BACnetApplicationTagBoolean for BACnetConstructedDataEgressActive must not be nil")
	}
	_result := &_BACnetConstructedDataEgressActive{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EgressActive:                  egressActive,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEgressActiveBuilder is a builder for BACnetConstructedDataEgressActive
type BACnetConstructedDataEgressActiveBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(egressActive BACnetApplicationTagBoolean) BACnetConstructedDataEgressActiveBuilder
	// WithEgressActive adds EgressActive (property field)
	WithEgressActive(BACnetApplicationTagBoolean) BACnetConstructedDataEgressActiveBuilder
	// WithEgressActiveBuilder adds EgressActive (property field) which is build by the builder
	WithEgressActiveBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEgressActiveBuilder
	// Build builds the BACnetConstructedDataEgressActive or returns an error if something is wrong
	Build() (BACnetConstructedDataEgressActive, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEgressActive
}

// NewBACnetConstructedDataEgressActiveBuilder() creates a BACnetConstructedDataEgressActiveBuilder
func NewBACnetConstructedDataEgressActiveBuilder() BACnetConstructedDataEgressActiveBuilder {
	return &_BACnetConstructedDataEgressActiveBuilder{_BACnetConstructedDataEgressActive: new(_BACnetConstructedDataEgressActive)}
}

type _BACnetConstructedDataEgressActiveBuilder struct {
	*_BACnetConstructedDataEgressActive

	err *utils.MultiError
}

var _ (BACnetConstructedDataEgressActiveBuilder) = (*_BACnetConstructedDataEgressActiveBuilder)(nil)

func (m *_BACnetConstructedDataEgressActiveBuilder) WithMandatoryFields(egressActive BACnetApplicationTagBoolean) BACnetConstructedDataEgressActiveBuilder {
	return m.WithEgressActive(egressActive)
}

func (m *_BACnetConstructedDataEgressActiveBuilder) WithEgressActive(egressActive BACnetApplicationTagBoolean) BACnetConstructedDataEgressActiveBuilder {
	m.EgressActive = egressActive
	return m
}

func (m *_BACnetConstructedDataEgressActiveBuilder) WithEgressActiveBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataEgressActiveBuilder {
	builder := builderSupplier(m.EgressActive.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.EgressActive, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataEgressActiveBuilder) Build() (BACnetConstructedDataEgressActive, error) {
	if m.EgressActive == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'egressActive' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataEgressActive.deepCopy(), nil
}

func (m *_BACnetConstructedDataEgressActiveBuilder) MustBuild() BACnetConstructedDataEgressActive {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataEgressActiveBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataEgressActiveBuilder()
}

// CreateBACnetConstructedDataEgressActiveBuilder creates a BACnetConstructedDataEgressActiveBuilder
func (m *_BACnetConstructedDataEgressActive) CreateBACnetConstructedDataEgressActiveBuilder() BACnetConstructedDataEgressActiveBuilder {
	if m == nil {
		return NewBACnetConstructedDataEgressActiveBuilder()
	}
	return &_BACnetConstructedDataEgressActiveBuilder{_BACnetConstructedDataEgressActive: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEgressActive) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEgressActive) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EGRESS_ACTIVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEgressActive) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEgressActive) GetEgressActive() BACnetApplicationTagBoolean {
	return m.EgressActive
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEgressActive) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetEgressActive())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEgressActive(structType any) BACnetConstructedDataEgressActive {
	if casted, ok := structType.(BACnetConstructedDataEgressActive); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEgressActive); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEgressActive) GetTypeName() string {
	return "BACnetConstructedDataEgressActive"
}

func (m *_BACnetConstructedDataEgressActive) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (egressActive)
	lengthInBits += m.EgressActive.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEgressActive) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEgressActive) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEgressActive BACnetConstructedDataEgressActive, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEgressActive"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEgressActive")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	egressActive, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "egressActive", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'egressActive' field"))
	}
	m.EgressActive = egressActive

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), egressActive)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEgressActive"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEgressActive")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEgressActive) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEgressActive) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEgressActive"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEgressActive")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "egressActive", m.GetEgressActive(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'egressActive' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEgressActive"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEgressActive")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEgressActive) IsBACnetConstructedDataEgressActive() {}

func (m *_BACnetConstructedDataEgressActive) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEgressActive) deepCopy() *_BACnetConstructedDataEgressActive {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEgressActiveCopy := &_BACnetConstructedDataEgressActive{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.EgressActive.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEgressActiveCopy
}

func (m *_BACnetConstructedDataEgressActive) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
