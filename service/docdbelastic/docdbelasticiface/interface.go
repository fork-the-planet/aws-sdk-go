// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package docdbelasticiface provides an interface to enable mocking the Amazon DocumentDB Elastic Clusters service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package docdbelasticiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/docdbelastic"
)

// DocDBElasticAPI provides an interface to enable mocking the
// docdbelastic.DocDBElastic service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon DocumentDB Elastic Clusters.
//	func myFunc(svc docdbelasticiface.DocDBElasticAPI) bool {
//	    // Make svc.CopyClusterSnapshot request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := docdbelastic.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockDocDBElasticClient struct {
//	    docdbelasticiface.DocDBElasticAPI
//	}
//	func (m *mockDocDBElasticClient) CopyClusterSnapshot(input *docdbelastic.CopyClusterSnapshotInput) (*docdbelastic.CopyClusterSnapshotOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockDocDBElasticClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DocDBElasticAPI interface {
	CopyClusterSnapshot(*docdbelastic.CopyClusterSnapshotInput) (*docdbelastic.CopyClusterSnapshotOutput, error)
	CopyClusterSnapshotWithContext(aws.Context, *docdbelastic.CopyClusterSnapshotInput, ...request.Option) (*docdbelastic.CopyClusterSnapshotOutput, error)
	CopyClusterSnapshotRequest(*docdbelastic.CopyClusterSnapshotInput) (*request.Request, *docdbelastic.CopyClusterSnapshotOutput)

	CreateCluster(*docdbelastic.CreateClusterInput) (*docdbelastic.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *docdbelastic.CreateClusterInput, ...request.Option) (*docdbelastic.CreateClusterOutput, error)
	CreateClusterRequest(*docdbelastic.CreateClusterInput) (*request.Request, *docdbelastic.CreateClusterOutput)

	CreateClusterSnapshot(*docdbelastic.CreateClusterSnapshotInput) (*docdbelastic.CreateClusterSnapshotOutput, error)
	CreateClusterSnapshotWithContext(aws.Context, *docdbelastic.CreateClusterSnapshotInput, ...request.Option) (*docdbelastic.CreateClusterSnapshotOutput, error)
	CreateClusterSnapshotRequest(*docdbelastic.CreateClusterSnapshotInput) (*request.Request, *docdbelastic.CreateClusterSnapshotOutput)

	DeleteCluster(*docdbelastic.DeleteClusterInput) (*docdbelastic.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *docdbelastic.DeleteClusterInput, ...request.Option) (*docdbelastic.DeleteClusterOutput, error)
	DeleteClusterRequest(*docdbelastic.DeleteClusterInput) (*request.Request, *docdbelastic.DeleteClusterOutput)

	DeleteClusterSnapshot(*docdbelastic.DeleteClusterSnapshotInput) (*docdbelastic.DeleteClusterSnapshotOutput, error)
	DeleteClusterSnapshotWithContext(aws.Context, *docdbelastic.DeleteClusterSnapshotInput, ...request.Option) (*docdbelastic.DeleteClusterSnapshotOutput, error)
	DeleteClusterSnapshotRequest(*docdbelastic.DeleteClusterSnapshotInput) (*request.Request, *docdbelastic.DeleteClusterSnapshotOutput)

	GetCluster(*docdbelastic.GetClusterInput) (*docdbelastic.GetClusterOutput, error)
	GetClusterWithContext(aws.Context, *docdbelastic.GetClusterInput, ...request.Option) (*docdbelastic.GetClusterOutput, error)
	GetClusterRequest(*docdbelastic.GetClusterInput) (*request.Request, *docdbelastic.GetClusterOutput)

	GetClusterSnapshot(*docdbelastic.GetClusterSnapshotInput) (*docdbelastic.GetClusterSnapshotOutput, error)
	GetClusterSnapshotWithContext(aws.Context, *docdbelastic.GetClusterSnapshotInput, ...request.Option) (*docdbelastic.GetClusterSnapshotOutput, error)
	GetClusterSnapshotRequest(*docdbelastic.GetClusterSnapshotInput) (*request.Request, *docdbelastic.GetClusterSnapshotOutput)

	ListClusterSnapshots(*docdbelastic.ListClusterSnapshotsInput) (*docdbelastic.ListClusterSnapshotsOutput, error)
	ListClusterSnapshotsWithContext(aws.Context, *docdbelastic.ListClusterSnapshotsInput, ...request.Option) (*docdbelastic.ListClusterSnapshotsOutput, error)
	ListClusterSnapshotsRequest(*docdbelastic.ListClusterSnapshotsInput) (*request.Request, *docdbelastic.ListClusterSnapshotsOutput)

	ListClusterSnapshotsPages(*docdbelastic.ListClusterSnapshotsInput, func(*docdbelastic.ListClusterSnapshotsOutput, bool) bool) error
	ListClusterSnapshotsPagesWithContext(aws.Context, *docdbelastic.ListClusterSnapshotsInput, func(*docdbelastic.ListClusterSnapshotsOutput, bool) bool, ...request.Option) error

	ListClusters(*docdbelastic.ListClustersInput) (*docdbelastic.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *docdbelastic.ListClustersInput, ...request.Option) (*docdbelastic.ListClustersOutput, error)
	ListClustersRequest(*docdbelastic.ListClustersInput) (*request.Request, *docdbelastic.ListClustersOutput)

	ListClustersPages(*docdbelastic.ListClustersInput, func(*docdbelastic.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *docdbelastic.ListClustersInput, func(*docdbelastic.ListClustersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*docdbelastic.ListTagsForResourceInput) (*docdbelastic.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *docdbelastic.ListTagsForResourceInput, ...request.Option) (*docdbelastic.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*docdbelastic.ListTagsForResourceInput) (*request.Request, *docdbelastic.ListTagsForResourceOutput)

	RestoreClusterFromSnapshot(*docdbelastic.RestoreClusterFromSnapshotInput) (*docdbelastic.RestoreClusterFromSnapshotOutput, error)
	RestoreClusterFromSnapshotWithContext(aws.Context, *docdbelastic.RestoreClusterFromSnapshotInput, ...request.Option) (*docdbelastic.RestoreClusterFromSnapshotOutput, error)
	RestoreClusterFromSnapshotRequest(*docdbelastic.RestoreClusterFromSnapshotInput) (*request.Request, *docdbelastic.RestoreClusterFromSnapshotOutput)

	StartCluster(*docdbelastic.StartClusterInput) (*docdbelastic.StartClusterOutput, error)
	StartClusterWithContext(aws.Context, *docdbelastic.StartClusterInput, ...request.Option) (*docdbelastic.StartClusterOutput, error)
	StartClusterRequest(*docdbelastic.StartClusterInput) (*request.Request, *docdbelastic.StartClusterOutput)

	StopCluster(*docdbelastic.StopClusterInput) (*docdbelastic.StopClusterOutput, error)
	StopClusterWithContext(aws.Context, *docdbelastic.StopClusterInput, ...request.Option) (*docdbelastic.StopClusterOutput, error)
	StopClusterRequest(*docdbelastic.StopClusterInput) (*request.Request, *docdbelastic.StopClusterOutput)

	TagResource(*docdbelastic.TagResourceInput) (*docdbelastic.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *docdbelastic.TagResourceInput, ...request.Option) (*docdbelastic.TagResourceOutput, error)
	TagResourceRequest(*docdbelastic.TagResourceInput) (*request.Request, *docdbelastic.TagResourceOutput)

	UntagResource(*docdbelastic.UntagResourceInput) (*docdbelastic.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *docdbelastic.UntagResourceInput, ...request.Option) (*docdbelastic.UntagResourceOutput, error)
	UntagResourceRequest(*docdbelastic.UntagResourceInput) (*request.Request, *docdbelastic.UntagResourceOutput)

	UpdateCluster(*docdbelastic.UpdateClusterInput) (*docdbelastic.UpdateClusterOutput, error)
	UpdateClusterWithContext(aws.Context, *docdbelastic.UpdateClusterInput, ...request.Option) (*docdbelastic.UpdateClusterOutput, error)
	UpdateClusterRequest(*docdbelastic.UpdateClusterInput) (*request.Request, *docdbelastic.UpdateClusterOutput)
}

var _ DocDBElasticAPI = (*docdbelastic.DocDBElastic)(nil)
