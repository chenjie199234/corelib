package common

func BkdrhashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	seed := uint64(131)
	hash := uint64(0)
	for _, v := range data {
		hash = hash*seed + uint64(v)
	}
	return hash % total
}
func DjbhashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(5381)
	for _, v := range data {
		hash = ((hash << 5) + hash) + uint64(v)
	}
	return hash % total
}
func FnvhashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(2166136261)
	for _, v := range data {
		hash *= uint64(16777619)
		hash ^= uint64(v)
	}
	return hash % total
}
func DekhashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(len(data))
	for _, v := range data {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint64(v)
	}
	return hash % total
}
func RshashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	seed := uint64(63689)
	hash := uint64(0)
	for _, v := range data {
		hash = hash*seed + uint64(v)
		seed *= uint64(378551)
	}
	return hash % total
}
func SdbmhashString(data string, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(0)
	for _, v := range data {
		hash = uint64(v) + (hash << 6) + (hash << 16) - hash
	}
	return hash % total
}
func BkdrhashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	seed := uint64(131)
	hash := uint64(0)
	for _, v := range data {
		hash = hash*seed + uint64(v)
	}
	return hash % total
}
func DjbhashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(5381)
	for _, v := range data {
		hash = ((hash << 5) + hash) + uint64(v)
	}
	return hash % total
}
func FnvhashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(2166136261)
	for _, v := range data {
		hash *= uint64(16777619)
		hash ^= uint64(v)
	}
	return hash % total
}
func DekhashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(len(data))
	for _, v := range data {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint64(v)
	}
	return hash % total
}
func RshashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	seed := uint64(63689)
	hash := uint64(0)
	for _, v := range data {
		hash = hash*seed + uint64(v)
		seed *= uint64(378551)
	}
	return hash % total
}
func SdbmhashByte(data []byte, total uint64) uint64 {
	if total == 1 {
		return 0
	}
	hash := uint64(0)
	for _, v := range data {
		hash = uint64(v) + (hash << 6) + (hash << 16) - hash
	}
	return hash % total
}
