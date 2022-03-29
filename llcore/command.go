package llcore

type CommandType int
const (
    PING CommandType = iota
    REG
    SEND
    INFO
)

type LLProtoBase struct {
    Command uint8
}
