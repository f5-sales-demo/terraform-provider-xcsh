---
page_title: "any_server"
subcategory: "Security"
description: "any_server for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1347, "body_sha256": "sha256:978c1abb88ea244c13d958efbaf6ff1d86061b7b1890d03466c0cef837820570", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:any_server", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:any_server", "parent_id": "xcsh-docs:resources:rate_limiter_policy:reference", "path": "docs/guides/resources--rate_limiter_policy--properties--any_server.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["any_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/any_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_server for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# any_server

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- any_server

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option

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

- [any_server](resources--rate_limiter_policy--properties--any_server.md#section)
- [server_name](resources--rate_limiter_policy--reference.md#schema-server_name)
- [server_name_matcher](resources--rate_limiter_policy--properties--server_name_matcher.md#section)
- [server_selector](resources--rate_limiter_policy--properties--server_selector.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--rate_limiter_policy--reference.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
