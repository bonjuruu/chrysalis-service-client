# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePostByID**](DefaultAPI.md#UpdatePostByID) | **Patch** /chrysalis/post/{id} | 



## UpdatePostByID

> UpdatePostByID(ctx, id).IUpdatePostRequestDTO(iUpdatePostRequestDTO).Execute()





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
	id := "id_example" // string | Post id.
	iUpdatePostRequestDTO := *openapiclient.NewIUpdatePostRequestDTO("Title_example", "Content_example", "Thumbnail_example") // IUpdatePostRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdatePostByID(context.Background(), id).IUpdatePostRequestDTO(iUpdatePostRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePostByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iUpdatePostRequestDTO** | [**IUpdatePostRequestDTO**](IUpdatePostRequestDTO.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

