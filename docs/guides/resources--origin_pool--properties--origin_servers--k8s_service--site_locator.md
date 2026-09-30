---
page_title: "origin_servers.k8s_service.site_locator"
subcategory: "Load Balancing"
description: "origin_servers.k8s_service.site_locator for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1864, "body_sha256": "sha256:d248a829d81b604146d6543de088c1f37633a89f16ff49a4dca9b1b64c68b9ca", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:site", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "path": "docs/guides/resources--origin_pool--properties--origin_servers--k8s_service--site_locator.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.k8s_service.site_locator for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [origin_servers.k8s_service](resources--origin_pool--properties--origin_servers--k8s_service.md)
- origin_servers.k8s_service.site_locator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md): complete subsection reference.

- [virtual_site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_servers.k8s_service.site_locator.site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md)
- [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md)
- [origin_servers.k8s_service](resources--origin_pool--properties--origin_servers--k8s_service.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
