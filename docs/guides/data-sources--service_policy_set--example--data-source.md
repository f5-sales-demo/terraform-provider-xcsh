---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_service_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:57be9b7d99e0845fcf6e0ce252cc751957bd4eaeee06be1e229d5fae72ee10fc", "canonical_id": "xcsh-docs:data-sources:service_policy_set:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a", "source_path": "examples/data-sources/xcsh_service_policy_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy_set:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy_set:examples", "path": "docs/guides/data-sources--service_policy_set--example--data-source.md", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_service_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_service_policy_set](../data-sources/service_policy_set.md)
- [Examples](data-sources--service_policy_set--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_set/data-source.tf`; digest `sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a`.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

## Next pages

- [Examples](data-sources--service_policy_set--examples.md)
- [xcsh_service_policy_set](../data-sources/service_policy_set.md)
