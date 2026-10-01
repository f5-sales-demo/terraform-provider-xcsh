---
page_title: "xcsh_voltstack_site"
subcategory: ""
description: "xcsh_voltstack_site for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1490, "body_sha256": "sha256:97234b3c803856c33d7777d0bd307fc87b61c8216740905e1eb5b993d33fcac4", "canonical_id": "xcsh-docs:resources:voltstack_site:fundamentals", "child_ids": ["xcsh-docs:resources:voltstack_site:reference", "xcsh-docs:resources:voltstack_site:examples", "xcsh-docs:resources:voltstack_site:import", "xcsh-docs:resources:voltstack_site:timeouts"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:fundamentals", "parent_id": null, "path": "docs/resources/voltstack_site.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_voltstack_site for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--voltstack_site--reference.md)
- [Examples](../guides/resources--voltstack_site--examples.md)
- [Import](../guides/resources--voltstack_site--import.md)
- [Timeouts](../guides/resources--voltstack_site--timeouts.md)
