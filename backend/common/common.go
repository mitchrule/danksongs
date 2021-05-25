package common

import "time"

// The length of time a cookie lasts before it expires
// NOTE: Reduce this for production (Makes it less annoying to test)
const SESSION_MINS = time.Duration(300) * time.Minute

// Essentially a representation of infinite time assigned to logout tokens
const LOGOUT_TIME = time.Duration(1000000) * time.Hour
