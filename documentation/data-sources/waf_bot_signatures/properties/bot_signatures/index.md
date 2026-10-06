---
page_title: "bot_signatures"
subcategory: ""
description: "Bot Signatures. A list of all supported bot signatures."
xcsh_docs: {"aliases": ["bot signatures"], "body_bytes": 1732, "body_sha256": "sha256:2154e2dd3959db2fdd17acaeb7e7f2793505635d82b04d9f0c64c2166466554e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "path": "documentation/data-sources/waf_bot_signatures/properties/bot_signatures/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1211023111120001-2203323101130003-2323130323200112-2121311121223131-1331320212010122-2123123111211233-1322212111133321-3302133331132013", "registry_path": "docs/guides/data-sources--waf_bot_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_signatures"], "schema_version": 1, "sections": [{"aliases": ["bot signatures bot class"], "anchor": "schema-bot_signatures--bot_class", "description": "Bot Class. Enumeration for Bot Class. Possible values are `None`, `Malicious`, `Trusted`, `Untrusted`.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures bot name"], "anchor": "schema-bot_signatures--bot_name", "description": "Bot Name. The Bot name.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures category"], "anchor": "schema-bot_signatures--category", "description": "Category. The Bot category.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "category"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures hostnames"], "anchor": "schema-bot_signatures--hostnames", "description": "List of hostnames associated with the bot.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "hostnames"], "syntax": "attribute", "type": "list"}, {"aliases": ["bot signatures id"], "anchor": "schema-bot_signatures--id", "description": "ID. The Signature ID.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures last updated"], "anchor": "schema-bot_signatures--last_updated", "description": "Last Update. The Signature last update time.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "last_updated"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures risk"], "anchor": "schema-bot_signatures--risk", "description": "Risk. The Bot risk.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "risk"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/bot_signatures/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Bot Signatures. A list of all supported bot signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_signatures

Breadcrumbs:

- [xcsh_waf_bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/)
- bot_signatures

<a id="section"></a>

Type: `"list"`. Computed.

Bot Signatures. A list of all supported bot signatures.

## Direct properties

<a id="schema-bot_signatures--bot_class"></a>

### bot_class property

Type: `"string"`. Computed.

\[Enum: None|Malicious|Trusted|Untrusted\] Bot Class. Enumeration for Bot Class. Possible values are
\`None\`, \`Malicious\`, \`Trusted\`, \`Untrusted\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("None",
    "Malicious",
    "Trusted",
    "Untrusted"),
}
```

<a id="schema-bot_signatures--bot_name"></a>

### bot_name property

Type: `"string"`. Computed.

Bot Name. The Bot name.

<a id="schema-bot_signatures--category"></a>

### category property

Type: `"string"`. Computed.

Category. The Bot category.

<a id="schema-bot_signatures--hostnames"></a>

### hostnames property

Type: `["list", "string"]`. Computed.

List of hostnames associated with the bot.

<a id="schema-bot_signatures--id"></a>

### id property

Type: `"string"`. Computed.

ID. The Signature ID.

<a id="schema-bot_signatures--last_updated"></a>

### last_updated property

Type: `"string"`. Computed.

Last Update. The Signature last update time.

<a id="schema-bot_signatures--risk"></a>

### risk property

Type: `"string"`. Computed.

Risk. The Bot risk.
