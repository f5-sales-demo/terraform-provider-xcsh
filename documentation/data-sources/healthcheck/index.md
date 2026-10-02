---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration."
xcsh_docs: {"aliases": ["healthcheck"], "body_bytes": 1504, "body_sha256": "sha256:74e7849ed24b58158bb47df341a2837f1910d45f027bfa5b4ae8e3478ee78e91", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:healthcheck:reference", "xcsh-docs:data-sources:healthcheck:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:healthcheck:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/healthcheck/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212", "registry_path": "docs/data-sources/healthcheck.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_healthcheck

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/healthcheck/examples/)
