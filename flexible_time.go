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
	"fmt"
	"strings"
	"time"
)

// A FlexibleTime is a time.Time that is flexible in the format used for
// unmarshaling.
type FlexibleTime time.Time

// flexibleTimeLayouts are the layouts used for parsing a FlexibleTime.
var flexibleTimeLayouts = [...]string{
	"2006-01-02",
	"2006-01-02 15:04",
	"2006-01-02 15:04:05",
	time.RFC3339,
}

// Strings implements the flag.Value interface.
func (ft FlexibleTime) String() string {
	return time.Time(ft).String()
}

// Set implements the flag.Value interface.
func (ft *FlexibleTime) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	for _, format := range flexibleTimeLayouts {
		t, err := time.Parse(format, s)
		if err == nil {
			*ft = FlexibleTime(t)
			return nil
		}
	}
	return fmt.Errorf("marshaler.FlexibleTime.Set: cannot parse \"%s\"", s)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (ft FlexibleTime) MarshalText() ([]byte, error) {
	return []byte(ft.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (ft *FlexibleTime) UnmarshalText(text []byte) error {
	return ft.Set(string(text))
}

// Format wraps time.Time.Format.
func (ft FlexibleTime) Format(layout string) string {
	return time.Time(ft).Format(layout)
}
