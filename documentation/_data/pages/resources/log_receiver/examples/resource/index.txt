---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_log_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1233, "body_sha256": "sha256:afd45808a7e02a952e4a3b85f3e1d765420eb7bc9e8a2f95877e214bb1913f69", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d", "source_path": "examples/resources/xcsh_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:log_receiver:example:resource", "parent_id": "xcsh-docs:resources:log_receiver:examples", "path": "documentation/resources/log_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1203001122000212-1132303023100012-3031323130113333-1231323300332032-3101211023032323-2020121020301202-0331323333021131-1021011023231322", "registry_path": "docs/guides/resources--log_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/examples/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
