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

// RunOptions defines the configuration for the operation.
type RunOptions struct {
	// From is the name of the resource to copy from.
	From string `json:"from"`

	// To is the name of the resource to copy to.
	To string `json:"to"`
}

// Run will perform the operation.
func Run(options *RunOptions) error {
	if err := validate(options); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("copy from %s to %s", options.From, options.To))
	return nil
}

// validate will validate the given options.
func validate(options *RunOptions) error {
	if len(strings.TrimSpace(options.From)) <= 0 {
		return errors.New("from file is required")
	}

	if len(strings.TrimSpace(options.To)) <= 0 {
		return errors.New("to file is required")
	}
	return nil
}
