---
page_title: "origin_servers.origin_servers.site_preferences"
subcategory: ""
description: "origin_servers.origin_servers.site_preferences for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1261, "body_sha256": "sha256:0cb579f2a7bdc2075cfbc5b1ca00a3398535abbd953a2aeb13ac56a37bffc7e7", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "path": "docs/guides/resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.origin_servers.site_preferences for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_servers.origin_servers.site_preferences

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- origin_servers.origin_servers.site_preferences

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Carries the references to one or more sites.

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
site_preferences {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.site_preferences.refs](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md)
- [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
