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

// NodeIdType is an enum
type NodeIdType uint8

type INodeIdType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	NodeIdType_nodeIdTypeTwoByte    NodeIdType = 0
	NodeIdType_nodeIdTypeFourByte   NodeIdType = 1
	NodeIdType_nodeIdTypeNumeric    NodeIdType = 2
	NodeIdType_nodeIdTypeString     NodeIdType = 3
	NodeIdType_nodeIdTypeGuid       NodeIdType = 4
	NodeIdType_nodeIdTypeByteString NodeIdType = 5
)

var NodeIdTypeValues []NodeIdType

func init() {
	_ = errors.New
	NodeIdTypeValues = []NodeIdType{
		NodeIdType_nodeIdTypeTwoByte,
		NodeIdType_nodeIdTypeFourByte,
		NodeIdType_nodeIdTypeNumeric,
		NodeIdType_nodeIdTypeString,
		NodeIdType_nodeIdTypeGuid,
		NodeIdType_nodeIdTypeByteString,
	}
}

func NodeIdTypeByValue(value uint8) (enum NodeIdType, ok bool) {
	switch value {
	case 0:
		return NodeIdType_nodeIdTypeTwoByte, true
	case 1:
		return NodeIdType_nodeIdTypeFourByte, true
	case 2:
		return NodeIdType_nodeIdTypeNumeric, true
	case 3:
		return NodeIdType_nodeIdTypeString, true
	case 4:
		return NodeIdType_nodeIdTypeGuid, true
	case 5:
		return NodeIdType_nodeIdTypeByteString, true
	}
	return 0, false
}

func NodeIdTypeByName(value string) (enum NodeIdType, ok bool) {
	switch value {
	case "nodeIdTypeTwoByte":
		return NodeIdType_nodeIdTypeTwoByte, true
	case "nodeIdTypeFourByte":
		return NodeIdType_nodeIdTypeFourByte, true
	case "nodeIdTypeNumeric":
		return NodeIdType_nodeIdTypeNumeric, true
	case "nodeIdTypeString":
		return NodeIdType_nodeIdTypeString, true
	case "nodeIdTypeGuid":
		return NodeIdType_nodeIdTypeGuid, true
	case "nodeIdTypeByteString":
		return NodeIdType_nodeIdTypeByteString, true
	}
	return 0, false
}

func NodeIdTypeKnows(value uint8) bool {
	for _, typeValue := range NodeIdTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNodeIdType(structType any) NodeIdType {
	castFunc := func(typ any) NodeIdType {
		if sNodeIdType, ok := typ.(NodeIdType); ok {
			return sNodeIdType
		}
		return 0
	}
	return castFunc(structType)
}

func (m NodeIdType) GetLengthInBits(ctx context.Context) uint16 {
	return 6
}

func (m NodeIdType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdTypeParse(ctx context.Context, theBytes []byte) (NodeIdType, error) {
	return NodeIdTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("NodeIdType", 6)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NodeIdType")
	}
	if enum, ok := NodeIdTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for NodeIdType")
		return NodeIdType(val), nil
	} else {
		return enum, nil
	}
}

func (e NodeIdType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NodeIdType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("NodeIdType", 6, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e NodeIdType) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NodeIdType) PLC4XEnumName() string {
	switch e {
	case NodeIdType_nodeIdTypeTwoByte:
		return "nodeIdTypeTwoByte"
	case NodeIdType_nodeIdTypeFourByte:
		return "nodeIdTypeFourByte"
	case NodeIdType_nodeIdTypeNumeric:
		return "nodeIdTypeNumeric"
	case NodeIdType_nodeIdTypeString:
		return "nodeIdTypeString"
	case NodeIdType_nodeIdTypeGuid:
		return "nodeIdTypeGuid"
	case NodeIdType_nodeIdTypeByteString:
		return "nodeIdTypeByteString"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e NodeIdType) String() string {
	return e.PLC4XEnumName()
}
