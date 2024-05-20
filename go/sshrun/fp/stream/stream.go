package stream

import "github.com/jucardi/go-streams/v2/streams"

func From[T comparable](set []T) streams.IStream[T] {
	return streams.FromArray(set)
}
