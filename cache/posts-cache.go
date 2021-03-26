package cache

import "rest_api/data"

type PostCache interface {
	Set(key string, val *data.Game)
	Get(key string) *data.Game
}
