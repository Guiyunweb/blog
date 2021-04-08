// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"blog/app/dao/internal"
)

// linksDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type linksDao struct {
	internal.LinksDao
}

var (
	// Links is globally public accessible object for table links operations.
	Links = linksDao{
		internal.Links,
	}
)

// Fill with you ideas below.
