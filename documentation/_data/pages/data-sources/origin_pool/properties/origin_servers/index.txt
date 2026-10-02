---
page_title: "origin_servers"
subcategory: "Load Balancing"
description: "List of origin servers in this pool."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers"], "body_bytes": 5437, "body_sha256": "sha256:e7264001074fb42ffcd75e0a6cdfda0e3cf902e8c81fa6eb84bb9373f4f6713e", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:cbip_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_name"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "documentation/data-sources/origin_pool/properties/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "cbip service", "origin servers", "upstream servers"], "anchor": "section", "description": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:cbip_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "cbip_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "consul service", "origin servers", "upstream servers"], "anchor": "section", "description": "Specify origin server with HashiCorp Consul service name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "consul_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "custom endpoint object", "origin servers", "upstream servers"], "anchor": "section", "description": "Specify origin server with a reference to endpoint object.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "custom_endpoint_object"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "k8s service", "origin servers", "upstream servers"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "k8s_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "labels", "origin servers", "upstream servers"], "anchor": "schema-origin_servers--labels", "description": "Add Labels for this origin server, these labels can be used to form subset.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["backend servers", "origin servers", "private ip", "upstream servers"], "anchor": "section", "description": "Specify origin server with private or public IP address and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "private name", "upstream servers"], "anchor": "section", "description": "Specify origin server with private or public DNS name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public ip", "upstream servers"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public name", "upstream servers"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "upstream servers", "vn private ip"], "anchor": "section", "description": "Specify origin server with IP on Virtual Network.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "upstream servers", "vn private name"], "anchor": "section", "description": "Specify origin server with DNS name on Virtual Network.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_name"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of origin servers in this pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- origin_servers

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

- [cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/cbip_service/): complete subsection reference.

- [consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/): complete subsection reference.

- [custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/custom_endpoint_object/): complete subsection reference.

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/): complete subsection reference.

<a id="schema-origin_servers--labels"></a>

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

- [private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/): complete subsection reference.

- [private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_name/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/public_name/): complete subsection reference.

- [vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/vn_private_ip/): complete subsection reference.

- [vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/vn_private_name/): complete subsection reference.

## Next pages

- [origin_servers.cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/cbip_service/)
- [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/consul_service/)
- [origin_servers.custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/custom_endpoint_object/)
- [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/)
- [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_ip/)
- [origin_servers.private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/private_name/)
- [origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/public_ip/)
- [origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/public_name/)
- [origin_servers.vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/vn_private_ip/)
- [origin_servers.vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/vn_private_name/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
