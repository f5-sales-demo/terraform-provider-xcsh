---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.private_ip for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 5701, "body_sha256": "sha256:e728f06ca3d24d7681e0ba5319f7a8f0a821bd1c53cb36ea35b12fc323fae18e", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:inside_network", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:outside_network", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:segment", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:site_locator", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.private_ip for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.private_ip

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- origin_pools.pools.origin_servers.origin_servers.private_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/inside_network/): complete subsection reference.

<a id="schema-origin_pools--pools--origin_servers--origin_servers--private_ip--ip"></a>

### ip property

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/outside_network/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/): complete subsection reference.

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/inside_network/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/outside_network/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
