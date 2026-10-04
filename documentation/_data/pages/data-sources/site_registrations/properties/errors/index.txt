---
page_title: "errors"
subcategory: ""
description: "Errors(if any) while listing items from collection."
xcsh_docs: {"aliases": ["errors"], "body_bytes": 2091, "body_sha256": "sha256:62d8348ebeeb2e23cec87f908d11accbe7139292d626e1c8477f3ecbb44917a2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:errors:error_obj"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:errors", "parent_id": "xcsh-docs:data-sources:site_registrations:reference", "path": "documentation/data-sources/site_registrations/properties/errors/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2113113221223112-0232201120213121-2000113033120020-0031101123032012-0033013212200001-0133322222133133-3302113302103213-0212030033002033", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["errors"], "schema_version": 1, "sections": [{"aliases": ["errors code"], "anchor": "schema-errors--code", "description": "Union of all possible error-codes from system - EOK: No error - EPERMS: Permissions error - EBADINPUT: Input is not correct - ENOTFOUND: Not found - EEXISTS: Already exists - EUNKNOWN: Unknown/catchall error - ESERIALIZE: Error in serializing/de-serializing - EINTERNAL: Server error - EPARTIAL.. Possible values are", "document_id": "xcsh-docs:data-sources:site_registrations:properties:errors", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "code"], "syntax": "attribute", "type": "string"}, {"aliases": ["errors error obj"], "anchor": "section", "description": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "document_id": "xcsh-docs:data-sources:site_registrations:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["errors", "error_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["errors message"], "anchor": "schema-errors--message", "description": "Message. A human readable string of the error.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:errors", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "message"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/errors/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Errors(if any) while listing items from collection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# errors

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- errors

<a id="section"></a>

Type: `"list"`. Computed.

Errors(if any) while listing items from collection.

## Direct properties

<a id="schema-errors--code"></a>

### code property

Type: `"string"`. Computed.

\[Enum: EOK|EPERMS|EBADINPUT|ENOTFOUND|EEXISTS|EUNKNOWN|ESERIALIZE|EINTERNAL|EPARTIAL\] Union of all
possible error-codes from system - EOK: No error - EPERMS: Permissions error - EBADINPUT: Input is
not correct - ENOTFOUND: Not found - EEXISTS: Already exists - EUNKNOWN: Unknown/catchall error -
ESERIALIZE: Error in serializing/de-serializing - EINTERNAL: Server error - EPARTIAL.. Possible
values are \`EOK\`, \`EPERMS\`, \`EBADINPUT\`, \`ENOTFOUND\`, \`EEXISTS\`, \`EUNKNOWN\`,
\`ESERIALIZE\`, \`EINTERNAL\`, \`EPARTIAL\`. Defaults to \`EOK\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EOK",
    "EPERMS",
    "EBADINPUT",
    "ENOTFOUND",
    "EEXISTS",
    "EUNKNOWN",
    "ESERIALIZE",
    "EINTERNAL",
    "EPARTIAL"),
}
```

- [error_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/errors/error_obj/): complete subsection reference.

<a id="schema-errors--message"></a>

### message property

Type: `"string"`. Computed.

Message. A human readable string of the error.

## Next pages

- [errors.error_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/errors/error_obj/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
