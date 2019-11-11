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
	"fmt"
	"strconv"
	"strings"
)

// A RobustUint32 is a uint32 that can be unmarshaled from a string.
type RobustUint32 uint32

// String implements the flag.Value interface.
func (ru RobustUint32) String() string {
	return strconv.FormatUint(uint64(ru), 10)
}

// Set implements the flag.Value interface.
func (ru *RobustUint32) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return fmt.Errorf("marshaler.RobustUint32.Set: cannot parse \"%s\"", s)
	}
	*ru = RobustUint32(u)
	return nil
}

// UnmarshalText implements the encoding.TextMarshaler interface.
func (ru *RobustUint32) UnmarshalText(text []byte) error {
	return ru.Set(string(text))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ru *RobustUint32) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		b = []byte(s)
	}
	return ru.UnmarshalText(b)
}
