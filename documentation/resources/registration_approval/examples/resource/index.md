---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_registration_approval."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1147, "body_sha256": "sha256:8d3872c4988e5177f400c4ab60b470dc87b49bce5fe53a0216e7aba0261156a8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5a69b599688eff851094350bc6a6e7fffd81e09b55087a523abb0c8c970908cd", "source_path": "examples/resources/xcsh_registration_approval/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:registration_approval:example:resource", "parent_id": "xcsh-docs:resources:registration_approval:examples", "path": "documentation/resources/registration_approval/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "registration_approval", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1031110231221333-0203300210212111-0223331033100203-1100320031333333-1222031031033003-3033130012233322-0223223213003023-0132133330322103", "registry_path": "docs/guides/resources--registration_approval--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_registration_approval.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
