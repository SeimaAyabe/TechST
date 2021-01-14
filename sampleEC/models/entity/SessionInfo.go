package entity

// SessionInfo はセッション情報を保持する構造体
type SessionInfo struct {
	UserID         interface{}
	UserName       interface{}
	IsSessionAlive bool
}
