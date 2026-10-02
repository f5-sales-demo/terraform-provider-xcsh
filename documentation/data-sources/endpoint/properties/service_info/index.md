---
page_title: "service_info"
subcategory: "Networking"
description: "Specifies whether endpoint service is discovered by name or labels."
xcsh_docs: {"aliases": ["service info"], "body_bytes": 4000, "body_sha256": "sha256:d19c9fb39045dbed23894df7cf0f576c1b41877ea46cd1a1b406bc7279d4b1ae", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:endpoint:properties:service_info:service_selector"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:endpoint:properties:service_info", "parent_id": "xcsh-docs:data-sources:endpoint:reference", "path": "documentation/data-sources/endpoint/properties/service_info/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2320321032020331-2133003222333220-3030332320301211-3130203301232322-2230031021333133-3303310130333212-1121103320212022-3100021322323222", "registry_path": "docs/guides/data-sources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service_info"], "schema_version": 1, "sections": [{"aliases": ["discovery type"], "anchor": "schema-service_info--discovery_type", "description": "Specifies the type of discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.", "document_id": "xcsh-docs:data-sources:endpoint:properties:service_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service_info", "discovery_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["service name"], "anchor": "schema-service_info--service_name", "description": "Exclusive with Name of the service to discover with an optional namespace and cluster identifier. The format is service_name.namespace_name:cluster_identifier for K8s and service_name:cluster_identifier for Consul Endpoint will be discovered in all discovery objects where the cluster identifier matches. If cluster", "document_id": "xcsh-docs:data-sources:endpoint:properties:service_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service_info", "service_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["service selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:endpoint:properties:service_info:service_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service_info", "service_selector"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/properties/service_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies whether endpoint service is discovered by name or labels.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service_info

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- service_info

<a id="section"></a>

Type: `"single"`. Computed.

Specifies whether endpoint service is discovered by name or labels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

## Direct properties

<a id="schema-service_info--discovery_type"></a>

### discovery_type property

Type: `"string"`. Computed.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Upstream description:

Specifies the type of discovery

Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover
from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID_DISCOVERY",
  "enum": [
    "INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-service_info--service_name"></a>

### service_name property

Type: `"string"`. Computed.

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the..

Upstream description:

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the cluster identifier matches. If cluster identifier is not specified then discovery will be
done in all discovery objects of the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/service_selector/): complete subsection reference.

## Next pages

- [service_info.service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/service_info/service_selector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/properties/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
