package vpc

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

// VSwitch is a nested struct in vpc response
type VSwitch struct {
	VSwitchId               string                  `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                   string                  `json:"VpcId" xml:"VpcId"`
	Status                  string                  `json:"Status" xml:"Status"`
	CidrBlock               string                  `json:"CidrBlock" xml:"CidrBlock"`
	Ipv6CidrBlock           string                  `json:"Ipv6CidrBlock" xml:"Ipv6CidrBlock"`
	ZoneId                  string                  `json:"ZoneId" xml:"ZoneId"`
	AvailableIpAddressCount int                     `json:"AvailableIpAddressCount" xml:"AvailableIpAddressCount"`
	Description             string                  `json:"Description" xml:"Description"`
	VSwitchName             string                  `json:"VSwitchName" xml:"VSwitchName"`
	CreationTime            string                  `json:"CreationTime" xml:"CreationTime"`
	IsDefault               bool                    `json:"IsDefault" xml:"IsDefault"`
	ResourceGroupId         string                  `json:"ResourceGroupId" xml:"ResourceGroupId"`
	NetworkAclId            string                  `json:"NetworkAclId" xml:"NetworkAclId"`
	RouteTable              RouteTable              `json:"RouteTable" xml:"RouteTable"`
	Tags                    TagsInDescribeVSwitches `json:"Tags" xml:"Tags"`
}
