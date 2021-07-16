// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workmailiface provides an interface to enable mocking the Amazon WorkMail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package workmailiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmail"
)

// WorkMailAPI provides an interface to enable mocking the
// workmail.WorkMail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon WorkMail.
//    func myFunc(svc workmailiface.WorkMailAPI) bool {
//        // Make svc.AssociateDelegateToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := workmail.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWorkMailClient struct {
//        workmailiface.WorkMailAPI
//    }
//    func (m *mockWorkMailClient) AssociateDelegateToResource(input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWorkMailClient{}
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
type WorkMailAPI interface {
	AssociateDelegateToResource(*workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateDelegateToResourceWithContext(aws.Context, *workmail.AssociateDelegateToResourceInput, ...request.Option) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateDelegateToResourceRequest(*workmail.AssociateDelegateToResourceInput) (*request.Request, *workmail.AssociateDelegateToResourceOutput)

	AssociateMemberToGroup(*workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error)
	AssociateMemberToGroupWithContext(aws.Context, *workmail.AssociateMemberToGroupInput, ...request.Option) (*workmail.AssociateMemberToGroupOutput, error)
	AssociateMemberToGroupRequest(*workmail.AssociateMemberToGroupInput) (*request.Request, *workmail.AssociateMemberToGroupOutput)

	CreateAlias(*workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error)
	CreateAliasWithContext(aws.Context, *workmail.CreateAliasInput, ...request.Option) (*workmail.CreateAliasOutput, error)
	CreateAliasRequest(*workmail.CreateAliasInput) (*request.Request, *workmail.CreateAliasOutput)

	CreateGroup(*workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *workmail.CreateGroupInput, ...request.Option) (*workmail.CreateGroupOutput, error)
	CreateGroupRequest(*workmail.CreateGroupInput) (*request.Request, *workmail.CreateGroupOutput)

	CreateResource(*workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error)
	CreateResourceWithContext(aws.Context, *workmail.CreateResourceInput, ...request.Option) (*workmail.CreateResourceOutput, error)
	CreateResourceRequest(*workmail.CreateResourceInput) (*request.Request, *workmail.CreateResourceOutput)

	CreateUser(*workmail.CreateUserInput) (*workmail.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *workmail.CreateUserInput, ...request.Option) (*workmail.CreateUserOutput, error)
	CreateUserRequest(*workmail.CreateUserInput) (*request.Request, *workmail.CreateUserOutput)

	DeleteAlias(*workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error)
	DeleteAliasWithContext(aws.Context, *workmail.DeleteAliasInput, ...request.Option) (*workmail.DeleteAliasOutput, error)
	DeleteAliasRequest(*workmail.DeleteAliasInput) (*request.Request, *workmail.DeleteAliasOutput)

	DeleteGroup(*workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *workmail.DeleteGroupInput, ...request.Option) (*workmail.DeleteGroupOutput, error)
	DeleteGroupRequest(*workmail.DeleteGroupInput) (*request.Request, *workmail.DeleteGroupOutput)

	DeleteMailboxPermissions(*workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteMailboxPermissionsWithContext(aws.Context, *workmail.DeleteMailboxPermissionsInput, ...request.Option) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteMailboxPermissionsRequest(*workmail.DeleteMailboxPermissionsInput) (*request.Request, *workmail.DeleteMailboxPermissionsOutput)

	DeleteResource(*workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error)
	DeleteResourceWithContext(aws.Context, *workmail.DeleteResourceInput, ...request.Option) (*workmail.DeleteResourceOutput, error)
	DeleteResourceRequest(*workmail.DeleteResourceInput) (*request.Request, *workmail.DeleteResourceOutput)

	DeleteUser(*workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *workmail.DeleteUserInput, ...request.Option) (*workmail.DeleteUserOutput, error)
	DeleteUserRequest(*workmail.DeleteUserInput) (*request.Request, *workmail.DeleteUserOutput)

	DeregisterFromWorkMail(*workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error)
	DeregisterFromWorkMailWithContext(aws.Context, *workmail.DeregisterFromWorkMailInput, ...request.Option) (*workmail.DeregisterFromWorkMailOutput, error)
	DeregisterFromWorkMailRequest(*workmail.DeregisterFromWorkMailInput) (*request.Request, *workmail.DeregisterFromWorkMailOutput)

	DescribeGroup(*workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error)
	DescribeGroupWithContext(aws.Context, *workmail.DescribeGroupInput, ...request.Option) (*workmail.DescribeGroupOutput, error)
	DescribeGroupRequest(*workmail.DescribeGroupInput) (*request.Request, *workmail.DescribeGroupOutput)

	DescribeOrganization(*workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error)
	DescribeOrganizationWithContext(aws.Context, *workmail.DescribeOrganizationInput, ...request.Option) (*workmail.DescribeOrganizationOutput, error)
	DescribeOrganizationRequest(*workmail.DescribeOrganizationInput) (*request.Request, *workmail.DescribeOrganizationOutput)

	DescribeResource(*workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error)
	DescribeResourceWithContext(aws.Context, *workmail.DescribeResourceInput, ...request.Option) (*workmail.DescribeResourceOutput, error)
	DescribeResourceRequest(*workmail.DescribeResourceInput) (*request.Request, *workmail.DescribeResourceOutput)

	DescribeUser(*workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *workmail.DescribeUserInput, ...request.Option) (*workmail.DescribeUserOutput, error)
	DescribeUserRequest(*workmail.DescribeUserInput) (*request.Request, *workmail.DescribeUserOutput)

	DisassociateDelegateFromResource(*workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateDelegateFromResourceWithContext(aws.Context, *workmail.DisassociateDelegateFromResourceInput, ...request.Option) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateDelegateFromResourceRequest(*workmail.DisassociateDelegateFromResourceInput) (*request.Request, *workmail.DisassociateDelegateFromResourceOutput)

	DisassociateMemberFromGroup(*workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error)
	DisassociateMemberFromGroupWithContext(aws.Context, *workmail.DisassociateMemberFromGroupInput, ...request.Option) (*workmail.DisassociateMemberFromGroupOutput, error)
	DisassociateMemberFromGroupRequest(*workmail.DisassociateMemberFromGroupInput) (*request.Request, *workmail.DisassociateMemberFromGroupOutput)

	GetMailboxDetails(*workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error)
	GetMailboxDetailsWithContext(aws.Context, *workmail.GetMailboxDetailsInput, ...request.Option) (*workmail.GetMailboxDetailsOutput, error)
	GetMailboxDetailsRequest(*workmail.GetMailboxDetailsInput) (*request.Request, *workmail.GetMailboxDetailsOutput)

	ListAliases(*workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error)
	ListAliasesWithContext(aws.Context, *workmail.ListAliasesInput, ...request.Option) (*workmail.ListAliasesOutput, error)
	ListAliasesRequest(*workmail.ListAliasesInput) (*request.Request, *workmail.ListAliasesOutput)

	ListAliasesPages(*workmail.ListAliasesInput, func(*workmail.ListAliasesOutput, bool) bool) error
	ListAliasesPagesWithContext(aws.Context, *workmail.ListAliasesInput, func(*workmail.ListAliasesOutput, bool) bool, ...request.Option) error

	ListGroupMembers(*workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersWithContext(aws.Context, *workmail.ListGroupMembersInput, ...request.Option) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersRequest(*workmail.ListGroupMembersInput) (*request.Request, *workmail.ListGroupMembersOutput)

	ListGroupMembersPages(*workmail.ListGroupMembersInput, func(*workmail.ListGroupMembersOutput, bool) bool) error
	ListGroupMembersPagesWithContext(aws.Context, *workmail.ListGroupMembersInput, func(*workmail.ListGroupMembersOutput, bool) bool, ...request.Option) error

	ListGroups(*workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *workmail.ListGroupsInput, ...request.Option) (*workmail.ListGroupsOutput, error)
	ListGroupsRequest(*workmail.ListGroupsInput) (*request.Request, *workmail.ListGroupsOutput)

	ListGroupsPages(*workmail.ListGroupsInput, func(*workmail.ListGroupsOutput, bool) bool) error
	ListGroupsPagesWithContext(aws.Context, *workmail.ListGroupsInput, func(*workmail.ListGroupsOutput, bool) bool, ...request.Option) error

	ListMailboxPermissions(*workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsWithContext(aws.Context, *workmail.ListMailboxPermissionsInput, ...request.Option) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsRequest(*workmail.ListMailboxPermissionsInput) (*request.Request, *workmail.ListMailboxPermissionsOutput)

	ListMailboxPermissionsPages(*workmail.ListMailboxPermissionsInput, func(*workmail.ListMailboxPermissionsOutput, bool) bool) error
	ListMailboxPermissionsPagesWithContext(aws.Context, *workmail.ListMailboxPermissionsInput, func(*workmail.ListMailboxPermissionsOutput, bool) bool, ...request.Option) error

	ListOrganizations(*workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsWithContext(aws.Context, *workmail.ListOrganizationsInput, ...request.Option) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsRequest(*workmail.ListOrganizationsInput) (*request.Request, *workmail.ListOrganizationsOutput)

	ListOrganizationsPages(*workmail.ListOrganizationsInput, func(*workmail.ListOrganizationsOutput, bool) bool) error
	ListOrganizationsPagesWithContext(aws.Context, *workmail.ListOrganizationsInput, func(*workmail.ListOrganizationsOutput, bool) bool, ...request.Option) error

	ListResourceDelegates(*workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesWithContext(aws.Context, *workmail.ListResourceDelegatesInput, ...request.Option) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesRequest(*workmail.ListResourceDelegatesInput) (*request.Request, *workmail.ListResourceDelegatesOutput)

	ListResourceDelegatesPages(*workmail.ListResourceDelegatesInput, func(*workmail.ListResourceDelegatesOutput, bool) bool) error
	ListResourceDelegatesPagesWithContext(aws.Context, *workmail.ListResourceDelegatesInput, func(*workmail.ListResourceDelegatesOutput, bool) bool, ...request.Option) error

	ListResources(*workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *workmail.ListResourcesInput, ...request.Option) (*workmail.ListResourcesOutput, error)
	ListResourcesRequest(*workmail.ListResourcesInput) (*request.Request, *workmail.ListResourcesOutput)

	ListResourcesPages(*workmail.ListResourcesInput, func(*workmail.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *workmail.ListResourcesInput, func(*workmail.ListResourcesOutput, bool) bool, ...request.Option) error

	ListUsers(*workmail.ListUsersInput) (*workmail.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *workmail.ListUsersInput, ...request.Option) (*workmail.ListUsersOutput, error)
	ListUsersRequest(*workmail.ListUsersInput) (*request.Request, *workmail.ListUsersOutput)

	ListUsersPages(*workmail.ListUsersInput, func(*workmail.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *workmail.ListUsersInput, func(*workmail.ListUsersOutput, bool) bool, ...request.Option) error

	PutMailboxPermissions(*workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error)
	PutMailboxPermissionsWithContext(aws.Context, *workmail.PutMailboxPermissionsInput, ...request.Option) (*workmail.PutMailboxPermissionsOutput, error)
	PutMailboxPermissionsRequest(*workmail.PutMailboxPermissionsInput) (*request.Request, *workmail.PutMailboxPermissionsOutput)

	RegisterToWorkMail(*workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error)
	RegisterToWorkMailWithContext(aws.Context, *workmail.RegisterToWorkMailInput, ...request.Option) (*workmail.RegisterToWorkMailOutput, error)
	RegisterToWorkMailRequest(*workmail.RegisterToWorkMailInput) (*request.Request, *workmail.RegisterToWorkMailOutput)

	ResetPassword(*workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error)
	ResetPasswordWithContext(aws.Context, *workmail.ResetPasswordInput, ...request.Option) (*workmail.ResetPasswordOutput, error)
	ResetPasswordRequest(*workmail.ResetPasswordInput) (*request.Request, *workmail.ResetPasswordOutput)

	UpdateMailboxQuota(*workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdateMailboxQuotaWithContext(aws.Context, *workmail.UpdateMailboxQuotaInput, ...request.Option) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdateMailboxQuotaRequest(*workmail.UpdateMailboxQuotaInput) (*request.Request, *workmail.UpdateMailboxQuotaOutput)

	UpdatePrimaryEmailAddress(*workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdatePrimaryEmailAddressWithContext(aws.Context, *workmail.UpdatePrimaryEmailAddressInput, ...request.Option) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdatePrimaryEmailAddressRequest(*workmail.UpdatePrimaryEmailAddressInput) (*request.Request, *workmail.UpdatePrimaryEmailAddressOutput)

	UpdateResource(*workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error)
	UpdateResourceWithContext(aws.Context, *workmail.UpdateResourceInput, ...request.Option) (*workmail.UpdateResourceOutput, error)
	UpdateResourceRequest(*workmail.UpdateResourceInput) (*request.Request, *workmail.UpdateResourceOutput)
}

var _ WorkMailAPI = (*workmail.WorkMail)(nil)
