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

type LoadControlObject struct {
	Object
}

func NewLoadControlObject(options ...Option) (*LoadControlObject, error) {
	l := new(LoadControlObject)
	objectType := "loadControl"
	_object_supports_cov := true
	properties := []Property{
		NewReadableProperty("presentValue", V2P(NewShedState)),
		NewOptionalProperty("stateDescription", V2P(NewCharacterString)),
		NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
		NewReadableProperty("eventState", V2P(NewEventState)),
		NewOptionalProperty("reliability", V2P(NewReliability)),
		NewWritableProperty("requestedShedLevel", V2P(NewShedLevel)),
		NewWritableProperty("startTime", V2P(NewDateTime)),
		NewWritableProperty("shedDuration", V2P(NewUnsigned)),
		NewWritableProperty("dutyWindow", V2P(NewUnsigned)),
		NewWritableProperty("enable", V2P(NewBoolean)),
		NewOptionalProperty("fullDutyBaseline", V2P(NewReal)),
		NewReadableProperty("expectedShedLevel", V2P(NewShedLevel)),
		NewReadableProperty("actualShedLevel", V2P(NewShedLevel)),
		NewWritableProperty("shedLevels", ArrayOfP(NewUnsigned, 0, 0)),
		NewReadableProperty("shedLevelDescriptions", ArrayOfP(NewCharacterString, 0, 0)),
		NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
		NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
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
		NewOptionalProperty("valueSource", V2P(NewValueSource)),
	}
	var err error
	l.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties), WithObjectSupportsCov(_object_supports_cov))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, l)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return l, nil
}
