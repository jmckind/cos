// Copyright 2020 COS Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cos

import (
	"errors"
	"fmt"
	"strings"
)

// Options defines the configuration for the operation.
type Options struct {
	// From is the name of the resource to copy from.
	From string `json:"from"`

	// To is the name of the resource to copy to.
	To string `json:"to"`
}

// Execute will perform the operation.
func Execute(o *Options) error {
	if err := validate(o); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("copy from %s to %s", o.From, o.To))
	return nil
}

// validate will validate the given options.
func validate(o *Options) error {
	if len(strings.TrimSpace(o.From)) <= 0 {
		return errors.New("from file is required")
	}

	if len(strings.TrimSpace(o.To)) <= 0 {
		return errors.New("to file is required")
	}
	return nil
}
