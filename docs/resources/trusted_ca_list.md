---
page_title: "xcsh_trusted_ca_list"
subcategory: ""
description: "xcsh_trusted_ca_list for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": [], "body_bytes": 1343, "body_sha256": "sha256:2e24bd51eb5910acf4bf15de8613e3813477fbdd8222e092643d3835c6745095", "canonical_id": "xcsh-docs:resources:trusted_ca_list:fundamentals", "child_ids": ["xcsh-docs:resources:trusted_ca_list:reference", "xcsh-docs:resources:trusted_ca_list:examples", "xcsh-docs:resources:trusted_ca_list:import", "xcsh-docs:resources:trusted_ca_list:timeouts"], "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "id": "xcsh-docs:resources:trusted_ca_list:fundamentals", "parent_id": null, "path": "docs/resources/trusted_ca_list.md", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_trusted_ca_list for xcsh_trusted_ca_list.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_trusted_ca_list

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--trusted_ca_list--reference.md)
- [Examples](../guides/resources--trusted_ca_list--examples.md)
- [Import](../guides/resources--trusted_ca_list--import.md)
- [Timeouts](../guides/resources--trusted_ca_list--timeouts.md)
