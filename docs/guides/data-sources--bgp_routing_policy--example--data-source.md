---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1050, "body_sha256": "sha256:9e6725dec82a9a5bdb78b19cb3a727a969c386b4a5a0093473e22bf6871cc5be", "canonical_id": "xcsh-docs:data-sources:bgp_routing_policy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb", "source_path": "examples/data-sources/xcsh_bgp_routing_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bgp_routing_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:examples", "path": "docs/guides/data-sources--bgp_routing_policy--example--data-source.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
- [Examples](data-sources--bgp_routing_policy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_routing_policy/data-source.tf`; digest `sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb`.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

## Next pages

- [Examples](data-sources--bgp_routing_policy--examples.md)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md)
