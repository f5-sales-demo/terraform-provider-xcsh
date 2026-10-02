---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_policy_based_routing."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1381, "body_sha256": "sha256:3e92c439e7a7fa7aeb36bd37955b79b287deb50bb820fac2b2bd5ff285fe8434", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc", "source_path": "examples/data-sources/xcsh_policy_based_routing/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:policy_based_routing:example:data-source", "parent_id": "xcsh-docs:data-sources:policy_based_routing:examples", "path": "documentation/data-sources/policy_based_routing/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0322032233221030-1213230211302320-2132120322330020-2331203100000203-0332212112033233-2003312020200311-3321022120333323-2311012110221103", "registry_path": "docs/guides/data-sources--policy_based_routing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_policy_based_routing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policy_based_routing/data-source.tf`; digest `sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc`.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/examples/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
