---
page_title: "default_pool.origin_servers.private_ip.site_locator"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.private_ip.site_locator for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2288, "body_sha256": "sha256:7f3b7af47c126846ae5ff8395d96e559f26cbad5da5f5d357fde4e4e704de9c1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:site", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:virtual_site"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.private_ip.site_locator for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_ip.site_locator

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md)
- default_pool.origin_servers.private_ip.site_locator

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

- [site](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator--site.md): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_ip.site_locator.site](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator--site.md)
- [default_pool.origin_servers.private_ip.site_locator.virtual_site](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--site_locator--virtual_site.md)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
