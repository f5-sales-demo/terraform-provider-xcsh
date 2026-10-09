---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_alert_template."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1071, "body_sha256": "sha256:8991d1ff0795a039c6c53086883ce5c75879d0280eda97fb1b90983044956b94", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_template:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:38bcb6351cb774c27b9a7119ac259c1eaff5879d005e5da296e0eb47b5f478f7", "source_path": "examples/data-sources/xcsh_alert_template/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_template:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_template:examples", "path": "documentation/data-sources/alert_template/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2123331100131133-2222022211322210-3111100201232122-2320003033322022-1130001200303213-2132022130203100-1222110321223111-3100230233222001", "registry_path": "docs/guides/data-sources--alert_template--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_template/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_alert_template.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["alert_templateCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_template/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_template/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_template/data-source.tf`; digest `sha256:38bcb6351cb774c27b9a7119ac259c1eaff5879d005e5da296e0eb47b5f478f7`.

```terraform
# AlertTemplate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertTemplate by name
data "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"
}

output "alert_template_id" {
  value = data.xcsh_alert_template.example.id
}
```
