package models

type EventType int32

type EventHandler func(interface{}) error
