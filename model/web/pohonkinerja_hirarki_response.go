package web

type PohonKinerjaHierarchyResponse struct {
	Strategic   []StrategicResponse   `json:"strategic"`
	Tactical    []TacticalResponse    `json:"tactical"`
	Operational []OperationalResponse `json:"operational"`
}

type StrategicResponse struct {
	ID        int    `json:"id"`
	Level     int    `json:"level"`
	NamaPohon string `json:"nama_pohon"`
}

type TacticalResponse struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	Level     int    `json:"level"`
	NamaPohon string `json:"nama_pohon"`
}

type OperationalResponse struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	Level     int    `json:"level"`
	NamaPohon string `json:"nama_pohon"`
}
