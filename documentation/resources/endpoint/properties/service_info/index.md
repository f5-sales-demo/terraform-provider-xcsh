---
page_title: "service_info"
subcategory: "Networking"
description: "service_info for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 4648, "body_sha256": "sha256:70e2e71bce0dd44639d4a6eb95d11f7c84da764ada5f14a82910c3663c244b20", "child_ids": ["xcsh-docs:resources:endpoint:properties:service_info:service_selector"], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:service_info", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "documentation/resources/endpoint/properties/service_info/index.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["service_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/service_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service_info for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
