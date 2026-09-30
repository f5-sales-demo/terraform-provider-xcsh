---
page_title: "blocked_services"
subcategory: ""
description: "blocked_services for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1909, "body_sha256": "sha256:ebc6a945231620f205a4de606bbb18fb40bf1b3ae3b8045df83c2c4fc311e340", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:blocked_services:blocked_service"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:blocked_services", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/blocked_services/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["blocked_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/blocked_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# blocked_services

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
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

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/blocked_services/#section)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/default_blocked_services/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/blocked_services/blocked_service/): complete subsection reference.

## Next pages

- [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/blocked_services/blocked_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
