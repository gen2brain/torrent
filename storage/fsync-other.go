//go:build !unix

package storage

import "os"

// FlushFileBuffers on Windows requires write access.
const fsyncOpenFlag = os.O_WRONLY
