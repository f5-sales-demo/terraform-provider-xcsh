---
page_title: "keys"
subcategory: ""
description: "Configuration parameter for keys"
xcsh_docs: {"aliases": ["keys"], "body_bytes": 1663, "body_sha256": "sha256:675d164736eb447424701304599484a2456549cff9985f4f71a75c5791a6f356", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "path": "documentation/data-sources/dns_zone_cryptokeys/properties/keys/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1222123200323131-0001133313121011-2322323223233231-3311220101230220-1003322320030102-3002333102002223-1102232023100020-1113121303230233", "registry_path": "docs/guides/data-sources--dns_zone_cryptokeys--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["keys"], "schema_version": 1, "sections": [{"aliases": ["keys active"], "anchor": "schema-keys--active", "description": "Whether the key is currently active for signing.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "active"], "syntax": "attribute", "type": "bool"}, {"aliases": ["keys algorithm"], "anchor": "schema-keys--algorithm", "description": "The DNSSEC signing algorithm used by this key.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["keys dnskey"], "anchor": "schema-keys--dnskey", "description": "DNSKEY. The DNSKEY record data for this key.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "dnskey"], "syntax": "attribute", "type": "string"}, {"aliases": ["keys key id"], "anchor": "schema-keys--key_id", "description": "Unique identifier for the cryptographic key.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "key_id"], "syntax": "attribute", "type": "number"}, {"aliases": ["keys key type"], "anchor": "schema-keys--key_type", "description": "The cryptographic key type (e.g., CSK, KSK, ZSK).", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "key_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["keys published"], "anchor": "schema-keys--published", "description": "Whether the key is published in the DNSKEY RRset.", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "published"], "syntax": "attribute", "type": "bool"}, {"aliases": ["keys type"], "anchor": "schema-keys--type", "description": "Type. Should always be 'CryptoKey'", "document_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["keys", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/properties/keys/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Configuration parameter for keys", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# keys

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/)
- keys

<a id="section"></a>

Type: `"list"`. Computed.

Configuration parameter for keys

## Direct properties

<a id="schema-keys--active"></a>

### active property

Type: `"bool"`. Computed.

Whether the key is currently active for signing.

<a id="schema-keys--algorithm"></a>

### algorithm property

Type: `"string"`. Computed.

The DNSSEC signing algorithm used by this key.

<a id="schema-keys--dnskey"></a>

### dnskey property

Type: `"string"`. Computed.

DNSKEY. The DNSKEY record data for this key.

<a id="schema-keys--key_id"></a>

### key_id property

Type: `"number"`. Computed.

Unique identifier for the cryptographic key.

<a id="schema-keys--key_type"></a>

### key_type property

Type: `"string"`. Computed.

The cryptographic key type (e.g., CSK, KSK, ZSK).

<a id="schema-keys--published"></a>

### published property

Type: `"bool"`. Computed.

Whether the key is published in the DNSKEY RRset.

<a id="schema-keys--type"></a>

### type property

Type: `"string"`. Computed.

Type. Should always be 'CryptoKey'

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/)
- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
