---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1062, "body_sha256": "sha256:bccc04045de3316a3bdb773ce43e5f43ecf7e410c534fb366b53dd22838f3950", "canonical_id": "xcsh-docs:data-sources:certificate:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e", "source_path": "examples/data-sources/xcsh_certificate/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate:examples", "path": "docs/guides/data-sources--certificate--example--data-source.md", "provider_name": "certificate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md)
- [Examples](data-sources--certificate--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate/data-source.tf`; digest `sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e`.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```

## Next pages

- [Examples](data-sources--certificate--examples.md)
- [xcsh_certificate](../data-sources/certificate.md)
