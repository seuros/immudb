/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sessions

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	op := Options{}

	op.WithMaxSessionAgeTime(time.Second).
		WithSessionGuardCheckInterval(2 * time.Second).
		WithMaxSessionInactivityTime(3 * time.Second).
		WithTimeout(4 * time.Second)

	assert.Equal(t, time.Second, op.MaxSessionAgeTime)
	assert.Equal(t, 2*time.Second, op.SessionGuardCheckInterval)
	assert.Equal(t, 3*time.Second, op.MaxSessionInactivityTime)
	assert.Equal(t, 4*time.Second, op.Timeout)
}

func TestOptionsValidate(t *testing.T) {
	op := DefaultOptions()
	err := op.Validate()
	require.NoError(t, err)

	for _, op := range []*Options{
		DefaultOptions().WithSessionGuardCheckInterval(0),
		DefaultOptions().WithSessionGuardCheckInterval(-1 * time.Second),
		DefaultOptions().WithMaxSessionInactivityTime(-1 * time.Second),
		DefaultOptions().WithMaxSessionAgeTime(-1 * time.Second),
		DefaultOptions().WithTimeout(-1 * time.Second),
	} {
		t.Run(fmt.Sprintf("%+v", op), func(t *testing.T) {
			err = op.Validate()
			require.ErrorIs(t, err, ErrInvalidOptionsProvided)
		})
	}

}
