---
page_title: "origin_servers.origin_servers.k8s_service.site_locator"
subcategory: ""
description: "origin_servers.origin_servers.k8s_service.site_locator for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2228, "body_sha256": "sha256:439482f7bb9e4eae1ed484a5801a6374236f65713ce9b284bbda92447ce7c80a", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:site", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.k8s_service.site_locator for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- origin_servers.origin_servers.k8s_service.site_locator

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

- [site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md): complete subsection reference.

- [virtual_site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service.site_locator.site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
