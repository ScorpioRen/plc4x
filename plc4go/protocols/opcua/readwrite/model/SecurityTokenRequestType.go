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

// SecurityTokenRequestType is an enum
type SecurityTokenRequestType uint32

type ISecurityTokenRequestType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	SecurityTokenRequestType_securityTokenRequestTypeIssue SecurityTokenRequestType = 0
	SecurityTokenRequestType_securityTokenRequestTypeRenew SecurityTokenRequestType = 1
)

var SecurityTokenRequestTypeValues []SecurityTokenRequestType

func init() {
	_ = errors.New
	SecurityTokenRequestTypeValues = []SecurityTokenRequestType{
		SecurityTokenRequestType_securityTokenRequestTypeIssue,
		SecurityTokenRequestType_securityTokenRequestTypeRenew,
	}
}

func SecurityTokenRequestTypeByValue(value uint32) (enum SecurityTokenRequestType, ok bool) {
	switch value {
	case 0:
		return SecurityTokenRequestType_securityTokenRequestTypeIssue, true
	case 1:
		return SecurityTokenRequestType_securityTokenRequestTypeRenew, true
	}
	return 0, false
}

func SecurityTokenRequestTypeByName(value string) (enum SecurityTokenRequestType, ok bool) {
	switch value {
	case "securityTokenRequestTypeIssue":
		return SecurityTokenRequestType_securityTokenRequestTypeIssue, true
	case "securityTokenRequestTypeRenew":
		return SecurityTokenRequestType_securityTokenRequestTypeRenew, true
	}
	return 0, false
}

func SecurityTokenRequestTypeKnows(value uint32) bool {
	for _, typeValue := range SecurityTokenRequestTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSecurityTokenRequestType(structType any) SecurityTokenRequestType {
	castFunc := func(typ any) SecurityTokenRequestType {
		if sSecurityTokenRequestType, ok := typ.(SecurityTokenRequestType); ok {
			return sSecurityTokenRequestType
		}
		return 0
	}
	return castFunc(structType)
}

func (m SecurityTokenRequestType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m SecurityTokenRequestType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityTokenRequestTypeParse(ctx context.Context, theBytes []byte) (SecurityTokenRequestType, error) {
	return SecurityTokenRequestTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityTokenRequestTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityTokenRequestType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("SecurityTokenRequestType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SecurityTokenRequestType")
	}
	if enum, ok := SecurityTokenRequestTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for SecurityTokenRequestType")
		return SecurityTokenRequestType(val), nil
	} else {
		return enum, nil
	}
}

func (e SecurityTokenRequestType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e SecurityTokenRequestType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("SecurityTokenRequestType", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e SecurityTokenRequestType) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SecurityTokenRequestType) PLC4XEnumName() string {
	switch e {
	case SecurityTokenRequestType_securityTokenRequestTypeIssue:
		return "securityTokenRequestTypeIssue"
	case SecurityTokenRequestType_securityTokenRequestTypeRenew:
		return "securityTokenRequestTypeRenew"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e SecurityTokenRequestType) String() string {
	return e.PLC4XEnumName()
}
