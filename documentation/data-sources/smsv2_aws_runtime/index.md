---
page_title: "xcsh_smsv2_aws_runtime"
subcategory: ""
description: "Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published physical-link status."
xcsh_docs: {"aliases": ["smsv2 aws runtime"], "body_bytes": 1782, "body_sha256": "sha256:8ee38dc703a90a78daec47d699acc88e84933ebfccb2c205052e2af53820ae4d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:reference", "xcsh-docs:data-sources:smsv2_aws_runtime:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/smsv2_aws_runtime/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1120113232301113-0213023023132120-1120122010310202-0111213100202201-0123213030213220-2303201220111233-2220132332120003-3013101222121310", "registry_path": "docs/data-sources/smsv2_aws_runtime.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published physical-link status.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": [], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/examples/)
