---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": [], "body_bytes": 1218, "body_sha256": "sha256:993549deb74301b4dae8d28a70750b3f02efa96fe331d756df3bd17147d89270", "child_ids": [], "collection_id": "xcsh-docs:resources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328", "source_path": "examples/resources/xcsh_trusted_ca_list/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:trusted_ca_list:example:resource", "parent_id": "xcsh-docs:resources:trusted_ca_list:examples", "path": "documentation/resources/trusted_ca_list/examples/resource/index.md", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/trusted_ca_list/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_trusted_ca_list.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/examples/)
- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/trusted_ca_list/)
