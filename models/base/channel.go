package base

const (
	CHANNEL_CALL_STATE_DOWN = iota
	CHANNEL_CALL_STATE_DIALING
	CHANNEL_CALL_STATE_RINGING
	CHANNEL_CALL_STATE_EARLY
	CHANNEL_CALL_STATE_ACTIVE
	CHANNEL_CALL_STATE_HELD
	CHANNEL_CALL_STATE_RING_WAIT
	CHANNEL_CALL_STATE_HANGUP
	CHANNEL_CALL_STATE_UNHELD
)

const (
	CHANNEL_STATE_NEW = iota
	CHANNEL_STATE_INIT
	CHANNEL_STATE_ROUTING
	CHANNEL_STATE_SOFT_EXECUTE
	CHANNEL_STATE_EXECUTE
	CHANNEL_STATE_EXCHANGE_MEDIA
	CHANNEL_STATE_PARK
	CHANNEL_STATE_CONSUME_MEDIA
	CHANNEL_STATE_HIBERNATE
	CHANNEL_STATE_RESET
	CHANNEL_STATE_HANGUP
	CHANNEL_STATE_REPORTING
	CHANNEL_STATE_DESTROY
	CHANNEL_STATE_NONE
)

type Channel struct {
	UUID             string `json:"uuid"`
	ChannelState     int    `json:"channel_state"`
	ChannelCallState int    `json:"channel_call_state"`
}

func NewChannel(uuid string, channelState, channelCallState int) *Channel {
	return &Channel{uuid, channelState, channelCallState}
}
