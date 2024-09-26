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

// BACnetLoggingType is an enum
type BACnetLoggingType uint8

type IBACnetLoggingType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLoggingType_POLLED                   BACnetLoggingType = 0
	BACnetLoggingType_COV                      BACnetLoggingType = 1
	BACnetLoggingType_TRIGGERED                BACnetLoggingType = 2
	BACnetLoggingType_VENDOR_PROPRIETARY_VALUE BACnetLoggingType = 0xFF
)

var BACnetLoggingTypeValues []BACnetLoggingType

func init() {
	_ = errors.New
	BACnetLoggingTypeValues = []BACnetLoggingType{
		BACnetLoggingType_POLLED,
		BACnetLoggingType_COV,
		BACnetLoggingType_TRIGGERED,
		BACnetLoggingType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLoggingTypeByValue(value uint8) (enum BACnetLoggingType, ok bool) {
	switch value {
	case 0:
		return BACnetLoggingType_POLLED, true
	case 0xFF:
		return BACnetLoggingType_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLoggingType_COV, true
	case 2:
		return BACnetLoggingType_TRIGGERED, true
	}
	return 0, false
}

func BACnetLoggingTypeByName(value string) (enum BACnetLoggingType, ok bool) {
	switch value {
	case "POLLED":
		return BACnetLoggingType_POLLED, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLoggingType_VENDOR_PROPRIETARY_VALUE, true
	case "COV":
		return BACnetLoggingType_COV, true
	case "TRIGGERED":
		return BACnetLoggingType_TRIGGERED, true
	}
	return 0, false
}

func BACnetLoggingTypeKnows(value uint8) bool {
	for _, typeValue := range BACnetLoggingTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLoggingType(structType any) BACnetLoggingType {
	castFunc := func(typ any) BACnetLoggingType {
		if sBACnetLoggingType, ok := typ.(BACnetLoggingType); ok {
			return sBACnetLoggingType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLoggingType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetLoggingType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLoggingTypeParse(ctx context.Context, theBytes []byte) (BACnetLoggingType, error) {
	return BACnetLoggingTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLoggingTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLoggingType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetLoggingType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLoggingType")
	}
	if enum, ok := BACnetLoggingTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLoggingType")
		return BACnetLoggingType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLoggingType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLoggingType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetLoggingType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetLoggingType) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLoggingType) PLC4XEnumName() string {
	switch e {
	case BACnetLoggingType_POLLED:
		return "POLLED"
	case BACnetLoggingType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLoggingType_COV:
		return "COV"
	case BACnetLoggingType_TRIGGERED:
		return "TRIGGERED"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetLoggingType) String() string {
	return e.PLC4XEnumName()
}
