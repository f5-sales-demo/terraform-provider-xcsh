---
page_title: "xcsh_bgp"
subcategory: ""
description: "xcsh_bgp for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1409, "body_sha256": "sha256:ca92850b7fcd0a69f3c3ce274b19b284661fcb932a403b21231a2011a2c65c19", "canonical_id": "xcsh-docs:resources:bgp:fundamentals", "child_ids": ["xcsh-docs:resources:bgp:reference", "xcsh-docs:resources:bgp:examples", "xcsh-docs:resources:bgp:import", "xcsh-docs:resources:bgp:timeouts"], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:fundamentals", "parent_id": null, "path": "docs/resources/bgp.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bgp for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--bgp--reference.md)
- [Examples](../guides/resources--bgp--examples.md)
- [Import](../guides/resources--bgp--import.md)
- [Timeouts](../guides/resources--bgp--timeouts.md)
