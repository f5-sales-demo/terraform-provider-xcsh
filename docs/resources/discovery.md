---
page_title: "xcsh_discovery"
subcategory: ""
description: "xcsh_discovery for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:f9ab2a61a64db3fe79b8df5b9cfeff2afd5ec13cfcd9293de343182963fcf1be", "canonical_id": "xcsh-docs:resources:discovery:fundamentals", "child_ids": ["xcsh-docs:resources:discovery:reference", "xcsh-docs:resources:discovery:examples", "xcsh-docs:resources:discovery:import", "xcsh-docs:resources:discovery:timeouts"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:fundamentals", "parent_id": null, "path": "docs/resources/discovery.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_discovery for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_discovery

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--discovery--reference.md)
- [Examples](../guides/resources--discovery--examples.md)
- [Import](../guides/resources--discovery--import.md)
- [Timeouts](../guides/resources--discovery--timeouts.md)
