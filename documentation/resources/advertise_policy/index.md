---
page_title: "xcsh_advertise_policy"
subcategory: ""
description: "xcsh_advertise_policy for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1824, "body_sha256": "sha256:e9215049908a59cc2e2e83ba2a9c7d8613cec7e6dfcec1ff54b7c319a23a3515", "child_ids": ["xcsh-docs:resources:advertise_policy:reference", "xcsh-docs:resources:advertise_policy:examples", "xcsh-docs:resources:advertise_policy:import", "xcsh-docs:resources:advertise_policy:timeouts"], "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:fundamentals", "parent_id": null, "path": "documentation/resources/advertise_policy/index.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_advertise_policy for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_advertise_policy

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/lifecycle/timeouts/)
