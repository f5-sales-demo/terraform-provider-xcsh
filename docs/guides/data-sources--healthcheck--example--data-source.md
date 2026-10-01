---
page_title: "Data source"
subcategory: "Monitoring"
description: "Data source for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1062, "body_sha256": "sha256:356d3e2fa8cf3185bb4e7751b3eea330a8e01080b008beb511ccda76f523c7f3", "canonical_id": "xcsh-docs:data-sources:healthcheck:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6", "source_path": "examples/data-sources/xcsh_healthcheck/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:healthcheck:example:data-source", "parent_id": "xcsh-docs:data-sources:healthcheck:examples", "path": "docs/guides/data-sources--healthcheck--example--data-source.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/healthcheck/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md)
- [Examples](data-sources--healthcheck--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_healthcheck/data-source.tf`; digest `sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6`.

```terraform
# Healthcheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Healthcheck by name
data "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"
}

output "healthcheck_id" {
  value = data.xcsh_healthcheck.example.id
}
```

## Next pages

- [Examples](data-sources--healthcheck--examples.md)
- [xcsh_healthcheck](../data-sources/healthcheck.md)
