package api

import (
	"access-tool/repo"
)

func QueryTeamIsExist(name string) bool {
	r := repo.AccessRepo{}
	t, err := r.QueryTeamByName(name)
	if err != nil {
		return false
	}
	if t != nil {
		return true
	}
	return false
}
