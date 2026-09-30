---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": [], "body_bytes": 1011, "body_sha256": "sha256:d9e03928dbb069bdf749ece33af07f69228a39c2d3fb94c123a5220975eb0e6b", "canonical_id": "xcsh-docs:data-sources:trusted_ca_list:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7", "source_path": "examples/data-sources/xcsh_trusted_ca_list/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:trusted_ca_list:example:data-source", "parent_id": "xcsh-docs:data-sources:trusted_ca_list:examples", "path": "docs/guides/data-sources--trusted_ca_list--example--data-source.md", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_trusted_ca_list.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md)
- [Examples](data-sources--trusted_ca_list--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

## Next pages

- [Examples](data-sources--trusted_ca_list--examples.md)
- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md)
