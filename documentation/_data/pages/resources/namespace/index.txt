---
page_title: "xcsh_namespace"
subcategory: ""
description: "xcsh_namespace for xcsh_namespace."
xcsh_docs: {"aliases": [], "body_bytes": 1485, "body_sha256": "sha256:099de36bbfe09bd6b9f96ad5d2ef829a169c968aeeadad9573458c8f9e66b877", "child_ids": ["xcsh-docs:resources:namespace:reference", "xcsh-docs:resources:namespace:examples", "xcsh-docs:resources:namespace:import", "xcsh-docs:resources:namespace:timeouts"], "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:resources:namespace:fundamentals", "parent_id": null, "path": "documentation/resources/namespace/index.md", "provider_name": "namespace", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_namespace for xcsh_namespace.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_namespace

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/lifecycle/timeouts/)
