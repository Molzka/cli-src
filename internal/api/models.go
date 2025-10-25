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
