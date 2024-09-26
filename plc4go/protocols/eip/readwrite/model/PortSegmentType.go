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
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// PortSegmentType is the corresponding interface of PortSegmentType
type PortSegmentType interface {
	PortSegmentTypeContract
	PortSegmentTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsPortSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegmentType()
}

// PortSegmentTypeContract provides a set of functions which can be overwritten by a sub struct
type PortSegmentTypeContract interface {
	// IsPortSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegmentType()
}

// PortSegmentTypeRequirements provides a set of functions which need to be implemented by a sub struct
type PortSegmentTypeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetExtendedLinkAddress returns ExtendedLinkAddress (discriminator field)
	GetExtendedLinkAddress() bool
}

// _PortSegmentType is the data-structure of this message
type _PortSegmentType struct {
	_SubType PortSegmentType
}

var _ PortSegmentTypeContract = (*_PortSegmentType)(nil)

// NewPortSegmentType factory function for _PortSegmentType
func NewPortSegmentType() *_PortSegmentType {
	return &_PortSegmentType{}
}

// Deprecated: use the interface for direct cast
func CastPortSegmentType(structType any) PortSegmentType {
	if casted, ok := structType.(PortSegmentType); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegmentType); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegmentType) GetTypeName() string {
	return "PortSegmentType"
}

func (m *_PortSegmentType) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (extendedLinkAddress)
	lengthInBits += 1

	return lengthInBits
}

func (m *_PortSegmentType) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func PortSegmentTypeParse[T PortSegmentType](ctx context.Context, theBytes []byte) (T, error) {
	return PortSegmentTypeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func PortSegmentTypeParseWithBufferProducer[T PortSegmentType]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := PortSegmentTypeParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func PortSegmentTypeParseWithBuffer[T PortSegmentType](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_PortSegmentType{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_PortSegmentType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__portSegmentType PortSegmentType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegmentType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	extendedLinkAddress, err := ReadDiscriminatorField[bool](ctx, "extendedLinkAddress", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedLinkAddress' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child PortSegmentType
	switch {
	case extendedLinkAddress == bool(false): // PortSegmentNormal
		if _child, err = new(_PortSegmentNormal).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type PortSegmentNormal for type-switch of PortSegmentType")
		}
	case extendedLinkAddress == bool(true): // PortSegmentExtended
		if _child, err = new(_PortSegmentExtended).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type PortSegmentExtended for type-switch of PortSegmentType")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [extendedLinkAddress=%v]", extendedLinkAddress)
	}

	if closeErr := readBuffer.CloseContext("PortSegmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegmentType")
	}

	return _child, nil
}

func (pm *_PortSegmentType) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child PortSegmentType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PortSegmentType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PortSegmentType")
	}

	if err := WriteDiscriminatorField(ctx, "extendedLinkAddress", m.GetExtendedLinkAddress(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'extendedLinkAddress' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("PortSegmentType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PortSegmentType")
	}
	return nil
}

func (m *_PortSegmentType) IsPortSegmentType() {}
