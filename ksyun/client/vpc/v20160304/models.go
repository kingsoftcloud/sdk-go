package v20160304

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeSubnetsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeRoutesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNetworkAclsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeSecurityGroupsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNatsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNatsTagKV struct {
	Name  *string `json:"Name,omitempty" name:"Name"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type DescribeInternetGatewaysFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeVpcPeeringConnectionsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNetworkInterfacesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeSubnetAvailableAddressesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeDirectConnectGatewaysFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeVpnGatewaysFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeVpnTunnelsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeSubnetAllocatedIpAddressesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeDirectConnectRoutesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeRouteTablesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeNatRateLimitFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeHaVipFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeDirectConnectGatewayRouteFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeVpnGatewayRoutesFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type DescribeFlowLogsFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type CreateVpcRequest struct {
	*ksyunhttp.BaseRequest
	VpcName               *string `json:"VpcName,omitempty" name:"VpcName"`
	CidrBlock             *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	ProvidedIpv6CidrBlock *bool   `json:"ProvidedIpv6CidrBlock,omitempty" name:"ProvidedIpv6CidrBlock"`
}

func (r *CreateVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcRequest struct {
	*ksyunhttp.BaseRequest
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsRequest struct {
	*ksyunhttp.BaseRequest
	VpcId      []*string `json:"VpcId,omitempty" name:"VpcId"`
	MaxResults *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	VpcSet    []struct {
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		VpcId                       *string `json:"VpcId" name:"VpcId"`
		VpcName                     *string `json:"VpcName" name:"VpcName"`
		CidrBlock                   *string `json:"CidrBlock" name:"CidrBlock"`
		IsDefault                   *bool   `json:"IsDefault" name:"IsDefault"`
		ProvidedIpv6CidrBlock       *bool   `json:"ProvidedIpv6CidrBlock" name:"ProvidedIpv6CidrBlock"`
		CenId                       *string `json:"CenId" name:"CenId"`
		Ipv6CidrBlockAssociationSet []struct {
			Ipv6CidrBlock *string `json:"Ipv6CidrBlock" name:"Ipv6CidrBlock"`
		} `json:"Ipv6CidrBlockAssociationSet" name:"Ipv6CidrBlockAssociationSet"`
		SecondaryCidrSet []struct {
			SecondaryCidrId *string `json:"SecondaryCidrId" name:"SecondaryCidrId"`
			Cidr            *string `json:"Cidr" name:"Cidr"`
			Type            *string `json:"Type" name:"Type"`
		} `json:"SecondaryCidrSet" name:"SecondaryCidrSet"`
	} `json:"VpcSet"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                 *string `json:"VpcId,omitempty" name:"VpcId"`
	SubnetName            *string `json:"SubnetName,omitempty" name:"SubnetName"`
	CidrBlock             *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	ProvidedIpv6CidrBlock *bool   `json:"ProvidedIpv6CidrBlock,omitempty" name:"ProvidedIpv6CidrBlock"`
	SubnetType            *string `json:"SubnetType,omitempty" name:"SubnetType"`
	DhcpIpFrom            *string `json:"DhcpIpFrom,omitempty" name:"DhcpIpFrom"`
	DhcpIpTo              *string `json:"DhcpIpTo,omitempty" name:"DhcpIpTo"`
	Dns1                  *string `json:"Dns1,omitempty" name:"Dns1"`
	Dns2                  *string `json:"Dns2,omitempty" name:"Dns2"`
	GatewayIp             *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
	SecondaryCidrId       *string `json:"SecondaryCidrId,omitempty" name:"SecondaryCidrId"`
	AvailabilityZone      *string `json:"AvailabilityZone,omitempty" name:"AvailabilityZone"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId   []*string                `json:"SubnetId,omitempty" name:"SubnetId"`
	Filter     []*DescribeSubnetsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                     `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                  `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	SubnetSet []struct {
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		SubnetId                    *string `json:"SubnetId" name:"SubnetId"`
		VpcId                       *string `json:"VpcId" name:"VpcId"`
		SubnetName                  *string `json:"SubnetName" name:"SubnetName"`
		CidrBlock                   *string `json:"CidrBlock" name:"CidrBlock"`
		SubnetType                  *string `json:"SubnetType" name:"SubnetType"`
		DhcpIpFrom                  *string `json:"DhcpIpFrom" name:"DhcpIpFrom"`
		DhcpIpTo                    *string `json:"DhcpIpTo" name:"DhcpIpTo"`
		Dns1                        *string `json:"Dns1" name:"Dns1"`
		Dns2                        *string `json:"Dns2" name:"Dns2"`
		GatewayIp                   *string `json:"GatewayIp" name:"GatewayIp"`
		AvailabilityZoneName        *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		ProvidedIpv6CidrBlock       *bool   `json:"ProvidedIpv6CidrBlock" name:"ProvidedIpv6CidrBlock"`
		SecondaryCidrId             *string `json:"SecondaryCidrId" name:"SecondaryCidrId"`
		AvailableIpNumber           *string `json:"AvailableIpNumber" name:"AvailableIpNumber"`
		NetworkAclId                *string `json:"NetworkAclId" name:"NetworkAclId"`
		NatId                       *string `json:"NatId" name:"NatId"`
		RouteTableId                *string `json:"RouteTableId" name:"RouteTableId"`
		Ipv6CidrBlockAssociationSet []struct {
			Ipv6CidrBlock *string `json:"Ipv6CidrBlock" name:"Ipv6CidrBlock"`
		} `json:"Ipv6CidrBlockAssociationSet" name:"Ipv6CidrBlockAssociationSet"`
	} `json:"SubnetSet"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkAclRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
}

func (r *AssociateNetworkAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNetworkAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateNetworkAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkAclRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
}

func (r *DisassociateNetworkAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNetworkAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DisassociateNetworkAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                         *string `json:"VpcId,omitempty" name:"VpcId"`
	RouteType                     *string `json:"RouteType,omitempty" name:"RouteType"`
	DestinationCidrBlock          *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	InstanceId                    *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcPeeringConnectionId        *string `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
	DirectConnectGatewayId        *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	VpnTunnelId                   *string `json:"VpnTunnelId,omitempty" name:"VpnTunnelId"`
	VpnGatewayId                  *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	NetworkInterfaceId            *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	HaVipId                       *string `json:"HaVipId,omitempty" name:"HaVipId"`
	HaVipMasterNetworkInterfaceId *string `json:"HaVipMasterNetworkInterfaceId,omitempty" name:"HaVipMasterNetworkInterfaceId"`
	CenId                         *string `json:"CenId,omitempty" name:"CenId"`
	Description                   *string `json:"Description,omitempty" name:"Description"`
	RouteTableId                  *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *CreateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	RouteId              *string `json:"RouteId" name:"RouteId"`
	RouteType            *string `json:"RouteType" name:"RouteType"`
	CreateTime           *string `json:"CreateTime" name:"CreateTime"`
	VpcId                *string `json:"VpcId" name:"VpcId"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
	System               *bool   `json:"System" name:"System"`
	Description          *string `json:"Description" name:"Description"`
	RouteTableId         *string `json:"RouteTableId" name:"RouteTableId"`
	RouteEntryType       *string `json:"RouteEntryType" name:"RouteEntryType"`
}

