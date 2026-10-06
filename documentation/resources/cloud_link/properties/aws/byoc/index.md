---
page_title: "aws.byoc"
subcategory: ""
description: "List of Bring You Own Connection."
xcsh_docs: {"aliases": ["aws byoc"], "body_bytes": 1184, "body_sha256": "sha256:6d283165456a07fac34be0d9cb68fe99512538985e5fcceb68bce338c33c6bc4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws", "path": "documentation/resources/cloud_link/properties/aws/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc"], "schema_version": 1, "sections": [{"aliases": ["aws byoc connections"], "anchor": "section", "description": "List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but will be used for connecting sites and REs.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-aws--byoc--connections--user_assigned_name", "enforcement": "provider-schema", "group": "aws.byoc.connections:ConflictingListObjectAttributes:system_generated_name,user_assigned_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws.byoc.connections:ConflictingListObjectAttributes:system_generated_name,user_assigned_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:system_generated_name", "type": "conflicts"}, {"anchor": "schema-aws--byoc--connections--bgp_asn", "enforcement": "provider-schema", "group": "aws.byoc.connections:RequiredListObjectAttributes:bgp_asn,connection_id,region,vlan", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}, {"anchor": "schema-aws--byoc--connections--connection_id", "enforcement": "provider-schema", "group": "aws.byoc.connections:RequiredListObjectAttributes:bgp_asn,connection_id,region,vlan", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}, {"anchor": "schema-aws--byoc--connections--region", "enforcement": "provider-schema", "group": "aws.byoc.connections:RequiredListObjectAttributes:bgp_asn,connection_id,region,vlan", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}, {"anchor": "schema-aws--byoc--connections--vlan", "enforcement": "provider-schema", "group": "aws.byoc.connections:RequiredListObjectAttributes:bgp_asn,connection_id,region,vlan", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}], "schema_path": ["aws", "byoc", "connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of Bring You Own Connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- aws.byoc

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bring Your Own Connections. List of Bring You Own Connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
```

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
byoc {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/): complete subsection reference.
