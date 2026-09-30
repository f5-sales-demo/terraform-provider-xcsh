---
page_title: "default_pool.origin_servers.vn_private_name"
subcategory: "Load Balancing"
description: "default_pool.origin_servers.vn_private_name for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2979, "body_sha256": "sha256:3b4603c4ee9a62c567ebc59f961d3c21fcd38162370f93c4b1e64105ff0b3c0c", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name:private_network"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:vn_private_name", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_name.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "origin_servers", "vn_private_name"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/vn_private_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.origin_servers.vn_private_name for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_pool.origin_servers.vn_private_name

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- default_pool.origin_servers.vn_private_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
vn_private_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-default_pool--origin_servers--vn_private_name--dns_name"></a>

### dns_name property

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [private_network](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_name--private_network.md): complete subsection reference.

## Next pages

- [default_pool.origin_servers.vn_private_name.private_network](resources--http_loadbalancer--properties--default_pool--origin_servers--vn_private_name--private_network.md)
- [default_pool.origin_servers](resources--http_loadbalancer--properties--default_pool--origin_servers.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
