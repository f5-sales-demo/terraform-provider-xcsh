---
page_title: "cname_pool.members"
subcategory: ""
description: "Configuration parameter for members"
xcsh_docs: {"aliases": ["cname pool members"], "body_bytes": 6450, "body_sha256": "sha256:8a950382776df92b329908d64639fea2e88d7ca99af967f2523b930e4f9f83f1", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool", "path": "documentation/data-sources/dns_lb_pool/properties/cname_pool/members/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0000213312220302-1011113232232320-0133221000321100-0121330130211212-0322202131221000-2013230022132131-1000102103303232-3021200300131033", "registry_path": "docs/guides/data-sources--dns_lb_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool", "members"], "schema_version": 1, "sections": [{"aliases": ["domain"], "anchor": "schema-cname_pool--members--domain", "description": "Specifies the fully qualified domain name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["final translation"], "anchor": "schema-cname_pool--members--final_translation", "description": "If this flag is true, the CNAME record will not be translated further.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "final_translation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-cname_pool--members--name", "description": "Pool member name.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["priority"], "anchor": "schema-cname_pool--members--priority", "description": "Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is routed to pool members. The lower the number, the higher the priority, making those members active while higher-numbered members act as backups.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["ratio"], "anchor": "schema-cname_pool--members--ratio", "description": "Used if the pool’s load balancing mode is set to Ratio-Member.", "document_id": "xcsh-docs:data-sources:dns_lb_pool:properties:cname_pool:members", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "ratio"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/cname_pool/members/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for members", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool.members

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/)
- cname_pool.members

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

<a id="schema-cname_pool--members--domain"></a>

### domain property

Type: `"string"`. Computed.

Specifies the fully qualified domain name.

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

<a id="schema-cname_pool--members--final_translation"></a>

### final_translation property

Type: `"bool"`. Computed.

If this flag is true, the CNAME record will not be translated further.

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

<a id="schema-cname_pool--members--name"></a>

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

<a id="schema-cname_pool--members--priority"></a>

### priority property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

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

<a id="schema-cname_pool--members--ratio"></a>

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

- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/cname_pool/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
