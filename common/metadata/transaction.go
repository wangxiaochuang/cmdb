package metadata

import "time"

type TxnOption struct {
	// transaction timeout time
	// min value: 5 * time.Second
	// default: 5min
	Timeout time.Duration
}

type TxnCapable struct {
	Timeout   time.Duration `json:"timeout"`
	SessionID string        `json:"session_id"`
}
