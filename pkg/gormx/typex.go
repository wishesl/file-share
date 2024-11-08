package gormx

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ListInt64 []int64

// Value 接口，Value 返回 json value any -> string
func (j ListInt64) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 接口，Scan 将 value 扫描至 Jsonb
func (j *ListInt64) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return err
	}
	return nil
}

type ListUint []uint

// Value 接口，Value 返回 json value any -> string
func (j ListUint) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 接口，Scan 将 value 扫描至 Jsonb
func (j *ListUint) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return err
	}
	return nil
}

type ListInt []int

// Value 接口，Value 返回 json value any -> string
func (j ListInt) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 接口，Scan 将 value 扫描至 Jsonb
func (j *ListInt) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return err
	}
	return nil
}

type ListString []string

// Value 接口，Value 返回 json value any -> string
func (j ListString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 接口，Scan 将 value 扫描至 Jsonb
func (j *ListString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return err
	}
	return nil
}

type MapString map[string]string

// Value 接口，Value 返回 json value any -> string
func (j MapString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 接口，Scan 将 value 扫描至 Jsonb
func (j *MapString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return err
	}
	return nil
}
