---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1395, "body_sha256": "sha256:a64885fdcab94ae673321d4b4fe3e4f1a7ff9ab8f6653f781f1e6d64d0e5d0e1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:sensitive_data_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1472df2a4b2a381ee1201907e79b19112e2352a74e50b30a503511f47b7c7e1a", "source_path": "examples/resources/xcsh_sensitive_data_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:sensitive_data_policy:example:resource", "parent_id": "xcsh-docs:resources:sensitive_data_policy:examples", "path": "documentation/resources/sensitive_data_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1113230001200101-2102113302103102-1303310101201210-3220023130013320-3222300323323201-0022031233132320-0203320223332000-2321223133011321", "registry_path": "docs/guides/resources--sensitive_data_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/sensitive_data_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_sensitive_data_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_sensitive_data_policy/resource.tf`; digest `sha256:1472df2a4b2a381ee1201907e79b19112e2352a74e50b30a503511f47b7c7e1a`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/examples/)
- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/)
