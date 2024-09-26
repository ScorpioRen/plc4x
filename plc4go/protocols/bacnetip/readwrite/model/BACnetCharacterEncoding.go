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

// BACnetCharacterEncoding is an enum
type BACnetCharacterEncoding byte

type IBACnetCharacterEncoding interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetCharacterEncoding_ISO_10646          BACnetCharacterEncoding = 0x0
	BACnetCharacterEncoding_IBM_Microsoft_DBCS BACnetCharacterEncoding = 0x1
	BACnetCharacterEncoding_JIS_X_0208         BACnetCharacterEncoding = 0x2
	BACnetCharacterEncoding_ISO_10646_4        BACnetCharacterEncoding = 0x3
	BACnetCharacterEncoding_ISO_10646_2        BACnetCharacterEncoding = 0x4
	BACnetCharacterEncoding_ISO_8859_1         BACnetCharacterEncoding = 0x5
)

var BACnetCharacterEncodingValues []BACnetCharacterEncoding

func init() {
	_ = errors.New
	BACnetCharacterEncodingValues = []BACnetCharacterEncoding{
		BACnetCharacterEncoding_ISO_10646,
		BACnetCharacterEncoding_IBM_Microsoft_DBCS,
		BACnetCharacterEncoding_JIS_X_0208,
		BACnetCharacterEncoding_ISO_10646_4,
		BACnetCharacterEncoding_ISO_10646_2,
		BACnetCharacterEncoding_ISO_8859_1,
	}
}

func BACnetCharacterEncodingByValue(value byte) (enum BACnetCharacterEncoding, ok bool) {
	switch value {
	case 0x0:
		return BACnetCharacterEncoding_ISO_10646, true
	case 0x1:
		return BACnetCharacterEncoding_IBM_Microsoft_DBCS, true
	case 0x2:
		return BACnetCharacterEncoding_JIS_X_0208, true
	case 0x3:
		return BACnetCharacterEncoding_ISO_10646_4, true
	case 0x4:
		return BACnetCharacterEncoding_ISO_10646_2, true
	case 0x5:
		return BACnetCharacterEncoding_ISO_8859_1, true
	}
	return 0, false
}

func BACnetCharacterEncodingByName(value string) (enum BACnetCharacterEncoding, ok bool) {
	switch value {
	case "ISO_10646":
		return BACnetCharacterEncoding_ISO_10646, true
	case "IBM_Microsoft_DBCS":
		return BACnetCharacterEncoding_IBM_Microsoft_DBCS, true
	case "JIS_X_0208":
		return BACnetCharacterEncoding_JIS_X_0208, true
	case "ISO_10646_4":
		return BACnetCharacterEncoding_ISO_10646_4, true
	case "ISO_10646_2":
		return BACnetCharacterEncoding_ISO_10646_2, true
	case "ISO_8859_1":
		return BACnetCharacterEncoding_ISO_8859_1, true
	}
	return 0, false
}

func BACnetCharacterEncodingKnows(value byte) bool {
	for _, typeValue := range BACnetCharacterEncodingValues {
		if byte(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetCharacterEncoding(structType any) BACnetCharacterEncoding {
	castFunc := func(typ any) BACnetCharacterEncoding {
		if sBACnetCharacterEncoding, ok := typ.(BACnetCharacterEncoding); ok {
			return sBACnetCharacterEncoding
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetCharacterEncoding) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetCharacterEncoding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCharacterEncodingParse(ctx context.Context, theBytes []byte) (BACnetCharacterEncoding, error) {
	return BACnetCharacterEncodingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCharacterEncodingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCharacterEncoding, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadByte("BACnetCharacterEncoding")
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetCharacterEncoding")
	}
	if enum, ok := BACnetCharacterEncodingByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetCharacterEncoding")
		return BACnetCharacterEncoding(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetCharacterEncoding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetCharacterEncoding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteByte("BACnetCharacterEncoding", byte(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetCharacterEncoding) GetValue() byte {
	return byte(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetCharacterEncoding) PLC4XEnumName() string {
	switch e {
	case BACnetCharacterEncoding_ISO_10646:
		return "ISO_10646"
	case BACnetCharacterEncoding_IBM_Microsoft_DBCS:
		return "IBM_Microsoft_DBCS"
	case BACnetCharacterEncoding_JIS_X_0208:
		return "JIS_X_0208"
	case BACnetCharacterEncoding_ISO_10646_4:
		return "ISO_10646_4"
	case BACnetCharacterEncoding_ISO_10646_2:
		return "ISO_10646_2"
	case BACnetCharacterEncoding_ISO_8859_1:
		return "ISO_8859_1"
	}
	return fmt.Sprintf("Unknown(%v)", byte(e))
}

func (e BACnetCharacterEncoding) String() string {
	return e.PLC4XEnumName()
}
