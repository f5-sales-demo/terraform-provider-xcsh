---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_oidc_oauth_discovery."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1042, "body_sha256": "sha256:243e9639af8ec310fc2b9f242b4c9d0417f9612e2bab61d1731ff3d41b1a2426", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:oidc_oauth_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a", "source_path": "examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:oidc_oauth_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:oidc_oauth_discovery:examples", "path": "documentation/data-sources/oidc_oauth_discovery/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "oidc_oauth_discovery", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0023303322201221-1101130221112332-1003310322131303-1310210330333031-0122200230000202-0020020100013300-3230230112031031-0303232301311133", "registry_path": "docs/guides/data-sources--oidc_oauth_discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/oidc_oauth_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_oidc_oauth_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
