// Copyright (C) 2018 Allen Li
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

// Package xdg implements access to environment variables defined by
// the XDG Base Directory Specification.
package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

var home = os.Getenv("HOME")

func ConfigHome() string {
	p := os.Getenv("XDG_CONFIG_HOME")
	if p == "" {
		return filepath.Join(home, ".config")
	}
	return p
}

func ConfigDirs() []string {
	p := os.Getenv("XDG_CONFIG_DIRS")
	if p == "" {
		return []string{"/etc/xdg"}
	}
	return strings.Split(p, ":")
}

func DataHome() string {
	p := os.Getenv("XDG_DATA_HOME")
	if p == "" {
		return filepath.Join(home, ".local", "share")
	}
	return p
}

func DataDirs() []string {
	p := os.Getenv("XDG_DATA_DIRS")
	if p == "" {
		return []string{"/usr/local/share", "/usr/share"}
	}
	return strings.Split(p, ":")
}

func CacheHome() string {
	p := os.Getenv("XDG_CACHE_HOME")
	if p == "" {
		return filepath.Join(home, ".cache")
	}
	return p
}

func StateHome() string {
	p := os.Getenv("XDG_STATE_HOME")
	if p == "" {
		return filepath.Join(home, ".local", "state")
	}
	return p
}

func RuntimeDir() string {
	return os.Getenv("XDG_RUNTIME_DIR")
}
