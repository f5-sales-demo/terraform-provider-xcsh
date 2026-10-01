---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1073, "body_sha256": "sha256:578068233dde178af8fdbbd2d0e9eec65453eb7ee308ef3a47d46122541fd350", "canonical_id": "xcsh-docs:data-sources:virtual_host:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72", "source_path": "examples/data-sources/xcsh_virtual_host/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_host:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_host:examples", "path": "docs/guides/data-sources--virtual_host--example--data-source.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Examples](data-sources--virtual_host--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_host/data-source.tf`; digest `sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72`.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```

## Next pages

- [Examples](data-sources--virtual_host--examples.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
