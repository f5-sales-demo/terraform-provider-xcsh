---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1304, "body_sha256": "sha256:27fa026ecc0248996f1debd9679e984298ae344de71fe765b0b6f3b3e3384535", "child_ids": [], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5b0dbb55293dd11d1f64bb7c693af3c05672ada5840fb242470c138bb98e2c6d", "source_path": "examples/resources/xcsh_fast_acl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:fast_acl:example:resource", "parent_id": "xcsh-docs:resources:fast_acl:examples", "path": "documentation/resources/fast_acl/examples/resource/index.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
