---
page_title: "multi_account_devices"
subcategory: ""
description: "multi_account_devices for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": [], "body_bytes": 1074, "body_sha256": "sha256:f9edda692b46ef4e51eec826c7d8eba9152a8993411209c4eda26e6adb25d781", "canonical_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:properties:multi_account_devices", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "path": "docs/guides/data-sources--device_intelligence_multi_account_devices--properties--multi_account_devices.md", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["multi_account_devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/properties/multi_account_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "multi_account_devices for xcsh_device_intelligence_multi_account_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# multi_account_devices

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference.md)
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

- [Property reference](data-sources--device_intelligence_multi_account_devices--reference.md)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md)
