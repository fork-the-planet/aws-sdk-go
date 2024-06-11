// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package accessanalyzeriface provides an interface to enable mocking the Access Analyzer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package accessanalyzeriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
)

// AccessAnalyzerAPI provides an interface to enable mocking the
// accessanalyzer.AccessAnalyzer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Access Analyzer.
//	func myFunc(svc accessanalyzeriface.AccessAnalyzerAPI) bool {
//	    // Make svc.ApplyArchiveRule request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := accessanalyzer.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAccessAnalyzerClient struct {
//	    accessanalyzeriface.AccessAnalyzerAPI
//	}
//	func (m *mockAccessAnalyzerClient) ApplyArchiveRule(input *accessanalyzer.ApplyArchiveRuleInput) (*accessanalyzer.ApplyArchiveRuleOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAccessAnalyzerClient{}
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
type AccessAnalyzerAPI interface {
	ApplyArchiveRule(*accessanalyzer.ApplyArchiveRuleInput) (*accessanalyzer.ApplyArchiveRuleOutput, error)
	ApplyArchiveRuleWithContext(aws.Context, *accessanalyzer.ApplyArchiveRuleInput, ...request.Option) (*accessanalyzer.ApplyArchiveRuleOutput, error)
	ApplyArchiveRuleRequest(*accessanalyzer.ApplyArchiveRuleInput) (*request.Request, *accessanalyzer.ApplyArchiveRuleOutput)

	CancelPolicyGeneration(*accessanalyzer.CancelPolicyGenerationInput) (*accessanalyzer.CancelPolicyGenerationOutput, error)
	CancelPolicyGenerationWithContext(aws.Context, *accessanalyzer.CancelPolicyGenerationInput, ...request.Option) (*accessanalyzer.CancelPolicyGenerationOutput, error)
	CancelPolicyGenerationRequest(*accessanalyzer.CancelPolicyGenerationInput) (*request.Request, *accessanalyzer.CancelPolicyGenerationOutput)

	CheckAccessNotGranted(*accessanalyzer.CheckAccessNotGrantedInput) (*accessanalyzer.CheckAccessNotGrantedOutput, error)
	CheckAccessNotGrantedWithContext(aws.Context, *accessanalyzer.CheckAccessNotGrantedInput, ...request.Option) (*accessanalyzer.CheckAccessNotGrantedOutput, error)
	CheckAccessNotGrantedRequest(*accessanalyzer.CheckAccessNotGrantedInput) (*request.Request, *accessanalyzer.CheckAccessNotGrantedOutput)

	CheckNoNewAccess(*accessanalyzer.CheckNoNewAccessInput) (*accessanalyzer.CheckNoNewAccessOutput, error)
	CheckNoNewAccessWithContext(aws.Context, *accessanalyzer.CheckNoNewAccessInput, ...request.Option) (*accessanalyzer.CheckNoNewAccessOutput, error)
	CheckNoNewAccessRequest(*accessanalyzer.CheckNoNewAccessInput) (*request.Request, *accessanalyzer.CheckNoNewAccessOutput)

	CheckNoPublicAccess(*accessanalyzer.CheckNoPublicAccessInput) (*accessanalyzer.CheckNoPublicAccessOutput, error)
	CheckNoPublicAccessWithContext(aws.Context, *accessanalyzer.CheckNoPublicAccessInput, ...request.Option) (*accessanalyzer.CheckNoPublicAccessOutput, error)
	CheckNoPublicAccessRequest(*accessanalyzer.CheckNoPublicAccessInput) (*request.Request, *accessanalyzer.CheckNoPublicAccessOutput)

	CreateAccessPreview(*accessanalyzer.CreateAccessPreviewInput) (*accessanalyzer.CreateAccessPreviewOutput, error)
	CreateAccessPreviewWithContext(aws.Context, *accessanalyzer.CreateAccessPreviewInput, ...request.Option) (*accessanalyzer.CreateAccessPreviewOutput, error)
	CreateAccessPreviewRequest(*accessanalyzer.CreateAccessPreviewInput) (*request.Request, *accessanalyzer.CreateAccessPreviewOutput)

	CreateAnalyzer(*accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error)
	CreateAnalyzerWithContext(aws.Context, *accessanalyzer.CreateAnalyzerInput, ...request.Option) (*accessanalyzer.CreateAnalyzerOutput, error)
	CreateAnalyzerRequest(*accessanalyzer.CreateAnalyzerInput) (*request.Request, *accessanalyzer.CreateAnalyzerOutput)

	CreateArchiveRule(*accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error)
	CreateArchiveRuleWithContext(aws.Context, *accessanalyzer.CreateArchiveRuleInput, ...request.Option) (*accessanalyzer.CreateArchiveRuleOutput, error)
	CreateArchiveRuleRequest(*accessanalyzer.CreateArchiveRuleInput) (*request.Request, *accessanalyzer.CreateArchiveRuleOutput)

	DeleteAnalyzer(*accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error)
	DeleteAnalyzerWithContext(aws.Context, *accessanalyzer.DeleteAnalyzerInput, ...request.Option) (*accessanalyzer.DeleteAnalyzerOutput, error)
	DeleteAnalyzerRequest(*accessanalyzer.DeleteAnalyzerInput) (*request.Request, *accessanalyzer.DeleteAnalyzerOutput)

	DeleteArchiveRule(*accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error)
	DeleteArchiveRuleWithContext(aws.Context, *accessanalyzer.DeleteArchiveRuleInput, ...request.Option) (*accessanalyzer.DeleteArchiveRuleOutput, error)
	DeleteArchiveRuleRequest(*accessanalyzer.DeleteArchiveRuleInput) (*request.Request, *accessanalyzer.DeleteArchiveRuleOutput)

	GenerateFindingRecommendation(*accessanalyzer.GenerateFindingRecommendationInput) (*accessanalyzer.GenerateFindingRecommendationOutput, error)
	GenerateFindingRecommendationWithContext(aws.Context, *accessanalyzer.GenerateFindingRecommendationInput, ...request.Option) (*accessanalyzer.GenerateFindingRecommendationOutput, error)
	GenerateFindingRecommendationRequest(*accessanalyzer.GenerateFindingRecommendationInput) (*request.Request, *accessanalyzer.GenerateFindingRecommendationOutput)

	GetAccessPreview(*accessanalyzer.GetAccessPreviewInput) (*accessanalyzer.GetAccessPreviewOutput, error)
	GetAccessPreviewWithContext(aws.Context, *accessanalyzer.GetAccessPreviewInput, ...request.Option) (*accessanalyzer.GetAccessPreviewOutput, error)
	GetAccessPreviewRequest(*accessanalyzer.GetAccessPreviewInput) (*request.Request, *accessanalyzer.GetAccessPreviewOutput)

	GetAnalyzedResource(*accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error)
	GetAnalyzedResourceWithContext(aws.Context, *accessanalyzer.GetAnalyzedResourceInput, ...request.Option) (*accessanalyzer.GetAnalyzedResourceOutput, error)
	GetAnalyzedResourceRequest(*accessanalyzer.GetAnalyzedResourceInput) (*request.Request, *accessanalyzer.GetAnalyzedResourceOutput)

	GetAnalyzer(*accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error)
	GetAnalyzerWithContext(aws.Context, *accessanalyzer.GetAnalyzerInput, ...request.Option) (*accessanalyzer.GetAnalyzerOutput, error)
	GetAnalyzerRequest(*accessanalyzer.GetAnalyzerInput) (*request.Request, *accessanalyzer.GetAnalyzerOutput)

	GetArchiveRule(*accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error)
	GetArchiveRuleWithContext(aws.Context, *accessanalyzer.GetArchiveRuleInput, ...request.Option) (*accessanalyzer.GetArchiveRuleOutput, error)
	GetArchiveRuleRequest(*accessanalyzer.GetArchiveRuleInput) (*request.Request, *accessanalyzer.GetArchiveRuleOutput)

	GetFinding(*accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error)
	GetFindingWithContext(aws.Context, *accessanalyzer.GetFindingInput, ...request.Option) (*accessanalyzer.GetFindingOutput, error)
	GetFindingRequest(*accessanalyzer.GetFindingInput) (*request.Request, *accessanalyzer.GetFindingOutput)

	GetFindingRecommendation(*accessanalyzer.GetFindingRecommendationInput) (*accessanalyzer.GetFindingRecommendationOutput, error)
	GetFindingRecommendationWithContext(aws.Context, *accessanalyzer.GetFindingRecommendationInput, ...request.Option) (*accessanalyzer.GetFindingRecommendationOutput, error)
	GetFindingRecommendationRequest(*accessanalyzer.GetFindingRecommendationInput) (*request.Request, *accessanalyzer.GetFindingRecommendationOutput)

	GetFindingRecommendationPages(*accessanalyzer.GetFindingRecommendationInput, func(*accessanalyzer.GetFindingRecommendationOutput, bool) bool) error
	GetFindingRecommendationPagesWithContext(aws.Context, *accessanalyzer.GetFindingRecommendationInput, func(*accessanalyzer.GetFindingRecommendationOutput, bool) bool, ...request.Option) error

	GetFindingV2(*accessanalyzer.GetFindingV2Input) (*accessanalyzer.GetFindingV2Output, error)
	GetFindingV2WithContext(aws.Context, *accessanalyzer.GetFindingV2Input, ...request.Option) (*accessanalyzer.GetFindingV2Output, error)
	GetFindingV2Request(*accessanalyzer.GetFindingV2Input) (*request.Request, *accessanalyzer.GetFindingV2Output)

	GetFindingV2Pages(*accessanalyzer.GetFindingV2Input, func(*accessanalyzer.GetFindingV2Output, bool) bool) error
	GetFindingV2PagesWithContext(aws.Context, *accessanalyzer.GetFindingV2Input, func(*accessanalyzer.GetFindingV2Output, bool) bool, ...request.Option) error

	GetGeneratedPolicy(*accessanalyzer.GetGeneratedPolicyInput) (*accessanalyzer.GetGeneratedPolicyOutput, error)
	GetGeneratedPolicyWithContext(aws.Context, *accessanalyzer.GetGeneratedPolicyInput, ...request.Option) (*accessanalyzer.GetGeneratedPolicyOutput, error)
	GetGeneratedPolicyRequest(*accessanalyzer.GetGeneratedPolicyInput) (*request.Request, *accessanalyzer.GetGeneratedPolicyOutput)

	ListAccessPreviewFindings(*accessanalyzer.ListAccessPreviewFindingsInput) (*accessanalyzer.ListAccessPreviewFindingsOutput, error)
	ListAccessPreviewFindingsWithContext(aws.Context, *accessanalyzer.ListAccessPreviewFindingsInput, ...request.Option) (*accessanalyzer.ListAccessPreviewFindingsOutput, error)
	ListAccessPreviewFindingsRequest(*accessanalyzer.ListAccessPreviewFindingsInput) (*request.Request, *accessanalyzer.ListAccessPreviewFindingsOutput)

	ListAccessPreviewFindingsPages(*accessanalyzer.ListAccessPreviewFindingsInput, func(*accessanalyzer.ListAccessPreviewFindingsOutput, bool) bool) error
	ListAccessPreviewFindingsPagesWithContext(aws.Context, *accessanalyzer.ListAccessPreviewFindingsInput, func(*accessanalyzer.ListAccessPreviewFindingsOutput, bool) bool, ...request.Option) error

	ListAccessPreviews(*accessanalyzer.ListAccessPreviewsInput) (*accessanalyzer.ListAccessPreviewsOutput, error)
	ListAccessPreviewsWithContext(aws.Context, *accessanalyzer.ListAccessPreviewsInput, ...request.Option) (*accessanalyzer.ListAccessPreviewsOutput, error)
	ListAccessPreviewsRequest(*accessanalyzer.ListAccessPreviewsInput) (*request.Request, *accessanalyzer.ListAccessPreviewsOutput)

	ListAccessPreviewsPages(*accessanalyzer.ListAccessPreviewsInput, func(*accessanalyzer.ListAccessPreviewsOutput, bool) bool) error
	ListAccessPreviewsPagesWithContext(aws.Context, *accessanalyzer.ListAccessPreviewsInput, func(*accessanalyzer.ListAccessPreviewsOutput, bool) bool, ...request.Option) error

	ListAnalyzedResources(*accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
	ListAnalyzedResourcesWithContext(aws.Context, *accessanalyzer.ListAnalyzedResourcesInput, ...request.Option) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
	ListAnalyzedResourcesRequest(*accessanalyzer.ListAnalyzedResourcesInput) (*request.Request, *accessanalyzer.ListAnalyzedResourcesOutput)

	ListAnalyzedResourcesPages(*accessanalyzer.ListAnalyzedResourcesInput, func(*accessanalyzer.ListAnalyzedResourcesOutput, bool) bool) error
	ListAnalyzedResourcesPagesWithContext(aws.Context, *accessanalyzer.ListAnalyzedResourcesInput, func(*accessanalyzer.ListAnalyzedResourcesOutput, bool) bool, ...request.Option) error

	ListAnalyzers(*accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error)
	ListAnalyzersWithContext(aws.Context, *accessanalyzer.ListAnalyzersInput, ...request.Option) (*accessanalyzer.ListAnalyzersOutput, error)
	ListAnalyzersRequest(*accessanalyzer.ListAnalyzersInput) (*request.Request, *accessanalyzer.ListAnalyzersOutput)

	ListAnalyzersPages(*accessanalyzer.ListAnalyzersInput, func(*accessanalyzer.ListAnalyzersOutput, bool) bool) error
	ListAnalyzersPagesWithContext(aws.Context, *accessanalyzer.ListAnalyzersInput, func(*accessanalyzer.ListAnalyzersOutput, bool) bool, ...request.Option) error

	ListArchiveRules(*accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error)
	ListArchiveRulesWithContext(aws.Context, *accessanalyzer.ListArchiveRulesInput, ...request.Option) (*accessanalyzer.ListArchiveRulesOutput, error)
	ListArchiveRulesRequest(*accessanalyzer.ListArchiveRulesInput) (*request.Request, *accessanalyzer.ListArchiveRulesOutput)

	ListArchiveRulesPages(*accessanalyzer.ListArchiveRulesInput, func(*accessanalyzer.ListArchiveRulesOutput, bool) bool) error
	ListArchiveRulesPagesWithContext(aws.Context, *accessanalyzer.ListArchiveRulesInput, func(*accessanalyzer.ListArchiveRulesOutput, bool) bool, ...request.Option) error

	ListFindings(*accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error)
	ListFindingsWithContext(aws.Context, *accessanalyzer.ListFindingsInput, ...request.Option) (*accessanalyzer.ListFindingsOutput, error)
	ListFindingsRequest(*accessanalyzer.ListFindingsInput) (*request.Request, *accessanalyzer.ListFindingsOutput)

	ListFindingsPages(*accessanalyzer.ListFindingsInput, func(*accessanalyzer.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *accessanalyzer.ListFindingsInput, func(*accessanalyzer.ListFindingsOutput, bool) bool, ...request.Option) error

	ListFindingsV2(*accessanalyzer.ListFindingsV2Input) (*accessanalyzer.ListFindingsV2Output, error)
	ListFindingsV2WithContext(aws.Context, *accessanalyzer.ListFindingsV2Input, ...request.Option) (*accessanalyzer.ListFindingsV2Output, error)
	ListFindingsV2Request(*accessanalyzer.ListFindingsV2Input) (*request.Request, *accessanalyzer.ListFindingsV2Output)

	ListFindingsV2Pages(*accessanalyzer.ListFindingsV2Input, func(*accessanalyzer.ListFindingsV2Output, bool) bool) error
	ListFindingsV2PagesWithContext(aws.Context, *accessanalyzer.ListFindingsV2Input, func(*accessanalyzer.ListFindingsV2Output, bool) bool, ...request.Option) error

	ListPolicyGenerations(*accessanalyzer.ListPolicyGenerationsInput) (*accessanalyzer.ListPolicyGenerationsOutput, error)
	ListPolicyGenerationsWithContext(aws.Context, *accessanalyzer.ListPolicyGenerationsInput, ...request.Option) (*accessanalyzer.ListPolicyGenerationsOutput, error)
	ListPolicyGenerationsRequest(*accessanalyzer.ListPolicyGenerationsInput) (*request.Request, *accessanalyzer.ListPolicyGenerationsOutput)

	ListPolicyGenerationsPages(*accessanalyzer.ListPolicyGenerationsInput, func(*accessanalyzer.ListPolicyGenerationsOutput, bool) bool) error
	ListPolicyGenerationsPagesWithContext(aws.Context, *accessanalyzer.ListPolicyGenerationsInput, func(*accessanalyzer.ListPolicyGenerationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *accessanalyzer.ListTagsForResourceInput, ...request.Option) (*accessanalyzer.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*accessanalyzer.ListTagsForResourceInput) (*request.Request, *accessanalyzer.ListTagsForResourceOutput)

	StartPolicyGeneration(*accessanalyzer.StartPolicyGenerationInput) (*accessanalyzer.StartPolicyGenerationOutput, error)
	StartPolicyGenerationWithContext(aws.Context, *accessanalyzer.StartPolicyGenerationInput, ...request.Option) (*accessanalyzer.StartPolicyGenerationOutput, error)
	StartPolicyGenerationRequest(*accessanalyzer.StartPolicyGenerationInput) (*request.Request, *accessanalyzer.StartPolicyGenerationOutput)

	StartResourceScan(*accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error)
	StartResourceScanWithContext(aws.Context, *accessanalyzer.StartResourceScanInput, ...request.Option) (*accessanalyzer.StartResourceScanOutput, error)
	StartResourceScanRequest(*accessanalyzer.StartResourceScanInput) (*request.Request, *accessanalyzer.StartResourceScanOutput)

	TagResource(*accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *accessanalyzer.TagResourceInput, ...request.Option) (*accessanalyzer.TagResourceOutput, error)
	TagResourceRequest(*accessanalyzer.TagResourceInput) (*request.Request, *accessanalyzer.TagResourceOutput)

	UntagResource(*accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *accessanalyzer.UntagResourceInput, ...request.Option) (*accessanalyzer.UntagResourceOutput, error)
	UntagResourceRequest(*accessanalyzer.UntagResourceInput) (*request.Request, *accessanalyzer.UntagResourceOutput)

	UpdateArchiveRule(*accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error)
	UpdateArchiveRuleWithContext(aws.Context, *accessanalyzer.UpdateArchiveRuleInput, ...request.Option) (*accessanalyzer.UpdateArchiveRuleOutput, error)
	UpdateArchiveRuleRequest(*accessanalyzer.UpdateArchiveRuleInput) (*request.Request, *accessanalyzer.UpdateArchiveRuleOutput)

	UpdateFindings(*accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error)
	UpdateFindingsWithContext(aws.Context, *accessanalyzer.UpdateFindingsInput, ...request.Option) (*accessanalyzer.UpdateFindingsOutput, error)
	UpdateFindingsRequest(*accessanalyzer.UpdateFindingsInput) (*request.Request, *accessanalyzer.UpdateFindingsOutput)

	ValidatePolicy(*accessanalyzer.ValidatePolicyInput) (*accessanalyzer.ValidatePolicyOutput, error)
	ValidatePolicyWithContext(aws.Context, *accessanalyzer.ValidatePolicyInput, ...request.Option) (*accessanalyzer.ValidatePolicyOutput, error)
	ValidatePolicyRequest(*accessanalyzer.ValidatePolicyInput) (*request.Request, *accessanalyzer.ValidatePolicyOutput)

	ValidatePolicyPages(*accessanalyzer.ValidatePolicyInput, func(*accessanalyzer.ValidatePolicyOutput, bool) bool) error
	ValidatePolicyPagesWithContext(aws.Context, *accessanalyzer.ValidatePolicyInput, func(*accessanalyzer.ValidatePolicyOutput, bool) bool, ...request.Option) error
}

var _ AccessAnalyzerAPI = (*accessanalyzer.AccessAnalyzer)(nil)