func (r *CreateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRequest struct {
	*ksyunhttp.BaseRequest
	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutesRequest struct {
	*ksyunhttp.BaseRequest
	RouteId    []*string               `json:"RouteId,omitempty" name:"RouteId"`
	Filter     []*DescribeRoutesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                    `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                 `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRoutesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	RouteSet  []struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		VpcId                *string `json:"VpcId" name:"VpcId"`
		DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
		RouteId              *string `json:"RouteId" name:"RouteId"`
		RouteType            *string `json:"RouteType" name:"RouteType"`
		Status               *string `json:"Status" name:"Status"`
		System               *bool   `json:"System" name:"System"`
		RoutePublishStatus   *string `json:"RoutePublishStatus" name:"RoutePublishStatus"`
		Description          *string `json:"Description" name:"Description"`
		RouteTableId         *string `json:"RouteTableId" name:"RouteTableId"`
		RouteEntryType       *string `json:"RouteEntryType" name:"RouteEntryType"`
		NextHopSet           []struct {
			GatewayId          *string `json:"GatewayId" name:"GatewayId"`
			GatewayName        *string `json:"GatewayName" name:"GatewayName"`
			NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			Mac                *string `json:"Mac" name:"Mac"`
		} `json:"NextHopSet" name:"NextHopSet"`
	} `json:"RouteSet"`
}

func (r *DescribeRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclRequest struct {
	*ksyunhttp.BaseRequest
	VpcId          *string `json:"VpcId,omitempty" name:"VpcId"`
	NetworkAclName *string `json:"NetworkAclName,omitempty" name:"NetworkAclName"`
	Description    *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateNetworkAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CreateTime     *string `json:"CreateTime" name:"CreateTime"`
	VpcId          *string `json:"VpcId" name:"VpcId"`
	NetworkAclName *string `json:"NetworkAclName" name:"NetworkAclName"`
	NetworkAclId   *string `json:"NetworkAclId" name:"NetworkAclId"`
	Description    *string `json:"Description" name:"Description"`
}

func (r *CreateNetworkAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclRequest struct {
	*ksyunhttp.BaseRequest
	NetworkAclId *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
}

func (r *DeleteNetworkAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteNetworkAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclEntryRequest struct {
	*ksyunhttp.BaseRequest
	NetworkAclId  *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
	Direction     *string `json:"Direction,omitempty" name:"Direction"`
	RuleNumber    *int    `json:"RuleNumber,omitempty" name:"RuleNumber"`
	Protocol      *string `json:"Protocol,omitempty" name:"Protocol"`
	IcmpType      *int    `json:"IcmpType,omitempty" name:"IcmpType"`
	IcmpCode      *int    `json:"IcmpCode,omitempty" name:"IcmpCode"`
	RuleAction    *string `json:"RuleAction,omitempty" name:"RuleAction"`
	PortRangeFrom *int    `json:"PortRangeFrom,omitempty" name:"PortRangeFrom"`
	PortRangeTo   *int    `json:"PortRangeTo,omitempty" name:"PortRangeTo"`
	CidrBlock     *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	Description   *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateNetworkAclEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkAclEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNetworkAclEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkAclEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *CreateNetworkAclEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkAclEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclEntryRequest struct {
	*ksyunhttp.BaseRequest
	NetworkAclId      *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" name:"NetworkAclEntryId"`
}

func (r *DeleteNetworkAclEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkAclEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNetworkAclEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkAclEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteNetworkAclEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkAclEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkAclsRequest struct {
	*ksyunhttp.BaseRequest
	NetworkAclId []*string                    `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
	Filter       []*DescribeNetworkAclsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults   *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string                      `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeNetworkAclsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNetworkAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkAclsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	NetworkAclSet []struct {
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		VpcId              *string `json:"VpcId" name:"VpcId"`
		NetworkAclName     *string `json:"NetworkAclName" name:"NetworkAclName"`
		NetworkAclId       *string `json:"NetworkAclId" name:"NetworkAclId"`
		Description        *string `json:"Description" name:"Description"`
		NetworkAclEntrySet []struct {
			Description       *string `json:"Description" name:"Description"`
			NetworkAclId      *string `json:"NetworkAclId" name:"NetworkAclId"`
			NetworkAclEntryId *string `json:"NetworkAclEntryId" name:"NetworkAclEntryId"`
			CidrBlock         *string `json:"CidrBlock" name:"CidrBlock"`
			RuleNumber        *int    `json:"RuleNumber" name:"RuleNumber"`
			Direction         *string `json:"Direction" name:"Direction"`
			RuleAction        *string `json:"RuleAction" name:"RuleAction"`
			Protocol          *string `json:"Protocol" name:"Protocol"`
			IcmpType          *int    `json:"IcmpType" name:"IcmpType"`
			IcmpCode          *int    `json:"IcmpCode" name:"IcmpCode"`
			PortRangeFrom     *int    `json:"PortRangeFrom" name:"PortRangeFrom"`
			PortRangeTo       *int    `json:"PortRangeTo" name:"PortRangeTo"`
		} `json:"NetworkAclEntrySet" name:"NetworkAclEntrySet"`
	} `json:"NetworkAclSet"`
}

func (r *DescribeNetworkAclsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	VpcId             *string `json:"VpcId,omitempty" name:"VpcId"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthorizeSecurityGroupEntryRequest struct {
	*ksyunhttp.BaseRequest
	Description     *string `json:"Description,omitempty" name:"Description"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	CidrBlock       *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
	Direction       *string `json:"Direction,omitempty" name:"Direction"`
	Protocol        *string `json:"Protocol,omitempty" name:"Protocol"`
	IcmpType        *int    `json:"IcmpType,omitempty" name:"IcmpType"`
	IcmpCode        *int    `json:"IcmpCode,omitempty" name:"IcmpCode"`
	PortRangeFrom   *int    `json:"PortRangeFrom,omitempty" name:"PortRangeFrom"`
	PortRangeTo     *int    `json:"PortRangeTo,omitempty" name:"PortRangeTo"`
	RuleTag         *string `json:"RuleTag,omitempty" name:"RuleTag"`
	Priority        *int    `json:"Priority,omitempty" name:"Priority"`
	Policy          *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *AuthorizeSecurityGroupEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthorizeSecurityGroupEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AuthorizeSecurityGroupEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AuthorizeSecurityGroupEntryResponse struct {
	*ksyunhttp.BaseResponse
	SecurityGroupEntryIdSet []*string `json:"SecurityGroupEntryIdSet" name:"SecurityGroupEntryIdSet"`
	RequestId               *string   `json:"RequestId" name:"RequestId"`
	Return                  *bool     `json:"Return" name:"Return"`
}

func (r *AuthorizeSecurityGroupEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthorizeSecurityGroupEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeSecurityGroupEntryRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId      *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupEntryId *string `json:"SecurityGroupEntryId,omitempty" name:"SecurityGroupEntryId"`
}

func (r *RevokeSecurityGroupEntryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeSecurityGroupEntryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RevokeSecurityGroupEntryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevokeSecurityGroupEntryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *RevokeSecurityGroupEntryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RevokeSecurityGroupEntryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId []*string                       `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	Filter          []*DescribeSecurityGroupsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults      *int                            `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken       *string                         `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	NextToken        *string `json:"NextToken" name:"NextToken"`
	SecurityGroupSet []struct {
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		VpcId                 *string `json:"VpcId" name:"VpcId"`
		SecurityGroupName     *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		SecurityGroupId       *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		Description           *string `json:"Description" name:"Description"`
		SecurityGroupType     *string `json:"SecurityGroupType" name:"SecurityGroupType"`
		SecurityGroupEntrySet []struct {
			CreateTime           *string `json:"CreateTime" name:"CreateTime"`
			Description          *string `json:"Description" name:"Description"`
			SecurityGroupEntryId *string `json:"SecurityGroupEntryId" name:"SecurityGroupEntryId"`
			CidrBlock            *string `json:"CidrBlock" name:"CidrBlock"`
			Direction            *string `json:"Direction" name:"Direction"`
			Protocol             *string `json:"Protocol" name:"Protocol"`
			IcmpType             *int    `json:"IcmpType" name:"IcmpType"`
			IcmpCode             *int    `json:"IcmpCode" name:"IcmpCode"`
			PortRangeFrom        *int    `json:"PortRangeFrom" name:"PortRangeFrom"`
			PortRangeTo          *int    `json:"PortRangeTo" name:"PortRangeTo"`
			RuleTag              *string `json:"RuleTag" name:"RuleTag"`
			Priority             *int    `json:"Priority" name:"Priority"`
			Policy               *string `json:"Policy" name:"Policy"`
		} `json:"SecurityGroupEntrySet" name:"SecurityGroupEntrySet"`
	} `json:"SecurityGroupSet"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRequest struct {
	*ksyunhttp.BaseRequest
	VpcId        *string `json:"VpcId,omitempty" name:"VpcId"`
	NatLineId    *string `json:"NatLineId,omitempty" name:"NatLineId"`
	BandWidth    *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	NatName      *string `json:"NatName,omitempty" name:"NatName"`
	NatType      *string `json:"NatType,omitempty" name:"NatType"`
	NatIpNumber  *int    `json:"NatIpNumber,omitempty" name:"NatIpNumber"`
	NatMode      *string `json:"NatMode,omitempty" name:"NatMode"`
	ProjectId    *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType   *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *CreateNatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string `json:"RequestId" name:"RequestId"`
	CreateTime  *string `json:"CreateTime" name:"CreateTime"`
	VpcId       *string `json:"VpcId" name:"VpcId"`
	NatId       *string `json:"NatId" name:"NatId"`
	NatName     *string `json:"NatName" name:"NatName"`
	NatMode     *string `json:"NatMode" name:"NatMode"`
	NatType     *string `json:"NatType" name:"NatType"`
	NatIpNumber *int    `json:"NatIpNumber" name:"NatIpNumber"`
	BandWidth   *int    `json:"BandWidth" name:"BandWidth"`
	ProjectId   *string `json:"ProjectId" name:"ProjectId"`
	ChargeType  *string `json:"ChargeType" name:"ChargeType"`
	NatIpSet    []struct {
		NatIp   *string `json:"NatIp" name:"NatIp"`
		NatIpId *string `json:"NatIpId" name:"NatIpId"`
		Enabled *bool   `json:"Enabled" name:"Enabled"`
	} `json:"NatIpSet"`
	ServiceEndTime                   *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	AssociateDirectConnectGatewaySet []struct {
		DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
	} `json:"AssociateDirectConnectGatewaySet"`
	AssociateVpnGatewaySet []struct {
		VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	} `json:"AssociateVpnGatewaySet"`
	AssociateInstanceSet []struct {
		PrivateIpAddress   *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"AssociateInstanceSet"`
	AssociateNatSet []struct {
		SubnetId *string `json:"SubnetId" name:"SubnetId"`
	} `json:"AssociateNatSet"`
	DnatSet []struct {
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		DnatId           *string `json:"DnatId" name:"DnatId"`
		NatId            *string `json:"NatId" name:"NatId"`
		DnatName         *string `json:"DnatName" name:"DnatName"`
		IpProtocol       *string `json:"IpProtocol" name:"IpProtocol"`
		NatIp            *string `json:"NatIp" name:"NatIp"`
		PublicPort       *string `json:"PublicPort" name:"PublicPort"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		PrivatePort      *string `json:"PrivatePort" name:"PrivatePort"`
		Description      *string `json:"Description" name:"Description"`
		Enabled          *bool   `json:"Enabled" name:"Enabled"`
	} `json:"DnatSet"`
}

func (r *CreateNatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatRequest struct {
	*ksyunhttp.BaseRequest
	NatId *string `json:"NatId,omitempty" name:"NatId"`
}

func (r *DeleteNatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteNatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId    []*string             `json:"ProjectId,omitempty" name:"ProjectId"`
	NatId        []*string             `json:"NatId,omitempty" name:"NatId"`
	Filter       []*DescribeNatsFilter `json:"Filter,omitempty" name:"Filter"`
	IsContainTag *bool                 `json:"IsContainTag,omitempty" name:"IsContainTag"`
	TagKey       []*string             `json:"TagKey,omitempty" name:"TagKey"`
	TagKV        []*DescribeNatsTagKV  `json:"TagKV,omitempty" name:"TagKV"`
	MaxResults   *int                  `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string               `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeNatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	NatSet    []struct {
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		NatId          *string `json:"NatId" name:"NatId"`
		NatName        *string `json:"NatName" name:"NatName"`
		NatMode        *string `json:"NatMode" name:"NatMode"`
		NatType        *string `json:"NatType" name:"NatType"`
		NatIpNumber    *int    `json:"NatIpNumber" name:"NatIpNumber"`
		BandWidth      *int    `json:"BandWidth" name:"BandWidth"`
		ProjectId      *string `json:"ProjectId" name:"ProjectId"`
		ChargeType     *string `json:"ChargeType" name:"ChargeType"`
		ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		NatIpSet       []struct {
			NatIp   *string `json:"NatIp" name:"NatIp"`
			NatIpId *string `json:"NatIpId" name:"NatIpId"`
			Enabled *bool   `json:"Enabled" name:"Enabled"`
		} `json:"NatIpSet" name:"NatIpSet"`
		AssociateDirectConnectGatewaySet []struct {
			DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
		} `json:"AssociateDirectConnectGatewaySet" name:"AssociateDirectConnectGatewaySet"`
		AssociateVpnGatewaySet []struct {
			VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
		} `json:"AssociateVpnGatewaySet" name:"AssociateVpnGatewaySet"`
		AssociateInstanceSet []struct {
			PrivateIpAddress   *string   `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			NetworkInterfaceId *string   `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
			NatIps             []*string `json:"NatIps" name:"NatIps"`
		} `json:"AssociateInstanceSet" name:"AssociateInstanceSet"`
		AssociateNatSet []struct {
			SubnetId *string   `json:"SubnetId" name:"SubnetId"`
			NatIps   []*string `json:"NatIps" name:"NatIps"`
		} `json:"AssociateNatSet" name:"AssociateNatSet"`
		DnatSet []struct {
			CreateTime       *string `json:"CreateTime" name:"CreateTime"`
			DnatId           *string `json:"DnatId" name:"DnatId"`
			NatId            *string `json:"NatId" name:"NatId"`
			DnatName         *string `json:"DnatName" name:"DnatName"`
			IpProtocol       *string `json:"IpProtocol" name:"IpProtocol"`
			NatIp            *string `json:"NatIp" name:"NatIp"`
			PublicPort       *string `json:"PublicPort" name:"PublicPort"`
			PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
			PrivatePort      *string `json:"PrivatePort" name:"PrivatePort"`
			Description      *string `json:"Description" name:"Description"`
			Enabled          *bool   `json:"Enabled" name:"Enabled"`
		} `json:"DnatSet" name:"DnatSet"`
		TagSet []struct {
			ResourceUuid *string `json:"ResourceUuid" name:"ResourceUuid"`
			TagId        *string `json:"TagId" name:"TagId"`
			TagKey       *string `json:"TagKey" name:"TagKey"`
			TagValue     *string `json:"TagValue" name:"TagValue"`
		} `json:"TagSet" name:"TagSet"`
	} `json:"NatSet"`
}

func (r *DescribeNatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNatRequest struct {
	*ksyunhttp.BaseRequest
	NatId    *string   `json:"NatId,omitempty" name:"NatId"`
	SubnetId *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	NatIpId  []*string `json:"NatIpId,omitempty" name:"NatIpId"`
}

func (r *AssociateNatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateNatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateNatResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateNatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateNatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatRequest struct {
	*ksyunhttp.BaseRequest
	NatId    *string `json:"NatId,omitempty" name:"NatId"`
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *DisassociateNatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateNatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateNatResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DisassociateNatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisassociateNatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetGatewaysRequest struct {
	*ksyunhttp.BaseRequest
	InternetGatewayId []*string                         `json:"InternetGatewayId,omitempty" name:"InternetGatewayId"`
	Filter            []*DescribeInternetGatewaysFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults        *int                              `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken         *string                           `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeInternetGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeInternetGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetGatewaysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NextToken          *string `json:"NextToken" name:"NextToken"`
	InternetGatewaySet []struct {
		CreateTime          *string `json:"CreateTime" name:"CreateTime"`
		VpcId               *string `json:"VpcId" name:"VpcId"`
		InternetGatewayName *string `json:"InternetGatewayName" name:"InternetGatewayName"`
		InternetGatewayId   *string `json:"InternetGatewayId" name:"InternetGatewayId"`
	} `json:"InternetGatewaySet"`
}

func (r *DescribeInternetGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternetGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeeringConnectionRequest struct {
	*ksyunhttp.BaseRequest
	VpcId         *string `json:"VpcId,omitempty" name:"VpcId"`
	PeeringName   *string `json:"PeeringName,omitempty" name:"PeeringName"`
	PeerVpcId     *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`
	PeerRegion    *string `json:"PeerRegion,omitempty" name:"PeerRegion"`
	PeerAccountId *string `json:"PeerAccountId,omitempty" name:"PeerAccountId"`
	BandWidth     *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	PurchaseTime  *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	ProjectId     *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType    *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateVpcPeeringConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcPeeringConnectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateVpcPeeringConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeeringConnectionRequest struct {
	*ksyunhttp.BaseRequest
	VpcPeeringConnectionId *string `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
}

func (r *DeleteVpcPeeringConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcPeeringConnectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteVpcPeeringConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeeringConnectionsRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId              []*string                              `json:"ProjectId,omitempty" name:"ProjectId"`
	VpcPeeringConnectionId []*string                              `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
	Filter                 []*DescribeVpcPeeringConnectionsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults             *int                                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken              *string                                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeVpcPeeringConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeeringConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpcPeeringConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcPeeringConnectionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	NextToken               *string `json:"NextToken" name:"NextToken"`
	VpcPeeringConnectionSet []struct {
		CreateTime               *string `json:"CreateTime" name:"CreateTime"`
		VpcId                    *string `json:"VpcId" name:"VpcId"`
		VpcPeeringConnectionId   *string `json:"VpcPeeringConnectionId" name:"VpcPeeringConnectionId"`
		VpcPeeringConnectionType *string `json:"VpcPeeringConnectionType" name:"VpcPeeringConnectionType"`
		PeeringName              *string `json:"PeeringName" name:"PeeringName"`
		State                    *string `json:"State" name:"State"`
		BandWidth                *int    `json:"BandWidth" name:"BandWidth"`
		ProjectId                *string `json:"ProjectId" name:"ProjectId"`
		ChargeType               *string `json:"ChargeType" name:"ChargeType"`
		ServiceEndTime           *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	} `json:"VpcPeeringConnectionSet"`
}

func (r *DescribeVpcPeeringConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcPeeringConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclRequest struct {
	*ksyunhttp.BaseRequest
	NetworkAclId   *string `json:"NetworkAclId,omitempty" name:"NetworkAclId"`
	NetworkAclName *string `json:"NetworkAclName,omitempty" name:"NetworkAclName"`
	Description    *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyNetworkAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkAclResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	CreateTime         *string `json:"CreateTime" name:"CreateTime"`
	VpcId              *string `json:"VpcId" name:"VpcId"`
	NetworkAclName     *string `json:"NetworkAclName" name:"NetworkAclName"`
	NetworkAclId       *string `json:"NetworkAclId" name:"NetworkAclId"`
	Description        *string `json:"Description" name:"Description"`
	NetworkAclEntrySet []struct {
		Description       *string `json:"Description" name:"Description"`
		NetworkAclId      *string `json:"NetworkAclId" name:"NetworkAclId"`
		NetworkAclEntryId *string `json:"NetworkAclEntryId" name:"NetworkAclEntryId"`
		CidrBlock         *string `json:"CidrBlock" name:"CidrBlock"`
		RuleNumber        *int    `json:"RuleNumber" name:"RuleNumber"`
		Direction         *string `json:"Direction" name:"Direction"`
		RuleAction        *string `json:"RuleAction" name:"RuleAction"`
		Protocol          *string `json:"Protocol" name:"Protocol"`
		IcmpType          *int    `json:"IcmpType" name:"IcmpType"`
		IcmpCode          *int    `json:"IcmpCode" name:"IcmpCode"`
		PortRangeFrom     *int    `json:"PortRangeFrom" name:"PortRangeFrom"`
		PortRangeTo       *int    `json:"PortRangeTo" name:"PortRangeTo"`
	} `json:"NetworkAclEntrySet"`
}

func (r *ModifyNetworkAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupRequest struct {
	*ksyunhttp.BaseRequest
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	CreateTime            *string `json:"CreateTime" name:"CreateTime"`
	VpcId                 *string `json:"VpcId" name:"VpcId"`
	SecurityGroupName     *string `json:"SecurityGroupName" name:"SecurityGroupName"`
	SecurityGroupId       *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	Description           *string `json:"Description" name:"Description"`
	SecurityGroupType     *string `json:"SecurityGroupType" name:"SecurityGroupType"`
	SecurityGroupEntrySet []struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		Description          *string `json:"Description" name:"Description"`
		SecurityGroupEntryId *string `json:"SecurityGroupEntryId" name:"SecurityGroupEntryId"`
		CidrBlock            *string `json:"CidrBlock" name:"CidrBlock"`
		Direction            *string `json:"Direction" name:"Direction"`
		Protocol             *string `json:"Protocol" name:"Protocol"`
		IcmpType             *int    `json:"IcmpType" name:"IcmpType"`
		IcmpCode             *int    `json:"IcmpCode" name:"IcmpCode"`
		PortRangeFrom        *int    `json:"PortRangeFrom" name:"PortRangeFrom"`
		PortRangeTo          *int    `json:"PortRangeTo" name:"PortRangeTo"`
		RuleTag              *string `json:"RuleTag" name:"RuleTag"`
		Priority             *int    `json:"Priority" name:"Priority"`
		Policy               *string `json:"Policy" name:"Policy"`
	} `json:"SecurityGroupEntrySet"`
}

func (r *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId   *string `json:"SubnetId,omitempty" name:"SubnetId"`
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	Dns1       *string `json:"Dns1,omitempty" name:"Dns1"`
	Dns2       *string `json:"Dns2,omitempty" name:"Dns2"`
}

func (r *ModifySubnetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifySubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubnetResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifySubnetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifySubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatRequest struct {
	*ksyunhttp.BaseRequest
	NatId     *string `json:"NatId,omitempty" name:"NatId"`
	NatName   *string `json:"NatName,omitempty" name:"NatName"`
	BandWidth *int    `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *ModifyNatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CreateTime     *string `json:"CreateTime" name:"CreateTime"`
	VpcId          *string `json:"VpcId" name:"VpcId"`
	NatId          *string `json:"NatId" name:"NatId"`
	NatName        *string `json:"NatName" name:"NatName"`
	NatMode        *string `json:"NatMode" name:"NatMode"`
	NatType        *string `json:"NatType" name:"NatType"`
	NatIpNumber    *int    `json:"NatIpNumber" name:"NatIpNumber"`
	BandWidth      *int    `json:"BandWidth" name:"BandWidth"`
	ProjectId      *string `json:"ProjectId" name:"ProjectId"`
	ChargeType     *string `json:"ChargeType" name:"ChargeType"`
	ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	NatIpSet       []struct {
		NatIp   *string `json:"NatIp" name:"NatIp"`
		NatIpId *string `json:"NatIpId" name:"NatIpId"`
		Enabled *bool   `json:"Enabled" name:"Enabled"`
	} `json:"NatIpSet"`
	AssociateDirectConnectGatewaySet []struct {
		DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
	} `json:"AssociateDirectConnectGatewaySet"`
	AssociateVpnGatewaySet []struct {
		VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	} `json:"AssociateVpnGatewaySet"`
	AssociateInstanceSet []struct {
		PrivateIpAddress   *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"AssociateInstanceSet"`
	AssociateNatSet []struct {
		SubnetId *string `json:"SubnetId" name:"SubnetId"`
	} `json:"AssociateNatSet"`
	DnatSet []struct {
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		DnatId           *string `json:"DnatId" name:"DnatId"`
		NatId            *string `json:"NatId" name:"NatId"`
		DnatName         *string `json:"DnatName" name:"DnatName"`
		IpProtocol       *string `json:"IpProtocol" name:"IpProtocol"`
		NatIp            *string `json:"NatIp" name:"NatIp"`
		PublicPort       *string `json:"PublicPort" name:"PublicPort"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		PrivatePort      *string `json:"PrivatePort" name:"PrivatePort"`
		Description      *string `json:"Description" name:"Description"`
		Enabled          *bool   `json:"Enabled" name:"Enabled"`
	} `json:"DnatSet"`
}

func (r *ModifyNatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId []*string                          `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	Filter             []*DescribeNetworkInterfacesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults         *int                               `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken          *string                            `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkInterfacesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId           *string `json:"RequestId" name:"RequestId"`
	NextToken           *string `json:"NextToken" name:"NextToken"`
	TotalCount          *int    `json:"TotalCount" name:"TotalCount"`
	NetworkInterfaceSet []struct {
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		VpcId                       *string `json:"VpcId" name:"VpcId"`
		VpcName                     *string `json:"VpcName" name:"VpcName"`
		SubnetName                  *string `json:"SubnetName" name:"SubnetName"`
		SubnetId                    *string `json:"SubnetId" name:"SubnetId"`
		CidrBlock                   *string `json:"CidrBlock" name:"CidrBlock"`
		SecurityGroupName           *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		SecurityGroupId             *string `json:"SecurityGroupId" name:"SecurityGroupId"`
		InstanceId                  *string `json:"InstanceId" name:"InstanceId"`
		PrivateIpAddress            *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		InstanceType                *string `json:"InstanceType" name:"InstanceType"`
		DNS1                        *string `json:"DNS1" name:"DNS1"`
		DNS2                        *string `json:"DNS2" name:"DNS2"`
		NetworkInterfaceType        *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
		MacAddress                  *string `json:"MacAddress" name:"MacAddress"`
		NetworkInterfaceId          *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		AvailabilityZoneName        *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
		AssignedPrivateIpAddressSet []struct {
			PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		} `json:"AssignedPrivateIpAddressSet" name:"AssignedPrivateIpAddressSet"`
		SecurityGroupSet []struct {
			SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
			SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
		} `json:"SecurityGroupSet" name:"SecurityGroupSet"`
		Ipv6PublicIpAddressSet []struct {
			Ipv6PublicIpAddress    *string `json:"Ipv6PublicIpAddress" name:"Ipv6PublicIpAddress"`
			Ipv6PublicIpAddressId  *string `json:"Ipv6PublicIpAddressId" name:"Ipv6PublicIpAddressId"`
			Ipv6PublicIpCreateTime *string `json:"Ipv6PublicIpCreateTime" name:"Ipv6PublicIpCreateTime"`
			BandWidth              *int    `json:"BandWidth" name:"BandWidth"`
		} `json:"Ipv6PublicIpAddressSet" name:"Ipv6PublicIpAddressSet"`
		NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
		Ipv6Public           *bool   `json:"Ipv6Public" name:"Ipv6Public"`
		Ipv6Address          *string `json:"Ipv6Address" name:"Ipv6Address"`
		State                *string `json:"State" name:"State"`
	} `json:"NetworkInterfaceSet"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAvailableAddressesRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*DescribeSubnetAvailableAddressesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                                      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                                   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeSubnetAvailableAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAvailableAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSubnetAvailableAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAvailableAddressesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string   `json:"RequestId" name:"RequestId"`
	NextToken          *string   `json:"NextToken" name:"NextToken"`
	AvailableIpAddress []*string `json:"AvailableIpAddress" name:"AvailableIpAddress"`
}

func (r *DescribeSubnetAvailableAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAvailableAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcRequest struct {
	*ksyunhttp.BaseRequest
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	VpcId   *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *ModifyVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcceptVpcPeeringConnectionRequest struct {
	*ksyunhttp.BaseRequest
	VpcPeeringConnectionId *string `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
}

func (r *AcceptVpcPeeringConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AcceptVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AcceptVpcPeeringConnectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AcceptVpcPeeringConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcceptVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectVpcPeeringConnectionRequest struct {
	*ksyunhttp.BaseRequest
	VpcPeeringConnectionId *string `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
}

func (r *RejectVpcPeeringConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "RejectVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RejectVpcPeeringConnectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *RejectVpcPeeringConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RejectVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcPeeringConnectionRequest struct {
	*ksyunhttp.BaseRequest
	VpcPeeringConnectionId *string `json:"VpcPeeringConnectionId,omitempty" name:"VpcPeeringConnectionId"`
	PeeringName            *string `json:"PeeringName,omitempty" name:"PeeringName"`
	BandWidth              *int    `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *ModifyVpcPeeringConnectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcPeeringConnectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyVpcPeeringConnectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailabilityZonesRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeAvailabilityZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailabilityZonesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	AvailabilityZoneInfo []struct {
		AvailabilityZoneName *string `json:"AvailabilityZoneName" name:"AvailabilityZoneName"`
	} `json:"AvailabilityZoneInfo"`
}

func (r *DescribeAvailabilityZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectId []*string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	MaxResults      *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken       *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeDirectConnectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	NextToken        *string `json:"NextToken" name:"NextToken"`
	DirectConnectSet []struct {
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		DirectConnectId   *string `json:"DirectConnectId" name:"DirectConnectId"`
		DirectConnectName *string `json:"DirectConnectName" name:"DirectConnectName"`
		Type              *string `json:"Type" name:"Type"`
		PopLocation       *string `json:"PopLocation" name:"PopLocation"`
		CustomerLocation  *string `json:"CustomerLocation" name:"CustomerLocation"`
		State             *string `json:"State" name:"State"`
		BandWidth         *int    `json:"BandWidth" name:"BandWidth"`
		Vlan              *bool   `json:"Vlan" name:"Vlan"`
		Distance          *int    `json:"Distance" name:"Distance"`
		VpcNOCId          *string `json:"VpcNOCId" name:"VpcNOCId"`
	} `json:"DirectConnectSet"`
}

func (r *DescribeDirectConnectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectId                 *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	VlanId                          *int    `json:"VlanId,omitempty" name:"VlanId"`
	RouteType                       *string `json:"RouteType,omitempty" name:"RouteType"`
	DirectConnectInterfaceName      *string `json:"DirectConnectInterfaceName,omitempty" name:"DirectConnectInterfaceName"`
	DirectConnectInterfaceAccountId *string `json:"DirectConnectInterfaceAccountId,omitempty" name:"DirectConnectInterfaceAccountId"`
	CustomerPeerIp                  *string `json:"CustomerPeerIp,omitempty" name:"CustomerPeerIp"`
	LocalPeerIp                     *string `json:"LocalPeerIp,omitempty" name:"LocalPeerIp"`
	HaDirectConnectId               *string `json:"HaDirectConnectId,omitempty" name:"HaDirectConnectId"`
	HaCustomerPeerIp                *string `json:"HaCustomerPeerIp,omitempty" name:"HaCustomerPeerIp"`
	HaLocalPeerIp                   *string `json:"HaLocalPeerIp,omitempty" name:"HaLocalPeerIp"`
	BgpPeer                         *string `json:"BgpPeer,omitempty" name:"BgpPeer"`
	ReliabilityMethod               *string `json:"ReliabilityMethod,omitempty" name:"ReliabilityMethod"`
	BfdConfigId                     *string `json:"BfdConfigId,omitempty" name:"BfdConfigId"`
	BgpClientToken                  *string `json:"BgpClientToken,omitempty" name:"BgpClientToken"`
	EnableIpv6                      *bool   `json:"EnableIpv6,omitempty" name:"EnableIpv6"`
	CustomerIpv6PeerIp              *string `json:"CustomerIpv6PeerIp,omitempty" name:"CustomerIpv6PeerIp"`
	LocalIpv6PeerIp                 *string `json:"LocalIpv6PeerIp,omitempty" name:"LocalIpv6PeerIp"`
}

func (r *CreateDirectConnectInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDirectConnectInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDirectConnectInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectInterfaceId *string `json:"DirectConnectInterfaceId,omitempty" name:"DirectConnectInterfaceId"`
}

func (r *DeleteDirectConnectInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDirectConnectInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteDirectConnectInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectInterfacesRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectInterfaceId []*string `json:"DirectConnectInterfaceId,omitempty" name:"DirectConnectInterfaceId"`
	MaxResults               *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken                *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeDirectConnectInterfacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectInterfacesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"RequestId" name:"RequestId"`
	NextToken                 *string `json:"NextToken" name:"NextToken"`
	DirectConnectInterfaceSet []struct {
		CreateTime                      *string `json:"CreateTime" name:"CreateTime"`
		DirectConnectInterfaceId        *string `json:"DirectConnectInterfaceId" name:"DirectConnectInterfaceId"`
		DirectConnectInterfaceName      *string `json:"DirectConnectInterfaceName" name:"DirectConnectInterfaceName"`
		DirectConnectId                 *string `json:"DirectConnectId" name:"DirectConnectId"`
		HaDirectConnectId               *string `json:"HaDirectConnectId" name:"HaDirectConnectId"`
		HaDirectConnectInterfaceName    *string `json:"HaDirectConnectInterfaceName" name:"HaDirectConnectInterfaceName"`
		HaDirectConnectInterfaceId      *string `json:"HaDirectConnectInterfaceId" name:"HaDirectConnectInterfaceId"`
		AccountId                       *string `json:"AccountId" name:"AccountId"`
		DirectConnectInterfaceAccountId *string `json:"DirectConnectInterfaceAccountId" name:"DirectConnectInterfaceAccountId"`
		CustomerPeerIp                  *string `json:"CustomerPeerIp" name:"CustomerPeerIp"`
		LocalPeerIp                     *string `json:"LocalPeerIp" name:"LocalPeerIp"`
		HaCustomerPeerIp                *string `json:"HaCustomerPeerIp" name:"HaCustomerPeerIp"`
		HaLocalPeerIp                   *string `json:"HaLocalPeerIp" name:"HaLocalPeerIp"`
		VlanId                          *int    `json:"VlanId" name:"VlanId"`
		HaVlanId                        *int    `json:"HaVlanId" name:"HaVlanId"`
		State                           *string `json:"State" name:"State"`
		RouteType                       *string `json:"RouteType" name:"RouteType"`
		BgpPeer                         *string `json:"BgpPeer" name:"BgpPeer"`
		ReliabilityMethod               *string `json:"ReliabilityMethod" name:"ReliabilityMethod"`
		BfdConfigId                     *string `json:"BfdConfigId" name:"BfdConfigId"`
		Priority                        *int    `json:"Priority" name:"Priority"`
		BgpClientToken                  *string `json:"BgpClientToken" name:"BgpClientToken"`
		EnableIpv6                      *bool   `json:"EnableIpv6" name:"EnableIpv6"`
		CustomerPeerIpv6                *string `json:"CustomerPeerIpv6" name:"CustomerPeerIpv6"`
		LocalPeerIpv6                   *string `json:"LocalPeerIpv6" name:"LocalPeerIpv6"`
	} `json:"DirectConnectInterfaceSet"`
}

func (r *DescribeDirectConnectInterfacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayRequest struct {
	*ksyunhttp.BaseRequest
	VpcId                    *string `json:"VpcId,omitempty" name:"VpcId"`
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
}

func (r *CreateDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
}

func (r *DeleteDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId []*string                              `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	Filter                 []*DescribeDirectConnectGatewaysFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults             *int                                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken              *string                                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeDirectConnectGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewaysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId               *string `json:"RequestId" name:"RequestId"`
	NextToken               *string `json:"NextToken" name:"NextToken"`
	DirectConnectGatewaySet []struct {
		CreateTime                    *string   `json:"CreateTime" name:"CreateTime"`
		VpcId                         *string   `json:"VpcId" name:"VpcId"`
		DirectConnectGatewayId        *string   `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
		NatId                         *string   `json:"NatId" name:"NatId"`
		DirectConnectInterfaceId      *string   `json:"DirectConnectInterfaceId" name:"DirectConnectInterfaceId"`
		DirectConnectGatewayName      *string   `json:"DirectConnectGatewayName" name:"DirectConnectGatewayName"`
		BandWidth                     *int      `json:"BandWidth" name:"BandWidth"`
		AssociatedInstanceType        *string   `json:"AssociatedInstanceType" name:"AssociatedInstanceType"`
		CenAccountId                  *string   `json:"CenAccountId" name:"CenAccountId"`
		Status                        *string   `json:"Status" name:"Status"`
		CenId                         *string   `json:"CenId" name:"CenId"`
		RemoteCidrSet                 []*string `json:"RemoteCidrSet" name:"RemoteCidrSet"`
		ExtraCidrSet                  []*string `json:"ExtraCidrSet" name:"ExtraCidrSet"`
		Version                       *string   `json:"Version" name:"Version"`
		DirectConnectInterfaceInfoSet []struct {
			DirectConnectInterfaceId *string `json:"DirectConnectInterfaceId" name:"DirectConnectInterfaceId"`
		} `json:"DirectConnectInterfaceInfoSet" name:"DirectConnectInterfaceInfoSet"`
	} `json:"DirectConnectGatewaySet"`
}

func (r *DescribeDirectConnectGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDirectConnectGatewayRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId   *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	DirectConnectInterfaceId *string `json:"DirectConnectInterfaceId,omitempty" name:"DirectConnectInterfaceId"`
}

func (r *AttachDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachDirectConnectGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AttachDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDirectConnectGatewayRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId   *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	DirectConnectInterfaceId *string `json:"DirectConnectInterfaceId,omitempty" name:"DirectConnectInterfaceId"`
}

func (r *DetachDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachDirectConnectGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectInterfaceId   *string `json:"DirectConnectInterfaceId,omitempty" name:"DirectConnectInterfaceId"`
	DirectConnectInterfaceName *string `json:"DirectConnectInterfaceName,omitempty" name:"DirectConnectInterfaceName"`
}

func (r *ModifyDirectConnectInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDirectConnectInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDirectConnectInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId   *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitempty" name:"DirectConnectGatewayName"`
}

func (r *ModifyDirectConnectGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDirectConnectGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRequest struct {
	*ksyunhttp.BaseRequest
	VpcId             *string `json:"VpcId,omitempty" name:"VpcId"`
	BandWidth         *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	VpnGatewayName    *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
	ProjectId         *string `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType        *string `json:"ChargeType,omitempty" name:"ChargeType"`
	PurchaseTime      *int    `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	VpnGatewayVersion *string `json:"VpnGatewayVersion,omitempty" name:"VpnGatewayVersion"`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayId   *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	BandWidth      *int    `json:"BandWidth,omitempty" name:"BandWidth"`
	VpnGatewayName *string `json:"VpnGatewayName,omitempty" name:"VpnGatewayName"`
}

func (r *ModifyVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysRequest struct {
	*ksyunhttp.BaseRequest
	ProjectId    []*string                    `json:"ProjectId,omitempty" name:"ProjectId"`
	VpnGatewayId []*string                    `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	Filter       []*DescribeVpnGatewaysFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults   *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string                      `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewaysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	VpnGatewaySet []struct {
		CreateTime        *string `json:"CreateTime" name:"CreateTime"`
		VpcId             *string `json:"VpcId" name:"VpcId"`
		VpnGatewayId      *string `json:"VpnGatewayId" name:"VpnGatewayId"`
		VpnGatewayName    *string `json:"VpnGatewayName" name:"VpnGatewayName"`
		BandWidth         *int    `json:"BandWidth" name:"BandWidth"`
		GatewayAddress    *string `json:"GatewayAddress" name:"GatewayAddress"`
		HaGatewayAddress  *string `json:"HaGatewayAddress" name:"HaGatewayAddress"`
		VpnSwitchType     *string `json:"VpnSwitchType" name:"VpnSwitchType"`
		ProjectId         *string `json:"ProjectId" name:"ProjectId"`
		ChargeType        *string `json:"ChargeType" name:"ChargeType"`
		ServiceEndTime    *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		VpnGatewayVersion *string `json:"VpnGatewayVersion" name:"VpnGatewayVersion"`
	} `json:"VpnGatewaySet"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnTunnelRequest struct {
	*ksyunhttp.BaseRequest
	CustomerGatewayId    *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	VpnGatewayId         *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	VpnTunnelName        *string `json:"VpnTunnelName,omitempty" name:"VpnTunnelName"`
	IpsecAuthenAlgorithm *string `json:"IpsecAuthenAlgorithm,omitempty" name:"IpsecAuthenAlgorithm"`
	IpsecEncryAlgorithm  *string `json:"IpsecEncryAlgorithm,omitempty" name:"IpsecEncryAlgorithm"`
	IkeAuthenAlgorithm   *string `json:"IkeAuthenAlgorithm,omitempty" name:"IkeAuthenAlgorithm"`
	IkeEncryAlgorithm    *string `json:"IkeEncryAlgorithm,omitempty" name:"IkeEncryAlgorithm"`
	Type                 *string `json:"Type,omitempty" name:"Type"`
	OpenHealthCheck      *bool   `json:"OpenHealthCheck,omitempty" name:"OpenHealthCheck"`
	PreSharedKey         *string `json:"PreSharedKey,omitempty" name:"PreSharedKey"`
	IpsecLifetimeSecond  *int    `json:"IpsecLifetimeSecond,omitempty" name:"IpsecLifetimeSecond"`
	IpsecLifetimeTraffic *int    `json:"IpsecLifetimeTraffic,omitempty" name:"IpsecLifetimeTraffic"`
	IkeDHGroup           *string `json:"IkeDHGroup,omitempty" name:"IkeDHGroup"`
	EnableNatTraversal   *bool   `json:"EnableNatTraversal,omitempty" name:"EnableNatTraversal"`
	VpnGreIp             *string `json:"VpnGreIp,omitempty" name:"VpnGreIp"`
	HaVpnGreIp           *string `json:"HaVpnGreIp,omitempty" name:"HaVpnGreIp"`
	CustomerGreIp        *string `json:"CustomerGreIp,omitempty" name:"CustomerGreIp"`
	HaCustomerGreIp      *string `json:"HaCustomerGreIp,omitempty" name:"HaCustomerGreIp"`
	HaMode               *string `json:"HaMode,omitempty" name:"HaMode"`
	LocalPeerIp          *string `json:"LocalPeerIp,omitempty" name:"LocalPeerIp"`
	CustomerPeerIp       *string `json:"CustomerPeerIp,omitempty" name:"CustomerPeerIp"`
	IkeVersion           *string `json:"IkeVersion,omitempty" name:"IkeVersion"`
}

func (r *CreateVpnTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVpnTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnTunnelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateVpnTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnTunnelRequest struct {
	*ksyunhttp.BaseRequest
	VpnTunnelId   *string `json:"VpnTunnelId,omitempty" name:"VpnTunnelId"`
	VpnTunnelName *string `json:"VpnTunnelName,omitempty" name:"VpnTunnelName"`
}

func (r *ModifyVpnTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVpnTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnTunnelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyVpnTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnTunnelRequest struct {
	*ksyunhttp.BaseRequest
	VpnTunnelId *string `json:"VpnTunnelId,omitempty" name:"VpnTunnelId"`
}

func (r *DeleteVpnTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnTunnelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVpnTunnelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnTunnelResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteVpnTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnTunnelsRequest struct {
	*ksyunhttp.BaseRequest
	VpnTunnelId []*string                   `json:"VpnTunnelId,omitempty" name:"VpnTunnelId"`
	Filter      []*DescribeVpnTunnelsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults  *int                        `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken   *string                     `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeVpnTunnelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnTunnelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpnTunnelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnTunnelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NextToken    *string `json:"NextToken" name:"NextToken"`
	VpnTunnelSet []struct {
		CreateTime           *string   `json:"CreateTime" name:"CreateTime"`
		OpenHealthCheck      *bool     `json:"OpenHealthCheck" name:"OpenHealthCheck"`
		State                *string   `json:"State" name:"State"`
		VpnTunnelId          *string   `json:"VpnTunnelId" name:"VpnTunnelId"`
		VpnGreIp             *string   `json:"VpnGreIp" name:"VpnGreIp"`
		Type                 *string   `json:"Type" name:"Type"`
		HaMode               *string   `json:"HaMode" name:"HaMode"`
		IkeVersion           *string   `json:"IkeVersion" name:"IkeVersion"`
		LocalPeerIp          *string   `json:"LocalPeerIp" name:"LocalPeerIp"`
		CustomerPeerIp       *string   `json:"CustomerPeerIp" name:"CustomerPeerIp"`
		CustomerGreIp        *string   `json:"CustomerGreIp" name:"CustomerGreIp"`
		HaVpnGreIp           *string   `json:"HaVpnGreIp" name:"HaVpnGreIp"`
		HaCustomerGreIp      *string   `json:"HaCustomerGreIp" name:"HaCustomerGreIp"`
		VpnGatewayId         *string   `json:"VpnGatewayId" name:"VpnGatewayId"`
		VpnGatewayVersion    *string   `json:"VpnGatewayVersion" name:"VpnGatewayVersion"`
		CustomerGatewayId    *string   `json:"CustomerGatewayId" name:"CustomerGatewayId"`
		PreSharedKey         *string   `json:"PreSharedKey" name:"PreSharedKey"`
		IkeAuthenAlgorithm   *string   `json:"IkeAuthenAlgorithm" name:"IkeAuthenAlgorithm"`
		IkeDHGroup           *string   `json:"IkeDHGroup" name:"IkeDHGroup"`
		IkeEncryAlgorithm    *string   `json:"IkeEncryAlgorithm" name:"IkeEncryAlgorithm"`
		IpsecAuthenAlgorithm *string   `json:"IpsecAuthenAlgorithm" name:"IpsecAuthenAlgorithm"`
		IpsecEncryAlgorithm  *string   `json:"IpsecEncryAlgorithm" name:"IpsecEncryAlgorithm"`
		IpsecLifetimeSecond  *int      `json:"IpsecLifetimeSecond" name:"IpsecLifetimeSecond"`
		IpsecLifetimeTraffic *int      `json:"IpsecLifetimeTraffic" name:"IpsecLifetimeTraffic"`
		VpnTunnelName        *string   `json:"VpnTunnelName" name:"VpnTunnelName"`
		VpnGwlName           *string   `json:"VpnGwlName" name:"VpnGwlName"`
		ExtraCidrSet         []*string `json:"ExtraCidrSet" name:"ExtraCidrSet"`
		NatId                *string   `json:"NatId" name:"NatId"`
		EnableNatTraversal   *bool     `json:"EnableNatTraversal" name:"EnableNatTraversal"`
	} `json:"VpnTunnelSet"`
}

func (r *DescribeVpnTunnelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnTunnelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayRequest struct {
	*ksyunhttp.BaseRequest
	CustomerGatewayName      *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
	CustomerGatewayAddress   *string `json:"CustomerGatewayAddress,omitempty" name:"CustomerGatewayAddress"`
	HaCustomerGatewayAddress *string `json:"HaCustomerGatewayAddress,omitempty" name:"HaCustomerGatewayAddress"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustomerGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayRequest struct {
	*ksyunhttp.BaseRequest
	CustomerGatewayId   *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	CustomerGatewayName *string `json:"CustomerGatewayName,omitempty" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomerGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyCustomerGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayRequest struct {
	*ksyunhttp.BaseRequest
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCustomerGatewayResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectId   *string `json:"DirectConnectId,omitempty" name:"DirectConnectId"`
	DirectConnectName *string `json:"DirectConnectName,omitempty" name:"DirectConnectName"`
}

func (r *ModifyDirectConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyDirectConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDirectConnectResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyDirectConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDirectConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysRequest struct {
	*ksyunhttp.BaseRequest
	CustomerGatewayId []*string `json:"CustomerGatewayId,omitempty" name:"CustomerGatewayId"`
	MaxResults        *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken         *string   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerGatewaysResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NextToken          *string `json:"NextToken" name:"NextToken"`
	CustomerGatewaySet []struct {
		CreateTime               *string `json:"CreateTime" name:"CreateTime"`
		CustomerGatewayId        *string `json:"CustomerGatewayId" name:"CustomerGatewayId"`
		CustomerGatewayName      *string `json:"CustomerGatewayName" name:"CustomerGatewayName"`
		CustomerGatewayAddress   *string `json:"CustomerGatewayAddress" name:"CustomerGatewayAddress"`
		HaCustomerGatewayAddress *string `json:"HaCustomerGatewayAddress" name:"HaCustomerGatewayAddress"`
	} `json:"CustomerGatewaySet"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAllocatedIpAddressesRequest struct {
	*ksyunhttp.BaseRequest
	Filter     []*DescribeSubnetAllocatedIpAddressesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                                        `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                                     `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeSubnetAllocatedIpAddressesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAllocatedIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeSubnetAllocatedIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubnetAllocatedIpAddressesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string   `json:"RequestId" name:"RequestId"`
	NextToken          *string   `json:"NextToken" name:"NextToken"`
	AvailableIpAddress []*string `json:"AvailableIpAddress" name:"AvailableIpAddress"`
}

func (r *DescribeSubnetAllocatedIpAddressesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubnetAllocatedIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNatIpRequest struct {
	*ksyunhttp.BaseRequest
	NatId     *string `json:"NatId,omitempty" name:"NatId"`
	AddNumber *int    `json:"AddNumber,omitempty" name:"AddNumber"`
}

func (r *AddNatIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddNatIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddNatIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CreateTime     *string `json:"CreateTime" name:"CreateTime"`
	VpcId          *string `json:"VpcId" name:"VpcId"`
	NatId          *string `json:"NatId" name:"NatId"`
	NatName        *string `json:"NatName" name:"NatName"`
	NatMode        *string `json:"NatMode" name:"NatMode"`
	NatType        *string `json:"NatType" name:"NatType"`
	NatIpNumber    *int    `json:"NatIpNumber" name:"NatIpNumber"`
	BandWidth      *int    `json:"BandWidth" name:"BandWidth"`
	ProjectId      *string `json:"ProjectId" name:"ProjectId"`
	ChargeType     *string `json:"ChargeType" name:"ChargeType"`
	ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	NatIpSet       []struct {
		NatIp   *string `json:"NatIp" name:"NatIp"`
		NatIpId *string `json:"NatIpId" name:"NatIpId"`
		Enabled *bool   `json:"Enabled" name:"Enabled"`
	} `json:"NatIpSet"`
	AssociateDirectConnectGatewaySet []struct {
		DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
	} `json:"AssociateDirectConnectGatewaySet"`
	AssociateVpnGatewaySet []struct {
		VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	} `json:"AssociateVpnGatewaySet"`
	AssociateInstanceSet []struct {
		PrivateIpAddress   *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"AssociateInstanceSet"`
	AssociateNatSet []struct {
		SubnetId *string `json:"SubnetId" name:"SubnetId"`
	} `json:"AssociateNatSet"`
	DnatSet []struct {
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		DnatId           *string `json:"DnatId" name:"DnatId"`
		NatId            *string `json:"NatId" name:"NatId"`
		DnatName         *string `json:"DnatName" name:"DnatName"`
		IpProtocol       *string `json:"IpProtocol" name:"IpProtocol"`
		NatIp            *string `json:"NatIp" name:"NatIp"`
		PublicPort       *string `json:"PublicPort" name:"PublicPort"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		PrivatePort      *string `json:"PrivatePort" name:"PrivatePort"`
		Description      *string `json:"Description" name:"Description"`
		Enabled          *bool   `json:"Enabled" name:"Enabled"`
	} `json:"DnatSet"`
}

func (r *AddNatIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNatIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatIpRequest struct {
	*ksyunhttp.BaseRequest
	NatId   *string `json:"NatId,omitempty" name:"NatId"`
	NatIpId *string `json:"NatIpId,omitempty" name:"NatIpId"`
}

func (r *DeleteNatIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNatIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	CreateTime     *string `json:"CreateTime" name:"CreateTime"`
	VpcId          *string `json:"VpcId" name:"VpcId"`
	NatId          *string `json:"NatId" name:"NatId"`
	NatName        *string `json:"NatName" name:"NatName"`
	NatMode        *string `json:"NatMode" name:"NatMode"`
	NatType        *string `json:"NatType" name:"NatType"`
	NatIpNumber    *int    `json:"NatIpNumber" name:"NatIpNumber"`
	BandWidth      *int    `json:"BandWidth" name:"BandWidth"`
	ProjectId      *string `json:"ProjectId" name:"ProjectId"`
	ChargeType     *string `json:"ChargeType" name:"ChargeType"`
	ServiceEndTime *string `json:"ServiceEndTime" name:"ServiceEndTime"`
	NatIpSet       []struct {
		NatIp   *string `json:"NatIp" name:"NatIp"`
		NatIpId *string `json:"NatIpId" name:"NatIpId"`
		Enabled *bool   `json:"Enabled" name:"Enabled"`
	} `json:"NatIpSet"`
	AssociateDirectConnectGatewaySet []struct {
		DirectConnectGatewayId *string `json:"DirectConnectGatewayId" name:"DirectConnectGatewayId"`
	} `json:"AssociateDirectConnectGatewaySet"`
	AssociateVpnGatewaySet []struct {
		VpnGatewayId *string `json:"VpnGatewayId" name:"VpnGatewayId"`
	} `json:"AssociateVpnGatewaySet"`
	AssociateInstanceSet []struct {
		PrivateIpAddress   *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
	} `json:"AssociateInstanceSet"`
	AssociateNatSet []struct {
		SubnetId *string `json:"SubnetId" name:"SubnetId"`
	} `json:"AssociateNatSet"`
	DnatSet []struct {
		CreateTime       *string `json:"CreateTime" name:"CreateTime"`
		DnatId           *string `json:"DnatId" name:"DnatId"`
		NatId            *string `json:"NatId" name:"NatId"`
		DnatName         *string `json:"DnatName" name:"DnatName"`
		IpProtocol       *string `json:"IpProtocol" name:"IpProtocol"`
		NatIp            *string `json:"NatIp" name:"NatIp"`
		PublicPort       *string `json:"PublicPort" name:"PublicPort"`
		PrivateIpAddress *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		PrivatePort      *string `json:"PrivatePort" name:"PrivatePort"`
		Description      *string `json:"Description" name:"Description"`
		Enabled          *bool   `json:"Enabled" name:"Enabled"`
	} `json:"DnatSet"`
}

func (r *DeleteNatIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressAttributeRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId         *string `json:"SubnetId,omitempty" name:"SubnetId"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	Status           *string `json:"Status,omitempty" name:"Status"`
	AllocateStatus   *string `json:"AllocateStatus,omitempty" name:"AllocateStatus"`
}

func (r *ModifyPrivateIpAddressAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateIpAddressAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyPrivateIpAddressAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateIpAddressAttributeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyPrivateIpAddressAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPrivateIpAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectRoutesRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectRouteId []*string                            `json:"DirectConnectRouteId,omitempty" name:"DirectConnectRouteId"`
	MaxResults           *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
	Filter               []*DescribeDirectConnectRoutesFilter `json:"Filter,omitempty" name:"Filter"`
	NextToken            *string                              `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeDirectConnectRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectRoutesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId             *string `json:"RequestId" name:"RequestId"`
	NextToken             *string `json:"NextToken" name:"NextToken"`
	DirectConnectRouteSet []struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		DirectConnectId      *string `json:"DirectConnectId" name:"DirectConnectId"`
		CenStatus            *string `json:"CenStatus" name:"CenStatus"`
		BgpStatus            *string `json:"BgpStatus" name:"BgpStatus"`
		DirectConnectRouteId *string `json:"DirectConnectRouteId" name:"DirectConnectRouteId"`
		DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
		RouteType            *string `json:"RouteType" name:"RouteType"`
		NextHopSet           []struct {
			GatewayId   *string `json:"GatewayId" name:"GatewayId"`
			GatewayName *string `json:"GatewayName" name:"GatewayName"`
		} `json:"NextHopSet" name:"NextHopSet"`
	} `json:"DirectConnectRouteSet"`
}

func (r *DescribeDirectConnectRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetachDirectConnectGatewayWithVpcRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DetachDirectConnectGatewayWithVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDirectConnectGatewayWithVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DetachDirectConnectGatewayWithVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetachDirectConnectGatewayWithVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DetachDirectConnectGatewayWithVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DetachDirectConnectGatewayWithVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachDirectConnectGatewayWithVpcRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	VpcId                  *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *AttachDirectConnectGatewayWithVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDirectConnectGatewayWithVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AttachDirectConnectGatewayWithVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AttachDirectConnectGatewayWithVpcResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AttachDirectConnectGatewayWithVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AttachDirectConnectGatewayWithVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableRequest struct {
	*ksyunhttp.BaseRequest
	VpcId          *string `json:"VpcId,omitempty" name:"VpcId"`
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	Description    *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteTableResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableRequest struct {
	*ksyunhttp.BaseRequest
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteTableResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableRequest struct {
	*ksyunhttp.BaseRequest
	RouteTableId   *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	Description    *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRouteTableResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
}

func (r *ModifyRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesRequest struct {
	*ksyunhttp.BaseRequest
	RouteTableId []*string                    `json:"RouteTableId,omitempty" name:"RouteTableId"`
	Filter       []*DescribeRouteTablesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults   *int                         `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string                      `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTablesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	NextToken     *string `json:"NextToken" name:"NextToken"`
	RouteTableSet []struct {
		CreateTime     *string `json:"CreateTime" name:"CreateTime"`
		VpcId          *string `json:"VpcId" name:"VpcId"`
		RouteTableId   *string `json:"RouteTableId" name:"RouteTableId"`
		RouteTableName *string `json:"RouteTableName" name:"RouteTableName"`
		Description    *string `json:"Description" name:"Description"`
		RouteTableType *string `json:"RouteTableType" name:"RouteTableType"`
	} `json:"RouteTableSet"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateRouteTableRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId     *string `json:"SubnetId,omitempty" name:"SubnetId"`
	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

func (r *AssociateRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateRouteTableResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNetworkInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId             *string   `json:"SubnetId,omitempty" name:"SubnetId"`
	NetworkInterfaceName *string   `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`
	PrivateIpAddress     *string   `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	SecurityGroupId      []*string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNetworkInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" name:"NetworkInterfaceName"`
	NetworkInterfaceId   *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *ModifyNetworkInterfaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkInterfaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *ModifyNetworkInterfaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	BandwidthLimit     *int    `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	InBandwidthLimit   *int    `json:"inBandwidthLimit,omitempty" name:"inBandwidthLimit"`
}

func (r *CreateNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	NatId  *string                       `json:"NatId,omitempty" name:"NatId"`
	Filter []*DescribeNatRateLimitFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	NextToken              *string `json:"NextToken" name:"NextToken"`
	NatNetworkInterfaceSet []struct {
		NetworkInterfaceId   *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BandwidthLimit       *int    `json:"BandwidthLimit" name:"BandwidthLimit"`
		InBandwidthLimit     *int    `json:"InBandwidthLimit" name:"InBandwidthLimit"`
		NatRateLimitId       *string `json:"NatRateLimitId" name:"NatRateLimitId"`
		NetworkInterfaceName *string `json:"NetworkInterfaceName" name:"NetworkInterfaceName"`
		NetworkInterfaceType *string `json:"NetworkInterfaceType" name:"NetworkInterfaceType"`
		InstanceId           *string `json:"InstanceId" name:"InstanceId"`
		PrivateIpAddress     *string `json:"PrivateIpAddress" name:"PrivateIpAddress"`
		InstanceType         *string `json:"InstanceType" name:"InstanceType"`
	} `json:"NatNetworkInterfaceSet"`
}

func (r *DescribeNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	NatRateLimitId   *string `json:"NatRateLimitId,omitempty" name:"NatRateLimitId"`
	BandwidthLimit   *int    `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	InBandwidthLimit *int    `json:"InBandwidthLimit,omitempty" name:"InBandwidthLimit"`
}

func (r *ModifyNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	NatRateLimitId *string `json:"NatRateLimitId,omitempty" name:"NatRateLimitId"`
}

func (r *DeleteNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHaVipRequest struct {
	*ksyunhttp.BaseRequest
	SubnetId  *string `json:"SubnetId,omitempty" name:"SubnetId"`
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateHaVipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipRequest struct {
	*ksyunhttp.BaseRequest
	HaVipId *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHaVipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateHaVipRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	HaVipId            *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *AssociateHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssociateHaVipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *AssociateHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssociateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnAssociateHaVipRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	HaVipId            *string `json:"HaVipId,omitempty" name:"HaVipId"`
}

func (r *UnAssociateHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnAssociateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnAssociateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnAssociateHaVipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UnAssociateHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnAssociateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipRequest struct {
	*ksyunhttp.BaseRequest
	HaVipId    []*string              `json:"HaVipId,omitempty" name:"HaVipId"`
	Filter     []*DescribeHaVipFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                   `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeHaVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHaVipResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	HaVipSet  []struct {
		HaVipId                    *string   `json:"HaVipId" name:"HaVipId"`
		SubnetId                   *string   `json:"SubnetId" name:"SubnetId"`
		MasterNetworkInterfaceId   *string   `json:"MasterNetworkInterfaceId" name:"MasterNetworkInterfaceId"`
		VpcId                      *string   `json:"VpcId" name:"VpcId"`
		AllocationId               *string   `json:"AllocationId" name:"AllocationId"`
		IpAddress                  *string   `json:"IpAddress" name:"IpAddress"`
		CreateTime                 *string   `json:"CreateTime" name:"CreateTime"`
		SlaveNetworkInterfaceIdSet []*string `json:"SlaveNetworkInterfaceIdSet" name:"SlaveNetworkInterfaceIdSet"`
	} `json:"HaVipSet"`
}

func (r *DescribeHaVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	DestinationCidrBlock   *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	NextHopType            *string `json:"NextHopType,omitempty" name:"NextHopType"`
	Priority               *int    `json:"Priority,omitempty" name:"Priority"`
	NextHopInstance        *string `json:"NextHopInstance,omitempty" name:"NextHopInstance"`
	EnableIpv6             *bool   `json:"EnableIpv6,omitempty" name:"EnableIpv6"`
}

func (r *CreateDirectConnectGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDirectConnectGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDirectConnectGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDirectConnectGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayRouteId *string `json:"DirectConnectGatewayRouteId,omitempty" name:"DirectConnectGatewayRouteId"`
}

func (r *DeleteDirectConnectGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDirectConnectGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteDirectConnectGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDirectConnectGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayId *string                                    `json:"DirectConnectGatewayId,omitempty" name:"DirectConnectGatewayId"`
	MaxResults             *int                                       `json:"MaxResults,omitempty" name:"MaxResults"`
	Filter                 []*DescribeDirectConnectGatewayRouteFilter `json:"Filter,omitempty" name:"Filter"`
	NextToken              *string                                    `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeDirectConnectGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                    *string `json:"RequestId" name:"RequestId"`
	NextToken                    *string `json:"NextToken" name:"NextToken"`
	DirectConnectGatewayRouteSet []struct {
		DirectConnectGatewayRouteId *string `json:"DirectConnectGatewayRouteId" name:"DirectConnectGatewayRouteId"`
		DestinationCidrBlock        *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
		NextHopInstance             *string `json:"NextHopInstance" name:"NextHopInstance"`
		NextHopInstanceName         *string `json:"NextHopInstanceName" name:"NextHopInstanceName"`
		NextHopType                 *string `json:"NextHopType" name:"NextHopType"`
		Priority                    *int    `json:"Priority" name:"Priority"`
		AsPath                      *int    `json:"AsPath" name:"AsPath"`
		CreateTime                  *string `json:"CreateTime" name:"CreateTime"`
		DirectConnectId             *string `json:"DirectConnectId" name:"DirectConnectId"`
		BgpStatus                   *string `json:"BgpStatus" name:"BgpStatus"`
		RouteType                   *string `json:"RouteType" name:"RouteType"`
	} `json:"DirectConnectGatewayRouteSet"`
}

func (r *DescribeDirectConnectGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishDirectConnectRouteRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayRouteId *string `json:"DirectConnectGatewayRouteId,omitempty" name:"DirectConnectGatewayRouteId"`
}

func (r *PublishDirectConnectRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishDirectConnectRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "PublishDirectConnectRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PublishDirectConnectRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *PublishDirectConnectRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishDirectConnectRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnpublishDirectConnectRouteRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectGatewayRouteId *string `json:"DirectConnectGatewayRouteId,omitempty" name:"DirectConnectGatewayRouteId"`
}

func (r *UnpublishDirectConnectRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnpublishDirectConnectRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnpublishDirectConnectRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnpublishDirectConnectRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UnpublishDirectConnectRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnpublishDirectConnectRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSecondaryCidrBlockRequest struct {
	*ksyunhttp.BaseRequest
	VpcId     *string `json:"VpcId,omitempty" name:"VpcId"`
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
}

func (r *AddSecondaryCidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSecondaryCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddSecondaryCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddSecondaryCidrBlockResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AddSecondaryCidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSecondaryCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecondaryCidrBlockRequest struct {
	*ksyunhttp.BaseRequest
	VpcId           *string `json:"VpcId,omitempty" name:"VpcId"`
	SecondaryCidrId *string `json:"SecondaryCidrId,omitempty" name:"SecondaryCidrId"`
}

func (r *DeleteSecondaryCidrBlockRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecondaryCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteSecondaryCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecondaryCidrBlockResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteSecondaryCidrBlockResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSecondaryCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId             *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	PrivateIpAddress               []*string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
	SecondaryPrivateIpAddressCount *int      `json:"SecondaryPrivateIpAddressCount,omitempty" name:"SecondaryPrivateIpAddressCount"`
}

func (r *AssignPrivateIpAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignPrivateIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssignPrivateIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssignPrivateIpAddressResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *AssignPrivateIpAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AssignPrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressRequest struct {
	*ksyunhttp.BaseRequest
	NetworkInterfaceId *string   `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
	PrivateIpAddress   []*string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *UnassignPrivateIpAddressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignPrivateIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UnassignPrivateIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnassignPrivateIpAddressResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *UnassignPrivateIpAddressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnassignPrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	BandwidthLimit     *int      `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	InBandwidthLimit   *int      `json:"InBandwidthLimit,omitempty" name:"InBandwidthLimit"`
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
}

func (r *BatchCreateNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BatchCreateNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NatRateLimit struct {
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BandwidthLimit     *int    `json:"BandwidthLimit" name:"BandwidthLimit"`
		NatRateLimitId     *string `json:"NatRateLimitId" name:"NatRateLimitId"`
		InBandwidthLimit   *int    `json:"InBandwidthLimit" name:"InBandwidthLimit"`
	} `json:"NatRateLimit"`
}

func (r *BatchCreateNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchCreateNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	BandwidthLimit   *string   `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`
	InBandwidthLimit *int      `json:"InBandwidthLimit,omitempty" name:"InBandwidthLimit"`
	NatRateLimitId   []*string `json:"NatRateLimitId,omitempty" name:"NatRateLimitId"`
}

func (r *BatchModifyNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchModifyNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BatchModifyNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchModifyNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	NatRateLimit struct {
		NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
		BandwidthLimit     *int    `json:"BandwidthLimit" name:"BandwidthLimit"`
		NatRateLimitId     *string `json:"NatRateLimitId" name:"NatRateLimitId"`
		InBandwidthLimit   *int    `json:"InBandwidthLimit" name:"InBandwidthLimit"`
	} `json:"NatRateLimit"`
}

func (r *BatchModifyNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchModifyNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteNatRateLimitRequest struct {
	*ksyunhttp.BaseRequest
	NatRateLimitId []*string `json:"NatRateLimitId,omitempty" name:"NatRateLimitId"`
}

func (r *BatchDeleteNatRateLimitRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteNatRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "BatchDeleteNatRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BatchDeleteNatRateLimitResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *BatchDeleteNatRateLimitResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDeleteNatRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayRoutesRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayId *string                           `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	Filter       []*DescribeVpnGatewayRoutesFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults   *int                              `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken    *string                           `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeVpnGatewayRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnGatewayRoutesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NextToken          *string `json:"NextToken" name:"NextToken"`
	VpnGatewayRouteSet []struct {
		CreateTime           *string `json:"CreateTime" name:"CreateTime"`
		VpnGatewayRouteId    *string `json:"VpnGatewayRouteId" name:"VpnGatewayRouteId"`
		DestinationCidrBlock *string `json:"DestinationCidrBlock" name:"DestinationCidrBlock"`
		RouteType            *string `json:"RouteType" name:"RouteType"`
		NextHopType          *string `json:"NextHopType" name:"NextHopType"`
		NextHopInstanceName  *string `json:"NextHopInstanceName" name:"NextHopInstanceName"`
		VpnGatewayId         *string `json:"VpnGatewayId" name:"VpnGatewayId"`
		Description          *string `json:"Description" name:"Description"`
	} `json:"VpnGatewayRouteSet"`
}

func (r *DescribeVpnGatewayRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayId         *string `json:"VpnGatewayId,omitempty" name:"VpnGatewayId"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	NextHopInstanceId    *string `json:"NextHopInstanceId,omitempty" name:"NextHopInstanceId"`
	NextHopType          *string `json:"NextHopType,omitempty" name:"NextHopType"`
	Description          *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateVpnGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpnGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	RouteId   *string `json:"RouteId" name:"RouteId"`
}

func (r *CreateVpnGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpnGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayRouteId *string `json:"VpnGatewayRouteId,omitempty" name:"VpnGatewayRouteId"`
}

func (r *DeleteVpnGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpnGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteVpnGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpnGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnTunnelIpsecStatusRequest struct {
	*ksyunhttp.BaseRequest
	IsMaster    *int      `json:"IsMaster,omitempty" name:"IsMaster"`
	VpnTunnelId []*string `json:"VpnTunnelId,omitempty" name:"VpnTunnelId"`
}

func (r *DescribeVpnTunnelIpsecStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnTunnelIpsecStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeVpnTunnelIpsecStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpnTunnelIpsecStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	VpnTunnelIpsecStatusList []struct {
		IsMaster    *int    `json:"IsMaster" name:"IsMaster"`
		IpsecStatus *bool   `json:"IpsecStatus" name:"IpsecStatus"`
		IkeStatus   *bool   `json:"IkeStatus" name:"IkeStatus"`
		Id          *string `json:"Id" name:"Id"`
	} `json:"VpnTunnelIpsecStatusList"`
	Return *bool `json:"Return" name:"Return"`
}

func (r *DescribeVpnTunnelIpsecStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpnTunnelIpsecStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryNatTopVifMonitorRequest struct {
	*ksyunhttp.BaseRequest
	NatId        *string `json:"NatId,omitempty" name:"NatId"`
	StartTime    *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime      *string `json:"EndTime,omitempty" name:"EndTime"`
	SortType     *string `json:"SortType,omitempty" name:"SortType"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	Ip           *string `json:"ip,omitempty" name:"ip"`
}

func (r *QueryNatTopVifMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNatTopVifMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryNatTopVifMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryNatTopVifMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NatMonitorDataList []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		Ip           *string `json:"Ip" name:"Ip"`
		InBound      *string `json:"InBound" name:"InBound"`
		OutBound     *string `json:"OutBound" name:"OutBound"`
		InPeakBound  *string `json:"InPeakBound" name:"InPeakBound"`
		OutPeakBound *string `json:"OutPeakBound" name:"OutPeakBound"`
		MemberData   []struct {
			Timestamp     *string `json:"Timestamp" name:"Timestamp"`
			UnixTimestamp *string `json:"UnixTimestamp" name:"UnixTimestamp"`
			InBoundValue  *string `json:"InBoundValue" name:"InBoundValue"`
			OutBoundValue *string `json:"OutBoundValue" name:"OutBoundValue"`
		} `json:"MemberData" name:"MemberData"`
		Num *string `json:"Num" name:"Num"`
	} `json:"NatMonitorDataList"`
}

func (r *QueryNatTopVifMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryNatTopVifMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatIpStatusRequest struct {
	*ksyunhttp.BaseRequest
	NatIpId *string `json:"NatIpId,omitempty" name:"NatIpId"`
	Enabled *bool   `json:"Enabled,omitempty" name:"Enabled"`
}

func (r *ModifyNatIpStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatIpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNatIpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNatIpStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNatIpStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNatIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPeerTopVifMonitorRequest struct {
	*ksyunhttp.BaseRequest
	VpcId     *string `json:"vpcId,omitempty" name:"vpcId"`
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	SortType  *string `json:"SortType,omitempty" name:"SortType"`
	Ip        *string `json:"ip,omitempty" name:"ip"`
}

func (r *QueryPeerTopVifMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPeerTopVifMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryPeerTopVifMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryPeerTopVifMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId          *string `json:"RequestId" name:"RequestId"`
	NatPeerMonitorData []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		InBound      *string `json:"InBound" name:"InBound"`
		OutBound     *string `json:"OutBound" name:"OutBound"`
		Ip           *string `json:"Ip" name:"Ip"`
	} `json:"NatPeerMonitorData"`
}

func (r *QueryPeerTopVifMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryPeerTopVifMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayRouteRequest struct {
	*ksyunhttp.BaseRequest
	VpnGatewayRouteId *string `json:"VpnGatewayRouteId,omitempty" name:"VpnGatewayRouteId"`
	Description       *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyVpnGatewayRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpnGatewayRouteResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	RouteId   *bool   `json:"RouteId" name:"RouteId"`
}

func (r *ModifyVpnGatewayRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpnGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDcNatIpRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CreateDcNatIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDcNatIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateDcNatIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDcNatIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateDcNatIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDcNatIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDcNatIpRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DeleteDcNatIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDcNatIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteDcNatIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDcNatIpResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteDcNatIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDcNatIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectInterfacesBgpStatusRequest struct {
	*ksyunhttp.BaseRequest
	DirectConnectInterfaceIdN *string `json:"DirectConnectInterfaceId.N,omitempty" name:"DirectConnectInterfaceId.N"`
}

func (r *DescribeDirectConnectInterfacesBgpStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInterfacesBgpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeDirectConnectInterfacesBgpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDirectConnectInterfacesBgpStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DescribeDirectConnectInterfacesBgpStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDirectConnectInterfacesBgpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeactiveFlowLogRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DeactiveFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeactiveFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeactiveFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeactiveFlowLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	FlowLogId    *string `json:"FlowLogId" name:"FlowLogId"`
	FlowLogName  *string `json:"FlowLogName" name:"FlowLogName"`
	ResourceType *string `json:"ResourceType" name:"ResourceType"`
	ResourceId   *string `json:"ResourceId" name:"ResourceId"`
	TrafficType  *string `json:"TrafficType" name:"TrafficType"`
	ProjectName  *string `json:"ProjectName" name:"ProjectName"`
	LogPoolName  *string `json:"LogPoolName" name:"LogPoolName"`
	WindowTime   *int    `json:"WindowTime" name:"WindowTime"`
	Description  *string `json:"Description" name:"Description"`
	CreateTime   *string `json:"CreateTime" name:"CreateTime"`
}

func (r *DeactiveFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeactiveFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveFlowLogRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *ActiveFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ActiveFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ActiveFlowLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	FlowLogId    *string `json:"FlowLogId" name:"FlowLogId"`
	FlowLogName  *string `json:"FlowLogName" name:"FlowLogName"`
	ResourceType *string `json:"ResourceType" name:"ResourceType"`
	ResourceId   *string `json:"ResourceId" name:"ResourceId"`
	TrafficType  *string `json:"TrafficType" name:"TrafficType"`
	ProjectName  *string `json:"ProjectName" name:"ProjectName"`
	LogPoolName  *string `json:"LogPoolName" name:"LogPoolName"`
	WindowTime   *int    `json:"WindowTime" name:"WindowTime"`
	Description  *string `json:"Description" name:"Description"`
	CreateTime   *string `json:"CreateTime" name:"CreateTime"`
}

func (r *ActiveFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowLogRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogId *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
}

func (r *DeleteFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFlowLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Return    *bool   `json:"Return" name:"Return"`
}

func (r *DeleteFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogId   *string `json:"FlowLogId,omitempty" name:"FlowLogId"`
	FlowLogName *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	WindowTime  *int    `json:"WindowTime,omitempty" name:"WindowTime"`
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFlowLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	FlowLogId    *string `json:"FlowLogId" name:"FlowLogId"`
	FlowLogName  *string `json:"FlowLogName" name:"FlowLogName"`
	ResourceType *string `json:"ResourceType" name:"ResourceType"`
	ResourceId   *string `json:"ResourceId" name:"ResourceId"`
	TrafficType  *string `json:"TrafficType" name:"TrafficType"`
	ProjectName  *string `json:"ProjectName" name:"ProjectName"`
	LogPoolName  *string `json:"LogPoolName" name:"LogPoolName"`
	WindowTime   *int    `json:"WindowTime" name:"WindowTime"`
	Description  *string `json:"Description" name:"Description"`
	CreateTime   *string `json:"CreateTime" name:"CreateTime"`
}

func (r *ModifyFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogId  []*string                 `json:"FlowLogId,omitempty" name:"FlowLogId"`
	Filter     []*DescribeFlowLogsFilter `json:"Filter,omitempty" name:"Filter"`
	MaxResults *int                      `json:"MaxResults,omitempty" name:"MaxResults"`
	NextToken  *string                   `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	NextToken *string `json:"NextToken" name:"NextToken"`
	FlowLogs  []struct {
		FlowLogId    *string `json:"FlowLogId" name:"FlowLogId"`
		FlowLogName  *string `json:"FlowLogName" name:"FlowLogName"`
		ResourceType *string `json:"ResourceType" name:"ResourceType"`
		ResourceId   *string `json:"ResourceId" name:"ResourceId"`
		TrafficType  *string `json:"TrafficType" name:"TrafficType"`
		ProjectName  *string `json:"ProjectName" name:"ProjectName"`
		LogPoolName  *string `json:"LogPoolName" name:"LogPoolName"`
		WindowTime   *int    `json:"WindowTime" name:"WindowTime"`
		Description  *string `json:"Description" name:"Description"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
	} `json:"FlowLogs"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowLogRequest struct {
	*ksyunhttp.BaseRequest
	FlowLogName  *string `json:"FlowLogName,omitempty" name:"FlowLogName"`
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	ResourceId   *string `json:"ResourceId,omitempty" name:"ResourceId"`
	TrafficType  *string `json:"TrafficType,omitempty" name:"TrafficType"`
	ProjectName  *string `json:"ProjectName,omitempty" name:"ProjectName"`
	LogPoolName  *string `json:"LogPoolName,omitempty" name:"LogPoolName"`
	WindowTime   *int    `json:"WindowTime,omitempty" name:"WindowTime"`
	Description  *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateFlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlowLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	FlowLogId    *string `json:"FlowLogId" name:"FlowLogId"`
	FlowLogName  *string `json:"FlowLogName" name:"FlowLogName"`
	ResourceType *string `json:"ResourceType" name:"ResourceType"`
	ResourceId   *string `json:"ResourceId" name:"ResourceId"`
	TrafficType  *string `json:"TrafficType" name:"TrafficType"`
	ProjectName  *string `json:"ProjectName" name:"ProjectName"`
	LogPoolName  *string `json:"LogPoolName" name:"LogPoolName"`
	WindowTime   *int    `json:"WindowTime" name:"WindowTime"`
	Description  *string `json:"Description" name:"Description"`
	CreateTime   *string `json:"CreateTime" name:"CreateTime"`
}

func (r *CreateFlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
