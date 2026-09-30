---
page_title: "a_pool"
subcategory: ""
description: "a_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2769, "body_sha256": "sha256:bb437031553e093d66f8fded0bdc25130e26cb0bfcad740a81b828e352a47012", "canonical_id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:disable_health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:health_check", "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool:members"], "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:a_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "docs/guides/data-sources--dns_lb_pool--properties--a_pool.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["a_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/a_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "a_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# a_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- a_pool

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

OneOf alternatives in this subsection:

- [a_pool](data-sources--dns_lb_pool--properties--a_pool.md#section)
- [aaaa_pool](data-sources--dns_lb_pool--properties--aaaa_pool.md#section)
- [cname_pool](data-sources--dns_lb_pool--properties--cname_pool.md#section)
- [mx_pool](data-sources--dns_lb_pool--properties--mx_pool.md#section)
- [srv_pool](data-sources--dns_lb_pool--properties--srv_pool.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_health_check](data-sources--dns_lb_pool--properties--a_pool--disable_health_check.md): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--properties--a_pool--health_check.md): complete subsection reference.

<a id="schema-a_pool--max_answers"></a>

### max_answers property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--properties--a_pool--members.md): complete subsection reference.

## Next pages

- [a_pool.disable_health_check](data-sources--dns_lb_pool--properties--a_pool--disable_health_check.md)
- [a_pool.health_check](data-sources--dns_lb_pool--properties--a_pool--health_check.md)
- [a_pool.members](data-sources--dns_lb_pool--properties--a_pool--members.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
