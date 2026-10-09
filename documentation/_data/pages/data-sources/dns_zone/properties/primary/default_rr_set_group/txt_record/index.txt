---
page_title: "primary.default_rr_set_group.txt_record"
subcategory: "DNS"
description: "DNSTXTResourceRecord."
xcsh_docs: {"aliases": ["primary default rr set group txt record"], "body_bytes": 3770, "body_sha256": "sha256:c27dddba05755ea48a288d44e0c91eb88fcef76905a1bedc737d6718c37a82dd", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:txt_record", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group", "path": "documentation/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2023003321332312-0023132123210230-1102330021223031-3112003211020100-3031102022010133-0321133333232212-0313320223232033-0232003031030313", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "txt_record"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group txt record name"], "anchor": "schema-primary--default_rr_set_group--txt_record--name", "description": "TXT Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:txt_record", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "txt_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary default rr set group txt record values"], "anchor": "schema-primary--default_rr_set_group--txt_record--values", "description": "Configuration parameter for values", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:default_rr_set_group:txt_record", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "txt_record", "values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "DNSTXTResourceRecord.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.txt_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/)
- primary.default_rr_set_group.txt_record

<a id="section"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="schema-primary--default_rr_set_group--txt_record--name"></a>

### name property

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="schema-primary--default_rr_set_group--txt_record--values"></a>

### values property

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
