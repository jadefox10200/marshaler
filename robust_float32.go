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

// A RobustFloat32 is a float32 that can be unmarshaled from a string.
type RobustFloat32 float32

// Strings implements the flag.Value interface.
func (rf RobustFloat32) String() string {
	return fmt.Sprintf("%f", rf)
}

// Set implements the flag.Value interface.
func (rf *RobustFloat32) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return fmt.Errorf("marshaler.RobustFloat32.Set: cannot parse \"%s\"", s)
	}
	*rf = RobustFloat32(f)
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (rf *RobustFloat32) UnmarshalText(text []byte) error {
	return rf.Set(string(text))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (rf *RobustFloat32) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		b = []byte(s)
	}
	return rf.UnmarshalText(b)
}
