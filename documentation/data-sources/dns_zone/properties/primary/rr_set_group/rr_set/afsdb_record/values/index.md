---
page_title: "primary.rr_set_group.rr_set.afsdb_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set afsdb record values"], "body_bytes": 4427, "body_sha256": "sha256:aa9c1922e5e0e43c792985810475bd2767a0231d02c85f33b130eda5d311cbb0", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "path": "documentation/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set afsdb record values hostname"], "anchor": "schema-primary--rr_set_group--rr_set--afsdb_record--values--hostname", "description": "Server name of the AFS cell database server or the DCE name server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set afsdb record values subtype"], "anchor": "schema-primary--rr_set_group--rr_set--afsdb_record--values--subtype", "description": "AFS Volume Location Server or DCE Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server - DCEAuthenticationServer: DCE Authentication Server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values", "subtype"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.afsdb_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="section"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--afsdb_record--values--hostname"></a>

### hostname property

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

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
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--afsdb_record--values--subtype"></a>

### subtype property

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [primary.rr_set_group.rr_set.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
