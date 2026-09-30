---
page_title: "proxy_advertisement"
subcategory: ""
description: "proxy_advertisement for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1594, "body_sha256": "sha256:01a6692abb164701e49b4e14dcb3892e480b8fcb44ad0d3395cb6af4f9ce7516", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--proxy_advertisement.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_advertisement for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_advertisement

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- proxy_advertisement

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_custom](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md): complete subsection reference.

- [do_not_advertise](resources--bigip_http_proxy--properties--proxy_advertisement--do_not_advertise.md): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md)
- [proxy_advertisement.do_not_advertise](resources--bigip_http_proxy--properties--proxy_advertisement--do_not_advertise.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
