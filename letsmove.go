//go:build darwin
// +build darwin

// Package letsmove moves the running MacOS application to the Applications folder.
package letsmove

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -framework Foundation -framework AppKit -framework Security

void PFMoveToApplicationsFolderIfNecessary(void);
int PFMoveIsInProgress(void);
*/
import "C"

// MoveToApplications moves the running application to ~/Applications or /Applications if the former does not exist.
// After the move, it relaunches app from the new location.
// DOES NOT work for sandboxed applications.
func MoveToApplications() {
	C.PFMoveToApplicationsFolderIfNecessary()
}

// IsMoveInProgress checks whether an app move is currently in progress.
// Returns YES if LetsMove is currently in-progress trying to move the app to the Applications folder, or NO otherwise.
// This can be used to work around a crash with apps that terminate after last window is closed.
// See https://github.com/potionfactory/LetsMove/issues/64 for details.
func IsMoveInProgress() bool {
	return C.PFMoveIsInProgress() == 1
}
