---
page_title: "errors.error_obj"
subcategory: ""
description: "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a.."
xcsh_docs: {"aliases": ["errors error obj"], "body_bytes": 1707, "body_sha256": "sha256:0d2b81f0a7d800c30e28888649162c95cace987c1b294e31fe30705d3545834c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors", "path": "documentation/data-sources/site_registrations_by_site/properties/errors/error_obj/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0330323012322031-0103031123333020-1101203222312200-1023010120223103-1330230001200002-2210131021103031-0001112323000220-2210223333022111", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["errors", "error_obj"], "schema_version": 1, "sections": [{"aliases": ["type url"], "anchor": "schema-errors--error_obj--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-errors--error_obj--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "error_obj", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/errors/error_obj/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [errors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/errors/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
