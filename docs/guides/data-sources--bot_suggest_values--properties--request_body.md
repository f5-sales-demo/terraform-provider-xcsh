---
page_title: "request_body"
subcategory: ""
description: "request_body for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 1212, "body_sha256": "sha256:aad31ff5e32ace49e750c95a80cc175a76ebd33dd088ff24749efc27ece1fb51", "canonical_id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "path": "docs/guides/data-sources--bot_suggest_values--properties--request_body.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["request_body"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/request_body/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "request_body for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# request_body

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
- [Property reference](data-sources--bot_suggest_values--reference.md)
- request_body

<a id="section"></a>

Type: `"single"`. Optional.

Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of
the serialized message. Protobuf library provides support to pack/unpack Any values in the form of
utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..

## Direct properties

<a id="schema-request_body--type_url"></a>

### type_url property

Type: `"string"`. Optional.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="schema-request_body--value"></a>

### value property

Type: `"string"`. Optional.

Must be a valid serialized protocol buffer of the above specified type.

## Next pages

- [Property reference](data-sources--bot_suggest_values--reference.md)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
