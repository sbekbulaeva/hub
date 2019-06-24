package api

import "time"

type CloudAccount struct {
	Id               string
	Name             string
	Kind             string
	Status           string
	BaseDomain       string `json:"baseDomain"`
	Parameters       []Parameter
	TeamsPermissions []Team `json:"teamsPermissions"`
}

type AwsSecurityCredentials struct {
	Cloud        string
	AccessKey    string
	SecretKey    string
	SessionToken string
	Ttl          int
}

type CloudAccountRequest struct {
	Name        string            `json:"name"`
	Kind        string            `json:"kind"`
	Parameters  []Parameter       `json:"parameters,omitempty"`
	Credentials map[string]string `json:"credentials,omitempty"`
}

type Parameter struct {
	Name      string      `json:"name"`
	Kind      string      `json:"kind,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	From      string      `json:"from,omitempty"`
	Component string      `json:"component,omitempty"`
	Origin    string      `json:"origin,omitempty"`
	Messenger string      `json:"messenger,omitempty"`
}

type Secret struct {
	Name   string
	Kind   string
	Values map[string]string
}

type Output struct {
	Name      string      `json:"name"`
	Component string      `json:"component,omitempty"`
	Kind      string      `json:"kind,omitempty"`
	Value     interface{} `json:"value"`
	Brief     string      `json:"brief,omitempty"`
	Messenger string      `json:"messenger,omitempty"`
}

type Provider struct {
	Kind       string      `json:"kind"`
	Name       string      `json:"name"`
	Provides   []string    `json:"provides,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Environment struct {
	Id               string
	Name             string
	Description      string
	Tags             []string
	CloudAccount     string `json:"cloudAccount"`
	Parameters       []Parameter
	Providers        []Provider
	TeamsPermissions []Team `json:"teamsPermissions"`
}

type EnvironmentRequest struct {
	Name         string      `json:"name"`
	CloudAccount string      `json:"cloudAccount"`
	Parameters   []Parameter `json:"parameters"` // TODO omitempty as soon as API is ready
	Providers    []Provider  `json:"providers"`
}

type StackComponent struct {
	Name        string
	Brief       string
	Description string
}

type BaseStack struct {
	Id         string
	Name       string
	Brief      string
	Tags       []string
	Components []StackComponent
	Parameters []Parameter
}

type StackRef struct {
	Id   string
	Name string
}

type GitRef struct {
	Public   string
	Template struct {
		Ref string
	}
	K8s struct {
		Ref string
	}
}

type StackTemplate struct {
	Id                string
	Name              string
	Description       string
	Status            string
	Tags              []string
	Stack             StackRef
	ComponentsEnabled []string `json:"componentsEnabled"`
	Verbs             []string
	GitRemote         GitRef `json:"gitRemote"`
	Parameters        []Parameter
	TeamsPermissions  []Team `json:"teamsPermissions"`
}

type StackTemplateRequest struct {
	Name              string      `json:"name"`
	Description       string      `json:"description,omitempty"`
	Tags              []string    `json:"tags,omitempty"`
	Stack             string      `json:"stack"`
	ComponentsEnabled []string    `json:"componentsEnabled,omitempty"`
	Verbs             []string    `json:"verbs,omitempty"`
	Parameters        []Parameter `json:"parameters,omitempty"`
	TeamsPermissions  []Team      `json:"teamsPermissions,omitempty"`
}

type EnvironmentRef struct {
	Id     string
	Name   string
	Domain string
}

type StackTemplateRef struct {
	Id   string
	Name string
}

type PlatformRef struct {
	Id     string
	Name   string
	Domain string
}

type ComponentStatus struct {
	Name    string            `json:"name"`
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Outputs map[string]string `json:"outputs,omitempty"`
}

type LifecyclePhase struct {
	Phase  string `json:"phase"`
	Status string `json:"status"`
}

type InflightOperation struct {
	Id          string                 `json:"id"`
	Operation   string                 `json:"operation"`
	Timestamp   time.Time              `json:"timestamp"`
	Status      string                 `json:"status,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"`
	Description string                 `json:"description,omitempty"`
	Initiator   string                 `json:"initiator,omitempty"`
	Logs        string                 `json:"logs,omitempty"`
	Phases      []LifecyclePhase       `json:"phases,omitempty"`
}

type TemplateStatus struct {
	Commit  string `json:"commit,omitempty"`
	Ref     string `json:"ref,omitempty"`
	Date    string `json:"date,omitempty"`
	Author  string `json:"author,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type StackInstanceStatus struct {
	Status     string            `json:"status,omitempty"`
	Template   *TemplateStatus   `json:"template,omitempty"`
	K8s        *TemplateStatus   `json:"k8s,omitempty"`
	Components []ComponentStatus `json:"components,omitempty"`
}

type StackInstance struct {
	Id                 string
	Name               string
	Domain             string
	Description        string
	Tags               []string
	Environment        EnvironmentRef
	Stack              StackRef
	Template           StackTemplateRef
	Platform           PlatformRef
	ComponentsEnabled  []string `json:"componentsEnabled"`
	GitRemote          GitRef   `json:"gitRemote"`
	Parameters         []Parameter
	Outputs            []Output
	Provides           map[string][]string
	StateFiles         []string `json:"stateFiles"`
	Status             StackInstanceStatus
	InflightOperations []InflightOperation `json:"inflightOperations,omitempty"`
}

type StackInstanceRequest struct {
	Name              string      `json:"name"`
	Description       string      `json:"description,omitempty"`
	Tags              []string    `json:"tags,omitempty"`
	Environment       string      `json:"environment"`
	Template          string      `json:"template"`
	Platform          string      `json:"platform,omitempty"`
	ComponentsEnabled []string    `json:"componentsEnabled,omitempty"`
	Parameters        []Parameter `json:"parameters,omitempty"`
}

type StackInstanceDeployResponse struct {
	JobId string `json:jobId`
}

type StackInstancePatch struct {
	ComponentsEnabled  []string             `json:"componentsEnabled,omitempty"`
	Parameters         []Parameter          `json:"parameters,omitempty"`
	StateFiles         []string             `json:"stateFiles,omitempty"`
	Status             *StackInstanceStatus `json:"status,omitempty"`
	InflightOperations []InflightOperation  `json:"inflightOperations,omitempty"`
	Outputs            []Output             `json:"outputs,omitempty"`
	Provides           map[string][]string  `json:"provides,omitempty"`
}

type Application struct {
	Id               string
	Name             string
	Description      string
	Tags             []string
	Environments     []EnvironmentRef
	Parameters       []Parameter
	GitRemote        GitRef `json:"gitRemote"`
	TeamsPermissions []Team `json:"teamsPermissions"`
}

type License struct {
	Component  string
	LicenseKey string
}

type ServiceAccount struct {
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	GroupId    string `json:"groupId"`
	LoginToken string `json:"loginToken"`
}

type DeploymentKey struct {
	DeploymentKey string `json:"deploymentKey"`
}
