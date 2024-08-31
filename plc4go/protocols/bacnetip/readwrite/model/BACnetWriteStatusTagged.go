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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetWriteStatusTagged is the corresponding interface of BACnetWriteStatusTagged
type BACnetWriteStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetWriteStatus
}

// BACnetWriteStatusTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetWriteStatusTagged.
// This is useful for switch cases.
type BACnetWriteStatusTaggedExactly interface {
	BACnetWriteStatusTagged
	isBACnetWriteStatusTagged() bool
}

// _BACnetWriteStatusTagged is the data-structure of this message
type _BACnetWriteStatusTagged struct {
	Header BACnetTagHeader
	Value  BACnetWriteStatus

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetWriteStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetWriteStatusTagged) GetValue() BACnetWriteStatus {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetWriteStatusTagged factory function for _BACnetWriteStatusTagged
func NewBACnetWriteStatusTagged(header BACnetTagHeader, value BACnetWriteStatus, tagNumber uint8, tagClass TagClass) *_BACnetWriteStatusTagged {
	return &_BACnetWriteStatusTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetWriteStatusTagged(structType any) BACnetWriteStatusTagged {
	if casted, ok := structType.(BACnetWriteStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetWriteStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetWriteStatusTagged) GetTypeName() string {
	return "BACnetWriteStatusTagged"
}

func (m *_BACnetWriteStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetWriteStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetWriteStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetWriteStatusTagged, error) {
	return BACnetWriteStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetWriteStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetWriteStatusTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetWriteStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetWriteStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetWriteStatusTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetWriteStatus](ctx, "value", readBuffer, func(ctx context.Context) (BACnetWriteStatus, error) {
		v, err := ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetWriteStatus_IDLE)
		var zero BACnetWriteStatus
		if err != nil {
			return zero, err
		}
		return v.(BACnetWriteStatus), err
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetWriteStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetWriteStatusTagged")
	}

	// Create the instance
	return &_BACnetWriteStatusTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetWriteStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetWriteStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetWriteStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetWriteStatusTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetWriteStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetWriteStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetWriteStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetWriteStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetWriteStatusTagged) isBACnetWriteStatusTagged() bool {
	return true
}

func (m *_BACnetWriteStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
