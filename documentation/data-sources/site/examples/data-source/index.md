---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 973, "body_sha256": "sha256:e25abd58462782badd665c4be54b7052dd3f005fbd4c5db7ccd3482d6ac4981d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348", "source_path": "examples/data-sources/xcsh_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site:example:data-source", "parent_id": "xcsh-docs:data-sources:site:examples", "path": "documentation/data-sources/site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3112312311123332-3110001102030110-2000232122320221-3102311002112103-0110201010112003-3110010033013022-3130233000000112-3122130202232102", "registry_path": "docs/guides/data-sources--site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site/data-source.tf`; digest `sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348`.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```
