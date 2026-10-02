---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protected_domain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1307, "body_sha256": "sha256:80649c27a0903235b797c3f0743e9bd4c80c16b227fc614f5f7c55064784e475", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92", "source_path": "examples/resources/xcsh_protected_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protected_domain:example:resource", "parent_id": "xcsh-docs:resources:protected_domain:examples", "path": "documentation/resources/protected_domain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2233220302022313-3333220310330113-3201111301233032-0312132232211230-2012203102002121-2222212020020301-3220020221200310-0133203022130133", "registry_path": "docs/guides/resources--protected_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_protected_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_domainCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_protected_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_domain/resource.tf`; digest `sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92`.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/examples/)
- [xcsh_protected_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/)
