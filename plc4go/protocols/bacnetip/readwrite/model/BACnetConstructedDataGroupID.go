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

// BACnetConstructedDataGroupID is the corresponding interface of BACnetConstructedDataGroupID
type BACnetConstructedDataGroupID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetGroupId returns GroupId (property field)
	GetGroupId() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataGroupID is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupID()
	// CreateBuilder creates a BACnetConstructedDataGroupIDBuilder
	CreateBACnetConstructedDataGroupIDBuilder() BACnetConstructedDataGroupIDBuilder
}

// _BACnetConstructedDataGroupID is the data-structure of this message
type _BACnetConstructedDataGroupID struct {
	BACnetConstructedDataContract
	GroupId BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataGroupID = (*_BACnetConstructedDataGroupID)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupID)(nil)

// NewBACnetConstructedDataGroupID factory function for _BACnetConstructedDataGroupID
func NewBACnetConstructedDataGroupID(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, groupId BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupID {
	if groupId == nil {
		panic("groupId of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataGroupID must not be nil")
	}
	_result := &_BACnetConstructedDataGroupID{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		GroupId:                       groupId,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupIDBuilder is a builder for BACnetConstructedDataGroupID
type BACnetConstructedDataGroupIDBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(groupId BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupIDBuilder
	// WithGroupId adds GroupId (property field)
	WithGroupId(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupIDBuilder
	// WithGroupIdBuilder adds GroupId (property field) which is build by the builder
	WithGroupIdBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupIDBuilder
	// Build builds the BACnetConstructedDataGroupID or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupID, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupID
}

// NewBACnetConstructedDataGroupIDBuilder() creates a BACnetConstructedDataGroupIDBuilder
func NewBACnetConstructedDataGroupIDBuilder() BACnetConstructedDataGroupIDBuilder {
	return &_BACnetConstructedDataGroupIDBuilder{_BACnetConstructedDataGroupID: new(_BACnetConstructedDataGroupID)}
}

type _BACnetConstructedDataGroupIDBuilder struct {
	*_BACnetConstructedDataGroupID

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupIDBuilder) = (*_BACnetConstructedDataGroupIDBuilder)(nil)

func (m *_BACnetConstructedDataGroupIDBuilder) WithMandatoryFields(groupId BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupIDBuilder {
	return m.WithGroupId(groupId)
}

func (m *_BACnetConstructedDataGroupIDBuilder) WithGroupId(groupId BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupIDBuilder {
	m.GroupId = groupId
	return m
}

func (m *_BACnetConstructedDataGroupIDBuilder) WithGroupIdBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupIDBuilder {
	builder := builderSupplier(m.GroupId.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.GroupId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataGroupIDBuilder) Build() (BACnetConstructedDataGroupID, error) {
	if m.GroupId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'groupId' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataGroupID.deepCopy(), nil
}

func (m *_BACnetConstructedDataGroupIDBuilder) MustBuild() BACnetConstructedDataGroupID {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataGroupIDBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataGroupIDBuilder()
}

// CreateBACnetConstructedDataGroupIDBuilder creates a BACnetConstructedDataGroupIDBuilder
func (m *_BACnetConstructedDataGroupID) CreateBACnetConstructedDataGroupIDBuilder() BACnetConstructedDataGroupIDBuilder {
	if m == nil {
		return NewBACnetConstructedDataGroupIDBuilder()
	}
	return &_BACnetConstructedDataGroupIDBuilder{_BACnetConstructedDataGroupID: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupID) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetGroupId() BACnetApplicationTagUnsignedInteger {
	return m.GroupId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetGroupId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupID(structType any) BACnetConstructedDataGroupID {
	if casted, ok := structType.(BACnetConstructedDataGroupID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupID) GetTypeName() string {
	return "BACnetConstructedDataGroupID"
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (groupId)
	lengthInBits += m.GroupId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupID) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupID BACnetConstructedDataGroupID, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	groupId, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "groupId", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupId' field"))
	}
	m.GroupId = groupId

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), groupId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupID")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupID")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "groupId", m.GetGroupId(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'groupId' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupID")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupID) IsBACnetConstructedDataGroupID() {}

func (m *_BACnetConstructedDataGroupID) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupID) deepCopy() *_BACnetConstructedDataGroupID {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupIDCopy := &_BACnetConstructedDataGroupID{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.GroupId.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupIDCopy
}

func (m *_BACnetConstructedDataGroupID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
