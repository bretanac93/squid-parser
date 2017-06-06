package main

import "errors"

// Storage doc
type Storage struct {
	m map[string]uint64
}

// InsertOrUpdate Inserts or updates a user in the storage
func (s Storage) InsertOrUpdate(username string, quota uint64) (commit bool, e error) {
	if username == "" {
		e = errors.New("Username field blank")
		return
	}
	if quota < 0 {
		e = errors.New("Quota field blank or less than zero")
		return
	}
	if _, ok := s.m[username]; ok {
		s.m[username] += quota
	} else {
		s.m[username] = quota
	}
	commit = true
	return
}

// GetQuota Doc
func (s Storage) GetQuota(username string) (quota uint64, e error) {
	if username == "" {
		e = errors.New("Username field blank")
		return
	}
	var ok bool
	if quota, ok = s.m[username]; ok {
		return
	}
	quota = s.m[username]
	return
}
