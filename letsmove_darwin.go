package letsmove

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -framework Foundation -framework AppKit -framework Security

void PFMoveToApplicationsFolderIfNecessary(void);
int PFMoveIsInProgress(void);
*/
import "C"

import "runtime"

func moveToApplications() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	C.PFMoveToApplicationsFolderIfNecessary()
}

func isMoveInProgress() bool {
	return C.PFMoveIsInProgress() == 1
}
