package schema

// File is one of the predefined ContentTypes
type File struct {
	SchemaVersion  int64                  `json:"_version"`
	ACL            map[string]interface{} `json:"acl"`
	ContentType    string                 `json:"content_type"`
	CreatedAt      string                 `json:"created_at,omitempty"`
	CreatedBy      string                 `json:"created_by,omitempty"`
	FileSize       string                 `json:"file_size,omitempty"`
	Filename       string                 `json:"filename,omitempty"`
	IsDir          *bool                  `json:"is_dir,omitempty"`
	ParentUID      string                 `json:"parent_uid,omitempty"`
	PublishDetails *FilePublishDetails    `json:"publish_details,omitempty"`
	Tags           []string               `json:"tags"`
	Title          string                 `json:"title"`
	UID            string                 `json:"uid"`
	Url            string                 `json:"url"`
	UpdatedAt      string                 `json:"updated_at,omitempty"`
	UpdatedBy      string                 `json:"updated_by,omitempty"`
}

type FilePublishDetails struct {
	Environment string `json:"environment,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Time        string `json:"time,omitempty"`
	User        string `json:"user,omitempty"`
}
