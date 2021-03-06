// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"blog/app/dao/internal"
)

// postTagsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type postTagsDao struct {
	internal.PostTagsDao
}

var (
	// PostTags is globally public accessible object for table post_tags operations.
	PostTags = postTagsDao{
		internal.PostTags,
	}
)

// Fill with you ideas below.
