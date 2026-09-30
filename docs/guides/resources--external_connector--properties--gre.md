---
page_title: "gre"
subcategory: ""
description: "gre for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1244, "body_sha256": "sha256:aa98904f1554cef6f6d3b4d39c098c92f247afa86d7d3d1d0cc7a07ea7fb972a", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre", "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre", "parent_id": "xcsh-docs:resources:external_connector:reference", "path": "docs/guides/resources--external_connector--properties--gre.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gre

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
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

- [gre](resources--external_connector--properties--gre.md#section)
- [ipsec](resources--external_connector--properties--ipsec.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
gre {
  # Configure direct properties listed below.
}
```

## Direct properties

- [gre_parameters](resources--external_connector--properties--gre--gre_parameters.md): complete subsection reference.

## Next pages

- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- [Property reference](resources--external_connector--reference.md)
- [xcsh_external_connector](../resources/external_connector.md)
