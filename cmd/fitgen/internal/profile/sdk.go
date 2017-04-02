package profile

var knownMesgNumButNoMsgPerSDK = map[string]map[string]bool{
	"16.20": {
		"GpsMetadata": true,
		"Pad":         true,
	},
	"20.14": {
		"Pad": true,
	},
	"20.27": {
		"Pad": true,
	},
}

func knownMesgNumButNoMsg(sdk, mesgNum string) bool {
	const fallbackSDK = "16.20"
	mnMap, found := knownMesgNumButNoMsgPerSDK[sdk]
	if !found {
		return knownMesgNumButNoMsgPerSDK[fallbackSDK][mesgNum]
	}
	return mnMap[mesgNum]
}
