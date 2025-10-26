package api

import "time"

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
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type MergePullRequestRequest struct {
	MergeParameters *MergeParameters `json:"merge_parameters,omitempty"`
	Silent          bool             `json:"silent,omitempty"`
}

type ReviewDecision string

const (
	ReviewDecisionApprove ReviewDecision = "approve"
	ReviewDecisionTrust   ReviewDecision = "trust"
	ReviewDecisionBlock   ReviewDecision = "block"
	ReviewDecisionAbstain ReviewDecision = "abstain"
)

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

type CreatePullRequestBody struct {
	Title        string   `json:"title"`
	Description  string   `json:"description,omitempty"`
	SourceBranch string   `json:"source_branch"`
	TargetBranch string   `json:"target_branch"`
	ForkRepoID   string   `json:"fork_repo_id,omitempty"`
	ReviewerIDs  []string `json:"reviewer_ids,omitempty"`
	Publish      bool     `json:"publish"`
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
	Parent       *RepoEmbedded        `json:"parent,omitempty"`
}

type OrganizationEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type ListIssuesAssignedToAuthenticatedUserResponse struct {
	Issues        []Issue `json:"issues"`
	NextPageToken string  `json:"next_page_token"`
}

type Issue struct {
	ID          string                `json:"id"`
	Slug        string                `json:"slug"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Status      IssueStatus           `json:"status"`
	Author      UserEmbedded          `json:"author"`
	UpdatedBy   UserEmbedded          `json:"updated_by"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Assignee    UserEmbedded          `json:"assignee"`
	Labels      []LabelEmbedded       `json:"labels"`
	LinkedPRs   []PullRequestEmbedded `json:"linked_prs"`
	Priority    Priority              `json:"priority"`
	Visibility  IssueVisibility       `json:"visibility"`
	Milestone   MilestoneEmbedded     `json:"milestone"`
	Deadline    *time.Time            `json:"deadline,omitempty"`
	StartedAt   *time.Time            `json:"started_at,omitempty"`
	CompletedAt *time.Time            `json:"completed_at,omitempty"`
}

type IssueStatus struct {
	ID         string     `json:"id"`
	Slug       string     `json:"slug"`
	Name       string     `json:"name"`
	StatusType StatusType `json:"status_type"`
}

type StatusType string

const (
	StatusTypeInitial    StatusType = "initial"
	StatusTypeInProgress StatusType = "in_progress"
	StatusTypePaused     StatusType = "paused"
	StatusTypeCompleted  StatusType = "completed"
	StatusTypeCancelled  StatusType = "cancelled"
)

type LabelEmbedded struct {
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PullRequestEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type Priority string

const (
	PriorityTrivial  Priority = "trivial"
	PriorityMinor    Priority = "minor"
	PriorityNormal   Priority = "normal"
	PriorityCritical Priority = "critical"
	PriorityBlocker  Priority = "blocker"
)

type IssueVisibility string

const (
	IssueVisibilityPublic  IssueVisibility = "public"
	IssueVisibilityPrivate IssueVisibility = "private"
)

type MilestoneEmbedded struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type CreateIssueBody struct {
	Title         string          `json:"title"`
	Description   string          `json:"description,omitempty"`
	StatusSlug    string          `json:"status_slug,omitempty"`
	Priority      Priority        `json:"priority,omitempty"`
	AssigneeID    string          `json:"assignee_id,omitempty"`
	MilestoneID   string          `json:"milestone_id,omitempty"`
	MilestoneSlug string          `json:"milestone_slug,omitempty"`
	Visibility    IssueVisibility `json:"visibility,omitempty"`
	LabelIDs      []string        `json:"label_ids,omitempty"`
	LabelSlugs    []string        `json:"label_slugs,omitempty"`
	LinkedPRIDs   []string        `json:"linked_pr_ids,omitempty"`
	LinkedPRSlugs []string        `json:"linked_pr_slugs,omitempty"`
	Deadline      string          `json:"deadline,omitempty"`
}

type Label struct {
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
