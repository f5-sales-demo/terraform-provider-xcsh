---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_global_log_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1072, "body_sha256": "sha256:e1e9cac478cb9e1e56fd7eedc6a41d619a761a950bad859966be60abf9d4839b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1d1066d34a2e865fbb60f6e1a0d99d1dfeb43aa5b8be24b88b8a109599e81802", "source_path": "examples/resources/xcsh_global_log_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:global_log_receiver:example:resource", "parent_id": "xcsh-docs:resources:global_log_receiver:examples", "path": "documentation/resources/global_log_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1302210003103331-0230103103033310-2330122030113120-1222212200232110-1031133321100222-3310031002121112-0030230022033313-1211010002100121", "registry_path": "docs/guides/resources--global_log_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_global_log_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
