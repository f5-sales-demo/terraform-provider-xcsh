---
page_title: "re_acl.fast_acl_rules"
subcategory: ""
description: "re_acl.fast_acl_rules for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2794, "body_sha256": "sha256:002ad6f09b75466fd7c77bb8987734fc5c91877d4c6190e61bdf299043cc1f2a", "canonical_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl", "path": "docs/guides/resources--fast_acl--properties--re_acl--fast_acl_rules.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_acl", "fast_acl_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_acl.fast_acl_rules for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [re_acl](resources--fast_acl--properties--re_acl.md)
- re_acl.fast_acl_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
fast_acl_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](resources--fast_acl--properties--re_acl--fast_acl_rules--action.md): complete subsection reference.

- [ip_prefix_set](resources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set.md): complete subsection reference.

- [metadata](resources--fast_acl--properties--re_acl--fast_acl_rules--metadata.md): complete subsection reference.

- [port](resources--fast_acl--properties--re_acl--fast_acl_rules--port.md): complete subsection reference.

- [prefix](resources--fast_acl--properties--re_acl--fast_acl_rules--prefix.md): complete subsection reference.

## Next pages

- [re_acl.fast_acl_rules.action](resources--fast_acl--properties--re_acl--fast_acl_rules--action.md)
- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--properties--re_acl--fast_acl_rules--ip_prefix_set.md)
- [re_acl.fast_acl_rules.metadata](resources--fast_acl--properties--re_acl--fast_acl_rules--metadata.md)
- [re_acl.fast_acl_rules.port](resources--fast_acl--properties--re_acl--fast_acl_rules--port.md)
- [re_acl.fast_acl_rules.prefix](resources--fast_acl--properties--re_acl--fast_acl_rules--prefix.md)
- [re_acl](resources--fast_acl--properties--re_acl.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
