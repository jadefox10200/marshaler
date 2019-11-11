// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package marshaler

import (
	"encoding/json"
	"strconv"
	"time"
)

// A UnixTimestamp is a time.Time that os marshaled and unmarshaled as a UNIX
// timestamp.
type UnixTimestamp time.Time

func (ut UnixTimestamp) String() string {
	return time.Time(ut).String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ut UnixTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(ut).Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ut *UnixTimestamp) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	*ut = UnixTimestamp(time.Unix(i, 0))
	return nil
}

// Format wraps time.Time.Format.
func (ut UnixTimestamp) Format(layout string) string {
	return time.Time(ut).Format(layout)
}
