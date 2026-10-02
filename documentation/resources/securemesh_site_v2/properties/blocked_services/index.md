---
page_title: "blocked_services"
subcategory: ""
description: "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site."
xcsh_docs: {"aliases": ["blocked services"], "body_bytes": 1544, "body_sha256": "sha256:a734621e20f2fa280dd6e95744388a2dd88f1bcaf4f997928df31d793aca2780", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/blocked_services/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "sections": [{"aliases": ["blocked service"], "anchor": "section", "description": "Blocking or denial configuration", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}], "schema_path": ["blocked_services", "blocked_service"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- blocked_services

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/blocked_services/blocked_service/): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/blocked_services/blocked_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
