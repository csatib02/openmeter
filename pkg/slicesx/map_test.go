// Copyright © 2024 Tailfin Cloud Inc.
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

package slicesx

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}

	actual := Map(s, func(v int) string {
		return fmt.Sprintf("%d", v)
	})

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("map failed")
	}
}
