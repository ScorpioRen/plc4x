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

// BACnetConfirmedServiceRequestConfirmedCOVNotification is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotification
type BACnetConfirmedServiceRequestConfirmedCOVNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() BACnetContextTagUnsignedInteger
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
	// IsBACnetConfirmedServiceRequestConfirmedCOVNotification is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedCOVNotification()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
}

// _BACnetConfirmedServiceRequestConfirmedCOVNotification is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedCOVNotification struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	LifetimeInSeconds           BACnetContextTagUnsignedInteger
	ListOfValues                BACnetPropertyValues
}

var _ BACnetConfirmedServiceRequestConfirmedCOVNotification = (*_BACnetConfirmedServiceRequestConfirmedCOVNotification)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestConfirmedCOVNotification)(nil)

// NewBACnetConfirmedServiceRequestConfirmedCOVNotification factory function for _BACnetConfirmedServiceRequestConfirmedCOVNotification
func NewBACnetConfirmedServiceRequestConfirmedCOVNotification(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, lifetimeInSeconds BACnetContextTagUnsignedInteger, listOfValues BACnetPropertyValues, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedCOVNotification {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotification must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestConfirmedCOVNotification must not be nil")
	}
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestConfirmedCOVNotification must not be nil")
	}
	if lifetimeInSeconds == nil {
		panic("lifetimeInSeconds of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotification must not be nil")
	}
	if listOfValues == nil {
		panic("listOfValues of type BACnetPropertyValues for BACnetConfirmedServiceRequestConfirmedCOVNotification must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedCOVNotification{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:            initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:             monitoredObjectIdentifier,
		LifetimeInSeconds:                     lifetimeInSeconds,
		ListOfValues:                          listOfValues,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder is a builder for BACnetConfirmedServiceRequestConfirmedCOVNotification
type BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, lifetimeInSeconds BACnetContextTagUnsignedInteger, listOfValues BACnetPropertyValues) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithSubscriberProcessIdentifier adds SubscriberProcessIdentifier (property field)
	WithSubscriberProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithSubscriberProcessIdentifierBuilder adds SubscriberProcessIdentifier (property field) which is build by the builder
	WithSubscriberProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithInitiatingDeviceIdentifier adds InitiatingDeviceIdentifier (property field)
	WithInitiatingDeviceIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithInitiatingDeviceIdentifierBuilder adds InitiatingDeviceIdentifier (property field) which is build by the builder
	WithInitiatingDeviceIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithMonitoredObjectIdentifier adds MonitoredObjectIdentifier (property field)
	WithMonitoredObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithMonitoredObjectIdentifierBuilder adds MonitoredObjectIdentifier (property field) which is build by the builder
	WithMonitoredObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithLifetimeInSeconds adds LifetimeInSeconds (property field)
	WithLifetimeInSeconds(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithLifetimeInSecondsBuilder adds LifetimeInSeconds (property field) which is build by the builder
	WithLifetimeInSecondsBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithListOfValues adds ListOfValues (property field)
	WithListOfValues(BACnetPropertyValues) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// WithListOfValuesBuilder adds ListOfValues (property field) which is build by the builder
	WithListOfValuesBuilder(func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedCOVNotification or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedCOVNotification, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedCOVNotification
}

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder() creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder{_BACnetConfirmedServiceRequestConfirmedCOVNotification: new(_BACnetConfirmedServiceRequestConfirmedCOVNotification)}
}

type _BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedCOVNotification

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder)(nil)

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, lifetimeInSeconds BACnetContextTagUnsignedInteger, listOfValues BACnetPropertyValues) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	return m.WithSubscriberProcessIdentifier(subscriberProcessIdentifier).WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier).WithMonitoredObjectIdentifier(monitoredObjectIdentifier).WithLifetimeInSeconds(lifetimeInSeconds).WithListOfValues(listOfValues)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithSubscriberProcessIdentifier(subscriberProcessIdentifier BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithSubscriberProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	builder := builderSupplier(m.SubscriberProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.SubscriberProcessIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithInitiatingDeviceIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	builder := builderSupplier(m.InitiatingDeviceIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.InitiatingDeviceIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithMonitoredObjectIdentifier(monitoredObjectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithMonitoredObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	builder := builderSupplier(m.MonitoredObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.MonitoredObjectIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithLifetimeInSeconds(lifetimeInSeconds BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	m.LifetimeInSeconds = lifetimeInSeconds
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithLifetimeInSecondsBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	builder := builderSupplier(m.LifetimeInSeconds.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.LifetimeInSeconds, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithListOfValues(listOfValues BACnetPropertyValues) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	m.ListOfValues = listOfValues
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) WithListOfValuesBuilder(builderSupplier func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	builder := builderSupplier(m.ListOfValues.CreateBACnetPropertyValuesBuilder())
	var err error
	m.ListOfValues, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetPropertyValuesBuilder failed"))
	}
	return m
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) Build() (BACnetConfirmedServiceRequestConfirmedCOVNotification, error) {
	if m.SubscriberProcessIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'subscriberProcessIdentifier' not set"))
	}
	if m.InitiatingDeviceIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'initiatingDeviceIdentifier' not set"))
	}
	if m.MonitoredObjectIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'monitoredObjectIdentifier' not set"))
	}
	if m.LifetimeInSeconds == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'lifetimeInSeconds' not set"))
	}
	if m.ListOfValues == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'listOfValues' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConfirmedServiceRequestConfirmedCOVNotification.deepCopy(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedCOVNotification {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder) DeepCopy() any {
	return m.CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder()
}

// CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder
func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder {
	if m == nil {
		return NewBACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedCOVNotificationBuilder{_BACnetConfirmedServiceRequestConfirmedCOVNotification: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLifetimeInSeconds() BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedCOVNotification(structType any) BACnetConfirmedServiceRequestConfirmedCOVNotification {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotification"
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits(ctx)

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestConfirmedCOVNotification BACnetConfirmedServiceRequestConfirmedCOVNotification, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	lifetimeInSeconds, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetimeInSeconds' field"))
	}
	m.LifetimeInSeconds = lifetimeInSeconds

	listOfValues, err := ReadSimpleField[BACnetPropertyValues](ctx, "listOfValues", ReadComplex[BACnetPropertyValues](BACnetPropertyValuesParseWithBufferProducer((uint8)(uint8(4)), (BACnetObjectType)(monitoredObjectIdentifier.GetObjectType())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	m.ListOfValues = listOfValues

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedCOVNotification")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedCOVNotification")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", m.GetLifetimeInSeconds(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetimeInSeconds' field")
		}

		if err := WriteSimpleField[BACnetPropertyValues](ctx, "listOfValues", m.GetListOfValues(), WriteComplex[BACnetPropertyValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedCOVNotification")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) IsBACnetConfirmedServiceRequestConfirmedCOVNotification() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) deepCopy() *_BACnetConfirmedServiceRequestConfirmedCOVNotification {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedCOVNotificationCopy := &_BACnetConfirmedServiceRequestConfirmedCOVNotification{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.SubscriberProcessIdentifier.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.InitiatingDeviceIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.MonitoredObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.LifetimeInSeconds.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ListOfValues.DeepCopy().(BACnetPropertyValues),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedCOVNotificationCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
