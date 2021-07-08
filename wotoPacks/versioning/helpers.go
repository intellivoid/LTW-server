package versioning

import "strings"

func VersionAcceptable(verStr, verHash string) bool {
	verStr = strings.ToLower(verStr)
	verHash = strings.ToLower(verHash)
	return verStr == currentVersion &&
		verHash == versionHash
}
