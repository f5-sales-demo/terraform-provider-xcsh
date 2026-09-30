---
page_title: "xcsh_registration"
subcategory: ""
description: "xcsh_registration for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 1400, "body_sha256": "sha256:e427766fd05661a5d43db4526a3d6c3fd1eb1fb9f2472b4a71415e5094394ca1", "canonical_id": "xcsh-docs:resources:registration:fundamentals", "child_ids": ["xcsh-docs:resources:registration:reference", "xcsh-docs:resources:registration:examples", "xcsh-docs:resources:registration:import", "xcsh-docs:resources:registration:timeouts"], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:fundamentals", "parent_id": null, "path": "docs/resources/registration.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_registration for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_registration

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `token`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--registration--reference.md)
- [Examples](../guides/resources--registration--examples.md)
- [Import](../guides/resources--registration--import.md)
- [Timeouts](../guides/resources--registration--timeouts.md)
