package oss

func getClient() *OssClient {
	c := NewOssClient(&Config{
		AK:     "kbHkkh2NjA_54eVMqt20PRlbA0SF8DozJF7yPWQR",
		SK:     "AVWVC4Mg42mNElszK7FO4CAtiY8b8BcZEJykvFec",
		Bucket: "live-video-1024",
		domain: "cdn-host.chajiuqqq.cn",
	})
	return c
}

// func TestOssClient_FileUpload(t *testing.T) {
// 	c := getClient()
// 	key, err := c.FileUpload("/Users/chajiu/Downloads/视频-盛邦.mp4", "视频-盛邦2.mp4")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "视频-盛邦2.mp4", key)
// }

// func TestOssClient_ResourceUrl(t *testing.T) {
// 	c := getClient()
// 	url := c.ResourceUrl("视频-盛邦2.mp4")
// 	assert.Equal(t, "cdn-host.chajiuqqq.cn/%E8%A7%86%E9%A2%91-%E7%9B%9B%E9%82%A62.mp4", url)
// }

// func TestOssClient_GetResource(t *testing.T) {
// 	c := getClient()
// 	bytes, err := c.GetResource("视频-盛邦2.mp4")
// 	assert.NoError(t, err)
// 	err = os.WriteFile("myfile.mp4", bytes, 0644)
// 	assert.NoError(t, err)
// }
