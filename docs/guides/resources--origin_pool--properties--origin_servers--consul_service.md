---
page_title: "origin_servers.consul_service"
subcategory: "Load Balancing"
description: "origin_servers.consul_service for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 3310, "body_sha256": "sha256:681278a567b509e08ed814c5f359ea7873a04ef1a80639987f44e9f05e21ca3c", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:inside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:site_locator", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "docs/guides/resources--origin_pool--properties--origin_servers--consul_service.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "consul_service"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/consul_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.consul_service for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.consul_service

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- origin_servers.consul_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with HashiCorp Consul service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

Terraform syntax:

```terraform
consul_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](resources--origin_pool--properties--origin_servers--consul_service--inside_network.md): complete subsection reference.

- [outside_network](resources--origin_pool--properties--origin_servers--consul_service--outside_network.md): complete subsection reference.

<a id="schema-origin_servers--consul_service--service_name"></a>

### service_name property

Type: `"string"`. Optional.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [site_locator](resources--origin_pool--properties--origin_servers--consul_service--site_locator.md): complete subsection reference.

- [snat_pool](resources--origin_pool--properties--origin_servers--consul_service--snat_pool.md): complete subsection reference.

## Next pages

- [origin_servers.consul_service.inside_network](resources--origin_pool--properties--origin_servers--consul_service--inside_network.md)
- [origin_servers.consul_service.outside_network](resources--origin_pool--properties--origin_servers--consul_service--outside_network.md)
- [origin_servers.consul_service.site_locator](resources--origin_pool--properties--origin_servers--consul_service--site_locator.md)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--properties--origin_servers--consul_service--snat_pool.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
