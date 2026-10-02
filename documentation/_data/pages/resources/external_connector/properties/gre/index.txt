---
page_title: "gre"
subcategory: ""
description: "External Connector with GRE tunnel."
xcsh_docs: {"aliases": ["gre"], "body_bytes": 1753, "body_sha256": "sha256:10713072ae6615e1c75ebde18e70e8c441c14ab5531e5c1149ef14ba9269d4ef", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre", "parent_id": "xcsh-docs:resources:external_connector:reference", "path": "documentation/resources/external_connector/properties/gre/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre"], "schema_version": 1, "sections": [{"aliases": ["gre parameters"], "anchor": "section", "description": "GRE configuration parameters required for GRE Connection type.", "document_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_inside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:segment,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "type": "conflicts"}, {"anchor": "schema-gre--gre_parameters--tunnel_mtu", "enforcement": "provider-schema", "group": "gre.gre_parameters:RequiredObjectAttributes:tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gre.gre_parameters:RequiredObjectAttributes:tunnel_eps,tunnel_mtu", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps", "type": "requires"}], "schema_path": ["gre", "gre_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "External Connector with GRE tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["external_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
