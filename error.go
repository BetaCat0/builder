// Copyright 2016 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builder

import "errors"

var (
	// ErrNotSupportType not supported SQL type error
	ErrNotSupportType = errors.New("Not supported SQL type")
	// ErrNoNotInConditions no NOT IN params error
	ErrNoNotInConditions = errors.New("No NOT IN conditions")
	// ErrNoInConditions no IN params error
	ErrNoInConditions = errors.New("No IN conditions")
	// ErrNeedMoreArguments need more arguments
	ErrNeedMoreArguments = errors.New("Need more sql arguments")
	// ErrNoTableName no table name
	ErrNoTableName = errors.New("No table indicated")
	// ErrNoColumnToInsert no column to update
	ErrNoColumnToUpdate = errors.New("No column(s) to update")
	// ErrNoColumnToInsert no column to update
	ErrNoColumnToInsert = errors.New("No column(s) to insert")
	// ErrNotSupportDialectType not supported dialect type error
	ErrNotSupportDialectType = errors.New("Not supported dialect type")
	// ErrNotUnexpectedUnionConditions using union in a wrong way
	ErrNotUnexpectedUnionConditions = errors.New("Unexpected conditional fields in UNION query")
	// ErrUnsupportedUnionMembers unexpected members in UNION query
	ErrUnsupportedUnionMembers = errors.New("Unexpected members in UNION query")
)
