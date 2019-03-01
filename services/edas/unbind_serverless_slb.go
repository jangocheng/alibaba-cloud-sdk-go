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

// UnbindServerlessSlb invokes the edas.UnbindServerlessSlb API synchronously
// api document: https://help.aliyun.com/api/edas/unbindserverlessslb.html
func (client *Client) UnbindServerlessSlb(request *UnbindServerlessSlbRequest) (response *UnbindServerlessSlbResponse, err error) {
	response = CreateUnbindServerlessSlbResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindServerlessSlbWithChan invokes the edas.UnbindServerlessSlb API asynchronously
// api document: https://help.aliyun.com/api/edas/unbindserverlessslb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindServerlessSlbWithChan(request *UnbindServerlessSlbRequest) (<-chan *UnbindServerlessSlbResponse, <-chan error) {
	responseChan := make(chan *UnbindServerlessSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindServerlessSlb(request)
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

// UnbindServerlessSlbWithCallback invokes the edas.UnbindServerlessSlb API asynchronously
// api document: https://help.aliyun.com/api/edas/unbindserverlessslb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindServerlessSlbWithCallback(request *UnbindServerlessSlbRequest, callback func(response *UnbindServerlessSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindServerlessSlbResponse
		var err error
		defer close(result)
		response, err = client.UnbindServerlessSlb(request)
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

// UnbindServerlessSlbRequest is the request struct for api UnbindServerlessSlb
type UnbindServerlessSlbRequest struct {
	*requests.RoaRequest
	Intranet requests.Boolean `position:"Query" name:"Intranet"`
	AppId    string           `position:"Query" name:"AppId"`
	Internet requests.Boolean `position:"Query" name:"Internet"`
}

// UnbindServerlessSlbResponse is the response struct for api UnbindServerlessSlb
type UnbindServerlessSlbResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
}

// CreateUnbindServerlessSlbRequest creates a request to invoke UnbindServerlessSlb API
func CreateUnbindServerlessSlbRequest() (request *UnbindServerlessSlbRequest) {
	request = &UnbindServerlessSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UnbindServerlessSlb", "/pop/v5/k8s/acs/serverless_slb_binding", "edas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateUnbindServerlessSlbResponse creates a response to parse from UnbindServerlessSlb response
func CreateUnbindServerlessSlbResponse() (response *UnbindServerlessSlbResponse) {
	response = &UnbindServerlessSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
