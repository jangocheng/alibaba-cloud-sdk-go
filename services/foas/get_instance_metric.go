package foas

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

// GetInstanceMetric invokes the foas.GetInstanceMetric API synchronously
// api document: https://help.aliyun.com/api/foas/getinstancemetric.html
func (client *Client) GetInstanceMetric(request *GetInstanceMetricRequest) (response *GetInstanceMetricResponse, err error) {
	response = CreateGetInstanceMetricResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceMetricWithChan invokes the foas.GetInstanceMetric API asynchronously
// api document: https://help.aliyun.com/api/foas/getinstancemetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceMetricWithChan(request *GetInstanceMetricRequest) (<-chan *GetInstanceMetricResponse, <-chan error) {
	responseChan := make(chan *GetInstanceMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceMetric(request)
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

// GetInstanceMetricWithCallback invokes the foas.GetInstanceMetric API asynchronously
// api document: https://help.aliyun.com/api/foas/getinstancemetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceMetricWithCallback(request *GetInstanceMetricRequest, callback func(response *GetInstanceMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceMetricResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceMetric(request)
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

// GetInstanceMetricRequest is the request struct for api GetInstanceMetric
type GetInstanceMetricRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Query" name:"instanceId"`
	MetricJson  string           `position:"Body" name:"metricJson"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceMetricResponse is the response struct for api GetInstanceMetric
type GetInstanceMetricResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Metrics   MetricsInGetInstanceMetric `json:"Metrics" xml:"Metrics"`
}

// CreateGetInstanceMetricRequest creates a request to invoke GetInstanceMetric API
func CreateGetInstanceMetricRequest() (request *GetInstanceMetricRequest) {
	request = &GetInstanceMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceMetric", "/api/v2/projects/[projectName]/jobs/[jobName]/metric", "foas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInstanceMetricResponse creates a response to parse from GetInstanceMetric response
func CreateGetInstanceMetricResponse() (response *GetInstanceMetricResponse) {
	response = &GetInstanceMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
