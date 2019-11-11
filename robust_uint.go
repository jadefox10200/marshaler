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

// A RobustUint is a uint that can be unmarshaled from a string.
type RobustUint uint

// String implements the flag.Value interface.
func (ri RobustUint) String() string {
	return strconv.FormatUint(uint64(ri), 10)
}

// Set implements the flag.Value interface.
func (ri *RobustUint) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return fmt.Errorf("marshaler.RobustUint.Set: cannot parse \"%s\"", s)
	}
	*ri = RobustUint(i)
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (ri *RobustUint) UnmarshalText(text []byte) error {
	return ri.Set(string(text))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ri *RobustUint) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		b = []byte(s)
	}
	return ri.UnmarshalText(b)
}
