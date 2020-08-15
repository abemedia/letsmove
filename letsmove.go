// Package letsmove moves the running MacOS application to the Applications folder.
package letsmove

// Moves the running application to ~/Applications or /Applications if the former does not exist.
// After the move, it relaunches app from the new location.
// DOES NOT work for sandboxed applications.
func MoveToApplications() {
	moveToApplications()
}

// Check whether an app move is currently in progress.
// Returns YES if LetsMove is currently in-progress trying to move the app to the Applications folder, or NO otherwise.
// This can be used to work around a crash with apps that terminate after last window is closed.
// See https://github.com/potionfactory/LetsMove/issues/64 for details.
func IsMoveInProgress() bool {
	return isMoveInProgress()
}
