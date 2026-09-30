---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_advertise_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1232, "body_sha256": "sha256:447f6f1b99d272b9e03690dc0ec4119e85bf4600132ffd4bfc90af2fe4b64e29", "child_ids": [], "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8", "source_path": "examples/data-sources/xcsh_advertise_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:advertise_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:advertise_policy:examples", "path": "documentation/data-sources/advertise_policy/examples/data-source/index.md", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_advertise_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_advertise_policy/data-source.tf`; digest `sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8`.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/examples/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
