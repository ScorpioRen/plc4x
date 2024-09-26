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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SALData is the corresponding interface of SALData
type SALData interface {
	SALDataContract
	SALDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsSALData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALData()
	// CreateBuilder creates a SALDataBuilder
	CreateSALDataBuilder() SALDataBuilder
}

// SALDataContract provides a set of functions which can be overwritten by a sub struct
type SALDataContract interface {
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// IsSALData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALData()
	// CreateBuilder creates a SALDataBuilder
	CreateSALDataBuilder() SALDataBuilder
}

// SALDataRequirements provides a set of functions which need to be implemented by a sub struct
type SALDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetApplicationId returns ApplicationId (discriminator field)
	GetApplicationId() ApplicationId
}

// _SALData is the data-structure of this message
type _SALData struct {
	_SubType SALData
	SalData  SALData
}

var _ SALDataContract = (*_SALData)(nil)

// NewSALData factory function for _SALData
func NewSALData(salData SALData) *_SALData {
	return &_SALData{SalData: salData}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataBuilder is a builder for SALData
type SALDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SALDataBuilder
	// WithSalData adds SalData (property field)
	WithOptionalSalData(SALData) SALDataBuilder
	// WithOptionalSalDataBuilder adds SalData (property field) which is build by the builder
	WithOptionalSalDataBuilder(func(SALDataBuilder) SALDataBuilder) SALDataBuilder
	// Build builds the SALData or returns an error if something is wrong
	Build() (SALDataContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataContract
}

// NewSALDataBuilder() creates a SALDataBuilder
func NewSALDataBuilder() SALDataBuilder {
	return &_SALDataBuilder{_SALData: new(_SALData)}
}

type _SALDataBuilder struct {
	*_SALData

	err *utils.MultiError
}

var _ (SALDataBuilder) = (*_SALDataBuilder)(nil)

func (m *_SALDataBuilder) WithMandatoryFields() SALDataBuilder {
	return m
}

func (m *_SALDataBuilder) WithOptionalSalData(salData SALData) SALDataBuilder {
	m.SalData = salData
	return m
}

func (m *_SALDataBuilder) WithOptionalSalDataBuilder(builderSupplier func(SALDataBuilder) SALDataBuilder) SALDataBuilder {
	builder := builderSupplier(m.SalData.CreateSALDataBuilder())
	var err error
	m.SalData, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "SALDataBuilder failed"))
	}
	return m
}

func (m *_SALDataBuilder) Build() (SALDataContract, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SALData.deepCopy(), nil
}

func (m *_SALDataBuilder) MustBuild() SALDataContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SALDataBuilder) DeepCopy() any {
	return m.CreateSALDataBuilder()
}

