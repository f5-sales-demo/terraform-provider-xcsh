---
page_title: "ddos_profile.enable_ddos_mitigation"
subcategory: ""
description: "ddos_profile.enable_ddos_mitigation for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 986, "body_sha256": "sha256:d2ed7b36923f3c94f33710d40a4b9ce04f1ad65e962923104a064b46fca94581", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "parent_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "path": "docs/guides/resources--dns_proxy--properties--ddos_profile--enable_ddos_mitigation.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ddos_profile.enable_ddos_mitigation for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile.enable_ddos_mitigation

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [ddos_profile](resources--dns_proxy--properties--ddos_profile.md)
- ddos_profile.enable_ddos_mitigation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
enable_ddos_mitigation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ddos_profile](resources--dns_proxy--properties--ddos_profile.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
