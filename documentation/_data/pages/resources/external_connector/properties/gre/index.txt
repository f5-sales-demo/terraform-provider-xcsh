---
page_title: "gre"
subcategory: ""
description: "gre for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1753, "body_sha256": "sha256:10713072ae6615e1c75ebde18e70e8c441c14ab5531e5c1149ef14ba9269d4ef", "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre", "parent_id": "xcsh-docs:resources:external_connector:reference", "path": "documentation/resources/external_connector/properties/gre/index.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["gre"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- gre

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: gre, ipsec\] GRE. External Connector with GRE tunnel.

Upstream description:

External Connector with GRE tunnel.

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

- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/#section)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
gre {
  # Configure direct properties listed below.
}
```

## Direct properties

- [gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/): complete subsection reference.

## Next pages

- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/gre/gre_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
