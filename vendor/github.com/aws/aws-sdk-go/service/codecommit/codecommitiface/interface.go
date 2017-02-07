// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codecommitiface provides an interface to enable mocking the AWS CodeCommit service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codecommitiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codecommit"
)

// CodeCommitAPI provides an interface to enable mocking the
// codecommit.CodeCommit service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeCommit.
//    func myFunc(svc codecommitiface.CodeCommitAPI) bool {
//        // Make svc.BatchGetRepositories request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codecommit.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeCommitClient struct {
//        codecommitiface.CodeCommitAPI
//    }
//    func (m *mockCodeCommitClient) BatchGetRepositories(input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeCommitClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CodeCommitAPI interface {
	BatchGetRepositoriesRequest(*codecommit.BatchGetRepositoriesInput) (*request.Request, *codecommit.BatchGetRepositoriesOutput)

	BatchGetRepositories(*codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error)

	CreateBranchRequest(*codecommit.CreateBranchInput) (*request.Request, *codecommit.CreateBranchOutput)

	CreateBranch(*codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error)

	CreateRepositoryRequest(*codecommit.CreateRepositoryInput) (*request.Request, *codecommit.CreateRepositoryOutput)

	CreateRepository(*codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error)

	DeleteRepositoryRequest(*codecommit.DeleteRepositoryInput) (*request.Request, *codecommit.DeleteRepositoryOutput)

	DeleteRepository(*codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error)

	GetBlobRequest(*codecommit.GetBlobInput) (*request.Request, *codecommit.GetBlobOutput)

	GetBlob(*codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error)

	GetBranchRequest(*codecommit.GetBranchInput) (*request.Request, *codecommit.GetBranchOutput)

	GetBranch(*codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error)

	GetCommitRequest(*codecommit.GetCommitInput) (*request.Request, *codecommit.GetCommitOutput)

	GetCommit(*codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error)

	GetDifferencesRequest(*codecommit.GetDifferencesInput) (*request.Request, *codecommit.GetDifferencesOutput)

	GetDifferences(*codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error)

	GetDifferencesPages(*codecommit.GetDifferencesInput, func(*codecommit.GetDifferencesOutput, bool) bool) error

	GetRepositoryRequest(*codecommit.GetRepositoryInput) (*request.Request, *codecommit.GetRepositoryOutput)

	GetRepository(*codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error)

	GetRepositoryTriggersRequest(*codecommit.GetRepositoryTriggersInput) (*request.Request, *codecommit.GetRepositoryTriggersOutput)

	GetRepositoryTriggers(*codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error)

	ListBranchesRequest(*codecommit.ListBranchesInput) (*request.Request, *codecommit.ListBranchesOutput)

	ListBranches(*codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error)

	ListBranchesPages(*codecommit.ListBranchesInput, func(*codecommit.ListBranchesOutput, bool) bool) error

	ListRepositoriesRequest(*codecommit.ListRepositoriesInput) (*request.Request, *codecommit.ListRepositoriesOutput)

	ListRepositories(*codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error)

	ListRepositoriesPages(*codecommit.ListRepositoriesInput, func(*codecommit.ListRepositoriesOutput, bool) bool) error

	PutRepositoryTriggersRequest(*codecommit.PutRepositoryTriggersInput) (*request.Request, *codecommit.PutRepositoryTriggersOutput)

	PutRepositoryTriggers(*codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error)

	TestRepositoryTriggersRequest(*codecommit.TestRepositoryTriggersInput) (*request.Request, *codecommit.TestRepositoryTriggersOutput)

	TestRepositoryTriggers(*codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error)

	UpdateDefaultBranchRequest(*codecommit.UpdateDefaultBranchInput) (*request.Request, *codecommit.UpdateDefaultBranchOutput)

	UpdateDefaultBranch(*codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error)

	UpdateRepositoryDescriptionRequest(*codecommit.UpdateRepositoryDescriptionInput) (*request.Request, *codecommit.UpdateRepositoryDescriptionOutput)

	UpdateRepositoryDescription(*codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error)

	UpdateRepositoryNameRequest(*codecommit.UpdateRepositoryNameInput) (*request.Request, *codecommit.UpdateRepositoryNameOutput)

	UpdateRepositoryName(*codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error)
}

var _ CodeCommitAPI = (*codecommit.CodeCommit)(nil)
