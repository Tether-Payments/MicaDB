package micadb

const defaultDBName = "micadb.bin"

// Options holds the settings used when creating a new MicaDB
type Options struct {
	// Filename is what the file on disk will be named. The default is micadb.bin
	Filename string

	// IsTest when enabled, the file will be created in the temp directory and removed afterwards
	IsTest bool

	// BackupFrequency is how often in seconds the in-memory data store should be persisted to disk
	// 0 (Default) will disable backups
	BackupFrequency int
}
