# IUpdatePostRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the post. | 
**Content** | **string** | Content of the post. | 
**Thumbnail** | **string** | URL of the post thumbnail. | 

## Methods

### NewIUpdatePostRequestDTO

`func NewIUpdatePostRequestDTO(title string, content string, thumbnail string, ) *IUpdatePostRequestDTO`

NewIUpdatePostRequestDTO instantiates a new IUpdatePostRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIUpdatePostRequestDTOWithDefaults

`func NewIUpdatePostRequestDTOWithDefaults() *IUpdatePostRequestDTO`

NewIUpdatePostRequestDTOWithDefaults instantiates a new IUpdatePostRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IUpdatePostRequestDTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IUpdatePostRequestDTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IUpdatePostRequestDTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *IUpdatePostRequestDTO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IUpdatePostRequestDTO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IUpdatePostRequestDTO) SetContent(v string)`

SetContent sets Content field to given value.


### GetThumbnail

`func (o *IUpdatePostRequestDTO) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *IUpdatePostRequestDTO) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *IUpdatePostRequestDTO) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

