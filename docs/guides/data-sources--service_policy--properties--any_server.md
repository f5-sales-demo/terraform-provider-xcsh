---
page_title: "any_server"
subcategory: "Security"
description: "any_server for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1430, "body_sha256": "sha256:a4069ff73557f05c207802bc391af1607928f780d74dde5fdbd61b1d46260171", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:any_server", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:any_server", "parent_id": "xcsh-docs:data-sources:service_policy:reference", "path": "docs/guides/data-sources--service_policy--properties--any_server.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["any_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/any_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_server for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_server

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- any_server

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [any_server](data-sources--service_policy--properties--any_server.md#section)
- [server_name](data-sources--service_policy--reference.md#schema-server_name)
- [server_name_matcher](data-sources--service_policy--properties--server_name_matcher.md#section)
- [server_selector](data-sources--service_policy--properties--server_selector.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--service_policy--reference.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
