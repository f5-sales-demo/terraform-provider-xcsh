---
page_title: "default_pool.origin_servers"
subcategory: "Load Balancing"
description: "List of origin servers in this pool."
xcsh_docs: {"aliases": ["backend servers", "default pool origin servers", "origin servers", "upstream servers"], "body_bytes": 4043, "body_sha256": "sha256:965e3c34d70782ddcb86a0d2939d5beb1725fc201db2872d6a17b115d839635c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_name", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["default pool origin servers cbip service"], "anchor": "section", "description": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "cbip_service"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers consul service"], "anchor": "section", "description": "Specify origin server with HashiCorp Consul service name and site information.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "consul_service"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers custom endpoint object"], "anchor": "section", "description": "Specify origin server with a reference to endpoint object.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "custom_endpoint_object"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "k8s_service"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers labels"], "anchor": "schema-default_pool--origin_servers--labels", "description": "Add Labels for this origin server, these labels can be used to form subset.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["default pool origin servers private ip"], "anchor": "section", "description": "Specify origin server with private or public IP address and site information.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers private name"], "anchor": "section", "description": "Specify origin server with private or public DNS name and site information.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_name"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "public_ip"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "public_name"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers vn private ip"], "anchor": "section", "description": "Specify origin server with IP on Virtual Network.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "vn_private_ip"], "syntax": "block", "type": "object"}, {"aliases": ["default pool origin servers vn private name"], "anchor": "section", "description": "Specify origin server with DNS name on Virtual Network.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "origin_servers", "vn_private_name"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of origin servers in this pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- default_pool.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/cbip_service/): complete subsection reference.

- [consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/consul_service/): complete subsection reference.

- [custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/custom_endpoint_object/): complete subsection reference.

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/): complete subsection reference.

<a id="schema-default_pool--origin_servers--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

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

- [private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/): complete subsection reference.

- [private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_name/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/public_name/): complete subsection reference.

- [vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_ip/): complete subsection reference.

- [vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_name/): complete subsection reference.
