---
page_title: "xcsh_sensitive_data_policy"
subcategory: "Security"
description: "Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["sensitive data policy"], "body_bytes": 1751, "body_sha256": "sha256:12e3ada6dc89c237711bebbf6fa1b9e73a21bc29e595b7c9040cc7e43bf7df32", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:sensitive_data_policy:reference", "xcsh-docs:resources:sensitive_data_policy:examples", "xcsh-docs:resources:sensitive_data_policy:import", "xcsh-docs:resources:sensitive_data_policy:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:sensitive_data_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/sensitive_data_policy/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010", "registry_path": "docs/resources/sensitive_data_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_sensitive_data_policy

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SensitiveDataPolicy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SensitiveDataPolicy configuration
resource "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/lifecycle/timeouts/)
