---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_advertise_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:c97e007277bb97a4417d6859efe042c8671cfd46b2f969682caf4bc6c8d14730", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8", "source_path": "examples/data-sources/xcsh_advertise_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:advertise_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:advertise_policy:examples", "path": "documentation/data-sources/advertise_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1233231313132130-3332030313113030-2321110130210211-3203001212202223-0033001301103321-2201031123102221-0331213233033331-1313122120323121", "registry_path": "docs/guides/data-sources--advertise_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_advertise_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_advertise_policy/data-source.tf`; digest `sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8`.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/examples/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
