---
page_title: "blocked_services"
subcategory: ""
description: "blocked_services for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1919, "body_sha256": "sha256:3e8eebcab1c0e18ee75aa8040bbb0f7348f52fb76df45d262753c356b2631ef7", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:blocked_services", "parent_id": "xcsh-docs:resources:securemesh_site:reference", "path": "documentation/resources/securemesh_site/properties/blocked_services/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
