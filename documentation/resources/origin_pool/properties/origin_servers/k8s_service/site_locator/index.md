---
page_title: "origin_servers.k8s_service.site_locator"
subcategory: "Load Balancing"
description: "origin_servers.k8s_service.site_locator for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2362, "body_sha256": "sha256:3d2e52513588b13b3a6c91cd9c4240d4c14e6fd7617640b2bdea391c69cbf9b6", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:site", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "path": "documentation/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.k8s_service.site_locator for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/)
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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/): complete subsection reference.

## Next pages

- [origin_servers.k8s_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/)
- [origin_servers.k8s_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/)
- [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
