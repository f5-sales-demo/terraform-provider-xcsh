---
page_title: "origin_servers.origin_servers"
subcategory: ""
description: "origin_servers.origin_servers for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3522, "body_sha256": "sha256:4871b3954c93ccd5a1ae6fc1b7200df8d524af22ef9ec9644dcbc0a8e8d1c8e8", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--origin_servers.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- origin_servers.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("no_preference",
    "site_preferences"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
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

- [k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md): complete subsection reference.

- [no_preference](resources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md): complete subsection reference.

- [public_ip](resources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md): complete subsection reference.

- [public_name](resources--dns_proxy--properties--origin_servers--origin_servers--public_name.md): complete subsection reference.

- [site_preferences](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md)
- [origin_servers.origin_servers.no_preference](resources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md)
- [origin_servers.origin_servers.public_ip](resources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md)
- [origin_servers.origin_servers.public_name](resources--dns_proxy--properties--origin_servers--origin_servers--public_name.md)
- [origin_servers.origin_servers.site_preferences](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
