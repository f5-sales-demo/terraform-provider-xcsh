---
page_title: "primary.soa_parameters"
subcategory: "DNS"
description: "Configuration parameter for soa parameters."
xcsh_docs: {"aliases": ["primary soa parameters"], "body_bytes": 5693, "body_sha256": "sha256:c4d60318468e8a8d0f2a9cdff63b39cc1aacd8d134a3717c7f5698dc7524593f", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "path": "documentation/data-sources/dns_zone/properties/primary/soa_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3111301200310120-2012211111103233-2001110322021102-1132112033133212-1100110333112313-0210200023120120-3233033130322010-1032311322222221", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "soa_parameters"], "schema_version": 1, "sections": [{"aliases": ["primary soa parameters expire"], "anchor": "schema-primary--soa_parameters--expire", "description": "Expire value indicates when secondary nameservers should stop answering request for this zone if primary does not respond.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "expire"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary soa parameters negative ttl"], "anchor": "schema-primary--soa_parameters--negative_ttl", "description": "Negative TTL value indicates how long to cache non-existent resource record for this zone.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "negative_ttl"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary soa parameters refresh"], "anchor": "schema-primary--soa_parameters--refresh", "description": "Refresh value indicates when secondary nameservers should query for the SOA record to detect zone changes.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "refresh"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary soa parameters retry"], "anchor": "schema-primary--soa_parameters--retry", "description": "Retry value indicates when secondary nameservers should retry to request the serial number if primary does not respond.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "retry"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary soa parameters ttl"], "anchor": "schema-primary--soa_parameters--ttl", "description": "SOA record time to live (in seconds)", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:soa_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/soa_parameters/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Configuration parameter for soa parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.soa_parameters

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- primary.soa_parameters

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for soa parameters.

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

<a id="schema-primary--soa_parameters--expire"></a>

### expire property

Type: `"number"`. Computed.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--negative_ttl"></a>

### negative_ttl property

Type: `"number"`. Computed.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--refresh"></a>

### refresh property

Type: `"number"`. Computed.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--retry"></a>

### retry property

Type: `"number"`. Computed.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--ttl"></a>

### ttl property

Type: `"number"`. Computed.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

## Next pages

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
