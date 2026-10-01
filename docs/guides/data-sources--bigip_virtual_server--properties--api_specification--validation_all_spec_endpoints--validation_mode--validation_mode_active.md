---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active"
subcategory: ""
description: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 3215, "body_sha256": "sha256:99f5d5ace855b298a92e1db447a1eb13a461abf7d436156bb2261b77307b5f27", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report"], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "docs/guides/data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
- [Property reference](data-sources--bigip_virtual_server--reference.md)
- [api_specification](data-sources--bigip_virtual_server--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

## Direct properties

- [enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md): complete subsection reference.

- [enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_report.md): complete subsection reference.

<a id="schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties"></a>

### request_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_block.md)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--enforcement_report.md)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--bigip_virtual_server--properties--api_specification--validation_all_spec_endpoints--validation_mode.md)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
