---
page_title: "primary.rr_set_group.rr_set.afsdb_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set afsdb record values"], "body_bytes": 3851, "body_sha256": "sha256:0ca368cf2f979f3940d468c9229ce359dd5b83b50b3643322e35a6473e5eea88", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record", "path": "documentation/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set afsdb record values hostname"], "anchor": "schema-primary--rr_set_group--rr_set--afsdb_record--values--hostname", "description": "Server name of the AFS cell database server or the DCE name server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set afsdb record values subtype"], "anchor": "schema-primary--rr_set_group--rr_set--afsdb_record--values--subtype", "description": "AFS Volume Location Server or DCE Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server - DCEAuthenticationServer: DCE Authentication Server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:rr_set_group:rr_set:afsdb_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "afsdb_record", "values", "subtype"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
