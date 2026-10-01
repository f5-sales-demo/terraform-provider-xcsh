---
page_title: "xcsh_ip_prefix_set"
subcategory: ""
description: "xcsh_ip_prefix_set for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": [], "body_bytes": 1432, "body_sha256": "sha256:729b667efe68b6cd7f66d650e78e3cd8c1416ffd504dd35c6b719a7c646d846a", "canonical_id": "xcsh-docs:resources:ip_prefix_set:fundamentals", "child_ids": ["xcsh-docs:resources:ip_prefix_set:reference", "xcsh-docs:resources:ip_prefix_set:examples", "xcsh-docs:resources:ip_prefix_set:import", "xcsh-docs:resources:ip_prefix_set:timeouts"], "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:ip_prefix_set:fundamentals", "parent_id": null, "path": "docs/resources/ip_prefix_set.md", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ip_prefix_set for xcsh_ip_prefix_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ip_prefix_set

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--ip_prefix_set--reference.md)
- [Examples](../guides/resources--ip_prefix_set--examples.md)
- [Import](../guides/resources--ip_prefix_set--import.md)
- [Timeouts](../guides/resources--ip_prefix_set--timeouts.md)
