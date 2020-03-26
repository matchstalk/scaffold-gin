package dto

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	"time"
)

type ApplicationQuery struct {
	Id      string         `search:"type:icontains;column:id" uri:"id"`
	Domain  string         `search:"type:icontains;column:domain" uri:"domain"`
	Version string         `search:"type:exact;column:version" uri:"version"`
	Status  []enums.Status `search:"type:in;column:status" uri:"status"`
	Start   time.Time      `search:"type:istartswith;column:created_at" uri:"start"`
	End     time.Time      `search:"type:iendswith;column:created_at" uri:"end"`
	ApplicationOrder
}

type ApplicationOrder struct {
	IdOrder string `search:"type:order;column:id" uri:"id_order"`
}
