---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_healthcheck."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1268, "body_sha256": "sha256:f7bba23aba8b1a123f891cf8adf8b32c5d128f4cdbb7ca3157ff9c5d9f13f61a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6", "source_path": "examples/data-sources/xcsh_healthcheck/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:healthcheck:example:data-source", "parent_id": "xcsh-docs:data-sources:healthcheck:examples", "path": "documentation/data-sources/healthcheck/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2133110110222303-3000322100302020-1220223200010030-2332031302010011-1123310010231103-3112132103313010-0022002302130231-2001222111222320", "registry_path": "docs/guides/data-sources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/examples/data-source/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Data source for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_healthcheck/data-source.tf`; digest `sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6`.

```terraform
# Healthcheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Healthcheck by name
data "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"
}

output "healthcheck_id" {
  value = data.xcsh_healthcheck.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/examples/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/)
