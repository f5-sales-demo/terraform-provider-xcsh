---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_bot_defense."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1381, "body_sha256": "sha256:da86f81a376541aa4a8007bbd89479d93e19f4bf6c32095641fb9ccafbe149cb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081", "source_path": "examples/data-sources/xcsh_network_bot_defense/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_bot_defense:example:data-source", "parent_id": "xcsh-docs:data-sources:network_bot_defense:examples", "path": "documentation/data-sources/network_bot_defense/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2110231032332131-1321211201201120-0310201033330000-2032302212220211-3330213230110222-0011120001300010-1101323001121133-0200120122011032", "registry_path": "docs/guides/data-sources--network_bot_defense--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_bot_defense.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_bot_defense/data-source.tf`; digest `sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/examples/)
- [xcsh_network_bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_bot_defense/)
