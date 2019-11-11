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
)

// A Percent64 is a float64 that can be marshaled and unmarshaled as a string
// in percent format (XX%).
type Percent64 float64

// Strings implements the flag.Value interface.
func (p Percent64) String() string {
	return fmt.Sprintf("%f%%", p*100)
}

// Set implements the flag.Value interface.
func (p *Percent64) Set(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	var f float64
	_, err := fmt.Sscanf(s, "%f%%", &f)
	if err != nil {
		return fmt.Errorf("marshaler.Percent64.Set: cannot parse \"%s\"", s)
	}
	*p = Percent64(f / 100)
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (p Percent64) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (p *Percent64) UnmarshalText(text []byte) error {
	return p.Set(string(text))
}
