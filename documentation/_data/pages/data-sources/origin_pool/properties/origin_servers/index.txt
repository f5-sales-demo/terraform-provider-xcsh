---
page_title: "origin_servers"
subcategory: "Load Balancing"
description: "List of origin servers in this pool."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers"], "body_bytes": 3581, "body_sha256": "sha256:b8fbd2cf47a582a4a176abb074ac6468156fe4c2aa230d1b2767a1d3e72e3b5f", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:cbip_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_name"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "documentation/data-sources/origin_pool/properties/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "sections": [{"aliases": ["origin servers cbip service"], "anchor": "section", "description": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:cbip_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "cbip_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers consul service"], "anchor": "section", "description": "Specify origin server with HashiCorp Consul service name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "consul_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers custom endpoint object"], "anchor": "section", "description": "Specify origin server with a reference to endpoint object.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "custom_endpoint_object"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "k8s_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers labels"], "anchor": "schema-origin_servers--labels", "description": "Add Labels for this origin server, these labels can be used to form subset.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["origin servers private ip"], "anchor": "section", "description": "Specify origin server with private or public IP address and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers private name"], "anchor": "section", "description": "Specify origin server with private or public DNS name and site information.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_name", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers vn private ip"], "anchor": "section", "description": "Specify origin server with IP on Virtual Network.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers vn private name"], "anchor": "section", "description": "Specify origin server with DNS name on Virtual Network.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_name"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of origin servers in this pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["origin_poolCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
