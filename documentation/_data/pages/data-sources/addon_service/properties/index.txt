---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_addon_service."
xcsh_docs: {"aliases": [], "body_bytes": 3168, "body_sha256": "sha256:90439b4963055c801dbd3b1ea3d3bf0195a9618c1aeaa2f456caf3555334802f", "child_ids": [], "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service:reference", "parent_id": "xcsh-docs:data-sources:addon_service:fundamentals", "path": "documentation/data-sources/addon_service/properties/index.md", "provider_name": "addon_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_addon_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
- Property reference

## Direct properties

<a id="schema-activation_type"></a>

### activation_type property

Type: `"string"`. Computed.

How the addon service is activated. Possible values: \`self\` (user can activate directly),
\`partial\` (requires partial SRE management), \`managed\` (requires full manual intervention).

<a id="schema-addon_service_group_display_name"></a>

### addon_service_group_display_name property

Type: `"string"`. Computed.

Display name of the addon service group.

<a id="schema-addon_service_group_name"></a>

### addon_service_group_name property

Type: `"string"`. Computed.

Name of the addon service group this service belongs to.

<a id="schema-display_name"></a>

### display_name property

Type: `"string"`. Computed.

Human-readable display name of the addon service.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the data source.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the addon service (e.g., \`bot\_defense\`, \`client\_side\_defense\`).

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the addon service is defined. Usually \`shared\`.

<a id="schema-tier"></a>

### tier property

Type: `"string"`. Computed.

Subscription tier required for this addon service. Possible values: \`NO\_TIER\`, \`BASIC\`,
\`STANDARD\`, \`ADVANCED\`, \`PREMIUM\`.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `activation_type` | [activation_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-activation_type) |
| `addon_service_group_display_name` | [addon_service_group_display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-addon_service_group_display_name) |
| `addon_service_group_name` | [addon_service_group_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-addon_service_group_name) |
| `display_name` | [display_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-display_name) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-namespace) |
| `tier` | [tier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/#schema-tier) |

## Next pages

- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
