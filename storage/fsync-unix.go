//go:build unix

package storage

import "os"

// fsync(2) works on a read-only descriptor, and needs one for a file promoted off its part file name,
// which is made read-only.
const fsyncOpenFlag = os.O_RDONLY
