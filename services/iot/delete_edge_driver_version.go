package iot

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

// DeleteEdgeDriverVersion invokes the iot.DeleteEdgeDriverVersion API synchronously
// api document: https://help.aliyun.com/api/iot/deleteedgedriverversion.html
func (client *Client) DeleteEdgeDriverVersion(request *DeleteEdgeDriverVersionRequest) (response *DeleteEdgeDriverVersionResponse, err error) {
	response = CreateDeleteEdgeDriverVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEdgeDriverVersionWithChan invokes the iot.DeleteEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEdgeDriverVersionWithChan(request *DeleteEdgeDriverVersionRequest) (<-chan *DeleteEdgeDriverVersionResponse, <-chan error) {
	responseChan := make(chan *DeleteEdgeDriverVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEdgeDriverVersion(request)
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

// DeleteEdgeDriverVersionWithCallback invokes the iot.DeleteEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/deleteedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEdgeDriverVersionWithCallback(request *DeleteEdgeDriverVersionRequest, callback func(response *DeleteEdgeDriverVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEdgeDriverVersionResponse
		var err error
		defer close(result)
		response, err = client.DeleteEdgeDriverVersion(request)
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

// DeleteEdgeDriverVersionRequest is the request struct for api DeleteEdgeDriverVersion
type DeleteEdgeDriverVersionRequest struct {
	*requests.RpcRequest
	DriverId      string `position:"Query" name:"DriverId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DriverVersion string `position:"Query" name:"DriverVersion"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteEdgeDriverVersionResponse is the response struct for api DeleteEdgeDriverVersion
type DeleteEdgeDriverVersionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteEdgeDriverVersionRequest creates a request to invoke DeleteEdgeDriverVersion API
func CreateDeleteEdgeDriverVersionRequest() (request *DeleteEdgeDriverVersionRequest) {
	request = &DeleteEdgeDriverVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteEdgeDriverVersion", "iot", "openAPI")
	return
}

// CreateDeleteEdgeDriverVersionResponse creates a response to parse from DeleteEdgeDriverVersion response
func CreateDeleteEdgeDriverVersionResponse() (response *DeleteEdgeDriverVersionResponse) {
	response = &DeleteEdgeDriverVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}