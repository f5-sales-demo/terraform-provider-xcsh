---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_registration."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1331, "body_sha256": "sha256:32aca4be090ac846a96f5e44a783875944a1c0500fdcbf5996be5da0fb0581ea", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc", "source_path": "examples/resources/xcsh_registration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:registration:example:resource", "parent_id": "xcsh-docs:resources:registration:examples", "path": "documentation/resources/registration/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1212203331303232-0301331003322222-3321232100112033-1311030031332120-0100310332321210-1131103213131131-3200111321322030-3131010323001010", "registry_path": "docs/guides/resources--registration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration/resource.tf`; digest `sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc`.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/examples/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
