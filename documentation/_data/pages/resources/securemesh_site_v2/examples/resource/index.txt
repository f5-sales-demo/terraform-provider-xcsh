---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1377, "body_sha256": "sha256:f6a8eb072d3c0ac1c588e7954301019295fe2adbe4f00682391df77943714594", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:dd55f175aeb63b8fcdedf3ae50d54522d7f389806ca40efe077ad7621a9cf6e2", "source_path": "examples/resources/xcsh_securemesh_site_v2/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:securemesh_site_v2:example:resource", "parent_id": "xcsh-docs:resources:securemesh_site_v2:examples", "path": "documentation/resources/securemesh_site_v2/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0023333232013220-2113230021310123-3120301200001322-3032302302200322-0133031223122031-3303032211013301-0202221332221222-3113320232122123", "registry_path": "docs/guides/resources--securemesh_site_v2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_securemesh_site_v2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
