---
page_title: "request_body"
subcategory: ""
description: "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a.."
xcsh_docs: {"aliases": ["request body"], "body_bytes": 1261, "body_sha256": "sha256:595b90261074a49f02bd683c200585651c06b313f85c93bab2b79bb786ab3501", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:reference", "path": "documentation/data-sources/bot_suggest_values/properties/request_body/index.md", "product": "distributed-cloud", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2302002121032203-2330322011313122-1000312001232330-3322130231231133-0032323203121310-2333031332013131-2122312101031013-0032330010330123", "registry_path": "docs/guides/data-sources--bot_suggest_values--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["request_body"], "schema_version": 1, "sections": [{"aliases": ["request body type url"], "anchor": "schema-request_body--type_url", "description": "URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one '/' character. The last segment of the URL path must represent the fully qualified name of the type (as in ).", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_body", "type_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["request body value"], "anchor": "schema-request_body--value", "description": "Must be a valid serialized protocol buffer of the above specified type.", "document_id": "xcsh-docs:data-sources:bot_suggest_values:properties:request_body", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_body", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/properties/request_body/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of the serialized message. Protobuf library provides support to pack/unpack Any values in the form of utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_body

Breadcrumbs:

- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/properties/)
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
