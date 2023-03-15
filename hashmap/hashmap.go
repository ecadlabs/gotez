// Package hashmap implements a generic wrapper around a built in map type which allows
// interface types like PublicKeyHash to be used as map keys
package hashmap

type Comparable[K any] interface {
	comparable
	ToKey() K
}

type ToComparable[K any, H Comparable[K]] interface {
	ToComparable() H
}

type KV[K, V any] struct {
	Key K
	Val V
}

type HashMap[H Comparable[K], K ToComparable[K, H], V any] map[H]V

func (m HashMap[H, K, V]) Insert(key K, val V) (V, bool) {
	k := key.ToComparable()
	v, ok := m[k]
	m[k] = val
	return v, ok
}

func (m HashMap[H, K, V]) Get(key K) (V, bool) {
	k := key.ToComparable()
	v, ok := m[k]
	return v, ok
}

func (m HashMap[H, K, V]) ForEach(cb func(key K, val V) bool) {
	for k, v := range m {
		if !cb(k.ToKey(), v) {
			break
		}
	}
}

func New[H Comparable[K], K ToComparable[K, H], V any](init []KV[K, V]) HashMap[H, K, V] {
	m := make(HashMap[H, K, V])
	for _, kv := range init {
		m.Insert(kv.Key, kv.Val)
	}
	return m
}
