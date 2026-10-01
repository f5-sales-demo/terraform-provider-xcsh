---
page_title: "xcsh_workload"
subcategory: "Container"
description: "xcsh_workload for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1412, "body_sha256": "sha256:1510e641fdd262c9badd2f64444cd1e1303e4173ce4e44affbb0a649e191179b", "child_ids": ["xcsh-docs:data-sources:workload:reference", "xcsh-docs:data-sources:workload:examples"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:fundamentals", "parent_id": null, "path": "documentation/data-sources/workload/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_workload for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_workload

Breadcrumbs:

- xcsh_workload

Manages a Workload resource in F5 Distributed Cloud for workload. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_k8s`.

- virtual_k8s: Namespace for workload deployment

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/examples/)
