---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": [], "body_bytes": 1196, "body_sha256": "sha256:ce47eb9050c9ed8cb3d3932ead89e5a5b2fb7fd60da735c121f4ded36278f1fb", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp_asn_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a", "source_path": "examples/resources/xcsh_bgp_asn_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp_asn_set:example:resource", "parent_id": "xcsh-docs:resources:bgp_asn_set:examples", "path": "documentation/resources/bgp_asn_set/examples/resource/index.md", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_asn_set/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_bgp_asn_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_asn_set/resource.tf`; digest `sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/examples/)
- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/)
