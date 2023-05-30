package schema

type Locale struct {
	Code      string `json:"code"`
	Uid       string `json:"uid"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	//ACL       struct {} `json:"ACL"`
	Version            int    `json:"_version"`
	FallbackLocaleCode string `json:"fallback_locale"`
}
