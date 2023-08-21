package model

type VkApiResponse struct {
	Response *VkApiResponseData `json:"response"`
	Error    *VkApiError        `json:"error"`
}

type VkApiError struct {
	ErrorMsg string `json:"error_msg"`
}

type VkApiResponseData struct {
	PostId int `json:"post_id"`
}

type VkApiGetPhotoServerResponse struct {
	Response *VkApiGetPhotoServerResponseData `json:"response"`
	Error    *VkApiError                      `json:"error"`
}

type VkApiGetPhotoServerResponseData struct {
	AlbumId   int    `json:"album_id"`
	UploadUrl string `json:"upload_url"`
	UserId    int    `json:"user_id"`
}

type VkApiLoadPhotoIntoServerResponse struct {
	Server *int    `json:"server"`
	Hash   *string `json:"hash"`
	Photo  *string `json:"photo"`
}

type VkApiSavePhotoResponse struct {
	Response []VkApiSavePhotoResponseData `json:"response"`
}

type VkApiSavePhotoResponseData struct {
	Id      *int `json:"id"`
	OwnerId *int `json:"owner_id"`
}
