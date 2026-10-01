---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_crl."
xcsh_docs: {"aliases": [], "body_bytes": 1164, "body_sha256": "sha256:36b3920214bb18a96df4f7e7518c8ca0c44155b7e4e6e734e89ed726ce45d7f2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a", "source_path": "examples/data-sources/xcsh_crl/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:crl:example:data-source", "parent_id": "xcsh-docs:data-sources:crl:examples", "path": "documentation/data-sources/crl/examples/data-source/index.md", "provider_name": "crl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_crl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_crl/data-source.tf`; digest `sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a`.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/)
