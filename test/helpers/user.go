package helpers

import (
	"fmt"
	"strings"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/aws/cluster-api-provider-cloudstack/pkg/cloud"
)

// GetDomainByPath fetches a domain by its path.
func GetDomainByPath(csClient *cloudstack.CloudStackClient, path string) (string, error, bool) {
	// Split path and get name.
	path = strings.Trim(path, "/")
	tokens := []string{}
	tokens = strings.Split(path, "/")

	// Ensure the path begins with ROOT.
	if !strings.EqualFold(tokens[0], "ROOT") {
		tokens = append([]string{"ROOT"}, tokens...)
	} else {
		tokens[0] = "ROOT"
	}
	path = strings.Join(tokens, "/")

	// Set present search/list parameters.
	p := csClient.Domain.NewListDomainsParams()
	p.SetListall(true)

	// If path was provided also use level narrow the search for domain.
	if level := len(tokens) - 1; level >= 0 {
		p.SetLevel(level)
	}

	if resp, err := csClient.Domain.ListDomains(p); err != nil {
		return "", err, false
	} else {
		for _, domain := range resp.Domains {
			if domain.Path == path {
				return domain.Id, nil, true
			}
		}
	}

	return "", nil, false
}

// CreateDomainUnderParent creates a domain as a sub-domain of the passed parent.
func CreateDomainUnderParent(csClient *cloudstack.CloudStackClient, parentID string, domainName string) (string, error) {
	p := csClient.Domain.NewCreateDomainParams(domainName)
	p.SetParentdomainid(parentID)
	resp, err := csClient.Domain.CreateDomain(p)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

// GetOrCreateDomain gets or creates a domain as specified in the passed domain object.
func GetOrCreateDomain(domain *cloud.Domain, csClient *cloudstack.CloudStackClient) error {
	// Split the specified domain path and prepend ROOT/ if it's missing.
	domain.Path = strings.Trim(domain.Path, "/")
	tokens := strings.Split(domain.Path, "/")
	if strings.EqualFold(tokens[0], "root") {
		tokens[0] = "ROOT"
	} else {
		tokens = append([]string{"ROOT"}, tokens...)
	}
	domain.Path = strings.Join(tokens, "/")

	// Fetch ROOT domain ID.
	rootID, err, _ := GetDomainByPath(csClient, "ROOT")
	if err != nil {
		return err
	}

	// Iteratively create the domain from its path.
	parentID := rootID
	currPath := "ROOT"
	for _, nextDomainName := range tokens[1:] {
		currPath = currPath + "/" + nextDomainName
		if nextId, err, found := GetDomainByPath(csClient, currPath); err != nil {
			return err
		} else if !found {
			if nextId, err := CreateDomainUnderParent(csClient, parentID, nextDomainName); err != nil {
				return err
			} else {
				parentID = nextId
			}
		} else {
			parentID = nextId
		}
	}
	domain.ID = parentID
	domain.Name = tokens[len(tokens)-1]
	domain.Path = strings.Join(tokens, "/")
	return nil
}

// DeleteDomain deletes a domain by ID.
func DeleteDomain(csClient *cloudstack.CloudStackClient, domainID string) error {
	p := csClient.Domain.NewDeleteDomainParams(domainID)
	p.SetCleanup(true)
	resp, err := csClient.Domain.DeleteDomain(p)
	if !resp.Success {
		return fmt.Errorf("unsuccessful deletion of domain with ID %s", domainID)
	}
	return err
}

// // CreateAccount creates a domain as specified in the passed account object.
// func CreateAccount(account *cloud.Account, csClient *cloudstack.CloudStackClient) error {
// 	return nil
// }

// // CreateUser creates a domain as specified in the passed account object.
// func CreateUser(user *cloud.User, csClient *cloudstack.CloudStackClient) error {
// 	return nil
// }
