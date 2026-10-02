---
page_title: "primary.default_rr_set_group.afsdb_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group afsdb record values"], "body_bytes": 4298, "body_sha256": "sha256:7427e1a4292c3f2673e959633c870a48139465b2da19973e7b99296b5f19c1ad", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record:values", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3323233331320022-1000000300301231-2103221320301002-1003123123131031-3002312031121300-0122330002032202-1000311132321011-2013131113033310", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "afsdb_record", "values"], "schema_version": 1, "sections": [{"aliases": ["hostname"], "anchor": "schema-primary--default_rr_set_group--afsdb_record--values--hostname", "description": "Server name of the AFS cell database server or the DCE name server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "afsdb_record", "values", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "subtype"], "anchor": "schema-primary--default_rr_set_group--afsdb_record--values--subtype", "description": "AFS Volume Location Server or DCE Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server - DCEAuthenticationServer: DCE Authentication Server.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:afsdb_record:values", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "afsdb_record", "values", "subtype"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.afsdb_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/)
- primary.default_rr_set_group.afsdb_record.values

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

<a id="schema-primary--default_rr_set_group--afsdb_record--values--hostname"></a>

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

<a id="schema-primary--default_rr_set_group--afsdb_record--values--subtype"></a>

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

- [primary.default_rr_set_group.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
