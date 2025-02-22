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

package object

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type AccessZoneObject struct {
	Object
}

func NewAccessZoneObject(options ...Option) (*AccessZoneObject, error) {
	a := new(AccessZoneObject)
	objectType := "accessZone"
	properties := []Property{
		NewWritableProperty("globalIdentifier", V2P(NewUnsigned)),
		NewReadableProperty("occupancyState", V2P(NewAccessZoneOccupancyState)),
		NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
		NewReadableProperty("eventState", V2P(NewEventState)),
		NewReadableProperty("reliability", V2P(NewReliability)),
		NewReadableProperty("outOfService", V2P(NewBoolean)),
		NewOptionalProperty("occupancyCount", V2P(NewUnsigned)),
		NewOptionalProperty("occupancyCountEnable", V2P(NewBoolean)),
		NewOptionalProperty("adjustValue", V2P(NewInteger)),
		NewOptionalProperty("occupancyUpperLimit", V2P(NewUnsigned)),
		NewOptionalProperty("occupancyLowerLimit", V2P(NewUnsigned)),
		NewOptionalProperty("credentialsInZone", ListOfP(NewDeviceObjectReference)),
		NewOptionalProperty("lastCredentialAdded", V2P(NewDeviceObjectReference)),
		NewOptionalProperty("lastCredentialAddedTime", V2P(NewDateTime)),
		NewOptionalProperty("lastCredentialRemoved", V2P(NewDeviceObjectReference)),
		NewOptionalProperty("lastCredentialRemovedTime", V2P(NewDateTime)),
		NewOptionalProperty("passbackMode", V2P(NewAccessPassbackMode)),
		NewOptionalProperty("passbackTimeout", V2P(NewUnsigned)),
		NewReadableProperty("entryPoints", ListOfP(NewDeviceObjectReference)),
		NewReadableProperty("exitPoints", ListOfP(NewDeviceObjectReference)),
		NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
		NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
		NewOptionalProperty("alarmValues", ListOfP(NewAccessZoneOccupancyState)),
		NewOptionalProperty("eventEnable", V2P(NewEventTransitionBits)),
		NewOptionalProperty("ackedTransitions", V2P(NewEventTransitionBits)),
		NewOptionalProperty("notifyType", V2P(NewNotifyType)),
		NewOptionalProperty("eventTimeStamps", ArrayOfP(NewTimeStamp, 3, 0)),
		NewOptionalProperty("eventMessageTexts", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventMessageTextsConfig", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventDetectionEnable", V2P(NewBoolean)),
		NewOptionalProperty("eventAlgorithmInhibitRef", V2P(NewObjectPropertyReference)),
		NewOptionalProperty("eventAlgorithmInhibit", V2P(NewBoolean)),
		NewOptionalProperty("timeDelayNormal", V2P(NewUnsigned)),
		NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
	}
	var err error
	a.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, a)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return a, nil
}
