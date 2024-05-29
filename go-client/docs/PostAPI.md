# \PostAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](PostAPI.md#CreatePost) | **Post** /chrysalis/post | 
[**DeletePostByID**](PostAPI.md#DeletePostByID) | **Delete** /chrysalis/post/{id} | 
[**GetPostByID**](PostAPI.md#GetPostByID) | **Get** /chrysalis/post/{id} | 
[**UpdatePostByID**](PostAPI.md#UpdatePostByID) | **Patch** /chrysalis/post/{id} | 



## CreatePost

> CreatePost(ctx).ICreatePostRequestDTO(iCreatePostRequestDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bonjuruu/chrysalis-service-client"
)

func main() {
	iCreatePostRequestDTO := *openapiclient.NewICreatePostRequestDTO("Author_example", "Title_example", "Content_example", "Thumbnail_example", "CreatedAt_example", "UpdatedAt_example") // ICreatePostRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostAPI.CreatePost(context.Background()).ICreatePostRequestDTO(iCreatePostRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.CreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iCreatePostRequestDTO** | [**ICreatePostRequestDTO**](ICreatePostRequestDTO.md) |  | 

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


## DeletePostByID

> DeletePostByID(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bonjuruu/chrysalis-service-client"
)

func main() {
	id := "id_example" // string | Post id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostAPI.DeletePostByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.DeletePostByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePostByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostByID

> IPostDTO GetPostByID(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bonjuruu/chrysalis-service-client"
)

func main() {
	id := "id_example" // string | Post id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostAPI.GetPostByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.GetPostByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostByID`: IPostDTO
	fmt.Fprintf(os.Stdout, "Response from `PostAPI.GetPostByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPostDTO**](IPostDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostByID

> UpdatePostByID(ctx, id).IUpdatePostRequestDTO(iUpdatePostRequestDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bonjuruu/chrysalis-service-client"
)

func main() {
	id := "id_example" // string | Post id.
	iUpdatePostRequestDTO := *openapiclient.NewIUpdatePostRequestDTO("Title_example", "Content_example", "Thumbnail_example") // IUpdatePostRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostAPI.UpdatePostByID(context.Background(), id).IUpdatePostRequestDTO(iUpdatePostRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostAPI.UpdatePostByID``: %v\n", err)
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

