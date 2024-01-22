package possumTorrentStorage

import (
	"github.com/anacrolix/log"
	possum "github.com/anacrolix/possum/go"
	possumResource "github.com/anacrolix/possum/go/resource"
	"github.com/anacrolix/torrent/storage"
	test_storage "github.com/anacrolix/torrent/storage/test"
	"testing"
)

func BenchmarkProvider(b *testing.B) {
	possumDir, err := possum.Open(b.TempDir())
	if err != nil {
		b.Fatal(err)
	}
	defer possumDir.Close()
	possumProvider := possumResource.Provider{Handle: possumDir}
	possumTorrentProvider := Provider{Provider: possumProvider, Logger: log.Default}
	clientStorageImpl := storage.NewResourcePiecesOpts(possumTorrentProvider, storage.ResourcePiecesOpts{})
	test_storage.BenchmarkPieceMarkComplete(b, clientStorageImpl, test_storage.DefaultPieceSize, test_storage.DefaultNumPieces, 0)
}
