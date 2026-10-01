---
page_title: "enable_trust_client_ip_headers"
subcategory: "Load Balancing"
description: "enable_trust_client_ip_headers for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3827, "body_sha256": "sha256:5669a8f44c1faba32370641908ab2113cf2f58a6c986d3c70b56561c65ad5ea3", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_trust_client_ip_headers", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_trust_client_ip_headers", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--enable_trust_client_ip_headers.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_trust_client_ip_headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_trust_client_ip_headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_trust_client_ip_headers for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_trust_client_ip_headers

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- enable_trust_client_ip_headers

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Trust Client IP Headers List. List of Client IP Headers.

Upstream description:

List of Client IP Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("client_ip_headers")}
```

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
enable_trust_client_ip_headers {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enable_trust_client_ip_headers--client_ip_headers"></a>

### client_ip_headers property

Type: `["list", "string"]`. Optional.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined..

Upstream description:

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
