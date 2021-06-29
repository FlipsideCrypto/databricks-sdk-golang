package models

type DockerImage struct {
	URL       string     `json:"url,omitempty" url:"url,omitempty"`
	BasicAuth *BasicAuth `json:"basic_auth,omitempty" url:"basic_auth,omitempty"`
}

type BasicAuth struct {
	Username string `json:"username,omitempty" url:"username,omitempty"`
	Password string `json:"password,omitempty" url:"password,omitempty"`
}
