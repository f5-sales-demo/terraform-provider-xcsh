---
page_title: "re_acl.fast_acl_rules.port"
subcategory: ""
description: "re_acl.fast_acl_rules.port for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2788, "body_sha256": "sha256:243d0439613ab0a352e9cb404aebfba101409bbadad2ceb58a066d5653156c82", "canonical_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port:dns"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules:port", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:re_acl:fast_acl_rules", "path": "docs/guides/data-sources--fast_acl--properties--re_acl--fast_acl_rules--port.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/re_acl/fast_acl_rules/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl.fast_acl_rules.port for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.port

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
- [Property reference](data-sources--fast_acl--reference.md)
- [re_acl](data-sources--fast_acl--properties--re_acl.md)
- [re_acl.fast_acl_rules](data-sources--fast_acl--properties--re_acl--fast_acl_rules.md)
- re_acl.fast_acl_rules.port

<a id="section"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

## Direct properties

- [all](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--all.md): complete subsection reference.

- [dns](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--dns.md): complete subsection reference.

<a id="schema-re_acl--fast_acl_rules--port--user_defined"></a>

### user_defined property

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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
    "minimum": 1
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

## Next pages

- [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--all.md)
- [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--properties--re_acl--fast_acl_rules--port--dns.md)
- [re_acl.fast_acl_rules](data-sources--fast_acl--properties--re_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
