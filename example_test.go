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

package marshaler_test

import (
	"fmt"
	"log"

	"github.com/tradyfinance/marshaler"
)

func ExampleCommaSeparatedString() {
	var css marshaler.CommaSeparatedString
	if err := css.UnmarshalText([]byte("foo,bar")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(css.String())
	// Output: foo,bar
}
