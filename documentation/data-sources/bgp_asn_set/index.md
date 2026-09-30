---
page_title: "xcsh_bgp_asn_set"
subcategory: ""
description: "xcsh_bgp_asn_set for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": [], "body_bytes": 1249, "body_sha256": "sha256:7b7f72ce9c82909a794dbc02f62cdb9e73a0875eb9a01b12764f2c251e2048dc", "child_ids": ["xcsh-docs:data-sources:bgp_asn_set:reference", "xcsh-docs:data-sources:bgp_asn_set:examples"], "collection_id": "xcsh-docs:data-sources:bgp_asn_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_asn_set:fundamentals", "parent_id": null, "path": "documentation/data-sources/bgp_asn_set/index.md", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_asn_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bgp_asn_set for xcsh_bgp_asn_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/examples/)
