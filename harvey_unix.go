// +build !plan9

package harvey

//go:generate stringer -type MountFlag

// MountFlag are flags for the mount syscall
type MountFlag int

const (
	// REPL Replace the old file by the new one.
	// Henceforth, an evaluation of old will be translated to the new file.
	// If they are directories (for mount, this condition is true by definition),
	// old becomes a union directory consisting of one directory (the new file).
	REPL MountFlag = 0x0000
	// BEFORE Both the old and new files must be directories.
	// Add the constituent files of the new directory to the
	// union directory at old so its contents appear first in the union.
	// After an BEFORE bind or mount, the new directory will be
	// searched first when evaluating file names in the union directory.
	BEFORE MountFlag = 0x0001
	// AFTER Like MBEFORE but the new directory goes at the end of the union.
	AFTER MountFlag = 0x0002
	// CREATE flag that can be OR'd with any of the above.
	// When a create system call (see open(2)) attempts to create in a union directory,
	// and the file does not exist, the elements of the union are searched in order until
	// one is found with CREATE set. The file is created in that directory;
	// if that attempt fails, the create fails.
	CREATE MountFlag = 0x0004
	// CACHE flag, valid for mount only, turns on caching for files made available by the mount.
	// By default, file contents are always retrieved from the server.
	// With caching enabled, the kernel may instead use a local cache
	// to satisfy read(5) requests for files accessible through this mount point.
	CACHE MountFlag = 0x0010
)

type Perm uint32

// Size is the type we use to encode 9P2000
type Size uint32

// Tag is the id field used to identify Transmit Messgaes
type Tag uint16

// FID is the id of the current file.
type FID uint32

const (
	// MSize is default message size (1048576+IOHdrSz)
	MSize = 2*1048576 + IOHDRSZ
	// IOHDRSZ non-data size of the Twrite messages
	IOHDRSZ = 24
	// DefaultVersion is the 9pversion
	DefaultVersion = "9P2000"

	NOTAG = 0xffff
	NOFID = 0xffffffff
	NOUID = 0xffffffff
)

// Error values
const (
	EPERM   = 1
	ENOENT  = 2
	EIO     = 5
	EACCES  = 13
	EEXIST  = 17
	ENOTDIR = 20
	EINVAL  = 22
)

// QID types
const (
	QTDIR     = 0x80 // directories
	QTAPPEND  = 0x40 // append only files
	QTEXCL    = 0x20 // exclusive use files
	QTMOUNT   = 0x10 // mounted channel
	QTAUTH    = 0x08 // authentication file
	QTTMP     = 0x04 // non-backed-up file
	QTSYMLINK = 0x02 // symbolic link (Unix, 9P2000.u)
	QTLINK    = 0x01 // hard link (Unix, 9P2000.u)
	QTFILE    = 0x00
)

// Flags for the mode field in Topen and Tcreate messages
const (
	OREAD   = 0x0    // open read-only
	OWRITE  = 0x1    // open write-only
	ORDWR   = 0x2    // open read-write
	OEXEC   = 0x3    // execute (== read but check execute permission)
	OTRUNC  = 0x10   // or'ed in (except for exec), truncate file first
	OCEXEC  = 0x20   // or'ed in, close on exec
	ORCLOSE = 0x40   // or'ed in, remove on close
	OAPPEND = 0x80   // or'ed in, append only
	OEXCL   = 0x1000 // or'ed in, exclusive client use
)

// File modes
const (
	DMDIR    = 0x80000000 // mode bit for directories
	DMAPPEND = 0x40000000 // mode bit for append only files
	DMEXCL   = 0x20000000 // mode bit for exclusive use files
	DMMOUNT  = 0x10000000 // mode bit for mounted channel
	DMAUTH   = 0x08000000 // mode bit for authentication file
	DMTMP    = 0x04000000 // mode bit for non-backed-up file
	DMREAD   = 0x4        // mode bit for read permission
	DMWRITE  = 0x2        // mode bit for write permission
	DMEXEC   = 0x1        // mode bit for execute permission
)

// A QID represents a 9P server's unique identification for a file.
type QID struct {
	Path uint64 // the file server's unique identification for the file
	Vers uint32 // version number for given Path
	Type uint8  // the type of the file (syscall.QTDIR for example)
}

// A Dir contains the metadata for a file.
type Dir struct {
	// system-modified data
	Type uint16 // server type
	Dev  uint32 // server subtype

	// file data
	QID    QID    // unique id from server
	Mode   Perm   // permissions
	Atime  uint32 // last read time
	Mtime  uint32 // last write time
	Length uint64 // file length
	Name   string // last element of path
	UID    string // owner name
	GID    string // group name
	Muid   string // last modifier name
}
