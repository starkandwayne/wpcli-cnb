/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// ScratchDir returns a safe scratch directory for tests to modify.
func ScratchDir(t *testing.T, prefix string) string {
	t.Helper()

	tmp, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatal(err)
	}

	abs, err := filepath.EvalSymlinks(tmp)
	if err != nil {
		t.Fatal(err)
	}

	return abs
}
