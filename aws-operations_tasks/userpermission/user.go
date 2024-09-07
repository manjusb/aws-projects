package userpermission

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

// GetUsers retrieves all IAM users in the account.
func GetUsers(svc *iam.IAM) ([]*iam.User, error) {
	var users []*iam.User
	input := &iam.ListUsersInput{}

	err := svc.ListUsersPages(input, func(page *iam.ListUsersOutput, lastPage bool) bool {
		users = append(users, page.Users...)
		return !lastPage
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserPermissions retrieves permissions for a specific IAM user.
func GetUserPermissions(svc *iam.IAM, userName string) (string, error) {
	var permissions []string

	// Get the user's attached policies
	policies, err := svc.ListAttachedUserPolicies(&iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		return "", err
	}
	for _, policy := range policies.AttachedPolicies {
		permissions = append(permissions, *policy.PolicyName)
	}

	// Get the user's inline policies
	inlinePolicies, err := svc.ListUserPolicies(&iam.ListUserPoliciesInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		return "", err
	}
	for _, policyName := range inlinePolicies.PolicyNames {
		permissions = append(permissions, *policyName)
	}

	return strings.Join(permissions, ", "), nil
}
