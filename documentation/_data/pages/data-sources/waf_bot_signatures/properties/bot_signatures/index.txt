---
page_title: "bot_signatures"
subcategory: ""
description: "Bot Signatures. A list of all supported bot signatures."
xcsh_docs: {"aliases": ["bot signatures"], "body_bytes": 1990, "body_sha256": "sha256:189b2bc495e7f906af7975ee471fcec65fee32267066a7d66f4342f1fe8c48ad", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "path": "documentation/data-sources/waf_bot_signatures/properties/bot_signatures/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1211023111120001-2203323101130003-2323130323200112-2121311121223131-1331320212010122-2123123111211233-1322212111133321-3302133331132013", "registry_path": "docs/guides/data-sources--waf_bot_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_signatures"], "schema_version": 1, "sections": [{"aliases": ["bot class"], "anchor": "schema-bot_signatures--bot_class", "description": "Bot Class. Enumeration for Bot Class. Possible values are `None`, `Malicious`, `Trusted`, `Untrusted`.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot name"], "anchor": "schema-bot_signatures--bot_name", "description": "Bot Name. The Bot name.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["category"], "anchor": "schema-bot_signatures--category", "description": "Category. The Bot category.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "category"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostnames"], "anchor": "schema-bot_signatures--hostnames", "description": "List of hostnames associated with the bot.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "hostnames"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-bot_signatures--id", "description": "ID. The Signature ID.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["last updated"], "anchor": "schema-bot_signatures--last_updated", "description": "Last Update. The Signature last update time.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "last_updated"], "syntax": "attribute", "type": "string"}, {"aliases": ["risk"], "anchor": "schema-bot_signatures--risk", "description": "Risk. The Bot risk.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "risk"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/bot_signatures/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Signatures. A list of all supported bot signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/)
- [xcsh_waf_bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/)
