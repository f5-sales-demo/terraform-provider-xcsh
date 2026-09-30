---
page_title: "default_pool.origin_servers.private_ip.snat_pool"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.private_ip.snat_pool for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2134, "body_sha256": "sha256:613d85dd841734871498c222f679fc8ade1f61274d9ffdd7f6e9dff3b3d7ed35", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool:snat_pool"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.private_ip.snat_pool for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md)
- default_pool.origin_servers.private_ip.snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_ip.snat_pool.no_snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool--no_snat_pool.md)
- [default_pool.origin_servers.private_ip.snat_pool.snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip--snat_pool--snat_pool.md)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
