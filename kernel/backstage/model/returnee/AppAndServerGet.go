package returnee

/**
 * @author eyesYeager
 * @date 2023/4/13 17:40
 */

type AppAndServerGet struct {
	Id          uint          `json:"id"`          // 应用id
	App         string        `json:"app"`         // 应用地址
	Pattern     int8          `json:"pattern"`     // 模式
	Status      int8          `json:"status"`      // 状态
	Remarks     string        `json:"remarks"`     // 备注
	CreateTime  string        `json:"create_time"` // 创建时间
	ServerSlice []ServerSlice `json:"serverSlice"` // 服务数组
}

type ServerSlice struct {
	Id         uint   `json:"id"`          // 服务id
	Server     string `json:"server"`      // 服务地址
	Status     uint8  `json:"status"`      // 服务状态
	Remarks    string `json:"remarks"`     // 备注
	CreateTime string `json:"create_time"` // 创建时间
}
