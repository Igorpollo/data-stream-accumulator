package models

import (
	"github.com/google/uuid"
	"net"






	

	"time"
)

const (
	DTYPE_JSON = iota
	DTYPE_STRING
	DTYPE_ARRAY_STRING
	DTYPE_IMAGE
	DTYPE_AUDIO
	DTYPE_OTHER
)


type DataPackage struct {
	UUID             uuid.UUID `json:"uuid"`
	Label            string    `json:"label"`
	DateTimeReceived time.Time `json:"date_recieved"`
	DataSize         int       `json:"data_size"`
	DataType         int       `json:"data_type"`
	Data             string    `json:"data"`
	OwnerID          int       `json:"owner_id"`
	IP               net.IP    `json:"ip"`
}
