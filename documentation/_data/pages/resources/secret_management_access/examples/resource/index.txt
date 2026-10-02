---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1462, "body_sha256": "sha256:a521379fdd97b52868b817462a5d019e7fffc9bd64be909d9a0eefa61639024e", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81", "source_path": "examples/resources/xcsh_secret_management_access/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:secret_management_access:example:resource", "parent_id": "xcsh-docs:resources:secret_management_access:examples", "path": "documentation/resources/secret_management_access/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0303230202110201-0203211233311100-3302130311223000-2223101332121131-3312313012232330-1131211023302311-0201003013011023-2331121003023322", "registry_path": "docs/guides/resources--secret_management_access--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_secret_management_access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
