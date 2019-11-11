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
	"strings"
)

// A CommaSeparatedString is a string slice that can be marshaled and
// unmarshaled as a comma-separated string.
type CommaSeparatedString []string

// Strings implements the flag.Value interface.
func (css CommaSeparatedString) String() string {
	return strings.Join(css, ",")
}

// Set implements the flag.Value interface.
func (css *CommaSeparatedString) Set(s string) error {
	*css = strings.Split(s, ",")
	for i := 0; i < len(*css); i++ {
		// Trim white space from the string.
		(*css)[i] = strings.TrimSpace((*css)[i])
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (css CommaSeparatedString) MarshalText() ([]byte, error) {
	return []byte(css.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (css *CommaSeparatedString) UnmarshalText(text []byte) error {
	return css.Set(string(text))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (css *CommaSeparatedString) UnmarshalJSON(b []byte) error {
	var ss []string
	if err := json.Unmarshal(b, &ss); err == nil {
		*css = CommaSeparatedString(ss)
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return css.Set(s)
}
