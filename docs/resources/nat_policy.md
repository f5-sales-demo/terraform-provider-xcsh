---
page_title: "xcsh_nat_policy"
subcategory: ""
description: "xcsh_nat_policy for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1358, "body_sha256": "sha256:99e2f633422aec0d987e05fa6cceecf0e66d01c973c603419ea9cde4cfad44ba", "canonical_id": "xcsh-docs:resources:nat_policy:fundamentals", "child_ids": ["xcsh-docs:resources:nat_policy:reference", "xcsh-docs:resources:nat_policy:examples", "xcsh-docs:resources:nat_policy:import", "xcsh-docs:resources:nat_policy:timeouts"], "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:fundamentals", "parent_id": null, "path": "docs/resources/nat_policy.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_nat_policy for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_nat_policy

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--nat_policy--reference.md)
- [Examples](../guides/resources--nat_policy--examples.md)
- [Import](../guides/resources--nat_policy--import.md)
- [Timeouts](../guides/resources--nat_policy--timeouts.md)
