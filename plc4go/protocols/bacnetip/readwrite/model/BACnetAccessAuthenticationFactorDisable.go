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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAccessAuthenticationFactorDisable is an enum
type BACnetAccessAuthenticationFactorDisable uint16

type IBACnetAccessAuthenticationFactorDisable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAccessAuthenticationFactorDisable_NONE                     BACnetAccessAuthenticationFactorDisable = 0
	BACnetAccessAuthenticationFactorDisable_DISABLED                 BACnetAccessAuthenticationFactorDisable = 1
	BACnetAccessAuthenticationFactorDisable_DISABLED_LOST            BACnetAccessAuthenticationFactorDisable = 2
	BACnetAccessAuthenticationFactorDisable_DISABLED_STOLEN          BACnetAccessAuthenticationFactorDisable = 3
	BACnetAccessAuthenticationFactorDisable_DISABLED_DAMAGED         BACnetAccessAuthenticationFactorDisable = 4
	BACnetAccessAuthenticationFactorDisable_DISABLED_DESTROYED       BACnetAccessAuthenticationFactorDisable = 5
	BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE BACnetAccessAuthenticationFactorDisable = 0xFFFF
)

var BACnetAccessAuthenticationFactorDisableValues []BACnetAccessAuthenticationFactorDisable

func init() {
	_ = errors.New
	BACnetAccessAuthenticationFactorDisableValues = []BACnetAccessAuthenticationFactorDisable{
		BACnetAccessAuthenticationFactorDisable_NONE,
		BACnetAccessAuthenticationFactorDisable_DISABLED,
		BACnetAccessAuthenticationFactorDisable_DISABLED_LOST,
		BACnetAccessAuthenticationFactorDisable_DISABLED_STOLEN,
		BACnetAccessAuthenticationFactorDisable_DISABLED_DAMAGED,
		BACnetAccessAuthenticationFactorDisable_DISABLED_DESTROYED,
		BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessAuthenticationFactorDisableByValue(value uint16) (enum BACnetAccessAuthenticationFactorDisable, ok bool) {
	switch value {
	case 0:
		return BACnetAccessAuthenticationFactorDisable_NONE, true
	case 0xFFFF:
		return BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAccessAuthenticationFactorDisable_DISABLED, true
	case 2:
		return BACnetAccessAuthenticationFactorDisable_DISABLED_LOST, true
	case 3:
		return BACnetAccessAuthenticationFactorDisable_DISABLED_STOLEN, true
	case 4:
		return BACnetAccessAuthenticationFactorDisable_DISABLED_DAMAGED, true
	case 5:
		return BACnetAccessAuthenticationFactorDisable_DISABLED_DESTROYED, true
	}
	return 0, false
}

func BACnetAccessAuthenticationFactorDisableByName(value string) (enum BACnetAccessAuthenticationFactorDisable, ok bool) {
	switch value {
	case "NONE":
		return BACnetAccessAuthenticationFactorDisable_NONE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE, true
	case "DISABLED":
		return BACnetAccessAuthenticationFactorDisable_DISABLED, true
	case "DISABLED_LOST":
		return BACnetAccessAuthenticationFactorDisable_DISABLED_LOST, true
	case "DISABLED_STOLEN":
		return BACnetAccessAuthenticationFactorDisable_DISABLED_STOLEN, true
	case "DISABLED_DAMAGED":
		return BACnetAccessAuthenticationFactorDisable_DISABLED_DAMAGED, true
	case "DISABLED_DESTROYED":
		return BACnetAccessAuthenticationFactorDisable_DISABLED_DESTROYED, true
	}
	return 0, false
}

func BACnetAccessAuthenticationFactorDisableKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessAuthenticationFactorDisableValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessAuthenticationFactorDisable(structType any) BACnetAccessAuthenticationFactorDisable {
	castFunc := func(typ any) BACnetAccessAuthenticationFactorDisable {
		if sBACnetAccessAuthenticationFactorDisable, ok := typ.(BACnetAccessAuthenticationFactorDisable); ok {
			return sBACnetAccessAuthenticationFactorDisable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessAuthenticationFactorDisable) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetAccessAuthenticationFactorDisable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessAuthenticationFactorDisableParse(ctx context.Context, theBytes []byte) (BACnetAccessAuthenticationFactorDisable, error) {
	return BACnetAccessAuthenticationFactorDisableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessAuthenticationFactorDisableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessAuthenticationFactorDisable, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("BACnetAccessAuthenticationFactorDisable", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessAuthenticationFactorDisable")
	}
	if enum, ok := BACnetAccessAuthenticationFactorDisableByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAccessAuthenticationFactorDisable")
		return BACnetAccessAuthenticationFactorDisable(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessAuthenticationFactorDisable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessAuthenticationFactorDisable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("BACnetAccessAuthenticationFactorDisable", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetAccessAuthenticationFactorDisable) GetValue() uint16 {
	return uint16(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessAuthenticationFactorDisable) PLC4XEnumName() string {
	switch e {
	case BACnetAccessAuthenticationFactorDisable_NONE:
		return "NONE"
	case BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessAuthenticationFactorDisable_DISABLED:
		return "DISABLED"
	case BACnetAccessAuthenticationFactorDisable_DISABLED_LOST:
		return "DISABLED_LOST"
	case BACnetAccessAuthenticationFactorDisable_DISABLED_STOLEN:
		return "DISABLED_STOLEN"
	case BACnetAccessAuthenticationFactorDisable_DISABLED_DAMAGED:
		return "DISABLED_DAMAGED"
	case BACnetAccessAuthenticationFactorDisable_DISABLED_DESTROYED:
		return "DISABLED_DESTROYED"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetAccessAuthenticationFactorDisable) String() string {
	return e.PLC4XEnumName()
}
