---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protected_application."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1346, "body_sha256": "sha256:7f6d3655b4dc8d58f33e94b290f12d859ca0e8b8d1abcaab390def53aa34c7be", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:868be3bcffc6df67dc51247c65c90b154557ee6fbafff96f6968ad86d503ff19", "source_path": "examples/resources/xcsh_protected_application/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protected_application:example:resource", "parent_id": "xcsh-docs:resources:protected_application:examples", "path": "documentation/resources/protected_application/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0120220300031032-0220001023101321-3110212110321113-1200123320302302-0322011311233030-3111120231220131-1211132021033233-3201031222020333", "registry_path": "docs/guides/resources--protected_application--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_protected_application.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_application/resource.tf`; digest `sha256:868be3bcffc6df67dc51247c65c90b154557ee6fbafff96f6968ad86d503ff19`.

```terraform
# ProtectedApplication Resource Example
# Manages applications protected by Bot Defense in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedApplication configuration
resource "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/examples/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
