// Copyright 2021 The Alpaca Authors
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

package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBlocklistExpiry(t *testing.T) {
	b := newBlocklist()
	var now time.Time
	b.now = func() time.Time { return now }

	b.add("foo")
	assert.True(t, b.contains("foo"))
	assert.False(t, b.contains("bar"))

	now = now.Add(3 * time.Minute)
	b.add("bar")
	assert.True(t, b.contains("foo"))
	assert.True(t, b.contains("bar"))

	now = now.Add(3 * time.Minute)
	assert.False(t, b.contains("foo"))
	assert.True(t, b.contains("bar"))

	now = now.Add(3 * time.Minute)
	assert.False(t, b.contains("foo"))
	assert.False(t, b.contains("bar"))
}

func TestBlocklistDuplicateEntry(t *testing.T) {
	b := newBlocklist()
	var now time.Time
	b.now = func() time.Time { return now }
	b.add("foo")
	now = now.Add(3*time.Minute)
	b.add("foo")
	now = now.Add(3*time.Minute)
	b.contains("foo")
	now = now.Add(3*time.Minute)
	b.contains("foo")
}
