---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1278, "body_sha256": "sha256:bdb429fdcf94c254f7f1670b806f557b5a0b85802c3034405f7cddf043945326", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:dd55f175aeb63b8fcdedf3ae50d54522d7f389806ca40efe077ad7621a9cf6e2", "source_path": "examples/resources/xcsh_securemesh_site_v2/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:securemesh_site_v2:example:resource", "parent_id": "xcsh-docs:resources:securemesh_site_v2:examples", "path": "documentation/resources/securemesh_site_v2/examples/resource/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site_v2/resource.tf`; digest `sha256:dd55f175aeb63b8fcdedf3ae50d54522d7f389806ca40efe077ad7621a9cf6e2`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/examples/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
