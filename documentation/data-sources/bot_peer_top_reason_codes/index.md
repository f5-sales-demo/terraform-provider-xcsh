---
page_title: "xcsh_bot_peer_top_reason_codes"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["bot peer top reason codes"], "body_bytes": 1296, "body_sha256": "sha256:8914aeec505b0a439df612b7c1bd51b8c623ae2260e152990bfad6ab10346a35", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_peer_top_reason_codes:reference", "xcsh-docs:data-sources:bot_peer_top_reason_codes:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_peer_top_reason_codes/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0013220013020012-2103133222320300-2320112130123102-0221023121031312-3010203123123211-1231322223322030-3232013030300102-1302220311210333", "registry_path": "docs/data-sources/bot_peer_top_reason_codes.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_peer_top_reason_codes

Breadcrumbs:

- xcsh_bot_peer_top_reason_codes

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_reason_codes/examples/)
