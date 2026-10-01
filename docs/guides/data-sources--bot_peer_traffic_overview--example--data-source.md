---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_traffic_overview."
xcsh_docs: {"aliases": [], "body_bytes": 1142, "body_sha256": "sha256:03d0e4cdaf1f7f6f16147de067e09092f20eebfff23259ea9a372bc124bc33ba", "canonical_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02ae4e984486512606c527da0d434568ac636c951e2a90ff31f351d9dd36be11", "source_path": "examples/data-sources/xcsh_bot_peer_traffic_overview/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_traffic_overview:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_traffic_overview:examples", "path": "docs/guides/data-sources--bot_peer_traffic_overview--example--data-source.md", "provider_name": "bot_peer_traffic_overview", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_traffic_overview/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_traffic_overview.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md)
- [Examples](data-sources--bot_peer_traffic_overview--examples.md)
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

- [Examples](data-sources--bot_peer_traffic_overview--examples.md)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md)
