package main

// Decoder for goim packet
// https://github.com/Terry-Mao/goim/blob/e742c99ad76e626d5f6df8b33bc47ca005501980/api/protocol/protocol.go#L25

const (
	// MaxBodySize max body size
	MaxBodySize = int32(1 << 12) // 4096
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

type Decoder struct {
	pack    []byte
	header  []byte
	version []byte
	seq     []byte
	heart   []byte
}

func Decode(buf []byte) *Decoder {
	decoder := &Decoder{}
	// Todo
	return decoder
}

func (d *Decoder) Pack() string {
	return string(d.pack)
}
