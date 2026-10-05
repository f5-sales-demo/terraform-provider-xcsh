---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1462, "body_sha256": "sha256:a521379fdd97b52868b817462a5d019e7fffc9bd64be909d9a0eefa61639024e", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81", "source_path": "examples/resources/xcsh_secret_management_access/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:secret_management_access:example:resource", "parent_id": "xcsh-docs:resources:secret_management_access:examples", "path": "documentation/resources/secret_management_access/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0303230202110201-0203211233311100-3302130311223000-2223101332121131-3312313012232330-1131211023302311-0201003013011023-2331121003023322", "registry_path": "docs/guides/resources--secret_management_access--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_secret_management_access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_management_access/resource.tf`; digest `sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81`.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/examples/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
