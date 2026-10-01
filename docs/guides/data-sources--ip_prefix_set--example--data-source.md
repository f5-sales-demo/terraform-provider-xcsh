---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": [], "body_bytes": 1084, "body_sha256": "sha256:9232023330d0a37f023348b7fffc9bbb42c65802580b3723bcfb0a7274961015", "canonical_id": "xcsh-docs:data-sources:ip_prefix_set:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b", "source_path": "examples/data-sources/xcsh_ip_prefix_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ip_prefix_set:example:data-source", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:examples", "path": "docs/guides/data-sources--ip_prefix_set--example--data-source.md", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_ip_prefix_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md)
- [Examples](data-sources--ip_prefix_set--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ip_prefix_set/data-source.tf`; digest `sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b`.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

## Next pages

- [Examples](data-sources--ip_prefix_set--examples.md)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md)
