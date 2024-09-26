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

// S7VarRequestParameterItem is the corresponding interface of S7VarRequestParameterItem
type S7VarRequestParameterItem interface {
	S7VarRequestParameterItemContract
	S7VarRequestParameterItemRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsS7VarRequestParameterItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7VarRequestParameterItem()
}

// S7VarRequestParameterItemContract provides a set of functions which can be overwritten by a sub struct
type S7VarRequestParameterItemContract interface {
	// IsS7VarRequestParameterItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7VarRequestParameterItem()
}

// S7VarRequestParameterItemRequirements provides a set of functions which need to be implemented by a sub struct
type S7VarRequestParameterItemRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
}

// _S7VarRequestParameterItem is the data-structure of this message
type _S7VarRequestParameterItem struct {
	_SubType S7VarRequestParameterItem
}

var _ S7VarRequestParameterItemContract = (*_S7VarRequestParameterItem)(nil)

// NewS7VarRequestParameterItem factory function for _S7VarRequestParameterItem
func NewS7VarRequestParameterItem() *_S7VarRequestParameterItem {
	return &_S7VarRequestParameterItem{}
}

// Deprecated: use the interface for direct cast
func CastS7VarRequestParameterItem(structType any) S7VarRequestParameterItem {
	if casted, ok := structType.(S7VarRequestParameterItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarRequestParameterItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarRequestParameterItem) GetTypeName() string {
	return "S7VarRequestParameterItem"
}

func (m *_S7VarRequestParameterItem) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7VarRequestParameterItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7VarRequestParameterItemParse[T S7VarRequestParameterItem](ctx context.Context, theBytes []byte) (T, error) {
	return S7VarRequestParameterItemParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7VarRequestParameterItemParseWithBufferProducer[T S7VarRequestParameterItem]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7VarRequestParameterItemParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func S7VarRequestParameterItemParseWithBuffer[T S7VarRequestParameterItem](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_S7VarRequestParameterItem{}).parse(ctx, readBuffer)
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

func (m *_S7VarRequestParameterItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__s7VarRequestParameterItem S7VarRequestParameterItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarRequestParameterItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemType, err := ReadDiscriminatorField[uint8](ctx, "itemType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7VarRequestParameterItem
	switch {
	case itemType == 0x12: // S7VarRequestParameterItemAddress
		if _child, err = new(_S7VarRequestParameterItemAddress).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7VarRequestParameterItemAddress for type-switch of S7VarRequestParameterItem")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [itemType=%v]", itemType)
	}

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarRequestParameterItem")
	}

	return _child, nil
}

func (pm *_S7VarRequestParameterItem) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7VarRequestParameterItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7VarRequestParameterItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarRequestParameterItem")
	}

	if err := WriteDiscriminatorField(ctx, "itemType", m.GetItemType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7VarRequestParameterItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarRequestParameterItem")
	}
	return nil
}

func (m *_S7VarRequestParameterItem) IsS7VarRequestParameterItem() {}
