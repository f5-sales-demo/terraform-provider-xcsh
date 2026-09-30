---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 859, "body_sha256": "sha256:6c8eb5f13760d3afdc7bdada5ecd288adef231350018623e5852f737e76a0743", "canonical_id": "xcsh-docs:data-sources:bgp:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b", "source_path": "examples/data-sources/xcsh_bgp/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bgp:example:data-source", "parent_id": "xcsh-docs:data-sources:bgp:examples", "path": "docs/guides/data-sources--bgp--example--data-source.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md)
- [Examples](data-sources--bgp--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp/data-source.tf`; digest `sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b`.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```

## Next pages

- [Examples](data-sources--bgp--examples.md)
- [xcsh_bgp](../data-sources/bgp.md)
