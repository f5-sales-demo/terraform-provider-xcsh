---
page_title: "xcsh_sensitive_data_policy"
subcategory: "Security"
description: "Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["sensitive data policy"], "body_bytes": 1492, "body_sha256": "sha256:3e0ab4540cad012ad5fe9276651a115dbea75410aef0f5dfdd1b5f91fbdc6df1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:sensitive_data_policy:reference", "xcsh-docs:data-sources:sensitive_data_policy:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:sensitive_data_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/sensitive_data_policy/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200", "registry_path": "docs/data-sources/sensitive_data_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
# SensitiveDataPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SensitiveDataPolicy by name
data "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}

output "sensitive_data_policy_id" {
  value = data.xcsh_sensitive_data_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/examples/)
