---
page_title: "re_acl.fast_acl_rules.port"
subcategory: ""
description: "re_acl.fast_acl_rules.port for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 3212, "body_sha256": "sha256:a1037f6a72f8003b46487a24c6dd08bd00dce67cba57e39cf1d6e7a8596405a2", "canonical_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "path": "docs/guides/resources--fast_acl--properties--re_acl--fast_acl_rules--port.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl.fast_acl_rules.port for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# re_acl.fast_acl_rules.port

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [re_acl](resources--fast_acl--properties--re_acl.md)
- [re_acl.fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md)
- re_acl.fast_acl_rules.port

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
```

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

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all](resources--fast_acl--properties--re_acl--fast_acl_rules--port--all.md): complete subsection reference.

- [dns](resources--fast_acl--properties--re_acl--fast_acl_rules--port--dns.md): complete subsection reference.

<a id="schema-re_acl--fast_acl_rules--port--user_defined"></a>

### user_defined property

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

- [re_acl.fast_acl_rules.port.all](resources--fast_acl--properties--re_acl--fast_acl_rules--port--all.md)
- [re_acl.fast_acl_rules.port.dns](resources--fast_acl--properties--re_acl--fast_acl_rules--port--dns.md)
- [re_acl.fast_acl_rules](resources--fast_acl--properties--re_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
