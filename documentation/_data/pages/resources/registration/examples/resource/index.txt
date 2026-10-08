---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_registration."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1109, "body_sha256": "sha256:2b835e300c3680c8dc38fdd583e87dc66a8776e8134a9aa1c234487c6f87ef2e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc", "source_path": "examples/resources/xcsh_registration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:registration:example:resource", "parent_id": "xcsh-docs:resources:registration:examples", "path": "documentation/resources/registration/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1212203331303232-0301331003322222-3321232100112033-1311030031332120-0100310332321210-1131103213131131-3200111321322030-3131010323001010", "registry_path": "docs/guides/resources--registration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["registrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
