---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_healthcheck."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1043, "body_sha256": "sha256:ea72f6ec9eb015685e62d349df4e88e5075fa4e7cdba2682f6b6425e1f0efa0f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6", "source_path": "examples/data-sources/xcsh_healthcheck/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:healthcheck:example:data-source", "parent_id": "xcsh-docs:data-sources:healthcheck:examples", "path": "documentation/data-sources/healthcheck/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2133110110222303-3000322100302020-1220223200010030-2332031302010011-1123310010231103-3112132103313010-0022002302130231-2001222111222320", "registry_path": "docs/guides/data-sources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
