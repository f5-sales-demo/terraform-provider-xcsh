---
page_title: "xcsh_smsv2_aws_runtime"
subcategory: ""
description: "xcsh_smsv2_aws_runtime for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 1684, "body_sha256": "sha256:c016538e1dbe710684bbb6cf3cdf9b3f48dd3bf8cc18b8d39b4dc00973098ae6", "canonical_id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "child_ids": ["xcsh-docs:data-sources:smsv2_aws_runtime:reference", "xcsh-docs:data-sources:smsv2_aws_runtime:examples"], "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_aws_runtime:fundamentals", "parent_id": null, "path": "docs/data-sources/smsv2_aws_runtime.md", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_smsv2_aws_runtime for xcsh_smsv2_aws_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/data-sources--smsv2_aws_runtime--reference.md)
- [Examples](../guides/data-sources--smsv2_aws_runtime--examples.md)
