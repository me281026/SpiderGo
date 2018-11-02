package uuid

import "github.com/satori/go.uuid"

func GetUUID() string {
	uuids := uuid.Must(uuid.NewV4())
	return uuids.String()
}
