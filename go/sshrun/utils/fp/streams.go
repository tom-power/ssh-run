package fp

import "github.com/jucardi/go-streams/v2/streams"

func Stream[T comparable](set []T) streams.IStream[T] {
	return streams.FromArray(set)
}
