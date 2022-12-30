package models

type Task struct {
	UUID            string `json:"uuid" gorm:"primaryKey"`
	Queue           string `json:"queue"`
	CurrentState    string `json:"current_state"`
	AutoTargetState string `json:"auto_target_state"`
	SubmitTime      int64  `json:"submit_time" gorm:"autoCreateTime:nano"`
	UpdateTime      int64  `json:"update_time" gorm:"autoUpdateTime:nano"`
	Timeout         int64  `json:"timeout"`
	Priority        int64  `json:"priority"`
	Payload         []byte `json:"payload" gorm:"type:bytea"`
}

type ArchivedTask struct {
	UUID            string `json:"uuid" gorm:"primaryKey"`
	Queue           string `json:"queue"`
	CurrentState    string `json:"current_state"`
	AutoTargetState string `json:"auto_target_state"`
	SubmitTime      int64  `json:"submit_time"`
	UpdateTime      int64  `json:"update_time" gorm:"autoUpdateTime:nano"`
}

func ConvertTaskForArchive(t Task) (a ArchivedTask) {
	a.UUID = t.UUID
	a.Queue = t.Queue
	a.CurrentState = t.CurrentState
	a.AutoTargetState = t.AutoTargetState
	a.SubmitTime = t.SubmitTime
	return a
}
