---
page_title: "origin_servers"
subcategory: "Load Balancing"
description: "origin_servers for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 4229, "body_sha256": "sha256:c69bed00ede578fc1721e972fec5037fd786ba9a8e438f32bd019332202b40ad", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:origin_servers:cbip_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:consul_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:custom_endpoint_object", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:private_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:public_name", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_ip", "xcsh-docs:data-sources:origin_pool:properties:origin_servers:vn_private_name"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "docs/guides/data-sources--origin_pool--properties--origin_servers.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Upstream description:

List of origin servers in this pool.

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

## Direct properties

- [cbip_service](data-sources--origin_pool--properties--origin_servers--cbip_service.md): complete subsection reference.

- [consul_service](data-sources--origin_pool--properties--origin_servers--consul_service.md): complete subsection reference.

- [custom_endpoint_object](data-sources--origin_pool--properties--origin_servers--custom_endpoint_object.md): complete subsection reference.

- [k8s_service](data-sources--origin_pool--properties--origin_servers--k8s_service.md): complete subsection reference.

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

- [private_ip](data-sources--origin_pool--properties--origin_servers--private_ip.md): complete subsection reference.

- [private_name](data-sources--origin_pool--properties--origin_servers--private_name.md): complete subsection reference.

- [public_ip](data-sources--origin_pool--properties--origin_servers--public_ip.md): complete subsection reference.

- [public_name](data-sources--origin_pool--properties--origin_servers--public_name.md): complete subsection reference.

- [vn_private_ip](data-sources--origin_pool--properties--origin_servers--vn_private_ip.md): complete subsection reference.

- [vn_private_name](data-sources--origin_pool--properties--origin_servers--vn_private_name.md): complete subsection reference.

## Next pages

- [origin_servers.cbip_service](data-sources--origin_pool--properties--origin_servers--cbip_service.md)
- [origin_servers.consul_service](data-sources--origin_pool--properties--origin_servers--consul_service.md)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--properties--origin_servers--custom_endpoint_object.md)
- [origin_servers.k8s_service](data-sources--origin_pool--properties--origin_servers--k8s_service.md)
- [origin_servers.private_ip](data-sources--origin_pool--properties--origin_servers--private_ip.md)
- [origin_servers.private_name](data-sources--origin_pool--properties--origin_servers--private_name.md)
- [origin_servers.public_ip](data-sources--origin_pool--properties--origin_servers--public_ip.md)
- [origin_servers.public_name](data-sources--origin_pool--properties--origin_servers--public_name.md)
- [origin_servers.vn_private_ip](data-sources--origin_pool--properties--origin_servers--vn_private_ip.md)
- [origin_servers.vn_private_name](data-sources--origin_pool--properties--origin_servers--vn_private_name.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
