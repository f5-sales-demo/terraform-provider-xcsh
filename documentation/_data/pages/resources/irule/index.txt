---
page_title: "xcsh_irule"
subcategory: ""
description: "xcsh_irule for xcsh_irule."
xcsh_docs: {"aliases": [], "body_bytes": 1560, "body_sha256": "sha256:79e888af43823d577f5b2022d46b4a5c389005817bd915a638df887582b26525", "child_ids": ["xcsh-docs:resources:irule:reference", "xcsh-docs:resources:irule:examples", "xcsh-docs:resources:irule:import", "xcsh-docs:resources:irule:timeouts"], "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "id": "xcsh-docs:resources:irule:fundamentals", "parent_id": null, "path": "documentation/resources/irule/index.md", "provider_name": "irule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_irule for xcsh_irule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_irule

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

## Root configuration

Required root properties: `description_spec`, `irule`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/lifecycle/timeouts/)
