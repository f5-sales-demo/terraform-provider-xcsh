---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1294, "body_sha256": "sha256:ca70428cb96249e71846fe2bd6221d2d0668cf4a07c7f6def48308fb8837dfd2", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a", "source_path": "examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:oidc_oauth_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:examples", "path": "documentation/data-sources/oidc_oauth_discovery/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0023303322201221-1101130221112332-1003310322131303-1310210330333031-0122200230000202-0020020100013300-3230230112031031-0303232301311133", "registry_path": "docs/guides/data-sources--oidc_oauth_discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_oidc_oauth_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf`; digest `sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a`.

```terraform
# OIDCOauthDiscovery DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_oidc_oauth_discovery" "example" {
  namespace = "example-value"
}

output "oidc_oauth_discovery_result" {
  value = data.xcsh_oidc_oauth_discovery.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/examples/)
- [xcsh_oidc_oauth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/oidc_oauth_discovery/)
