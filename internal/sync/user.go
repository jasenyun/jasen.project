package sync

import (
	"log"
	"time"
)

type SyncHandler struct {
}

func (SyncHandler) SyncUser() {
	log.Printf("user--111 time = %d\n", time.Now().Unix())
}

func (SyncHandler) SyncDept() {
	log.Printf("dept--222")
}
