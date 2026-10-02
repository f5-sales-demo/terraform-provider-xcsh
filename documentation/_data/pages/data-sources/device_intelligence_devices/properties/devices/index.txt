---
page_title: "devices"
subcategory: ""
description: "Devices. List of devices for this page."
xcsh_docs: {"aliases": ["devices"], "body_bytes": 2040, "body_sha256": "sha256:d96b7b72e706a1c0b1affeb0066baeb616a0348b58a0adced6d53734fef212d9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "path": "documentation/data-sources/device_intelligence_devices/properties/devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1023202112002210-1320033332001301-0031231330321303-0320220203210312-2133220130012221-0331123213330123-0002030021201233-2021212010131213", "registry_path": "docs/guides/data-sources--device_intelligence_devices--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["devices"], "schema_version": 1, "sections": [{"aliases": ["action taken"], "anchor": "schema-devices--action_taken", "description": "Action taken or recommended based on the risk assessment.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "action_taken"], "syntax": "attribute", "type": "string"}, {"aliases": ["confidence"], "anchor": "schema-devices--confidence", "description": "Confidence level of the risk assessment (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "confidence"], "syntax": "attribute", "type": "number"}, {"aliases": ["device id"], "anchor": "schema-devices--device_id", "description": "Device ID. Unique identifier for the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "device_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["high risk txn count"], "anchor": "schema-devices--high_risk_txn_count", "description": "Number of high-risk transactions linked to the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "high_risk_txn_count"], "syntax": "attribute", "type": "string"}, {"aliases": ["latest txn id"], "anchor": "schema-devices--latest_txn_id", "description": "Identifier of the most recent transaction associated with the device.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "latest_txn_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["linked accounts"], "anchor": "schema-devices--linked_accounts", "description": "Number of distinct accounts linked to this device.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "linked_accounts"], "syntax": "attribute", "type": "string"}, {"aliases": ["risk score"], "anchor": "schema-devices--risk_score", "description": "Overall risk score for the device (0–100).", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "risk_score"], "syntax": "attribute", "type": "number"}, {"aliases": ["risk signals"], "anchor": "schema-devices--risk_signals", "description": "List of risk signals detected for this device.", "document_id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "risk_signals"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/devices/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Devices. List of devices for this page.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# devices

Breadcrumbs:

- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/)
- devices

<a id="section"></a>

Type: `"list"`. Computed.

Devices. List of devices for this page.

## Direct properties

<a id="schema-devices--action_taken"></a>

### action_taken property

Type: `"string"`. Computed.

Action taken or recommended based on the risk assessment.

<a id="schema-devices--confidence"></a>

### confidence property

Type: `"number"`. Computed.

Confidence level of the risk assessment (0–100).

<a id="schema-devices--device_id"></a>

### device_id property

Type: `"string"`. Computed.

Device ID. Unique identifier for the device.

<a id="schema-devices--high_risk_txn_count"></a>

### high_risk_txn_count property

Type: `"string"`. Computed.

Number of high-risk transactions linked to the device.

<a id="schema-devices--latest_txn_id"></a>

### latest_txn_id property

Type: `"string"`. Computed.

Identifier of the most recent transaction associated with the device.

<a id="schema-devices--linked_accounts"></a>

### linked_accounts property

Type: `"string"`. Computed.

Number of distinct accounts linked to this device.

<a id="schema-devices--risk_score"></a>

### risk_score property

Type: `"number"`. Computed.

Overall risk score for the device (0–100).

<a id="schema-devices--risk_signals"></a>

### risk_signals property

Type: `["list", "string"]`. Computed.

List of risk signals detected for this device.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/)
- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
