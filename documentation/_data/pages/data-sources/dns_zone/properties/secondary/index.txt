---
page_title: "secondary"
subcategory: "DNS"
description: "SecondaryDNSCreateSpecType."
xcsh_docs: {"aliases": ["secondary"], "body_bytes": 4605, "body_sha256": "sha256:5653f76c2d3ba4879fcdaf328b7dd85ccb4cadb6ae65a7bc779b9d23d0631650", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "parent_id": "xcsh-docs:data-sources:dns_zone:reference", "path": "documentation/data-sources/dns_zone/properties/secondary/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["secondary"], "schema_version": 1, "sections": [{"aliases": ["primary servers"], "anchor": "schema-secondary--primary_servers", "description": "Configuration parameter for primary servers", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["secondary", "primary_servers"], "syntax": "attribute", "type": "list"}, {"aliases": ["tsig key algorithm"], "anchor": "schema-secondary--tsig_key_algorithm", "description": "TSIG key value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC_MD5: HMAC_MD5 - HMAC_SHA1: HMAC_SHA1 - HMAC_SHA224: HMAC_SHA224 - HMAC_SHA256: HMAC_SHA256 - HMAC_SHA384: HMAC_SHA384 - HMAC_SHA512: HMAC_SHA512.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["secondary", "tsig_key_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["tsig key name"], "anchor": "schema-secondary--tsig_key_name", "description": "TSIG key name as used in TSIG protocol extension.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["secondary", "tsig_key_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["tsig key value"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:secondary:tsig_key_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["secondary", "tsig_key_value"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/secondary/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecondaryDNSCreateSpecType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
