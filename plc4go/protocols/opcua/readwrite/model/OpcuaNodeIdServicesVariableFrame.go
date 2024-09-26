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

// OpcuaNodeIdServicesVariableFrame is an enum
type OpcuaNodeIdServicesVariableFrame int32

type IOpcuaNodeIdServicesVariableFrame interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableFrame_FrameType_Orientation                     OpcuaNodeIdServicesVariableFrame = 18787
	OpcuaNodeIdServicesVariableFrame_FrameType_Constant                        OpcuaNodeIdServicesVariableFrame = 18788
	OpcuaNodeIdServicesVariableFrame_FrameType_BaseFrame                       OpcuaNodeIdServicesVariableFrame = 18789
	OpcuaNodeIdServicesVariableFrame_FrameType_FixedBase                       OpcuaNodeIdServicesVariableFrame = 18790
	OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates            OpcuaNodeIdServicesVariableFrame = 18801
	OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates_LengthUnit OpcuaNodeIdServicesVariableFrame = 18802
	OpcuaNodeIdServicesVariableFrame_FrameType_Orientation_AngleUnit           OpcuaNodeIdServicesVariableFrame = 18803
)

var OpcuaNodeIdServicesVariableFrameValues []OpcuaNodeIdServicesVariableFrame

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableFrameValues = []OpcuaNodeIdServicesVariableFrame{
		OpcuaNodeIdServicesVariableFrame_FrameType_Orientation,
		OpcuaNodeIdServicesVariableFrame_FrameType_Constant,
		OpcuaNodeIdServicesVariableFrame_FrameType_BaseFrame,
		OpcuaNodeIdServicesVariableFrame_FrameType_FixedBase,
		OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates,
		OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates_LengthUnit,
		OpcuaNodeIdServicesVariableFrame_FrameType_Orientation_AngleUnit,
	}
}

func OpcuaNodeIdServicesVariableFrameByValue(value int32) (enum OpcuaNodeIdServicesVariableFrame, ok bool) {
	switch value {
	case 18787:
		return OpcuaNodeIdServicesVariableFrame_FrameType_Orientation, true
	case 18788:
		return OpcuaNodeIdServicesVariableFrame_FrameType_Constant, true
	case 18789:
		return OpcuaNodeIdServicesVariableFrame_FrameType_BaseFrame, true
	case 18790:
		return OpcuaNodeIdServicesVariableFrame_FrameType_FixedBase, true
	case 18801:
		return OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates, true
	case 18802:
		return OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates_LengthUnit, true
	case 18803:
		return OpcuaNodeIdServicesVariableFrame_FrameType_Orientation_AngleUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFrameByName(value string) (enum OpcuaNodeIdServicesVariableFrame, ok bool) {
	switch value {
	case "FrameType_Orientation":
		return OpcuaNodeIdServicesVariableFrame_FrameType_Orientation, true
	case "FrameType_Constant":
		return OpcuaNodeIdServicesVariableFrame_FrameType_Constant, true
	case "FrameType_BaseFrame":
		return OpcuaNodeIdServicesVariableFrame_FrameType_BaseFrame, true
	case "FrameType_FixedBase":
		return OpcuaNodeIdServicesVariableFrame_FrameType_FixedBase, true
	case "FrameType_CartesianCoordinates":
		return OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates, true
	case "FrameType_CartesianCoordinates_LengthUnit":
		return OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates_LengthUnit, true
	case "FrameType_Orientation_AngleUnit":
		return OpcuaNodeIdServicesVariableFrame_FrameType_Orientation_AngleUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFrameKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableFrameValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableFrame(structType any) OpcuaNodeIdServicesVariableFrame {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableFrame {
		if sOpcuaNodeIdServicesVariableFrame, ok := typ.(OpcuaNodeIdServicesVariableFrame); ok {
			return sOpcuaNodeIdServicesVariableFrame
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableFrame) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableFrameParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableFrame, error) {
	return OpcuaNodeIdServicesVariableFrameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableFrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableFrame, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableFrame", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableFrame")
	}
	if enum, ok := OpcuaNodeIdServicesVariableFrameByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableFrame")
		return OpcuaNodeIdServicesVariableFrame(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableFrame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableFrame", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableFrame) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableFrame) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableFrame_FrameType_Orientation:
		return "FrameType_Orientation"
	case OpcuaNodeIdServicesVariableFrame_FrameType_Constant:
		return "FrameType_Constant"
	case OpcuaNodeIdServicesVariableFrame_FrameType_BaseFrame:
		return "FrameType_BaseFrame"
	case OpcuaNodeIdServicesVariableFrame_FrameType_FixedBase:
		return "FrameType_FixedBase"
	case OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates:
		return "FrameType_CartesianCoordinates"
	case OpcuaNodeIdServicesVariableFrame_FrameType_CartesianCoordinates_LengthUnit:
		return "FrameType_CartesianCoordinates_LengthUnit"
	case OpcuaNodeIdServicesVariableFrame_FrameType_Orientation_AngleUnit:
		return "FrameType_Orientation_AngleUnit"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableFrame) String() string {
	return e.PLC4XEnumName()
}
