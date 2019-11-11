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

// A UnixTimestampMS is a time.Time that is marshaled and unmarshaled as a UNIX
// timestamp in milliseconds (ms).
type UnixTimestampMS time.Time

func (utms UnixTimestampMS) String() string {
	return time.Time(utms).String()
}

// MarshalJSON implements the json.Marshaler interface.
func (utms UnixTimestampMS) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(utms).UnixNano()/1000000, 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (utms *UnixTimestampMS) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	*utms = UnixTimestampMS(time.Unix(i/1000, i%1000*1000000))
	return nil
}

// Format wraps time.Time.Format.
func (utms UnixTimestampMS) Format(layout string) string {
	return time.Time(utms).Format(layout)
}
