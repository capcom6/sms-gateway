package version

import "strconv"

const notSet string = "not set"

// these information will be collected when build, by `-ldflags "-X main.appVersion=0.1"`
var (
	AppVersion = notSet
	AppRelease = notSet
)

func AppReleaseID() int {
	id, _ := strconv.Atoi(AppRelease)

	return id
}
