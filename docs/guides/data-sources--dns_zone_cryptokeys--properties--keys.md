---
page_title: "keys"
subcategory: ""
description: "keys for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": [], "body_bytes": 1455, "body_sha256": "sha256:43ed6abe6cbe16e22fd5d6b6403e537f470d7eca4c102b4d04f48a8fe44656ab", "canonical_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:properties:keys", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "path": "docs/guides/data-sources--dns_zone_cryptokeys--properties--keys.md", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["keys"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/properties/keys/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "keys for xcsh_dns_zone_cryptokeys.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# keys

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md)
- [Property reference](data-sources--dns_zone_cryptokeys--reference.md)
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

- [Property reference](data-sources--dns_zone_cryptokeys--reference.md)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md)
