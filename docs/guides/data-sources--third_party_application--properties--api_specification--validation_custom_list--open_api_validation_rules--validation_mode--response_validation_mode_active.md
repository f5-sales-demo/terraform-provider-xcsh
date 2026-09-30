---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active"
subcategory: ""
description: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active for xcsh_third_party_application."
xcsh_docs: {"aliases": [], "body_bytes": 3581, "body_sha256": "sha256:ea636b53f355f9f5478ff98d836f426cf914e4009037da51d26b57cbcc850f09", "canonical_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report"], "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "docs/guides/data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active.md", "provider_name": "third_party_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active for xcsh_third_party_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md)
- [Property reference](data-sources--third_party_application--reference.md)
- [api_specification](data-sources--third_party_application--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--third_party_application--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

## Direct properties

- [enforcement_block](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_block.md): complete subsection reference.

- [enforcement_report](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_report.md): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties"></a>

### response_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_block.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--enforcement_report.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- [xcsh_third_party_application](../data-sources/third_party_application.md)
