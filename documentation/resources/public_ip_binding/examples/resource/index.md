---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_public_ip_binding."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1198, "body_sha256": "sha256:734a0705142d048eaa32d6d18253e93589cff698cc6c20bab93b76b50053e6d8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:public_ip_binding:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:193afa89349b62b6fa18ffe2657f80b77590d2dd479611710090adfb88effdd9", "source_path": "examples/resources/xcsh_public_ip_binding/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:public_ip_binding:example:resource", "parent_id": "xcsh-docs:resources:public_ip_binding:examples", "path": "documentation/resources/public_ip_binding/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "public_ip_binding", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2223123313212300-3023032303202012-2012211231031331-1311020231001102-3213120012231210-1013001311303121-0003302020113300-2213222303130231", "registry_path": "docs/guides/resources--public_ip_binding--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/public_ip_binding/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_public_ip_binding.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_public_ip_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_public_ip_binding/resource.tf`; digest `sha256:193afa89349b62b6fa18ffe2657f80b77590d2dd479611710090adfb88effdd9`.

```terraform
# PublicIPBinding Resource Example
# Manage the regional virtual-site binding of an already allocated public IP.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PublicIPBinding configuration
resource "xcsh_public_ip_binding" "example" {
  name      = "example-public-ip-binding"
  namespace = "staging"

  expected_ip            = "example-value"
  virtual_site           = "example-value"
  virtual_site_namespace = "example-value"
}
```
