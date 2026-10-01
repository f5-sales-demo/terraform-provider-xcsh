---
page_title: "re_select.specific_re"
subcategory: ""
description: "re_select.specific_re for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2489, "body_sha256": "sha256:efaf92dc4b0b1994030a70718857d7dd3c69337ba50b311eb22dfc8639fe2863", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select:specific_re", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:re_select", "path": "docs/guides/data-sources--securemesh_site_v2--properties--re_select--specific_re.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["re_select", "specific_re"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/re_select/specific_re/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "re_select.specific_re for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_select.specific_re

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [re_select](data-sources--securemesh_site_v2--properties--re_select.md)
- re_select.specific_re

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-re_select--specific_re--backup_re"></a>

### backup_re property

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

Primary RE Geography. Select primary RE for this site.

Upstream description:

Select primary RE for this site.

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

- [re_select](data-sources--securemesh_site_v2--properties--re_select.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
