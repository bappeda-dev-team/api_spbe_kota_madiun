package web

type PohonKinerjaHierarchyResponse struct {
	Strategic   []StrategicResponse   `json:"Strategic"`
	Tactical    []TacticalResponse    `json:"Tactical"`
	Operational []OperationalResponse `json:"Operational"`
}

type StrategicResponse struct {
	IDStrategic int    `json:"id_strategic"`
	Level       int    `json:"level"`
	NamaPohon   string `json:"nama_pohon"`
}

type TacticalResponse struct {
	IDTactical int    `json:"id_tactical"`
	Parent     int    `json:"parent"`
	Level      int    `json:"level"`
	NamaPohon  string `json:"nama_pohon"`
}

type OperationalResponse struct {
	IDOperational int    `json:"id_operational"`
	Parent        int    `json:"parent"`
	Level         int    `json:"level"`
	NamaPohon     string `json:"nama_pohon"`
}
