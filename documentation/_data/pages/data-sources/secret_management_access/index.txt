---
page_title: "xcsh_secret_management_access"
subcategory: ""
description: "Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["secret management access"], "body_bytes": 1487, "body_sha256": "sha256:645c4f3f2d54f0b6baf35f9405dd61c1565252184a518a941dd2463b83bced1f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:reference", "xcsh-docs:data-sources:secret_management_access:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/secret_management_access/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123", "registry_path": "docs/data-sources/secret_management_access.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_secret_management_access

Breadcrumbs:

- xcsh_secret_management_access

Manages secret\_management\_access creates a new object in storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/examples/)
