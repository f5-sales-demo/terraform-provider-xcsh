---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 54152, "body_sha256": "sha256:f16c9c5988a7fb16de91cf61208d1c169ac5e18a0544122bac9daec227c4783e", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:reference", "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification", "xcsh-docs:data-sources:bigip_virtual_server:properties:default_sensitive_data_policy", "xcsh-docs:data-sources:bigip_virtual_server:properties:disable_api_definition", "xcsh-docs:data-sources:bigip_virtual_server:properties:disable_api_discovery", "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery", "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy", "xcsh-docs:data-sources:bigip_virtual_server:properties:service_discovery"], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:reference", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:fundamentals", "path": "docs/guides/data-sources--bigip_virtual_server--reference.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [api_specification](data-sources--bigip_virtual_server--properties--api_specification.md): complete subsection reference.

<a id="schema-bigip_hostname"></a>

### bigip_hostname property

Type: `"string"`. Computed.

Hostname. BIG-IP Hostname.

<a id="schema-bigip_version"></a>

### bigip_version property

Type: `"string"`. Computed.

Version of the BIG-IP which hosts the virtual server.

<a id="schema-bigip_vs_description"></a>

### bigip_vs_description property

Type: `"string"`. Computed.

Description. BIG-IP Virtual Server Description.

