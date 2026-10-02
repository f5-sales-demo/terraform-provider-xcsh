---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1348, "body_sha256": "sha256:e6ffb38f58e7adebceb1c9806c90376c74f4251e3c5d5ef602cb45bec812ac60", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11", "source_path": "examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:examples", "path": "documentation/data-sources/bot_peer_traffic_overview/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2021013333003013-2022132231103003-2110023121003113-0331223002011203-3131000012302013-2022221320002110-2200322122122320-0112033001031231", "registry_path": "docs/guides/data-sources--bot_peer_traffic_overview--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_traffic_overview.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/examples/)
- [xcsh_bot_peer_traffic_overview](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_traffic_overview/)
