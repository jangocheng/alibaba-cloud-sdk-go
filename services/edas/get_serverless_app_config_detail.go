package edas

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

// GetServerlessAppConfigDetail invokes the edas.GetServerlessAppConfigDetail API synchronously
// api document: https://help.aliyun.com/api/edas/getserverlessappconfigdetail.html
func (client *Client) GetServerlessAppConfigDetail(request *GetServerlessAppConfigDetailRequest) (response *GetServerlessAppConfigDetailResponse, err error) {
	response = CreateGetServerlessAppConfigDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetServerlessAppConfigDetailWithChan invokes the edas.GetServerlessAppConfigDetail API asynchronously
// api document: https://help.aliyun.com/api/edas/getserverlessappconfigdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServerlessAppConfigDetailWithChan(request *GetServerlessAppConfigDetailRequest) (<-chan *GetServerlessAppConfigDetailResponse, <-chan error) {
	responseChan := make(chan *GetServerlessAppConfigDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServerlessAppConfigDetail(request)
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

// GetServerlessAppConfigDetailWithCallback invokes the edas.GetServerlessAppConfigDetail API asynchronously
// api document: https://help.aliyun.com/api/edas/getserverlessappconfigdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServerlessAppConfigDetailWithCallback(request *GetServerlessAppConfigDetailRequest, callback func(response *GetServerlessAppConfigDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServerlessAppConfigDetailResponse
		var err error
		defer close(result)
		response, err = client.GetServerlessAppConfigDetail(request)
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

// GetServerlessAppConfigDetailRequest is the request struct for api GetServerlessAppConfigDetail
type GetServerlessAppConfigDetailRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// GetServerlessAppConfigDetailResponse is the response struct for api GetServerlessAppConfigDetail
type GetServerlessAppConfigDetailResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateGetServerlessAppConfigDetailRequest creates a request to invoke GetServerlessAppConfigDetail API
func CreateGetServerlessAppConfigDetailRequest() (request *GetServerlessAppConfigDetailRequest) {
	request = &GetServerlessAppConfigDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetServerlessAppConfigDetail", "/pop/v5/k8s/pop/pop_serverless_app_config_detail", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetServerlessAppConfigDetailResponse creates a response to parse from GetServerlessAppConfigDetail response
func CreateGetServerlessAppConfigDetailResponse() (response *GetServerlessAppConfigDetailResponse) {
	response = &GetServerlessAppConfigDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
