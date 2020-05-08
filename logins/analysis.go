package logins

// Analysis represents a resulting analysis of a singular login event. It
// contains the current geo locational resolution to the IP address of a login
// event, comparative details of the preceding login event if it exists,
// comparative details of the subsequent login event if it exists, and whether
// or not the comparison of the events result in suspicious activity.
type Analysis struct {
	CurrentLocation             *Location `json:"currentGeo"`
	PrecedingAccess             *IPAccess `json:"precedingIpAccess,omitempty"`
	SuspiciousPrecedingAccess   bool      `json:"travelToCurrentGeoSuspicious"`
	SubsequentAccess            *IPAccess `json:"subsequentIpAccess,omitempty"`
	SuspiciiousSubsequentAccess bool      `json:"travelFromCurrentGeoSuspicious"`
}
