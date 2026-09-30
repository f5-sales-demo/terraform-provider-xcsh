---
page_title: "no_interception"
subcategory: ""
description: "no_interception for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1054, "body_sha256": "sha256:7abd833af1c85214a8ba447ad830f99ec6375537ac5e62d682205590cf7a80c1", "canonical_id": "xcsh-docs:data-sources:proxy:properties:no_interception", "child_ids": [], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:no_interception", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "docs/guides/data-sources--proxy--properties--no_interception.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_interception"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/no_interception/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_interception for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# no_interception

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- no_interception

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_interception, tls\_intercept; Default: no\_interception\] Configuration parameter for
no interception.

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

OneOf alternatives in this subsection:

- [no_interception](data-sources--proxy--properties--no_interception.md#section)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--proxy--reference.md)
- [xcsh_proxy](../data-sources/proxy.md)
