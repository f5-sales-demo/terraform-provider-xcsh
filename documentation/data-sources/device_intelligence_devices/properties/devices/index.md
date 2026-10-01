---
page_title: "devices"
subcategory: ""
description: "devices for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 2040, "body_sha256": "sha256:d96b7b72e706a1c0b1affeb0066baeb616a0348b58a0adced6d53734fef212d9", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:properties:devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:reference", "path": "documentation/data-sources/device_intelligence_devices/properties/devices/index.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/properties/devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "devices for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
