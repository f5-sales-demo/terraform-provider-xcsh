---
page_title: "blocking"
subcategory: "Security"
description: "blocking for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1453, "body_sha256": "sha256:3a3a7c1e49a5fd555354d6c10e1ec170c6cf7396d743c61b6a890a2203144d5c", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:blocking", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/blocking/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["blocking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/blocking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocking for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocking

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- blocking

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: blocking, monitoring\] Enable this option

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

- [blocking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/blocking/#section)
- [monitoring](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/monitoring/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
