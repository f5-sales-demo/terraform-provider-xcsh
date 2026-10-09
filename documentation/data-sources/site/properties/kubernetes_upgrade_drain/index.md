---
page_title: "kubernetes_upgrade_drain"
subcategory: "Infrastructure"
description: "Specify how worker nodes within a site will be upgraded."
xcsh_docs: {"aliases": ["kubernetes upgrade drain"], "body_bytes": 883, "body_sha256": "sha256:403b615c8f31b272cdc66fdba83f000a7618a4fd0dcb6e073b73ea545c7835f0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:site:reference", "path": "documentation/data-sources/site/properties/kubernetes_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2002303221332323-3233110303213322-0311210120200132-3112301030213001-0203120331301121-0100202030200110-2012032300320120-0333202231032000", "registry_path": "docs/guides/data-sources--site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["kubernetes upgrade drain disable upgrade drain"], "anchor": "section", "description": "Configuration parameter for disable upgrade drain.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain"], "anchor": "section", "description": "Specify batch upgrade settings for worker nodes within a site.", "document_id": "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Specify how worker nodes within a site will be upgraded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/)
- kubernetes_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

## Direct properties

- [disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/): complete subsection reference.

- [enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/): complete subsection reference.
