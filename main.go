package main

import (
    "crypto/sha1"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func peerId() string {
	sid := "Taipei_tor_" + strconv.Itoa(os.Getpid()) + "______________"
	return sid[0:20]
}

func binaryToDottedPort(port string) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", port[0], port[1], port[2], port[3],
		(port[4]<<8)|port[5])
}

func doTorrent() (err os.Error) {
	log.Stderr("Fetching torrent.")
	// testBencode()
	m, err := getMetaInfo("http://releases.ubuntu.com/9.10/ubuntu-9.10-desktop-amd64.iso.torrent")
	if err != nil {
		return
	}
	log.Stderr("Tracker: ", m.Announce, " Comment: ", m.Comment, " Encoding: ", m.Encoding)
	
	fileStore, totalSize, err := NewFileStore(&m.Info)
	if err != nil {
	    return
	}
	defer fileStore.Close()
	
	log.Stderr("Computing pieces left")
	good, bad, err := checkPieces(fileStore, totalSize, m)
	log.Stderr("Good pieces: ", good, " Bad pieces: ", bad)

	si := &SessionInfo{PeerId: peerId(), Port: 6881, Left: bad * m.Info.PieceLength}

	tr, err := getTrackerInfo(m, si)
	if err != nil {
		return
	}
	
	log.Stderr("Torrent has ", tr.Complete, " seeders and ", tr.Incomplete, " leachers.")

	ip := tr.Peers[0:6]
	c, err := net.Dial("tcp", "", binaryToDottedPort(ip))
	if err != nil {
		return
	}
	var header [28]byte
	_, err = c.Read(&header)
	if err != nil {
		return
	}
	log.Stderr(header[1:20])
	log.Stderr(tr)
	return
}

func checkPieces(fs FileStore, totalLength int64, m *MetaInfo) (good, bad int64, err os.Error) {
	currentSums, err := computeSums(fs, totalLength, m.Info.PieceLength)
	if err != nil {
	    return
	}
	pieceLength := m.Info.PieceLength
    numPieces := (totalLength + pieceLength - 1) / pieceLength;
    ref := m.Info.Pieces
	for i := int64(0); i < numPieces; i++ {
	    base := i * sha1.Size
	    end := base + sha1.Size
	    if checkEqual(ref[base:end], currentSums[base:end]) {
	        good++
	    } else {
	        bad++
	    }
	}
	return
}

func checkEqual(ref string, current []byte) bool {
    for i := 0; i < len(current); i++ {
        if ref[i] != current[i] {
            return false
        }
    }
    return true
}
	
func computeSums(fs FileStore, totalLength int64, pieceLength int64) (sums []byte, err os.Error) {
    numPieces := (totalLength + pieceLength - 1) / pieceLength;
    sums = make([]byte, sha1.Size * numPieces)
    hasher := sha1.New()
    piece := make([]byte, pieceLength)
    for i := int64(0); i < numPieces; i++ {
        _, err := fs.ReadAt(piece, i * pieceLength)
        if err != nil {
            return
        }
        hasher.Reset()
        _, err = hasher.Write(piece)
        if err != nil {
            return
        }
        copy(sums[i * sha1.Size:], hasher.Sum())
    }
    return
}


func main() {
	log.Stderr("Starting.")
	err := doTorrent()
	if err != nil {
	    log.Stderr("Failed: ", err)
	} else {
	    log.Stderr("Done")
	}
}

