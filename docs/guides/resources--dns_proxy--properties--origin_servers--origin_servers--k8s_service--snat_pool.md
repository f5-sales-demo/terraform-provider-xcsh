---
page_title: "origin_servers.origin_servers.k8s_service.snat_pool"
subcategory: ""
description: "origin_servers.origin_servers.k8s_service.snat_pool for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2173, "body_sha256": "sha256:625b7fe1fba4177add3bf0b1a049c4a7114de81adc4559384b0d520822fc14a3", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:no_snat_pool", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool:snat_pool"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.k8s_service.snat_pool for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service.snat_pool

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- origin_servers.origin_servers.k8s_service.snat_pool

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

- [no_snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md)
- [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md)
- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
