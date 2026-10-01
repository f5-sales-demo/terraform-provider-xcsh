---
page_title: "default_pool.origin_servers.k8s_service"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.k8s_service for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5476, "body_sha256": "sha256:b368f294497b3be1f2e45fdc6c5bdc62525885c1c12c6da4db0cb48c0d33d66a", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:inside_network", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:outside_network", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:site_locator", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:snat_pool", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service:vk8s_networks"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:k8s_service", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "k8s_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/k8s_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.k8s_service for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.k8s_service

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.k8s_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--inside_network.md): complete subsection reference.

- [outside_network](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--outside_network.md): complete subsection reference.

<a id="schema-default_pool--origin_servers--k8s_service--protocol"></a>

### protocol property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-default_pool--origin_servers--k8s_service--service_name"></a>

### service_name property

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator.md): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool.md): complete subsection reference.

- [vk8s_networks](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--vk8s_networks.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.k8s_service.inside_network](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--inside_network.md)
- [default_pool.origin_servers.k8s_service.outside_network](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--outside_network.md)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--site_locator.md)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--snat_pool.md)
- [default_pool.origin_servers.k8s_service.vk8s_networks](resources--http_loadbalancer--properties--default_pool--origin_servers--k8s_service--vk8s_networks.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
