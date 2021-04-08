// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"blog/app/dao/internal"
)

// optionsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type optionsDao struct {
	internal.OptionsDao
}

var (
	// Options is globally public accessible object for table options operations.
	Options = optionsDao{
		internal.Options,
	}
)

// Fill with you ideas below.