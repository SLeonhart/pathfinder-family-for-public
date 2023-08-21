package utils

import "pathfinder-family/model"

func PaginatePushToken(x []model.PushToken, pageNum int, pageSize int) []model.PushToken {
	start := pageNum * pageSize
	sliceLength := len(x)

	if start > sliceLength {
		return make([]model.PushToken, 0)
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return x[start:end]
}
