---
page_title: "xcsh_securemesh_site"
subcategory: ""
description: "xcsh_securemesh_site for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1538, "body_sha256": "sha256:fefbfc95b67d0078f6518e73559bf48348c0403260ef4e39cce1bd2162c3d128", "canonical_id": "xcsh-docs:resources:securemesh_site:fundamentals", "child_ids": ["xcsh-docs:resources:securemesh_site:reference", "xcsh-docs:resources:securemesh_site:examples", "xcsh-docs:resources:securemesh_site:import", "xcsh-docs:resources:securemesh_site:timeouts"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:fundamentals", "parent_id": null, "path": "docs/resources/securemesh_site.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_securemesh_site for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_securemesh_site

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--securemesh_site--reference.md)
- [Examples](../guides/resources--securemesh_site--examples.md)
- [Import](../guides/resources--securemesh_site--import.md)
- [Timeouts](../guides/resources--securemesh_site--timeouts.md)
