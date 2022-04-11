package config

type Server struct {
	Mongo      Mongo      `json:"mongo"`
	TencentCOS TencentCOS `json:"tencent-cos"`
}
