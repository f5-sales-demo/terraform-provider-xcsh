---
page_title: "aaaa_pool.members"
subcategory: ""
description: "Configuration parameter for members"
xcsh_docs: {"aliases": ["aaaa pool members"], "body_bytes": 6063, "body_sha256": "sha256:9b687fe511cddacf0b54d4c44a7a2f8e60791c3fc2ae16cb4aebc5eebb288e94", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool", "path": "documentation/data-sources/dns_lb_pool/properties/aaaa_pool/members/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3330231313131112-2230323222113321-0232330201121330-1021130323222322-3001220022312012-1323112332011303-2003132101303123-0233222033231313", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aaaa_pool", "members"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "schema-aaaa_pool--members--disable_spec", "description": "Value of true will disable the pool-member.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "members", "disable_spec"], "syntax": "attribute", "type": "bool"}, {"aliases": ["ip endpoint"], "anchor": "schema-aaaa_pool--members--ip_endpoint", "description": "Public IP address.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "members", "ip_endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-aaaa_pool--members--name", "description": "Pool member name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "members", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["priority"], "anchor": "schema-aaaa_pool--members--priority", "description": "Used if the pool’s load balancing mode is set to Priority.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "members", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["ratio"], "anchor": "schema-aaaa_pool--members--ratio", "description": "Used if the pool’s load balancing mode is set to Ratio-Member.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:aaaa_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aaaa_pool", "members", "ratio"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/aaaa_pool/members/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for members", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aaaa_pool.members

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/aaaa_pool/)
- aaaa_pool.members

<a id="section"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-aaaa_pool--members--disable_spec"></a>

### disable_spec property

Type: `"bool"`. Computed.

Value of true will disable the pool-member.

<a id="schema-aaaa_pool--members--ip_endpoint"></a>

### ip_endpoint property

Type: `"string"`. Computed.

Public IP. Public IP address.

Upstream description:

Public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-aaaa_pool--members--name"></a>

### name property

Type: `"string"`. Computed.

Name. Pool member name.

Upstream description:

Pool member name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-aaaa_pool--members--priority"></a>

### priority property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="schema-aaaa_pool--members--ratio"></a>

### ratio property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

## Next pages

- [aaaa_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/aaaa_pool/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
