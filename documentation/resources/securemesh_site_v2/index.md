---
page_title: "xcsh_securemesh_site_v2"
subcategory: ""
description: "xcsh_securemesh_site_v2 for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1703, "body_sha256": "sha256:bf42d0cba0d68669e9c6352db90046967ffce05d99834e8fa92c7d6b64d5bcc5", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:reference", "xcsh-docs:resources:securemesh_site_v2:examples", "xcsh-docs:resources:securemesh_site_v2:import", "xcsh-docs:resources:securemesh_site_v2:timeouts"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:fundamentals", "parent_id": null, "path": "documentation/resources/securemesh_site_v2/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_securemesh_site_v2 for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_securemesh_site_v2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Resource Example
# Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with security and networking controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSiteV2 configuration
resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/lifecycle/timeouts/)
