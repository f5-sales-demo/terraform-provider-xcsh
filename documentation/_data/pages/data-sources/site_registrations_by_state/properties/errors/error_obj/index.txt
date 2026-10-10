---
page_title: "errors.error_obj"
subcategory: ""
description: "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a.."
xcsh_docs: {"aliases": ["errors error obj"], "body_bytes": 1434, "body_sha256": "sha256:a7989529aec094faf8db83aca45c07f5d99a8e1c88f7d7ef5206576acbad5943", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors:error_obj", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors", "path": "documentation/data-sources/site_registrations_by_state/properties/errors/error_obj/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1010211013112301-1131221010330231-3133220301030231-0302212300122302-3200131000020122-0110213101233100-3120011230322302-0313320102333033", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["errors", "error_obj"], "schema_version": 1, "sections": [{"aliases": ["errors error obj type url"], "anchor": "schema-errors--error_obj--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors:error_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["errors error obj value"], "anchor": "schema-errors--error_obj--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors:error_obj", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/errors/error_obj/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# errors.error_obj

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [errors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/errors/)
- errors.error_obj

<a id="section"></a>

Type: `"single"`. Computed.

Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of
the serialized message. Protobuf library provides support to pack/unpack Any values in the form of
utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..

## Direct properties

<a id="schema-errors--error_obj--type_url"></a>

### type_url property

Type: `"string"`. Computed.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="schema-errors--error_obj--value"></a>

### value property

Type: `"string"`. Computed.

Must be a valid serialized protocol buffer of the above specified type.
