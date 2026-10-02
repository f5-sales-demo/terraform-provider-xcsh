---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_service_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1305, "body_sha256": "sha256:c126a2b55d6812107c6c4037e3fe63cf33de8c5a96022eb53a0056d40987e7df", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9189dd7384a926dd448e27affcd3a13f4fb5015e7cdfa0bb42501f00fe48f58d", "source_path": "examples/data-sources/xcsh_service_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy:examples", "path": "documentation/data-sources/service_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3133021133101133-1331033230023211-3202032232200331-2023300021323301-0130231331111303-2301211203303113-2131112300200131-1331020312233013", "registry_path": "docs/guides/data-sources--service_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_service_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/examples/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
