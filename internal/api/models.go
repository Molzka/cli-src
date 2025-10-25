package api

type PullRequest struct {
	ID           string            `json:"id"`
	Slug         string            `json:"slug"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Status       PRStatus          `json:"status"`
	SourceBranch string            `json:"source_branch"`
	TargetBranch string            `json:"target_branch"`
	Author       map[string]string `json:"author"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	UpdatedBy    map[string]string `json:"updated_by"`
	Repository   map[string]string `json:"repository"`
	MergeInfo    *MergeInfo        `json:"merge_info"`
}

type UserEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type RepoEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type MergeInfo struct {
	Merger           UserEmbedded    `json:"merger"`
	MergeParameters  MergeParameters `json:"merge_parameters"`
	TargetCommitHash string          `json:"target_commit_hash"`
	Error            string          `json:"error"`
	MergeCommitHash  string          `json:"merge_commit_hash"`
}

type MergeParameters struct {
	Rebase       bool `json:"rebase"`
	Squash       bool `json:"squash"`
	DeleteBranch bool `json:"delete_branch"`
}

type PRStatus string

const (
	PRStatusDraft     PRStatus = "draft"
	PRStatusOpen      PRStatus = "open"
	PRStatusMerging   PRStatus = "merging"
	PRStatusMerged    PRStatus = "merged"
	PRStatusDiscarded PRStatus = "discarded"
)

type UpdatePullRequestRequest struct {
	Status      PRStatus `json:"status,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
}

type ListRepositoriesResponse struct {
	Repositories  []Repository `json:"repositories"`
	NextPageToken string       `json:"next_page_token,omitempty"`
}

type Repository struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	Description   string         `json:"description"`
	Visibility    string         `json:"visibility"`
	DefaultBranch string         `json:"default_branch"`
	Organization  Organization   `json:"organization"`
	CloneURL      CloneURL       `json:"clone_url"`
	Counters      RepoCounters   `json:"counters"`
	LastUpdated   string         `json:"last_updated"`
	Language      *Language      `json:"language,omitempty"`
	Parent        *RepositoryRef `json:"parent,omitempty"`
}

type Organization struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type CloneURL struct {
	HTTPS string `json:"https"`
	SSH   string `json:"ssh"`
}

type RepoCounters struct {
	Forks        string `json:"forks"`
	PullRequests string `json:"pull_requests"`
	Issues       string `json:"issues"`
	Tags         string `json:"tags"`
	Branches     string `json:"branches"`
}

type Language struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type RepositoryRef struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type InitRepoSettings struct {
	DefaultBranch       string   `json:"default_branch,omitempty"`
	CreateReadme        bool     `json:"create_readme"`
	GitignorePresets    []string `json:"gitignore_presets,omitempty"`
	LicenseSlug         string   `json:"license_slug,omitempty"`
	SrcYamlTemplateSlug string   `json:"src_yaml_template_slug,omitempty"`
}

type TemplatingOptions struct {
	TemplateID string `json:"template_id"`
}

type CreateRepositoryBody struct {
	Name              string             `json:"name"`
	Slug              string             `json:"slug"`
	Description       string             `json:"description,omitempty"`
	Visibility        string             `json:"visibility,omitempty"`
	InitSettings      *InitRepoSettings  `json:"init_settings,omitempty"`
	TemplatingOptions *TemplatingOptions `json:"templating_options,omitempty"`
}

type ForkRepositoryBody struct {
	OrgSlug           string `json:"org_slug,omitempty"`
	OrgID             string `json:"org_id,omitempty"`
	Slug              string `json:"slug,omitempty"`
	DefaultBranchOnly bool   `json:"default_branch_only,omitempty"`
}

type ForkRepositoryResponse struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Slug         string               `json:"slug"`
	Description  string               `json:"description"`
	CloneURL     CloneURL             `json:"clone_url"`
	Organization OrganizationEmbedded `json:"organization"`
	Parent       *RepositoryEmbedded  `json:"parent,omitempty"`
}

type OrganizationEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type RepositoryEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}
