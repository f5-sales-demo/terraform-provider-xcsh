---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_bot_defense."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1381, "body_sha256": "sha256:da86f81a376541aa4a8007bbd89479d93e19f4bf6c32095641fb9ccafbe149cb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_bot_defense:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081", "source_path": "examples/data-sources/xcsh_network_bot_defense/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_bot_defense:example:data-source", "parent_id": "xcsh-docs:data-sources:network_bot_defense:examples", "path": "documentation/data-sources/network_bot_defense/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_bot_defense", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2110231032332131-1321211201201120-0310201033330000-2032302212220211-3330213230110222-0011120001300010-1101323001121133-0200120122011032", "registry_path": "docs/guides/data-sources--network_bot_defense--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_bot_defense/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_network_bot_defense.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
