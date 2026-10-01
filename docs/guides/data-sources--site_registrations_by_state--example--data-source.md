---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 1153, "body_sha256": "sha256:7adc4719e6ff8f891a1ffbbf2fa9d8ea44970d74d3e3e8d19a8d9acd590287f2", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a", "source_path": "examples/data-sources/xcsh_site_registrations_by_state/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registrations_by_state:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:examples", "path": "docs/guides/data-sources--site_registrations_by_state--example--data-source.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Examples](data-sources--site_registrations_by_state--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_state/data-source.tf`; digest `sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a`.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```

## Next pages

- [Examples](data-sources--site_registrations_by_state--examples.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
