package msg

import (
	"github.com/carsonsx/net4g"
)

var Serializer = net4g.NewJsonSerializer()


func jsonByString() {

	Serializer.DeserializeId(new(UserLogin))
	Serializer.SerializeId(new(UserLoginReply), "user_login_reply")
	Serializer.SerializeId(new(UserOnline))
	Serializer.SerializeId(new(UserOffline))
	Serializer.DeserializeId(new(ChangeName))
}

func InitSerializer() {
	Serializer.DeserializeId(new(UserLogin))
	Serializer.SerializeId(new(UserLoginReply))
	Serializer.SerializeId(new(UserOnline))
	Serializer.SerializeId(new(UserOffline))
	Serializer.DeserializeId(new(ChangeName))
}