- [default_sensitive_data_policy](data-sources--bigip_virtual_server--properties--default_sensitive_data_policy.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

- [disable_api_definition](data-sources--bigip_virtual_server--properties--disable_api_definition.md): complete subsection reference.

- [disable_api_discovery](data-sources--bigip_virtual_server--properties--disable_api_discovery.md): complete subsection reference.

- [enable_api_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BigIPVirtualServer to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the BigIPVirtualServer.

- [sensitive_data_policy](data-sources--bigip_virtual_server--properties--sensitive_data_policy.md): complete subsection reference.

<a id="schema-server_name"></a>

### server_name property

Type: `"string"`. Computed.

Server Name. Virtual Server name.

- [service_discovery](data-sources--bigip_virtual_server--properties--service_discovery.md): complete subsection reference.

<a id="schema-type"></a>

### type property

Type: `"string"`. Computed.

\[Enum: INVALID\_VIRTUAL\_SERVER|BIGIP\_VIRTUAL\_SERVER\] VirtualServerType could be of type classic
BIG-IP or BIG-IP-NEXT. BIG-IP-NEXT will be added later. Specifies the virtual server type Invalid
Virtual Server Type Classic BIG-IP Virtual Server. Possible values are \`INVALID\_VIRTUAL\_SERVER\`,
\`BIGIP\_VIRTUAL\_SERVER\`. Defaults to \`INVALID\_VIRTUAL\_SERVER\`.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bigip_virtual_server--reference.md#schema-annotations) |
| `api_specification` | [api_specification](data-sources--bigip_virtual_server--properties--api_specification.md#section) |
| `api_specification.api_definition` | [api_specification.api_definition](data-sources--bigip_virtual_server--properties--api_specification--api_definition.md#section) |
| `api_specification.api_definition.name` | [api_specification.api_definition.name](data-sources--bigip_virtual_server--properties--api_specification--api_definition.md#schema-api_specification--api_definition--name) |
| `api_specification.api_definition.namespace` | [api_specification.api_definition.namespace](data-sources--bigip_virtual_server--properties--api_specification--api_definition.md#schema-api_specification--api_definition--namespace) |
| `api_specification.api_definition.tenant` | [api_specification.api_definition.tenant](data-sources--bigip_virtual_server--properties--api_specification--api_definition.md#schema-api_specification--api_definition--tenant) |
| `api_specification.validation_all_spec_endpoints` | [api_specification.validation_all_spec_endpoints](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode` | [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_allow.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_block.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_report.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_skip.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_group) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--base_path) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#section) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--description_spec) |
| `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` | [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--name) |
| `api_specification.validation_all_spec_endpoints.settings` | [api_specification.validation_all_spec_endpoints.settings](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation` | [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_fail_validation.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation` | [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_skip_validation.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom` | [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters` | [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` | [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` | [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md#section) |
| `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default` | [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_default.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode` | [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active` | [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block` | [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active--enforcement_block.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report` | [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active--enforcement_report.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties` | [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active.md#schema-api_specification--validation_all_spec_endpoints--validation_mode--response_validation_mode_active--response_validation_properties) |
| `api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation` | [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_response_validation.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.skip_validation` | [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--skip_validation.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active` | [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block` | [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report` | [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_report.md#section) |
| `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties` | [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md#schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties) |
| `api_specification.validation_custom_list` | [api_specification.validation_custom_list](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list.md#section) |
| `api_specification.validation_custom_list.fall_through_mode` | [api_specification.validation_custom_list.fall_through_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_allow.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_block.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_report.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--action_skip.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_group) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--base_path) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#section) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--description_spec) |
| `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` | [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata.md#schema-api_specification--validation_custom_list--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--metadata--name) |
| `api_specification.validation_custom_list.open_api_validation_rules` | [api_specification.validation_custom_list.open_api_validation_rules](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.any_domain` | [api_specification.validation_custom_list.open_api_validation_rules.any_domain](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--any_domain.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint` | [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--api_endpoint.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.methods` | [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.methods](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--methods) |
| `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.path` | [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.path](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--api_endpoint.md#schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path) |
| `api_specification.validation_custom_list.open_api_validation_rules.api_group` | [api_specification.validation_custom_list.open_api_validation_rules.api_group](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules.md#schema-api_specification--validation_custom_list--open_api_validation_rules--api_group) |
| `api_specification.validation_custom_list.open_api_validation_rules.base_path` | [api_specification.validation_custom_list.open_api_validation_rules.base_path](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules.md#schema-api_specification--validation_custom_list--open_api_validation_rules--base_path) |
| `api_specification.validation_custom_list.open_api_validation_rules.metadata` | [api_specification.validation_custom_list.open_api_validation_rules.metadata](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--metadata.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.metadata.description_spec` | [api_specification.validation_custom_list.open_api_validation_rules.metadata.description_spec](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--metadata.md#schema-api_specification--validation_custom_list--open_api_validation_rules--metadata--description_spec) |
| `api_specification.validation_custom_list.open_api_validation_rules.metadata.name` | [api_specification.validation_custom_list.open_api_validation_rules.metadata.name](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--metadata.md#schema-api_specification--validation_custom_list--open_api_validation_rules--metadata--name) |
| `api_specification.validation_custom_list.open_api_validation_rules.specific_domain` | [api_specification.validation_custom_list.open_api_validation_rules.specific_domain](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules.md#schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_block.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_report.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.response_validation_properties` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.response_validation_properties](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md#schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_response_validation.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--skip_validation.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--enforcement_block.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--enforcement_report.md#section) |
| `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.request_validation_properties` | [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.request_validation_properties](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active.md#schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--request_validation_properties) |
| `api_specification.validation_custom_list.settings` | [api_specification.validation_custom_list.settings](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings.md#section) |
| `api_specification.validation_custom_list.settings.oversized_body_fail_validation` | [api_specification.validation_custom_list.settings.oversized_body_fail_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--oversized_body_fail_validation.md#section) |
| `api_specification.validation_custom_list.settings.oversized_body_skip_validation` | [api_specification.validation_custom_list.settings.oversized_body_skip_validation](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--oversized_body_skip_validation.md#section) |
| `api_specification.validation_custom_list.settings.property_validation_settings_custom` | [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom.md#section) |
| `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters` | [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters.md#section) |
| `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` | [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md#section) |
| `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` | [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md#section) |
| `api_specification.validation_custom_list.settings.property_validation_settings_default` | [api_specification.validation_custom_list.settings.property_validation_settings_default](data-sources--bigip_virtual_server--properties--api_specification--validation_custom_list--settings--property_validation_settings_default.md#section) |
| `api_specification.validation_disabled` | [api_specification.validation_disabled](data-sources--bigip_virtual_server--properties--api_specification--validation_disabled.md#section) |
| `bigip_hostname` | [bigip_hostname](data-sources--bigip_virtual_server--reference.md#schema-bigip_hostname) |
| `bigip_version` | [bigip_version](data-sources--bigip_virtual_server--reference.md#schema-bigip_version) |
| `bigip_vs_description` | [bigip_vs_description](data-sources--bigip_virtual_server--reference.md#schema-bigip_vs_description) |
| `default_sensitive_data_policy` | [default_sensitive_data_policy](data-sources--bigip_virtual_server--properties--default_sensitive_data_policy.md#section) |
| `description` | [description](data-sources--bigip_virtual_server--reference.md#schema-description) |
| `disable_api_definition` | [disable_api_definition](data-sources--bigip_virtual_server--properties--disable_api_definition.md#section) |
| `disable_api_discovery` | [disable_api_discovery](data-sources--bigip_virtual_server--properties--disable_api_discovery.md#section) |
| `enable_api_discovery` | [enable_api_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery.md#section) |
| `enable_api_discovery.api_crawler` | [enable_api_discovery.api_crawler](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config` | [enable_api_discovery.api_crawler.api_crawler_config](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains` | [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` | [enable_api_discovery.api_crawler.api_crawler_config.domains.domain](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--domain) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--decryption_provider) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--location) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--store_provider) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md#section) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--provider_ref) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--url) |
| `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` | [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md#schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user) |
| `enable_api_discovery.api_crawler.disable_api_crawler` | [enable_api_discovery.api_crawler.disable_api_crawler](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--disable_api_crawler.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan` | [enable_api_discovery.api_discovery_from_code_scan](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--all_repos.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md#schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration--name) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md#schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration--namespace) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration.md#schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--code_base_integration--tenant) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md#section) |
| `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` | [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos.md#schema-enable_api_discovery--api_discovery_from_code_scan--code_base_integrations--selected_repos--api_code_repo) |
| `enable_api_discovery.custom_api_auth_discovery` | [enable_api_discovery.custom_api_auth_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery--custom_api_auth_discovery.md#section) |
| `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` | [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](data-sources--bigip_virtual_server--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md#section) |
| `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` | [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name](data-sources--bigip_virtual_server--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md#schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name) |
| `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` | [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace](data-sources--bigip_virtual_server--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md#schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--namespace) |
| `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` | [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant](data-sources--bigip_virtual_server--properties--enable_api_discovery--custom_api_auth_discovery--api_discovery_ref.md#schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--tenant) |
| `enable_api_discovery.default_api_auth_discovery` | [enable_api_discovery.default_api_auth_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery--default_api_auth_discovery.md#section) |
| `enable_api_discovery.disable_learn_from_redirect_traffic` | [enable_api_discovery.disable_learn_from_redirect_traffic](data-sources--bigip_virtual_server--properties--enable_api_discovery--disable_learn_from_redirect_traffic.md#section) |
| `enable_api_discovery.discovered_api_settings` | [enable_api_discovery.discovered_api_settings](data-sources--bigip_virtual_server--properties--enable_api_discovery--discovered_api_settings.md#section) |
| `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` | [enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis](data-sources--bigip_virtual_server--properties--enable_api_discovery--discovered_api_settings.md#schema-enable_api_discovery--discovered_api_settings--purge_duration_for_inactive_discovered_apis) |
| `enable_api_discovery.enable_learn_from_redirect_traffic` | [enable_api_discovery.enable_learn_from_redirect_traffic](data-sources--bigip_virtual_server--properties--enable_api_discovery--enable_learn_from_redirect_traffic.md#section) |
| `id` | [id](data-sources--bigip_virtual_server--reference.md#schema-id) |
| `labels` | [labels](data-sources--bigip_virtual_server--reference.md#schema-labels) |
| `name` | [name](data-sources--bigip_virtual_server--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--bigip_virtual_server--reference.md#schema-namespace) |
| `sensitive_data_policy` | [sensitive_data_policy](data-sources--bigip_virtual_server--properties--sensitive_data_policy.md#section) |
| `sensitive_data_policy.sensitive_data_policy_ref` | [sensitive_data_policy.sensitive_data_policy_ref](data-sources--bigip_virtual_server--properties--sensitive_data_policy--sensitive_data_policy_ref.md#section) |
| `sensitive_data_policy.sensitive_data_policy_ref.name` | [sensitive_data_policy.sensitive_data_policy_ref.name](data-sources--bigip_virtual_server--properties--sensitive_data_policy--sensitive_data_policy_ref.md#schema-sensitive_data_policy--sensitive_data_policy_ref--name) |
| `sensitive_data_policy.sensitive_data_policy_ref.namespace` | [sensitive_data_policy.sensitive_data_policy_ref.namespace](data-sources--bigip_virtual_server--properties--sensitive_data_policy--sensitive_data_policy_ref.md#schema-sensitive_data_policy--sensitive_data_policy_ref--namespace) |
| `sensitive_data_policy.sensitive_data_policy_ref.tenant` | [sensitive_data_policy.sensitive_data_policy_ref.tenant](data-sources--bigip_virtual_server--properties--sensitive_data_policy--sensitive_data_policy_ref.md#schema-sensitive_data_policy--sensitive_data_policy_ref--tenant) |
| `server_name` | [server_name](data-sources--bigip_virtual_server--reference.md#schema-server_name) |
| `service_discovery` | [service_discovery](data-sources--bigip_virtual_server--properties--service_discovery.md#section) |
| `service_discovery.name` | [service_discovery.name](data-sources--bigip_virtual_server--properties--service_discovery.md#schema-service_discovery--name) |
| `service_discovery.namespace` | [service_discovery.namespace](data-sources--bigip_virtual_server--properties--service_discovery.md#schema-service_discovery--namespace) |
| `service_discovery.tenant` | [service_discovery.tenant](data-sources--bigip_virtual_server--properties--service_discovery.md#schema-service_discovery--tenant) |
| `type` | [type](data-sources--bigip_virtual_server--reference.md#schema-type) |

## Next pages

- [api_specification](data-sources--bigip_virtual_server--properties--api_specification.md)
- [default_sensitive_data_policy](data-sources--bigip_virtual_server--properties--default_sensitive_data_policy.md)
- [disable_api_definition](data-sources--bigip_virtual_server--properties--disable_api_definition.md)
- [disable_api_discovery](data-sources--bigip_virtual_server--properties--disable_api_discovery.md)
- [enable_api_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery.md)
- [sensitive_data_policy](data-sources--bigip_virtual_server--properties--sensitive_data_policy.md)
- [service_discovery](data-sources--bigip_virtual_server--properties--service_discovery.md)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
