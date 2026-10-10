---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1094, "body_sha256": "sha256:d9995b7385b3ce7eba45a308ca0dd5a4130dd1998ae41034b84f726800470e67", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5b0dbb55293dd11d1f64bb7c693af3c05672ada5840fb242470c138bb98e2c6d", "source_path": "examples/resources/xcsh_fast_acl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl:example:resource", "parent_id": "xcsh-docs:resources:fast_acl:examples", "path": "documentation/resources/fast_acl/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0212022213133203-3122131312110013-1233313220132120-3022321020311112-1030331020321311-0230013303231233-3122312122312030-2012302323133211", "registry_path": "docs/guides/resources--fast_acl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_fast_acl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_aclCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
