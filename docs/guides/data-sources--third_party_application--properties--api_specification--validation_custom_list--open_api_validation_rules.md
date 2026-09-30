---
page_title: "api_specification.validation_custom_list.open_api_validation_rules"
subcategory: ""
description: "api_specification.validation_custom_list.open_api_validation_rules for xcsh_third_party_application."
xcsh_docs: {"aliases": [], "body_bytes": 3256, "body_sha256": "sha256:5df98f8d4b99be4c47bed2564f8b7c9becfee05518b69ea1b8050f20dc9d92ce", "canonical_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode"], "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list", "path": "docs/guides/data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules.md", "provider_name": "third_party_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.open_api_validation_rules for xcsh_third_party_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.open_api_validation_rules

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md)
- [Property reference](data-sources--third_party_application--reference.md)
- [api_specification](data-sources--third_party_application--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--third_party_application--properties--api_specification--validation_custom_list.md)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="section"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

## Direct properties

- [any_domain](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--any_domain.md): complete subsection reference.

- [api_endpoint](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--api_endpoint.md): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_group"></a>

### api_group property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--base_path"></a>

### base_path property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

- [metadata](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--metadata.md): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

- [validation_mode](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--any_domain.md)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--api_endpoint.md)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--metadata.md)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--third_party_application--properties--api_specification--validation_custom_list--open_api_validation_rules--validation_mode.md)
- [api_specification.validation_custom_list](data-sources--third_party_application--properties--api_specification--validation_custom_list.md)
- [xcsh_third_party_application](../data-sources/third_party_application.md)
