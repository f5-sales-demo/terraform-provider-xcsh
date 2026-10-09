---
page_title: "Import"
subcategory: "Infrastructure"
description: "Import for xcsh_virtual_site."
xcsh_docs: {"aliases": ["virtual site"], "body_bytes": 361, "body_sha256": "sha256:71ace4381974e96d2cfc0cd135fd8b16eef586407aa314b1d5b7a9198cc282a6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_site:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:virtual_site:fundamentals", "path": "documentation/resources/virtual_site/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "virtual_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2203103200003320-0013322321202333-1222011210030202-0100100121022331-0211122210202331-0002010031312231-2303212101221333-0021010102001201", "registry_path": "docs/guides/resources--virtual_site--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_site/lifecycle/import/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Import for xcsh_virtual_site.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_site/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_virtual_site.example system/example
```
