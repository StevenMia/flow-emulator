/*
 * Flow Emulator
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package unittest

import (
	"github.com/onflow/cadence/encoding/ccf"
	"github.com/onflow/flow-emulator/convert"
	sdk "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/test"
	flowgo "github.com/onflow/flow-go/model/flow"

	"github.com/onflow/flow-emulator/types"
)

func TransactionFixture() flowgo.TransactionBody {
	return *convert.SDKTransactionToFlow(*test.TransactionGenerator().New())
}

func StorableTransactionResultFixture() types.StorableTransactionResult {
	events := test.EventGenerator()

	eventA, _ := SDKEventToFlow(events.New())
	eventB, _ := SDKEventToFlow(events.New())

	return types.StorableTransactionResult{
		ErrorCode:    42,
		ErrorMessage: "foo",
		Logs:         []string{"a", "b", "c"},
		Events: []flowgo.Event{
			eventA,
			eventB,
		},
	}
}

func SDKEventToFlow(event sdk.Event) (flowgo.Event, error) {
	payload, err := ccf.Encode(event.Value)
	if err != nil {
		return flowgo.Event{}, err
	}

	return flowgo.Event{
		Type:             flowgo.EventType(event.Type),
		TransactionID:    convert.SDKIdentifierToFlow(event.TransactionID),
		TransactionIndex: uint32(event.TransactionIndex),
		EventIndex:       uint32(event.EventIndex),
		Payload:          payload,
	}, nil
}
