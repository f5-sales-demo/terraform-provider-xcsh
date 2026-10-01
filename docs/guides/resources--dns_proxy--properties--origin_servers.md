---
page_title: "origin_servers"
subcategory: ""
description: "origin_servers for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1385, "body_sha256": "sha256:db8119590abda58e09cf542dae8d887d8cd668161ad881ed42235810c71eabe0", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "docs/guides/resources--dns_proxy--properties--origin_servers.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- origin_servers

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the DNS proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers")}
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
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md): complete subsection reference.

- [origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md): complete subsection reference.

## Next pages

- [origin_servers.health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- [Property reference](resources--dns_proxy--reference.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
