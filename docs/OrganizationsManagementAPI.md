# \OrganizationsManagementAPI

All URIs are relative to *https://organizations.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](OrganizationsManagementAPI.md#CreateAccessKey) | **Post** /v1/organizations/{orgId}/accessKey | Create access key for organization.
[**CreateOrganization**](OrganizationsManagementAPI.md#CreateOrganization) | **Post** /v1/organizations | Create a new organization.
[**DeleteOrg**](OrganizationsManagementAPI.md#DeleteOrg) | **Delete** /v1/organizations/{orgId} | Deactivate an organization.
[**DelinkOrg**](OrganizationsManagementAPI.md#DelinkOrg) | **Put** /v1/organizations/delink | Delink a child Org.
[**GetAllocatedCredits**](OrganizationsManagementAPI.md#GetAllocatedCredits) | **Get** /v1/organizations/allocatedCredits | Get allocated credits.
[**GetDeployments**](OrganizationsManagementAPI.md#GetDeployments) | **Get** /v1/deployments | List available deployments.
[**GetOrgUsage**](OrganizationsManagementAPI.md#GetOrgUsage) | **Get** /v1/organizations/usages/{orgId} | Get detailed usage of an organization.
[**GetOrganization**](OrganizationsManagementAPI.md#GetOrganization) | **Get** /v1/organizations/{orgId} | Get an organization&#39;s details.
[**GetParentOrgDetails**](OrganizationsManagementAPI.md#GetParentOrgDetails) | **Get** /v1/organizations/usages | Get usage details.
[**GetParentOrgInfo**](OrganizationsManagementAPI.md#GetParentOrgInfo) | **Get** /v1/organizations/parentOrg | Get parent organization information.
[**GetProvisioning**](OrganizationsManagementAPI.md#GetProvisioning) | **Get** /v1/organizations/provisioning/{orgId} | Get provisioning status for a child organization.
[**GetSubdomainLoginUrl**](OrganizationsManagementAPI.md#GetSubdomainLoginUrl) | **Get** /v1/organizations/{orgId}/subdomainLoginUrl | Get an organization&#39;s subdomain login URL.
[**GetUsages**](OrganizationsManagementAPI.md#GetUsages) | **Post** /v1/organizations/usages | Get credits usages for a list of organizations.
[**ListOrganizations**](OrganizationsManagementAPI.md#ListOrganizations) | **Get** /v1/organizations | Get a list of organizations.
[**UpdateSubscription**](OrganizationsManagementAPI.md#UpdateSubscription) | **Put** /v1/organizations/{orgId} | Update an organization.



## CreateAccessKey

> AccessKey CreateAccessKey(ctx, orgId).ParentDeploymentId(parentDeploymentId).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()

Create access key for organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization for which the access ID and key pair is required.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	accessKeyCreateRequest := *openapiclient.NewAccessKeyCreateRequest("automation access key") // AccessKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.CreateAccessKey(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.CreateAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessKey`: AccessKey
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.CreateAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization for which the access ID and key pair is required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **accessKeyCreateRequest** | [**AccessKeyCreateRequest**](AccessKeyCreateRequest.md) |  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> ReadOrganizationResponse CreateOrganization(ctx).ParentDeploymentId(parentDeploymentId).OrganizationWithSubscriptionDetails(organizationWithSubscriptionDetails).Execute()

Create a new organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	organizationWithSubscriptionDetails := *openapiclient.NewOrganizationWithSubscriptionDetails("sumo@example.com", "SumoLogic", "Jane", "us2") // OrganizationWithSubscriptionDetails | Details about the organization to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.CreateOrganization(context.Background()).ParentDeploymentId(parentDeploymentId).OrganizationWithSubscriptionDetails(organizationWithSubscriptionDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: ReadOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **organizationWithSubscriptionDetails** | [**OrganizationWithSubscriptionDetails**](OrganizationWithSubscriptionDetails.md) | Details about the organization to create. | 

### Return type

[**ReadOrganizationResponse**](ReadOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrg

> ReadOrganizationResponse DeleteOrg(ctx, orgId).ParentDeploymentId(parentDeploymentId).Execute()

Deactivate an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization to deactivate.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.DeleteOrg(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.DeleteOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrg`: ReadOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.DeleteOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization to deactivate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**ReadOrganizationResponse**](ReadOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelinkOrg

> string DelinkOrg(ctx).ParentDeploymentId(parentDeploymentId).DelinkChildOrgsRequest(delinkChildOrgsRequest).Execute()

Delink a child Org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	delinkChildOrgsRequest := *openapiclient.NewDelinkChildOrgsRequest([]string{"us2-00000000FF42A0C3"}) // DelinkChildOrgsRequest | Identifier of organizations to delink.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.DelinkOrg(context.Background()).ParentDeploymentId(parentDeploymentId).DelinkChildOrgsRequest(delinkChildOrgsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.DelinkOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelinkOrg`: string
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.DelinkOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelinkOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **delinkChildOrgsRequest** | [**DelinkChildOrgsRequest**](DelinkChildOrgsRequest.md) | Identifier of organizations to delink. | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllocatedCredits

> AllocatedCredits GetAllocatedCredits(ctx).ParentDeploymentId(parentDeploymentId).Execute()

Get allocated credits.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetAllocatedCredits(context.Background()).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetAllocatedCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllocatedCredits`: AllocatedCredits
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetAllocatedCredits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllocatedCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**AllocatedCredits**](AllocatedCredits.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployments

> []Deployment GetDeployments(ctx).ParentDeploymentId(parentDeploymentId).Execute()

List available deployments.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetDeployments(context.Background()).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployments`: []Deployment
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgUsage

> DetailedUsage GetOrgUsage(ctx, orgId).ParentDeploymentId(parentDeploymentId).Execute()

Get detailed usage of an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization for which the details are required.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetOrgUsage(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetOrgUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgUsage`: DetailedUsage
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetOrgUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization for which the details are required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**DetailedUsage**](DetailedUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> ReadOrganizationResponse GetOrganization(ctx, orgId).ParentDeploymentId(parentDeploymentId).Execute()

Get an organization's details.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization for which the details are required.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetOrganization(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: ReadOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization for which the details are required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**ReadOrganizationResponse**](ReadOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentOrgDetails

> ParentUsage GetParentOrgDetails(ctx).ParentDeploymentId(parentDeploymentId).Execute()

Get usage details.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetParentOrgDetails(context.Background()).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetParentOrgDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentOrgDetails`: ParentUsage
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetParentOrgDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParentOrgDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**ParentUsage**](ParentUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentOrgInfo

> ParentOrgInfo GetParentOrgInfo(ctx).ParentDeploymentId(parentDeploymentId).Execute()

Get parent organization information.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetParentOrgInfo(context.Background()).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetParentOrgInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentOrgInfo`: ParentOrgInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetParentOrgInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParentOrgInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**ParentOrgInfo**](ParentOrgInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioning

> Provisioning GetProvisioning(ctx, orgId).ParentDeploymentId(parentDeploymentId).Execute()

Get provisioning status for a child organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization for which the details are required.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetProvisioning(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetProvisioning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioning`: Provisioning
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization for which the details are required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**Provisioning**](Provisioning.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubdomainLoginUrl

> Subdomain GetSubdomainLoginUrl(ctx, orgId).ParentDeploymentId(parentDeploymentId).Execute()

Get an organization's subdomain login URL.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the child organization for which the loginUrl is required.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetSubdomainLoginUrl(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetSubdomainLoginUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubdomainLoginUrl`: Subdomain
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetSubdomainLoginUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the child organization for which the loginUrl is required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubdomainLoginUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 

### Return type

[**Subdomain**](Subdomain.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsages

> ListUsagesResponse GetUsages(ctx).ParentDeploymentId(parentDeploymentId).RequestBody(requestBody).Limit(limit).Token(token).Execute()

Get credits usages for a list of organizations.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	requestBody := []string{"us2-00000000FF42A0C3"} // []string | List of the organizations for which usage needs to be fetched.
	limit := int32(56) // int32 | Limit the number of results returned in the response. The number of results returned may be less than the `limit`. (optional) (default to 100)
	token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.GetUsages(context.Background()).ParentDeploymentId(parentDeploymentId).RequestBody(requestBody).Limit(limit).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.GetUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsages`: ListUsagesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.GetUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **requestBody** | **[]string** | List of the organizations for which usage needs to be fetched. | 
 **limit** | **int32** | Limit the number of results returned in the response. The number of results returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListUsagesResponse**](ListUsagesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> ListOrganizationResponse ListOrganizations(ctx).ParentDeploymentId(parentDeploymentId).Limit(limit).Token(token).Status(status).Execute()

Get a list of organizations.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	limit := int32(56) // int32 | Limit the number of organizations returned in the response. The number of organizations returned may be less than the `limit`. (optional) (default to 100)
	token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
	status := "Active" // string | Status of an organization, based on its subscription. Valid values are 'Active', 'Inactive', and 'All'. By default, only active organizations are listed. (optional) (default to "Active")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.ListOrganizations(context.Background()).ParentDeploymentId(parentDeploymentId).Limit(limit).Token(token).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.ListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizations`: ListOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **limit** | **int32** | Limit the number of organizations returned in the response. The number of organizations returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **status** | **string** | Status of an organization, based on its subscription. Valid values are &#39;Active&#39;, &#39;Inactive&#39;, and &#39;All&#39;. By default, only active organizations are listed. | [default to &quot;Active&quot;]

### Return type

[**ListOrganizationResponse**](ListOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> ReadOrganizationResponse UpdateSubscription(ctx, orgId).ParentDeploymentId(parentDeploymentId).Baselines(baselines).Execute()

Update an organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "orgId_example" // string | Identifier of the organization to update.
	parentDeploymentId := "us2" // string | Deployment on which the calling organization resides.
	baselines := *openapiclient.NewBaselines() // Baselines | The utilization baselines for the organization and the updated credits allocation. For organizations on Trial/Free plans, reallocating credits will upgrade their plan type to your plan. The plan change cannot be rolled back.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsManagementAPI.UpdateSubscription(context.Background(), orgId).ParentDeploymentId(parentDeploymentId).Baselines(baselines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsManagementAPI.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: ReadOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsManagementAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Identifier of the organization to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentDeploymentId** | **string** | Deployment on which the calling organization resides. | 
 **baselines** | [**Baselines**](Baselines.md) | The utilization baselines for the organization and the updated credits allocation. For organizations on Trial/Free plans, reallocating credits will upgrade their plan type to your plan. The plan change cannot be rolled back. | 

### Return type

[**ReadOrganizationResponse**](ReadOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

