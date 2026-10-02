---
page_title: "origin_servers.consul_service.snat_pool"
subcategory: "Load Balancing"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers consul service snat pool", "upstream servers"], "body_bytes": 2442, "body_sha256": "sha256:c858f9c1879a90ea8836f32bd05f8504b1cc9bffb4747b8d8716f86c90a37b3a", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "path": "documentation/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:snat_pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "consul_service", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "consul_service", "snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "consul_service", "snat_pool", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.consul_service.snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/)
- origin_servers.consul_service.snat_pool

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

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [origin_servers.consul_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/)
- [origin_servers.consul_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
