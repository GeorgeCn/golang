package main

import (
	"time"
	"github.com/gogo/protobuf/io"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // 比如 "MP3"、"WAV"
}

type Video interface {
	 Stream() (io.ReadCloser, error)
	 RunningTime() time.Duration
	 Format() string //
	 Resolution() (x, y int)
}