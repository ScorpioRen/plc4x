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

// MediaTransportControlDataPlay is the corresponding interface of MediaTransportControlDataPlay
type MediaTransportControlDataPlay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// IsMediaTransportControlDataPlay is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataPlay()
	// CreateBuilder creates a MediaTransportControlDataPlayBuilder
	CreateMediaTransportControlDataPlayBuilder() MediaTransportControlDataPlayBuilder
}

// _MediaTransportControlDataPlay is the data-structure of this message
type _MediaTransportControlDataPlay struct {
	MediaTransportControlDataContract
}

var _ MediaTransportControlDataPlay = (*_MediaTransportControlDataPlay)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataPlay)(nil)

// NewMediaTransportControlDataPlay factory function for _MediaTransportControlDataPlay
func NewMediaTransportControlDataPlay(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataPlay {
	_result := &_MediaTransportControlDataPlay{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataPlayBuilder is a builder for MediaTransportControlDataPlay
type MediaTransportControlDataPlayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MediaTransportControlDataPlayBuilder
	// Build builds the MediaTransportControlDataPlay or returns an error if something is wrong
	Build() (MediaTransportControlDataPlay, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataPlay
}

// NewMediaTransportControlDataPlayBuilder() creates a MediaTransportControlDataPlayBuilder
func NewMediaTransportControlDataPlayBuilder() MediaTransportControlDataPlayBuilder {
	return &_MediaTransportControlDataPlayBuilder{_MediaTransportControlDataPlay: new(_MediaTransportControlDataPlay)}
}

type _MediaTransportControlDataPlayBuilder struct {
	*_MediaTransportControlDataPlay

	err *utils.MultiError
}

var _ (MediaTransportControlDataPlayBuilder) = (*_MediaTransportControlDataPlayBuilder)(nil)

func (m *_MediaTransportControlDataPlayBuilder) WithMandatoryFields() MediaTransportControlDataPlayBuilder {
	return m
}

func (m *_MediaTransportControlDataPlayBuilder) Build() (MediaTransportControlDataPlay, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._MediaTransportControlDataPlay.deepCopy(), nil
}

func (m *_MediaTransportControlDataPlayBuilder) MustBuild() MediaTransportControlDataPlay {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_MediaTransportControlDataPlayBuilder) DeepCopy() any {
	return m.CreateMediaTransportControlDataPlayBuilder()
}

// CreateMediaTransportControlDataPlayBuilder creates a MediaTransportControlDataPlayBuilder
func (m *_MediaTransportControlDataPlay) CreateMediaTransportControlDataPlayBuilder() MediaTransportControlDataPlayBuilder {
	if m == nil {
		return NewMediaTransportControlDataPlayBuilder()
	}
	return &_MediaTransportControlDataPlayBuilder{_MediaTransportControlDataPlay: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataPlay) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataPlay(structType any) MediaTransportControlDataPlay {
	if casted, ok := structType.(MediaTransportControlDataPlay); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataPlay); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataPlay) GetTypeName() string {
	return "MediaTransportControlDataPlay"
}

func (m *_MediaTransportControlDataPlay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MediaTransportControlDataPlay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataPlay) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataPlay MediaTransportControlDataPlay, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataPlay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataPlay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataPlay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataPlay")
	}

	return m, nil
}

func (m *_MediaTransportControlDataPlay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataPlay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataPlay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataPlay")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataPlay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataPlay")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataPlay) IsMediaTransportControlDataPlay() {}

func (m *_MediaTransportControlDataPlay) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataPlay) deepCopy() *_MediaTransportControlDataPlay {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataPlayCopy := &_MediaTransportControlDataPlay{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
	}
	m.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataPlayCopy
}

func (m *_MediaTransportControlDataPlay) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
