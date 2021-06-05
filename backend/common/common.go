package common

// Can use this for global variables or functions used throughout

import "time"

// The length of time a cookie lasts before it expires
// NOTE: Reduce this for production (Makes it less annoying to test)
const SESSION_MINS = time.Duration(30000) * time.Minute

// Essentially a representation of infinite time assigned to logout tokens
const LOGOUT_TIME = time.Duration(1000000) * time.Hour
