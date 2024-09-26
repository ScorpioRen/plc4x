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

// OpcuaNodeIdServicesVariableRequest is an enum
type OpcuaNodeIdServicesVariableRequest int32

type IOpcuaNodeIdServicesVariableRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableRequest_RequestServerStateChangeMethodType_InputArguments OpcuaNodeIdServicesVariableRequest = 12889
	OpcuaNodeIdServicesVariableRequest_RequestTicketsMethodType_OutputArguments          OpcuaNodeIdServicesVariableRequest = 25728
)

var OpcuaNodeIdServicesVariableRequestValues []OpcuaNodeIdServicesVariableRequest

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableRequestValues = []OpcuaNodeIdServicesVariableRequest{
		OpcuaNodeIdServicesVariableRequest_RequestServerStateChangeMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRequest_RequestTicketsMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableRequestByValue(value int32) (enum OpcuaNodeIdServicesVariableRequest, ok bool) {
	switch value {
	case 12889:
		return OpcuaNodeIdServicesVariableRequest_RequestServerStateChangeMethodType_InputArguments, true
	case 25728:
		return OpcuaNodeIdServicesVariableRequest_RequestTicketsMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRequestByName(value string) (enum OpcuaNodeIdServicesVariableRequest, ok bool) {
	switch value {
	case "RequestServerStateChangeMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRequest_RequestServerStateChangeMethodType_InputArguments, true
	case "RequestTicketsMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableRequest_RequestTicketsMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRequestKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableRequestValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableRequest(structType any) OpcuaNodeIdServicesVariableRequest {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRequest {
		if sOpcuaNodeIdServicesVariableRequest, ok := typ.(OpcuaNodeIdServicesVariableRequest); ok {
			return sOpcuaNodeIdServicesVariableRequest
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRequest) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableRequestParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRequest, error) {
	return OpcuaNodeIdServicesVariableRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRequest, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRequest", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRequest")
	}
	if enum, ok := OpcuaNodeIdServicesVariableRequestByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRequest")
		return OpcuaNodeIdServicesVariableRequest(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRequest", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableRequest) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRequest) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRequest_RequestServerStateChangeMethodType_InputArguments:
		return "RequestServerStateChangeMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRequest_RequestTicketsMethodType_OutputArguments:
		return "RequestTicketsMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRequest) String() string {
	return e.PLC4XEnumName()
}
