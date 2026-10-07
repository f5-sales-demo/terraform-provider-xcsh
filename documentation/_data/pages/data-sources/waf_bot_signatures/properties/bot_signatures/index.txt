---
page_title: "bot_signatures"
subcategory: ""
description: "Bot Signatures. A list of all supported bot signatures."
xcsh_docs: {"aliases": ["bot signatures"], "body_bytes": 2009, "body_sha256": "sha256:1eb5316fb9463cead8a78e8935022145a964fd71c51e5c09a64d084ad1d02fde", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "path": "documentation/data-sources/waf_bot_signatures/properties/bot_signatures/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1211023111120001-2203323101130003-2323130323200112-2121311121223131-1331320212010122-2123123111211233-1322212111133321-3302133331132013", "registry_path": "docs/guides/data-sources--waf_bot_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_signatures"], "schema_version": 1, "sections": [{"aliases": ["bot signatures bot class"], "anchor": "schema-bot_signatures--bot_class", "description": "Bot Class. Enumeration for Bot Class. Possible values are `None`, `Malicious`, `Trusted`, `Untrusted`.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["Malicious", "None", "Trusted", "Untrusted"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures bot name"], "anchor": "schema-bot_signatures--bot_name", "description": "Bot Name. The Bot name.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "bot_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures category"], "anchor": "schema-bot_signatures--category", "description": "Category. The Bot category.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "category"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures hostnames"], "anchor": "schema-bot_signatures--hostnames", "description": "List of hostnames associated with the bot.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "hostnames"], "syntax": "attribute", "type": "list"}, {"aliases": ["bot signatures id"], "anchor": "schema-bot_signatures--id", "description": "ID. The Signature ID.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures last updated"], "anchor": "schema-bot_signatures--last_updated", "description": "Last Update. The Signature last update time.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "last_updated"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot signatures risk"], "anchor": "schema-bot_signatures--risk", "description": "Risk. The Bot risk.", "document_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_signatures", "risk"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/bot_signatures/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Bot Signatures. A list of all supported bot signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Malicious","None","Trusted","Untrusted"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
