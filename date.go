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

// A Date is a time.Time that can be marshaled and unmarshaled as a string in
// YYYY-MM-DD format.
type Date time.Time

// Strings implements the flag.Value interface.
func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}

// Set implements the flag.Value interface.
func (d *Date) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("marshaler.Date.Set: cannot parse \"%s\"", s)
	}
	*d = Date(t)
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Date) UnmarshalText(text []byte) error {
	return d.Set(string(text))
}

// Format wraps time.Time.Format.
func (d Date) Format(layout string) string {
	return time.Time(d).Format(layout)
}
