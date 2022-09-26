package main

//DescribeRegions

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

var (
	Ak = ""
	Sk = ""
)

func vmRegions() {
	client, err := ecs.NewClientWithAccessKey("cn-qingdao", Ak, Sk)
	if err != nil {
		panic(err)
	}

	req := ecs.CreateDescribeRegionsRequest()
	res, err := client.DescribeRegions(req)
	if err != nil {
		panic(err)
	}

	for _, r := range res.Regions.Region {
		fmt.Println(r.RegionId, r.LocalName, r.RegionEndpoint)
	}
}

func rdsRegions() {
	client, err := rds.NewClientWithAccessKey("cn-qingdao", Ak, Sk)
	if err != nil {
		panic(err)
	}

	req := rds.CreateDescribeRegionsRequest()
	res, err := client.DescribeRegions(req)
	if err != nil {
		panic(err)
	}

	for _, r := range res.Regions.RDSRegion {
		fmt.Println(r.RegionId, r.LocalName, r.RegionEndpoint)
	}
}

func vmList(regionId string) {
	client, err := ecs.NewClientWithAccessKey(regionId, Ak, Sk)
	if err != nil {
		panic(err)
	}

	req := ecs.CreateDescribeInstancesRequest()
	res, err := client.DescribeInstances(req)
	if err != nil {
		panic(err)
	}

	for _, ins := range res.Instances.Instance {
		// fmt.Printf("%+v\n", ins)
		// fmt.Println(ins.InstanceType)
		fmt.Println(ins.InstanceName)
		fmt.Println(ins.InstanceId)
		fmt.Println(ins.OSType)
		// fmt.Println(ins.OSName)
		// fmt.Println(ins.OsVersion)
		fmt.Println(ins.VpcAttributes.VpcId)
		fmt.Println(ins.PublicIpAddress.IpAddress)
		fmt.Println(ins.VpcAttributes.PrivateIpAddress.IpAddress)
	}

}

func rdsList(regionId string) {
	client, err := rds.NewClientWithAccessKey(regionId, Ak, Sk)
	if err != nil {
		panic(err)
	}

	req := rds.CreateDescribeDBInstancesRequest()
	res, err := client.DescribeDBInstances(req)
	if err != nil {
		panic(err)
	}

	for _, ins := range res.Items.DBInstance {
		fmt.Println(ins.Engine)                // 数据库类型 MySQL
		fmt.Println(ins.DBInstanceId)          // 实例ID
		fmt.Println(ins.ConnectionString)      // 实例连接地址
		fmt.Println(ins.DBInstanceDescription) // 实例描述
		// fmt.Printf("%+v\n", ins)
		// fmt.Println(ins.InstanceType)
		// fmt.Println(ins.InstanceName)
		// fmt.Println(ins.InstanceId)
		// fmt.Println(ins.OSType)
		// // fmt.Println(ins.OSName)
		// // fmt.Println(ins.OsVersion)
		// fmt.Println(ins.VpcAttributes.VpcId)
		// fmt.Println(ins.PublicIpAddress.IpAddress)
		// fmt.Println(ins.VpcAttributes.PrivateIpAddress.IpAddress)
	}
}

func main() {
	// vmRegions()
	// rdsRegions()
	// vmList("cn-guangzhou")
	rdsList("cn-beijing")
}
