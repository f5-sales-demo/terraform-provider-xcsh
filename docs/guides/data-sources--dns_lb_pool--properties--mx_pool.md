---
page_title: "mx_pool"
subcategory: ""
description: "mx_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1813, "body_sha256": "sha256:6f2fd1b396b2873444028a52dbe7f11ed7ff8416b5a2c4da3ecb82d10051ee43", "canonical_id": "xcsh-docs:data-sources:dns_lb_pool:properties:mx_pool", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:properties:mx_pool:members"], "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:properties:mx_pool", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:reference", "path": "docs/guides/data-sources--dns_lb_pool--properties--mx_pool.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mx_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/properties/mx_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mx_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mx_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- mx_pool

<a id="section"></a>

Type: `"single"`. Computed.

Pool for MX Record.

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

<a id="schema-mx_pool--max_answers"></a>

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

- [members](data-sources--dns_lb_pool--properties--mx_pool--members.md): complete subsection reference.

## Next pages

- [mx_pool.members](data-sources--dns_lb_pool--properties--mx_pool--members.md)
- [Property reference](data-sources--dns_lb_pool--reference.md)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md)
