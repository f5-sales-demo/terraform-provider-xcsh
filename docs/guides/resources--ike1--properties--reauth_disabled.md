---
page_title: "reauth_disabled"
subcategory: ""
description: "reauth_disabled for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 1262, "body_sha256": "sha256:5566bc1b2c3bac7b4d4a40adc543fe9c946a67d734d1e350e672273705884a60", "canonical_id": "xcsh-docs:resources:ike1:properties:reauth_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:reauth_disabled", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "docs/guides/resources--ike1--properties--reauth_disabled.md", "provider_name": "ike1", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["reauth_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/reauth_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "reauth_disabled for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# reauth_disabled

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md)
- [Property reference](resources--ike1--reference.md)
- reauth_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [reauth_disabled](resources--ike1--properties--reauth_disabled.md#section)
- [reauth_timeout_days](resources--ike1--properties--reauth_timeout_days.md#section)
- [reauth_timeout_hours](resources--ike1--properties--reauth_timeout_hours.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--ike1--reference.md)
- [xcsh_ike1](../resources/ike1.md)
