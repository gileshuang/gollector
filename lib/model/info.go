package model

import (
	"time"
)

//BaseInfo is the base of Infomation.
type BaseInfo struct {
	ID int
}

// AtomInfo for storage infomation of each value.
type AtomInfo struct {
	UpdateTime time.Time
	Method     string
	Value      string `json:"value"`
}

// HostInfo for storage infomation of each host.
type HostInfo struct {
	BaseInfo
	HostName string               `json:"hostname"`
	Info     map[string]*AtomInfo `json:"info"`
}
