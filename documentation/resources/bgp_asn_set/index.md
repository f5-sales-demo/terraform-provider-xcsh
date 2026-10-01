---
page_title: "xcsh_bgp_asn_set"
subcategory: ""
description: "xcsh_bgp_asn_set for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": [], "body_bytes": 1631, "body_sha256": "sha256:a96b736af01171fe868cde027db8bc36d8176613f5e7887efdb0a26b5e1ec2b0", "child_ids": ["xcsh-docs:resources:bgp_asn_set:reference", "xcsh-docs:resources:bgp_asn_set:examples", "xcsh-docs:resources:bgp_asn_set:import", "xcsh-docs:resources:bgp_asn_set:timeouts"], "collection_id": "xcsh-docs:resources:bgp_asn_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_asn_set:fundamentals", "parent_id": null, "path": "documentation/resources/bgp_asn_set/index.md", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_asn_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bgp_asn_set for xcsh_bgp_asn_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp_asn_set

Breadcrumbs:

- xcsh_bgp_asn_set

Manages bgp\_asn\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPAsnSet configuration
resource "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"

  as_numbers = [1]
}
```

## Root configuration

Required root properties: `as_numbers`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/lifecycle/timeouts/)
