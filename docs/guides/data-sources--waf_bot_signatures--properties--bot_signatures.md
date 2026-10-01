---
page_title: "bot_signatures"
subcategory: ""
description: "bot_signatures for xcsh_waf_bot_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 1782, "body_sha256": "sha256:b04b0e994beea22f0b939df3a962ea026e3174302efaa56a92acc1e01381d75f", "canonical_id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:properties:bot_signatures", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:reference", "path": "docs/guides/data-sources--waf_bot_signatures--properties--bot_signatures.md", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/properties/bot_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_signatures for xcsh_waf_bot_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_signatures

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md)
- [Property reference](data-sources--waf_bot_signatures--reference.md)
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

- [Property reference](data-sources--waf_bot_signatures--reference.md)
- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md)
