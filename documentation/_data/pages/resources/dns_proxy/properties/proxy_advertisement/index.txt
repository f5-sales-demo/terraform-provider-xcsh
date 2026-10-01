---
page_title: "proxy_advertisement"
subcategory: ""
description: "proxy_advertisement for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 7545, "body_sha256": "sha256:68d7c85b5282ebe9a0cfe853c84393ae5ed28d69b18ecf2a2b44f7d1256dbc88", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_dualstack_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_dualstack_vip", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_ipv6_vip", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_on_public_default_vip", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:advertise_v6_on_public", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement:do_not_advertise"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "documentation/resources/dns_proxy/properties/proxy_advertisement/index.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- proxy_advertisement

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_dualstack_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_dualstack_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_dualstack_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_ipv6_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_dualstack_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_on_public_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_ipv6_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "advertise_v6_on_public"),
  validators.ConflictingObjectAttributes("advertise_on_public_default_vip",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_v6_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_dualstack_on_public\",\"advertise_on_public\",\"advertise_on_public_default_dualstack_vip\",\"advertise_on_public_default_ipv6_vip\",\"advertise_on_public_default_vip\",\"advertise_v6_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/): complete subsection reference.

- [advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public/): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_dualstack_vip/): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_ipv6_vip/): complete subsection reference.

- [advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_vip/): complete subsection reference.

- [advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/do_not_advertise/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/)
- [proxy_advertisement.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public/)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_dualstack_vip/)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_ipv6_vip/)
- [proxy_advertisement.advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_vip/)
- [proxy_advertisement.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/)
- [proxy_advertisement.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/proxy_advertisement/do_not_advertise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
