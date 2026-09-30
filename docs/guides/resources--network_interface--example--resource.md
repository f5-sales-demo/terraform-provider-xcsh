---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1055, "body_sha256": "sha256:c9b14eb661cbfb8619de0079fc92a446a6715d2f3bc583c7e0855cdb835ad18d", "canonical_id": "xcsh-docs:resources:network_interface:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5", "source_path": "examples/resources/xcsh_network_interface/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_interface:example:resource", "parent_id": "xcsh-docs:resources:network_interface:examples", "path": "docs/guides/resources--network_interface--example--resource.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Examples](resources--network_interface--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_interface/resource.tf`; digest `sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5`.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--network_interface--examples.md)
- [xcsh_network_interface](../resources/network_interface.md)
