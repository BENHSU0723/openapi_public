package models

type MulticastGroup struct {
	// it's unique among all MulticastGroups under a single 5G VN Group
	MultiGroupId        string           `json:"multiGroupId"`
	SourcePduSessIpAddr IpAddress        `json:"sourcePduSessIpAddr"`
	GroupIpAddr         IpAddress        `json:"groupIpAddr"`
	GroupServiceType    GroupServiceType `json:"groupServiceType"`
	QueryType           QueryType        `json:"queryType,omitempty"`
}

type QueryType string

const (
	QueryType_General                QueryType = "General"
	QueryType_GroupSpecific          QueryType = "GroupSpecific"
	QueryType_GroupAndSourceSpecific QueryType = "GroupAndSourceSpecific"
)

type GroupServiceType string

const (
	GroupServiceType_Streaming    GroupServiceType = "Streaming"
	GroupServiceType_TextMessage  GroupServiceType = "TextMessage"
	GroupServiceType_FileTransfer GroupServiceType = "FileTransfer"
)
