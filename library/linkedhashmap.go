package aoc

type Linkedhashmap[K comparable, V interface{}] struct {
	mapOfKeys map[K]*linkedHashmapNode[K, V]
	startNode *linkedHashmapNode[K, V]
}
type linkedHashmapNode[K comparable, V interface{}] struct {
	before *linkedHashmapNode[K, V]
	next   *linkedHashmapNode[K, V]
	value  V
	key    K
}

func (l *Linkedhashmap[K, V]) Put(key K, value V) {
	if l == nil {
		panic("assignment in null linkedhashmap")
	}

	// if key already exists in the map, just update the value
	if _, ok := l.mapOfKeys[key]; ok {
		l.mapOfKeys[key].value = value
		return
	}

	node := &linkedHashmapNode[K, V]{
		value: value,
		key:   key,
	}

	// if no key exists in the map, simple create a node and assign it in map
	if l.mapOfKeys == nil {
		l.mapOfKeys = make(map[K]*linkedHashmapNode[K, V])
		l.mapOfKeys[key] = node
		l.startNode = node
		return
	}

	// if key does not exits, add it to the end of the linked list
	l.mapOfKeys[key] = node
	lastNode := l.startNode.before
	if lastNode == nil {
		l.startNode.next = node
		l.startNode.before = node
		node.before = l.startNode
		node.next = l.startNode
		return
	}

	lastNode.next = node
	node.next = l.startNode
	node.before = lastNode
	l.startNode.before = node
}

func (l *Linkedhashmap[K, V]) Get(key K) (V, bool) {
	var value V
	if node, ok := l.mapOfKeys[key]; ok {
		return node.value, ok
	}
	return value, false
}

func (l *Linkedhashmap[K, V]) Delete(key K) {
	// if key does not exist in map, return early
	if _, ok := l.mapOfKeys[key]; !ok {
		return
	}

	if len(l.mapOfKeys) <= 1 {
		l.mapOfKeys = nil
		l.startNode = nil
		return
	}

	before := l.mapOfKeys[key].before
	after := l.mapOfKeys[key].next

	before.next = after
	after.before = before

	if l.mapOfKeys[key] == l.startNode {
		l.startNode = after
	}
	delete(l.mapOfKeys, key)
}

func (l *Linkedhashmap[K, V]) GetAllValues() (values []V) {
	if l == nil || l.startNode == nil {
		return
	}

	startNode := l.startNode
	for i := 0; i < len(l.mapOfKeys); i++ {
		values = append(values, startNode.value)
		startNode = startNode.next
	}
	return values
}
