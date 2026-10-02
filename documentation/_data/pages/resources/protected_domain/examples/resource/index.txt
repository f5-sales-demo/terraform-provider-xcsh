---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protected_domain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1307, "body_sha256": "sha256:80649c27a0903235b797c3f0743e9bd4c80c16b227fc614f5f7c55064784e475", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92", "source_path": "examples/resources/xcsh_protected_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protected_domain:example:resource", "parent_id": "xcsh-docs:resources:protected_domain:examples", "path": "documentation/resources/protected_domain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233220302022313-3333220310330113-3201111301233032-0312132232211230-2012203102002121-2222212020020301-3220020221200310-0133203022130133", "registry_path": "docs/guides/resources--protected_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protected_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
