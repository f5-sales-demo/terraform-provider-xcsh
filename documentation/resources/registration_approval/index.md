---
page_title: "xcsh_registration_approval"
subcategory: ""
description: "Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval. configuration."
xcsh_docs: {"aliases": ["registration approval"], "body_bytes": 1617, "body_sha256": "sha256:9544141a6c4c1c54f1e091439197adcc44a075ceb222f22f74e1af7fc1407fbd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration_approval:reference", "xcsh-docs:resources:registration_approval:examples", "xcsh-docs:resources:registration_approval:import"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration_approval:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration_approval:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/registration_approval/index.md", "product": "distributed-cloud", "provider_name": "registration_approval", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1331103323121310-0002130210000313-2020131122020200-0233121113323131-3323011213331230-0033000123022130-3131333113210323-3013000301112322", "registry_path": "docs/resources/registration_approval.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration_approval/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_registration_approval

Breadcrumbs:

- xcsh_registration_approval

Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `cluster_size`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration_approval/lifecycle/import/)
