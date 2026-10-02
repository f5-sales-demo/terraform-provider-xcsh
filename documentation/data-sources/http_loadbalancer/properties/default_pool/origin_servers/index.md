---
page_title: "default_pool.origin_servers"
subcategory: "Load Balancing"
description: "List of origin servers in this pool."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers", "origin servers", "upstream servers"], "body_bytes": 6160, "body_sha256": "sha256:cbafab2205dd825101b08821d3c967e6097b680236273a3c56c7d8d556c29fa5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:public_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:public_name", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0012023132220102-0022302223220122-2213110010110133-1121000333032021-3031332302320102-1320111030302223-0112023003033302-3321303330233000", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["cbip service"], "anchor": "section", "description": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "cbip_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["consul service"], "anchor": "section", "description": "Specify origin server with HashiCorp Consul service name and site information.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "consul_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom endpoint object"], "anchor": "section", "description": "Specify origin server with a reference to endpoint object.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "custom_endpoint_object"], "syntax": "attribute", "type": "object"}, {"aliases": ["k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-default_pool--origin_servers--labels", "description": "Add Labels for this origin server, these labels can be used to form subset.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["private ip"], "anchor": "section", "description": "Specify origin server with private or public IP address and site information.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["private name"], "anchor": "section", "description": "Specify origin server with private or public DNS name and site information.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:public_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn private ip"], "anchor": "section", "description": "Specify origin server with IP on Virtual Network.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "vn_private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn private name"], "anchor": "section", "description": "Specify origin server with DNS name on Virtual Network.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "vn_private_name"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of origin servers in this pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- default_pool.origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/cbip_service/): complete subsection reference.

- [consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/): complete subsection reference.

- [custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/): complete subsection reference.

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/): complete subsection reference.

<a id="schema-default_pool--origin_servers--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Add Labels for this origin server, these labels can be used to form subset.

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

- [private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/): complete subsection reference.

- [private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/public_name/): complete subsection reference.

- [vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/): complete subsection reference.

- [vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_name/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/cbip_service/)
- [default_pool.origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/)
- [default_pool.origin_servers.custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/)
- [default_pool.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/)
- [default_pool.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/)
- [default_pool.origin_servers.private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/private_name/)
- [default_pool.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/public_ip/)
- [default_pool.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/public_name/)
- [default_pool.origin_servers.vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/)
- [default_pool.origin_servers.vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_name/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
