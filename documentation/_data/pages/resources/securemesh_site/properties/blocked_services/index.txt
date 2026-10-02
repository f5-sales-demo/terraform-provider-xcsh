---
page_title: "blocked_services"
subcategory: ""
description: "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site."
xcsh_docs: {"aliases": ["blocked services"], "body_bytes": 2018, "body_sha256": "sha256:93c81902d306bbbe4606b7d5b87a262951be9948d81708a8725aa17b8893398c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:blocked_services", "parent_id": "xcsh-docs:resources:securemesh_site:reference", "path": "documentation/resources/securemesh_site/properties/blocked_services/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3001130311132323-3122130302103132-1323300023312300-2032011103233332-2301123302122222-2313012323332111-2033321031100322-2002130010211231", "registry_path": "docs/guides/resources--securemesh_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "sections": [{"aliases": ["blocked service"], "anchor": "section", "description": "Blocking or denial configuration", "document_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,ssh", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:ssh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:dns,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "blocked_services.blocked_service:ConflictingListObjectAttributes:ssh,web_user_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service:web_user_interface", "type": "conflicts"}], "schema_path": ["blocked_services", "blocked_service"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- blocked_services

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: blocked\_services, default\_blocked\_services; Default: default\_blocked\_services\]
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

OneOf alternatives in this subsection:

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/#section)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_blocked_services/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
