---
page_title: "allow_list"
subcategory: "Security"
description: "allow_list for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2476, "body_sha256": "sha256:ec5f0223ffb59268e756877c7605b9660a25a51b083078cfaf793b2b9b36fc83", "canonical_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_allow", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_deny", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:dest_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:http_list", "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list:tls_list"], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "docs/guides/data-sources--forward_proxy_policy--properties--allow_list.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/allow_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_list for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- allow_list

<a id="section"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

## Direct properties

- [default_action_allow](data-sources--forward_proxy_policy--properties--allow_list--default_action_allow.md): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--properties--allow_list--default_action_deny.md): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--properties--allow_list--default_action_next_policy.md): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--properties--allow_list--dest_list.md): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--properties--allow_list--http_list.md): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--properties--allow_list--tls_list.md): complete subsection reference.

## Next pages

- [allow_list.default_action_allow](data-sources--forward_proxy_policy--properties--allow_list--default_action_allow.md)
- [allow_list.default_action_deny](data-sources--forward_proxy_policy--properties--allow_list--default_action_deny.md)
- [allow_list.default_action_next_policy](data-sources--forward_proxy_policy--properties--allow_list--default_action_next_policy.md)
- [allow_list.dest_list](data-sources--forward_proxy_policy--properties--allow_list--dest_list.md)
- [allow_list.http_list](data-sources--forward_proxy_policy--properties--allow_list--http_list.md)
- [allow_list.tls_list](data-sources--forward_proxy_policy--properties--allow_list--tls_list.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
