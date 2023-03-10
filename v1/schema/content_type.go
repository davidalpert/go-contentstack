package schema

// Generated by https://quicktype.io

type LastActivity interface{}

type ContentType struct {
	Title             string               `json:"title"`
	Description       string               `json:"description"`
	Options           Options              `json:"options"`
	Schema            []Field              `json:"schema"`
	UID               string               `json:"uid"`
	DefaultACL        DefaultACL           `json:"DEFAULT_ACL"`
	SysACL            SysACL               `json:"SYS_ACL"`
	CreatedAt         string               `json:"created_at"`
	UpdatedAt         string               `json:"updated_at"`
	InbuiltClass      bool                 `json:"inbuilt_class"`
	Abilities         ContentTypeAbilities `json:"abilities"`
	LastActivity      LastActivity         `json:"last_activity"`
	MaintainRevisions bool                 `json:"maintain_revisions"`
	Version           int64                `json:"_version"`
	Url               string               `json:"url,omitempty"`
}

type ContentTypeAbilities struct {
	GetOneObject     bool `json:"get_one_object"`
	GetAllObjects    bool `json:"get_all_objects"`
	CreateObject     bool `json:"create_object"`
	UpdateObject     bool `json:"update_object"`
	DeleteObject     bool `json:"delete_object"`
	DeleteAllObjects bool `json:"delete_all_objects"`
}

type DefaultACL struct {
	Others ACL           `json:"others"`
	Users  []interface{} `json:"users"`
}

type Options struct {
	IsPage    bool     `json:"is_page"`
	Singleton bool     `json:"singleton"`
	SubTitle  []string `json:"sub_title"`
	Title     string   `json:"title"`
}

type SysACL struct {
	Others ACL         `json:"others"`
	Roles  interface{} `json:"roles"`
}

type ACL struct {
	Read    bool  `json:"read"`
	Create  *bool `json:"create,omitempty"`
	Update  *bool `json:"update,omitempty"`
	Delete  *bool `json:"delete,omitempty"`
	SubACL  *ACL  `json:"sub_acl,omitempty"`
	Publish *bool `json:"publish,omitempty"`
}
