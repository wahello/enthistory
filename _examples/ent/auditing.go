// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"

	"github.com/flume/enthistory/_examples/ent/characterhistory"
	"github.com/flume/enthistory/_examples/ent/friendshiphistory"
)

type Change struct {
	FieldName string
	Old       any
	New       any
}

func NewChange(fieldName string, old, new any) Change {
	return Change{
		FieldName: fieldName,
		Old:       old,
		New:       new,
	}
}

type HistoryDiff[T any] struct {
	Old     *T
	New     *T
	Changes []Change
}

var (
	MismatchedRefError    = errors.New("cannot take diff of histories with different Refs")
	IdenticalHistoryError = errors.New("cannot take diff of identical history")
)

func (ch *CharacterHistory) changes(new *CharacterHistory) []Change {
	var changes []Change
	if !ch.CreatedAt.Equal(new.CreatedAt) {
		changes = append(changes, NewChange(characterhistory.FieldCreatedAt, ch.CreatedAt, new.CreatedAt))
	}
	if !ch.UpdatedAt.Equal(new.UpdatedAt) {
		changes = append(changes, NewChange(characterhistory.FieldUpdatedAt, ch.UpdatedAt, new.UpdatedAt))
	}
	if ch.Age != new.Age {
		changes = append(changes, NewChange(characterhistory.FieldAge, ch.Age, new.Age))
	}
	if ch.Name != new.Name {
		changes = append(changes, NewChange(characterhistory.FieldName, ch.Name, new.Name))
	}
	return changes
}

func (ch *CharacterHistory) Diff(history *CharacterHistory) (*HistoryDiff[CharacterHistory], error) {
	if ch.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	chUnix, historyUnix := ch.HistoryTime.Unix(), history.HistoryTime.Unix()
	chOlder := chUnix < historyUnix || (chUnix == historyUnix && ch.ID < history.ID)
	historyOlder := chUnix > historyUnix || (chUnix == historyUnix && ch.ID > history.ID)

	if chOlder {
		return &HistoryDiff[CharacterHistory]{
			Old:     ch,
			New:     history,
			Changes: ch.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[CharacterHistory]{
			Old:     history,
			New:     ch,
			Changes: history.changes(ch),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (fh *FriendshipHistory) changes(new *FriendshipHistory) []Change {
	var changes []Change
	if !fh.CreatedAt.Equal(new.CreatedAt) {
		changes = append(changes, NewChange(friendshiphistory.FieldCreatedAt, fh.CreatedAt, new.CreatedAt))
	}
	if !fh.UpdatedAt.Equal(new.UpdatedAt) {
		changes = append(changes, NewChange(friendshiphistory.FieldUpdatedAt, fh.UpdatedAt, new.UpdatedAt))
	}
	if fh.CharacterID != new.CharacterID {
		changes = append(changes, NewChange(friendshiphistory.FieldCharacterID, fh.CharacterID, new.CharacterID))
	}
	if fh.FriendID != new.FriendID {
		changes = append(changes, NewChange(friendshiphistory.FieldFriendID, fh.FriendID, new.FriendID))
	}
	return changes
}

func (fh *FriendshipHistory) Diff(history *FriendshipHistory) (*HistoryDiff[FriendshipHistory], error) {
	if fh.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	fhUnix, historyUnix := fh.HistoryTime.Unix(), history.HistoryTime.Unix()
	fhOlder := fhUnix < historyUnix || (fhUnix == historyUnix && fh.ID < history.ID)
	historyOlder := fhUnix > historyUnix || (fhUnix == historyUnix && fh.ID > history.ID)

	if fhOlder {
		return &HistoryDiff[FriendshipHistory]{
			Old:     fh,
			New:     history,
			Changes: fh.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[FriendshipHistory]{
			Old:     history,
			New:     fh,
			Changes: history.changes(fh),
		}, nil
	}
	return nil, IdenticalHistoryError
}