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

// OpcuaNodeIdServicesVariablePriority is an enum
type OpcuaNodeIdServicesVariablePriority int32

type IOpcuaNodeIdServicesVariablePriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_PriorityMapppingEntries                   OpcuaNodeIdServicesVariablePriority = 25228
	OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_AddPriorityMappingEntry_InputArguments    OpcuaNodeIdServicesVariablePriority = 25230
	OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments OpcuaNodeIdServicesVariablePriority = 25232
)

var OpcuaNodeIdServicesVariablePriorityValues []OpcuaNodeIdServicesVariablePriority

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariablePriorityValues = []OpcuaNodeIdServicesVariablePriority{
		OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_PriorityMapppingEntries,
		OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_AddPriorityMappingEntry_InputArguments,
		OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments,
	}
}

func OpcuaNodeIdServicesVariablePriorityByValue(value int32) (enum OpcuaNodeIdServicesVariablePriority, ok bool) {
	switch value {
	case 25228:
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_PriorityMapppingEntries, true
	case 25230:
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_AddPriorityMappingEntry_InputArguments, true
	case 25232:
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariablePriorityByName(value string) (enum OpcuaNodeIdServicesVariablePriority, ok bool) {
	switch value {
	case "PriorityMappingTableType_PriorityMapppingEntries":
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_PriorityMapppingEntries, true
	case "PriorityMappingTableType_AddPriorityMappingEntry_InputArguments":
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_AddPriorityMappingEntry_InputArguments, true
	case "PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments":
		return OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariablePriorityKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariablePriorityValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariablePriority(structType any) OpcuaNodeIdServicesVariablePriority {
	castFunc := func(typ any) OpcuaNodeIdServicesVariablePriority {
		if sOpcuaNodeIdServicesVariablePriority, ok := typ.(OpcuaNodeIdServicesVariablePriority); ok {
			return sOpcuaNodeIdServicesVariablePriority
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariablePriority) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariablePriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariablePriorityParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariablePriority, error) {
	return OpcuaNodeIdServicesVariablePriorityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariablePriorityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariablePriority, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariablePriority", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariablePriority")
	}
	if enum, ok := OpcuaNodeIdServicesVariablePriorityByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariablePriority")
		return OpcuaNodeIdServicesVariablePriority(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariablePriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariablePriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariablePriority", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariablePriority) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariablePriority) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_PriorityMapppingEntries:
		return "PriorityMappingTableType_PriorityMapppingEntries"
	case OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_AddPriorityMappingEntry_InputArguments:
		return "PriorityMappingTableType_AddPriorityMappingEntry_InputArguments"
	case OpcuaNodeIdServicesVariablePriority_PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments:
		return "PriorityMappingTableType_DeletePriorityMappingEntry_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariablePriority) String() string {
	return e.PLC4XEnumName()
}
