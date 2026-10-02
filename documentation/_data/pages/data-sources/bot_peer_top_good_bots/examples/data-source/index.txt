---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_top_good_bots."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1314, "body_sha256": "sha256:222557e4014689ec357494e5de347c944a50f87105ecd6914ab40f0beb31379a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cbd5508dce9276d9af60eef5a236b141ce290117e051aad973c941688a149920", "source_path": "examples/data-sources/xcsh_bot_peer_top_good_bots/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_top_good_bots:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_top_good_bots:examples", "path": "documentation/data-sources/bot_peer_top_good_bots/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_top_good_bots", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1012323120023221-1201133133303300-1201113000133113-2120332011222220-3132023213221212-2202301130031111-2300232320200221-0013111233133031", "registry_path": "docs/guides/data-sources--bot_peer_top_good_bots--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_good_bots/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_top_good_bots.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_good_bots/data-source.tf`; digest `sha256:cbd5508dce9276d9af60eef5a236b141ce290117e051aad973c941688a149920`.

```terraform
# BotPeerTopGoodBots DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_good_bots" "example" {
  namespace = "example-value"
}

output "bot_peer_top_good_bots_result" {
  value = data.xcsh_bot_peer_top_good_bots.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/examples/)
- [xcsh_bot_peer_top_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_top_good_bots/)
