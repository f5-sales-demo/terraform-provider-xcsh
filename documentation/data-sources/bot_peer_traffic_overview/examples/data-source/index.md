---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1081, "body_sha256": "sha256:33794a6602efc38ed038c93ee474ddca23a9a8fe82b195bd8979654e74197484", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11", "source_path": "examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:examples", "path": "documentation/data-sources/bot_peer_traffic_overview/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2021013333003013-2022132231103003-2110023121003113-0331223002011203-3131000012302013-2022221320002110-2200322122122320-0112033001031231", "registry_path": "docs/guides/data-sources--bot_peer_traffic_overview--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/examples/data-source/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Data source for xcsh_bot_peer_traffic_overview.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": [], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf`; digest `sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11`.

```terraform
# BotPeerTrafficOverview DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_traffic_overview" "example" {
  namespace = "example-value"
}

output "bot_peer_traffic_overview_result" {
  value = data.xcsh_bot_peer_traffic_overview.example
}
```
