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

// BACnetFaultParameterFaultLifeSafety is the corresponding interface of BACnetFaultParameterFaultLifeSafety
type BACnetFaultParameterFaultLifeSafety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultLifeSafety is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultLifeSafety()
	// CreateBuilder creates a BACnetFaultParameterFaultLifeSafetyBuilder
	CreateBACnetFaultParameterFaultLifeSafetyBuilder() BACnetFaultParameterFaultLifeSafetyBuilder
}

// _BACnetFaultParameterFaultLifeSafety is the data-structure of this message
type _BACnetFaultParameterFaultLifeSafety struct {
	BACnetFaultParameterContract
	OpeningTag            BACnetOpeningTag
	ListOfFaultValues     BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	ModePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag            BACnetClosingTag
}

var _ BACnetFaultParameterFaultLifeSafety = (*_BACnetFaultParameterFaultLifeSafety)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultLifeSafety)(nil)

// NewBACnetFaultParameterFaultLifeSafety factory function for _BACnetFaultParameterFaultLifeSafety
func NewBACnetFaultParameterFaultLifeSafety(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultLifeSafety {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if listOfFaultValues == nil {
		panic("listOfFaultValues of type BACnetFaultParameterFaultLifeSafetyListOfFaultValues for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if modePropertyReference == nil {
		panic("modePropertyReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultLifeSafety must not be nil")
	}
	_result := &_BACnetFaultParameterFaultLifeSafety{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfFaultValues:            listOfFaultValues,
		ModePropertyReference:        modePropertyReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultLifeSafetyBuilder is a builder for BACnetFaultParameterFaultLifeSafety
type BACnetFaultParameterFaultLifeSafetyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithListOfFaultValues adds ListOfFaultValues (property field)
	WithListOfFaultValues(BACnetFaultParameterFaultLifeSafetyListOfFaultValues) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithListOfFaultValuesBuilder adds ListOfFaultValues (property field) which is build by the builder
	WithListOfFaultValuesBuilder(func(BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithModePropertyReference adds ModePropertyReference (property field)
	WithModePropertyReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithModePropertyReferenceBuilder adds ModePropertyReference (property field) which is build by the builder
	WithModePropertyReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultLifeSafetyBuilder
	// Build builds the BACnetFaultParameterFaultLifeSafety or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultLifeSafety, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultLifeSafety
}

// NewBACnetFaultParameterFaultLifeSafetyBuilder() creates a BACnetFaultParameterFaultLifeSafetyBuilder
func NewBACnetFaultParameterFaultLifeSafetyBuilder() BACnetFaultParameterFaultLifeSafetyBuilder {
	return &_BACnetFaultParameterFaultLifeSafetyBuilder{_BACnetFaultParameterFaultLifeSafety: new(_BACnetFaultParameterFaultLifeSafety)}
}

type _BACnetFaultParameterFaultLifeSafetyBuilder struct {
	*_BACnetFaultParameterFaultLifeSafety

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultLifeSafetyBuilder) = (*_BACnetFaultParameterFaultLifeSafetyBuilder)(nil)

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyBuilder {
	return m.WithOpeningTag(openingTag).WithListOfFaultValues(listOfFaultValues).WithModePropertyReference(modePropertyReference).WithClosingTag(closingTag)
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetFaultParameterFaultLifeSafetyBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultLifeSafetyBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithListOfFaultValues(listOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues) BACnetFaultParameterFaultLifeSafetyBuilder {
	m.ListOfFaultValues = listOfFaultValues
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithListOfFaultValuesBuilder(builderSupplier func(BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) BACnetFaultParameterFaultLifeSafetyBuilder {
	builder := builderSupplier(m.ListOfFaultValues.CreateBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder())
	var err error
	m.ListOfFaultValues, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithModePropertyReference(modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetFaultParameterFaultLifeSafetyBuilder {
	m.ModePropertyReference = modePropertyReference
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithModePropertyReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetFaultParameterFaultLifeSafetyBuilder {
	builder := builderSupplier(m.ModePropertyReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	m.ModePropertyReference, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultLifeSafetyBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) Build() (BACnetFaultParameterFaultLifeSafety, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.ListOfFaultValues == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'listOfFaultValues' not set"))
	}
	if m.ModePropertyReference == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'modePropertyReference' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetFaultParameterFaultLifeSafety.deepCopy(), nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) MustBuild() BACnetFaultParameterFaultLifeSafety {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetFaultParameterFaultLifeSafetyBuilder) DeepCopy() any {
	return m.CreateBACnetFaultParameterFaultLifeSafetyBuilder()
}

// CreateBACnetFaultParameterFaultLifeSafetyBuilder creates a BACnetFaultParameterFaultLifeSafetyBuilder
func (m *_BACnetFaultParameterFaultLifeSafety) CreateBACnetFaultParameterFaultLifeSafetyBuilder() BACnetFaultParameterFaultLifeSafetyBuilder {
	if m == nil {
		return NewBACnetFaultParameterFaultLifeSafetyBuilder()
	}
	return &_BACnetFaultParameterFaultLifeSafetyBuilder{_BACnetFaultParameterFaultLifeSafety: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultLifeSafety) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultLifeSafety) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetListOfFaultValues() BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultLifeSafety(structType any) BACnetFaultParameterFaultLifeSafety {
	if casted, ok := structType.(BACnetFaultParameterFaultLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetTypeName() string {
	return "BACnetFaultParameterFaultLifeSafety"
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultLifeSafety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultLifeSafety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultLifeSafety BACnetFaultParameterFaultLifeSafety, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultLifeSafety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	modePropertyReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modePropertyReference' field"))
	}
	m.ModePropertyReference = modePropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultLifeSafety")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultLifeSafety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultLifeSafety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultLifeSafetyListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", m.GetModePropertyReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'modePropertyReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultLifeSafety")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultLifeSafety) IsBACnetFaultParameterFaultLifeSafety() {}

func (m *_BACnetFaultParameterFaultLifeSafety) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultLifeSafety) deepCopy() *_BACnetFaultParameterFaultLifeSafety {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultLifeSafetyCopy := &_BACnetFaultParameterFaultLifeSafety{
		m.BACnetFaultParameterContract.(*_BACnetFaultParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ListOfFaultValues.DeepCopy().(BACnetFaultParameterFaultLifeSafetyListOfFaultValues),
		m.ModePropertyReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultLifeSafetyCopy
}

func (m *_BACnetFaultParameterFaultLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
