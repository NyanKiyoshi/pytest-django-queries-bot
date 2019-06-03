package consts

// MaxUploadSize defines the max rapport size that one can upload.
// The value is in bytes which is calculated to `{x} times 1 MiB`.
const MaxUploadSize = 2 * (1 << 20)

// CommitHashHeaderName defines the header from which we read
// the SHA1 commit hash.
const CommitHashHeaderName = "X-Commit-Rev"
const CommitHashHeaderNameLower = "x-commit-rev"

// SHA1Length is the length of a SHA1 commit hash.
const SHA1Length = 40
