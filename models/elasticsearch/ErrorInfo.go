package elasticsearch

type ErrorInfo struct {
	RootCause []*ErrorInfo
	Type      string
	Reason    string
	Phase     string
}
