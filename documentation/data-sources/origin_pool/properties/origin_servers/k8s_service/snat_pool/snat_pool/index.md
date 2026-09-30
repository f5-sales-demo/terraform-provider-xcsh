---
page_title: "origin_servers.k8s_service.snat_pool.snat_pool"
subcategory: "Load Balancing"
description: "origin_servers.k8s_service.snat_pool.snat_pool for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2467, "body_sha256": "sha256:8c382d9cc997876368b10f1547635482e7599883aec48d048aeef05438e19bab", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:snat_pool:snat_pool", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:snat_pool", "path": "documentation/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/snat_pool/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["origin_servers", "k8s_service", "snat_pool", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.k8s_service.snat_pool.snat_pool for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.k8s_service.snat_pool.snat_pool

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/)
- [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/)
- [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-origin_servers--k8s_service--snat_pool--snat_pool--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/origin_servers/k8s_service/snat_pool/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
