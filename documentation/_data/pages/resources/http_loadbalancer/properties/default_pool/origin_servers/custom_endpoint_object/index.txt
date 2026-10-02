---
page_title: "default_pool.origin_servers.custom_endpoint_object"
subcategory: "Load Balancing"
description: "Specify origin server with a reference to endpoint object."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers custom endpoint object", "origin servers", "upstream servers"], "body_bytes": 1886, "body_sha256": "sha256:c3db7d5ebef2907cd732792ad9196d569730ffc52d6a5e275a590de7ea48bbdc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object:endpoint"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3320101200101333-0101332020212110-3023110121121203-0211302010232331-0301203130200233-3221233212020201-1322313330211100-1110201120130123", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "custom_endpoint_object"], "schema_version": 1, "sections": [{"aliases": ["endpoint"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object:endpoint", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--origin_servers--custom_endpoint_object--endpoint--name", "enforcement": "provider-schema", "group": "default_pool.origin_servers.custom_endpoint_object.endpoint:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object:endpoint", "type": "requires"}], "schema_path": ["default_pool", "origin_servers", "custom_endpoint_object", "endpoint"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify origin server with a reference to endpoint object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.custom_endpoint_object

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- default_pool.origin_servers.custom_endpoint_object

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

## Direct properties

- [endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/endpoint/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.custom_endpoint_object.endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/endpoint/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
