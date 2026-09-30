---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1000, "body_sha256": "sha256:fe681a21fa1179e421589188fc8487a9d60c3270d5ac9fb02ebb793e47bb066d", "canonical_id": "xcsh-docs:data-sources:service_policy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9189dd7384a926dd448e27affcd3a13f4fb5015e7cdfa0bb42501f00fe48f58d", "source_path": "examples/data-sources/xcsh_service_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy:examples", "path": "docs/guides/data-sources--service_policy--example--data-source.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Examples](data-sources--service_policy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy/data-source.tf`; digest `sha256:9189dd7384a926dd448e27affcd3a13f4fb5015e7cdfa0bb42501f00fe48f58d`.

```terraform
# ServicePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicy by name
data "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}

output "service_policy_id" {
  value = data.xcsh_service_policy.example.id
}
```

## Next pages

- [Examples](data-sources--service_policy--examples.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
