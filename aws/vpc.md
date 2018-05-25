* [亚马逊AWS学习——VPC里面几个概念的关系](https://www.cnblogs.com/mrchige/p/5916632.html)
* [官方网页](https://amazonaws-china.com/cn/vpc/)

# 选型
* [创建VPC向导](https://ap-southeast-1.console.aws.amazon.com/vpc/home?region=ap-southeast-1#wizardSelector:)

  带有公有和私有子网的VPC：除了包含公有子网之外，此配置还添加了一个私有子网，该子网的实例无法从 Internet 寻址。私有子网中的实例可以使用 Network Address Translation (NAT) 通过公有子网与 Internet 建立出站连接。
  
  **创建**:
  
  具有两个 /24 子网的 /16 网络。公共子网实例，使用弹性 IP 地址访问 Internet。私有子网实例通过 Network Address Translation (NAT) 实例访问 Internet。 (NAT 实例按小时收费)
