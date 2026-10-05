---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_alert_template."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1305, "body_sha256": "sha256:ac0b398d0b8f78ecd76ff842d771eb7f29d14099619cf183b76d0ce5bb4a647e", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_template:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:38bcb6351cb774c27b9a7119ac259c1eaff5879d005e5da296e0eb47b5f478f7", "source_path": "examples/data-sources/xcsh_alert_template/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_template:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_template:examples", "path": "documentation/data-sources/alert_template/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "alert_template", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2123331100131133-2222022211322210-3111100201232122-2320003033322022-1130001200303213-2132022130203100-1222110321223111-3100230233222001", "registry_path": "docs/guides/data-sources--alert_template--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_template/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_alert_template.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["alert_templateCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_template/examples/)
- [xcsh_alert_template](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_template/)
