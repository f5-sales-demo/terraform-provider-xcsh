---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1266, "body_sha256": "sha256:db5e79e03fc4735e6ae7a0fe97ab75cde175d8d304420fcf2126bc8bfc03658e", "child_ids": [], "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd", "source_path": "examples/data-sources/xcsh_api_testing/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_testing:example:data-source", "parent_id": "xcsh-docs:data-sources:api_testing:examples", "path": "documentation/data-sources/api_testing/examples/data-source/index.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_testing/data-source.tf`; digest `sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd`.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/examples/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
