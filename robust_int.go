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

// A RobustInt is an int that can be unmarshaled from a string.
type RobustInt int

// String implements the flag.Value interface.
func (ri RobustInt) String() string {
	return strconv.Itoa(int(ri))
}

// Set implements the flag.Value interface.
func (ri *RobustInt) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("marshaler.RobustInt.Set: cannot parse \"%s\"", s)
	}
	*ri = RobustInt(i)
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (ri *RobustInt) UnmarshalText(text []byte) error {
	return ri.Set(string(text))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ri *RobustInt) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		b = []byte(s)
	}
	return ri.UnmarshalText(b)
}