// CreateSALDataBuilder creates a SALDataBuilder
func (m *_SALData) CreateSALDataBuilder() SALDataBuilder {
	if m == nil {
		return NewSALDataBuilder()
	}
	return &_SALDataBuilder{_SALData: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALData) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALData(structType any) SALData {
	if casted, ok := structType.(SALData); ok {
		return casted
	}
	if casted, ok := structType.(*SALData); ok {
		return *casted
	}
	return nil
}

func (m *_SALData) GetTypeName() string {
	return "SALData"
}

func (m *_SALData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_SALData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func SALDataParse[T SALData](ctx context.Context, theBytes []byte, applicationId ApplicationId) (T, error) {
	return SALDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataParseWithBufferProducer[T SALData](applicationId ApplicationId) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := SALDataParseWithBuffer[T](ctx, readBuffer, applicationId)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func SALDataParseWithBuffer[T SALData](ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (T, error) {
	v, err := (&_SALData{}).parse(ctx, readBuffer, applicationId)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_SALData) parse(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (__sALData SALData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child SALData
	switch {
	case applicationId == ApplicationId_RESERVED: // SALDataReserved
		if _child, err = new(_SALDataReserved).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataReserved for type-switch of SALData")
		}
	case applicationId == ApplicationId_FREE_USAGE: // SALDataFreeUsage
		if _child, err = new(_SALDataFreeUsage).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataFreeUsage for type-switch of SALData")
		}
	case applicationId == ApplicationId_TEMPERATURE_BROADCAST: // SALDataTemperatureBroadcast
		if _child, err = new(_SALDataTemperatureBroadcast).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataTemperatureBroadcast for type-switch of SALData")
		}
	case applicationId == ApplicationId_ROOM_CONTROL_SYSTEM: // SALDataRoomControlSystem
		if _child, err = new(_SALDataRoomControlSystem).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataRoomControlSystem for type-switch of SALData")
		}
	case applicationId == ApplicationId_LIGHTING: // SALDataLighting
		if _child, err = new(_SALDataLighting).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataLighting for type-switch of SALData")
		}
	case applicationId == ApplicationId_VENTILATION: // SALDataVentilation
		if _child, err = new(_SALDataVentilation).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataVentilation for type-switch of SALData")
		}
	case applicationId == ApplicationId_IRRIGATION_CONTROL: // SALDataIrrigationControl
		if _child, err = new(_SALDataIrrigationControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataIrrigationControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_POOLS_SPAS_PONDS_FOUNTAINS_CONTROL: // SALDataPoolsSpasPondsFountainsControl
		if _child, err = new(_SALDataPoolsSpasPondsFountainsControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataPoolsSpasPondsFountainsControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_HEATING: // SALDataHeating
		if _child, err = new(_SALDataHeating).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataHeating for type-switch of SALData")
		}
	case applicationId == ApplicationId_AIR_CONDITIONING: // SALDataAirConditioning
		if _child, err = new(_SALDataAirConditioning).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataAirConditioning for type-switch of SALData")
		}
	case applicationId == ApplicationId_TRIGGER_CONTROL: // SALDataTriggerControl
		if _child, err = new(_SALDataTriggerControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataTriggerControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_ENABLE_CONTROL: // SALDataEnableControl
		if _child, err = new(_SALDataEnableControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataEnableControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_AUDIO_AND_VIDEO: // SALDataAudioAndVideo
		if _child, err = new(_SALDataAudioAndVideo).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataAudioAndVideo for type-switch of SALData")
		}
	case applicationId == ApplicationId_SECURITY: // SALDataSecurity
		if _child, err = new(_SALDataSecurity).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataSecurity for type-switch of SALData")
		}
	case applicationId == ApplicationId_METERING: // SALDataMetering
		if _child, err = new(_SALDataMetering).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataMetering for type-switch of SALData")
		}
	case applicationId == ApplicationId_ACCESS_CONTROL: // SALDataAccessControl
		if _child, err = new(_SALDataAccessControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataAccessControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_CLOCK_AND_TIMEKEEPING: // SALDataClockAndTimekeeping
		if _child, err = new(_SALDataClockAndTimekeeping).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataClockAndTimekeeping for type-switch of SALData")
		}
	case applicationId == ApplicationId_TELEPHONY_STATUS_AND_CONTROL: // SALDataTelephonyStatusAndControl
		if _child, err = new(_SALDataTelephonyStatusAndControl).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataTelephonyStatusAndControl for type-switch of SALData")
		}
	case applicationId == ApplicationId_MEASUREMENT: // SALDataMeasurement
		if _child, err = new(_SALDataMeasurement).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataMeasurement for type-switch of SALData")
		}
	case applicationId == ApplicationId_TESTING: // SALDataTesting
		if _child, err = new(_SALDataTesting).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataTesting for type-switch of SALData")
		}
	case applicationId == ApplicationId_MEDIA_TRANSPORT_CONTROL: // SALDataMediaTransport
		if _child, err = new(_SALDataMediaTransport).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataMediaTransport for type-switch of SALData")
		}
	case applicationId == ApplicationId_ERROR_REPORTING: // SALDataErrorReporting
		if _child, err = new(_SALDataErrorReporting).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataErrorReporting for type-switch of SALData")
		}
	case applicationId == ApplicationId_HVAC_ACTUATOR: // SALDataHvacActuator
		if _child, err = new(_SALDataHvacActuator).parse(ctx, readBuffer, m, applicationId); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SALDataHvacActuator for type-switch of SALData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [applicationId=%v]", applicationId)
	}

	var salData SALData
	_salData, err := ReadOptionalField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(applicationId)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	if _salData != nil {
		salData = *_salData
		m.SalData = salData
	}

	if closeErr := readBuffer.CloseContext("SALData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALData")
	}

	return _child, nil
}

func (pm *_SALData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SALData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SALData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SALData")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteOptionalField[SALData](ctx, "salData", GetRef(m.GetSalData()), WriteComplex[SALData](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'salData' field")
	}

	if popErr := writeBuffer.PopContext("SALData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SALData")
	}
	return nil
}

func (m *_SALData) IsSALData() {}

func (m *_SALData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALData) deepCopy() *_SALData {
	if m == nil {
		return nil
	}
	_SALDataCopy := &_SALData{
		nil, // will be set by child
		m.SalData.DeepCopy().(SALData),
	}
	return _SALDataCopy
}
