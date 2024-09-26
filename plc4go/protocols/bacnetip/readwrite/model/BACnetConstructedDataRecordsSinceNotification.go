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

// BACnetConstructedDataRecordsSinceNotification is the corresponding interface of BACnetConstructedDataRecordsSinceNotification
type BACnetConstructedDataRecordsSinceNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRecordsSinceNotifications returns RecordsSinceNotifications (property field)
	GetRecordsSinceNotifications() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataRecordsSinceNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRecordsSinceNotification()
	// CreateBuilder creates a BACnetConstructedDataRecordsSinceNotificationBuilder
	CreateBACnetConstructedDataRecordsSinceNotificationBuilder() BACnetConstructedDataRecordsSinceNotificationBuilder
}

// _BACnetConstructedDataRecordsSinceNotification is the data-structure of this message
type _BACnetConstructedDataRecordsSinceNotification struct {
	BACnetConstructedDataContract
	RecordsSinceNotifications BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRecordsSinceNotification = (*_BACnetConstructedDataRecordsSinceNotification)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRecordsSinceNotification)(nil)

// NewBACnetConstructedDataRecordsSinceNotification factory function for _BACnetConstructedDataRecordsSinceNotification
func NewBACnetConstructedDataRecordsSinceNotification(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, recordsSinceNotifications BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRecordsSinceNotification {
	if recordsSinceNotifications == nil {
		panic("recordsSinceNotifications of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataRecordsSinceNotification must not be nil")
	}
	_result := &_BACnetConstructedDataRecordsSinceNotification{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RecordsSinceNotifications:     recordsSinceNotifications,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRecordsSinceNotificationBuilder is a builder for BACnetConstructedDataRecordsSinceNotification
type BACnetConstructedDataRecordsSinceNotificationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recordsSinceNotifications BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordsSinceNotificationBuilder
	// WithRecordsSinceNotifications adds RecordsSinceNotifications (property field)
	WithRecordsSinceNotifications(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordsSinceNotificationBuilder
	// WithRecordsSinceNotificationsBuilder adds RecordsSinceNotifications (property field) which is build by the builder
	WithRecordsSinceNotificationsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRecordsSinceNotificationBuilder
	// Build builds the BACnetConstructedDataRecordsSinceNotification or returns an error if something is wrong
	Build() (BACnetConstructedDataRecordsSinceNotification, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRecordsSinceNotification
}

// NewBACnetConstructedDataRecordsSinceNotificationBuilder() creates a BACnetConstructedDataRecordsSinceNotificationBuilder
func NewBACnetConstructedDataRecordsSinceNotificationBuilder() BACnetConstructedDataRecordsSinceNotificationBuilder {
	return &_BACnetConstructedDataRecordsSinceNotificationBuilder{_BACnetConstructedDataRecordsSinceNotification: new(_BACnetConstructedDataRecordsSinceNotification)}
}

type _BACnetConstructedDataRecordsSinceNotificationBuilder struct {
	*_BACnetConstructedDataRecordsSinceNotification

	err *utils.MultiError
}

var _ (BACnetConstructedDataRecordsSinceNotificationBuilder) = (*_BACnetConstructedDataRecordsSinceNotificationBuilder)(nil)

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) WithMandatoryFields(recordsSinceNotifications BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordsSinceNotificationBuilder {
	return m.WithRecordsSinceNotifications(recordsSinceNotifications)
}

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) WithRecordsSinceNotifications(recordsSinceNotifications BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRecordsSinceNotificationBuilder {
	m.RecordsSinceNotifications = recordsSinceNotifications
	return m
}

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) WithRecordsSinceNotificationsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRecordsSinceNotificationBuilder {
	builder := builderSupplier(m.RecordsSinceNotifications.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.RecordsSinceNotifications, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) Build() (BACnetConstructedDataRecordsSinceNotification, error) {
	if m.RecordsSinceNotifications == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'recordsSinceNotifications' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataRecordsSinceNotification.deepCopy(), nil
}

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) MustBuild() BACnetConstructedDataRecordsSinceNotification {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataRecordsSinceNotificationBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataRecordsSinceNotificationBuilder()
}

// CreateBACnetConstructedDataRecordsSinceNotificationBuilder creates a BACnetConstructedDataRecordsSinceNotificationBuilder
func (m *_BACnetConstructedDataRecordsSinceNotification) CreateBACnetConstructedDataRecordsSinceNotificationBuilder() BACnetConstructedDataRecordsSinceNotificationBuilder {
	if m == nil {
		return NewBACnetConstructedDataRecordsSinceNotificationBuilder()
	}
	return &_BACnetConstructedDataRecordsSinceNotificationBuilder{_BACnetConstructedDataRecordsSinceNotification: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORDS_SINCE_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetRecordsSinceNotifications() BACnetApplicationTagUnsignedInteger {
	return m.RecordsSinceNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRecordsSinceNotification) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordsSinceNotifications())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRecordsSinceNotification(structType any) BACnetConstructedDataRecordsSinceNotification {
	if casted, ok := structType.(BACnetConstructedDataRecordsSinceNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecordsSinceNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetTypeName() string {
	return "BACnetConstructedDataRecordsSinceNotification"
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (recordsSinceNotifications)
	lengthInBits += m.RecordsSinceNotifications.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRecordsSinceNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRecordsSinceNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRecordsSinceNotification BACnetConstructedDataRecordsSinceNotification, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecordsSinceNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRecordsSinceNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recordsSinceNotifications, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordsSinceNotifications", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordsSinceNotifications' field"))
	}
	m.RecordsSinceNotifications = recordsSinceNotifications

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), recordsSinceNotifications)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecordsSinceNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRecordsSinceNotification")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRecordsSinceNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecordsSinceNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRecordsSinceNotification")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "recordsSinceNotifications", m.GetRecordsSinceNotifications(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'recordsSinceNotifications' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecordsSinceNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRecordsSinceNotification")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRecordsSinceNotification) IsBACnetConstructedDataRecordsSinceNotification() {
}

func (m *_BACnetConstructedDataRecordsSinceNotification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRecordsSinceNotification) deepCopy() *_BACnetConstructedDataRecordsSinceNotification {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRecordsSinceNotificationCopy := &_BACnetConstructedDataRecordsSinceNotification{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.RecordsSinceNotifications.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRecordsSinceNotificationCopy
}

func (m *_BACnetConstructedDataRecordsSinceNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
