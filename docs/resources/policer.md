---
page_title: "xcsh_policer"
subcategory: ""
description: "xcsh_policer for xcsh_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1298, "body_sha256": "sha256:91dc47d6771944ab76ef6df5846243b4aaea4cdd84ec6b5a6ec4c5595cc6af3d", "canonical_id": "xcsh-docs:resources:policer:fundamentals", "child_ids": ["xcsh-docs:resources:policer:reference", "xcsh-docs:resources:policer:examples", "xcsh-docs:resources:policer:import", "xcsh-docs:resources:policer:timeouts"], "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:fundamentals", "parent_id": null, "path": "docs/resources/policer.md", "provider_name": "policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_policer for xcsh_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_policer

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```

## Root configuration

Required root properties: `burst_size`, `committed_information_rate`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--policer--reference.md)
- [Examples](../guides/resources--policer--examples.md)
- [Import](../guides/resources--policer--import.md)
- [Timeouts](../guides/resources--policer--timeouts.md)
