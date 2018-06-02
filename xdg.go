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
