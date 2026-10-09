---
page_title: "multi_account_devices"
subcategory: ""
description: "Distribution of devices by account range buckets."
xcsh_docs: {"aliases": ["multi account devices"], "body_bytes": 955, "body_sha256": "sha256:8161d4b13fba91ec692dcf665280eb72e2de73cc443170d57cdfcb8bf66669dd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3032111311010121-1003310120003002-3320123213232122-1020212011031201-2003202210303021-0220233030001321-1110223203232233-3203200022222311", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["multi_account_devices"], "schema_version": 1, "sections": [{"aliases": ["multi account devices account range"], "anchor": "schema-multi_account_devices--account_range", "description": "Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "account_range"], "syntax": "attribute", "type": "string"}, {"aliases": ["multi account devices device count"], "anchor": "schema-multi_account_devices--device_count", "description": "Number of devices in this account range.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "device_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Distribution of devices by account range buckets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# multi_account_devices

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- multi_account_devices

<a id="section"></a>

Type: `"list"`. Computed.

Distribution of devices by account range buckets.

## Direct properties

<a id="schema-multi_account_devices--account_range"></a>

### account_range property

Type: `"string"`. Computed.

Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').

<a id="schema-multi_account_devices--device_count"></a>

### device_count property

Type: `"string"`. Computed.

Number of devices in this account range.
