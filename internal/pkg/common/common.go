package common

func GetLimitAndOffset(page, size int32) (int32, int32) {
	var limit int32 = 10
	if size > 0 {
		limit = size
	}
	offset := (page - 1) * limit

	return limit, offset
}
