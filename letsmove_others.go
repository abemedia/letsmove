// +build !darwin

package letsmove

func moveToApplications() {
	// no-op on all but darwin
}

func isMoveInProgress() bool {
	// no-op on all but darwin
	return false
}
