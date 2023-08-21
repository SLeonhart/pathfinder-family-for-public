package config

type VkAPI struct {
	Send           string `env:"VK_API_SEND"`
	GetPhotoServer string `env:"VK_API_GET_PHOTO_SERVER"`
	SavePhoto      string `env:"VK_API_SAVE_PHOTO"`
	AccessToken    string `env:"VK_API_ACCESS_TOKEN"`
	GroupId        int64  `env:"VK_API_GROUP_ID"`
	TestGroupId    int64  `env:"VK_API_TEST_GROUP_ID"`
}
