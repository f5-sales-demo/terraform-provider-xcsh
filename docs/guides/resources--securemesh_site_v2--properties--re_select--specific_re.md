---
page_title: "re_select.specific_re"
subcategory: ""
description: "re_select.specific_re for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2636, "body_sha256": "sha256:422328e6e53d3f235fd3df9437dd98de85d0a1818a4c07fc32ffd4285dfabf58", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select:specific_re", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:re_select", "path": "docs/guides/resources--securemesh_site_v2--properties--re_select--specific_re.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_select", "specific_re"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/re_select/specific_re/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_select.specific_re for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# re_select.specific_re

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [re_select](resources--securemesh_site_v2--properties--re_select.md)
- re_select.specific_re

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

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

Terraform syntax:

```terraform
specific_re {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-re_select--specific_re--backup_re"></a>

### backup_re property

Type: `"string"`. Optional.

Select backup RE for this site, cannot be the same as Primary RE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-re_select--specific_re--primary_re"></a>

### primary_re property

Type: `"string"`. Optional.

Primary RE Geography. Select primary RE for this site.

Upstream description:

Select primary RE for this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [re_select](resources--securemesh_site_v2--properties--re_select.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
