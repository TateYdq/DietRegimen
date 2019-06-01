package disease


type DiseaseInfo struct{
	DiseaseID int `json:"disease_id"`
	Name string `json:"name"`
	DiseaseKind string `json:"disease_kind"`
	Info string `json:"info"`
	PhotoPath string `json:"photo_path"`
	VoicePath string `json:"voice_path"`
	ViewCount int `json:"view_count"`
	CollectCount string `json:"collect_count"`
}