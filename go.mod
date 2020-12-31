module github.com/Mongey/terraform-provider-kafka

go 1.12

require (
	github.com/Shopify/sarama v1.26.4
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/terraform v0.14.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.3
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
