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

// BACnetConstructedDataNegativeAccessRules is the corresponding interface of BACnetConstructedDataNegativeAccessRules
type BACnetConstructedDataNegativeAccessRules interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetNegativeAccessRules returns NegativeAccessRules (property field)
	GetNegativeAccessRules() []BACnetAccessRule
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataNegativeAccessRules is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNegativeAccessRules()
	// CreateBuilder creates a BACnetConstructedDataNegativeAccessRulesBuilder
	CreateBACnetConstructedDataNegativeAccessRulesBuilder() BACnetConstructedDataNegativeAccessRulesBuilder
}

// _BACnetConstructedDataNegativeAccessRules is the data-structure of this message
type _BACnetConstructedDataNegativeAccessRules struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	NegativeAccessRules  []BACnetAccessRule
}

var _ BACnetConstructedDataNegativeAccessRules = (*_BACnetConstructedDataNegativeAccessRules)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNegativeAccessRules)(nil)

// NewBACnetConstructedDataNegativeAccessRules factory function for _BACnetConstructedDataNegativeAccessRules
func NewBACnetConstructedDataNegativeAccessRules(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, negativeAccessRules []BACnetAccessRule, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNegativeAccessRules {
	_result := &_BACnetConstructedDataNegativeAccessRules{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		NegativeAccessRules:           negativeAccessRules,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNegativeAccessRulesBuilder is a builder for BACnetConstructedDataNegativeAccessRules
type BACnetConstructedDataNegativeAccessRulesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(negativeAccessRules []BACnetAccessRule) BACnetConstructedDataNegativeAccessRulesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNegativeAccessRulesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNegativeAccessRulesBuilder
	// WithNegativeAccessRules adds NegativeAccessRules (property field)
	WithNegativeAccessRules(...BACnetAccessRule) BACnetConstructedDataNegativeAccessRulesBuilder
	// Build builds the BACnetConstructedDataNegativeAccessRules or returns an error if something is wrong
	Build() (BACnetConstructedDataNegativeAccessRules, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNegativeAccessRules
}

// NewBACnetConstructedDataNegativeAccessRulesBuilder() creates a BACnetConstructedDataNegativeAccessRulesBuilder
func NewBACnetConstructedDataNegativeAccessRulesBuilder() BACnetConstructedDataNegativeAccessRulesBuilder {
	return &_BACnetConstructedDataNegativeAccessRulesBuilder{_BACnetConstructedDataNegativeAccessRules: new(_BACnetConstructedDataNegativeAccessRules)}
}

type _BACnetConstructedDataNegativeAccessRulesBuilder struct {
	*_BACnetConstructedDataNegativeAccessRules

	err *utils.MultiError
}

var _ (BACnetConstructedDataNegativeAccessRulesBuilder) = (*_BACnetConstructedDataNegativeAccessRulesBuilder)(nil)

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) WithMandatoryFields(negativeAccessRules []BACnetAccessRule) BACnetConstructedDataNegativeAccessRulesBuilder {
	return m.WithNegativeAccessRules(negativeAccessRules...)
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNegativeAccessRulesBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNegativeAccessRulesBuilder {
	builder := builderSupplier(m.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) WithNegativeAccessRules(negativeAccessRules ...BACnetAccessRule) BACnetConstructedDataNegativeAccessRulesBuilder {
	m.NegativeAccessRules = negativeAccessRules
	return m
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) Build() (BACnetConstructedDataNegativeAccessRules, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataNegativeAccessRules.deepCopy(), nil
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) MustBuild() BACnetConstructedDataNegativeAccessRules {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataNegativeAccessRulesBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataNegativeAccessRulesBuilder()
}

// CreateBACnetConstructedDataNegativeAccessRulesBuilder creates a BACnetConstructedDataNegativeAccessRulesBuilder
func (m *_BACnetConstructedDataNegativeAccessRules) CreateBACnetConstructedDataNegativeAccessRulesBuilder() BACnetConstructedDataNegativeAccessRulesBuilder {
	if m == nil {
		return NewBACnetConstructedDataNegativeAccessRulesBuilder()
	}
	return &_BACnetConstructedDataNegativeAccessRulesBuilder{_BACnetConstructedDataNegativeAccessRules: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNegativeAccessRules) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNegativeAccessRules) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NEGATIVE_ACCESS_RULES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNegativeAccessRules) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNegativeAccessRules) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataNegativeAccessRules) GetNegativeAccessRules() []BACnetAccessRule {
	return m.NegativeAccessRules
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNegativeAccessRules) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNegativeAccessRules(structType any) BACnetConstructedDataNegativeAccessRules {
	if casted, ok := structType.(BACnetConstructedDataNegativeAccessRules); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNegativeAccessRules); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNegativeAccessRules) GetTypeName() string {
	return "BACnetConstructedDataNegativeAccessRules"
}

func (m *_BACnetConstructedDataNegativeAccessRules) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.NegativeAccessRules) > 0 {
		for _, element := range m.NegativeAccessRules {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataNegativeAccessRules) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNegativeAccessRules) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNegativeAccessRules BACnetConstructedDataNegativeAccessRules, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNegativeAccessRules"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNegativeAccessRules")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	negativeAccessRules, err := ReadTerminatedArrayField[BACnetAccessRule](ctx, "negativeAccessRules", ReadComplex[BACnetAccessRule](BACnetAccessRuleParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'negativeAccessRules' field"))
	}
	m.NegativeAccessRules = negativeAccessRules

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNegativeAccessRules"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNegativeAccessRules")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNegativeAccessRules) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNegativeAccessRules) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNegativeAccessRules"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNegativeAccessRules")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "negativeAccessRules", m.GetNegativeAccessRules(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'negativeAccessRules' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNegativeAccessRules"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNegativeAccessRules")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNegativeAccessRules) IsBACnetConstructedDataNegativeAccessRules() {}

func (m *_BACnetConstructedDataNegativeAccessRules) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNegativeAccessRules) deepCopy() *_BACnetConstructedDataNegativeAccessRules {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNegativeAccessRulesCopy := &_BACnetConstructedDataNegativeAccessRules{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetAccessRule, BACnetAccessRule](m.NegativeAccessRules),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNegativeAccessRulesCopy
}

func (m *_BACnetConstructedDataNegativeAccessRules) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
