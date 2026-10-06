---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_namespace."
xcsh_docs: {"aliases": ["namespace"], "body_bytes": 383, "body_sha256": "sha256:cc77fab036d29a5660214b2f2d6b321192401dc52466f64e1fc7bd7c2bf41115", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:resources:namespace:import", "import_guidance": "This tenant-level resource uses its bare name. The `namespace` argument is omitted.", "parent_id": "xcsh-docs:resources:namespace:fundamentals", "path": "documentation/resources/namespace/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0320000300030320-2021002123203112-0303313311033303-1233033211003313-1002122203320111-2330321121202110-2211110330123011-3002101211000321", "registry_path": "docs/guides/resources--namespace--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Import for xcsh_namespace.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/)
- Import

This tenant-level resource uses its bare name. The `namespace` argument is omitted.

```shell
terraform import xcsh_namespace.this example-namespace
```
