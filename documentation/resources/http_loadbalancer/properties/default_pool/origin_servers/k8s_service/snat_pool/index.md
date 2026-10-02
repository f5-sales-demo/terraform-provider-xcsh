---
page_title: "default_pool.origin_servers.k8s_service.snat_pool"
subcategory: "Load Balancing"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers k8s service snat pool", "origin servers", "upstream servers"], "body_bytes": 2787, "body_sha256": "sha256:40a2e6f48e793bff11a57d0ee29e5bbb53424e92b36ddaa6c4c1e79806fec735", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2301203302113232-3211031310211322-3111320331112100-1303312113222221-0231203303112033-0220333300001100-2202001011012031-0130213321022303", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.k8s_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.k8s_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:no_snat_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service", "snat_pool", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service.snat_pool

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- default_pool.origin_servers.k8s_service.snat_pool

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

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/no_snat_pool/)
- [default_pool.origin_servers.k8s_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/snat_pool/snat_pool/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
