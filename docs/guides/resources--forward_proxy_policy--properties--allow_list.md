---
page_title: "allow_list"
subcategory: "Security"
description: "allow_list for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2836, "body_sha256": "sha256:615448caf365a761f69f7ebb951bb404e89a25e91b618495f6fedf2450a9772f", "canonical_id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "child_ids": ["xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_allow", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_deny", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:default_action_next_policy", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:dest_list", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:http_list", "xcsh-docs:resources:forward_proxy_policy:properties:allow_list:tls_list"], "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:allow_list", "parent_id": "xcsh-docs:resources:forward_proxy_policy:reference", "path": "docs/guides/resources--forward_proxy_policy--properties--allow_list.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/allow_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_list for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- allow_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
allow_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_action_allow](resources--forward_proxy_policy--properties--allow_list--default_action_allow.md): complete subsection reference.

- [default_action_deny](resources--forward_proxy_policy--properties--allow_list--default_action_deny.md): complete subsection reference.

- [default_action_next_policy](resources--forward_proxy_policy--properties--allow_list--default_action_next_policy.md): complete subsection reference.

- [dest_list](resources--forward_proxy_policy--properties--allow_list--dest_list.md): complete subsection reference.

- [http_list](resources--forward_proxy_policy--properties--allow_list--http_list.md): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--properties--allow_list--tls_list.md): complete subsection reference.

## Next pages

- [allow_list.default_action_allow](resources--forward_proxy_policy--properties--allow_list--default_action_allow.md)
- [allow_list.default_action_deny](resources--forward_proxy_policy--properties--allow_list--default_action_deny.md)
- [allow_list.default_action_next_policy](resources--forward_proxy_policy--properties--allow_list--default_action_next_policy.md)
- [allow_list.dest_list](resources--forward_proxy_policy--properties--allow_list--dest_list.md)
- [allow_list.http_list](resources--forward_proxy_policy--properties--allow_list--http_list.md)
- [allow_list.tls_list](resources--forward_proxy_policy--properties--allow_list--tls_list.md)
- [Property reference](resources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md)
