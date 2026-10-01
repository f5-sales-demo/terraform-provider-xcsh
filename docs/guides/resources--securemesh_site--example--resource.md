---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1169, "body_sha256": "sha256:806eed2633ea80520c8e1a39960c6c263b920bd6deaaa49e43d9a79c83c903e3", "canonical_id": "xcsh-docs:resources:securemesh_site:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d12d5ac4d45f23102fc4cf022413c9bdb3fd5f59873115079a2e3a8fe4e2c28a", "source_path": "examples/resources/xcsh_securemesh_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:securemesh_site:example:resource", "parent_id": "xcsh-docs:resources:securemesh_site:examples", "path": "docs/guides/resources--securemesh_site--example--resource.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Examples](resources--securemesh_site--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site/resource.tf`; digest `sha256:d12d5ac4d45f23102fc4cf022413c9bdb3fd5f59873115079a2e3a8fe4e2c28a`.

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

## Next pages

- [Examples](resources--securemesh_site--examples.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
