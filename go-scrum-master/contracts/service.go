package contracts

type Service interface {
	IssueExists(issueId string) bool
}
