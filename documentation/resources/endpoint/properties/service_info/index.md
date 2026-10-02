---
page_title: "service_info"
subcategory: "Networking"
description: "Specifies whether endpoint service is discovered by name or labels."
xcsh_docs: {"aliases": ["service info"], "body_bytes": 4648, "body_sha256": "sha256:70e2e71bce0dd44639d4a6eb95d11f7c84da764ada5f14a82910c3663c244b20", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:endpoint:properties:service_info:service_selector"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:service_info", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "documentation/resources/endpoint/properties/service_info/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3023332231310113-2111321122221101-1300032112031023-0133010100210121-2002200120323201-3322322203231112-3301330002231202-3022103031323110", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [{"anchor": "schema-service_info--service_name", "enforcement": "provider-schema", "group": "service_info:ConflictingObjectAttributes:service_name,service_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:service_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service_info:ConflictingObjectAttributes:service_name,service_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:service_info:service_selector", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service_info"], "schema_version": 1, "sections": [{"aliases": ["discovery type"], "anchor": "schema-service_info--discovery_type", "description": "Specifies the type of discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.", "document_id": "xcsh-docs:resources:endpoint:properties:service_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service_info", "discovery_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["service name"], "anchor": "schema-service_info--service_name", "description": "Exclusive with Name of the service to discover with an optional namespace and cluster identifier. The format is service_name.namespace_name:cluster_identifier for K8s and service_name:cluster_identifier for Consul Endpoint will be discovered in all discovery objects where the cluster identifier matches. If cluster", "document_id": "xcsh-docs:resources:endpoint:properties:service_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service_info", "service_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["service selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:endpoint:properties:service_info:service_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service_info--service_selector--expressions", "enforcement": "provider-schema", "group": "service_info.service_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:service_info:service_selector", "type": "requires"}], "schema_path": ["service_info", "service_selector"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/service_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies whether endpoint service is discovered by name or labels.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service_info

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- service_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether endpoint service is discovered by name or labels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("service_name",
    "service_selector")}
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
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

Terraform syntax:

```terraform
service_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service_info--discovery_type"></a>

### discovery_type property

Type: `"string"`. Optional.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Upstream description:

Specifies the type of discovery

Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover
from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"),
}
```

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

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/service_selector/): complete subsection reference.

## Next pages

- [service_info.service_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/service_info/service_selector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
