package handler

func int32ToInt(i32 *int32) *int {
	converted := int(*i32)

	return &converted
}
