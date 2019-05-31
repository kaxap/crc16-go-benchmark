package crc16ccitt

// https://www.ccsinfo.com/forum/viewtopic.php?t=24977
type Crc16CCITT struct {
	crc    uint16
	crcRes []byte
}

func NewCrc16CCITT() *Crc16CCITT {
	return &Crc16CCITT{0xFFFF, make([]byte, 2, 2)}
}

func (h *Crc16CCITT) Sum(b []byte) []byte {
	h.crcRes[0] = byte(h.crc >> 8)
	h.crcRes[1] = byte(h.crc & 0xff)
	return h.crcRes
}

func (h *Crc16CCITT) Reset() {
	h.crc = 0xFFFF
}

func (h *Crc16CCITT) Size() int {
	return 2
}

func (h *Crc16CCITT) BlockSize() int {
	return 1
}

func (h *Crc16CCITT) Write(p []byte) (n int, err error) {
	/*
		// ported from this code
		int16 Calc_CRC_C(int8 *Buffer, int16 Len)
		{
		   int16 x;
		   int16 crc = 0xFFFF;

		   while(Len--)
		   {
			  x = make8(crc,1) ^ *Buffer++;
			  x ^= x>>4;

			  crc = (crc << 8) ^ (x << 12) ^ (x <<5) ^ x;
		   }
		   return crc;
		}
		} */
	oldCrc := h.crc
	var x uint16
	for _, d := range p {
		x = uint16(byte(oldCrc>>8) ^ d)
		x ^= x >> 4
		oldCrc = (oldCrc << 8) ^ (x << 12) ^ (x << 5) ^ x
	}

	h.crc = oldCrc
	return len(p), nil
}
