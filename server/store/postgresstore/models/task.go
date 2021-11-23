package models

type Task struct {
	UUID 			string `json:"uuid" gorm:"primaryKey"`
	Queue           string `json:"queue"`
	CurrentState    string `json:"current_state"`
	AutoTargetState string `json:"auto_target_state"`
	SubmitTime 		int64 `json:"submit_time" gorm:"autoCreateTime:nano"`
	UpdateTime 		int64 `json:"update_time" gorm:"autoUpdateTime:nano"`
	Timeout    		int64 `json:"timeout"`
	Payload 		[]byte `json:"payload" gorm:"type:bytea"`
}
