package helper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Parser struct {
	err error
}

func convertToUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return u, nil
	case json.Number:
		i64, err := v.Int64()
		if err != nil {
			return 0, err
		}
		u := uint64(i64)
		return u, nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		u := v.(uint64)
		return u, nil
	default:
		return 0, fmt.Errorf("%v cannot convert to uint64", v)
	}
}

func (p *Parser) ParseString(value interface{}) *string {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	switch v := value.(type) {
	case string:
		return StringToPtr(v)
	case json.Number:
		return StringToPtr(v.String())
	default:
		return StringToPtr(fmt.Sprint(v))
	}
}

func (p *Parser) parseUint8(value interface{}) *uint8 {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	u64, err := convertToUint64(value)
	if err != nil {
		p.err = err
		return nil
	}
	return Uint8ToPtr(uint8(u64))
}

func (p *Parser) parseUint32(value interface{}) *uint32 {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	u64, err := convertToUint64(value)
	if err != nil {
		p.err = err
		return nil
	}
	return Uint32ToPtr(uint32(u64))
}

func (p *Parser) parseUint64(value interface{}) *uint64 {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	u64, err := convertToUint64(value)
	if err != nil {
		p.err = err
		return nil
	}
	return Uint64ToPtr(u64)
}

func (p *Parser) parseInt(value interface{}) *int {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	u64, err := convertToUint64(value)
	if err != nil {
		p.err = err
		return nil
	}
	i := int(u64)
	return &i
}

func (p *Parser) parseBool(value interface{}) *bool {
	if value == "" || value == nil || p.err != nil {
		return nil
	}
	strToBoolPtr := func(s string) *bool {
		b, err := strconv.ParseBool(s)
		if err != nil {
			p.err = err
			return nil
		}
		return BoolToPtr(b)
	}
	switch v := value.(type) {
	case bool:
		return BoolToPtr(v)
	case string:
		return strToBoolPtr(v)
	case json.Number:
		return strToBoolPtr(v.String())
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		u64 := value.(uint64)
		return BoolToPtr(u64 != 0)
	default:
		return nil
	}
}

func (p *Parser) parseTime(value string) time.Time {
	if value == "" || p.err != nil {
		return time.Time{}
	}
	v, err := time.Parse(time.RFC3339, value)
	if err != nil {
		p.err = err
		return v
	}
	return v
}
