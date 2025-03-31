package awsclient

// ObjectInfo represents information about an S3 object
type ObjectInfo struct {
	Key  string
	Size int64
}

// S3Client defines the interface for interacting with S3
type S3Client interface {
	// ListObjects lists objects in a bucket with an optional prefix
	ListObjects(bucket string, prefix string) ([]ObjectInfo, error)
}

// ECSClient defines the interface for interacting with ECS
type ECSClient interface {
	// UpdateServiceTaskCount updates the desired count of tasks for a service
	UpdateServiceTaskCount(cluster string, service string, desiredCount int) error
} 