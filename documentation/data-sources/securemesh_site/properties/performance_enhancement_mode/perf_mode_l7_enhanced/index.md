---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "L7 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "body_bytes": 2241, "body_sha256": "sha256:ce87ba248e7ba42ed9f377120d910d14c8216369f8bbe46e0a3be34d95599b6f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode", "path": "documentation/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "sections": [{"aliases": ["jumbo disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["jumbo enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "L7 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

## Direct properties

- [jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/): complete subsection reference.

- [jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/performance_enhancement_mode/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
