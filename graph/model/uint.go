package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUint(t uint) graphql.Marshaler {
	return MarshalUint64(uint64(t))
}

func UnmarshalUint(v interface{}) (uint, error) {
	i, err := UnmarshalUint64(v)
	return uint(i), err
}

func MarshalUint64(t uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.FormatUint(t, 10))
		if err != nil {
			return
		}
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseUint(t, 10, 64)
	case int:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case json.Number:
		i, err := t.Int64()
		return uint64(i), err
	case float64:
		return uint64(t), nil
	}

	return 0, fmt.Errorf("unable to unmarshal uint64: %#v %T", v, v)
}

func MarshalUint32(t uint32) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.FormatUint(uint64(t), 10))
		if err != nil {
			return
		}
	})
}

func UnmarshalUint32(v interface{}) (uint32, error) {
	switch t := v.(type) {
	case string:
		u, err := strconv.ParseUint(t, 10, 32)
		return uint32(u), err
	case int:
		return uint32(t), nil
	case int64:
		return uint32(t), nil
	case json.Number:
		i, err := t.Int64()
		return uint32(i), err
	case float64:
		return uint32(t), nil
	}

	return 0, fmt.Errorf("unable to unmarshal uint32: %#v %T", v, v)
}
