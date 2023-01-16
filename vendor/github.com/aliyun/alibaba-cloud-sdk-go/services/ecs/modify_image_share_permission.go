package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyImageSharePermission invokes the ecs.ModifyImageSharePermission API synchronously
func (client *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
	response = CreateModifyImageSharePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyImageSharePermissionWithChan invokes the ecs.ModifyImageSharePermission API asynchronously
func (client *Client) ModifyImageSharePermissionWithChan(request *ModifyImageSharePermissionRequest) (<-chan *ModifyImageSharePermissionResponse, <-chan error) {
	responseChan := make(chan *ModifyImageSharePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyImageSharePermission(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyImageSharePermissionWithCallback invokes the ecs.ModifyImageSharePermission API asynchronously
func (client *Client) ModifyImageSharePermissionWithCallback(request *ModifyImageSharePermissionRequest, callback func(response *ModifyImageSharePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyImageSharePermissionResponse
		var err error
		defer close(result)
		response, err = client.ModifyImageSharePermission(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyImageSharePermissionRequest is the request struct for api ModifyImageSharePermission
type ModifyImageSharePermissionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId              string           `position:"Query" name:"ImageId"`
	IsPublic             requests.Boolean `position:"Query" name:"IsPublic"`
	LaunchPermission     string           `position:"Query" name:"LaunchPermission"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AddAccount           *[]string        `position:"Query" name:"AddAccount"  type:"Repeated"`
	RemoveAccount        *[]string        `position:"Query" name:"RemoveAccount"  type:"Repeated"`
}

// ModifyImageSharePermissionResponse is the response struct for api ModifyImageSharePermission
type ModifyImageSharePermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyImageSharePermissionRequest creates a request to invoke ModifyImageSharePermission API
func CreateModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
	request = &ModifyImageSharePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageSharePermission", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyImageSharePermissionResponse creates a response to parse from ModifyImageSharePermission response
func CreateModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
	response = &ModifyImageSharePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}