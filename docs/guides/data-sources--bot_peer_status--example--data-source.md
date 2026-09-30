---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_status."
xcsh_docs: {"aliases": [], "body_bytes": 934, "body_sha256": "sha256:5d1b4da1ecddcce306d94c4d30ab07f5546862884c963688ed98fa2ecec9874e", "canonical_id": "xcsh-docs:data-sources:bot_peer_status:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc", "source_path": "examples/data-sources/xcsh_bot_peer_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_status:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_status:examples", "path": "docs/guides/data-sources--bot_peer_status--example--data-source.md", "provider_name": "bot_peer_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md)
- [Examples](data-sources--bot_peer_status--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_status/data-source.tf`; digest `sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc`.

```terraform
# BotPeerStatus DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_status" "example" {
  namespace = "example-value"
}

output "bot_peer_status_result" {
  value = data.xcsh_bot_peer_status.example
}
```

## Next pages

- [Examples](data-sources--bot_peer_status--examples.md)
- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md)
