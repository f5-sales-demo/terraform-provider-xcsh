---
page_title: "secondary"
subcategory: "DNS"
description: "secondary for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 4605, "body_sha256": "sha256:5653f76c2d3ba4879fcdaf328b7dd85ccb4cadb6ae65a7bc779b9d23d0631650", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "parent_id": "xcsh-docs:data-sources:dns_zone:reference", "path": "documentation/data-sources/dns_zone/properties/secondary/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["secondary"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/secondary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "secondary for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# secondary

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- secondary

<a id="section"></a>

Type: `"single"`. Computed.

SecondaryDNSCreateSpecType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-secondary--primary_servers"></a>

### primary_servers property

Type: `["list", "string"]`. Computed.

Configuration parameter for primary servers.

Upstream description:

Configuration parameter for primary servers

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-secondary--tsig_key_algorithm"></a>

### tsig_key_algorithm property

Type: `"string"`. Computed.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Upstream description:

TSIG key value must be compatible with the specified algorithm

&#8203;- UNDEFINED: UNDEFINED

&#8203;- HMAC\_MD5: HMAC\_MD5

&#8203;- HMAC\_SHA1: HMAC\_SHA1

&#8203;- HMAC\_SHA224: HMAC\_SHA224

&#8203;- HMAC\_SHA256: HMAC\_SHA256

&#8203;- HMAC\_SHA384: HMAC\_SHA384

&#8203;- HMAC\_SHA512: HMAC\_SHA512.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-secondary--tsig_key_name"></a>

### tsig_key_name property

Type: `"string"`. Computed.

TSIG key name as used in TSIG protocol extension.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/): complete subsection reference.

## Next pages

- [secondary.tsig_key_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
