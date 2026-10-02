---
page_title: "aws.byoc"
subcategory: ""
description: "List of Bring You Own Connection."
xcsh_docs: {"aliases": ["aws byoc"], "body_bytes": 1595, "body_sha256": "sha256:021507ae470c849d427605e011b26120d0469bbb38b205f4702dfd07031c4881", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:aws:byoc:connections"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws", "path": "documentation/resources/cloud_link/properties/aws/byoc/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws.byoc:RequiredObjectAttributes:connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc"], "schema_version": 1, "sections": [{"aliases": ["connections"], "anchor": "section", "description": "List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but will be used for connecting sites and REs.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["aws", "byoc", "connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Bring You Own Connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

List of Bring You Own Connection.

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

## Next pages

- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
