package crc16ccitt

import "github.com/sigurn/crc16"

type sgWrapper struct {
	crc    uint16
	crcRes []byte
	table  *crc16.Table
}

func newsgWrapper(params crc16.Params) *sgWrapper {
	table := crc16.MakeTable(params)
	return &sgWrapper{crc16.Init(table), make([]byte, 2, 2), table}
}

func (h *sgWrapper) Sum(b []byte) []byte {
	h.crc = crc16.Complete(h.crc, h.table)
	h.crcRes[0] = byte(h.crc >> 8)
	h.crcRes[1] = byte(h.crc & 0xff)
	return h.crcRes
}

func (h *sgWrapper) Reset() {
	h.crc = crc16.Init(h.table)
}

func (h *sgWrapper) Size() int {
	return 2
}

func (h *sgWrapper) BlockSize() int {
	return 1
}

func (h *sgWrapper) Write(p []byte) (n int, err error) {

	h.crc = crc16.Update(h.crc, p, h.table)
	return len(p), nil
}
