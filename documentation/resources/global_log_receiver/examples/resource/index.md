---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_global_log_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1315, "body_sha256": "sha256:9a559c57808455cceba3270f10556f3db32811e6dc151f7186bb9ca099830e0b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802", "source_path": "examples/resources/xcsh_global_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:global_log_receiver:example:resource", "parent_id": "xcsh-docs:resources:global_log_receiver:examples", "path": "documentation/resources/global_log_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1302210003103331-0230103103033310-2330122030113120-1222212200232110-1031133321100222-3310031002121112-0030230022033313-1211010002100121", "registry_path": "docs/guides/resources--global_log_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_global_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_global_log_receiver/resource.tf`; digest `sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802`.

```terraform
# GlobalLogReceiver Resource Example
# Manages new Global Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GlobalLogReceiver configuration
resource "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/examples/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
