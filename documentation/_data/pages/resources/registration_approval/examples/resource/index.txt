---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_registration_approval."
xcsh_docs: {"aliases": [], "body_bytes": 1297, "body_sha256": "sha256:252074e4cee625795dcb8d45bfdfaed6859fc3d1d8552a849db5a34ffd6bde28", "child_ids": [], "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5a69b599688eff851094350bc6a6e7fffd81e09b55087a523abb0c8c970908cd", "source_path": "examples/resources/xcsh_registration_approval/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:registration_approval:example:resource", "parent_id": "xcsh-docs:resources:registration_approval:examples", "path": "documentation/resources/registration_approval/examples/resource/index.md", "provider_name": "registration_approval", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_registration_approval.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_registration_approval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration_approval/resource.tf`; digest `sha256:5a69b599688eff851094350bc6a6e7fffd81e09b55087a523abb0c8c970908cd`.

```terraform
# RegistrationApproval Resource Example
# Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RegistrationApproval configuration
resource "xcsh_registration_approval" "example" {
  name      = "example-registration-approval"
  namespace = "staging"

  cluster_size = 1
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/examples/)
- [xcsh_registration_approval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/)
