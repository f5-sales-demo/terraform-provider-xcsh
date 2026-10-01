---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1088, "body_sha256": "sha256:de80374f0e7474ee935f2cf3ff26e0d6c04284b98097d415fff036fb0818841e", "canonical_id": "xcsh-docs:data-sources:oidc_oauth_discovery:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a", "source_path": "examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:oidc_oauth_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:examples", "path": "docs/guides/data-sources--oidc_oauth_discovery--example--data-source.md", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_oidc_oauth_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
- [Examples](data-sources--oidc_oauth_discovery--examples.md)
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

- [Examples](data-sources--oidc_oauth_discovery--examples.md)
- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md)
