package models

type MulticastGroup struct {
	// It'a group ID that the multicast-group belongs to, a 5GVN group can contains many multicast-groups
	ExternalGroupId string `json:"externalGroupId"`
	// it's an unique ID among all Multicast Groups under a single 5G VN Group
	MultiGroupId        string           `json:"multiGroupId"`
	SourcePduSessIpAddr IpAddress        `json:"sourcePduSessIpAddr"`
	SourceUeGpsi        string           `json:"sourceUeGpsi"`
	GroupIpAddr         IpAddress        `json:"groupIpAddr"`
	MembersGpsi         []string         `json:"membersGpsi"`
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
