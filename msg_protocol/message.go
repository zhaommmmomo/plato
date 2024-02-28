package msg_protocol

type Message struct {
	version    byte
	cmd        byte
	len        uint32
	varHeadLen uint32
	crc32Sum   uint32
}
