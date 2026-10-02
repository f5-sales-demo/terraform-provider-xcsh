---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_threat_types."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1304, "body_sha256": "sha256:821b65f3fbf51c09c060ccf17f3f78b3e7b936715fc3350fcf5eca24a1d669f0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_peer_threat_types:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa", "source_path": "examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_threat_types:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_threat_types:examples", "path": "documentation/data-sources/bot_peer_threat_types/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_peer_threat_types", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3303302102011220-1331032122300301-3302330113313021-3120312120332211-0110023003331220-2003113313201210-2011131202320023-2303030320200122", "registry_path": "docs/guides/data-sources--bot_peer_threat_types--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_threat_types/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bot_peer_threat_types.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_threat_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf`; digest `sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa`.

```terraform
# BotPeerThreatTypes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_threat_types" "example" {
  namespace = "example-value"
}

output "bot_peer_threat_types_result" {
  value = data.xcsh_bot_peer_threat_types.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/examples/)
- [xcsh_bot_peer_threat_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_peer_threat_types/)
