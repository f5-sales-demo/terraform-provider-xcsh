---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 971, "body_sha256": "sha256:4445b7541f83e71318b327d2cbdbd979e11dd62762f69c724a94a8b7b1b3021e", "canonical_id": "xcsh-docs:data-sources:ike1:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f", "source_path": "examples/data-sources/xcsh_ike1/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike1:example:data-source", "parent_id": "xcsh-docs:data-sources:ike1:examples", "path": "docs/guides/data-sources--ike1--example--data-source.md", "provider_name": "ike1", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md)
- [Examples](data-sources--ike1--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike1/data-source.tf`; digest `sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f`.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```

## Next pages

- [Examples](data-sources--ike1--examples.md)
- [xcsh_ike1](../data-sources/ike1.md)
