---
page_title: "no_policer"
subcategory: ""
description: "no_policer for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 1201, "body_sha256": "sha256:8a4811cdf6c84cc5fa63aa3e1e7a05447c617eb8e0f29ae9a6a29aa702e22b13", "canonical_id": "xcsh-docs:resources:forwarding_class:properties:no_policer", "child_ids": [], "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:properties:no_policer", "parent_id": "xcsh-docs:resources:forwarding_class:reference", "path": "docs/guides/resources--forwarding_class--properties--no_policer.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_policer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/properties/no_policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_policer for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_policer

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md)
- [Property reference](resources--forwarding_class--reference.md)
- no_policer

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_policer, policer; Default: no\_policer\] Enable this option

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

- [no_policer](resources--forwarding_class--properties--no_policer.md#section)
- [policer](resources--forwarding_class--properties--policer.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_policer = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--forwarding_class--reference.md)
- [xcsh_forwarding_class](../resources/forwarding_class.md)
