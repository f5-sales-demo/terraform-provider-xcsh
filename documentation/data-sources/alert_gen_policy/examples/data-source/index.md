---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_alert_gen_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1329, "body_sha256": "sha256:6c944b18f848ae6df3963ed669d4576d59321da24aa94cf3d511c19ef9bb32dd", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:alert_gen_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed", "source_path": "examples/data-sources/xcsh_alert_gen_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:alert_gen_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:alert_gen_policy:examples", "path": "documentation/data-sources/alert_gen_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "alert_gen_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1103023110200103-2022110230331111-2121301210003300-0333200010031033-2013330110211221-2313110133210132-2012222310222220-2032213300022312", "registry_path": "docs/guides/data-sources--alert_gen_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_gen_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_alert_gen_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["alert_gen_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_gen_policy/data-source.tf`; digest `sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed`.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/examples/)
- [xcsh_alert_gen_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_gen_policy/)
