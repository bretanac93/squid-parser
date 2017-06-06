package main

import "errors"

// Stat doc
type Stat struct {
	current, quota uint64
}

// Storage doc
type Storage struct {
	m map[string]Stat
}

// InsertOrUpdate Inserts or updates a user in the storage
func (s Storage) InsertOrUpdate(username string, usage uint64) (commit bool, e error) {
	if username == "" {
		e = errors.New("Username field blank")
		return
	}
	if usage < 0 {
		e = errors.New("Quota field blank or less than zero")
		return
	}
	if _, exists := s.m[username]; exists {
		x := s.m[username]
		x.current += usage

		s.m[username] = x
	} else {
		quota := GetQuota(username)
		s.m[username] = Stat{current: usage, quota: quota}
	}
	commit = true
	return
}

// GetUsage gets the total usage recorded in the storage
func (s Storage) GetUsage(username string) (usage uint64, e error) {
	if username == "" {
		e = errors.New("Username field blank")
		return
	}
	if stat, ok := s.m[username]; ok {
		usage = stat.current
		return
	}
	return
}

// GetQuota gets the quota of a user by checking the AD group where he/she belongs
func (s Storage) GetQuota(username string) (quota uint64, e error) {
	// Implement this method, right now every used is treat as student
	return uint64(30), nil
}
