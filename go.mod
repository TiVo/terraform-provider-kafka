module github.com/Mongey/terraform-provider-kafka

go 1.15

require (
	github.com/Shopify/sarama v1.26.4
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/hcl/v2 v2.8.1 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.3
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	github.com/zclconf/go-cty v1.7.1 // indirect
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/hashicorp/terraform-plugin-test/v2 => github.com/Mongey/terraform-plugin-test/v2 v2.1.3-0.20201231030340-31624e2320cd

