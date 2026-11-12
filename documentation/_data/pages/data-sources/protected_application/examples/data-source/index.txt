---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protected_application."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1141, "body_sha256": "sha256:91dec3a998f2ba8d9c2d49e7a4a16c61338a4c758fd3eef87a0a6aa979c30eec", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b", "source_path": "examples/data-sources/xcsh_protected_application/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protected_application:example:data-source", "parent_id": "xcsh-docs:data-sources:protected_application:examples", "path": "documentation/data-sources/protected_application/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-2302323202133100-1210020003010131-3221021002120333-3303321302132100-2013231130003331-1031000313023100-0032230130121211-1030222020202023", "registry_path": "docs/guides/data-sources--protected_application--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/examples/data-source/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Data source for xcsh_protected_application.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_application/data-source.tf`; digest `sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b`.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```
