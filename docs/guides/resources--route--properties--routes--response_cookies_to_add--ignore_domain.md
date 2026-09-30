---
page_title: "routes.response_cookies_to_add.ignore_domain"
subcategory: ""
description: "routes.response_cookies_to_add.ignore_domain for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 1018, "body_sha256": "sha256:d295ef391ee8de613a884d79143409d63547d01124158bf80a4f184e90bc32e4", "canonical_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_domain", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_domain", "parent_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "path": "docs/guides/resources--route--properties--routes--response_cookies_to_add--ignore_domain.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "response_cookies_to_add", "ignore_domain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/response_cookies_to_add/ignore_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.response_cookies_to_add.ignore_domain for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.response_cookies_to_add.ignore_domain

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- [routes.response_cookies_to_add](resources--route--properties--routes--response_cookies_to_add.md)
- routes.response_cookies_to_add.ignore_domain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.response_cookies_to_add](resources--route--properties--routes--response_cookies_to_add.md)
- [xcsh_route](../resources/route.md)
