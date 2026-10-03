---
page_title: "xcsh_smsv2_aws_runtime"
subcategory: ""
description: "Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published physical-link status."
xcsh_docs: {"aliases": ["smsv2 aws runtime"], "body_bytes": 1769, "body_sha256": "sha256:295257e2067ee5706138fe868ce996b666fd35aff2ed113a699910057c8998d7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:reference", "xcsh-docs:data-sources:smsv2_aws_runtime:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/smsv2_aws_runtime/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1120113232301113-0213023023132120-1120122010310202-0111213100202201-0123213030213220-2303201220111233-2220132332120003-3013101222121310", "registry_path": "docs/data-sources/smsv2_aws_runtime.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published physical-link status.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_smsv2_aws_runtime

Breadcrumbs:

- xcsh_smsv2_aws_runtime

Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published
physical-link status.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Correlate stable logical node keys and AWS-authoritative ENI MAC addresses
# with the SMSv2 interface configuration and runtime health observed by F5 XC.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_aws_runtime" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  nodes = {
    node_0_slo = {
      node = "node-0"
      role = "slo"
      mac  = "02:00:00:00:00:10"
    }
    node_0_sli = {
      node = "node-0"
      role = "sli"
      mac  = "02:00:00:00:00:11"
    }
  }
}

output "smsv2_interfaces" {
  value = data.xcsh_smsv2_aws_runtime.site.interfaces
}

output "smsv2_healthy" {
  value = data.xcsh_smsv2_aws_runtime.site.healthy
}
```

## Root configuration

Required root properties: `namespace`, `nodes`, `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/examples/)
