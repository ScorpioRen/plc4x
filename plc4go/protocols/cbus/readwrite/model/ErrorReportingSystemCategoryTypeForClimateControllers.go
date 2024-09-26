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

// ErrorReportingSystemCategoryTypeForClimateControllers is an enum
type ErrorReportingSystemCategoryTypeForClimateControllers uint8

type IErrorReportingSystemCategoryTypeForClimateControllers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ErrorReportingSystemCategoryTypeForClimateControllers_AIR_CONDITIONING_SYSTEM  ErrorReportingSystemCategoryTypeForClimateControllers = 0x0
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_1               ErrorReportingSystemCategoryTypeForClimateControllers = 0x1
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_2               ErrorReportingSystemCategoryTypeForClimateControllers = 0x2
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_3               ErrorReportingSystemCategoryTypeForClimateControllers = 0x3
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_4               ErrorReportingSystemCategoryTypeForClimateControllers = 0x4
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_5               ErrorReportingSystemCategoryTypeForClimateControllers = 0x5
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_6               ErrorReportingSystemCategoryTypeForClimateControllers = 0x6
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_7               ErrorReportingSystemCategoryTypeForClimateControllers = 0x7
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_8               ErrorReportingSystemCategoryTypeForClimateControllers = 0x8
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_9               ErrorReportingSystemCategoryTypeForClimateControllers = 0x9
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_10              ErrorReportingSystemCategoryTypeForClimateControllers = 0xA
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_11              ErrorReportingSystemCategoryTypeForClimateControllers = 0xB
	ErrorReportingSystemCategoryTypeForClimateControllers_GLOBAL_WARMING_MODULATOR ErrorReportingSystemCategoryTypeForClimateControllers = 0xC
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_13              ErrorReportingSystemCategoryTypeForClimateControllers = 0xD
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_14              ErrorReportingSystemCategoryTypeForClimateControllers = 0xE
	ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_15              ErrorReportingSystemCategoryTypeForClimateControllers = 0xF
)

var ErrorReportingSystemCategoryTypeForClimateControllersValues []ErrorReportingSystemCategoryTypeForClimateControllers

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryTypeForClimateControllersValues = []ErrorReportingSystemCategoryTypeForClimateControllers{
		ErrorReportingSystemCategoryTypeForClimateControllers_AIR_CONDITIONING_SYSTEM,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_1,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_2,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_3,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_4,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_5,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_6,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_7,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_8,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_9,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_10,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_11,
		ErrorReportingSystemCategoryTypeForClimateControllers_GLOBAL_WARMING_MODULATOR,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_13,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_14,
		ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_15,
	}
}

func ErrorReportingSystemCategoryTypeForClimateControllersByValue(value uint8) (enum ErrorReportingSystemCategoryTypeForClimateControllers, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryTypeForClimateControllers_AIR_CONDITIONING_SYSTEM, true
	case 0x1:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_1, true
	case 0x2:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_3, true
	case 0x4:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_4, true
	case 0x5:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_5, true
	case 0x6:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_6, true
	case 0x7:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_7, true
	case 0x8:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_8, true
	case 0x9:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_9, true
	case 0xA:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_10, true
	case 0xB:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_11, true
	case 0xC:
		return ErrorReportingSystemCategoryTypeForClimateControllers_GLOBAL_WARMING_MODULATOR, true
	case 0xD:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_13, true
	case 0xE:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_14, true
	case 0xF:
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForClimateControllersByName(value string) (enum ErrorReportingSystemCategoryTypeForClimateControllers, ok bool) {
	switch value {
	case "AIR_CONDITIONING_SYSTEM":
		return ErrorReportingSystemCategoryTypeForClimateControllers_AIR_CONDITIONING_SYSTEM, true
	case "RESERVED_1":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_3, true
	case "RESERVED_4":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_4, true
	case "RESERVED_5":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_5, true
	case "RESERVED_6":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_6, true
	case "RESERVED_7":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_7, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_8, true
	case "RESERVED_9":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_9, true
	case "RESERVED_10":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_10, true
	case "RESERVED_11":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_11, true
	case "GLOBAL_WARMING_MODULATOR":
		return ErrorReportingSystemCategoryTypeForClimateControllers_GLOBAL_WARMING_MODULATOR, true
	case "RESERVED_13":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_13, true
	case "RESERVED_14":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_14, true
	case "RESERVED_15":
		return ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForClimateControllersKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryTypeForClimateControllersValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryTypeForClimateControllers(structType any) ErrorReportingSystemCategoryTypeForClimateControllers {
	castFunc := func(typ any) ErrorReportingSystemCategoryTypeForClimateControllers {
		if sErrorReportingSystemCategoryTypeForClimateControllers, ok := typ.(ErrorReportingSystemCategoryTypeForClimateControllers); ok {
			return sErrorReportingSystemCategoryTypeForClimateControllers
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryTypeForClimateControllers) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryTypeForClimateControllers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeForClimateControllersParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryTypeForClimateControllers, error) {
	return ErrorReportingSystemCategoryTypeForClimateControllersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryTypeForClimateControllersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryTypeForClimateControllers, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("ErrorReportingSystemCategoryTypeForClimateControllers", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryTypeForClimateControllers")
	}
	if enum, ok := ErrorReportingSystemCategoryTypeForClimateControllersByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ErrorReportingSystemCategoryTypeForClimateControllers")
		return ErrorReportingSystemCategoryTypeForClimateControllers(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryTypeForClimateControllers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryTypeForClimateControllers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("ErrorReportingSystemCategoryTypeForClimateControllers", 4, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e ErrorReportingSystemCategoryTypeForClimateControllers) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryTypeForClimateControllers) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryTypeForClimateControllers_AIR_CONDITIONING_SYSTEM:
		return "AIR_CONDITIONING_SYSTEM"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_4:
		return "RESERVED_4"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_5:
		return "RESERVED_5"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_6:
		return "RESERVED_6"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_7:
		return "RESERVED_7"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_9:
		return "RESERVED_9"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_10:
		return "RESERVED_10"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_11:
		return "RESERVED_11"
	case ErrorReportingSystemCategoryTypeForClimateControllers_GLOBAL_WARMING_MODULATOR:
		return "GLOBAL_WARMING_MODULATOR"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_13:
		return "RESERVED_13"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_14:
		return "RESERVED_14"
	case ErrorReportingSystemCategoryTypeForClimateControllers_RESERVED_15:
		return "RESERVED_15"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ErrorReportingSystemCategoryTypeForClimateControllers) String() string {
	return e.PLC4XEnumName()
}
