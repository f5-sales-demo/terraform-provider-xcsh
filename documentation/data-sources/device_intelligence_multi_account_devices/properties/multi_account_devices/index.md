---
page_title: "multi_account_devices"
subcategory: ""
description: "Distribution of devices by account range buckets."
xcsh_docs: {"aliases": ["multi account devices"], "body_bytes": 955, "body_sha256": "sha256:8161d4b13fba91ec692dcf665280eb72e2de73cc443170d57cdfcb8bf66669dd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3032111311010121-1003310120003002-3320123213232122-1020212011031201-2003202210303021-0220233030001321-1110223203232233-3203200022222311", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["multi_account_devices"], "schema_version": 1, "sections": [{"aliases": ["multi account devices account range"], "anchor": "schema-multi_account_devices--account_range", "description": "Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "account_range"], "syntax": "attribute", "type": "string"}, {"aliases": ["multi account devices device count"], "anchor": "schema-multi_account_devices--device_count", "description": "Number of devices in this account range.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "device_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Distribution of devices by account range buckets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
