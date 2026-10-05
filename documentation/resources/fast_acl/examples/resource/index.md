---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1304, "body_sha256": "sha256:27fa026ecc0248996f1debd9679e984298ae344de71fe765b0b6f3b3e3384535", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5b0dbb55293dd11d1f64bb7c693af3c05672ada5840fb242470c138bb98e2c6d", "source_path": "examples/resources/xcsh_fast_acl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl:example:resource", "parent_id": "xcsh-docs:resources:fast_acl:examples", "path": "documentation/resources/fast_acl/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0212022213133203-3122131312110013-1233313220132120-3022321020311112-1030331020321311-0230013303231233-3122312122312030-2012302323133211", "registry_path": "docs/guides/resources--fast_acl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_fast_acl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fast_acl/resource.tf`; digest `sha256:5b0dbb55293dd11d1f64bb7c693af3c05672ada5840fb242470c138bb98e2c6d`.

```terraform
# FastACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACL configuration
resource "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/examples/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
