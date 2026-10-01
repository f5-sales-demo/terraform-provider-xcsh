---
page_title: "xcsh_voltstack_site"
subcategory: ""
description: "xcsh_voltstack_site for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:61beab7a01dd08a5e2827529f53d6af95cc354d7bfa5de4fb0b09de87eeb7579", "canonical_id": "xcsh-docs:data-sources:voltstack_site:fundamentals", "child_ids": ["xcsh-docs:data-sources:voltstack_site:reference", "xcsh-docs:data-sources:voltstack_site:examples"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:fundamentals", "parent_id": null, "path": "docs/data-sources/voltstack_site.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_voltstack_site for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_voltstack_site

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--voltstack_site--reference.md)
- [Examples](../guides/data-sources--voltstack_site--examples.md)
