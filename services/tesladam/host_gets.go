package tesladam

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

func (client *Client) HostGets(request *HostGetsRequest) (response *HostGetsResponse, err error) {
	response = CreateHostGetsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) HostGetsWithChan(request *HostGetsRequest) (<-chan *HostGetsResponse, <-chan error) {
	responseChan := make(chan *HostGetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HostGets(request)
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

func (client *Client) HostGetsWithCallback(request *HostGetsRequest, callback func(response *HostGetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HostGetsResponse
		var err error
		defer close(result)
		response, err = client.HostGets(request)
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

type HostGetsRequest struct {
	*requests.RpcRequest
	QueryType string           `position:"Query" name:"QueryType"`
	Query     string           `position:"Query" name:"Query"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
}

type HostGetsResponse struct {
	*responses.BaseResponse
	Status  bool             `json:"Status" xml:"Status"`
	Message string           `json:"Message" xml:"Message"`
	Data    []DataInHostGets `json:"Data" xml:"Data"`
}

func CreateHostGetsRequest() (request *HostGetsRequest) {
	request = &HostGetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaDam", "2018-01-18", "HostGets", "", "")
	return
}

func CreateHostGetsResponse() (response *HostGetsResponse) {
	response = &HostGetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
