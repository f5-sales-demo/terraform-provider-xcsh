---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_top_reason_codes."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1347, "body_sha256": "sha256:63591921fc00474c18af338ca6379d9589880893bda9e52c1ef7da8a4bd42611", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516", "source_path": "examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:examples", "path": "documentation/data-sources/bot_peer_top_reason_codes/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2230133111230322-1202023033331030-3200213012023022-3302133031002223-3110033012111233-0101223211133032-1310121002113011-2023333223200301", "registry_path": "docs/guides/data-sources--bot_peer_top_reason_codes--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_bot_peer_top_reason_codes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf`; digest `sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516`.

```terraform
# BotPeerTopReasonCodes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_reason_codes" "example" {
  namespace = "example-value"
}

output "bot_peer_top_reason_codes_result" {
  value = data.xcsh_bot_peer_top_reason_codes.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/examples/)
- [xcsh_bot_peer_top_reason_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/)
