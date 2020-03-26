package search

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/ffmt.v1"
	"testing"
	"time"
)

type ApplicationQuery struct {
	Id      string         `search:"type:icontains;column:id" uri:"id"`
	Domain  string         `search:"type:icontains;column:domain" uri:"domain"`
	Version string         `search:"type:exact;column:version" uri:"version"`
	Status  []enums.Status `search:"type:in;column:status" uri:"status"`
	Start   time.Time      `search:"type:gte;column:created_at" uri:"start"`
	End     time.Time      `search:"type:lte;column:created_at" uri:"end"`
	ApplicationOrder
}

type ApplicationOrder struct {
	IdOrder string `search:"type:order;column:id" uri:"id_order"`
}

func TestResolveSearchQuery(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		d := ApplicationQuery{
			Id:               "aaa",
			Domain:           "bbb",
			Version:          "ccc",
			Status:           []enums.Status{1, 2},
			Start:            time.Now().Add(-8 * time.Hour),
			End:              time.Now(),
			ApplicationOrder: ApplicationOrder{IdOrder: "desc"},
		}
		condition, order := ResolveSearchQuery("mysql", d, nil, nil)
		_, _ = ffmt.P(condition, order)
	})
}
