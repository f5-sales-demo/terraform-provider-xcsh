---
page_title: "origin_servers"
subcategory: "Load Balancing"
description: "List of origin servers in this pool."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers"], "body_bytes": 3668, "body_sha256": "sha256:6ee9a664065009b2a2f1823044336c83314b4eabbc5de03614f8624a73829100", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "documentation/resources/origin_pool/properties/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "sections": [{"aliases": ["origin servers cbip service"], "anchor": "section", "description": "Specify origin server with Classic BIG-IP Service (Virtual Server)", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "cbip_service"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers consul service"], "anchor": "section", "description": "Specify origin server with HashiCorp Consul service name and site information.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "consul_service"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers custom endpoint object"], "anchor": "section", "description": "Specify origin server with a reference to endpoint object.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "custom_endpoint_object"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "k8s_service"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers labels"], "anchor": "schema-origin_servers--labels", "description": "Add Labels for this origin server, these labels can be used to form subset.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["origin servers private ip"], "anchor": "section", "description": "Specify origin server with private or public IP address and site information.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_ip"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers private name"], "anchor": "section", "description": "Specify origin server with private or public DNS name and site information.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "private_name"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_ip"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "public_name"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers vn private ip"], "anchor": "section", "description": "Specify origin server with IP on Virtual Network.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_ip"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers vn private name"], "anchor": "section", "description": "Specify origin server with DNS name on Virtual Network.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "vn_private_name"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of origin servers in this pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["origin_poolCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- origin_servers

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/cbip_service/): complete subsection reference.

- [consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/): complete subsection reference.

- [custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/): complete subsection reference.

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/): complete subsection reference.

<a id="schema-origin_servers--labels"></a>

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

- [private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/): complete subsection reference.

- [private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_name/): complete subsection reference.

- [vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/): complete subsection reference.

- [vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/): complete subsection reference.
