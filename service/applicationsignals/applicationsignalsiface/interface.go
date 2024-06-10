// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationsignalsiface provides an interface to enable mocking the Amazon CloudWatch Application Signals service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationsignalsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/applicationsignals"
)

// ApplicationSignalsAPI provides an interface to enable mocking the
// applicationsignals.ApplicationSignals service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon CloudWatch Application Signals.
//	func myFunc(svc applicationsignalsiface.ApplicationSignalsAPI) bool {
//	    // Make svc.BatchGetServiceLevelObjectiveBudgetReport request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := applicationsignals.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockApplicationSignalsClient struct {
//	    applicationsignalsiface.ApplicationSignalsAPI
//	}
//	func (m *mockApplicationSignalsClient) BatchGetServiceLevelObjectiveBudgetReport(input *applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput) (*applicationsignals.BatchGetServiceLevelObjectiveBudgetReportOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockApplicationSignalsClient{}
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
type ApplicationSignalsAPI interface {
	BatchGetServiceLevelObjectiveBudgetReport(*applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput) (*applicationsignals.BatchGetServiceLevelObjectiveBudgetReportOutput, error)
	BatchGetServiceLevelObjectiveBudgetReportWithContext(aws.Context, *applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput, ...request.Option) (*applicationsignals.BatchGetServiceLevelObjectiveBudgetReportOutput, error)
	BatchGetServiceLevelObjectiveBudgetReportRequest(*applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput) (*request.Request, *applicationsignals.BatchGetServiceLevelObjectiveBudgetReportOutput)

	CreateServiceLevelObjective(*applicationsignals.CreateServiceLevelObjectiveInput) (*applicationsignals.CreateServiceLevelObjectiveOutput, error)
	CreateServiceLevelObjectiveWithContext(aws.Context, *applicationsignals.CreateServiceLevelObjectiveInput, ...request.Option) (*applicationsignals.CreateServiceLevelObjectiveOutput, error)
	CreateServiceLevelObjectiveRequest(*applicationsignals.CreateServiceLevelObjectiveInput) (*request.Request, *applicationsignals.CreateServiceLevelObjectiveOutput)

	DeleteServiceLevelObjective(*applicationsignals.DeleteServiceLevelObjectiveInput) (*applicationsignals.DeleteServiceLevelObjectiveOutput, error)
	DeleteServiceLevelObjectiveWithContext(aws.Context, *applicationsignals.DeleteServiceLevelObjectiveInput, ...request.Option) (*applicationsignals.DeleteServiceLevelObjectiveOutput, error)
	DeleteServiceLevelObjectiveRequest(*applicationsignals.DeleteServiceLevelObjectiveInput) (*request.Request, *applicationsignals.DeleteServiceLevelObjectiveOutput)

	GetService(*applicationsignals.GetServiceInput) (*applicationsignals.GetServiceOutput, error)
	GetServiceWithContext(aws.Context, *applicationsignals.GetServiceInput, ...request.Option) (*applicationsignals.GetServiceOutput, error)
	GetServiceRequest(*applicationsignals.GetServiceInput) (*request.Request, *applicationsignals.GetServiceOutput)

	GetServiceLevelObjective(*applicationsignals.GetServiceLevelObjectiveInput) (*applicationsignals.GetServiceLevelObjectiveOutput, error)
	GetServiceLevelObjectiveWithContext(aws.Context, *applicationsignals.GetServiceLevelObjectiveInput, ...request.Option) (*applicationsignals.GetServiceLevelObjectiveOutput, error)
	GetServiceLevelObjectiveRequest(*applicationsignals.GetServiceLevelObjectiveInput) (*request.Request, *applicationsignals.GetServiceLevelObjectiveOutput)

	ListServiceDependencies(*applicationsignals.ListServiceDependenciesInput) (*applicationsignals.ListServiceDependenciesOutput, error)
	ListServiceDependenciesWithContext(aws.Context, *applicationsignals.ListServiceDependenciesInput, ...request.Option) (*applicationsignals.ListServiceDependenciesOutput, error)
	ListServiceDependenciesRequest(*applicationsignals.ListServiceDependenciesInput) (*request.Request, *applicationsignals.ListServiceDependenciesOutput)

	ListServiceDependenciesPages(*applicationsignals.ListServiceDependenciesInput, func(*applicationsignals.ListServiceDependenciesOutput, bool) bool) error
	ListServiceDependenciesPagesWithContext(aws.Context, *applicationsignals.ListServiceDependenciesInput, func(*applicationsignals.ListServiceDependenciesOutput, bool) bool, ...request.Option) error

	ListServiceDependents(*applicationsignals.ListServiceDependentsInput) (*applicationsignals.ListServiceDependentsOutput, error)
	ListServiceDependentsWithContext(aws.Context, *applicationsignals.ListServiceDependentsInput, ...request.Option) (*applicationsignals.ListServiceDependentsOutput, error)
	ListServiceDependentsRequest(*applicationsignals.ListServiceDependentsInput) (*request.Request, *applicationsignals.ListServiceDependentsOutput)

	ListServiceDependentsPages(*applicationsignals.ListServiceDependentsInput, func(*applicationsignals.ListServiceDependentsOutput, bool) bool) error
	ListServiceDependentsPagesWithContext(aws.Context, *applicationsignals.ListServiceDependentsInput, func(*applicationsignals.ListServiceDependentsOutput, bool) bool, ...request.Option) error

	ListServiceLevelObjectives(*applicationsignals.ListServiceLevelObjectivesInput) (*applicationsignals.ListServiceLevelObjectivesOutput, error)
	ListServiceLevelObjectivesWithContext(aws.Context, *applicationsignals.ListServiceLevelObjectivesInput, ...request.Option) (*applicationsignals.ListServiceLevelObjectivesOutput, error)
	ListServiceLevelObjectivesRequest(*applicationsignals.ListServiceLevelObjectivesInput) (*request.Request, *applicationsignals.ListServiceLevelObjectivesOutput)

	ListServiceLevelObjectivesPages(*applicationsignals.ListServiceLevelObjectivesInput, func(*applicationsignals.ListServiceLevelObjectivesOutput, bool) bool) error
	ListServiceLevelObjectivesPagesWithContext(aws.Context, *applicationsignals.ListServiceLevelObjectivesInput, func(*applicationsignals.ListServiceLevelObjectivesOutput, bool) bool, ...request.Option) error

	ListServiceOperations(*applicationsignals.ListServiceOperationsInput) (*applicationsignals.ListServiceOperationsOutput, error)
	ListServiceOperationsWithContext(aws.Context, *applicationsignals.ListServiceOperationsInput, ...request.Option) (*applicationsignals.ListServiceOperationsOutput, error)
	ListServiceOperationsRequest(*applicationsignals.ListServiceOperationsInput) (*request.Request, *applicationsignals.ListServiceOperationsOutput)

	ListServiceOperationsPages(*applicationsignals.ListServiceOperationsInput, func(*applicationsignals.ListServiceOperationsOutput, bool) bool) error
	ListServiceOperationsPagesWithContext(aws.Context, *applicationsignals.ListServiceOperationsInput, func(*applicationsignals.ListServiceOperationsOutput, bool) bool, ...request.Option) error

	ListServices(*applicationsignals.ListServicesInput) (*applicationsignals.ListServicesOutput, error)
	ListServicesWithContext(aws.Context, *applicationsignals.ListServicesInput, ...request.Option) (*applicationsignals.ListServicesOutput, error)
	ListServicesRequest(*applicationsignals.ListServicesInput) (*request.Request, *applicationsignals.ListServicesOutput)

	ListServicesPages(*applicationsignals.ListServicesInput, func(*applicationsignals.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *applicationsignals.ListServicesInput, func(*applicationsignals.ListServicesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*applicationsignals.ListTagsForResourceInput) (*applicationsignals.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *applicationsignals.ListTagsForResourceInput, ...request.Option) (*applicationsignals.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*applicationsignals.ListTagsForResourceInput) (*request.Request, *applicationsignals.ListTagsForResourceOutput)

	StartDiscovery(*applicationsignals.StartDiscoveryInput) (*applicationsignals.StartDiscoveryOutput, error)
	StartDiscoveryWithContext(aws.Context, *applicationsignals.StartDiscoveryInput, ...request.Option) (*applicationsignals.StartDiscoveryOutput, error)
	StartDiscoveryRequest(*applicationsignals.StartDiscoveryInput) (*request.Request, *applicationsignals.StartDiscoveryOutput)

	TagResource(*applicationsignals.TagResourceInput) (*applicationsignals.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *applicationsignals.TagResourceInput, ...request.Option) (*applicationsignals.TagResourceOutput, error)
	TagResourceRequest(*applicationsignals.TagResourceInput) (*request.Request, *applicationsignals.TagResourceOutput)

	UntagResource(*applicationsignals.UntagResourceInput) (*applicationsignals.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *applicationsignals.UntagResourceInput, ...request.Option) (*applicationsignals.UntagResourceOutput, error)
	UntagResourceRequest(*applicationsignals.UntagResourceInput) (*request.Request, *applicationsignals.UntagResourceOutput)

	UpdateServiceLevelObjective(*applicationsignals.UpdateServiceLevelObjectiveInput) (*applicationsignals.UpdateServiceLevelObjectiveOutput, error)
	UpdateServiceLevelObjectiveWithContext(aws.Context, *applicationsignals.UpdateServiceLevelObjectiveInput, ...request.Option) (*applicationsignals.UpdateServiceLevelObjectiveOutput, error)
	UpdateServiceLevelObjectiveRequest(*applicationsignals.UpdateServiceLevelObjectiveInput) (*request.Request, *applicationsignals.UpdateServiceLevelObjectiveOutput)
}

var _ ApplicationSignalsAPI = (*applicationsignals.ApplicationSignals)(nil)
