---
page_title: "risk_score_distribution"
subcategory: ""
description: "Distribution of devices across risk ranks."
xcsh_docs: {"aliases": ["risk score distribution"], "body_bytes": 1251, "body_sha256": "sha256:cc52e15e16d65082dcd9808c308edcc16da3ab887e6a1c15a3a72204b62d83d5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "parent_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:reference", "path": "documentation/data-sources/device_intelligence_risk_score_distribution/properties/risk_score_distribution/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_risk_score_distribution", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3013013212003322-2122223133210121-1302130312133021-0330132120321103-1033110031033023-1221330002310030-2111012231022010-0313313031320032", "registry_path": "docs/guides/data-sources--device_intelligence_risk_score_distribution--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["risk_score_distribution"], "schema_version": 1, "sections": [{"aliases": ["device count"], "anchor": "schema-risk_score_distribution--device_count", "description": "Device Count. Number of devices in this risk rank.", "document_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["risk_score_distribution", "device_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["risk rank"], "anchor": "schema-risk_score_distribution--risk_rank", "description": "Risk rank label (e.g., low, medium, high).", "document_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:properties:risk_score_distribution", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["risk_score_distribution", "risk_rank"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_risk_score_distribution/properties/risk_score_distribution/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Distribution of devices across risk ranks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/properties/)
- [xcsh_device_intelligence_risk_score_distribution](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/)
