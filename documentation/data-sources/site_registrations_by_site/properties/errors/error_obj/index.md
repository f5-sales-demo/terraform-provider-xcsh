---
page_title: "errors.error_obj"
subcategory: ""
description: "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a.."
xcsh_docs: {"aliases": ["errors error obj"], "body_bytes": 1430, "body_sha256": "sha256:2ba812fc9a2f5d041fc962a6b3cf86c9379fb86aef62b19d28904c057697d4b0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors", "path": "documentation/data-sources/site_registrations_by_site/properties/errors/error_obj/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0330323012322031-0103031123333020-1101203222312200-1023010120223103-1330230001200002-2210131021103031-0001112323000220-2210223333022111", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["errors", "error_obj"], "schema_version": 1, "sections": [{"aliases": ["errors error obj type url"], "anchor": "schema-errors--error_obj--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["errors error obj value"], "anchor": "schema-errors--error_obj--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/errors/error_obj/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# errors.error_obj

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [errors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/errors/)
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
