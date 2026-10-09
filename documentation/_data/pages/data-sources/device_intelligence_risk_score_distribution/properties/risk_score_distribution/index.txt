---
page_title: "risk_score_distribution"
subcategory: ""
description: "Distribution of devices across risk ranks."
xcsh_docs: {"aliases": ["risk score distribution"], "body_bytes": 918, "body_sha256": "sha256:8af559ea8a8a48af24014ebd9c92eee5323488ef61241df4980bc7999f59b8bb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "parent_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:reference", "path": "documentation/data-sources/device_intelligence_risk_score_distribution/properties/risk_score_distribution/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_risk_score_distribution", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3013013212003322-2122223133210121-1302130312133021-0330132120321103-1033110031033023-1221330002310030-2111012231022010-0313313031320032", "registry_path": "docs/guides/data-sources--device_intelligence_risk_score_distribution--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["risk_score_distribution"], "schema_version": 1, "sections": [{"aliases": ["risk score distribution device count"], "anchor": "schema-risk_score_distribution--device_count", "description": "Device Count. Number of devices in this risk rank.", "document_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["risk_score_distribution", "device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["risk score distribution risk rank"], "anchor": "schema-risk_score_distribution--risk_rank", "description": "Risk rank label (e.g., low, medium, high).", "document_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["risk_score_distribution", "risk_rank"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_risk_score_distribution/properties/risk_score_distribution/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Distribution of devices across risk ranks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# risk_score_distribution

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/properties/)
- risk_score_distribution

<a id="section"></a>

Type: `"list"`. Computed.

Distribution of devices across risk ranks.

## Direct properties

<a id="schema-risk_score_distribution--device_count"></a>

### device_count property

Type: `"string"`. Computed.

Device Count. Number of devices in this risk rank.

<a id="schema-risk_score_distribution--risk_rank"></a>

### risk_rank property

Type: `"string"`. Computed.

Risk rank label (e.g., low, medium, high).
