---
page_title: "site_virtual_sites.advertise_where"
subcategory: ""
description: "site_virtual_sites.advertise_where for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3307, "body_sha256": "sha256:63d8e2103f68a8339e0af9f2dde1e66f639f2c10bea1762a6eec207621c09dd0", "canonical_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "child_ids": ["xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:site", "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "parent_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites", "path": "docs/guides/data-sources--proxy--properties--site_virtual_sites--advertise_where.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_virtual_sites/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_virtual_sites.advertise_where for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# site_virtual_sites.advertise_where

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md)
- site_virtual_sites.advertise_where

<a id="section"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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

<a id="schema-site_virtual_sites--advertise_where--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Upstream description:

Exclusive with \[use\_default\_port\] TCP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [site](data-sources--proxy--properties--site_virtual_sites--advertise_where--site.md): complete subsection reference.

- [use_default_port](data-sources--proxy--properties--site_virtual_sites--advertise_where--use_default_port.md): complete subsection reference.

- [virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site.md): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where.site](data-sources--proxy--properties--site_virtual_sites--advertise_where--site.md)
- [site_virtual_sites.advertise_where.use_default_port](data-sources--proxy--properties--site_virtual_sites--advertise_where--use_default_port.md)
- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site.md)
- [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md)
- [xcsh_proxy](../data-sources/proxy.md)
