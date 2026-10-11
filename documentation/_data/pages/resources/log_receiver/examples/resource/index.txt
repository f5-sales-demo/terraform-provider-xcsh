---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_log_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1011, "body_sha256": "sha256:8436b56dfe2bb58108aca057be8329f8996d955124ee448fe58c1e015a0fba66", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d", "source_path": "examples/resources/xcsh_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:log_receiver:example:resource", "parent_id": "xcsh-docs:resources:log_receiver:examples", "path": "documentation/resources/log_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1203001122000212-1132303023100012-3031323130113333-1231323300332032-3101211023032323-2020121020301202-0331323333021131-1021011023231322", "registry_path": "docs/guides/resources--log_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_log_receiver/resource.tf`; digest `sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d`.

```terraform
# LogReceiver Resource Example
# Manages new Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic LogReceiver configuration
resource "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}
```
