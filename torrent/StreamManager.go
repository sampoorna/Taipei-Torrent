package torrent

import (
	"fmt"
)

type StreamManager struct {

var readylist = []peerState{}
var activelist = []peerState{}


//TODO : Change Active Piece to set of Active pieces
int activepieceindex=0// do pieces start with 1 or more

int failpeer=0;

int threshold //lets set this to 3/4 of the total peers

}

func InitializeStreamManager(int thresh){
	threshold=thresh
	readylist=readylist[0:0]
	activelist=activelist[0:0]
}


func (*t TorrentSession) RequestBlock(p *peerState, piecereturning bool){
	
	if !t.si.HaveTorrent { // We can't request a block without a torrent
		return
	}

	AddtoReadyList(p) //It Either has come for the first time and returning after providing a piece

	if (piecereturning){ //If the peer returned a PIECE, then we delete it from active list (its already in the Ready List now)
		delete(activelist, p)
	}
	
	CycleReadyList()
	
	}



}

func (*t TorrentSession) CycleReadyList (){
	for i,element := range readylist{
		if (!t.pieceSet.IsSet(activepieceindex)) && readylist[i].have.IsSet(activepieceindex) {
			//We have found the piece we are looking for 
			RemoveFromReadyList(readylist[i]) //now we are removing chosen peer from ready list to being download
			AddtoActiveList(readylist[i])  // we will add the peer to activelist while downloading
			//REQUEST THE BLOCK
			activepieceindex++ // we are interested in the next piece
			failpeer=0; // fail peer is reset for this piece
			CycleReadyList() // recusrively call the Cycle function to start searching for the next piece
			// Note this time, the peer that was previosly in the ready list will not be there since we removed him
		}
		else {
			//we did not find the piece we are looking for 
			failpeer++ // we increase fail peer counter because this was a peer that did not have the needed PIECE
			if (failpeer>threshold){
				// CALL A FRESH PEER FUNCTION
			}
		}
	}


}




func AddPeertoActiveList(p *peerState){
	activelist=append(activelist,p)
}

func AddtoReadyList(p *peerState){
	readylist=append(readylist,p)
}

func RemoveFromActiveList(p *peerState){
	for i,element := range activelist{
		if(activelist[i].id==p.id){
			activelist = append(activelist[:i], activelist[i+1:]...)
		}
	}
}

func RemoveFromReadyList(p *peerState){
	for i,element := range readylist{
		if(readylist[i].id==p.id){
			readylist = append(readylist[:i], readylist[i+1:]...)
		}
	}
}