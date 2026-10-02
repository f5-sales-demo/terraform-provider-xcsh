---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_healthcheck."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1268, "body_sha256": "sha256:f7bba23aba8b1a123f891cf8adf8b32c5d128f4cdbb7ca3157ff9c5d9f13f61a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6", "source_path": "examples/data-sources/xcsh_healthcheck/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:healthcheck:example:data-source", "parent_id": "xcsh-docs:data-sources:healthcheck:examples", "path": "documentation/data-sources/healthcheck/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2133110110222303-3000322100302020-1220223200010030-2332031302010011-1123310010231103-3112132103313010-0022002302130231-2001222111222320", "registry_path": "docs/guides/data-sources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
