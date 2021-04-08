// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"blog/app/dao/internal"
)

// attachmentsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type attachmentsDao struct {
	internal.AttachmentsDao
}

var (
	// Attachments is globally public accessible object for table attachments operations.
	Attachments = attachmentsDao{
		internal.Attachments,
	}
)

// Fill with you ideas below.
