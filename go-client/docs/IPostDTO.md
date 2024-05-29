# IPostDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Post id. | 
**Author** | **string** | Name of the author. | 
**Title** | **string** | Title of the post. | 
**Content** | **string** | Content of the post. | 
**Thumbnail** | **string** | URL of the post thumbnail. | 
**CreatedAt** | **string** | Post creation date. | 
**UpdatedAt** | **string** | Post update date. | 

## Methods

### NewIPostDTO

`func NewIPostDTO(id string, author string, title string, content string, thumbnail string, createdAt string, updatedAt string, ) *IPostDTO`

NewIPostDTO instantiates a new IPostDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPostDTOWithDefaults

`func NewIPostDTOWithDefaults() *IPostDTO`

NewIPostDTOWithDefaults instantiates a new IPostDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPostDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPostDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPostDTO) SetId(v string)`

SetId sets Id field to given value.


### GetAuthor

`func (o *IPostDTO) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *IPostDTO) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *IPostDTO) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetTitle

`func (o *IPostDTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IPostDTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IPostDTO) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *IPostDTO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IPostDTO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IPostDTO) SetContent(v string)`

SetContent sets Content field to given value.


### GetThumbnail

`func (o *IPostDTO) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *IPostDTO) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *IPostDTO) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.


### GetCreatedAt

`func (o *IPostDTO) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPostDTO) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPostDTO) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IPostDTO) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPostDTO) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPostDTO) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


