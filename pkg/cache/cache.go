package cache

import "sync"

var Cache = make(map[string]string)
var CacheMu sync.RWMutex

func Init() {

}
