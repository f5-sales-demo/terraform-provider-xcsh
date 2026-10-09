---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_swagger_object."
xcsh_docs: {"aliases": ["swagger object"], "body_bytes": 2410, "body_sha256": "sha256:a833f974106de20f6dcbada310419b3226df32866e6f180475134574190f864e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:swagger_object:collection", "completeness": "complete", "id": "xcsh-docs:resources:swagger_object:reference", "parent_id": "xcsh-docs:resources:swagger_object:fundamentals", "path": "documentation/resources/swagger_object/properties/index.md", "product": "distributed-cloud", "provider_name": "swagger_object", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1201120211111103-2221030103330020-2302333103300223-1010321323231300-1333012002311332-0113101012000220-0113100312031013-1131032312102331", "registry_path": "docs/guides/resources--swagger_object--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["content"], "anchor": "schema-content", "description": "Exact UTF-8 OpenAPI 3.0/3.1 or Swagger 2.0 JSON bytes. Use file(). Changed bytes require a new object name and a reviewed replacement.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["content"], "syntax": "attribute", "type": "string"}, {"aliases": ["adopt existing object", "id", "import existing resource"], "anchor": "schema-id", "description": "Import identity: namespace/name/version.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Object DNS label. Include a content digest so changed content has a distinct immutable path.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Owning namespace DNS label.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["path"], "anchor": "schema-path", "description": "Immutable object-store path for API definition swagger_specs.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["path"], "syntax": "attribute", "type": "string"}, {"aliases": ["sha256"], "anchor": "schema-sha256", "description": "SHA-256 of the exact verified content bytes.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-version", "description": "Exact server-issued immutable version.", "document_id": "xcsh-docs:resources:swagger_object:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/swagger_object/properties/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Property reference for xcsh_swagger_object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_swagger_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/)
- Property reference

## Direct properties

<a id="schema-content"></a>

### content property

Type: `"string"`. Required.

Exact UTF-8 OpenAPI 3.0/3.1 or Swagger 2.0 JSON bytes. Use file(). Changed bytes require a new
object name and a reviewed replacement.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Import identity: namespace/name/version.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Object DNS label. Include a content digest so changed content has a distinct immutable path.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Owning namespace DNS label.

<a id="schema-path"></a>

### path property

Type: `"string"`. Computed.

Immutable object-store path for API definition swagger\_specs.

<a id="schema-sha256"></a>

### sha256 property

Type: `"string"`. Computed.

SHA-256 of the exact verified content bytes.

<a id="schema-version"></a>

### version property

Type: `"string"`. Computed.

Exact server-issued immutable version.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `content` | [content](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-content) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-namespace) |
| `path` | [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-path) |
| `sha256` | [sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-sha256) |
| `version` | [version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/swagger_object/properties/#schema-version) |
