---
page_title: "default_pool.origin_servers"
subcategory: "Load Balancing"
description: "default_pool.origin_servers for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 8836, "body_sha256": "sha256:b291dc9331ed0bb8745d813f44683171e68365031891820adfa7df7374eb95bf", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:cbip_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:consul_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:custom_endpoint_object", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_name", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:public_name", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_ip", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- default_pool.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cbip_service",
    "consul_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("cbip_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "custom_endpoint_object"),
  validators.ConflictingListObjectAttributes("consul_service",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("consul_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "k8s_service"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "private_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "public_name"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("custom_endpoint_object",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "private_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "private_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("private_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_ip",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_ip"),
  validators.ConflictingListObjectAttributes("public_name",
    "vn_private_name"),
  validators.ConflictingListObjectAttributes("vn_private_ip",
    "vn_private_name")}
```

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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cbip_service](resources--http_loadbalancer--properties--default_pool--origin_servers--cbip_service.md): complete subsection reference.

- [consul_service](resources--http_loadbalancer--properties--default_pool--origin_servers--consul_service.md): complete subsection reference.

- [custom_endpoint_object](resources--http_loadbalancer--properties--default_pool--origin_servers--custom_endpoint_object.md): complete subsection reference.

- [k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md): complete subsection reference.

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

- [private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md): complete subsection reference.

- [private_name](resources--http_loadbalancer--properties--default_pool--origin_servers--private_name.md): complete subsection reference.

- [public_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--public_ip.md): complete subsection reference.

- [public_name](resources--http_loadbalancer--properties--default_pool--origin_servers--public_name.md): complete subsection reference.

- [vn_private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_ip.md): complete subsection reference.

- [vn_private_name](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_name.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.cbip_service](resources--http_loadbalancer--properties--default_pool--origin_servers--cbip_service.md)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--properties--default_pool--origin_servers--consul_service.md)
- [default_pool.origin_servers.custom_endpoint_object](resources--http_loadbalancer--properties--default_pool--origin_servers--custom_endpoint_object.md)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--private_ip.md)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--properties--default_pool--origin_servers--private_name.md)
- [default_pool.origin_servers.public_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--public_ip.md)
- [default_pool.origin_servers.public_name](resources--http_loadbalancer--properties--default_pool--origin_servers--public_name.md)
- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_ip.md)
- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_name.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
