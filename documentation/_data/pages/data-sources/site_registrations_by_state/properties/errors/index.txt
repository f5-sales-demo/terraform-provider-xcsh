---
page_title: "errors"
subcategory: ""
description: "Errors(if any) while listing items from collection."
xcsh_docs: {"aliases": ["errors"], "body_bytes": 2163, "body_sha256": "sha256:c3e56804c20c1397d5fa118ee4fd95a38e45aae6f1b6ea827e4f470d2093ec96", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:errors:error_obj"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:reference", "path": "documentation/data-sources/site_registrations_by_state/properties/errors/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3123201032113231-1110203102030102-0321102211032320-2011122333202121-2011102030213113-2300113212013233-3311310101312121-2030132221300302", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["errors"], "schema_version": 1, "sections": [{"aliases": ["code"], "anchor": "schema-errors--code", "description": "Union of all possible error-codes from system - EOK: No error - EPERMS: Permissions error - EBADINPUT: Input is not correct - ENOTFOUND: Not found - EEXISTS: Already exists - EUNKNOWN: Unknown/catchall error - ESERIALIZE: Error in serializing/de-serializing - EINTERNAL: Server error - EPARTIAL.. Possible values are `EO", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "code"], "syntax": "attribute", "type": "string"}, {"aliases": ["error obj"], "anchor": "section", "description": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors:error_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["errors", "error_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["message"], "anchor": "schema-errors--message", "description": "Message. A human readable string of the error.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:errors", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["errors", "message"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/errors/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Errors(if any) while listing items from collection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# errors

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
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

- [error_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/errors/error_obj/): complete subsection reference.

<a id="schema-errors--message"></a>

### message property

Type: `"string"`. Computed.

Message. A human readable string of the error.

## Next pages

- [errors.error_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/errors/error_obj/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
