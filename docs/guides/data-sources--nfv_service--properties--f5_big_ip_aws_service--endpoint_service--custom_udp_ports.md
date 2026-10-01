---
page_title: "f5_big_ip_aws_service.endpoint_service.custom_udp_ports"
subcategory: ""
description: "f5_big_ip_aws_service.endpoint_service.custom_udp_ports for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2204, "body_sha256": "sha256:68c165e34eb95b94083e9b4778080792fc49f944d9d5c2417e2f8db1288c4dfc", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service:custom_udp_ports", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:endpoint_service", "path": "docs/guides/data-sources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service--custom_udp_ports.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["f5_big_ip_aws_service", "endpoint_service", "custom_udp_ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/endpoint_service/custom_udp_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "f5_big_ip_aws_service.endpoint_service.custom_udp_ports for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.endpoint_service.custom_udp_ports

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [f5_big_ip_aws_service](data-sources--nfv_service--properties--f5_big_ip_aws_service.md)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="section"></a>

Type: `"single"`. Computed.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

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

## Direct properties

<a id="schema-f5_big_ip_aws_service--endpoint_service--custom_udp_ports--ports"></a>

### ports property

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

## Next pages

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--properties--f5_big_ip_aws_service--endpoint_service.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
