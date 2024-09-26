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

// BACnetUnconfirmedServiceRequestUnconfirmedTextMessage is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
type BACnetUnconfirmedServiceRequestUnconfirmedTextMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetTextMessageSourceDevice returns TextMessageSourceDevice (property field)
	GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier
	// GetMessageClass returns MessageClass (property field)
	GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetMessagePriority returns MessagePriority (property field)
	GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	// GetMessage returns Message (property field)
	GetMessage() BACnetContextTagCharacterString
	// IsBACnetUnconfirmedServiceRequestUnconfirmedTextMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestUnconfirmedTextMessage()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	CreateBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder() BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
}

// _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage struct {
	BACnetUnconfirmedServiceRequestContract
	TextMessageSourceDevice BACnetContextTagObjectIdentifier
	MessageClass            BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	MessagePriority         BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	Message                 BACnetContextTagCharacterString
}

var _ BACnetUnconfirmedServiceRequestUnconfirmedTextMessage = (*_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage)(nil)

// NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessage factory function for _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
func NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessage(textMessageSourceDevice BACnetContextTagObjectIdentifier, messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	if textMessageSourceDevice == nil {
		panic("textMessageSourceDevice of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage must not be nil")
	}
	if messagePriority == nil {
		panic("messagePriority of type BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage must not be nil")
	}
	if message == nil {
		panic("message of type BACnetContextTagCharacterString for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		TextMessageSourceDevice:                 textMessageSourceDevice,
		MessageClass:                            messageClass,
		MessagePriority:                         messagePriority,
		Message:                                 message,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder is a builder for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
type BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(textMessageSourceDevice BACnetContextTagObjectIdentifier, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithTextMessageSourceDevice adds TextMessageSourceDevice (property field)
	WithTextMessageSourceDevice(BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithTextMessageSourceDeviceBuilder adds TextMessageSourceDevice (property field) which is build by the builder
	WithTextMessageSourceDeviceBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithMessageClass adds MessageClass (property field)
	WithOptionalMessageClass(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithOptionalMessageClassBuilder adds MessageClass (property field) which is build by the builder
	WithOptionalMessageClassBuilder(func(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithMessagePriority adds MessagePriority (property field)
	WithMessagePriority(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithMessagePriorityBuilder adds MessagePriority (property field) which is build by the builder
	WithMessagePriorityBuilder(func(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithMessage adds Message (property field)
	WithMessage(BACnetContextTagCharacterString) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
	// Build builds the BACnetUnconfirmedServiceRequestUnconfirmedTextMessage or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestUnconfirmedTextMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
}

// NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder() creates a BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
func NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder() BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	return &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder{_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage: new(_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage)}
}

type _BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder struct {
	*_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) = (*_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder)(nil)

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithMandatoryFields(textMessageSourceDevice BACnetContextTagObjectIdentifier, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	return m.WithTextMessageSourceDevice(textMessageSourceDevice).WithMessagePriority(messagePriority).WithMessage(message)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithTextMessageSourceDevice(textMessageSourceDevice BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	m.TextMessageSourceDevice = textMessageSourceDevice
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithTextMessageSourceDeviceBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	builder := builderSupplier(m.TextMessageSourceDevice.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.TextMessageSourceDevice, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithOptionalMessageClass(messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	m.MessageClass = messageClass
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithOptionalMessageClassBuilder(builderSupplier func(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	builder := builderSupplier(m.MessageClass.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder())
	var err error
	m.MessageClass, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder failed"))
	}
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithMessagePriority(messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	m.MessagePriority = messagePriority
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithMessagePriorityBuilder(builderSupplier func(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	builder := builderSupplier(m.MessagePriority.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder())
	var err error
	m.MessagePriority, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithMessage(message BACnetContextTagCharacterString) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	m.Message = message
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) WithMessageBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	builder := builderSupplier(m.Message.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	m.Message, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return m
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) Build() (BACnetUnconfirmedServiceRequestUnconfirmedTextMessage, error) {
	if m.TextMessageSourceDevice == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'textMessageSourceDevice' not set"))
	}
	if m.MessagePriority == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'messagePriority' not set"))
	}
	if m.Message == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetUnconfirmedServiceRequestUnconfirmedTextMessage.deepCopy(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) MustBuild() BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder) DeepCopy() any {
	return m.CreateBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder()
}

// CreateBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder creates a BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder
func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) CreateBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder() BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder {
	if m == nil {
		return NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageBuilder{_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier {
	return m.TextMessageSourceDevice
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m.MessageClass
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	return m.MessagePriority
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessage() BACnetContextTagCharacterString {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedTextMessage(structType any) BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedTextMessage); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedTextMessage); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (textMessageSourceDevice)
	lengthInBits += m.TextMessageSourceDevice.GetLengthInBits(ctx)

	// Optional Field (messageClass)
	if m.MessageClass != nil {
		lengthInBits += m.MessageClass.GetLengthInBits(ctx)
	}

	// Simple field (messagePriority)
	lengthInBits += m.MessagePriority.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUnconfirmedTextMessage BACnetUnconfirmedServiceRequestUnconfirmedTextMessage, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	textMessageSourceDevice, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "textMessageSourceDevice", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'textMessageSourceDevice' field"))
	}
	m.TextMessageSourceDevice = textMessageSourceDevice

	var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	_messageClass, err := ReadOptionalField[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx, "messageClass", ReadComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParseWithBufferProducer[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass]((uint8)(uint8(1))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageClass' field"))
	}
	if _messageClass != nil {
		messageClass = *_messageClass
		m.MessageClass = messageClass
	}

	messagePriority, err := ReadSimpleField[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](ctx, "messagePriority", ReadComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messagePriority' field"))
	}
	m.MessagePriority = messagePriority

	message, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "message", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "textMessageSourceDevice", m.GetTextMessageSourceDevice(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'textMessageSourceDevice' field")
		}

		if err := WriteOptionalField[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](ctx, "messageClass", GetRef(m.GetMessageClass()), WriteComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'messageClass' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](ctx, "messagePriority", m.GetMessagePriority(), WriteComplex[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'messagePriority' field")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "message", m.GetMessage(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) IsBACnetUnconfirmedServiceRequestUnconfirmedTextMessage() {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) deepCopy() *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestUnconfirmedTextMessageCopy := &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		m.TextMessageSourceDevice.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.MessageClass.DeepCopy().(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass),
		m.MessagePriority.DeepCopy().(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged),
		m.Message.DeepCopy().(BACnetContextTagCharacterString),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestUnconfirmedTextMessageCopy
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
