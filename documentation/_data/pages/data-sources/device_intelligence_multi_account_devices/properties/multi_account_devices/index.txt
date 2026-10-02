---
page_title: "multi_account_devices"
subcategory: ""
description: "Distribution of devices by account range buckets."
xcsh_docs: {"aliases": ["multi account devices"], "body_bytes": 1282, "body_sha256": "sha256:c0296cdffa9354ae422d56035a885f8e19d6ad163b56aa2d8f388bbe13105860", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "path": "documentation/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3032111311010121-1003310120003002-3320123213232122-1020212011031201-2003202210303021-0220233030001321-1110223203232233-3203200022222311", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["multi_account_devices"], "schema_version": 1, "sections": [{"aliases": ["account range"], "anchor": "schema-multi_account_devices--account_range", "description": "Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "account_range"], "syntax": "attribute", "type": "string"}, {"aliases": ["device count"], "anchor": "schema-multi_account_devices--device_count", "description": "Number of devices in this account range.", "document_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["multi_account_devices", "device_count"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Distribution of devices by account range buckets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
