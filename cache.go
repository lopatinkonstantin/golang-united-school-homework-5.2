package cache

import "time"

type Value struct {
	value string
	no_dead bool
	deadline time.Time
}
type Cache struct {
	data map[string] *Value
}

func NewCache() Cache {
	return Cache{data:make(map[string] *Value, 0)}
}

func (in *Cache) Get(key string) (string, bool) {
	in.RemoveOld()
	out,ok:=in.data[key]
	if ok!=true {
		return "",false
	}
	return out.value,true
}

func (in *Cache) Put(key, value string) {
	in.data[key]=&Value {value:value,no_dead:true}
}


func (in *Cache) Keys() []string {
	in.RemoveOld()
	out:=make([]string,len(in.data))
	i:=0
	for k:= range in.data {
		out[i]=k
		i++
	}
	return out
}

func (in *Cache) PutTill(key, value string, deadline time.Time) {
	in.data[key]=&Value {value:value,no_dead:false,deadline:deadline}
}

func (in *Cache) RemoveOld() {
	t:=time.Now()
	for k:= range in.data {
		if in.data[k].no_dead==false && t.After(in.data[k].deadline) {
			delete(in.data,k)
		}
	}
}