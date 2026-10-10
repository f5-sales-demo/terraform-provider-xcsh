---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_smsv2_aws_runtime."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1455, "body_sha256": "sha256:da6752f7bb495fc34a9e8ab51876f323f775eee5011e997a928d0ca7d39e2556", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_aws_runtime:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e299ec8ab27b0aebf0736cf9e2b336fdc7524df91d37a6ad7dac1a9e71242dd5", "source_path": "examples/data-sources/xcsh_smsv2_aws_runtime/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:smsv2_aws_runtime:example:data-source", "parent_id": "xcsh-docs:data-sources:smsv2_aws_runtime:examples", "path": "documentation/data-sources/smsv2_aws_runtime/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "smsv2_aws_runtime", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1220032022202012-2231333113102233-0210321323030212-0110131031312222-2033303300200222-3221110300202013-0031301303112132-2100130123303312", "registry_path": "docs/guides/data-sources--smsv2_aws_runtime--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_aws_runtime/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_smsv2_aws_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_smsv2_aws_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_aws_runtime/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_aws_runtime/data-source.tf`; digest `sha256:e299ec8ab27b0aebf0736cf9e2b336fdc7524df91d37a6ad7dac1a9e71242dd5`.

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
