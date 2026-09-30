---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1008, "body_sha256": "sha256:67a51a71c1ce9b620dd962866e82a73edb2a986236de4fd6dea73ee5ee7064f6", "canonical_id": "xcsh-docs:resources:service_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e0fbf1c5446df6211fe69296e8020455f7d5f52c63778a8e6218ac43e92a6906", "source_path": "examples/resources/xcsh_service_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:resource", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "docs/guides/resources--service_policy--example--resource.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Examples](resources--service_policy--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/resource.tf`; digest `sha256:e0fbf1c5446df6211fe69296e8020455f7d5f52c63778a8e6218ac43e92a6906`.

```terraform
# ServicePolicy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicy configuration
resource "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--service_policy--examples.md)
- [xcsh_service_policy](../resources/service_policy.md)
