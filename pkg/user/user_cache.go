package user

import (
	"errors"
	"sync"
)

type UserCache struct {
	cache map[int]User
}

var cacheMutex sync.RWMutex

func NewUserCache() *UserCache {
	return &UserCache{
		cache: make(map[int]User),
	}
}

func (uc *UserCache) AddUser(user User) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	uc.cache[len(uc.cache)+1] = user
}

func (uc *UserCache) GetUser(id int) (User, error) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	user, ok := uc.cache[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (uc *UserCache) DeleteUser(id int) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if _, ok := uc.cache[id]; !ok {
		return errors.New("user not found")
	}

	delete(uc.cache, id)
	return nil
}
