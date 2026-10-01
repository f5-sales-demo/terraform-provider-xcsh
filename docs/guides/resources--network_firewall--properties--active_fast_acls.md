---
page_title: "active_fast_acls"
subcategory: "Security"
description: "active_fast_acls for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1624, "body_sha256": "sha256:f51f9a81997e9d12fea2a957f18c4005ed24dab4275cc86829a690e21edfb66a", "canonical_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--active_fast_acls.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_fast_acls for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_fast_acls

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- active_fast_acls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

OneOf alternatives in this subsection:

- [active_fast_acls](resources--network_firewall--properties--active_fast_acls.md#section)
- [disable_fast_acl](resources--network_firewall--properties--disable_fast_acl.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fast_acls](resources--network_firewall--properties--active_fast_acls--fast_acls.md): complete subsection reference.

## Next pages

- [active_fast_acls.fast_acls](resources--network_firewall--properties--active_fast_acls--fast_acls.md)
- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
